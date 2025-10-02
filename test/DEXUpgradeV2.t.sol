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
        (uint32 makerFee, uint32 takerFee) = MARKET.getMarketFees();
        assertEq(makerFee, MARKET_MAKER_FEE, "Maker fee should be preserved after upgrade");
        assertEq(takerFee, MARKET_TAKER_FEE, "Taker fee should be preserved after upgrade");

        // Test new V2 functionality
        vm.prank(OWNER);
        MARKET.setMarketFees(100, 200);

        (uint32 newMakerFee, uint32 newTakerFee) = MARKET.getMarketFees();
        assertEq(newMakerFee, 100, "New maker fee should be set");
        assertEq(newTakerFee, 200, "New taker fee should be set");
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
        uint256 expectedFee = Math.mulDiv(_toQuote(100), MARKET_MAKER_FEE, 10000);
        assertEq(feeCollected, expectedFee, "Fee amount should match expected");
    }

    // Test storage layout compatibility
    function test_storage_layout_compatibility() external {
        // Test that basic storage variables are preserved
        assertEq(MARKET.feeCollector(), FEE_COLLECTOR, "Fee collector should be preserved");
        assertEq(MARKET.makerFeeBps(), MARKET_MAKER_FEE, "Maker fee should be preserved");

        // Test pair storage
        assertEq(address(PAIR.BASE()), address(BASE), "BASE token should be preserved");
        assertEq(address(PAIR.QUOTE()), address(QUOTE), "QUOTE token should be preserved");
        assertEq(PAIR.DENOMINATOR(), BASE_DECIMALS, "Denominator should be preserved");
    }

    // Test edge cases for fee values
    function test_fee_edge_cases() external {
        // Test maximum valid fee (9999 BPS = 99.99%)
        vm.prank(OWNER);
        PAIR.setPairFees(9999, 9999);

        (uint32 pairMakerFee, uint32 pairTakerFee) = PAIR.getPairFees();
        assertEq(pairMakerFee, 9999, "Should accept maximum valid fee");
        assertEq(pairTakerFee, 9999, "Should accept maximum valid fee");

        // Test boundary value exactly at NO_FEE_BPS
        vm.prank(OWNER);
        PAIR.setPairFees(NO_FEE_BPS, 9999);

        (uint32 effectiveMakerFee, uint32 effectiveTakerFee) = PAIR.getEffectiveFees();
        assertEq(effectiveMakerFee, MARKET_MAKER_FEE, "NO_FEE_BPS should fallback to market fee");
        assertEq(effectiveTakerFee, 9999, "9999 should use pair fee");
    }

    // Test fee application in complex trading scenarios
    function test_complex_trading_scenario() external {
        // Set different pair fees
        uint32 makerFee = 15; // 0.15%
        uint32 takerFee = 25; // 0.25%

        vm.prank(OWNER);
        PAIR.setPairFees(makerFee, takerFee);

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

        // Calculate expected fees: 50 QUOTE at maker fee + 100 QUOTE at maker fee
        uint256 expectedFees = Math.mulDiv(_toQuote(50), makerFee, 10000) + Math.mulDiv(_toQuote(100), makerFee, 10000);
        assertEq(totalFeesCollected, expectedFees, "Total fees should match expected from multiple orders");
    }

    // Test that fee changes don't affect existing orders
    function test_fee_changes_dont_affect_existing_orders() external {
        uint32 initialMakerFee = 10;
        uint32 initialTakerFee = 20;

        vm.prank(OWNER);
        PAIR.setPairFees(initialMakerFee, initialTakerFee);

        // Place a limit order
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify the order has the initial maker fee stored
        IPair.Order memory order = PAIR.orderById(orderId);
        assertEq(order.feeBps, initialMakerFee, "Order should store the maker fee at creation time");

        // Change the fees
        vm.prank(OWNER);
        PAIR.setPairFees(50, 60);

        uint256 initialFeeBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // Execute the order
        vm.prank(USER2);
        ROUTER.submitBuyMarket(address(PAIR), _toQuote(100), type(uint256).max);

        uint256 finalFeeBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 feeCollected = finalFeeBalance - initialFeeBalance;

        // Fee should still be calculated using the original rate
        uint256 expectedFee = Math.mulDiv(_toQuote(100), initialMakerFee, 10000);
        assertEq(feeCollected, expectedFee, "Existing orders should use their original fee rate");
    }
}
