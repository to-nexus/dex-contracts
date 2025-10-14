// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {CrossDexImplV2} from "../src/CrossDexImplV2.sol";
import {CrossDexRouter} from "../src/CrossDexRouter.sol";
import {MarketImplV2} from "../src/MarketImplV2.sol";
import {PairImplV2} from "../src/PairImplV2.sol";
import {WETH} from "../src/WETH.sol";

import {IMarketV2, NO_FEE_BPS} from "../src/interfaces/IMarket.sol";
import {IPair} from "../src/interfaces/IPair.sol";

import {T20} from "./mock/T20.sol";

contract DEXFeeV2Test is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    address public constant USER1 = address(bytes20("USER1"));
    address public constant USER2 = address(bytes20("USER2"));

    CrossDexImplV2 public CROSS_DEX;
    CrossDexRouter public ROUTER;
    WETH public CROSS;

    IERC20 public QUOTE;
    IERC20 public BASE;

    MarketImplV2 public MARKET;
    PairImplV2 public PAIR;

    uint256 public FIND_PREV_PRICE_COUNT = type(uint256).max;
    uint256 public MAX_MATCH_COUNT = type(uint256).max;
    uint256 public CANCEL_LIMIT = type(uint256).max;

    // Market-level fees (in BPS) - 4가지 수수료
    uint256 public MARKET_SELLER_MAKER_FEE = 20; // 0.2%
    uint256 public MARKET_SELLER_TAKER_FEE = 30; // 0.3%
    uint256 public MARKET_BUYER_MAKER_FEE = 15; // 0.15%
    uint256 public MARKET_BUYER_TAKER_FEE = 25; // 0.25%

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
            address routerImpl = address(new CrossDexRouter());
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
            ROUTER = CrossDexRouter(CROSS_DEX.ROUTER());
            CROSS = WETH(payable(address(ROUTER.CROSS())));
        }
        {
            // deploy base and quote tokens
            QUOTE = new T20("QUOTE", "QUOTE", quote_decimals);
            BASE = new T20("BASE", "BASE", base_decimals);
        }
        {
            // create market with V2 fees (4가지 수수료)
            bytes memory marketInitializeData = abi.encode(
                uint32(MARKET_SELLER_MAKER_FEE),
                uint32(MARKET_SELLER_TAKER_FEE),
                uint32(MARKET_BUYER_MAKER_FEE),
                uint32(MARKET_BUYER_TAKER_FEE)
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
                NO_FEE_BPS, // use market buyer maker fee
                NO_FEE_BPS // use market buyer taker fee
            );
            PAIR = PairImplV2(pair);
        }

        // Setup initial balances
        QUOTE.transfer(USER1, _toQuote(10000));
        QUOTE.transfer(USER2, _toQuote(10000));
        BASE.transfer(USER1, _toBase(10000));
        BASE.transfer(USER2, _toBase(10000));

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

    // Test MarketImplV2 initialization with 4가지 수수료
    function test_marketV2_initialization() external {
        IMarketV2.FeeConfig memory feeConfig = MARKET.getFeeConfig();
        assertEq(feeConfig.sellerMakerFeeBps, MARKET_SELLER_MAKER_FEE, "Seller maker fee should match");
        assertEq(feeConfig.sellerTakerFeeBps, MARKET_SELLER_TAKER_FEE, "Seller taker fee should match");
        assertEq(feeConfig.buyerMakerFeeBps, MARKET_BUYER_MAKER_FEE, "Buyer maker fee should match");
        assertEq(feeConfig.buyerTakerFeeBps, MARKET_BUYER_TAKER_FEE, "Buyer taker fee should match");
    }

    // Test PairImplV2 initialization with default fees (use market fees)
    function test_pairV2_initialization_with_market_fees() external {
        (uint32 pairSellerMakerFee, uint32 pairSellerTakerFee, uint32 pairBuyerMakerFee, uint32 pairBuyerTakerFee) =
            PAIR.feeConfig();
        assertEq(pairSellerMakerFee, NO_FEE_BPS, "Pair seller maker fee should be NO_FEE_BPS (use market)");
        assertEq(pairSellerTakerFee, NO_FEE_BPS, "Pair seller taker fee should be NO_FEE_BPS (use market)");
        assertEq(pairBuyerMakerFee, NO_FEE_BPS, "Pair buyer maker fee should be NO_FEE_BPS (use market)");
        assertEq(pairBuyerTakerFee, NO_FEE_BPS, "Pair buyer taker fee should be NO_FEE_BPS (use market)");

        (
            uint32 effectiveSellerMakerFee,
            uint32 effectiveSellerTakerFee,
            uint32 effectiveBuyerMakerFee,
            uint32 effectiveBuyerTakerFee
        ) = PAIR.getEffectiveFees();
        assertEq(effectiveSellerMakerFee, MARKET_SELLER_MAKER_FEE, "Effective seller maker fee should use market fee");
        assertEq(effectiveSellerTakerFee, MARKET_SELLER_TAKER_FEE, "Effective seller taker fee should use market fee");
        assertEq(effectiveBuyerMakerFee, MARKET_BUYER_MAKER_FEE, "Effective buyer maker fee should use market fee");
        assertEq(effectiveBuyerTakerFee, MARKET_BUYER_TAKER_FEE, "Effective buyer taker fee should use market fee");
    }

    // Test setting market fees
    function test_marketV2_set_fees() external {
        uint32 newSellerMakerFee = 15; // 0.15%
        uint32 newSellerTakerFee = 25; // 0.25%
        uint32 newBuyerMakerFee = 12; // 0.12%
        uint32 newBuyerTakerFee = 22; // 0.22%

        vm.prank(OWNER);
        MARKET.setMarketFees(newSellerMakerFee, newSellerTakerFee, newBuyerMakerFee, newBuyerTakerFee);

        IMarketV2.FeeConfig memory updatedFeeConfig = MARKET.getFeeConfig();
        assertEq(updatedFeeConfig.sellerMakerFeeBps, newSellerMakerFee, "Seller maker fee should be updated");
        assertEq(updatedFeeConfig.sellerTakerFeeBps, newSellerTakerFee, "Seller taker fee should be updated");
        assertEq(updatedFeeConfig.buyerMakerFeeBps, newBuyerMakerFee, "Buyer maker fee should be updated");
        assertEq(updatedFeeConfig.buyerTakerFeeBps, newBuyerTakerFee, "Buyer taker fee should be updated");
    }

    // Test setting pair-specific fees
    function test_pairV2_set_pair_specific_fees() external {
        uint32 pairSellerMakerFee = 10; // 0.1%
        uint32 pairSellerTakerFee = 40; // 0.4%
        uint32 pairBuyerMakerFee = 8; // 0.08%
        uint32 pairBuyerTakerFee = 35; // 0.35%

        vm.prank(OWNER);
        PAIR.setPairFees(pairSellerMakerFee, pairSellerTakerFee, pairBuyerMakerFee, pairBuyerTakerFee);

        (uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee) = PAIR.feeConfig();
        assertEq(sellerMakerFee, pairSellerMakerFee, "Pair seller maker fee should be updated");
        assertEq(sellerTakerFee, pairSellerTakerFee, "Pair seller taker fee should be updated");
        assertEq(buyerMakerFee, pairBuyerMakerFee, "Pair buyer maker fee should be updated");
        assertEq(buyerTakerFee, pairBuyerTakerFee, "Pair buyer taker fee should be updated");

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

    // Test fee resolution: pair fees == NO_FEE_BPS should use market fees
    function test_pairV2_fee_resolution_fallback_to_market() external {
        vm.prank(OWNER);
        PAIR.setPairFees(NO_FEE_BPS, NO_FEE_BPS, NO_FEE_BPS, NO_FEE_BPS); // All == NO_FEE_BPS, should use market fees

        (
            uint32 effectiveSellerMakerFee,
            uint32 effectiveSellerTakerFee,
            uint32 effectiveBuyerMakerFee,
            uint32 effectiveBuyerTakerFee
        ) = PAIR.getEffectiveFees();
        assertEq(effectiveSellerMakerFee, MARKET_SELLER_MAKER_FEE, "Should fallback to market seller maker fee");
        assertEq(effectiveSellerTakerFee, MARKET_SELLER_TAKER_FEE, "Should fallback to market seller taker fee");
        assertEq(effectiveBuyerMakerFee, MARKET_BUYER_MAKER_FEE, "Should fallback to market buyer maker fee");
        assertEq(effectiveBuyerTakerFee, MARKET_BUYER_TAKER_FEE, "Should fallback to market buyer taker fee");
    }

    // Test seller maker fee application on limit orders
    function test_pairV2_seller_maker_fee_on_limit_order() external {
        uint32 pairSellerMakerFee = 10; // 0.1%
        uint32 pairSellerTakerFee = 40; // 0.4%
        uint32 pairBuyerMakerFee = 8; // 0.08%
        uint32 pairBuyerTakerFee = 35; // 0.35%

        vm.prank(OWNER);
        PAIR.setPairFees(pairSellerMakerFee, pairSellerTakerFee, pairBuyerMakerFee, pairBuyerTakerFee);

        uint256 initialQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // USER1 places limit sell order (seller maker)
        vm.prank(USER1);
        uint256 sellOrderId = ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER2 places buy market order (buyer taker) - should match with USER1's limit order
        vm.prank(USER2);
        ROUTER.submitBuyMarket(address(PAIR), _toQuote(100), type(uint256).max);

        uint256 finalQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 feeCollected = finalQuoteBalance - initialQuoteBalance;

        // Fee should be calculated based on seller maker fee (since USER1's limit order acts as seller maker)
        uint256 expectedFee = Math.mulDiv(_toQuote(100), pairSellerMakerFee, 10000);
        assertEq(feeCollected, expectedFee, "Fee should be calculated using seller maker fee rate");
    }

    // Test taker fee application on market orders
    function test_pairV2_taker_fee_on_market_order() external {
        uint32 pairSellerMakerFee = 10; // 0.1%
        uint32 pairSellerTakerFee = 40; // 0.4%
        uint32 pairBuyerMakerFee = 15; // 0.15%
        uint32 pairBuyerTakerFee = 50; // 0.5%

        vm.prank(OWNER);
        PAIR.setPairFees(pairSellerMakerFee, pairSellerTakerFee, pairBuyerMakerFee, pairBuyerTakerFee);

        // USER1 places limit buy order (buyer maker)
        vm.prank(USER1);
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 initialQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // USER2 places sell market order (seller taker) - should match with USER1's limit order
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(100), type(uint256).max);

        uint256 finalQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 feeCollected = finalQuoteBalance - initialQuoteBalance;

        // Fee should be calculated based on both fees:
        // Buyer (USER1) maker fee + Seller (USER2) taker fee
        uint256 expectedBuyerFee = Math.mulDiv(_toQuote(100), pairBuyerMakerFee, 10000);
        uint256 expectedSellerFee = Math.mulDiv(_toQuote(100), pairSellerTakerFee, 10000);
        uint256 expectedTotalFee = expectedBuyerFee + expectedSellerFee;
        assertEq(
            feeCollected, expectedTotalFee, "Fee should be calculated using both buyer maker and seller taker fees"
        );
    }

    // Test mixed scenario: limit-limit order matching (seller maker fee)
    function test_pairV2_limit_limit_order_matching() external {
        uint32 pairSellerMakerFee = 10; // 0.1%
        uint32 pairSellerTakerFee = 40; // 0.4%
        uint32 pairBuyerMakerFee = 8; // 0.08%
        uint32 pairBuyerTakerFee = 35; // 0.35%

        vm.prank(OWNER);
        PAIR.setPairFees(pairSellerMakerFee, pairSellerTakerFee, pairBuyerMakerFee, pairBuyerTakerFee);

        uint256 initialQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // USER1 places limit sell order first (seller maker)
        vm.prank(USER1);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER2 places limit buy order that matches (buyer taker)
        vm.prank(USER2);
        ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 finalQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 feeCollected = finalQuoteBalance - initialQuoteBalance;

        // Fee should be based on seller maker fee (USER1's order was placed first as seller)
        uint256 expectedFee = Math.mulDiv(_toQuote(100), pairSellerMakerFee, 10000);
        assertEq(feeCollected, expectedFee, "Fee should use seller maker rate for first-placed sell order");
    }

    // Test zero fees
    function test_pairV2_zero_fees() external {
        vm.prank(OWNER);
        PAIR.setPairFees(0, 0, 0, 0); // Zero fees for all

        uint256 initialQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);

        // Execute a trade
        vm.prank(USER1);
        ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER2);
        ROUTER.submitBuyMarket(address(PAIR), _toQuote(100), type(uint256).max);

        uint256 finalQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);
        uint256 feeCollected = finalQuoteBalance - initialQuoteBalance;

        assertEq(feeCollected, 0, "No fees should be collected when fees are set to 0");
    }

    // Test backward compatibility with market fee updates
    function test_pairV2_market_fee_propagation() external {
        // Keep pair fees at default (NO_FEE_BPS = use market fees)
        (uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps) =
            PAIR.feeConfig();
        assertEq(sellerMakerFeeBps, NO_FEE_BPS, "Pair should use market fees for seller maker");
        assertEq(sellerTakerFeeBps, NO_FEE_BPS, "Pair should use market fees for seller taker");
        assertEq(buyerMakerFeeBps, NO_FEE_BPS, "Pair should use market fees for buyer maker");
        assertEq(buyerTakerFeeBps, NO_FEE_BPS, "Pair should use market fees for buyer taker");

        // Update market fees
        uint32 newSellerMakerFee = 50; // 0.5%
        uint32 newSellerTakerFee = 80; // 0.8%
        uint32 newBuyerMakerFee = 60; // 0.6%
        uint32 newBuyerTakerFee = 90; // 0.9%

        vm.prank(OWNER);
        MARKET.setMarketFees(newSellerMakerFee, newSellerTakerFee, newBuyerMakerFee, newBuyerTakerFee);

        // Effective fees should reflect the new market fees
        (
            uint32 effectiveSellerMakerFee,
            uint32 effectiveSellerTakerFee,
            uint32 effectiveBuyerMakerFee,
            uint32 effectiveBuyerTakerFee
        ) = PAIR.getEffectiveFees();
        assertEq(effectiveSellerMakerFee, newSellerMakerFee, "Effective seller maker fee should use updated market fee");
        assertEq(effectiveSellerTakerFee, newSellerTakerFee, "Effective seller taker fee should use updated market fee");
        assertEq(effectiveBuyerMakerFee, newBuyerMakerFee, "Effective buyer maker fee should use updated market fee");
        assertEq(effectiveBuyerTakerFee, newBuyerTakerFee, "Effective buyer taker fee should use updated market fee");
    }

    // Test creating new pair with custom fees
    function test_pairV2_create_pair_with_custom_fees() external {
        IERC20 newBase = new T20("NEWBASE", "NEWBASE", 18);

        vm.prank(OWNER);
        address newPairAddr = MARKET.createPair(
            address(newBase),
            QUOTE_DECIMALS / 1e2,
            BASE_DECIMALS / 1e6,
            25,
            35,
            30,
            40 // seller maker, seller taker, buyer maker, buyer taker
        );

        PairImplV2 newPair = PairImplV2(newPairAddr);

        // Check pair fees
        {
            (uint32 pairSellerMakerFee, uint32 pairSellerTakerFee, uint32 pairBuyerMakerFee, uint32 pairBuyerTakerFee) =
                newPair.feeConfig();
            assertEq(pairSellerMakerFee, 25, "New pair should have custom seller maker fee");
            assertEq(pairSellerTakerFee, 35, "New pair should have custom seller taker fee");
            assertEq(pairBuyerMakerFee, 30, "New pair should have custom buyer maker fee");
            assertEq(pairBuyerTakerFee, 40, "New pair should have custom buyer taker fee");
        }

        // Check effective fees
        {
            (
                uint32 effectiveSellerMakerFee,
                uint32 effectiveSellerTakerFee,
                uint32 effectiveBuyerMakerFee,
                uint32 effectiveBuyerTakerFee
            ) = newPair.getEffectiveFees();
            assertEq(effectiveSellerMakerFee, 25, "Effective seller maker fee should use custom fee");
            assertEq(effectiveSellerTakerFee, 35, "Effective seller taker fee should use custom fee");
            assertEq(effectiveBuyerMakerFee, 30, "Effective buyer maker fee should use custom fee");
            assertEq(effectiveBuyerTakerFee, 40, "Effective buyer taker fee should use custom fee");
        }
    }

    // Test access control for fee setting functions
    function test_pairV2_access_control() external {
        vm.prank(USER1);
        vm.expectRevert();
        PAIR.setPairFees(50, 60, 70, 80);

        vm.prank(USER1);
        vm.expectRevert();
        MARKET.setMarketFees(50, 60, 70, 80);
    }

    // Test events are emitted correctly
    function test_pairV2_events() external {
        uint32 newSellerMakerFee = 15;
        uint32 newSellerTakerFee = 25;
        uint32 newBuyerMakerFee = 35;
        uint32 newBuyerTakerFee = 45;

        vm.prank(OWNER);
        vm.expectEmit(true, true, false, false);
        emit PairFeesUpdated(newSellerMakerFee, newSellerTakerFee, newBuyerMakerFee, newBuyerTakerFee);
        PAIR.setPairFees(newSellerMakerFee, newSellerTakerFee, newBuyerMakerFee, newBuyerTakerFee);

        vm.prank(OWNER);
        vm.expectEmit(true, true, false, false);
        emit MarketFeesUpdated(newSellerMakerFee, newSellerTakerFee, newBuyerMakerFee, newBuyerTakerFee);
        MARKET.setMarketFees(newSellerMakerFee, newSellerTakerFee, newBuyerMakerFee, newBuyerTakerFee);
    }

    // Events (to make the test compile)
    event PairFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee);
    event MarketFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee);
}
