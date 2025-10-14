// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./DEXFeeV2.t.sol";

contract DEXUpgradeV2Test is DEXFeeV2Test {
    // Test upgrading from MarketImpl to MarketImplV2
    function test_upgrade_market_to_v2() external {
        // Create a new MarketImplV2 implementation
        address newMarketImpl = address(new MarketImplV2());

        // Upgrade the market to V2
        vm.prank(OWNER);
        MARKET.upgradeToAndCall(newMarketImpl, "");

        // Verify the upgrade worked
        IMarketV2.FeeConfig memory feeConfig = MARKET.getFeeConfig();
        assertEq(
            feeConfig.sellerMakerFeeBps, MARKET_SELLER_MAKER_FEE, "Seller maker fee should be preserved after upgrade"
        );
        assertEq(
            feeConfig.sellerTakerFeeBps, MARKET_SELLER_TAKER_FEE, "Seller taker fee should be preserved after upgrade"
        );
        assertEq(
            feeConfig.buyerMakerFeeBps, MARKET_BUYER_MAKER_FEE, "Buyer maker fee should be preserved after upgrade"
        );
        assertEq(
            feeConfig.buyerTakerFeeBps, MARKET_BUYER_TAKER_FEE, "Buyer taker fee should be preserved after upgrade"
        );

        // Test new V2 functionality
        vm.prank(OWNER);
        MARKET.setMarketFees(100, 200, 300, 400);

        IMarketV2.FeeConfig memory newFeeConfig = MARKET.getFeeConfig();
        assertEq(newFeeConfig.sellerMakerFeeBps, 100, "New seller maker fee should be set");
        assertEq(newFeeConfig.sellerTakerFeeBps, 200, "New seller taker fee should be set");
        assertEq(newFeeConfig.buyerMakerFeeBps, 300, "New buyer maker fee should be set");
        assertEq(newFeeConfig.buyerTakerFeeBps, 400, "New buyer taker fee should be set");
    }

    // Test that V2 pairs work with existing V1 infrastructure
    function test_v2_pair_backward_compatibility() external {
        // Test that all existing functions still work
        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(sellPrices.length, 0, "Should start with no sell orders");
        assertEq(buyPrices.length, 0, "Should start with no buy orders");

        // Test order placement works the same
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        IPair.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, USER1, "Order owner should be correct");
        assertEq(order.amount, _toBase(100), "Order amount should be correct");
    }

    // Test mixed V1/V2 market interactions
    function test_mixed_v1_v2_operations() external {
        // Test that fee collection still works correctly
        uint256 initialFeeBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // Execute a trade that should generate fees
        vm.prank(USER1);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER2);
        ROUTER.submitBuyMarket(address(PAIR), _toQuote(100), type(uint256).max);

        uint256 finalFeeBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 feeCollected = finalFeeBalance - initialFeeBalance;

        // Should collect fees using the effective rates
        assertTrue(feeCollected > 0, "Fees should be collected");

        // Verify fee calculation matches expected
        // In V2, both seller (maker) and buyer (taker) pay fees
        uint256 expectedSellerFee = Math.mulDiv(_toQuote(100), MARKET_SELLER_MAKER_FEE, 10000);
        uint256 expectedBuyerFee = Math.mulDiv(_toQuote(100), MARKET_BUYER_TAKER_FEE, 10000);
        uint256 expectedTotalFee = expectedSellerFee + expectedBuyerFee;
        assertEq(feeCollected, expectedTotalFee, "Fee amount should match expected total");
    }

    // Test storage layout compatibility
    function test_storage_layout_compatibility() external {
        // Test that basic storage variables are preserved
        assertEq(MARKET.feeCollector(), FEE_COLLECTOR, "Fee collector should be preserved");

        // Test that the new 4-parameter fees are accessible
        IMarketV2.FeeConfig memory feeConfig = MARKET.getFeeConfig();
        assertEq(feeConfig.sellerMakerFeeBps, MARKET_SELLER_MAKER_FEE, "Seller maker fee should be preserved");

        // Test pair storage
        assertEq(address(PAIR.BASE()), address(BASE), "BASE token should be preserved");
        assertEq(address(PAIR.QUOTE()), address(QUOTE), "QUOTE token should be preserved");
        assertEq(PAIR.DENOMINATOR(), BASE_DECIMALS, "Denominator should be preserved");
    }

    // Test edge cases for fee values
    function test_fee_edge_cases() external {
        // Test maximum valid fee (9999 BPS = 99.99%)
        vm.prank(OWNER);
        PAIR.setPairFees(9999, 9999, 9999, 9999);

        (uint32 pairSellerMakerFee, uint32 pairSellerTakerFee, uint32 pairBuyerMakerFee, uint32 pairBuyerTakerFee) =
            PAIR.feeConfig();
        assertEq(pairSellerMakerFee, 9999, "Should accept maximum valid fee for seller maker");
        assertEq(pairSellerTakerFee, 9999, "Should accept maximum valid fee for seller taker");
        assertEq(pairBuyerMakerFee, 9999, "Should accept maximum valid fee for buyer maker");
        assertEq(pairBuyerTakerFee, 9999, "Should accept maximum valid fee for buyer taker");

        // Test boundary value exactly at NO_FEE_BPS
        vm.prank(OWNER);
        PAIR.setPairFees(NO_FEE_BPS, 9999, NO_FEE_BPS, 9999);

        (
            uint32 effectiveSellerMakerFee,
            uint32 effectiveSellerTakerFee,
            uint32 effectiveBuyerMakerFee,
            uint32 effectiveBuyerTakerFee
        ) = PAIR.getEffectiveFees();
        assertEq(
            effectiveSellerMakerFee,
            MARKET_SELLER_MAKER_FEE,
            "NO_FEE_BPS should fallback to market fee for seller maker"
        );
        assertEq(effectiveSellerTakerFee, 9999, "9999 should use pair fee for seller taker");
        assertEq(
            effectiveBuyerMakerFee, MARKET_BUYER_MAKER_FEE, "NO_FEE_BPS should fallback to market fee for buyer maker"
        );
        assertEq(effectiveBuyerTakerFee, 9999, "9999 should use pair fee for buyer taker");
    }

    // Test fee application in complex trading scenarios
    function test_complex_trading_scenario() external {
        // Set different pair fees
        uint32 sellerMakerFee = 15; // 0.15%
        uint32 sellerTakerFee = 25; // 0.25%
        uint32 buyerMakerFee = 20; // 0.20%
        uint32 buyerTakerFee = 30; // 0.30%

        vm.prank(OWNER);
        PAIR.setPairFees(sellerMakerFee, sellerTakerFee, buyerMakerFee, buyerTakerFee);

        uint256 initialFeeBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // Create multiple limit orders (makers)
        vm.prank(USER1);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(50), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER1);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(2), _toBase(50), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Execute partial market order (taker) that matches both
        vm.prank(USER2);
        ROUTER.submitBuyMarket(address(PAIR), _toQuote(150), type(uint256).max);

        uint256 finalFeeBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 totalFeesCollected = finalFeeBalance - initialFeeBalance;

        // Calculate expected fees:
        // Seller (USER1) fees: 50 QUOTE at seller maker fee + 100 QUOTE at seller maker fee
        // Buyer (USER2) fees: 50 QUOTE at buyer taker fee + 100 QUOTE at buyer taker fee
        uint256 expectedSellerFees =
            Math.mulDiv(_toQuote(50), sellerMakerFee, 10000) + Math.mulDiv(_toQuote(100), sellerMakerFee, 10000);
        uint256 expectedBuyerFees =
            Math.mulDiv(_toQuote(50), buyerTakerFee, 10000) + Math.mulDiv(_toQuote(100), buyerTakerFee, 10000);
        uint256 expectedTotalFees = expectedSellerFees + expectedBuyerFees;
        assertEq(
            totalFeesCollected,
            expectedTotalFees,
            "Total fees should match expected from multiple orders with both seller and buyer fees"
        );
    }

    // Test that fee changes don't affect existing orders
    function test_fee_changes_dont_affect_existing_orders() external {
        uint32 initialSellerMakerFee = 10;
        uint32 initialSellerTakerFee = 20;
        uint32 initialBuyerMakerFee = 15;
        uint32 initialBuyerTakerFee = 25;

        vm.prank(OWNER);
        PAIR.setPairFees(initialSellerMakerFee, initialSellerTakerFee, initialBuyerMakerFee, initialBuyerTakerFee);

        // Place a limit order
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify the order has the initial maker fee stored
        IPair.Order memory order = PAIR.orderById(orderId);
        assertEq(order.feeBps, initialSellerMakerFee, "Order should store the seller maker fee at creation time");

        // Change the fees
        vm.prank(OWNER);
        PAIR.setPairFees(50, 60, 70, 80);

        uint256 initialFeeBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // Execute the order
        vm.prank(USER2);
        ROUTER.submitBuyMarket(address(PAIR), _toQuote(100), type(uint256).max);

        uint256 finalFeeBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 feeCollected = finalFeeBalance - initialFeeBalance;

        // Fee should be calculated using the original rates for both seller and buyer
        // Seller (USER1) uses original seller maker fee, Buyer (USER2) uses current buyer taker fee (80)
        uint256 expectedSellerFee = Math.mulDiv(_toQuote(100), initialSellerMakerFee, 10000);
        uint256 expectedBuyerFee = Math.mulDiv(_toQuote(100), 80, 10000); // Current buyer taker fee
        uint256 expectedTotalFee = expectedSellerFee + expectedBuyerFee;
        assertEq(
            feeCollected,
            expectedTotalFee,
            "Existing orders should use their original fee rate, but current buyer fees apply"
        );
    }
}
