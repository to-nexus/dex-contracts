// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {CrossDexImplV3} from "../src/CrossDexImplV3.sol";
import {CrossDexRouterV3} from "../src/CrossDexRouterV3.sol";
import {FeeControllerV3Split} from "../src/FeeControllerV3Split.sol";
import {MarketImplV3} from "../src/MarketImplV3.sol";
import {PairImplV3} from "../src/PairImplV3.sol";
import {WETH} from "../src/WETH.sol";

import {BPS_DENOMINATOR} from "../src/interfaces/IFeeController.sol";
import {IPairV3} from "../src/interfaces/IPairV3.sol";

import {T20} from "./mock/T20.sol";

/// @title DEXV3E2EScenarioV3SplitTest
/// @notice End-to-end scenario tests for V3 matching engine with FeeControllerV3Split (3-way split)
contract DEXV3E2EScenarioV3SplitTest is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    address public constant CREATOR = address(bytes20("CREATOR"));
    address public constant USER1 = address(bytes20("USER1"));
    address public constant USER2 = address(bytes20("USER2"));
    address public constant USER3 = address(bytes20("USER3"));
    address public constant USER4 = address(bytes20("USER4"));

    CrossDexImplV3 public CROSS_DEX;
    CrossDexRouterV3 public ROUTER;
    WETH public CROSS;

    IERC20 public QUOTE;
    IERC20 public BASE;

    MarketImplV3 public MARKET;
    PairImplV3 public PAIR;
    FeeControllerV3Split public FEE_CONTROLLER;

    uint256 public FIND_PREV_PRICE_COUNT = type(uint256).max;
    uint256 public MAX_MATCH_COUNT = type(uint256).max;
    uint256 public CANCEL_LIMIT = type(uint256).max;

    // V3Split fee structure: taker-only with 3-way split
    uint32 public constant TAKER_FEE_BPS = 100; // 1%
    uint32 public constant CREATOR_SHARE_BPS = 3000; // 30% of taker fee
    uint32 public constant MAKER_REBATE_SHARE_BPS = 2000; // 20% of taker fee
    // System (feeCollector) gets remaining 50%

    uint256 public QUOTE_DECIMALS;
    uint256 public BASE_DECIMALS;

    uint256[2] internal _searchPrices;

    // Price constants (initialized in setUp to reduce stack depth)
    uint256 internal _price100;
    uint256 internal _price105;
    uint256 internal _price110;
    uint256 internal _price95;
    uint256 internal _price120;

    // Order ID tracking for E2E tests (contract-level to reduce stack depth)
    uint256 internal _sellOrder1;
    uint256 internal _sellOrder2;
    uint256 internal _buyOrder1;
    uint256 internal _buyOrder3;
    uint256 internal _sellOrder4;

    // Cumulative fee tracking
    uint256 internal _totalTakerFees;
    uint256 internal _totalCreatorFees;
    uint256 internal _totalMakerRebates;
    uint256 internal _totalSystemFees;

    function setUp() external {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");
        vm.label(CREATOR, "creator");
        vm.label(USER1, "user1");
        vm.label(USER2, "user2");
        vm.label(USER3, "user3");
        vm.label(USER4, "user4");

        vm.startPrank(OWNER);

        QUOTE_DECIMALS = 10 ** 18;
        BASE_DECIMALS = 10 ** 18;

        // Deploy FeeControllerV3Split implementation
        FEE_CONTROLLER = new FeeControllerV3Split();

        // Deploy impl contracts
        address routerImpl = address(new CrossDexRouterV3());
        address marketImpl = address(new MarketImplV3());
        address pairImpl = address(new PairImplV3());

        // Deploy cross dex
        address crossDexImpl = address(new CrossDexImplV3());
        ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
        CROSS_DEX = CrossDexImplV3(address(proxy));
        CROSS_DEX.initialize(
            OWNER,
            routerImpl,
            FIND_PREV_PRICE_COUNT,
            MAX_MATCH_COUNT,
            CANCEL_LIMIT,
            marketImpl,
            pairImpl,
            address(0) // tickSizeSetter
        );

        // Allow V3Split fee controller
        CROSS_DEX.setFeeControllerAllow(address(FEE_CONTROLLER), true);

        // Get contracts from CROSS_DEX
        ROUTER = CrossDexRouterV3(CROSS_DEX.ROUTER());
        CROSS = WETH(payable(address(ROUTER.CROSS())));

        // Deploy base and quote tokens
        QUOTE = new T20("QUOTE", "QUOTE", 18);
        BASE = new T20("BASE", "BASE", 18);

        // Create market with V3Split FeeController
        address market = CROSS_DEX.createMarket(OWNER, address(QUOTE), address(FEE_CONTROLLER), "");
        MARKET = MarketImplV3(market);

        // Encode V3Split fee controller init data
        bytes memory feeControllerInitData =
            abi.encode(FEE_COLLECTOR, CREATOR, TAKER_FEE_BPS, CREATOR_SHARE_BPS, MAKER_REBATE_SHARE_BPS);

        // Create pair with tick/lot sizes
        address pair =
            MARKET.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeControllerInitData);
        PAIR = PairImplV3(pair);

        // Setup initial balances
        QUOTE.transfer(USER1, _toQuote(50000));
        QUOTE.transfer(USER2, _toQuote(50000));
        QUOTE.transfer(USER3, _toQuote(50000));
        QUOTE.transfer(USER4, _toQuote(50000));
        BASE.transfer(USER1, _toBase(50000));
        BASE.transfer(USER2, _toBase(50000));
        BASE.transfer(USER3, _toBase(50000));
        BASE.transfer(USER4, _toBase(50000));

        vm.stopPrank();

        // Setup approvals for all users
        address[4] memory users = [USER1, USER2, USER3, USER4];
        for (uint256 i = 0; i < users.length; i++) {
            vm.prank(users[i]);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            vm.prank(users[i]);
            BASE.approve(address(ROUTER), type(uint256).max);
        }

        // Initialize price constants
        _price100 = _toQuote(100);
        _price105 = _toQuote(105);
        _price110 = _toQuote(110);
        _price95 = _toQuote(95);
        _price120 = _toQuote(120);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Helper Functions
    // ─────────────────────────────────────────────────────────────────────────────

    function _toBase(uint256 x) internal view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) internal view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }

    function _toTradeVolume(uint256 price, uint256 amount) internal view returns (uint256) {
        return Math.mulDiv(price, amount, BASE_DECIMALS);
    }

    function _calcFee(uint256 amount, uint32 bps) internal pure returns (uint256) {
        return Math.mulDiv(amount, bps, BPS_DENOMINATOR);
    }

    /// @notice Calculate V3Split fee distribution
    function _calcV3SplitFees(uint256 tradeVolume)
        internal
        pure
        returns (uint256 takerFee, uint256 creatorFee, uint256 makerRebate, uint256 systemFee)
    {
        takerFee = _calcFee(tradeVolume, TAKER_FEE_BPS);
        creatorFee = _calcFee(takerFee, CREATOR_SHARE_BPS);
        makerRebate = _calcFee(takerFee, MAKER_REBATE_SHARE_BPS);
        systemFee = takerFee - creatorFee - makerRebate;
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // E2E Full Scenario Test
    // ═══════════════════════════════════════════════════════════════════════════

    /// @notice Full E2E scenario testing order lifecycle with V3Split fees.
    /// @dev Steps: submit → match → partial fill → cancel → verify reserves zero
    ///      V3Split: taker pays fee, maker receives rebate immediately during match
    function test_e2e_v3split_fullScenario() external {
        // Reset order IDs and fee tracking (use contract-level storage)
        _sellOrder1 = 0;
        _sellOrder2 = 0;
        _buyOrder1 = 0;
        _buyOrder3 = 0;
        _sellOrder4 = 0;
        _totalTakerFees = 0;
        _totalCreatorFees = 0;
        _totalMakerRebates = 0;
        _totalSystemFees = 0;

        // ═══════════════════════════════════════════════════════════════════════
        // Step 1: Initial State Verification
        // ═══════════════════════════════════════════════════════════════════════
        uint256 user1BaseBefore = BASE.balanceOf(USER1);
        uint256 user2BaseBefore = BASE.balanceOf(USER2);
        uint256 user3BaseBefore = BASE.balanceOf(USER3);
        uint256 feeCollectorInit = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 creatorInit = QUOTE.balanceOf(CREATOR);

        assertEq(PAIR.baseReserve(), 0, "Step1: Initial baseReserve should be 0");
        assertEq(PAIR.quoteReserve(), 0, "Step1: Initial quoteReserve should be 0");

        // ═══════════════════════════════════════════════════════════════════════
        // Step 2: Multiple SELL Limit Orders (USER2, USER3) - Makers
        // V3Split: Maker fee is always 0
        // ═══════════════════════════════════════════════════════════════════════
        {
            // USER2: SELL 100 BASE @ price 100
            vm.prank(USER2);
            _sellOrder1 = ROUTER.submitSellLimit(
                address(PAIR), _price100, _toBase(100), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            // Verify order was created with 0 maker fee
            IPairV3.Order memory order = PAIR.orderById(_sellOrder1);
            assertEq(order.amount, _toBase(100), "Step2: _sellOrder1 amount");
            assertEq(order.feeBps, 0, "Step2: _sellOrder1 maker fee = 0 (V3Split)");

            // USER3: SELL 50 BASE @ price 110
            vm.prank(USER3);
            _sellOrder2 = ROUTER.submitSellLimit(
                address(PAIR), _price110, _toBase(50), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            order = PAIR.orderById(_sellOrder2);
            assertEq(order.amount, _toBase(50), "Step2: _sellOrder2 amount");
            assertEq(order.feeBps, 0, "Step2: _sellOrder2 maker fee = 0");

            // Verify reserves
            assertEq(PAIR.baseReserve(), _toBase(150), "Step2: baseReserve = 150");
            assertEq(PAIR.quoteReserve(), 0, "Step2: quoteReserve = 0");

            // Verify user balances
            assertEq(BASE.balanceOf(USER2), user2BaseBefore - _toBase(100), "Step2: USER2 BASE decreased");
            assertEq(BASE.balanceOf(USER3), user3BaseBefore - _toBase(50), "Step2: USER3 BASE decreased");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 3: BUY Limit Order Partial Match (USER1) - Taker
        // USER1 BUY 60 BASE @ 105 → matches _sellOrder1 @ 100
        // V3Split: USER1 pays taker fee, USER2 (maker) receives rebate IMMEDIATELY
        // ═══════════════════════════════════════════════════════════════════════
        {
            uint256 user1QuoteBeforeStep3 = QUOTE.balanceOf(USER1);
            uint256 user2QuoteBeforeStep3 = QUOTE.balanceOf(USER2);
            uint256 feeCollectorBeforeStep3 = QUOTE.balanceOf(FEE_COLLECTOR);
            uint256 creatorBeforeStep3 = QUOTE.balanceOf(CREATOR);

            // Calculate expected fees
            uint256 tradeVolume = _toTradeVolume(_price100, _toBase(60));
            (uint256 takerFee, uint256 creatorFee, uint256 makerRebate, uint256 systemFee) =
                _calcV3SplitFees(tradeVolume);

            vm.prank(USER1);
            _buyOrder1 = ROUTER.submitBuyLimit(
                address(PAIR), _price105, _toBase(60), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            // _buyOrder1 should be fully filled
            IPairV3.Order memory order = PAIR.orderById(_buyOrder1);
            assertEq(order.amount, 0, "Step3: _buyOrder1 should be fully filled");

            // _sellOrder1 should have 40 remaining
            order = PAIR.orderById(_sellOrder1);
            assertEq(order.amount, _toBase(40), "Step3: _sellOrder1 should have 40 remaining");

            // Verify reserves
            assertEq(PAIR.baseReserve(), _toBase(90), "Step3: baseReserve = 90");
            assertEq(PAIR.quoteReserve(), 0, "Step3: quoteReserve = 0 (buy fully filled)");

            // Verify matchedPrice updated
            assertEq(PAIR.matchedPrice(), _price100, "Step3: matchedPrice = 100");

            // Verify balances
            // USER1 (buyer taker): receives 60 BASE, pays tradeVolume + takerFee
            assertEq(BASE.balanceOf(USER1) - user1BaseBefore, _toBase(60), "Step3: USER1 received 60 BASE");
            uint256 user1QuotePaid = user1QuoteBeforeStep3 - QUOTE.balanceOf(USER1);
            assertEq(user1QuotePaid, tradeVolume + takerFee, "Step3: USER1 paid volume + taker fee");

            // USER2 (seller maker): receives tradeVolume + makerRebate (NO fee deduction!)
            uint256 user2QuoteReceived = QUOTE.balanceOf(USER2) - user2QuoteBeforeStep3;
            assertEq(user2QuoteReceived, tradeVolume + makerRebate, "Step3: USER2 received volume + rebate");

            // Creator: receives creatorFee
            uint256 creatorReceived = QUOTE.balanceOf(CREATOR) - creatorBeforeStep3;
            assertEq(creatorReceived, creatorFee, "Step3: CREATOR received creator fee");

            // FeeCollector: receives systemFee only
            uint256 systemReceived = QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBeforeStep3;
            assertEq(systemReceived, systemFee, "Step3: FeeCollector received system fee");

            // Track totals
            _totalTakerFees += takerFee;
            _totalCreatorFees += creatorFee;
            _totalMakerRebates += makerRebate;
            _totalSystemFees += systemFee;
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 4: BUY Market Order Matching Multiple Prices (USER4) - Taker
        // ═══════════════════════════════════════════════════════════════════════
        // Step 4a: Execute buy market and track fees
        {
            uint256 user4BaseBefore = BASE.balanceOf(USER4);
            uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);
            uint256 creatorBefore = QUOTE.balanceOf(CREATOR);

            // _sellOrder1: 40 @ 100, _sellOrder2: 50 @ 110
            uint256 trade1Vol = _toTradeVolume(_price100, _toBase(40));
            uint256 trade2Vol = _toTradeVolume(_price110, _toBase(50));
            uint256 totalVol = trade1Vol + trade2Vol;

            (uint256 takerFee, uint256 creatorFee, uint256 makerRebate, uint256 systemFee) = _calcV3SplitFees(totalVol);

            vm.prank(USER4);
            ROUTER.submitBuyMarket(address(PAIR), totalVol + takerFee + _toQuote(100), 0);

            // Verify orders filled
            assertEq(PAIR.orderById(_sellOrder1).amount, 0, "Step4: _sellOrder1 filled");
            assertEq(PAIR.orderById(_sellOrder2).amount, 0, "Step4: _sellOrder2 filled");

            // Verify reserves and price
            assertEq(PAIR.baseReserve(), 0, "Step4: baseReserve = 0");
            assertEq(PAIR.quoteReserve(), 0, "Step4: quoteReserve = 0");
            assertEq(PAIR.matchedPrice(), _price110, "Step4: matchedPrice = 110");

            // Verify USER4 received BASE
            assertEq(BASE.balanceOf(USER4) - user4BaseBefore, _toBase(90), "Step4: USER4 received 90 BASE");

            // Track totals
            _totalTakerFees += takerFee;
            _totalCreatorFees += creatorFee;
            _totalMakerRebates += makerRebate;
            _totalSystemFees += systemFee;

            // Verify fee distribution
            assertEq(QUOTE.balanceOf(CREATOR) - creatorBefore, creatorFee, "Step4: CREATOR received fee");
            assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, systemFee, "Step4: FeeCollector received fee");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 5: SELL Market Order on Empty BUY Book (USER2)
        // ═══════════════════════════════════════════════════════════════════════
        {
            uint256 user2BaseBeforeStep5 = BASE.balanceOf(USER2);

            vm.prank(USER2);
            ROUTER.submitSellMarket(address(PAIR), _toBase(10), 0);

            // No BUY orders, so BASE should be returned
            assertEq(BASE.balanceOf(USER2), user2BaseBeforeStep5, "Step5: USER2 BASE unchanged (returned)");

            // Reserves unchanged
            assertEq(PAIR.baseReserve(), 0, "Step5: baseReserve still 0");
            assertEq(PAIR.quoteReserve(), 0, "Step5: quoteReserve still 0");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 6: Additional Orders and Partial Cancel
        // V3Split: No maker fee, so quoteReserve = just the trade volume
        // ═══════════════════════════════════════════════════════════════════════
        {
            // USER1: BUY 100 @ 95 (maker, no fee)
            uint256 buyVolume = _toTradeVolume(_price95, _toBase(100));
            // V3Split: maker fee = 0, so deposit = just volume
            uint256 expectedDeposit = buyVolume;

            uint256 user1QuoteBeforeStep6 = QUOTE.balanceOf(USER1);

            vm.prank(USER1);
            _buyOrder3 = ROUTER.submitBuyLimit(
                address(PAIR), _price95, _toBase(100), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            // Verify order created with 0 maker fee
            IPairV3.Order memory order = PAIR.orderById(_buyOrder3);
            assertEq(order.feeBps, 0, "Step6: _buyOrder3 maker fee = 0 (V3Split)");
            assertEq(order.amount, _toBase(100), "Step6: _buyOrder3 amount");

            // Verify quote reserve
            assertEq(PAIR.quoteReserve(), expectedDeposit, "Step6: quoteReserve = volume only (no maker fee)");

            // Verify USER1 paid deposit
            uint256 user1QuotePaid = user1QuoteBeforeStep6 - QUOTE.balanceOf(USER1);
            assertEq(user1QuotePaid, expectedDeposit, "Step6: USER1 paid deposit");

            // USER2: SELL 30 @ 120
            vm.prank(USER2);
            _sellOrder4 = ROUTER.submitSellLimit(
                address(PAIR), _price120, _toBase(30), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            assertEq(PAIR.baseReserve(), _toBase(30), "Step6: baseReserve = 30");

            // Cancel _buyOrder3
            uint256 user1QuoteBeforeCancel = QUOTE.balanceOf(USER1);
            uint256[] memory cancelIds = new uint256[](1);
            cancelIds[0] = _buyOrder3;

            vm.prank(USER1);
            ROUTER.cancelOrder(address(PAIR), cancelIds);

            // Verify full refund (V3Split: no maker fee, so full volume returned)
            uint256 user1QuoteAfterCancel = QUOTE.balanceOf(USER1);
            assertEq(user1QuoteAfterCancel - user1QuoteBeforeCancel, expectedDeposit, "Step6: Full refund on cancel");

            // Verify quoteReserve reduced
            assertEq(PAIR.quoteReserve(), 0, "Step6: quoteReserve = 0 after cancel");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Step 7: Cancel All Remaining Orders and Final Verification
        // ═══════════════════════════════════════════════════════════════════════
        {
            // Only _sellOrder4 remains (30 BASE @ 120)
            uint256 user2BaseBeforeCancel = BASE.balanceOf(USER2);

            uint256[] memory cancelIds = new uint256[](1);
            cancelIds[0] = _sellOrder4;

            vm.prank(USER2);
            ROUTER.cancelOrder(address(PAIR), cancelIds);

            // Verify refund
            assertEq(BASE.balanceOf(USER2) - user2BaseBeforeCancel, _toBase(30), "Step7: USER2 received BASE refund");

            // Final state: all reserves should be 0
            assertEq(PAIR.baseReserve(), 0, "Final: baseReserve = 0");
            assertEq(PAIR.quoteReserve(), 0, "Final: quoteReserve = 0");
        }

        // ═══════════════════════════════════════════════════════════════════════
        // Final Fee Verification
        // ═══════════════════════════════════════════════════════════════════════
        uint256 actualCreatorFees = QUOTE.balanceOf(CREATOR) - creatorInit;
        uint256 actualSystemFees = QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorInit;

        assertEq(actualCreatorFees, _totalCreatorFees, "Final: Creator fees match");
        assertEq(actualSystemFees, _totalSystemFees, "Final: System fees match");

        console.log("=== V3Split E2E Scenario Complete ===");
        console.log("Total taker fees:", _totalTakerFees);
        console.log("Creator fees:", _totalCreatorFees);
        console.log("Maker rebates:", _totalMakerRebates);
        console.log("System fees:", _totalSystemFees);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Additional Scenario Tests
    // ═══════════════════════════════════════════════════════════════════════════

    /// @notice Test 3-way fee split with single match
    function test_e2e_v3split_threeWayFeeSplit() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 tradeVolume = _toTradeVolume(price, amount);

        // USER1: Place BUY limit order (maker)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 makerQuoteBefore = QUOTE.balanceOf(USER1);
        uint256 sellerQuoteBefore = QUOTE.balanceOf(USER2);
        uint256 creatorBefore = QUOTE.balanceOf(CREATOR);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        // Calculate expected fees
        (uint256 takerFee, uint256 creatorFee, uint256 makerRebate, uint256 systemFee) = _calcV3SplitFees(tradeVolume);

        // USER2: SELL market order (taker)
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Verify 3-way split:
        // 1. Maker (USER1) received rebate immediately
        assertEq(QUOTE.balanceOf(USER1) - makerQuoteBefore, makerRebate, "Maker received rebate");

        // 2. Seller (taker, USER2) received (tradeVolume - takerFee)
        assertEq(QUOTE.balanceOf(USER2) - sellerQuoteBefore, tradeVolume - takerFee, "Seller received net");

        // 3. Creator received creatorFee
        assertEq(QUOTE.balanceOf(CREATOR) - creatorBefore, creatorFee, "Creator received fee");

        // 4. FeeCollector received systemFee
        assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, systemFee, "FeeCollector received fee");

        // Verify sum: takerFee = creatorFee + makerRebate + systemFee
        assertEq(takerFee, creatorFee + makerRebate + systemFee, "Fee split sums correctly");
    }

    /// @notice Test maker rebate is paid immediately during matching (not at settlement)
    function test_e2e_v3split_makerRebateImmediatePayment() external {
        uint256 price = _toQuote(100);

        // USER1 and USER3: Place BUY orders at same price (both makers)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(50), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        vm.prank(USER3);
        ROUTER.submitBuyLimit(
            address(PAIR), price, _toBase(50), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 maker1QuoteBefore = QUOTE.balanceOf(USER1);
        uint256 maker2QuoteBefore = QUOTE.balanceOf(USER3);

        // Calculate expected rebates per maker
        uint256 tradeVolumePerMaker = _toTradeVolume(price, _toBase(50));
        uint256 rebatePerMaker = _calcFee(_calcFee(tradeVolumePerMaker, TAKER_FEE_BPS), MAKER_REBATE_SHARE_BPS);

        // USER2: SELL market order matching both
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(100), 0);

        // Each maker should have received their rebate immediately during match
        assertEq(QUOTE.balanceOf(USER1) - maker1QuoteBefore, rebatePerMaker, "Maker1 received rebate immediately");
        assertEq(QUOTE.balanceOf(USER3) - maker2QuoteBefore, rebatePerMaker, "Maker2 received rebate immediately");
    }

    /// @notice Test creator share accumulation across multiple trades
    function test_e2e_v3split_creatorShareAccumulation() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 tradeVolume = _toTradeVolume(price, amount);

        uint256 creatorInit = QUOTE.balanceOf(CREATOR);

        // Trade 1
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Trade 2
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        uint256 takerFee = _calcFee(tradeVolume, TAKER_FEE_BPS);
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), tradeVolume + takerFee, 0);

        // Calculate expected total creator fees (2 trades)
        (, uint256 creatorFeePerTrade,,) = _calcV3SplitFees(tradeVolume);
        uint256 expectedTotalCreatorFee = creatorFeePerTrade * 2;

        assertEq(
            QUOTE.balanceOf(CREATOR) - creatorInit, expectedTotalCreatorFee, "Creator accumulated fees from 2 trades"
        );
    }

    /// @notice Test that cancel returns full amount (no maker fee deduction in V3Split)
    function test_e2e_v3split_cancelNoFeeDeduction() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(100);
        uint256 tradeVolume = _toTradeVolume(price, amount);

        uint256 userQuoteBefore = QUOTE.balanceOf(USER1);

        // USER1: BUY limit order (maker)
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify deposit equals trade volume (no maker fee in V3Split)
        uint256 depositPaid = userQuoteBefore - QUOTE.balanceOf(USER1);
        assertEq(depositPaid, tradeVolume, "Deposit = volume (no maker fee)");

        // Verify order.feeBps = 0
        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.feeBps, 0, "Maker fee = 0");

        // Cancel order
        uint256[] memory cancelIds = new uint256[](1);
        cancelIds[0] = orderId;
        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), cancelIds);

        // Verify full refund
        assertEq(QUOTE.balanceOf(USER1), userQuoteBefore, "Full refund on cancel");
    }
}
