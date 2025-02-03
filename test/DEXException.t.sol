// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./DEXBase.t.sol";
import {Address} from "@openzeppelin-contracts-5.2.0/utils/Address.sol";

contract DEXExceptionTest is DEXBaseTest {
    using Address for address payable;

    function setUp() external {
        MAX_MATCH_COUNT = 5;
        FEE_PERMIL = 50;

        _deploy(6, 18, 1e2, 1e4);
    }

    // // [ROUTER] 제거된 Pair 는 거래 할 수 없다.
    // function test_exception_router_case1() external {
    //     vm.startPrank(OWNER);
    //     address pair = address(PAIR);
    //     ROUTER.removePair(pair); // 등록된 PAIR 제거

    //     BASE.approve(address(ROUTER), type(uint256).max);
    //     QUOTE.approve(address(ROUTER), type(uint256).max);

    //     // 거래 및 실패 확인
    //     vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
    //     ROUTER.limitSell(pair, 1e18, 1e18, 0, 0);

    //     vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
    //     ROUTER.limitBuy(pair, 1e18, 1e18, 0, 0);

    //     vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
    //     ROUTER.marketSell(pair, 1e18, 0);

    //     vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", pair));
    //     ROUTER.marketBuy(pair, 1e18, 0);

    //     vm.stopPrank();
    // }

    // [ROUTER] coin 을 전송 받을 수 없다.
    function test_exception_wcross_case1() external {
        vm.startPrank(OWNER);
        vm.deal(OWNER, 1);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        payable(address(ROUTER)).sendValue(1);
    }

    // [MARKET] 등록된 BASE 토큰은 중복해서 등록할 수 없다.
    function test_exception_market_case1() external {
        vm.prank(OWNER);

        vm.expectRevert(abi.encodeWithSignature("MarketAlreadyCreatedBaseAddress(address)", address(BASE)));
        MARKET.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e4, FEE_PERMIL);
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

    // // [Pair] 다른 ROUTER 으로 거래를 요청할 수 없다.
    // function test_exception_pair_case3() external {
    //     address pair = address(PAIR);

    //     vm.startPrank(OWNER);
    //     // ROUTER 로는 거래할 수 있다.
    //     BASE.approve(address(ROUTER), type(uint256).max);
    //     QUOTE.approve(address(ROUTER), type(uint256).max);
    //     ROUTER.limitSell(pair, _toQuote(1), _toBase(1), 0, 0);
    //     ROUTER.limitBuy(pair, _toQuote(1), _toBase(1), 0, 0);
    //     ROUTER.marketSell(pair, _toBase(1), 0);
    //     ROUTER.marketBuy(pair, _toQuote(1), 0);

    //     // R 배포
    //     address routerImpl = address(new RouterImpl());
    //     address router = address(new ERC1967Proxy(routerImpl, ""));
    //     RouterImpl R = RouterImpl(payable(router));
    //     R.initialize(payable(address(WCross)), type(uint256).max);

    //     // R 에 Pair 등록
    //     R.addPair(pair);

    //     BASE.approve(router, type(uint256).max);
    //     QUOTE.approve(router, type(uint256).max);

    //     // R 로 주문은 넣을수 없다.
    //     vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", router));
    //     R.limitSell(pair, _toQuote(1), _toBase(1), 0, 0);
    //     vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", router));
    //     R.limitBuy(pair, _toQuote(1), _toBase(1), 0, 0);
    //     vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", router));
    //     R.marketSell(pair, _toBase(1), 0);
    //     vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", router));
    //     R.marketBuy(pair, _toQuote(1), 0);

    //     vm.stopPrank();
    // }
}
