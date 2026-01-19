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

import {BPS_DENOMINATOR, IFeeController} from "../src/interfaces/IFeeController.sol";
import {IPairV3} from "../src/interfaces/IPairV3.sol";

import {T20} from "./mock/T20.sol";

/// @title DEXV3FeeControllerV3SplitTest
/// @notice Tests for V3 fee controller with 3-way split (creator/maker rebate/system)
contract DEXV3FeeControllerV3SplitTest is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    address public constant CREATOR = address(bytes20("CREATOR"));
    address public constant USER1 = address(bytes20("USER1"));
    address public constant USER2 = address(bytes20("USER2"));

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

    // V3Split default fees (taker only, 1% = 100 bps)
    uint32 public TAKER_FEE_BPS = 100; // 1%
    uint32 public CREATOR_SHARE_BPS = 3000; // 30% of taker fee
    uint32 public MAKER_REBATE_SHARE_BPS = 2000; // 20% of taker fee
    // System (feeCollector) gets remaining 50%

    uint256 public QUOTE_DECIMALS;
    uint256 public BASE_DECIMALS;

    uint256[2] internal _searchPrices;

    function setUp() external {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");
        vm.label(CREATOR, "creator");
        vm.label(USER1, "user1");
        vm.label(USER2, "user2");

        vm.startPrank(OWNER);

        QUOTE_DECIMALS = 10 ** 18;
        BASE_DECIMALS = 10 ** 18;

        // Deploy FeeController implementation
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

        // Allow fee controller
        CROSS_DEX.setFeeControllerAllow(address(FEE_CONTROLLER), true);

        // Get contracts from CROSS_DEX
        ROUTER = CrossDexRouterV3(CROSS_DEX.ROUTER());
        CROSS = WETH(payable(address(ROUTER.CROSS())));

        // Deploy base and quote tokens
        QUOTE = new T20("QUOTE", "QUOTE", 18);
        BASE = new T20("BASE", "BASE", 18);

        // Create market with FeeController
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
        BASE.transfer(USER1, _toBase(50000));
        BASE.transfer(USER2, _toBase(50000));

        vm.stopPrank();

        // Setup approvals
        vm.prank(USER1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER1);
        BASE.approve(address(ROUTER), type(uint256).max);

        vm.prank(USER2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER2);
        BASE.approve(address(ROUTER), type(uint256).max);
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

    // ─────────────────────────────────────────────────────────────────────────────
    // Initialization Tests
    // ─────────────────────────────────────────────────────────────────────────────

    function test_initialization() external {
        assertEq(address(PAIR.feeController()), address(FEE_CONTROLLER), "Pair should have feeController set");

        // Verify calcBuyVolumeWithFee works correctly
        uint256 testVolume = _toQuote(1000);
        uint256 expectedWithFee = testVolume + _calcFee(testVolume, TAKER_FEE_BPS);
        assertEq(PAIR.calcBuyVolumeWithFee(testVolume), expectedWithFee, "Taker fee should be applied");
    }

    function test_makerFee_always_zero() external {
        // Verify maker fee bps are 0
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // Place BUY limit order (maker)
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify order.feeBps is 0 (buyer maker)
        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.feeBps, 0, "Buyer maker fee should be 0");

        // Place SELL limit order (maker)
        vm.prank(USER2);
        orderId = ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify order.feeBps is 0 (seller maker)
        order = PAIR.orderById(orderId);
        assertEq(order.feeBps, 0, "Seller maker fee should be 0");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // SELL Taker Tests (seller takes from BUY order book)
    // ─────────────────────────────────────────────────────────────────────────────

    function test_sell_taker_fee_distribution() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        // USER1: Place a BUY limit order (maker)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // Record balances before
        uint256 sellerQuoteBefore = QUOTE.balanceOf(USER2);
        uint256 makerQuoteBefore = QUOTE.balanceOf(USER1); // maker gets rebate
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 creatorBefore = QUOTE.balanceOf(CREATOR);

        // USER2: Sell market order (taker)
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Calculate expected values
        uint256 takerFee = _calcFee(quoteVolume, TAKER_FEE_BPS);
        uint256 creatorFee = _calcFee(takerFee, CREATOR_SHARE_BPS);
        uint256 makerRebate = _calcFee(takerFee, MAKER_REBATE_SHARE_BPS);
        uint256 systemFee = takerFee - creatorFee - makerRebate;

        // Verify seller received (quoteVolume - takerFee)
        uint256 sellerQuoteAfter = QUOTE.balanceOf(USER2);
        assertEq(sellerQuoteAfter - sellerQuoteBefore, quoteVolume - takerFee, "Seller should receive net QUOTE");

        // Verify maker received rebate
        uint256 makerQuoteAfter = QUOTE.balanceOf(USER1);
        assertEq(makerQuoteAfter - makerQuoteBefore, makerRebate, "Maker should receive rebate");

        // Verify creator received their share
        uint256 creatorAfter = QUOTE.balanceOf(CREATOR);
        assertEq(creatorAfter - creatorBefore, creatorFee, "Creator should receive creator fee");

        // Verify feeCollector received system fee
        uint256 feeCollectorAfter = QUOTE.balanceOf(FEE_COLLECTOR);
        assertEq(feeCollectorAfter - feeCollectorBefore, systemFee, "FeeCollector should receive system fee");
    }

    function test_sell_taker_multiple_makers_rebate() external {
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
        uint256 makerQuoteBefore = QUOTE.balanceOf(USER1);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 creatorBefore = QUOTE.balanceOf(CREATOR);

        // USER2: Sell market order matching both
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), totalAmount, 0);

        // Calculate expected fees
        uint256 takerFee = _calcFee(quoteVolumeTotal, TAKER_FEE_BPS);
        uint256 creatorFee = _calcFee(takerFee, CREATOR_SHARE_BPS);
        uint256 makerRebate = _calcFee(takerFee, MAKER_REBATE_SHARE_BPS);
        uint256 systemFee = takerFee - creatorFee - makerRebate;

        // Verify maker got total rebate (both matches go to same maker USER1)
        uint256 makerQuoteAfter = QUOTE.balanceOf(USER1);
        assertEq(makerQuoteAfter - makerQuoteBefore, makerRebate, "Maker should receive total rebate");

        assertEq(QUOTE.balanceOf(CREATOR) - creatorBefore, creatorFee, "Creator should receive fee");
        assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, systemFee, "FeeCollector should receive fee");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // BUY Taker Tests (buyer takes from SELL order book)
    // ─────────────────────────────────────────────────────────────────────────────

    function test_buy_taker_fee_distribution() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);
        uint256 takerFee = _calcFee(quoteVolume, TAKER_FEE_BPS);

        // USER2: Place a SELL limit order (maker)
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Record balances before
        uint256 makerQuoteBefore = QUOTE.balanceOf(USER2); // seller maker gets rebate
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 creatorBefore = QUOTE.balanceOf(CREATOR);

        // USER1: Buy market order (taker)
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteVolume + takerFee, 0);

        // Calculate expected values
        uint256 creatorFee = _calcFee(takerFee, CREATOR_SHARE_BPS);
        uint256 makerRebate = _calcFee(takerFee, MAKER_REBATE_SHARE_BPS);
        uint256 systemFee = takerFee - creatorFee - makerRebate;

        // Verify maker (seller) received quoteVolume + rebate (no maker fee deduction)
        uint256 makerQuoteAfter = QUOTE.balanceOf(USER2);
        assertEq(
            makerQuoteAfter - makerQuoteBefore, quoteVolume + makerRebate, "Seller maker should receive full + rebate"
        );

        // Verify distribution
        assertEq(QUOTE.balanceOf(CREATOR) - creatorBefore, creatorFee, "Creator should receive fee");
        assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, systemFee, "FeeCollector should receive fee");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Event Tests
    // ─────────────────────────────────────────────────────────────────────────────

    function test_v3_fees_settled_event_emitted() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        // Place BUY limit order
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        // Calculate expected fees
        uint256 takerFee = _calcFee(quoteVolume, TAKER_FEE_BPS);
        uint256 creatorFee = _calcFee(takerFee, CREATOR_SHARE_BPS);
        uint256 makerRebate = _calcFee(takerFee, MAKER_REBATE_SHARE_BPS);
        uint256 systemFee = takerFee - creatorFee - makerRebate;

        // Expect FeeControllerV3FeesSettled event
        vm.expectEmit(true, true, true, true, address(PAIR));
        emit FeeControllerV3Split.FeeControllerV3FeesSettled(
            2, // takerId (sell order is 2, buy order is 1)
            takerFee,
            creatorFee,
            systemFee,
            makerRebate,
            CREATOR,
            FEE_COLLECTOR
        );

        // Sell market order
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Transient Storage Tests
    // ─────────────────────────────────────────────────────────────────────────────

    function test_transient_storage_resets_after_settle() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        // First trade
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Second trade should work without takerId mismatch
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 quoteVolume = _toTradeVolume(price, amount);
        uint256 takerFee = _calcFee(quoteVolume, TAKER_FEE_BPS);

        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteVolume + takerFee, 0);

        // If we got here without revert, transient storage was properly reset
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Cancel/Refund Tests (maker fee is 0, so full refund expected)
    // ─────────────────────────────────────────────────────────────────────────────

    function test_buy_maker_cancel_full_refund() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);
        // Since maker fee is 0, deposit = quoteVolume
        uint256 expectedDeposit = quoteVolume;

        uint256 buyerQuoteBefore = QUOTE.balanceOf(USER1);

        // Place BUY limit order
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // Verify deposit
        uint256 buyerQuoteAfterSubmit = QUOTE.balanceOf(USER1);
        assertEq(
            buyerQuoteBefore - buyerQuoteAfterSubmit, expectedDeposit, "Deposit should equal volume (no maker fee)"
        );

        // Cancel
        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;
        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), orderIds);

        // Verify full refund
        uint256 buyerQuoteAfterCancel = QUOTE.balanceOf(USER1);
        assertEq(buyerQuoteAfterCancel, buyerQuoteBefore, "Full deposit should be refunded");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Configuration Change Tests
    // ─────────────────────────────────────────────────────────────────────────────

    function test_different_fee_configuration() external {
        // Change fee configuration
        uint32 newTakerFeeBps = 200; // 2%
        uint32 newCreatorShareBps = 4000; // 40%
        uint32 newMakerRebateShareBps = 1000; // 10%

        bytes memory newFeeData =
            abi.encode(FEE_COLLECTOR, CREATOR, newTakerFeeBps, newCreatorShareBps, newMakerRebateShareBps);
        address[] memory pairs = new address[](1);
        pairs[0] = address(PAIR);
        vm.prank(OWNER);
        MARKET.setFeeController(pairs, address(FEE_CONTROLLER), newFeeData);

        // Test with new configuration
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 makerQuoteBefore = QUOTE.balanceOf(USER1);
        uint256 creatorBefore = QUOTE.balanceOf(CREATOR);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Calculate expected with new rates
        uint256 takerFee = _calcFee(quoteVolume, newTakerFeeBps);
        uint256 creatorFee = _calcFee(takerFee, newCreatorShareBps);
        uint256 makerRebate = _calcFee(takerFee, newMakerRebateShareBps);
        uint256 systemFee = takerFee - creatorFee - makerRebate;

        assertEq(QUOTE.balanceOf(USER1) - makerQuoteBefore, makerRebate, "New maker rebate should be applied");
        assertEq(QUOTE.balanceOf(CREATOR) - creatorBefore, creatorFee, "New creator fee should be applied");
        assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, systemFee, "New system fee should be applied");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Zero Fee Edge Case
    // ─────────────────────────────────────────────────────────────────────────────

    function test_zero_taker_fee() external {
        // Set taker fee to 0
        bytes memory zeroFeeData =
            abi.encode(FEE_COLLECTOR, CREATOR, uint32(0), CREATOR_SHARE_BPS, MAKER_REBATE_SHARE_BPS);
        address[] memory pairs = new address[](1);
        pairs[0] = address(PAIR);
        vm.prank(OWNER);
        MARKET.setFeeController(pairs, address(FEE_CONTROLLER), zeroFeeData);

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 sellerQuoteBefore = QUOTE.balanceOf(USER2);
        uint256 creatorBefore = QUOTE.balanceOf(CREATOR);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Seller should receive full volume
        assertEq(QUOTE.balanceOf(USER2) - sellerQuoteBefore, quoteVolume, "Seller should receive full volume");
        // No fees collected
        assertEq(QUOTE.balanceOf(CREATOR), creatorBefore, "No creator fee with 0 taker fee");
        assertEq(QUOTE.balanceOf(FEE_COLLECTOR), feeCollectorBefore, "No system fee with 0 taker fee");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // View Functions Tests
    // ─────────────────────────────────────────────────────────────────────────────

    function test_view_sellerMakerFeeBps() external view {
        assertEq(FEE_CONTROLLER.sellerMakerFeeBps(), 0, "Seller maker fee should always be 0");
    }

    function test_view_buyerMakerFeeBps() external view {
        assertEq(FEE_CONTROLLER.buyerMakerFeeBps(), 0, "Buyer maker fee should always be 0");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Configuration View Functions Tests (reading from PAIR's storage via delegatecall context)
    // ─────────────────────────────────────────────────────────────────────────────

    /// @dev ERC-7201 storage slot for FeeControllerV3Split
    /// keccak256(abi.encode(uint256(keccak256("cross.storage.FeeControllerV3Split")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant FeeControllerV3SplitStorageLocation =
        0xc063cf2dfd170dd7c3dd72990f0424c6080fb35c03fb56e3a8e6c5b294381000;

    function test_view_feeCollector_afterInit() external view {
        // Read slot 0: feeCollector address (right-aligned in 32 bytes)
        bytes32 slot0 = vm.load(address(PAIR), FeeControllerV3SplitStorageLocation);
        address storedFeeCollector = address(uint160(uint256(slot0)));

        assertEq(storedFeeCollector, FEE_COLLECTOR, "FeeCollector should match initialized value");
    }

    function test_view_creator_afterInit() external view {
        // Read slot 1: creator (20 bytes) + takerFeeBps (4) + creatorShareBps (4) + makerRebateShareBps (4) = 32 bytes
        // Layout (right to left): makerRebateShareBps | creatorShareBps | takerFeeBps | creator
        bytes32 slot1 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 1));

        // Extract creator (lowest 20 bytes = 160 bits)
        address storedCreator = address(uint160(uint256(slot1)));

        assertEq(storedCreator, CREATOR, "Creator should match initialized value");
    }

    function test_view_takerFeeBps_afterInit() external view {
        bytes32 slot1 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 1));

        // Extract takerFeeBps (bits 160-191, shift right 160 bits, mask 32 bits)
        uint32 storedTakerFeeBps = uint32(uint256(slot1) >> 160);

        assertEq(storedTakerFeeBps, TAKER_FEE_BPS, "TakerFeeBps should match initialized value");
    }

    function test_view_creatorShareBps_afterInit() external view {
        bytes32 slot1 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 1));

        // Extract creatorShareBps (bits 192-223, shift right 192 bits, mask 32 bits)
        uint32 storedCreatorShareBps = uint32(uint256(slot1) >> 192);

        assertEq(storedCreatorShareBps, CREATOR_SHARE_BPS, "CreatorShareBps should match initialized value");
    }

    function test_view_makerRebateShareBps_afterInit() external view {
        bytes32 slot1 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 1));

        // Extract makerRebateShareBps (bits 224-255, shift right 224 bits)
        uint32 storedMakerRebateShareBps = uint32(uint256(slot1) >> 224);

        assertEq(
            storedMakerRebateShareBps, MAKER_REBATE_SHARE_BPS, "MakerRebateShareBps should match initialized value"
        );
    }

    function test_view_quote_afterInit() external view {
        // Read slot 2: quote address
        bytes32 slot2 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 2));
        address storedQuote = address(uint160(uint256(slot2)));

        assertEq(storedQuote, address(QUOTE), "Quote should match initialized value");
    }

    function test_view_denominator_afterInit() external view {
        // Read slot 3: denominator (full uint256)
        bytes32 slot3 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 3));
        uint256 storedDenominator = uint256(slot3);

        assertEq(storedDenominator, BASE_DECIMALS, "Denominator should match initialized value");
    }

    function test_view_allConfigValues_afterInit() external view {
        // Comprehensive test: verify all config values in one test
        bytes32 slot0 = vm.load(address(PAIR), FeeControllerV3SplitStorageLocation);
        bytes32 slot1 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 1));
        bytes32 slot2 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 2));
        bytes32 slot3 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 3));

        // Decode all values
        address storedFeeCollector = address(uint160(uint256(slot0)));
        address storedCreator = address(uint160(uint256(slot1)));
        uint32 storedTakerFeeBps = uint32(uint256(slot1) >> 160);
        uint32 storedCreatorShareBps = uint32(uint256(slot1) >> 192);
        uint32 storedMakerRebateShareBps = uint32(uint256(slot1) >> 224);
        address storedQuote = address(uint160(uint256(slot2)));
        uint256 storedDenominator = uint256(slot3);

        // Verify all values
        assertEq(storedFeeCollector, FEE_COLLECTOR, "FeeCollector mismatch");
        assertEq(storedCreator, CREATOR, "Creator mismatch");
        assertEq(storedTakerFeeBps, TAKER_FEE_BPS, "TakerFeeBps mismatch");
        assertEq(storedCreatorShareBps, CREATOR_SHARE_BPS, "CreatorShareBps mismatch");
        assertEq(storedMakerRebateShareBps, MAKER_REBATE_SHARE_BPS, "MakerRebateShareBps mismatch");
        assertEq(storedQuote, address(QUOTE), "Quote mismatch");
        assertEq(storedDenominator, BASE_DECIMALS, "Denominator mismatch");
    }

    function test_view_configValues_afterUpdate() external {
        // Test that config values are properly updated after setFeeController
        uint32 newTakerFeeBps = 50; // 0.5%
        uint32 newCreatorShareBps = 5000; // 50%
        uint32 newMakerRebateShareBps = 1000; // 10%
        address newCreator = makeAddr("newCreator");
        address newFeeCollector = makeAddr("newFeeCollector");

        // Update configuration
        bytes memory newFeeData =
            abi.encode(newFeeCollector, newCreator, newTakerFeeBps, newCreatorShareBps, newMakerRebateShareBps);
        address[] memory pairs = new address[](1);
        pairs[0] = address(PAIR);
        vm.prank(OWNER);
        MARKET.setFeeController(pairs, address(FEE_CONTROLLER), newFeeData);

        // Read and verify updated values
        bytes32 slot0 = vm.load(address(PAIR), FeeControllerV3SplitStorageLocation);
        bytes32 slot1 = vm.load(address(PAIR), bytes32(uint256(FeeControllerV3SplitStorageLocation) + 1));

        address storedFeeCollector = address(uint160(uint256(slot0)));
        address storedCreator = address(uint160(uint256(slot1)));
        uint32 storedTakerFeeBps = uint32(uint256(slot1) >> 160);
        uint32 storedCreatorShareBps = uint32(uint256(slot1) >> 192);
        uint32 storedMakerRebateShareBps = uint32(uint256(slot1) >> 224);

        assertEq(storedFeeCollector, newFeeCollector, "Updated FeeCollector mismatch");
        assertEq(storedCreator, newCreator, "Updated Creator mismatch");
        assertEq(storedTakerFeeBps, newTakerFeeBps, "Updated TakerFeeBps mismatch");
        assertEq(storedCreatorShareBps, newCreatorShareBps, "Updated CreatorShareBps mismatch");
        assertEq(storedMakerRebateShareBps, newMakerRebateShareBps, "Updated MakerRebateShareBps mismatch");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Initialize Validation Tests
    // ─────────────────────────────────────────────────────────────────────────────

    function test_initialize_revert_zeroFeeCollector() external {
        // Create a new pair with invalid config
        T20 newBase = new T20("NEW", "NEW", 18);
        bytes memory invalidFeeData = abi.encode(
            address(0), // feeCollector = 0 (invalid)
            CREATOR,
            TAKER_FEE_BPS,
            CREATOR_SHARE_BPS,
            MAKER_REBATE_SHARE_BPS
        );

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(IFeeController.FeeControllerInvalidFeeCollector.selector));
        MARKET.createPair(address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, invalidFeeData);
    }

    function test_initialize_revert_zeroCreator() external {
        T20 newBase = new T20("NEW", "NEW", 18);
        bytes memory invalidFeeData = abi.encode(
            FEE_COLLECTOR,
            address(0), // creator = 0 (invalid)
            TAKER_FEE_BPS,
            CREATOR_SHARE_BPS,
            MAKER_REBATE_SHARE_BPS
        );

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(IFeeController.FeeControllerInvalidFeeCollector.selector));
        MARKET.createPair(address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, invalidFeeData);
    }

    function test_initialize_revert_takerFeeTooHigh() external {
        T20 newBase = new T20("NEW", "NEW", 18);
        bytes memory invalidFeeData = abi.encode(
            FEE_COLLECTOR,
            CREATOR,
            uint32(10000), // takerFeeBps >= 10000 (invalid)
            CREATOR_SHARE_BPS,
            MAKER_REBATE_SHARE_BPS
        );

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(IFeeController.FeeControllerInvalidFeeBps.selector));
        MARKET.createPair(address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, invalidFeeData);
    }

    function test_initialize_revert_shareSumTooHigh() external {
        T20 newBase = new T20("NEW", "NEW", 18);
        bytes memory invalidFeeData = abi.encode(
            FEE_COLLECTOR,
            CREATOR,
            TAKER_FEE_BPS,
            uint32(6000), // creatorShareBps
            uint32(5000) // makerRebateShareBps, sum = 11000 > 10000 (invalid)
        );

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(IFeeController.FeeControllerInvalidFeeBps.selector));
        MARKET.createPair(address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, invalidFeeData);
    }

    function test_initialize_revert_directCall_notDelegateCall() external {
        // Calling initialize directly (not via delegatecall) should revert
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, CREATOR, TAKER_FEE_BPS, CREATOR_SHARE_BPS, MAKER_REBATE_SHARE_BPS);

        vm.expectRevert(abi.encodeWithSelector(IFeeController.FeeControllerNotDelegateCall.selector));
        FEE_CONTROLLER.initialize(address(QUOTE), BASE_DECIMALS, feeData);
    }

    function test_calcBuyVolumeWithFee_maker() external view {
        uint256 volume = _toQuote(1000);
        // Maker has no fee in V3Split - this is a pure calculation, no storage needed
        uint256 result = FEE_CONTROLLER.calcBuyVolumeWithFee(true, volume);
        assertEq(result, volume, "Maker should have no fee added");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Zero Rebate Share Tests
    // ─────────────────────────────────────────────────────────────────────────────

    function test_zero_maker_rebate_share() external {
        // Set maker rebate to 0
        bytes memory zeroRebateData = abi.encode(
            FEE_COLLECTOR,
            CREATOR,
            TAKER_FEE_BPS,
            CREATOR_SHARE_BPS,
            uint32(0) // makerRebateShareBps = 0
        );
        address[] memory pairs = new address[](1);
        pairs[0] = address(PAIR);
        vm.prank(OWNER);
        MARKET.setFeeController(pairs, address(FEE_CONTROLLER), zeroRebateData);

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 makerQuoteBefore = QUOTE.balanceOf(USER1);
        uint256 creatorBefore = QUOTE.balanceOf(CREATOR);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Calculate expected with 0 rebate
        uint256 takerFee = _calcFee(quoteVolume, TAKER_FEE_BPS);
        uint256 creatorFee = _calcFee(takerFee, CREATOR_SHARE_BPS);
        uint256 systemFee = takerFee - creatorFee; // No rebate deduction

        // Maker should NOT receive rebate
        assertEq(QUOTE.balanceOf(USER1) - makerQuoteBefore, 0, "Maker should not receive rebate");
        assertEq(QUOTE.balanceOf(CREATOR) - creatorBefore, creatorFee, "Creator gets their share");
        assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, systemFee, "FeeCollector gets the rest");
    }

    function test_zero_creator_share() external {
        // Set creator share to 0
        bytes memory zeroCreatorData = abi.encode(
            FEE_COLLECTOR,
            CREATOR,
            TAKER_FEE_BPS,
            uint32(0), // creatorShareBps = 0
            MAKER_REBATE_SHARE_BPS
        );
        address[] memory pairs = new address[](1);
        pairs[0] = address(PAIR);
        vm.prank(OWNER);
        MARKET.setFeeController(pairs, address(FEE_CONTROLLER), zeroCreatorData);

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 creatorBefore = QUOTE.balanceOf(CREATOR);
        uint256 feeCollectorBefore = QUOTE.balanceOf(FEE_COLLECTOR);

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        // Calculate expected with 0 creator share
        uint256 takerFee = _calcFee(quoteVolume, TAKER_FEE_BPS);
        uint256 makerRebate = _calcFee(takerFee, MAKER_REBATE_SHARE_BPS);
        uint256 systemFee = takerFee - makerRebate; // No creator deduction

        // Creator should NOT receive anything
        assertEq(QUOTE.balanceOf(CREATOR) - creatorBefore, 0, "Creator should not receive fee");
        assertEq(
            QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorBefore, systemFee, "FeeCollector gets creator's share too"
        );
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Comprehensive Scenario
    // ─────────────────────────────────────────────────────────────────────────────

    function test_comprehensive_v3split_scenario() external {
        // Track fees
        uint256 totalTakerFeesCollected;
        uint256 totalCreatorFees;
        uint256 totalMakerRebates;
        uint256 totalSystemFees;

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        uint256 creatorInit = QUOTE.balanceOf(CREATOR);
        uint256 feeCollectorInit = QUOTE.balanceOf(FEE_COLLECTOR);

        // Trade 1: USER1 buys, USER2 sells (USER2 is taker)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        {
            uint256 takerFee = _calcFee(quoteVolume, TAKER_FEE_BPS);
            uint256 creatorFee = _calcFee(takerFee, CREATOR_SHARE_BPS);
            uint256 makerRebate = _calcFee(takerFee, MAKER_REBATE_SHARE_BPS);
            uint256 systemFee = takerFee - creatorFee - makerRebate;
            totalTakerFeesCollected += takerFee;
            totalCreatorFees += creatorFee;
            totalMakerRebates += makerRebate;
            totalSystemFees += systemFee;
        }

        // Trade 2: USER2 sells, USER1 buys (USER1 is taker)
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 takerFee2 = _calcFee(quoteVolume, TAKER_FEE_BPS);
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteVolume + takerFee2, 0);

        {
            uint256 creatorFee = _calcFee(takerFee2, CREATOR_SHARE_BPS);
            uint256 makerRebate = _calcFee(takerFee2, MAKER_REBATE_SHARE_BPS);
            uint256 systemFee = takerFee2 - creatorFee - makerRebate;
            totalTakerFeesCollected += takerFee2;
            totalCreatorFees += creatorFee;
            totalMakerRebates += makerRebate;
            totalSystemFees += systemFee;
        }

        // Verify totals
        assertEq(QUOTE.balanceOf(CREATOR) - creatorInit, totalCreatorFees, "Total creator fees");
        assertEq(QUOTE.balanceOf(FEE_COLLECTOR) - feeCollectorInit, totalSystemFees, "Total system fees");

        console.log("=== V3Split Comprehensive Scenario Complete ===");
        console.log("Total taker fees:", totalTakerFeesCollected);
        console.log("Creator fees:", totalCreatorFees);
        console.log("Maker rebates:", totalMakerRebates);
        console.log("System fees:", totalSystemFees);
    }
}
