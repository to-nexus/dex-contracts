// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {PairImplV3} from "../src/PairImplV3.sol";
import {IPairV3} from "../src/interfaces/IPairV3.sol";
import {DEXV3BaseTest} from "./DEXV3Base.t.sol";

contract PairImplV3Test is DEXV3BaseTest {
    address public constant USER3 = address(bytes20("USER3"));
    address public constant USER4 = address(bytes20("USER4"));

    function setUp() external {
        _deployV3(18, 18, 1e2, 1e6);

        vm.startPrank(OWNER);
        vm.label(USER3, "user3");
        vm.label(USER4, "user4");
        QUOTE.transfer(USER3, _toQuote(50000));
        QUOTE.transfer(USER4, _toQuote(50000));
        BASE.transfer(USER3, _toBase(50000));
        BASE.transfer(USER4, _toBase(50000));
        vm.stopPrank();

        vm.prank(USER3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER3);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER4);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER4);
        BASE.approve(address(ROUTER), type(uint256).max);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // View Functions Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_getConfig() external view {
        IPairV3.Config memory config = PAIR.getConfig();
        assertEq(address(config.QUOTE), address(QUOTE));
        assertEq(address(config.BASE), address(BASE));
        assertEq(config.DENOMINATOR, BASE_DECIMALS);
    }

    function test_orderById_nonExistent() external view {
        IPairV3.Order memory order = PAIR.orderById(999);
        assertEq(order.owner, address(0));
        assertEq(order.amount, 0);
    }

    function test_accountReserves_initial() external view {
        (uint256 baseAmount, uint256 quoteAmount) = PAIR.accountReserves(USER1);
        assertEq(baseAmount, 0);
        assertEq(quoteAmount, 0);
    }

    function test_ticks_empty() external view {
        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(sellPrices.length, 0);
        assertEq(buyPrices.length, 0);
    }

    function test_ticks_withOrders() external {
        uint256 buyPrice = _toQuote(100);
        uint256 sellPrice = _toQuote(110);
        uint256 amount = _toBase(10);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), buyPrice, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), sellPrice, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(sellPrices.length, 1);
        assertEq(buyPrices.length, 1);
        assertEq(sellPrices[0], sellPrice);
        assertEq(buyPrices[0], buyPrice);
    }

    function test_tickSizes() external view {
        (uint256 tick, uint256 lot) = PAIR.tickSizes();
        assertEq(tick, QUOTE_DECIMALS / 1e2); // tickSize
        assertEq(lot, BASE_DECIMALS / 1e6); // lotSize
    }

    function test_ordersByPrices() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        vm.prank(USER1);
        uint256 orderId1 = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER2);
        uint256 orderId2 = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256[] memory prices = new uint256[](1);
        prices[0] = price;
        uint256[][] memory orderIds = PAIR.ordersByPrices(IPairV3.OrderSide.BUY, prices);

        assertEq(orderIds.length, 1);
        assertEq(orderIds[0].length, 2);
        assertEq(orderIds[0][0], orderId1);
        assertEq(orderIds[0][1], orderId2);
    }

    function test_owner() external view {
        assertEq(PAIR.owner(), OWNER);
    }

    function test_findPrevPrice_sell() external {
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(100), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(200), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 prev = PAIR.findPrevPrice(IPairV3.OrderSide.SELL, _toQuote(150), _searchPrices, 100);
        assertEq(prev, _toQuote(100));
    }

    function test_findPrevPrice_buy() external {
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(200), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(100), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 prev = PAIR.findPrevPrice(IPairV3.OrderSide.BUY, _toQuote(150), _searchPrices, 100);
        assertEq(prev, _toQuote(200));
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Limit Order - GOOD_TILL_CANCEL Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_submitSellLimit_GTC_noMatch_addsToBook() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        vm.prank(USER2);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(uint8(order.side), uint8(IPairV3.OrderSide.SELL));
        assertEq(order.owner, USER2);
        assertEq(order.price, price);
        assertEq(order.amount, amount);

        assertEq(PAIR.baseReserve(), amount);
        (uint256 userBase,) = PAIR.accountReserves(USER2);
        assertEq(userBase, amount);
    }

    function test_submitSellLimit_GTC_partialMatch() external {
        uint256 price = _toQuote(100);

        // USER1: BUY 5 BASE
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER2: SELL 10 BASE -> matches 5, 5 remains on book
        vm.prank(USER2);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.amount, _toBase(5));
        assertEq(PAIR.baseReserve(), _toBase(5));
    }

    function test_submitSellLimit_GTC_fullMatch() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // USER1: BUY 10 BASE
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 user2BaseBefore = BASE.balanceOf(USER2);

        // USER2: SELL 10 BASE -> fully matched
        vm.prank(USER2);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Order should be closed
        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
        assertEq(PAIR.baseReserve(), 0);
        assertEq(BASE.balanceOf(USER2), user2BaseBefore - amount);
    }

    function test_submitBuyLimit_GTC_noMatch_addsToBook() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(uint8(order.side), uint8(IPairV3.OrderSide.BUY));
        assertEq(order.owner, USER1);
        assertEq(order.price, price);
        assertEq(order.amount, amount);

        assertEq(PAIR.quoteReserve(), quoteVolume);
        (, uint256 userQuote) = PAIR.accountReserves(USER1);
        assertEq(userQuote, quoteVolume);
    }

    function test_submitBuyLimit_GTC_partialMatch() external {
        uint256 price = _toQuote(100);

        // USER2: SELL 5 BASE
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER1: BUY 10 BASE -> matches 5, 5 remains on book
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.amount, _toBase(5));
    }

    function test_submitBuyLimit_GTC_fullMatch() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // USER2: SELL 10 BASE
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 user1BaseBefore = BASE.balanceOf(USER1);

        // USER1: BUY 10 BASE -> fully matched
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
        assertEq(BASE.balanceOf(USER1), user1BaseBefore + amount);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Limit Order - IMMEDIATE_OR_CANCEL Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_submitSellLimit_IOC_noMatch_returnsAll() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        uint256 user2BaseBefore = BASE.balanceOf(USER2);

        // No buy orders, so SELL IOC should return all BASE
        vm.prank(USER2);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.IMMEDIATE_OR_CANCEL, _searchPrices, 0
        );

        // Order closed, all BASE returned
        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
        assertEq(BASE.balanceOf(USER2), user2BaseBefore);
        assertEq(PAIR.baseReserve(), 0);
    }

    function test_submitSellLimit_IOC_partialMatch_returnsRemaining() external {
        uint256 price = _toQuote(100);

        // USER1: BUY 5 BASE
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 user2BaseBefore = BASE.balanceOf(USER2);

        // USER2: SELL 10 BASE IOC -> matches 5, returns 5
        vm.prank(USER2);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.IMMEDIATE_OR_CANCEL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
        assertEq(BASE.balanceOf(USER2), user2BaseBefore - _toBase(5)); // Only 5 was sold
        assertEq(PAIR.baseReserve(), 0);
    }

    function test_submitBuyLimit_IOC_noMatch_returnsAll() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);

        // No sell orders, so BUY IOC should return all QUOTE
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.IMMEDIATE_OR_CANCEL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
        assertEq(QUOTE.balanceOf(USER1), user1QuoteBefore);
        assertEq(PAIR.quoteReserve(), 0);
    }

    function test_submitBuyLimit_IOC_partialMatch_returnsRemaining() external {
        uint256 price = _toQuote(100);

        // USER2: SELL 5 BASE
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);
        uint256 quoteFor5Base = _toTradeVolume(price, _toBase(5));

        // USER1: BUY 10 BASE IOC -> matches 5, returns quote for remaining 5
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.IMMEDIATE_OR_CANCEL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
        assertEq(QUOTE.balanceOf(USER1), user1QuoteBefore - quoteFor5Base);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Limit Order - FILL_OR_KILL Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_submitSellLimit_FOK_fullMatch_succeeds() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // USER1: BUY 10 BASE
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // USER2: SELL 10 BASE FOK -> fully matched
        vm.prank(USER2);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.FILL_OR_KILL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
    }

    function test_submitSellLimit_FOK_partialMatch_reverts() external {
        uint256 price = _toQuote(100);

        // USER1: BUY only 5 BASE
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER2: SELL 10 BASE FOK -> should revert
        vm.prank(USER2);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairFillOrKill.selector, USER2));
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.FILL_OR_KILL, _searchPrices, 0
        );
    }

    function test_submitBuyLimit_FOK_fullMatch_succeeds() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // USER2: SELL 10 BASE
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER1: BUY 10 BASE FOK -> fully matched
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.FILL_OR_KILL, _searchPrices, 0
        );

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
    }

    function test_submitBuyLimit_FOK_partialMatch_reverts() external {
        uint256 price = _toQuote(100);

        // USER2: SELL only 5 BASE
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER1: BUY 10 BASE FOK -> should revert
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairFillOrKill.selector, USER1));
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.FILL_OR_KILL, _searchPrices, 0
        );
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Market Order Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_submitSellMarket_matchesMultipleBuyOrders() external {
        uint256 price1 = _toQuote(100);
        uint256 price2 = _toQuote(105);

        // USER1: BUY 5 BASE @ 100
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price1, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER1: BUY 5 BASE @ 105
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price2, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 user2QuoteBefore = QUOTE.balanceOf(USER2);

        // USER2: SELL 10 BASE -> matches both orders (best price first: 105)
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(10), 0);

        // Price matched at 105 last
        assertEq(PAIR.matchedPrice(), price1);

        // USER2 should receive quote for 5@105 + 5@100 minus fees
        uint256 expectedQuote = _toTradeVolume(price2, _toBase(5)) + _toTradeVolume(price1, _toBase(5));
        uint256 takerFee = _calcFee(expectedQuote, SELLER_TAKER_FEE);
        assertEq(QUOTE.balanceOf(USER2), user2QuoteBefore + expectedQuote - takerFee);
    }

    function test_submitSellMarket_emptyOrderBook_returnsAll() external {
        uint256 amount = _toBase(10);
        uint256 user2BaseBefore = BASE.balanceOf(USER2);

        // No buy orders
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // All BASE returned
        assertEq(BASE.balanceOf(USER2), user2BaseBefore);
    }

    function test_submitBuyMarket_matchesMultipleSellOrders() external {
        uint256 price1 = _toQuote(100);
        uint256 price2 = _toQuote(105);

        // USER2: SELL 5 BASE @ 100
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price1, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER2: SELL 5 BASE @ 105
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price2, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 user1BaseBefore = BASE.balanceOf(USER1);
        uint256 quoteToSpend = _toTradeVolume(price1, _toBase(5)) + _toTradeVolume(price2, _toBase(5));

        // USER1: BUY market -> matches both (best price first: 100)
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteToSpend, 0);

        assertEq(PAIR.matchedPrice(), price2);
        assertEq(BASE.balanceOf(USER1), user1BaseBefore + _toBase(10));
    }

    function test_submitBuyMarket_revert_insufficientVolume() external {
        // minTradeVolume is computed as tickSize * lotSize / DENOMINATOR
        uint256 minVol = PAIR.minTradeVolume();

        // Submit with less than minTradeVolume
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInsufficientTradeVolume.selector, minVol - 1, minVol));
        ROUTER.submitBuyMarket(address(PAIR), minVol - 1, 0);
    }

    function test_submitMarket_updatesMatchedPrice() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // USER1: BUY
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        assertEq(PAIR.matchedPrice(), 0);
        assertEq(PAIR.matchedAt(), 0);

        // USER2: SELL -> matches
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        assertEq(PAIR.matchedPrice(), price);
        assertEq(PAIR.matchedAt(), block.timestamp);
    }

    function test_submitMarket_maxMatchCountLimits() external {
        uint256 price = _toQuote(100);

        // Create 5 buy orders
        for (uint256 i = 0; i < 5; i++) {
            vm.prank(USER1);
            ROUTER.submitBuyLimit(
                address(PAIR), price, _toBase(2), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }

        // SELL with maxMatchCount = 2 -> only matches 2 orders
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(10), 2);

        // 3 orders should remain (matched 2, remaining 3)
        uint256[] memory prices = new uint256[](1);
        prices[0] = price;
        uint256[][] memory orderIds = PAIR.ordersByPrices(IPairV3.OrderSide.BUY, prices);
        assertEq(orderIds[0].length, 3);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Cancel Order Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_cancelOrder_sellOrder_returnsBase() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        uint256 user2BaseBefore = BASE.balanceOf(USER2);

        vm.prank(USER2);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        assertEq(BASE.balanceOf(USER2), user2BaseBefore - amount);
        assertEq(PAIR.baseReserve(), amount);

        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;
        vm.prank(USER2);
        ROUTER.cancelOrder(address(PAIR), orderIds);

        assertEq(BASE.balanceOf(USER2), user2BaseBefore);
        assertEq(PAIR.baseReserve(), 0);
    }

    function test_cancelOrder_buyOrder_returnsQuote() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);

        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        assertEq(QUOTE.balanceOf(USER1), user1QuoteBefore - quoteVolume);

        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;
        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), orderIds);

        assertEq(QUOTE.balanceOf(USER1), user1QuoteBefore);
        assertEq(PAIR.quoteReserve(), 0);
    }

    function test_cancelOrder_multipleOrders() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        vm.prank(USER1);
        uint256 orderId1 = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER1);
        uint256 orderId2 = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256[] memory orderIds = new uint256[](2);
        orderIds[0] = orderId1;
        orderIds[1] = orderId2;

        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), orderIds);

        assertEq(PAIR.quoteReserve(), 0);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Price List Integrity Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_priceList_sellOrdersAscending() external {
        // Create sell orders at various prices
        vm.startPrank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(150), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(100), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(200), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(125), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        vm.stopPrank();

        (uint256[] memory sellPrices,) = PAIR.ticks();
        assertEq(sellPrices.length, 4);
        // Should be ascending: 100, 125, 150, 200
        assertEq(sellPrices[0], _toQuote(100));
        assertEq(sellPrices[1], _toQuote(125));
        assertEq(sellPrices[2], _toQuote(150));
        assertEq(sellPrices[3], _toQuote(200));
    }

    function test_priceList_buyOrdersDescending() external {
        // Create buy orders at various prices
        vm.startPrank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(100), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(150), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(75), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(125), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        vm.stopPrank();

        (, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(buyPrices.length, 4);
        // Should be descending: 150, 125, 100, 75
        assertEq(buyPrices[0], _toQuote(150));
        assertEq(buyPrices[1], _toQuote(125));
        assertEq(buyPrices[2], _toQuote(100));
        assertEq(buyPrices[3], _toQuote(75));
    }

    function test_priceList_removePriceWhenLastOrderMatched() external {
        uint256 price = _toQuote(100);

        // Single buy order
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        (, uint256[] memory buyPricesBefore) = PAIR.ticks();
        assertEq(buyPricesBefore.length, 1);

        // Match the order
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(10), 0);

        (, uint256[] memory buyPricesAfter) = PAIR.ticks();
        assertEq(buyPricesAfter.length, 0);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Reserve Tracking Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_baseReserve_tracksCorrectly() external {
        uint256 price = _toQuote(100);

        assertEq(PAIR.baseReserve(), 0);

        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(PAIR.baseReserve(), _toBase(10));

        vm.prank(USER3);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(PAIR.baseReserve(), _toBase(15));
    }

    function test_quoteReserve_tracksCorrectly() external {
        uint256 price = _toQuote(100);

        assertEq(PAIR.quoteReserve(), 0);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(PAIR.quoteReserve(), _toTradeVolume(price, _toBase(10)));

        vm.prank(USER3);
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(PAIR.quoteReserve(), _toTradeVolume(price, _toBase(15)));
    }

    function test_accountReserves_tracksPerUser() external {
        uint256 price = _toQuote(100);

        // USER1: BUY
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER2: SELL
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price * 2, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        (uint256 user1Base, uint256 user1Quote) = PAIR.accountReserves(USER1);
        assertEq(user1Base, 0);
        assertEq(user1Quote, _toTradeVolume(price, _toBase(10)));

        (uint256 user2Base, uint256 user2Quote) = PAIR.accountReserves(USER2);
        assertEq(user2Base, _toBase(5));
        assertEq(user2Quote, 0);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Admin Functions Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_setTickSize_success() external {
        // First, set a tickSizeSetter
        address TICK_SIZE_SETTER = address(bytes20("TICK_SIZE_SETTER"));
        vm.prank(OWNER);
        CROSS_DEX.setTickSizeSetter(TICK_SIZE_SETTER);

        (uint256 tickBefore, uint256 lotBefore) = PAIR.tickSizes();

        uint256 newTickSize = tickBefore * 2;
        uint256 newLotSize = lotBefore * 2;

        vm.prank(TICK_SIZE_SETTER);
        PAIR.setTickSize(newLotSize, newTickSize);

        (uint256 tickAfter, uint256 lotAfter) = PAIR.tickSizes();
        assertEq(tickAfter, newTickSize);
        assertEq(lotAfter, newLotSize);
    }

    function test_setTickSize_revert_unauthorized() external {
        // Without tickSizeSetter set, even OWNER can't change tick sizes
        vm.prank(USER1);
        vm.expectRevert();
        PAIR.setTickSize(1e6, 1e16);
    }

    function test_setTickSize_revert_zeroTickSize() external {
        address TICK_SIZE_SETTER = address(bytes20("TICK_SIZE_SETTER"));
        vm.prank(OWNER);
        CROSS_DEX.setTickSizeSetter(TICK_SIZE_SETTER);

        vm.prank(TICK_SIZE_SETTER);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInvalidInitializeData.selector, bytes32("tickSize")));
        PAIR.setTickSize(1e6, 0);
    }

    function test_setTickSize_revert_zeroLotSize() external {
        address TICK_SIZE_SETTER = address(bytes20("TICK_SIZE_SETTER"));
        vm.prank(OWNER);
        CROSS_DEX.setTickSizeSetter(TICK_SIZE_SETTER);

        vm.prank(TICK_SIZE_SETTER);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInvalidInitializeData.selector, bytes32("lotSize")));
        PAIR.setTickSize(0, 1e16);
    }

    function test_setTickSize_revert_invalidRatio() external {
        address TICK_SIZE_SETTER = address(bytes20("TICK_SIZE_SETTER"));
        vm.prank(OWNER);
        CROSS_DEX.setTickSizeSetter(TICK_SIZE_SETTER);

        // tickSize * lotSize must be divisible by DENOMINATOR
        vm.prank(TICK_SIZE_SETTER);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInvalidTickSize.selector, 3, 7, BASE_DECIMALS));
        PAIR.setTickSize(7, 3);
    }

    function test_skim_success_excessTokens() external {
        uint256 excessAmount = _toBase(100);

        // Send excess BASE to pair
        vm.prank(OWNER);
        BASE.transfer(address(PAIR), excessAmount);

        uint256 ownerBaseBefore = BASE.balanceOf(OWNER);

        vm.prank(OWNER);
        PAIR.skim(BASE, OWNER, excessAmount);

        assertEq(BASE.balanceOf(OWNER), ownerBaseBefore + excessAmount);
    }

    function test_skim_revert_insufficientBalance() external {
        // Create order to add to reserve
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(100), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Try to skim more than excess
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInvalidReserve.selector, address(BASE)));
        PAIR.skim(BASE, OWNER, _toBase(1)); // No excess, should fail
    }

    function test_skim_zeroAmount() external {
        // Should do nothing
        vm.prank(OWNER);
        PAIR.skim(BASE, OWNER, 0);
    }

    function test_setPause_blocksSubmit() external {
        vm.prank(OWNER);
        PAIR.setPause(true);

        vm.prank(USER1);
        vm.expectRevert();
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(100), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    function test_setPause_unpause() external {
        vm.prank(OWNER);
        PAIR.setPause(true);

        vm.prank(OWNER);
        PAIR.setPause(false);

        // Should work now
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(100), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    function test_emergencyCancelOrder_whenPaused() external {
        // Create order
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(100), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);

        // Pause
        vm.prank(OWNER);
        PAIR.setPause(true);

        // Emergency cancel
        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;
        vm.prank(OWNER);
        PAIR.emergencyCancelOrder(orderIds);

        // User should get refund
        assertGt(QUOTE.balanceOf(USER1), user1QuoteBefore);
    }

    function test_emergencyCancelOrder_revert_whenNotPaused() external {
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(100), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;

        vm.prank(OWNER);
        vm.expectRevert();
        PAIR.emergencyCancelOrder(orderIds);
    }

    function test_setFeeController_success() external {
        address newFeeController = address(FEE_CONTROLLER);
        bytes memory initData = abi.encode(FEE_COLLECTOR, uint32(10), uint32(20), uint32(0), uint32(0));

        vm.prank(OWNER);
        PAIR.setFeeController(newFeeController, initData);

        assertEq(address(PAIR.feeController()), newFeeController);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Error Condition Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_submitLimit_revert_invalidPrice_zero() external {
        // Router calls findPrevPrice first which reverts with ListZeroData for price 0
        vm.prank(USER2);
        vm.expectRevert();
        ROUTER.submitSellLimit(
            address(PAIR), 0, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    function test_submitLimit_revert_invalidPrice_notTickAligned() external {
        (uint256 tick,) = PAIR.tickSizes();
        uint256 badPrice = tick + 1; // Not aligned

        vm.prank(USER2);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInvalidPrice.selector, badPrice));
        ROUTER.submitSellLimit(
            address(PAIR), badPrice, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    function test_submitLimit_revert_invalidAmount_zero() external {
        vm.prank(USER2);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInvalidAmount.selector, 0));
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(100), 0, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    function test_submitLimit_revert_invalidAmount_notLotAligned() external {
        (, uint256 lot) = PAIR.tickSizes();
        uint256 badAmount = lot + 1; // Not aligned

        vm.prank(USER2);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInvalidAmount.selector, badAmount));
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(100), badAmount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    function test_submitSellMarket_revert_invalidAmount_zero() external {
        vm.prank(USER2);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInvalidAmount.selector, 0));
        ROUTER.submitSellMarket(address(PAIR), 0, 0);
    }

    function test_submitSellMarket_revert_invalidAmount_notLotAligned() external {
        (, uint256 lot) = PAIR.tickSizes();
        uint256 badAmount = lot + 1;

        vm.prank(USER2);
        vm.expectRevert(abi.encodeWithSelector(PairImplV3.PairInvalidAmount.selector, badAmount));
        ROUTER.submitSellMarket(address(PAIR), badAmount, 0);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // E2E Scenario Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_e2e_fullOrderLifecycle() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // 1. Place limit order
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify order exists
        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.amount, amount);

        // 2. Partial match
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(5), 0);

        order = PAIR.orderById(orderId);
        assertEq(order.amount, _toBase(5));

        // 3. Cancel remaining
        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;
        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), orderIds);

        // Order should be deleted
        order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
    }

    function test_e2e_multiUserTradingSession() external {
        // Multiple users placing orders and trading
        uint256 price100 = _toQuote(100);
        uint256 price105 = _toQuote(105);
        uint256 price110 = _toQuote(110);

        // USER1: BUY @ 100
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price100, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER2: SELL @ 110
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price110, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER3: BUY @ 105
        vm.prank(USER3);
        ROUTER.submitBuyLimit(
            address(PAIR), price105, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER4: SELL market -> matches USER3's 105 first (best buy)
        vm.prank(USER4);
        ROUTER.submitSellMarket(address(PAIR), _toBase(5), 0);

        assertEq(PAIR.matchedPrice(), price105);

        // USER4: BUY market -> matches USER2's 110 (only sell)
        vm.prank(USER4);
        ROUTER.submitBuyMarket(address(PAIR), _toTradeVolume(price110, _toBase(5)), 0);

        assertEq(PAIR.matchedPrice(), price110);
    }

    function test_e2e_priceDiscovery() external {
        // Simulate price discovery through order matching
        uint256[] memory prices = new uint256[](5);
        prices[0] = _toQuote(98);
        prices[1] = _toQuote(99);
        prices[2] = _toQuote(100);
        prices[3] = _toQuote(101);
        prices[4] = _toQuote(102);

        // Place buy orders at various prices
        for (uint256 i = 0; i < 3; i++) {
            vm.prank(USER1);
            ROUTER.submitBuyLimit(
                address(PAIR), prices[i], _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }

        // Place sell orders at various prices
        for (uint256 i = 3; i < 5; i++) {
            vm.prank(USER2);
            ROUTER.submitSellLimit(
                address(PAIR), prices[i], _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }

        // Market order to discover price
        vm.prank(USER3);
        ROUTER.submitBuyMarket(address(PAIR), _toTradeVolume(prices[3], _toBase(1)), 0);

        assertEq(PAIR.matchedPrice(), prices[3]);
    }
}
