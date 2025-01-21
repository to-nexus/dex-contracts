// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Address} from "@openzeppelin-contracts-5.2.0/utils/Address.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {MarketImpl} from "../src/MarketImpl.sol";
import {PairImpl} from "../src/PairImpl.sol";
import {RouterImpl} from "../src/RouterImpl.sol";
import {WETH} from "../src/WETH.sol";
import {IPair} from "../src/interfaces/IPair.sol";

import {T20} from "./mock/T20.sol";

contract DexWrapBaseTest is Test {
    using Address for address payable;

    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));

    WETH public Weth;
    IERC20 public BASE;
    IERC20 public QUOTE;

    uint256 public BASE_DECIMALS;
    uint256 public QUOTE_DECIMALS;

    RouterImpl public ROUTER;
    PairImpl public PAIR;

    function setUp() external {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");
    }

    // [WRAP BASE] create limit sell -> buy market
    function test_wrap_trade_base_case1() external wrapBase {
        address seller = address(0x1);
        address buyer = address(0x2);
        vm.label(seller, "seller");
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        vm.deal(seller, amount);
        vm.prank(seller);
        ROUTER.limitSell{value: amount}(address(PAIR), price, amount, 0, 0);

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, Weth.balanceOf(seller));
        vm.assertEq(0, Weth.balanceOf(buyer));
        vm.assertEq(0, Weth.balanceOf(address(ROUTER)));
        vm.assertEq(amount, Weth.balanceOf(address(PAIR)));

        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);
        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketBuy(address(PAIR), volume, 0);

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(amount, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, Weth.balanceOf(seller));
        vm.assertEq(0, Weth.balanceOf(buyer));
        vm.assertEq(0, Weth.balanceOf(address(ROUTER)));
        vm.assertEq(0, Weth.balanceOf(address(PAIR)));
    }

    // [WRAP BASE] create limit buy -> sell market
    function test_wrap_trade_base_case2() external wrapBase {
        address seller = address(0x1);
        address buyer = address(0x2);
        vm.label(seller, "seller");
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);
        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
        vm.stopPrank();

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, Weth.balanceOf(seller));
        vm.assertEq(0, Weth.balanceOf(buyer));
        vm.assertEq(0, Weth.balanceOf(address(ROUTER)));
        vm.assertEq(0, Weth.balanceOf(address(PAIR)));

        vm.deal(seller, amount);
        vm.prank(seller);
        ROUTER.marketSell{value: amount}(address(PAIR), amount, 0);

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(amount, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, Weth.balanceOf(seller));
        vm.assertEq(0, Weth.balanceOf(buyer));
        vm.assertEq(0, Weth.balanceOf(address(ROUTER)));
        vm.assertEq(0, Weth.balanceOf(address(PAIR)));
    }

    // [WRAP QUOTE] create limit sell -> buy market
    function test_wrap_trade_quote_case1() external wrapQuote {
        address seller = address(0x1);
        address buyer = address(0x2);
        vm.label(seller, "seller");
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
        vm.stopPrank();

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, Weth.balanceOf(seller));
        vm.assertEq(0, Weth.balanceOf(buyer));
        vm.assertEq(0, Weth.balanceOf(address(ROUTER)));
        vm.assertEq(0, Weth.balanceOf(address(PAIR)));

        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.deal(buyer, volume);
        vm.prank(buyer);
        ROUTER.marketBuy{value: volume}(address(PAIR), volume, 0);

        vm.assertEq(volume, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, Weth.balanceOf(seller));
        vm.assertEq(0, Weth.balanceOf(buyer));
        vm.assertEq(0, Weth.balanceOf(address(ROUTER)));
        vm.assertEq(0, Weth.balanceOf(address(PAIR)));
    }

    // [WRAP QUOTE] create limit buy -> sell market
    function test_wrap_trade_quote_case2() external wrapQuote {
        address seller = address(0x1);
        address buyer = address(0x2);
        vm.label(seller, "seller");
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.deal(buyer, volume);
        vm.prank(buyer);
        ROUTER.limitBuy{value: volume}(address(PAIR), price, amount, 0, 0);

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, Weth.balanceOf(seller));
        vm.assertEq(0, Weth.balanceOf(buyer));
        vm.assertEq(0, Weth.balanceOf(address(ROUTER)));
        vm.assertEq(volume, Weth.balanceOf(address(PAIR)));

        vm.prank(OWNER);
        BASE.transfer(seller, amount);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketSell(address(PAIR), amount, 0);
        vm.stopPrank();

        vm.assertEq(volume, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, Weth.balanceOf(seller));
        vm.assertEq(0, Weth.balanceOf(buyer));
        vm.assertEq(0, Weth.balanceOf(address(ROUTER)));
        vm.assertEq(0, Weth.balanceOf(address(PAIR)));
    }

    // EOA 는 토큰 형태로 WETH를 가질 수 없다.
    function test_wrap_check() external wrapBase {
        vm.startPrank(OWNER);
        vm.deal(OWNER, 100);
        uint256 beforeBalance = OWNER.balance;
        assertNotEq(0, beforeBalance);

        payable(address(Weth)).sendValue(10);
        assertEq(0, Weth.balanceOf(OWNER));
        assertEq(beforeBalance, OWNER.balance);
    }

    // msg.value 가 맞지 않으면 에러를 리턴한다.
    function test_wrap_exception_case1() external wrapBase {
        address seller = address(0x1);
        vm.label(seller, "seller");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        vm.deal(seller, amount + 1);
        vm.prank(seller);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        ROUTER.limitSell{value: amount + 1}(address(PAIR), price, amount, 0, 0);

        vm.prank(seller);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        ROUTER.marketSell{value: amount + 1}(address(PAIR), amount, 0);
    }

    // msg.value 가 맞지 않으면 에러를 리턴한다.
    function test_wrap_exception_case2() external wrapQuote {
        address buyer = address(0x1);
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.deal(buyer, volume + 1);
        vm.prank(buyer);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        ROUTER.limitBuy{value: volume + 1}(address(PAIR), price, amount, 0, 0);

        vm.prank(buyer);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        ROUTER.marketBuy{value: volume + 1}(address(PAIR), volume, 0);
    }

    function _toBase(uint256 x) private view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) private view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }

    modifier wrapBase() {
        _setWrapBase();
        _;
    }

    modifier wrapQuote() {
        _setWrapQuote();
        _;
    }

    function _setWrapBase() private {
        vm.startPrank(OWNER);
        QUOTE = IERC20(address(new T20("QUOTE", "Q", 6)));
        QUOTE_DECIMALS = 10 ** IERC20Metadata(address(QUOTE)).decimals();

        address routerImpl = address(new RouterImpl());
        address router = address(new ERC1967Proxy(routerImpl, ""));
        Weth = new WETH("Wrap Cross", "WCross", payable(router));
        ROUTER = RouterImpl(payable(router));
        ROUTER.initialize(payable(address(Weth)), type(uint256).max);

        BASE = IERC20(address(Weth));
        BASE_DECIMALS = 10 ** IERC20Metadata(address(BASE)).decimals();

        address pairImpl = address(new PairImpl());
        address factoryImpl = address(new MarketImpl());
        address factory = address(
            new ERC1967Proxy(
                factoryImpl,
                abi.encodeWithSelector(
                    MarketImpl.initialize.selector, address(ROUTER), FEE_COLLECTOR, address(QUOTE), pairImpl
                )
            )
        );
        MarketImpl F = MarketImpl(factory);

        PAIR = PairImpl(F.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e4, 0, 0));

        ROUTER.addPair(address(PAIR));

        vm.label(address(ROUTER), "ROUTER");
        vm.label(address(PAIR), "PAIR");
        vm.label(address(QUOTE), "QUOTE");
        vm.label(address(BASE), "BASE");

        vm.stopPrank();
    }

    function _setWrapQuote() private {
        vm.startPrank(OWNER);
        // deploy wrap quote
        BASE = IERC20(address(new T20("BASE", "Q", 6)));
        BASE_DECIMALS = 10 ** IERC20Metadata(address(BASE)).decimals();

        address routerImpl = address(new RouterImpl());
        address router = address(new ERC1967Proxy(routerImpl, ""));
        Weth = new WETH("Wrap Cross", "WCross", payable(router));
        ROUTER = RouterImpl(payable(router));
        ROUTER.initialize(payable(address(Weth)), type(uint256).max);

        QUOTE = IERC20(address(Weth));
        QUOTE_DECIMALS = 10 ** IERC20Metadata(address(QUOTE)).decimals();

        address pairImpl = address(new PairImpl());
        address factoryImpl = address(new MarketImpl());
        address factory = address(
            new ERC1967Proxy(
                factoryImpl,
                abi.encodeWithSelector(
                    MarketImpl.initialize.selector, address(ROUTER), FEE_COLLECTOR, address(QUOTE), pairImpl
                )
            )
        );
        MarketImpl F = MarketImpl(factory);

        PAIR = PairImpl(F.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e4, 0, 0));

        ROUTER.addPair(address(PAIR));

        vm.label(address(ROUTER), "ROUTER");
        vm.label(address(PAIR), "PAIR");
        vm.label(address(QUOTE), "QUOTE");
        vm.label(address(BASE), "BASE");

        vm.stopPrank();
    }
}
