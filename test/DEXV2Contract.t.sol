// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {CrossDexImplV2} from "../src/CrossDexImplV2.sol";
import {CrossDexRouterV2} from "../src/CrossDexRouterV2.sol";
import {MarketImplV2} from "../src/MarketImplV2.sol";
import {PairImplV2} from "../src/PairImplV2.sol";
import {WETH} from "../src/WETH.sol";

import {BPS_DENOMINATOR, IMarketV2, NO_FEE_BPS} from "../src/interfaces/IMarket.sol";
import {IPair, IPairV2} from "../src/interfaces/IPair.sol";

import {T20} from "./mock/T20.sol";

contract DEXV2ContractTest is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    address public constant USER1 = address(bytes20("USER1"));
    address public constant USER2 = address(bytes20("USER2"));

    CrossDexImplV2 public CROSS_DEX;
    CrossDexRouterV2 public ROUTER;
    WETH public CROSS;

    IERC20 public QUOTE;
    IERC20 public BASE;

    MarketImplV2 public MARKET;
    PairImplV2 public PAIR;

    uint256 public FIND_PREV_PRICE_COUNT = type(uint256).max;
    uint256 public MAX_MATCH_COUNT = type(uint256).max;
    uint256 public CANCEL_LIMIT = type(uint256).max;

    // Market-level fees (in BPS) - seller만 수수료
    uint256 public MARKET_SELLER_MAKER_FEE = 20; // 0.2%
    uint256 public MARKET_SELLER_TAKER_FEE = 30; // 0.3%

    uint256 public QUOTE_DECIMALS;
    uint256 public BASE_DECIMALS;

    uint256[2] internal _searchPrices;

    function setUp() external {
        _deployV2(18, 18, 1e2, 1e6);
    }

    function _deployV2(uint8 quote_decimals, uint8 base_decimals, uint256 quote_tick_size, uint256 base_tick_size)
        internal
    {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");
        vm.label(USER1, "user1");
        vm.label(USER2, "user2");

        vm.startPrank(OWNER);

        QUOTE_DECIMALS = 10 ** quote_decimals;
        BASE_DECIMALS = 10 ** base_decimals;

        {
            // deploy impl contracts (using V2 versions)
            address routerImpl = address(new CrossDexRouterV2());
            address marketImpl = address(new MarketImplV2());
            address pairImpl = address(new PairImplV2());

            // deploy cross dex
            address crossDexImpl = address(new CrossDexImplV2());
            ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
            CROSS_DEX = CrossDexImplV2(address(proxy));
            CROSS_DEX.initialize(
                OWNER, routerImpl, FIND_PREV_PRICE_COUNT, MAX_MATCH_COUNT, CANCEL_LIMIT, marketImpl, pairImpl
            );
        }
        {
            // get contracts from CROSS_DEX
            ROUTER = CrossDexRouterV2(CROSS_DEX.ROUTER());
            CROSS = WETH(payable(address(ROUTER.CROSS())));
        }
        {
            // deploy base and quote tokens
            QUOTE = new T20("QUOTE", "QUOTE", quote_decimals);
            BASE = new T20("BASE", "BASE", base_decimals);
        }
        {
            // create market with V2 fees (seller만 수수료)
            bytes memory marketInitializeData = abi.encode(
                uint32(MARKET_SELLER_MAKER_FEE),
                uint32(MARKET_SELLER_TAKER_FEE),
                uint32(0), // buyer maker fee = 0
                uint32(0) // buyer taker fee = 0
            );
            address market = CROSS_DEX.createMarket(OWNER, address(QUOTE), FEE_COLLECTOR, marketInitializeData);
            MARKET = MarketImplV2(market);

            // create pair with default fees (NO_FEE_BPS = use market fees)
            address pair = MARKET.createPair(
                address(BASE),
                QUOTE_DECIMALS / quote_tick_size,
                BASE_DECIMALS / base_tick_size,
                NO_FEE_BPS, // use market seller maker fee
                NO_FEE_BPS, // use market seller taker fee
                NO_FEE_BPS, // use market buyer maker fee (0)
                NO_FEE_BPS // use market buyer taker fee (0)
            );
            PAIR = PairImplV2(pair);
        }

        // Setup initial balances - buyer는 충분한 자금이 필요 (거래량 + 수수료)
        QUOTE.transfer(USER1, _toQuote(50000)); // buyer에게 충분한 QUOTE
        QUOTE.transfer(USER2, _toQuote(50000));
        BASE.transfer(USER1, _toBase(50000)); // seller에게 충분한 BASE
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

    function _toBase(uint256 x) internal view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) internal view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }

    // Test V2 system initialization with seller-only fees
    function test_v2_initialization_seller_only_fees() external view {
        IMarketV2.FeeConfig memory feeConfig = MARKET.getFeeConfig();
        assertEq(feeConfig.sellerMakerFeeBps, MARKET_SELLER_MAKER_FEE, "Seller maker fee should match");
        assertEq(feeConfig.sellerTakerFeeBps, MARKET_SELLER_TAKER_FEE, "Seller taker fee should match");
        assertEq(feeConfig.buyerMakerFeeBps, 0, "Buyer maker fee should be 0");
        assertEq(feeConfig.buyerTakerFeeBps, 0, "Buyer taker fee should be 0");
    }

    // Test setting new market fees with buyer fees
    function test_v2_set_market_fees_with_buyer_fees() external {
        uint32 newSellerMakerFee = 15; // 0.15%
        uint32 newSellerTakerFee = 25; // 0.25%
        uint32 newBuyerMakerFee = 10; // 0.10%
        uint32 newBuyerTakerFee = 20; // 0.20%

        vm.prank(OWNER);
        MARKET.setMarketFees(newSellerMakerFee, newSellerTakerFee, newBuyerMakerFee, newBuyerTakerFee);

        IMarketV2.FeeConfig memory updatedFeeConfig = MARKET.getFeeConfig();
        assertEq(updatedFeeConfig.sellerMakerFeeBps, newSellerMakerFee, "Seller maker fee should be updated");
        assertEq(updatedFeeConfig.sellerTakerFeeBps, newSellerTakerFee, "Seller taker fee should be updated");
        assertEq(updatedFeeConfig.buyerMakerFeeBps, newBuyerMakerFee, "Buyer maker fee should be updated");
        assertEq(updatedFeeConfig.buyerTakerFeeBps, newBuyerTakerFee, "Buyer taker fee should be updated");
    }

    // Test pair-specific fees override market fees
    function test_v2_pair_specific_fees() external {
        uint32 pairSellerMakerFee = 50; // 0.5%
        uint32 pairSellerTakerFee = 60; // 0.6%
        uint32 pairBuyerMakerFee = 30; // 0.3%
        uint32 pairBuyerTakerFee = 40; // 0.4%

        vm.prank(OWNER);
        PAIR.setPairFees(pairSellerMakerFee, pairSellerTakerFee, pairBuyerMakerFee, pairBuyerTakerFee);

        (uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee) = PAIR.feeConfig();
        assertEq(sellerMakerFee, pairSellerMakerFee, "Pair seller maker fee should be set");
        assertEq(sellerTakerFee, pairSellerTakerFee, "Pair seller taker fee should be set");
        assertEq(buyerMakerFee, pairBuyerMakerFee, "Pair buyer maker fee should be set");
        assertEq(buyerTakerFee, pairBuyerTakerFee, "Pair buyer taker fee should be set");

        // Check effective fees use pair fees instead of market fees
        (
            uint32 effectiveSellerMakerFee,
            uint32 effectiveSellerTakerFee,
            uint32 effectiveBuyerMakerFee,
            uint32 effectiveBuyerTakerFee
        ) = PAIR.getEffectiveFees();
        assertEq(effectiveSellerMakerFee, pairSellerMakerFee, "Effective seller maker fee should use pair fee");
        assertEq(effectiveSellerTakerFee, pairSellerTakerFee, "Effective seller taker fee should use pair fee");
        assertEq(effectiveBuyerMakerFee, pairBuyerMakerFee, "Effective buyer maker fee should use pair fee");
        assertEq(effectiveBuyerTakerFee, pairBuyerTakerFee, "Effective buyer taker fee should use pair fee");
    }

    // Test simple limit order without fees first
    function test_v2_simple_limit_order_no_fees() external {
        uint256 tradePrice = _toQuote(1);
        uint256 tradeAmount = _toBase(100);
        uint256 tradeVolume = Math.mulDiv(tradePrice, tradeAmount, _toQuote(1)); // 100 QUOTE

        uint256 initialBuyerBalance = QUOTE.balanceOf(USER1);

        // Buyer places limit order without any fees
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), tradePrice, tradeAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 afterOrderBuyerBalance = QUOTE.balanceOf(USER1);

        // Should only deduct the trade volume (no fees)
        assertEq(
            initialBuyerBalance - afterOrderBuyerBalance, tradeVolume, "Buyer should only pay trade volume when no fees"
        );

        // Execute the trade by selling to this buy order
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), tradeAmount, type(uint256).max);
    }

    // Test buyer pays fees in addition to trade amount - BUY LIMIT ORDER
    function test_v2_buyer_pays_additional_fees_buy_limit() external {
        // Set buyer maker fee and buyer taker fee (taker >= maker)
        uint32 buyerMakerFee = 100; // 1%
        uint32 buyerTakerFee = 150; // 1.5% (must be >= maker fee)

        vm.prank(OWNER);
        PAIR.setPairFees(NO_FEE_BPS, NO_FEE_BPS, buyerMakerFee, buyerTakerFee);

        uint256 tradeVolume = _toQuote(100);
        uint256 initialBuyerBalance = QUOTE.balanceOf(USER1);

        // Buyer places limit order - RouterV2 calculates required amount using taker fee (higher fee)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 afterOrderBuyerBalance = QUOTE.balanceOf(USER1);

        // Expected fee calculation based on taker fee: 100 QUOTE * 1.5% = 1.5 QUOTE
        uint256 expectedTakerFee = Math.mulDiv(tradeVolume, buyerTakerFee, BPS_DENOMINATOR);
        uint256 expectedTotalDeduction = tradeVolume + expectedTakerFee; // 100 + 1.5 = 101.5 QUOTE

        assertEq(
            initialBuyerBalance - afterOrderBuyerBalance,
            expectedTotalDeduction,
            "Buyer should pay trade volume + taker fee (worst case)"
        );

        uint256 initialFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // Execute the trade by selling to this buy order
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(100), type(uint256).max);

        uint256 finalFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 collectedFees = finalFeeCollectorBalance - initialFeeCollectorBalance;

        // Since the limit order gets filled, it becomes a maker, so buyer maker fee is charged
        // But seller acts as taker and pays seller taker fee (from market default)
        uint256 expectedBuyerMakerFee = Math.mulDiv(tradeVolume, buyerMakerFee, BPS_DENOMINATOR);
        uint256 expectedSellerTakerFee = Math.mulDiv(tradeVolume, 30, BPS_DENOMINATOR); // market default seller taker fee
        uint256 expectedTotalFees = expectedBuyerMakerFee + expectedSellerTakerFee;

        assertEq(
            collectedFees, expectedTotalFees, "Fee collector should receive both buyer maker and seller taker fees"
        );
    }
    // Test buyer pays fees in addition to trade amount - BUY MARKET ORDER

    function test_v2_buyer_pays_additional_fees_buy_market() external {
        // Set seller maker fee and buyer taker fee
        uint32 sellerMakerFee = 30; // 0.3%
        uint32 sellerTakerFee = 50; // 0.5%
        uint32 buyerTakerFee = 75; // 0.75%

        vm.prank(OWNER);
        PAIR.setPairFees(sellerMakerFee, sellerTakerFee, NO_FEE_BPS, buyerTakerFee);

        uint256 tradePrice = _toQuote(1);
        uint256 tradeAmount = _toBase(100);

        // First, create a sell limit order for the buyer to match against
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), tradePrice, tradeAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 initialBuyerBalance = QUOTE.balanceOf(USER1);
        uint256 initialFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        uint256 buyVolume = _toQuote(100); // Buyer wants to spend 100 QUOTE

        // Buyer places market order - RouterV2 calculates additional fee requirement
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), buyVolume, type(uint256).max);

        uint256 afterOrderBuyerBalance = QUOTE.balanceOf(USER1);
        uint256 finalFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // Expected fees: buyer taker fee + seller maker fee
        uint256 expectedBuyerFee = Math.mulDiv(buyVolume, buyerTakerFee, BPS_DENOMINATOR);
        uint256 expectedSellerFee = Math.mulDiv(buyVolume, sellerMakerFee, BPS_DENOMINATOR);
        uint256 expectedTotalFees = expectedBuyerFee + expectedSellerFee;

        // Buyer should pay: original volume + buyer taker fee
        uint256 expectedBuyerPayment = buyVolume + expectedBuyerFee;

        assertEq(
            initialBuyerBalance - afterOrderBuyerBalance,
            expectedBuyerPayment,
            "Buyer should pay trade volume + taker fee"
        );

        uint256 collectedFees = finalFeeCollectorBalance - initialFeeCollectorBalance;
        assertEq(collectedFees, expectedTotalFees, "Fee collector should receive both buyer and seller fees");
    }

    // Test limit-limit order matching with both buyer and seller fees
    function test_v2_limit_limit_matching_with_both_fees() external {
        uint32 sellerMakerFee = 30; // 0.3%
        uint32 buyerMakerFee = 25; // 0.25%
        uint32 buyerTakerFee = 35; // 0.35% (must be >= maker fee)

        vm.prank(OWNER);
        PAIR.setPairFees(sellerMakerFee, NO_FEE_BPS, buyerMakerFee, buyerTakerFee);

        uint256 tradePrice = _toQuote(1);
        uint256 tradeAmount = _toBase(100);
        uint256 tradeVolume = _toQuote(100);

        uint256 initialFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // USER1 places buy limit order first (buyer maker)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), tradePrice, tradeAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER2 places sell limit order that matches (seller taker - but since it's limit order, becomes maker)
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), tradePrice, tradeAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 finalFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 collectedFees = finalFeeCollectorBalance - initialFeeCollectorBalance;

        // Expected fees: both orders are makers, so both pay maker fees
        uint256 expectedBuyerFee = Math.mulDiv(tradeVolume, buyerMakerFee, BPS_DENOMINATOR);
        uint256 expectedSellerFee = Math.mulDiv(tradeVolume, sellerMakerFee, BPS_DENOMINATOR);
        uint256 expectedTotalFees = expectedBuyerFee + expectedSellerFee;

        assertEq(collectedFees, expectedTotalFees, "Should collect both buyer and seller maker fees");
    }

    // Test market order fee calculation when matching multiple limit orders
    function test_v2_market_order_multiple_matches() external {
        uint32 sellerMakerFee = 20; // 0.2%
        uint32 buyerTakerFee = 30; // 0.3%

        vm.prank(OWNER);
        PAIR.setPairFees(sellerMakerFee, NO_FEE_BPS, NO_FEE_BPS, buyerTakerFee);

        // Create multiple sell limit orders at different prices
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(50), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(2), _toBase(50), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 initialFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 initialBuyerBalance = QUOTE.balanceOf(USER1);

        // Buyer places market order that matches both
        uint256 buyVolume = _toQuote(150); // 50 + 100 = 150 QUOTE total
        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), buyVolume, type(uint256).max);

        uint256 finalFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 finalBuyerBalance = QUOTE.balanceOf(USER1);

        // Calculate expected fees
        uint256 expectedBuyerFee = Math.mulDiv(buyVolume, buyerTakerFee, BPS_DENOMINATOR);
        uint256 expectedSellerFee = Math.mulDiv(buyVolume, sellerMakerFee, BPS_DENOMINATOR);
        uint256 expectedTotalFees = expectedBuyerFee + expectedSellerFee;

        uint256 collectedFees = finalFeeCollectorBalance - initialFeeCollectorBalance;
        assertEq(collectedFees, expectedTotalFees, "Should collect fees from multiple order matches");

        // Buyer should pay trade volume + buyer taker fee
        uint256 buyerPayment = initialBuyerBalance - finalBuyerBalance;
        uint256 expectedBuyerPayment = buyVolume + expectedBuyerFee;
        assertEq(buyerPayment, expectedBuyerPayment, "Buyer should pay volume + taker fee");
    }

    // Test fee validation (should reject fees >= BPS_DENOMINATOR)
    function test_v2_fee_validation() external {
        // Test market fee validation
        vm.prank(OWNER);
        vm.expectRevert();
        MARKET.setMarketFees(BPS_DENOMINATOR, 0, 0, 0); // 100% fee should be rejected

        vm.prank(OWNER);
        vm.expectRevert();
        MARKET.setMarketFees(0, BPS_DENOMINATOR + 1, 0, 0); // >100% fee should be rejected

        // Test pair fee validation
        vm.prank(OWNER);
        vm.expectRevert();
        PAIR.setPairFees(0, 0, BPS_DENOMINATOR, 0); // 100% fee should be rejected

        // Test valid maximum fee (BPS_DENOMINATOR - 1)
        vm.prank(OWNER);
        MARKET.setMarketFees(BPS_DENOMINATOR - 1, BPS_DENOMINATOR - 1, BPS_DENOMINATOR - 1, BPS_DENOMINATOR - 1);

        IMarketV2.FeeConfig memory feeConfig = MARKET.getFeeConfig();
        assertEq(feeConfig.sellerMakerFeeBps, BPS_DENOMINATOR - 1, "Should accept maximum valid fee");
    }

    // Test fee structure validation (maker fee <= taker fee)
    function test_v2_fee_structure_validation() external {
        // Test that setPairFees reverts when maker fee > taker fee
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeStructure(uint32,uint32)", 100, 50));
        PAIR.setPairFees(100, 50, 0, 0); // seller maker > seller taker

        // Test buyer fee validation
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeStructure(uint32,uint32)", 75, 50));
        PAIR.setPairFees(0, 0, 75, 50); // buyer maker > buyer taker

        // Test that valid configuration works
        vm.prank(OWNER);
        PAIR.setPairFees(50, 100, 25, 75); // maker <= taker for both

        // This should work fine
        PAIR.getEffectiveFees();
    }

    // Test access control for fee setting functions
    function test_v2_access_control() external {
        vm.prank(USER1);
        vm.expectRevert();
        PAIR.setPairFees(50, 60, 70, 80);

        vm.prank(USER1);
        vm.expectRevert();
        MARKET.setMarketFees(50, 60, 70, 80);

        vm.prank(USER1);
        vm.expectRevert();
        MARKET.setFeeCollector(USER1);
    }

    // Test fee collector changes
    function test_v2_fee_collector_management() external {
        address newFeeCollector = address(bytes20("NEW_FEE_COLLECTOR"));

        vm.prank(OWNER);
        vm.expectEmit(true, true, false, false);
        emit FeeCollectorChanged(FEE_COLLECTOR, newFeeCollector);
        MARKET.setFeeCollector(newFeeCollector);

        assertEq(MARKET.feeCollector(), newFeeCollector, "Fee collector should be updated");

        // Test zero address rejection
        vm.prank(OWNER);
        vm.expectRevert();
        MARKET.setFeeCollector(address(0));
    }

    // Test zero fees work correctly
    function test_v2_zero_fees() external {
        vm.prank(OWNER);
        PAIR.setPairFees(0, 0, 0, 0);

        uint256 initialFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 initialBuyerBalance = QUOTE.balanceOf(USER1);

        uint256 tradeVolume = _toQuote(100);

        // Execute a trade
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(100), type(uint256).max);

        uint256 finalFeeCollectorBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 finalBuyerBalance = QUOTE.balanceOf(USER1);

        assertEq(finalFeeCollectorBalance, initialFeeCollectorBalance, "No fees should be collected");
        assertEq(initialBuyerBalance - finalBuyerBalance, tradeVolume, "Buyer should only pay trade volume");
    }

    // Test events are emitted correctly
    function test_v2_events() external {
        // Test market fee update event
        vm.prank(OWNER);
        vm.expectEmit(true, true, true, true);
        emit MarketFeesUpdated(10, 20, 30, 40);
        MARKET.setMarketFees(10, 20, 30, 40);

        // Test pair fee update event
        vm.prank(OWNER);
        vm.expectEmit(true, true, true, true);
        emit PairFeesUpdated(15, 25, 35, 45);
        PAIR.setPairFees(15, 25, 35, 45);
    }

    // Test RouterV2 calculates correct amount for buy operations
    function test_v2_router_calculates_correct_buy_amounts() external {
        uint32 buyerMakerFee = 200; // 2%
        uint32 buyerTakerFee = 250; // 2.5% (RouterV2 uses this for calculation)

        vm.prank(OWNER);
        PAIR.setPairFees(NO_FEE_BPS, NO_FEE_BPS, buyerMakerFee, buyerTakerFee);

        uint256 tradeVolume = _toQuote(1000);
        uint256 expectedFee = Math.mulDiv(tradeVolume, buyerTakerFee, BPS_DENOMINATOR); // 25 QUOTE (taker fee)
        uint256 expectedTotal = tradeVolume + expectedFee; // 1025 QUOTE

        uint256 initialBalance = QUOTE.balanceOf(USER1);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(1), _toBase(1000), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 finalBalance = QUOTE.balanceOf(USER1);
        uint256 actualDeduction = initialBalance - finalBalance;

        assertEq(actualDeduction, expectedTotal, "RouterV2 should deduct trade volume + buyer taker fee (worst case)");
    }

    // Events (to make the test compile)
    event MarketFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee);
    event PairFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee);
    event FeeCollectorChanged(address indexed before, address indexed current);
}
