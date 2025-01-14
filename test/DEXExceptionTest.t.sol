// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.1.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.1.0/token/ERC20/extensions/IERC20Metadata.sol";

import {Address} from "@openzeppelin-contracts-5.1.0/utils/Address.sol";
import {Math} from "@openzeppelin-contracts-5.1.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {FactoryImpl} from "../src/FactoryImpl.sol";
import {PairImpl} from "../src/PairImpl.sol";
import {RouterImpl} from "../src/RouterImpl.sol";
import {WETH} from "../src/WETH.sol";
import {IPair} from "../src/interfaces/IPair.sol";

import {T20} from "./mock/T20.sol";

contract DEXExceptionTest is Test {
    using Address for address payable;

    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));

    uint256 public constant MAX_MATCH_COUNT = 5;
    uint256 public constant MAKER_FEE_PERMIL = 50; // 5%
    uint256 public constant TAKER_FEE_PERMIL = 20; // 2%

    WETH public Weth;
    IERC20 public BASE;
    IERC20 public QUOTE;
    uint256 public BASE_DECIMALS;
    uint256 public QUOTE_DECIMALS;

    RouterImpl public ROUTER;
    FactoryImpl public FACTORY;
    PairImpl public PAIR;

    function setUp() external {
        vm.startPrank(OWNER);

        BASE = IERC20(address(new T20("BASE", "B", 18)));
        QUOTE = IERC20(address(new T20("QUOTE", "Q", 6)));
        BASE_DECIMALS = 10 ** IERC20Metadata(address(BASE)).decimals();
        QUOTE_DECIMALS = 10 ** IERC20Metadata(address(QUOTE)).decimals();

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
                abi.encodeWithSelector(FactoryImpl.initialize.selector, router, FEE_COLLECTOR, address(QUOTE), pairImpl)
            )
        );
        FACTORY = FactoryImpl(factory);

        address pair = FACTORY.createPair(
            address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e4, MAKER_FEE_PERMIL, TAKER_FEE_PERMIL
        );
        ROUTER.addPair(pair);
        PAIR = PairImpl(pair);

        vm.stopPrank();

        vm.label(OWNER, "OWNER");
        vm.label(FEE_COLLECTOR, "FEE_COLLECTOR");
        vm.label(address(BASE), "BASE");
        vm.label(address(QUOTE), "QUOTE");
        vm.label(address(ROUTER), "ROUTER");
        vm.label(address(FACTORY), "FACTORY");
        vm.label(address(PAIR), "PAIR");
    }

    function _toBase(uint256 x) private view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) private view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }

    // [ROUTER] 한번 등록된 Pair 는 다시 등록 할 수 없다.
    function test_exception_router_case1() external {
        address pair = address(PAIR);
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("RouterAlreadyAddedPair(address)", pair));
        ROUTER.addPair(pair);
    }

    // [ROUTER] 등록되지 않은 Pair 는 거래 할 수 없다.
    function test_exception_router_case2() external {
        vm.startPrank(OWNER);

        IERC20 BASE2 = IERC20(address(new T20("BASE2", "B2", 18)));

        address pair =
            FACTORY.createPair(address(BASE2), QUOTE_DECIMALS / 1e2, 1e18 / 1e4, MAKER_FEE_PERMIL, TAKER_FEE_PERMIL);

        BASE2.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        // 거래 및 실패 확인
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
        ROUTER.limitSell(pair, 1e18, 1e18, 0, 0);

        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
        ROUTER.limitBuy(pair, 1e18, 1e18, 0, 0);

        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
        ROUTER.marketSell(pair, 1e18, 0);

        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
        ROUTER.marketBuy(pair, 1e18, 0);

        vm.stopPrank();
    }

    // [ROUTER] 제거된 Pair 는 거래 할 수 없다.
    function test_exception_router_case3() external {
        vm.startPrank(OWNER);
        address pair = address(PAIR);
        ROUTER.removePair(pair); // 등록된 PAIR 제거

        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        // 거래 및 실패 확인
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
        ROUTER.limitSell(pair, 1e18, 1e18, 0, 0);

        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
        ROUTER.limitBuy(pair, 1e18, 1e18, 0, 0);

        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
        ROUTER.marketSell(pair, 1e18, 0);

        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
        ROUTER.marketBuy(pair, 1e18, 0);

        vm.stopPrank();
    }

    // [ROUTER] coin 을 전송 받을 수 없다.
    function test_exception_weth_case1() external {
        vm.startPrank(OWNER);
        vm.deal(OWNER, 1);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        payable(address(ROUTER)).sendValue(1);
    }

    // [Factory] 등록된 BASE 토큰은 중복해서 등록할 수 없다.
    function test_exception_factory_case1() external {
        vm.prank(OWNER);

        vm.expectRevert(abi.encodeWithSignature("FactoryAlreadyCreatedBaseAddress(address)", address(BASE)));
        FACTORY.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e4, MAKER_FEE_PERMIL, TAKER_FEE_PERMIL);
    }

    // [Pair] Tick Size 보다 낮은 단위로 거래할 수 없다.
    function test_exception_pair_case1() external {
        (uint256 quoteTickSize, uint256 baseTickSize) = (PAIR.quoteTickSize(), PAIR.baseTickSize());

        uint256 invalidPrice = (quoteTickSize * 11) / 10;
        uint256 invalidAmount = (baseTickSize * 11) / 10;

        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        // [limit] 성공 확인
        address pair = address(PAIR);
        ROUTER.limitSell(pair, quoteTickSize * 2, baseTickSize, 0, 0);
        ROUTER.limitBuy(pair, quoteTickSize, baseTickSize, 0, 0);

        // [limit] 실패 확인
        vm.expectRevert(abi.encodeWithSignature("PairInvalidPrice(uint256)", invalidPrice));
        ROUTER.limitSell(pair, invalidPrice, baseTickSize, 0, 0);

        // [limit] 실패 확인
        vm.expectRevert(abi.encodeWithSignature("PairInvalidAmount(uint256)", invalidAmount));
        ROUTER.limitSell(pair, quoteTickSize, invalidAmount, 0, 0);

        // [market] 성공 확인
        uint256 denominator = PAIR.DENOMINATOR();
        uint256 volume = Math.mulDiv(quoteTickSize, baseTickSize, denominator);
        ROUTER.marketSell(pair, baseTickSize, 0);
        ROUTER.marketBuy(pair, volume, 0);

        // [market] 실패 확인
        vm.expectRevert(abi.encodeWithSignature("PairInvalidAmount(uint256)", invalidAmount));
        ROUTER.marketSell(pair, invalidAmount, 0);

        // [market] 실패 확인
        vm.expectRevert(abi.encodeWithSignature("PairInsufficientTradeVolume(uint256,uint256)", volume - 1, volume));
        ROUTER.marketBuy(pair, volume - 1, 0);
    }

    // [Pair] 사용자는 직접 Pair 으로 거래를 요청할 수 없다.
    function test_exception_pair_case2() external {
        vm.startPrank(OWNER);
        BASE.approve(address(PAIR), type(uint256).max);
        QUOTE.approve(address(PAIR), type(uint256).max);

        // sell limit
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.limit(
            IPair.Order({
                _type: IPair.OrderType.SELL,
                owner: address(OWNER),
                feePermil: 0,
                price: _toQuote(1),
                amount: _toBase(1)
            }),
            0,
            0
        );

        // buy limit
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.limit(
            IPair.Order({
                _type: IPair.OrderType.BUY,
                owner: address(OWNER),
                feePermil: 0,
                price: _toQuote(1),
                amount: _toBase(1)
            }),
            0,
            0
        );

        // sell market
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.market(
            IPair.Order({_type: IPair.OrderType.SELL, owner: address(OWNER), feePermil: 0, price: 0, amount: 0}),
            _toBase(1),
            0
        );

        // buy market
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.market(
            IPair.Order({_type: IPair.OrderType.BUY, owner: address(OWNER), feePermil: 0, price: 0, amount: 0}),
            _toQuote(1),
            0
        );
    }

    // [Pair] 다른 ROUTER 으로 거래를 요청할 수 없다.
    function test_exception_pair_case3() external {
        address pair = address(PAIR);

        vm.startPrank(OWNER);
        // ROUTER 로는 거래할 수 있다.
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitSell(pair, _toQuote(1), _toBase(1), 0, 0);
        ROUTER.limitBuy(pair, _toQuote(1), _toBase(1), 0, 0);
        ROUTER.marketSell(pair, _toBase(1), 0);
        ROUTER.marketBuy(pair, _toQuote(1), 0);

        // R 배포
        address routerImpl = address(new RouterImpl());
        address router = address(new ERC1967Proxy(routerImpl, ""));
        RouterImpl R = RouterImpl(payable(router));
        R.initialize(payable(address(Weth)), type(uint256).max);

        // R 에 Pair 등록
        R.addPair(pair);

        BASE.approve(router, type(uint256).max);
        QUOTE.approve(router, type(uint256).max);

        // R 로 주문은 넣을수 없다.
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", router));
        R.limitSell(pair, _toQuote(1), _toBase(1), 0, 0);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", router));
        R.limitBuy(pair, _toQuote(1), _toBase(1), 0, 0);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", router));
        R.marketSell(pair, _toBase(1), 0);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", router));
        R.marketBuy(pair, _toQuote(1), 0);

        vm.stopPrank();
    }
}
