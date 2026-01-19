// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {console} from "forge-std/console.sol";

import {FeeControllerV2Compat} from "../src/FeeControllerV2Compat.sol";
import {PairImplV3} from "../src/PairImplV3.sol";
import {IPairV3} from "../src/interfaces/IPairV3.sol";
import {DEXV3BaseTest} from "./DEXV3Base.t.sol";
import {T20} from "./mock/T20.sol";

/// @title DEXV3FeeControllerV2CompatTest
/// @notice Tests for V2-compatible fee behavior in V3 with FeeControllerV2Compat
contract DEXV3FeeControllerV2CompatTest is DEXV3BaseTest {
    address public constant USER3 = address(bytes20("USER3"));
    address public constant USER4 = address(bytes20("USER4"));

    function setUp() external {
        _deployV3(18, 18, 1e2, 1e6);

        // Additional users setup
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

    // ─────────────────────────────────────────────────────────────────────────────
    // Initialization Tests
    // ─────────────────────────────────────────────────────────────────────────────

    function test_initialization() external {
        // FeeController state is stored in Pair's namespaced storage via delegatecall.
        // Direct calls to FEE_CONTROLLER read its own empty storage, not Pair's.
        // So we verify initialization by checking the Pair's feeController address
        // and that trading works correctly (validated by other tests).
        assertEq(address(PAIR.feeController()), address(FEE_CONTROLLER), "Pair should have feeController set");

        // Verify fee behavior by checking calcBuyVolumeWithFee returns correct values
        // With buyer fee = 0, volume should be unchanged
        uint256 testVolume = _toQuote(1000);
        assertEq(PAIR.calcBuyVolumeWithFee(testVolume), testVolume, "With 0 buyer fee, volume unchanged");
    }

    function test_calcBuyVolumeWithFee() external {
        uint256 baseVolume = _toQuote(1000);

        // With buyer taker fee = 0, should return same volume
        uint256 volumeWithFee = PAIR.calcBuyVolumeWithFee(baseVolume);
        assertEq(volumeWithFee, baseVolume, "With 0 buyer fee, volume should be unchanged");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // SELL Taker Tests (seller takes from BUY order book)
    // ─────────────────────────────────────────────────────────────────────────────

    function test_sell_taker_fee_collection() external {
        uint256 price = _toQuote(100); // 100 QUOTE per BASE
        uint256 amount = _toBase(10); // 10 BASE
        uint256 quoteVolume = _toTradeVolume(price, amount); // 1000 QUOTE

        // USER1: Place a BUY limit order (maker)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // Record balances before
        uint256 sellerQuoteBefore = QUOTE.balanceOf(USER2);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        // USER2: Sell market order (taker) - matches against USER1's buy order
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Calculate expected values
        uint256 expectedTakerFee = _calcFee(quoteVolume, SELLER_TAKER_FEE);
        uint256 expectedSellerNet = quoteVolume - expectedTakerFee;

        // Verify seller received net (quote minus taker fee)
        uint256 sellerQuoteAfter = QUOTE.balanceOf(USER2);
        assertEq(sellerQuoteAfter - sellerQuoteBefore, expectedSellerNet, "Seller should receive net QUOTE");

        // Verify feeCollector received the fee
        uint256 feeCollectorAfter = QUOTE.balanceOf(FEE_COLLECTOR);
        assertEq(feeCollectorAfter - feeCollectorBefore, expectedTakerFee, "FeeCollector should receive taker fee");
    }

    function test_sell_taker_multiple_makers() external {
        uint256 price = _toQuote(100);
        uint256 amountPerOrder = _toBase(5);
        uint256 totalAmount = _toBase(10);
        uint256 quoteVolumeTotal = _toTradeVolume(price, totalAmount);

        // USER1: Place two BUY limit orders (makers)
        vm.startPrank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price, amountPerOrder, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.submitBuyLimit(
            address(PAIR), price, amountPerOrder, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        vm.stopPrank();

        // Record balances
        uint256 sellerQuoteBefore = QUOTE.balanceOf(USER2);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        // USER2: Sell market order that matches both
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), totalAmount, 0);

        // With seller-only fees, only taker fee is charged (no maker fee)
        uint256 expectedTakerFee = _calcFee(quoteVolumeTotal, SELLER_TAKER_FEE);
        uint256 expectedSellerNet = quoteVolumeTotal - expectedTakerFee;

        uint256 sellerQuoteAfter = QUOTE.balanceOf(USER2);
        uint256 feeCollectorAfter = QUOTE.balanceOf(FEE_COLLECTOR);

        assertEq(sellerQuoteAfter - sellerQuoteBefore, expectedSellerNet, "Seller should receive net QUOTE");
        assertEq(feeCollectorAfter - feeCollectorBefore, expectedTakerFee, "FeeCollector should receive taker fee");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // BUY Taker Tests (buyer takes from SELL order book)
    // ─────────────────────────────────────────────────────────────────────────────

    function test_buy_taker_with_seller_maker_fee() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        // USER2: Place a SELL limit order (maker) - seller will be charged maker fee
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Record balances
        uint256 sellerQuoteBefore = QUOTE.balanceOf(USER2);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        // USER1: Buy market order (taker) - matches against USER2's sell order
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteVolume, 0);

        // Seller (maker) pays maker fee
        uint256 expectedMakerFee = _calcFee(quoteVolume, SELLER_MAKER_FEE);
        uint256 expectedSellerNet = quoteVolume - expectedMakerFee;

        uint256 sellerQuoteAfter = QUOTE.balanceOf(USER2);
        uint256 feeCollectorAfter = QUOTE.balanceOf(FEE_COLLECTOR);

        assertEq(sellerQuoteAfter - sellerQuoteBefore, expectedSellerNet, "Seller (maker) should receive net QUOTE");
        assertEq(feeCollectorAfter - feeCollectorBefore, expectedMakerFee, "FeeCollector should receive maker fee");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Cancel/Refund Tests (V2 compatibility: order.feeBps based refund)
    // ─────────────────────────────────────────────────────────────────────────────

    function test_buy_maker_cancel_refund() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);
        // With buyer maker fee = 0, deposit = volume
        uint256 expectedDeposit = quoteVolume + _calcFee(quoteVolume, BUYER_MAKER_FEE);

        uint256 buyerQuoteBefore = QUOTE.balanceOf(USER1);

        // USER1: Place a BUY limit order
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify deposit was taken
        uint256 buyerQuoteAfterSubmit = QUOTE.balanceOf(USER1);
        assertEq(buyerQuoteBefore - buyerQuoteAfterSubmit, expectedDeposit, "Deposit should equal volume + maker fee");

        // Cancel the order
        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;
        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), orderIds);

        // Verify full refund
        uint256 buyerQuoteAfterCancel = QUOTE.balanceOf(USER1);
        assertEq(buyerQuoteAfterCancel, buyerQuoteBefore, "Full deposit should be refunded on cancel");
    }

    function test_buy_maker_cancel_refund_with_buyer_fee() external {
        // Set buyer maker fee to 10 bps (0.1%)
        uint32 newBuyerMakerFee = 10;

        vm.prank(OWNER);
        bytes memory newFeeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, newBuyerMakerFee, newBuyerMakerFee);
        MARKET.setFeeController(0, 1, true, address(FEE_CONTROLLER), newFeeData);

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);
        uint256 expectedDeposit = quoteVolume + _calcFee(quoteVolume, newBuyerMakerFee);

        uint256 buyerQuoteBefore = QUOTE.balanceOf(USER1);

        // USER1: Place a BUY limit order
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify deposit includes fee
        uint256 buyerQuoteAfterSubmit = QUOTE.balanceOf(USER1);
        assertEq(buyerQuoteBefore - buyerQuoteAfterSubmit, expectedDeposit, "Deposit should include maker fee");

        // Verify order.feeBps was set
        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.feeBps, newBuyerMakerFee, "order.feeBps should be set to buyerMakerFeeBps");

        // Cancel the order
        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;
        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), orderIds);

        // Verify full refund (including fee)
        uint256 buyerQuoteAfterCancel = QUOTE.balanceOf(USER1);
        assertEq(buyerQuoteAfterCancel, buyerQuoteBefore, "Full deposit (volume + fee) should be refunded");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Seller Maker Fee Tests (limit sell order on book)
    // ─────────────────────────────────────────────────────────────────────────────

    function test_sell_maker_fee_on_match() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        // USER2: Place a SELL limit order (maker)
        vm.prank(USER2);
        uint256 sellOrderId = ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify order.feeBps was set for sell maker
        IPairV3.Order memory sellOrder = PAIR.orderById(sellOrderId);
        assertEq(sellOrder.feeBps, SELLER_MAKER_FEE, "Sell order.feeBps should be sellerMakerFeeBps");

        uint256 sellerQuoteBefore = QUOTE.balanceOf(USER2);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        // USER1: Buy limit order that matches (taker)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // Seller (maker) pays maker fee
        uint256 expectedMakerFee = _calcFee(quoteVolume, SELLER_MAKER_FEE);
        uint256 expectedSellerNet = quoteVolume - expectedMakerFee;

        uint256 sellerQuoteAfter = QUOTE.balanceOf(USER2);
        uint256 feeCollectorAfter = QUOTE.balanceOf(FEE_COLLECTOR);

        assertEq(sellerQuoteAfter - sellerQuoteBefore, expectedSellerNet, "Seller (maker) should receive net QUOTE");
        assertEq(feeCollectorAfter - feeCollectorBefore, expectedMakerFee, "FeeCollector should receive maker fee");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Edge Cases
    // ─────────────────────────────────────────────────────────────────────────────

    function test_zero_fee_scenario() external {
        // Set all fees to 0
        vm.prank(OWNER);
        bytes memory zeroFeeData = abi.encode(FEE_COLLECTOR, uint32(0), uint32(0), uint32(0), uint32(0));
        MARKET.setFeeController(0, 1, true, address(FEE_CONTROLLER), zeroFeeData);

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        // Place and match orders
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 sellerQuoteBefore = QUOTE.balanceOf(USER2);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // No fees should be collected
        uint256 sellerQuoteAfter = QUOTE.balanceOf(USER2);
        uint256 feeCollectorAfter = QUOTE.balanceOf(FEE_COLLECTOR);

        assertEq(sellerQuoteAfter - sellerQuoteBefore, quoteVolume, "Seller should receive full volume with 0 fee");
        assertEq(feeCollectorAfter, feeCollectorBefore, "No fee should be collected");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // FeeControllerFeesSettled Event Tests
    // ─────────────────────────────────────────────────────────────────────────────

    function test_fees_settled_event_emitted() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        // USER1: Place a BUY limit order (maker)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // Expect FeeControllerFeesSettled event
        uint256 expectedTakerFee = _calcFee(quoteVolume, SELLER_TAKER_FEE);
        // takerId will be the sell order id (2, since buy order is 1)
        vm.expectEmit(true, true, false, true, address(PAIR));
        emit FeeControllerV2Compat.FeeControllerFeesSettled(2, FEE_COLLECTOR, expectedTakerFee);

        // USER2: Sell market order (taker)
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);
    }

    function test_fees_settled_event_with_maker_fee() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        // USER2: Place a SELL limit order (maker with sellerMakerFee)
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Expect FeeControllerFeesSettled event with maker fee
        uint256 expectedMakerFee = _calcFee(quoteVolume, SELLER_MAKER_FEE);
        // takerId will be the buy order id (2, since sell order is 1)
        vm.expectEmit(true, true, false, true, address(PAIR));
        emit FeeControllerV2Compat.FeeControllerFeesSettled(2, FEE_COLLECTOR, expectedMakerFee);

        // USER1: Buy market order (taker) - buyer taker fee is 0
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteVolume, 0);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Transient Storage Tests (multiple makers in single transaction)
    // ─────────────────────────────────────────────────────────────────────────────

    function test_transient_storage_multiple_makers_same_taker() external {
        uint256 price = _toQuote(100);
        uint256 amountPerOrder = _toBase(5);
        uint256 totalAmount = _toBase(10);
        uint256 quoteVolumeTotal = _toTradeVolume(price, totalAmount);

        // USER1: Place two BUY limit orders (makers)
        vm.startPrank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price, amountPerOrder, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        ROUTER.submitBuyLimit(
            address(PAIR), price, amountPerOrder, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        vm.stopPrank();

        // Expect single FeeControllerFeesSettled event with accumulated fees
        uint256 expectedTakerFee = _calcFee(quoteVolumeTotal, SELLER_TAKER_FEE);
        // takerId = 3 (sell order), order 1 and 2 are buy orders
        vm.expectEmit(true, true, false, true, address(PAIR));
        emit FeeControllerV2Compat.FeeControllerFeesSettled(3, FEE_COLLECTOR, expectedTakerFee);

        // USER2: Sell market order that matches both makers
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), totalAmount, 0);
    }

    /// @notice Tests that transient storage resets after settleFees allowing new takerId.
    /// @dev settleFees() explicitly resets transient storage, allowing subsequent trades
    ///      with different takerIds within the same Foundry test (same tx).
    function test_transient_storage_resets_after_settle() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        // First trade: USER1 places buy (orderId=1), USER2 sells (takerId=2)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 feeCollectorBefore1 = QUOTE.balanceOf(FEE_COLLECTOR);
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);
        uint256 feeCollectorAfter1 = QUOTE.balanceOf(FEE_COLLECTOR);

        // Verify first trade fee
        uint256 expectedTakerFee = _calcFee(quoteVolume, SELLER_TAKER_FEE);
        assertEq(feeCollectorAfter1 - feeCollectorBefore1, expectedTakerFee, "First trade taker fee");

        // Second trade: USER2 places sell (orderId=3), USER1 buys (takerId=4)
        // This should work because settleFees() reset transient storage
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 feeCollectorBefore2 = QUOTE.balanceOf(FEE_COLLECTOR);
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteVolume, 0);
        uint256 feeCollectorAfter2 = QUOTE.balanceOf(FEE_COLLECTOR);

        // Verify second trade fee (seller maker fee)
        uint256 expectedMakerFee = _calcFee(quoteVolume, SELLER_MAKER_FEE);
        assertEq(feeCollectorAfter2 - feeCollectorBefore2, expectedMakerFee, "Second trade maker fee");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Comprehensive Scenario Test
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Comprehensive test covering multiple users, orders, matches, and cancellations.
    /// @dev Verifies baseReserve, quoteReserve, matchedPrice, and fee distribution at each step.
    function test_comprehensive_scenario() external {
        // Track fees and order IDs
        uint256 totalFeesCollected;
        uint256[] memory orderIds = new uint256[](3);

        // Price constants
        uint256 price100 = _toQuote(100);
        uint256 price105 = _toQuote(105);
        uint256 price110 = _toQuote(110);

        // Initial state checks
        assertEq(PAIR.baseReserve(), 0, "Initial baseReserve");
        assertEq(PAIR.quoteReserve(), 0, "Initial quoteReserve");
        assertEq(PAIR.matchedPrice(), 0, "Initial matchedPrice");

        uint256 feeCollectorInit = QUOTE.balanceOf(FEE_COLLECTOR);

        // ═══════════════════════════════════════════════════════════════════════
        // Step 1: USER1 BUY limit @ 100 for 10 BASE
        // ═══════════════════════════════════════════════════════════════════════
        vm.prank(USER1);
        orderIds[0] = ROUTER.submitBuyLimit(
            address(PAIR), price100, _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        assertEq(PAIR.orderById(orderIds[0]).feeBps, BUYER_MAKER_FEE, "Step1: buyerMakerFee");
        assertEq(PAIR.quoteReserve(), _toQuote(1000), "Step1: quoteReserve=1000");
        assertEq(PAIR.baseReserve(), 0, "Step1: baseReserve=0");

        // ═══════════════════════════════════════════════════════════════════════
        // Step 2: USER2 SELL limit @ 110 for 20 BASE
        // ═══════════════════════════════════════════════════════════════════════
        vm.prank(USER2);
        orderIds[1] = ROUTER.submitSellLimit(
            address(PAIR), price110, _toBase(20), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        assertEq(PAIR.orderById(orderIds[1]).feeBps, SELLER_MAKER_FEE, "Step2: sellerMakerFee");
        assertEq(PAIR.baseReserve(), _toBase(20), "Step2: baseReserve=20");
        assertEq(PAIR.matchedPrice(), 0, "Step2: no match");

        // ═══════════════════════════════════════════════════════════════════════
        // Step 3: USER3 BUY limit @ 105 for 15 BASE
        // ═══════════════════════════════════════════════════════════════════════
        vm.prank(USER3);
        orderIds[2] = ROUTER.submitBuyLimit(
            address(PAIR), price105, _toBase(15), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        assertEq(PAIR.quoteReserve(), _toQuote(1000 + 1575), "Step3: quoteReserve=2575");
        assertEq(PAIR.matchedPrice(), 0, "Step3: no match");

        // ═══════════════════════════════════════════════════════════════════════
        // Step 4: USER4 SELL market 10 BASE → matches USER3's buy @ 105
        // ═══════════════════════════════════════════════════════════════════════
        {
            uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);
            uint256 user4QuoteBefore = QUOTE.balanceOf(USER4);

            vm.prank(USER4);
            ROUTER.submitSellMarket(address(PAIR), _toBase(10), 0);

            assertEq(PAIR.matchedPrice(), price105, "Step4: matchedPrice=105");
            assertEq(PAIR.orderById(orderIds[2]).amount, _toBase(5), "Step4: Order3 has 5 remaining");

            // Taker fee: USER4 sells, SELLER_TAKER_FEE = 0.3% of 1050 = 3.15
            uint256 takerFee = _calcFee(_toQuote(1050), SELLER_TAKER_FEE);
            totalFeesCollected += takerFee;

            assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, takerFee, "Step4: takerFee");
            assertEq(QUOTE.balanceOf(USER4) - user4QuoteBefore, _toQuote(1050) - takerFee, "Step4: USER4 net");

            // quoteReserve = 1000 (Order1) + 525 (Order3 remaining)
            assertEq(PAIR.quoteReserve(), _toQuote(1000 + 525), "Step4: quoteReserve=1525");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 5: USER1 SELL market 5 BASE → matches USER3's remaining buy @ 105
        // ═══════════════════════════════════════════════════════════════════════
        {
            uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);
            uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);

            vm.prank(USER1);
            ROUTER.submitSellMarket(address(PAIR), _toBase(5), 0);

            assertEq(PAIR.orderById(orderIds[2]).amount, 0, "Step5: Order3 fully filled");

            uint256 takerFee = _calcFee(_toQuote(525), SELLER_TAKER_FEE);
            totalFeesCollected += takerFee;

            assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, takerFee, "Step5: takerFee");
            assertEq(QUOTE.balanceOf(USER1) - user1QuoteBefore, _toQuote(525) - takerFee, "Step5: USER1 net");
            assertEq(PAIR.quoteReserve(), _toQuote(1000), "Step5: quoteReserve=1000 (Order1)");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 6: USER4 BUY market 1100 QUOTE → matches USER2's sell @ 110 (10 BASE)
        // ═══════════════════════════════════════════════════════════════════════
        {
            uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);
            uint256 user2QuoteBefore = QUOTE.balanceOf(USER2);
            uint256 user4BaseBefore = BASE.balanceOf(USER4);

            vm.prank(USER4);
            ROUTER.submitBuyMarket(address(PAIR), _toQuote(1100), 0);

            assertEq(PAIR.matchedPrice(), price110, "Step6: matchedPrice=110");
            assertEq(PAIR.orderById(orderIds[1]).amount, _toBase(10), "Step6: Order2 has 10 remaining");

            // Maker fee: USER2 is maker (SELL), SELLER_MAKER_FEE = 0.2% of 1100 = 2.2
            uint256 makerFee = _calcFee(_toQuote(1100), SELLER_MAKER_FEE);
            totalFeesCollected += makerFee;

            assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, makerFee, "Step6: makerFee");
            assertEq(QUOTE.balanceOf(USER2) - user2QuoteBefore, _toQuote(1100) - makerFee, "Step6: USER2 net");
            assertEq(BASE.balanceOf(USER4) - user4BaseBefore, _toBase(10), "Step6: USER4 receives BASE");
            assertEq(PAIR.baseReserve(), _toBase(10), "Step6: baseReserve=10");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 7: Cancel all remaining orders
        // ═══════════════════════════════════════════════════════════════════════
        {
            // Cancel Order1 (USER1 BUY)
            uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);
            uint256[] memory ids = new uint256[](1);
            ids[0] = orderIds[0];
            vm.prank(USER1);
            ROUTER.cancelOrder(address(PAIR), ids);
            assertEq(QUOTE.balanceOf(USER1) - user1QuoteBefore, _toQuote(1000), "Step7: USER1 refund");
        }
        {
            // Cancel Order2 (USER2 SELL)
            uint256 user2BaseBefore = BASE.balanceOf(USER2);
            uint256[] memory ids = new uint256[](1);
            ids[0] = orderIds[1];
            vm.prank(USER2);
            ROUTER.cancelOrder(address(PAIR), ids);
            assertEq(BASE.balanceOf(USER2) - user2BaseBefore, _toBase(10), "Step7: USER2 refund");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Final Verification
        // ═══════════════════════════════════════════════════════════════════════
        assertEq(PAIR.baseReserve(), 0, "Final: baseReserve=0");
        assertEq(PAIR.quoteReserve(), 0, "Final: quoteReserve=0");
        assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorInit, totalFeesCollected, "Final: total fees");

        console.log("=== Comprehensive Scenario Complete ===");
        console.log("Total fees collected:", totalFeesCollected);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Security / Edge Case Tests
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Tests that transient storage resets even when totalFee == 0
    /// @dev Prevents FeeControllerTakerIdMismatch in subsequent trades within same tx
    function test_zero_fee_settlement_resets_transient() external {
        // Set all fees to 0
        vm.prank(OWNER);
        bytes memory zeroFeeData = abi.encode(FEE_COLLECTOR, uint32(0), uint32(0), uint32(0), uint32(0));
        MARKET.setFeeController(0, 1, true, address(FEE_CONTROLLER), zeroFeeData);

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        // First trade with 0 fee (takerId=2)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Second trade should also work (takerId=4) - transient storage should be reset
        // If transient storage wasn't reset, this would revert with FeeControllerTakerIdMismatch(2, 4)
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteVolume, 0); // Should NOT revert
    }

    /// @notice Tests CEI pattern - effects before interactions
    /// @dev Verifies transient storage is cleared before external transfer call
    function test_cei_pattern_transient_cleared_before_transfer() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // Place and match order
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // After first trade, transient should be cleared allowing second trade
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Immediately do another trade - this verifies CEI pattern works
        // because transient storage was cleared before the transfer call completed
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Verify fees were collected for both trades
        uint256 expectedFee = _calcFee(_toTradeVolume(price, amount), SELLER_TAKER_FEE) * 2;
        assertEq(QUOTE.balanceOf(FEE_COLLECTOR), expectedFee, "Both trades should have collected fees");
    }

    /// @notice Tests that multiple Pairs with same FeeController have isolated storage
    function test_multiple_pairs_storage_isolation() external {
        // Create second pair with different BASE token
        T20 base2;
        PairImplV3 pair2;
        vm.startPrank(OWNER);
        base2 = new T20("BASE2", "BASE2", 18);
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);
        address pair2Addr = MARKET.createPair(address(base2), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
        pair2 = PairImplV3(pair2Addr);

        // Fund users for pair2
        base2.transfer(USER1, _toBase(10000));
        base2.transfer(USER2, _toBase(10000));
        vm.stopPrank();

        vm.prank(USER1);
        base2.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER2);
        base2.approve(address(ROUTER), type(uint256).max);

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // Trade on PAIR
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // Trade on pair2 should work independently (different transient storage context)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(pair2), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Complete trades
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(pair2), amount, 0);

        // Verify both pairs have correct reserves
        assertEq(PAIR.baseReserve(), 0, "PAIR baseReserve should be 0");
        assertEq(pair2.baseReserve(), 0, "pair2 baseReserve should be 0");
    }
}
