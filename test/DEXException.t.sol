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

    // [MARKET] QUOTE 와 같은 주소로 BASE를 등록할 수 없다.
    function test_exception_market_case2() external {
        vm.prank(OWNER);

        vm.expectRevert(abi.encodeWithSignature("MarketInvalidBaseAddress(address)", address(QUOTE)));
        MARKET.createPair(address(QUOTE), QUOTE_DECIMALS / 1e2, QUOTE_DECIMALS / 1e4, FEE_PERMIL);
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
        ROUTER.limitSell(pair, quoteTickSize * 2, baseTickSize, IPair.LimitConstraints.GOOD_TILL_CANCEL, 0, 0);
        ROUTER.limitBuy(pair, quoteTickSize, baseTickSize, IPair.LimitConstraints.GOOD_TILL_CANCEL, 0, 0);

        // [limit] 실패 확인
        vm.expectRevert(abi.encodeWithSignature("PairInvalidPrice(uint256)", invalidPrice));
        ROUTER.limitSell(pair, invalidPrice, baseTickSize, IPair.LimitConstraints.GOOD_TILL_CANCEL, 0, 0);

        // [limit] 실패 확인
        vm.expectRevert(abi.encodeWithSignature("PairInvalidAmount(uint256)", invalidAmount));
        ROUTER.limitSell(pair, quoteTickSize, invalidAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, 0, 0);

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
                side: IPair.OrderSide.SELL,
                owner: address(OWNER),
                feePermil: 0,
                price: _toQuote(1),
                amount: _toBase(1)
            }),
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            0,
            0
        );

        // buy limit
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.limit(
            IPair.Order({
                side: IPair.OrderSide.BUY,
                owner: address(OWNER),
                feePermil: 0,
                price: _toQuote(1),
                amount: _toBase(1)
            }),
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            0,
            0
        );

        // sell market
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.market(
            IPair.Order({side: IPair.OrderSide.SELL, owner: address(OWNER), feePermil: 0, price: 0, amount: 0}),
            _toBase(1),
            0
        );

        // buy market
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.market(
            IPair.Order({side: IPair.OrderSide.BUY, owner: address(OWNER), feePermil: 0, price: 0, amount: 0}),
            _toQuote(1),
            0
        );
    }
}
