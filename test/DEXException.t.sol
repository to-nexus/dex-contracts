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

    // [MARKET] A registered BASE token cannot be registered again.
    function test_exception_market_case1() external {
        vm.prank(OWNER);

        vm.expectRevert(abi.encodeWithSignature("MarketAlreadyCreatedBaseAddress(address)", address(BASE)));
        MARKET.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e4, FEE_PERMIL);
    }

    // [MARKET] BASE cannot be registered with the same address as QUOTE.
    function test_exception_market_case2() external {
        vm.prank(OWNER);

        vm.expectRevert(abi.encodeWithSignature("MarketInvalidBaseAddress(address)", address(QUOTE)));
        MARKET.createPair(address(QUOTE), QUOTE_DECIMALS / 1e2, QUOTE_DECIMALS / 1e4, FEE_PERMIL);
    }

    // [Pair] Trades cannot be executed in units smaller than the Tick Size.
    function test_exception_pair_case1() external {
        (uint256 quoteTickSize, uint256 baseTickSize) = (PAIR.quoteTickSize(), PAIR.baseTickSize());

        uint256 invalidPrice = (quoteTickSize * 11) / 10;
        uint256 invalidAmount = (baseTickSize * 11) / 10;

        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        // [limit] check success.
        address pair = address(PAIR);
        ROUTER.limitSell(
            pair, quoteTickSize * 2, baseTickSize, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.limitBuy(pair, quoteTickSize, baseTickSize, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // [limit] check fail.
        vm.expectRevert(abi.encodeWithSignature("PairInvalidPrice(uint256)", invalidPrice));
        ROUTER.limitSell(pair, invalidPrice, baseTickSize, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // [limit] check fail.
        vm.expectRevert(abi.encodeWithSignature("PairInvalidAmount(uint256)", invalidAmount));
        ROUTER.limitSell(pair, quoteTickSize, invalidAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // [market] check success.
        uint256 denominator = PAIR.DENOMINATOR();
        uint256 volume = Math.mulDiv(quoteTickSize, baseTickSize, denominator);
        ROUTER.marketSell(pair, baseTickSize, 0);
        ROUTER.marketBuy(pair, volume, 0);

        // [market] check fail.
        vm.expectRevert(abi.encodeWithSignature("PairInvalidAmount(uint256)", invalidAmount));
        ROUTER.marketSell(pair, invalidAmount, 0);

        // [market] check fail.
        vm.expectRevert(abi.encodeWithSignature("PairInsufficientTradeVolume(uint256,uint256)", volume - 1, volume));
        ROUTER.marketBuy(pair, volume - 1, 0);
    }

    // [Pair] Users cannot directly request trades through the Pair.
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
            _searchPrices,
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
            _searchPrices,
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
