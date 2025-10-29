// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import "./DEXBase.t.sol";
import {Address} from "@openzeppelin-contracts-5.2.0/utils/Address.sol";

contract DEXExceptionTest is DEXBaseTest {
    using Address for address payable;

    function setUp() external {
        MAX_MATCH_COUNT = 5;
        FEE_BPS = 50;

        _deploy(6, 18, 1e2, 1e4);
    }

    // [MARKET] A registered BASE token cannot be registered again.
    function test_exception_market_case1() external {
        vm.prank(OWNER);

        vm.expectRevert(abi.encodeWithSignature("MarketAlreadyCreatedBaseAddress(address)", address(BASE)));
        MARKET.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e4);
    }

    // [MARKET] BASE cannot be registered with the same address as QUOTE.
    function test_exception_market_case2() external {
        vm.prank(OWNER);

        vm.expectRevert(abi.encodeWithSignature("MarketInvalidBaseAddress(address)", address(QUOTE)));
        MARKET.createPair(address(QUOTE), QUOTE_DECIMALS / 1e2, QUOTE_DECIMALS / 1e4);
    }

    // [Pair] Trades cannot be executed in units smaller than the Tick Size.
    function test_exception_pair_case1() external {
        (uint256 quoteTickSize, uint256 baseTickSize) = (PAIR.tickSize(), PAIR.lotSize());

        uint256 invalidPrice = (quoteTickSize * 11) / 10;
        uint256 invalidAmount = (baseTickSize * 11) / 10;

        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        // [limit] check success.
        address pair = address(PAIR);
        ROUTER.submitSellLimit(
            pair, quoteTickSize * 2, baseTickSize, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.submitBuyLimit(
            pair, quoteTickSize, baseTickSize, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // [limit] check fail.
        vm.expectRevert(abi.encodeWithSignature("PairInvalidPrice(uint256)", invalidPrice));
        ROUTER.submitSellLimit(
            pair, invalidPrice, baseTickSize, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // [limit] check fail.
        vm.expectRevert(abi.encodeWithSignature("PairInvalidAmount(uint256)", invalidAmount));
        ROUTER.submitSellLimit(
            pair, quoteTickSize, invalidAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // [market] check success.
        uint256 denominator = PAIR.DENOMINATOR();
        uint256 volume = Math.mulDiv(quoteTickSize, baseTickSize, denominator);
        ROUTER.submitSellMarket(pair, baseTickSize, 0);
        ROUTER.submitBuyMarket(pair, volume, 0);

        // [market] check fail.
        vm.expectRevert(abi.encodeWithSignature("PairInvalidAmount(uint256)", invalidAmount));
        ROUTER.submitSellMarket(pair, invalidAmount, 0);

        // [market] check fail.
        vm.expectRevert(abi.encodeWithSignature("PairInsufficientTradeVolume(uint256,uint256)", volume - 1, volume));
        ROUTER.submitBuyMarket(pair, volume - 1, 0);
    }

    // [Pair] Users cannot directly request trades through the Pair.
    function test_exception_pair_case2() external {
        vm.startPrank(OWNER);
        BASE.approve(address(PAIR), type(uint256).max);
        QUOTE.approve(address(PAIR), type(uint256).max);

        // sell limit
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.submitLimitOrder(
            IPair.Order({
                side: IPair.OrderSide.SELL,
                owner: address(OWNER),
                feeBps: 0,
                price: _toQuote(1),
                amount: _toBase(1)
            }),
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            0,
            0
        );

        // buy limit
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.submitLimitOrder(
            IPair.Order({
                side: IPair.OrderSide.BUY,
                owner: address(OWNER),
                feeBps: 0,
                price: _toQuote(1),
                amount: _toBase(1)
            }),
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            0,
            0
        );

        // sell market
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.submitMarketOrder(
            IPair.Order({side: IPair.OrderSide.SELL, owner: address(OWNER), feeBps: 0, price: 0, amount: 0}),
            _toBase(1),
            0
        );

        // buy market
        vm.expectRevert(abi.encodeWithSignature("PairInvalidRouter(address)", address(OWNER)));
        PAIR.submitMarketOrder(
            IPair.Order({side: IPair.OrderSide.BUY, owner: address(OWNER), feeBps: 0, price: 0, amount: 0}),
            _toQuote(1),
            0
        );
    }

    // If the match count is insufficient match count, the order will be linked.
    // order books with the same price can appear on both sides.
    function test_exception_duplicate_orderbook_case1() external {
        vm.startPrank(OWNER);
        MARKET.setFeeBps(0); // for simple calculation
        BASE.approve(address(ROUTER), type(uint256).max);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(1);
        uint256 volume = Math.mulDiv(price, amount, PAIR.DENOMINATOR());
        // [limit] check success.
        address pair = address(PAIR);
        ROUTER.submitSellLimit(pair, price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        ROUTER.submitSellLimit(pair, price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        {
            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length);
            assertEq(0, buyPrices.length);
        }

        address buyer = address(0x1);
        QUOTE.transfer(buyer, volume * 2);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(pair, price, amount * 2, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 1);
        {
            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length); // remain other order
            assertEq(0, buyPrices.length, "todo fix");
        }

        assertEq(volume, QUOTE.balanceOf(buyer)); // check return balance
        assertEq(amount, BASE.balanceOf(buyer)); // check match balance
    }

    function test_exception_duplicate_orderbook_case2() external {
        vm.startPrank(OWNER);
        MARKET.setFeeBps(0); // for simple calculation
        QUOTE.approve(address(ROUTER), type(uint256).max);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(1);
        uint256 volume = Math.mulDiv(price, amount, PAIR.DENOMINATOR());

        // [limit] check success.
        address pair = address(PAIR);
        ROUTER.submitBuyLimit(pair, price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        ROUTER.submitBuyLimit(pair, price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        {
            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(1, buyPrices.length);
        }

        address seller = address(0x1);
        BASE.transfer(seller, amount * 2);
        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(pair, price, amount * 2, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 1);
        {
            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length, "todo fix");
            assertEq(1, buyPrices.length); // remain other order
        }
        assertEq(amount, BASE.balanceOf(seller)); // check return balance
        assertEq(volume, QUOTE.balanceOf(seller)); // check match balance
    }
}
