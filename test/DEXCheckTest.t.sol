// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.1.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.1.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin-contracts-5.1.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {FactoryImpl} from "../src/FactoryImpl.sol";
import {PairImpl} from "../src/PairImpl.sol";
import {RouterImpl} from "../src/RouterImpl.sol";
import {WETH} from "../src/WETH.sol";
import {IPair} from "../src/interfaces/IPair.sol";

import {T20} from "./mock/T20.sol";

contract DEXCheckTest is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));

    WETH public Weth;
    IERC20 public BASE;
    IERC20 public QUOTE;

    uint256 public BASE_DECIMALS;
    uint256 public QUOTE_DECIMALS;

    RouterImpl public ROUTER;
    PairImpl public PAIR;

    uint256 public MAKER_FEE_PERMIL = 50; // 5%
    uint256 public TAKER_FEE_PERMIL = 10; // 1%

    function setUp() external {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");

        vm.startPrank(OWNER);

        QUOTE = IERC20(address(new T20("QUOTE", "Q", 18)));
        BASE = IERC20(address(new T20("BASE", "B", 18)));
        QUOTE_DECIMALS = 10 ** IERC20Metadata(address(QUOTE)).decimals();
        BASE_DECIMALS = 10 ** IERC20Metadata(address(BASE)).decimals();

        address routerImpl = address(new RouterImpl());
        address router = address(new ERC1967Proxy(routerImpl, ""));
        Weth = new WETH("Wrap Cross", "WCross", payable(router));
        ROUTER = RouterImpl(payable(router));
        ROUTER.initialize(payable(address(Weth)), type(uint256).max);

        address pairImpl = address(new PairImpl());
        address factoryImpl = address(new FactoryImpl());
        address factory = address(
            new ERC1967Proxy(
                factoryImpl,
                abi.encodeWithSelector(
                    FactoryImpl.initialize.selector, address(ROUTER), FEE_COLLECTOR, address(QUOTE), pairImpl
                )
            )
        );
        FactoryImpl F = FactoryImpl(factory);

        PAIR = PairImpl(
            F.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e4, MAKER_FEE_PERMIL, TAKER_FEE_PERMIL)
        );

        ROUTER.addPair(address(PAIR));

        vm.label(address(ROUTER), "ROUTER");
        vm.label(address(PAIR), "PAIR");
        vm.label(address(QUOTE), "QUOTE");
        vm.label(address(BASE), "BASE");

        vm.stopPrank();
    }

    function _toBase(uint256 x) private view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) private view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }

    // [AMOUNT] sell == buy
    function test_check_fee_case1() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        uint256 makerFee = Math.mulDiv(volume, MAKER_FEE_PERMIL, 1000);
        uint256 takerFee = Math.mulDiv(amount, TAKER_FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(makerFee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid maker fee");
        assertEq(takerFee, BASE.balanceOf(FEE_COLLECTOR), "invalid taker fee");

        // 잔액 확인
        assertEq(volume - makerFee, QUOTE.balanceOf(seller));
        assertEq(amount - takerFee, BASE.balanceOf(buyer));
    }

    // [AMOUNT] sell < buy
    function test_check_fee_case2() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 sellAmount = _toBase(700);
        uint256 buyAmount = _toBase(900);
        uint256 buyerVolume = Math.mulDiv(price, buyAmount, BASE_DECIMALS);

        uint256 matchAmount = Math.min(sellAmount, buyAmount);
        uint256 matchVolume = Math.mulDiv(price, matchAmount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, sellAmount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, buyerVolume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price, sellAmount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price, buyAmount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        uint256 makerFee = Math.mulDiv(matchVolume, MAKER_FEE_PERMIL, 1000);
        uint256 takerFee = Math.mulDiv(matchAmount, TAKER_FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(makerFee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid maker fee");
        assertEq(takerFee, BASE.balanceOf(FEE_COLLECTOR), "invalid taker fee");

        // 잔액 확인
        assertEq(matchVolume - makerFee, QUOTE.balanceOf(seller));
        assertEq(matchAmount - takerFee, BASE.balanceOf(buyer));
    }

    // [AMOUNT] sell > buy
    function test_check_fee_case3() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 sellAmount = _toBase(700);
        uint256 buyAmount = _toBase(300);
        uint256 buyerVolume = Math.mulDiv(price, buyAmount, BASE_DECIMALS);

        uint256 matchAmount = Math.min(sellAmount, buyAmount);
        uint256 matchVolume = Math.mulDiv(price, matchAmount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, sellAmount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, buyerVolume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price, sellAmount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price, buyAmount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        uint256 makerFee = Math.mulDiv(matchVolume, MAKER_FEE_PERMIL, 1000);
        uint256 takerFee = Math.mulDiv(matchAmount, TAKER_FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(makerFee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid maker fee");
        assertEq(takerFee, BASE.balanceOf(FEE_COLLECTOR), "invalid taker fee");

        // 잔액 확인
        assertEq(matchVolume - makerFee, QUOTE.balanceOf(seller));
        assertEq(matchAmount - takerFee, BASE.balanceOf(buyer));
    }

    // [AMOUNT] buy == sell
    function test_check_fee_case4() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(7);
        uint256 amount = _toBase(300);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        uint256 makerFee = Math.mulDiv(amount, MAKER_FEE_PERMIL, 1000);
        uint256 takerFee = Math.mulDiv(volume, TAKER_FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(makerFee, BASE.balanceOf(FEE_COLLECTOR), "invalid taker fee");
        assertEq(takerFee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid maker fee");

        // 잔액 확인
        assertEq(volume - takerFee, QUOTE.balanceOf(seller));
        assertEq(amount - makerFee, BASE.balanceOf(buyer));
    }

    // [AMOUNT] buy < sell
    function test_check_fee_case5() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 sellAmount = _toBase(700);
        uint256 buyAmount = _toBase(500);
        uint256 buyerVolume = Math.mulDiv(price, buyAmount, BASE_DECIMALS);

        uint256 matchAmount = Math.min(sellAmount, buyAmount);
        uint256 matchVolume = Math.mulDiv(price, matchAmount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, sellAmount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, buyerVolume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price, buyAmount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price, sellAmount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        uint256 makerFee = Math.mulDiv(matchAmount, MAKER_FEE_PERMIL, 1000);
        uint256 takerFee = Math.mulDiv(matchVolume, TAKER_FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(makerFee, BASE.balanceOf(FEE_COLLECTOR), "invalid maker fee");
        assertEq(takerFee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid taker fee");

        // 잔액 확인
        assertEq(matchVolume - takerFee, QUOTE.balanceOf(seller));
        assertEq(matchAmount - makerFee, BASE.balanceOf(buyer));
    }

    // [AMOUNT] buy > sell
    function test_check_fee_case6() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 sellAmount = _toBase(700);
        uint256 buyAmount = _toBase(900);
        uint256 buyerVolume = Math.mulDiv(price, buyAmount, BASE_DECIMALS);

        uint256 matchAmount = Math.min(sellAmount, buyAmount);
        uint256 matchVolume = Math.mulDiv(price, matchAmount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, sellAmount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, buyerVolume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price, buyAmount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price, sellAmount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        uint256 makerFee = Math.mulDiv(matchAmount, MAKER_FEE_PERMIL, 1000);
        uint256 takerFee = Math.mulDiv(matchVolume, TAKER_FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(makerFee, BASE.balanceOf(FEE_COLLECTOR), "invalid maker fee");
        assertEq(takerFee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid taker fee");

        // 잔액 확인
        assertEq(matchVolume - takerFee, QUOTE.balanceOf(seller));
        assertEq(matchAmount - makerFee, BASE.balanceOf(buyer));
    }

    // [199981] SELL OnePrice -> BUY (1 orders)
    function test_check_gas_case1() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(1, sellPrices.length);
        assertEq(0, buyPrices.length);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        // 199981
        ROUTER.buyMarketOrder(address(PAIR), volume, 0);
        assertEq(0, QUOTE.balanceOf(buyer)); // 정확이 매칭 되었는지 확인

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    // [258882(+58901)] SELL OnePrice -> BUY (2 orders)
    function test_check_gas_case2() external {
        address seller1 = address(0x1);
        address seller2 = address(0x2);
        address buyer = address(0x3);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount * 2, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller1, amount);
        vm.prank(OWNER);
        BASE.transfer(seller2, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        vm.startPrank(seller1);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller1));
        vm.startPrank(seller2);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(1, sellPrices.length);
        assertEq(0, buyPrices.length);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        // 258882
        ROUTER.buyMarketOrder(address(PAIR), volume, 0);
        assertEq(0, QUOTE.balanceOf(buyer)); // 정확이 매칭 되었는지 확인

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    // [266886(+66905(+8004))] SELL TwoPrices -> BUY (2 orders)
    function test_check_gas_case3() external {
        address seller1 = address(0x1);
        address seller2 = address(0x2);
        address buyer = address(0x3);

        uint256 price1 = _toQuote(3);
        uint256 price2 = _toQuote(5);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price1, amount, BASE_DECIMALS) + Math.mulDiv(price2, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller1, amount);
        vm.prank(OWNER);
        BASE.transfer(seller2, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        vm.startPrank(seller1);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price1, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller1));
        vm.startPrank(seller2);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.sellLimitOrder(address(PAIR), price2, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(2, sellPrices.length);
        assertEq(0, buyPrices.length);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        // 266886
        ROUTER.buyMarketOrder(address(PAIR), volume, 0);
        assertEq(0, QUOTE.balanceOf(buyer)); // 정확이 매칭 되었는지 확인

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    // [195076] BUY OnePrice -> SELL (1 orders)
    function test_check_gas_case4() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(1, buyPrices.length);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        // 195076
        ROUTER.sellMarketOrder(address(PAIR), amount, 0);
        assertEq(0, BASE.balanceOf(seller)); // 정확이 매칭 되었는지 확인

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    // [253878(+58802)] BUY OnePrice -> SELL (2 orders)
    function test_check_gas_case5() external {
        address seller = address(0x1);
        address buyer1 = address(0x2);
        address buyer2 = address(0x3);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount * 2);
        vm.prank(OWNER);
        QUOTE.transfer(buyer1, volume);
        vm.prank(OWNER);
        QUOTE.transfer(buyer2, volume);

        vm.startPrank(buyer1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer1));

        vm.startPrank(buyer2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(1, buyPrices.length);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        // 253878
        ROUTER.sellMarketOrder(address(PAIR), amount * 2, 0);
        assertEq(0, BASE.balanceOf(seller)); // 정확이 매칭 되었는지 확인

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    // [261216(+66140(+7338))] BUY TwoPrices -> SELL (2 orders)
    function test_check_gas_case6() external {
        address seller = address(0x1);
        address buyer1 = address(0x2);
        address buyer2 = address(0x3);

        uint256 price1 = _toQuote(3);
        uint256 price2 = _toQuote(1);
        uint256 amount = _toBase(700);
        uint256 volume1 = Math.mulDiv(price1, amount, BASE_DECIMALS);
        uint256 volume2 = Math.mulDiv(price2, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount * 2);
        vm.prank(OWNER);
        QUOTE.transfer(buyer1, volume1);
        vm.prank(OWNER);
        QUOTE.transfer(buyer2, volume2);

        vm.startPrank(buyer1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price1, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer1));

        vm.startPrank(buyer2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyLimitOrder(address(PAIR), price2, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(2, buyPrices.length);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        // 261216
        ROUTER.sellMarketOrder(address(PAIR), amount * 2, 0);
        assertEq(0, BASE.balanceOf(seller)); // 정확이 매칭 되었는지 확인

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }
}
