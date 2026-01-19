// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {console} from "forge-std/console.sol";

import {IPairV3} from "../src/interfaces/IPairV3.sol";
import {DEXV3BaseTest} from "./DEXV3Base.t.sol";

/// @title DEXV3E2EScenarioV2CompatTest
/// @notice End-to-end scenario tests for V3 matching engine with FeeControllerV2Compat
contract DEXV3E2EScenarioV2CompatTest is DEXV3BaseTest {
    address public constant USER3 = address(bytes20("USER3"));
    address public constant USER4 = address(bytes20("USER4"));

    // Custom fee structure for more comprehensive testing
    uint32 constant TEST_SELLER_MAKER_FEE = 20; // 0.2%
    uint32 constant TEST_SELLER_TAKER_FEE = 30; // 0.3%
    uint32 constant TEST_BUYER_MAKER_FEE = 10; // 0.1%
    uint32 constant TEST_BUYER_TAKER_FEE = 25; // 0.25%

    // Contract-level price constants to reduce stack depth
    uint256 internal _price100;
    uint256 internal _price105;
    uint256 internal _price110;
    uint256 internal _price95;
    uint256 internal _price120;

    function setUp() external {
        // Override default fees before deployment
        SELLER_MAKER_FEE = TEST_SELLER_MAKER_FEE;
        SELLER_TAKER_FEE = TEST_SELLER_TAKER_FEE;
        BUYER_MAKER_FEE = TEST_BUYER_MAKER_FEE;
        BUYER_TAKER_FEE = TEST_BUYER_TAKER_FEE;

        _deployV3(18, 18, 1e2, 1e6);

        // Initialize price constants
        _price100 = _toQuote(100);
        _price105 = _toQuote(105);
        _price110 = _toQuote(110);
        _price95 = _toQuote(95);
        _price120 = _toQuote(120);

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

    // ═══════════════════════════════════════════════════════════════════════════
    // E2E Full Scenario Test
    // ═══════════════════════════════════════════════════════════════════════════

    /// @notice Full E2E scenario testing order lifecycle with V2Compat fees.
    /// @dev Steps: submit → match → partial fill → cancel → verify reserves zero
    function test_e2e_v2compat_fullScenario() external {
        uint256 feeCollectorInit = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 totalFeesCollected;

        assertEq(PAIR.baseReserve(), 0, "Step1: Initial baseReserve should be 0");
        assertEq(PAIR.quoteReserve(), 0, "Step1: Initial quoteReserve should be 0");

        // ═══════════════════════════════════════════════════════════════════════
        // Step 2: Multiple SELL Limit Orders (USER2, USER3)
        // ═══════════════════════════════════════════════════════════════════════
        uint256 sellOrder1;
        uint256 sellOrder2;
        {
            uint256 user2BaseBefore = BASE.balanceOf(USER2);
            uint256 user3BaseBefore = BASE.balanceOf(USER3);

            vm.prank(USER2);
            sellOrder1 = ROUTER.submitSellLimit(
                address(PAIR), _price100, _toBase(100), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            vm.prank(USER3);
            sellOrder2 = ROUTER.submitSellLimit(
                address(PAIR), _price110, _toBase(50), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            assertEq(PAIR.baseReserve(), _toBase(150), "Step2: baseReserve = 150");
            assertEq(PAIR.quoteReserve(), 0, "Step2: quoteReserve = 0");
            assertEq(BASE.balanceOf(USER2), user2BaseBefore - _toBase(100), "Step2: USER2 BASE decreased");
            assertEq(BASE.balanceOf(USER3), user3BaseBefore - _toBase(50), "Step2: USER3 BASE decreased");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 3: BUY Limit Order Partial Match (USER1)
        // ═══════════════════════════════════════════════════════════════════════
        uint256 buyOrder1;
        {
            uint256 user1BaseBefore = BASE.balanceOf(USER1);
            uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

            uint256 tradeVolume = _toTradeVolume(_price100, _toBase(60));
            uint256 sellerMakerFee = _calcFee(tradeVolume, SELLER_MAKER_FEE);
            uint256 buyerTakerFee = _calcFee(tradeVolume, BUYER_TAKER_FEE);
            uint256 expectedTotalFee = sellerMakerFee + buyerTakerFee;

            vm.prank(USER1);
            buyOrder1 = ROUTER.submitBuyLimit(
                address(PAIR), _price105, _toBase(60), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            assertEq(PAIR.orderById(buyOrder1).amount, 0, "Step3: buyOrder1 fully filled");
            assertEq(PAIR.orderById(sellOrder1).amount, _toBase(40), "Step3: sellOrder1 has 40 remaining");
            assertEq(PAIR.baseReserve(), _toBase(90), "Step3: baseReserve = 90");
            assertEq(PAIR.matchedPrice(), _price100, "Step3: matchedPrice = 100");
            assertEq(BASE.balanceOf(USER1) - user1BaseBefore, _toBase(60), "Step3: USER1 received 60 BASE");

            uint256 feesCollected = QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore;
            assertEq(feesCollected, expectedTotalFee, "Step3: FeeCollector received total fees");
            totalFeesCollected += feesCollected;
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 4: BUY Market Order Matching Multiple Prices (USER4)
        // ═══════════════════════════════════════════════════════════════════════
        {
            uint256 user4BaseBefore = BASE.balanceOf(USER4);
            uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

            uint256 trade1Volume = _toTradeVolume(_price100, _toBase(40));
            uint256 trade2Volume = _toTradeVolume(_price110, _toBase(50));
            uint256 totalTradeVolume = trade1Volume + trade2Volume;
            uint256 buyerTakerFee = _calcFee(totalTradeVolume, BUYER_TAKER_FEE);

            uint256 quoteToSpend = totalTradeVolume + buyerTakerFee + _toQuote(100);

            vm.prank(USER4);
            ROUTER.submitBuyMarket(address(PAIR), quoteToSpend, 0);

            assertEq(PAIR.orderById(sellOrder1).amount, 0, "Step4: sellOrder1 fully filled");
            assertEq(PAIR.orderById(sellOrder2).amount, 0, "Step4: sellOrder2 fully filled");
            assertEq(PAIR.baseReserve(), 0, "Step4: baseReserve = 0");
            assertEq(PAIR.matchedPrice(), _price110, "Step4: matchedPrice = 110");
            assertEq(BASE.balanceOf(USER4) - user4BaseBefore, _toBase(90), "Step4: USER4 received 90 BASE");

            totalFeesCollected += QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore;
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 5: SELL Market Order on Empty BUY Book (USER2)
        // ═══════════════════════════════════════════════════════════════════════
        {
            uint256 user2BaseBefore = BASE.balanceOf(USER2);

            vm.prank(USER2);
            ROUTER.submitSellMarket(address(PAIR), _toBase(10), 0);

            assertEq(BASE.balanceOf(USER2), user2BaseBefore, "Step5: USER2 BASE unchanged (returned)");
            assertEq(PAIR.baseReserve(), 0, "Step5: baseReserve still 0");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 6: Additional Orders and Partial Cancel
        // ═══════════════════════════════════════════════════════════════════════
        uint256 sellOrder4;
        {
            uint256 buyVolume = _toTradeVolume(_price95, _toBase(100));
            uint256 buyerMakerFee = _calcFee(buyVolume, BUYER_MAKER_FEE);
            uint256 expectedDeposit = buyVolume + buyerMakerFee;

            uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);

            vm.prank(USER1);
            uint256 buyOrder3 = ROUTER.submitBuyLimit(
                address(PAIR), _price95, _toBase(100), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            assertEq(PAIR.quoteReserve(), expectedDeposit, "Step6: quoteReserve = deposit with fee");
            assertEq(user1QuoteBefore - QUOTE.balanceOf(USER1), expectedDeposit, "Step6: USER1 paid deposit with fee");

            vm.prank(USER2);
            sellOrder4 = ROUTER.submitSellLimit(
                address(PAIR), _price120, _toBase(30), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            assertEq(PAIR.baseReserve(), _toBase(30), "Step6: baseReserve = 30");

            // Cancel buyOrder3
            uint256 user1QuoteBeforeCancel = QUOTE.balanceOf(USER1);
            uint256[] memory cancelIds = new uint256[](1);
            cancelIds[0] = buyOrder3;

            vm.prank(USER1);
            ROUTER.cancelOrder(address(PAIR), cancelIds);

            assertEq(QUOTE.balanceOf(USER1) - user1QuoteBeforeCancel, expectedDeposit, "Step6: Full refund on cancel");
            assertEq(PAIR.quoteReserve(), 0, "Step6: quoteReserve = 0 after cancel");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 7: Cancel All Remaining Orders and Final Verification
        // ═══════════════════════════════════════════════════════════════════════
        {
            uint256 user2BaseBefore = BASE.balanceOf(USER2);

            uint256[] memory cancelIds = new uint256[](1);
            cancelIds[0] = sellOrder4;

            vm.prank(USER2);
            ROUTER.cancelOrder(address(PAIR), cancelIds);

            assertEq(BASE.balanceOf(USER2) - user2BaseBefore, _toBase(30), "Step7: USER2 received BASE refund");
            assertEq(PAIR.baseReserve(), 0, "Final: baseReserve = 0");
            assertEq(PAIR.quoteReserve(), 0, "Final: quoteReserve = 0");
        }

        // Final Fee Verification
        uint256 actualFeesCollected = QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorInit;
        assertEq(actualFeesCollected, totalFeesCollected, "Final: Total fees match");

        console.log("=== V2Compat E2E Scenario Complete ===");
        console.log("Total fees collected:", totalFeesCollected);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Additional Scenario Tests
    // ═══════════════════════════════════════════════════════════════════════════

    /// @notice Test sell taker matching multiple buy makers at same price
    function test_e2e_v2compat_sellTakerMultipleMakers() external {
        // USER1 and USER3: Place BUY orders at same price
        vm.prank(USER1);
        uint256 buyId1 = ROUTER.submitBuyLimit(
            address(PAIR), _price100, _toBase(50), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        vm.prank(USER3);
        uint256 buyId2 = ROUTER.submitBuyLimit(
            address(PAIR), _price100, _toBase(50), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 totalTradeVolume = _toTradeVolume(_price100, _toBase(100));
        uint256 sellerTakerFee = _calcFee(totalTradeVolume, SELLER_TAKER_FEE);
        uint256 buyerMakerFee = _calcFee(totalTradeVolume, BUYER_MAKER_FEE);

        uint256 sellerQuoteBefore = QUOTE.balanceOf(USER2);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        // USER2: SELL market order matching both
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(100), 0);

        uint256 sellerQuoteReceived = QUOTE.balanceOf(USER2) - sellerQuoteBefore;
        assertEq(sellerQuoteReceived, totalTradeVolume - sellerTakerFee, "Seller received net");

        uint256 feesCollected = QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore;
        assertEq(feesCollected, sellerTakerFee + buyerMakerFee, "Correct fees collected");

        assertEq(PAIR.orderById(buyId1).amount, 0, "buyId1 filled");
        assertEq(PAIR.orderById(buyId2).amount, 0, "buyId2 filled");
    }

    /// @notice Test buy taker matching multiple sell makers at different prices
    function test_e2e_v2compat_buyTakerMultipleMakers() external {
        // USER2 and USER3: Place SELL orders at different prices
        vm.prank(USER2);
        uint256 sellId1 = ROUTER.submitSellLimit(
            address(PAIR), _price100, _toBase(50), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        vm.prank(USER3);
        uint256 sellId2 = ROUTER.submitSellLimit(
            address(PAIR), _price105, _toBase(50), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 trade1Volume = _toTradeVolume(_price100, _toBase(50));
        uint256 trade2Volume = _toTradeVolume(_price105, _toBase(50));
        uint256 totalTradeVolume = trade1Volume + trade2Volume;

        uint256 seller1MakerFee = _calcFee(trade1Volume, SELLER_MAKER_FEE);
        uint256 seller2MakerFee = _calcFee(trade2Volume, SELLER_MAKER_FEE);
        uint256 buyerTakerFee = _calcFee(totalTradeVolume, BUYER_TAKER_FEE);

        uint256 user2QuoteBefore = QUOTE.balanceOf(USER2);
        uint256 user3QuoteBefore = QUOTE.balanceOf(USER3);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        uint256 quoteToSpend = totalTradeVolume + buyerTakerFee + _toQuote(10);
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteToSpend, 0);

        assertEq(QUOTE.balanceOf(USER2) - user2QuoteBefore, trade1Volume - seller1MakerFee, "USER2 received net");
        assertEq(QUOTE.balanceOf(USER3) - user3QuoteBefore, trade2Volume - seller2MakerFee, "USER3 received net");

        uint256 feesCollected = QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore;
        assertEq(feesCollected, seller1MakerFee + seller2MakerFee + buyerTakerFee, "Correct fees collected");

        assertEq(PAIR.orderById(sellId1).amount, 0, "sellId1 filled");
        assertEq(PAIR.orderById(sellId2).amount, 0, "sellId2 filled");
    }

    /// @notice Test partial fill followed by cancel with maker fee refund
    function test_e2e_v2compat_partialFillAndCancel() external {
        uint256 buyVolume = _toTradeVolume(_price100, _toBase(100));
        uint256 buyerMakerFee = _calcFee(buyVolume, BUYER_MAKER_FEE);
        uint256 expectedDeposit = buyVolume + buyerMakerFee;

        uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);

        vm.prank(USER1);
        uint256 buyId = ROUTER.submitBuyLimit(
            address(PAIR), _price100, _toBase(100), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        assertEq(user1QuoteBefore - QUOTE.balanceOf(USER1), expectedDeposit, "Deposit taken");

        // USER2: SELL 30 @ 100 (partial fill)
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(30), 0);

        assertEq(PAIR.orderById(buyId).amount, _toBase(70), "70 remaining after partial fill");

        uint256 remainingVolume = _toTradeVolume(_price100, _toBase(70));
        uint256 remainingMakerFee = _calcFee(remainingVolume, BUYER_MAKER_FEE);
        uint256 expectedRemainingReserve = remainingVolume + remainingMakerFee;

        assertEq(PAIR.quoteReserve(), expectedRemainingReserve, "Correct remaining reserve");

        uint256 user1QuoteBeforeCancel = QUOTE.balanceOf(USER1);
        uint256[] memory cancelIds = new uint256[](1);
        cancelIds[0] = buyId;

        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), cancelIds);

        uint256 user1QuoteRefund = QUOTE.balanceOf(USER1) - user1QuoteBeforeCancel;
        assertEq(user1QuoteRefund, expectedRemainingReserve, "Full remaining refunded");
        assertEq(PAIR.quoteReserve(), 0, "Reserve 0 after cancel");
    }

    /// @notice Test that maker fees are correctly charged on match
    function test_e2e_v2compat_makerFeeOnMatch() external {
        uint256 amount = _toBase(10);
        uint256 tradeVolume = _toTradeVolume(_price100, amount);

        vm.prank(USER2);
        uint256 sellId = ROUTER.submitSellLimit(
            address(PAIR), _price100, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        assertEq(PAIR.orderById(sellId).feeBps, SELLER_MAKER_FEE, "Seller maker fee set");

        uint256 user2QuoteBefore = QUOTE.balanceOf(USER2);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        uint256 buyerTakerFee = _calcFee(tradeVolume, BUYER_TAKER_FEE);
        uint256 sellerMakerFee = _calcFee(tradeVolume, SELLER_MAKER_FEE);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), _price100, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 user2QuoteReceived = QUOTE.balanceOf(USER2) - user2QuoteBefore;
        assertEq(user2QuoteReceived, tradeVolume - sellerMakerFee, "Seller received net");

        uint256 feesCollected = QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore;
        assertEq(feesCollected, buyerTakerFee + sellerMakerFee, "Total fees collected");
    }
}
