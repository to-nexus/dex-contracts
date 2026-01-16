// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";

import {BPS_DENOMINATOR} from "../src/interfaces/IFeeController.sol";
import {IPairV3} from "../src/interfaces/IPairV3.sol";
import {DEXV3BaseTest} from "./DEXV3Base.t.sol";

/// @title DEXV3FeeControllerV2CompatTest
/// @notice Tests for V2-compatible fee behavior in V3 with FeeControllerV2Compat
contract DEXV3FeeControllerV2CompatTest is DEXV3BaseTest {
    function setUp() external {
        _deployV3(18, 18, 1e2, 1e6);
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
}
