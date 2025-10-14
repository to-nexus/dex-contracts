// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {CrossDexImpl} from "../src/CrossDexImpl.sol";
import {CrossDexRouter} from "../src/CrossDexRouter.sol";
import {MarketImplV2} from "../src/MarketImplV2.sol";
import {PairImplV2} from "../src/PairImplV2.sol";
import {WETH} from "../src/WETH.sol";

import {IMarketV2, NO_FEE_BPS} from "../src/interfaces/IMarket.sol";
import {IPair} from "../src/interfaces/IPair.sol";

import {T20} from "./mock/T20.sol";

// contract DEXFeeV2Test is Test {
//     address public constant OWNER = address(bytes20("OWNER"));
//     address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
//     address public constant USER1 = address(bytes20("USER1"));
//     address public constant USER2 = address(bytes20("USER2"));

//     CrossDexImpl public CROSS_DEX;
//     CrossDexRouter public ROUTER;
//     WETH public CROSS;

//     IERC20 public QUOTE;
//     IERC20 public BASE;

//     MarketImplV2 public MARKET;
//     PairImplV2 public PAIR;

//     uint256 public FIND_PREV_PRICE_COUNT = type(uint256).max;
//     uint256 public MAX_MATCH_COUNT = type(uint256).max;
//     uint256 public CANCEL_LIMIT = type(uint256).max;

//     // Market-level fees (in BPS)
//     uint256 public MARKET_MAKER_FEE = 20; // 0.2%
//     uint256 public MARKET_TAKER_FEE = 30; // 0.3%

//     uint256 public QUOTE_DECIMALS;
//     uint256 public BASE_DECIMALS;

//     uint256[2] internal _searchPrices;

//     function setUp() external {
//         _deployV2(18, 18, 1e2, 1e6);
//     }

//     function _deployV2(uint8 quote_decimals, uint8 base_decimals, uint256 quote_tick_size, uint256 base_tick_size)
//         internal
//     {
//         vm.label(OWNER, "owner");
//         vm.label(FEE_COLLECTOR, "feeCollector");
//         vm.label(USER1, "user1");
//         vm.label(USER2, "user2");

//         vm.startPrank(OWNER);

//         QUOTE_DECIMALS = 10 ** quote_decimals;
//         BASE_DECIMALS = 10 ** base_decimals;

//         {
//             // deploy impl contracts (using V2 versions)
//             address routerImpl = address(new CrossDexRouter());
//             address marketImpl = address(new MarketImplV2());
//             address pairImpl = address(new PairImplV2());

//             // deploy cross dex
//             address crossDexImpl = address(new CrossDexImpl());
//             ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
//             CROSS_DEX = CrossDexImpl(address(proxy));
//             CROSS_DEX.initialize(
//                 OWNER, routerImpl, FIND_PREV_PRICE_COUNT, MAX_MATCH_COUNT, CANCEL_LIMIT, marketImpl, pairImpl
//             );
//         }
//         {
//             // get contracts from CROSS_DEX
//             ROUTER = CrossDexRouter(CROSS_DEX.ROUTER());
//             CROSS = WETH(payable(address(ROUTER.CROSS())));
//         }
//         {
//             // deploy base and quote tokens
//             QUOTE = new T20("QUOTE", "QUOTE", quote_decimals);
//             BASE = new T20("BASE", "BASE", base_decimals);
//         }
//         {
//             // create market with V2 fees (maker and taker)
//             address market = CROSS_DEX.createMarket(OWNER, address(QUOTE), FEE_COLLECTOR, MARKET_MAKER_FEE);
//             MARKET = MarketImplV2(market);

//             // Set market-level taker fee (different from maker fee)
//             MARKET.setMarketFees(uint32(MARKET_MAKER_FEE), uint32(MARKET_TAKER_FEE));

//             // create pair with default fees (NO_FEE_BPS = use market fees)
//             address pair = MARKET.createPair(
//                 address(BASE),
//                 QUOTE_DECIMALS / quote_tick_size,
//                 BASE_DECIMALS / base_tick_size,
//                 NO_FEE_BPS, // use market maker fee
//                 NO_FEE_BPS // use market taker fee
//             );
//             PAIR = PairImplV2(pair);
//         }

//         // Setup initial balances
//         QUOTE.transfer(USER1, _toQuote(10000));
//         QUOTE.transfer(USER2, _toQuote(10000));
//         BASE.transfer(USER1, _toBase(10000));
//         BASE.transfer(USER2, _toBase(10000));

//         vm.stopPrank();

//         // Setup approvals
//         vm.prank(USER1);
//         QUOTE.approve(address(ROUTER), type(uint256).max);
//         vm.prank(USER1);
//         BASE.approve(address(ROUTER), type(uint256).max);

//         vm.prank(USER2);
//         QUOTE.approve(address(ROUTER), type(uint256).max);
//         vm.prank(USER2);
//         BASE.approve(address(ROUTER), type(uint256).max);
//     }

//     function _toBase(uint256 x) internal view returns (uint256) {
//         return x * BASE_DECIMALS;
//     }

//     function _toQuote(uint256 x) internal view returns (uint256) {
//         return x * QUOTE_DECIMALS;
//     }

//     // Test MarketImplV2 initialization with separate maker/taker fees
//     function test_marketV2_initialization() external {
//         (uint32 makerFee, uint32 takerFee) = MARKET.getMarketFees();
//         assertEq(makerFee, MARKET_MAKER_FEE, "Maker fee should match");
//         assertEq(takerFee, MARKET_TAKER_FEE, "Taker fee should match");

//         assertEq(MARKET.makerFeeBps(), MARKET_MAKER_FEE, "Legacy feeBps should be maker fee");
//         assertEq(MARKET.takerFeeBps(), MARKET_TAKER_FEE, "Taker fee should be set correctly");
//     }

//     // Test PairImplV2 initialization with default fees (use market fees)
//     function test_pairV2_initialization_with_market_fees() external {
//         (uint32 pairMakerFee, uint32 pairTakerFee) = PAIR.getPairFees();
//         assertEq(pairMakerFee, NO_FEE_BPS, "Pair maker fee should be NO_FEE_BPS (use market)");
//         assertEq(pairTakerFee, NO_FEE_BPS, "Pair taker fee should be NO_FEE_BPS (use market)");

//         (uint32 effectiveMakerFee, uint32 effectiveTakerFee) = PAIR.getEffectiveFees();
//         assertEq(effectiveMakerFee, MARKET_MAKER_FEE, "Effective maker fee should use market fee");
//         assertEq(effectiveTakerFee, MARKET_TAKER_FEE, "Effective taker fee should use market fee");
//     }

//     // Test setting market fees
//     function test_marketV2_set_fees() external {
//         uint32 newMakerFee = 15; // 0.15%
//         uint32 newTakerFee = 25; // 0.25%

//         vm.prank(OWNER);
//         MARKET.setMarketFees(newMakerFee, newTakerFee);

//         (uint32 makerFee, uint32 takerFee) = MARKET.getMarketFees();
//         assertEq(makerFee, newMakerFee, "Maker fee should be updated");
//         assertEq(takerFee, newTakerFee, "Taker fee should be updated");
//     }

//     // Test setting pair-specific fees
//     function test_pairV2_set_pair_specific_fees() external {
//         uint32 pairMakerFee = 10; // 0.1%
//         uint32 pairTakerFee = 40; // 0.4%

//         vm.prank(OWNER);
//         PAIR.setPairFees(pairMakerFee, pairTakerFee);

//         (uint32 makerFee, uint32 takerFee) = PAIR.getPairFees();
//         assertEq(makerFee, pairMakerFee, "Pair maker fee should be updated");
//         assertEq(takerFee, pairTakerFee, "Pair taker fee should be updated");

//         (uint32 effectiveMakerFee, uint32 effectiveTakerFee) = PAIR.getEffectiveFees();
//         assertEq(effectiveMakerFee, pairMakerFee, "Effective maker fee should use pair fee");
//         assertEq(effectiveTakerFee, pairTakerFee, "Effective taker fee should use pair fee");
//     }

//     // Test fee resolution: pair fees == NO_FEE_BPS should use market fees
//     function test_pairV2_fee_resolution_fallback_to_market() external {
//         vm.prank(OWNER);
//         PAIR.setPairFees(NO_FEE_BPS, NO_FEE_BPS); // Both == NO_FEE_BPS, should use market fees

//         (uint32 effectiveMakerFee, uint32 effectiveTakerFee) = PAIR.getEffectiveFees();
//         assertEq(effectiveMakerFee, MARKET_MAKER_FEE, "Should fallback to market maker fee");
//         assertEq(effectiveTakerFee, MARKET_TAKER_FEE, "Should fallback to market taker fee");
//     }

//     // Test maker fee application on limit orders
//     function test_pairV2_maker_fee_on_limit_order() external {
//         uint32 pairMakerFee = 10; // 0.1%
//         uint32 pairTakerFee = 40; // 0.4%

//         vm.prank(OWNER);
//         PAIR.setPairFees(pairMakerFee, pairTakerFee);

//         uint256 initialQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);

//         // USER1 places limit sell order (maker)
//         vm.prank(USER1);
//         uint256 sellOrderId = ROUTER.submitSellLimit(
//             address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
//         );

//         // USER2 places buy market order (taker) - should match with USER1's limit order
//         vm.prank(USER2);
//         ROUTER.submitBuyMarket(address(PAIR), _toQuote(100), type(uint256).max);

//         uint256 finalQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);
//         uint256 feeCollected = finalQuoteBalance - initialQuoteBalance;

//         // Fee should be calculated based on maker fee (since USER1's limit order acts as maker)
//         uint256 expectedFee = Math.mulDiv(_toQuote(100), pairMakerFee, 10000);
//         assertEq(feeCollected, expectedFee, "Fee should be calculated using maker fee rate");
//     }

//     // Test taker fee application on market orders
//     function test_pairV2_taker_fee_on_market_order() external {
//         uint32 pairMakerFee = 10; // 0.1%
//         uint32 pairTakerFee = 40; // 0.4%

//         vm.prank(OWNER);
//         PAIR.setPairFees(pairMakerFee, pairTakerFee);

//         // USER1 places limit buy order (maker)
//         vm.prank(USER1);
//         ROUTER.submitBuyLimit(
//             address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
//         );

//         uint256 initialQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);

//         // USER2 places sell market order (taker) - should match with USER1's limit order
//         vm.prank(USER2);
//         ROUTER.submitSellMarket(address(PAIR), _toBase(100), type(uint256).max);

//         uint256 finalQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);
//         uint256 feeCollected = finalQuoteBalance - initialQuoteBalance;

//         // Fee should be calculated based on taker fee (since USER2's market order acts as taker)
//         uint256 expectedFee = Math.mulDiv(_toQuote(100), pairTakerFee, 10000);
//         assertEq(feeCollected, expectedFee, "Fee should be calculated using taker fee rate");
//     }

//     // Test mixed scenario: limit-limit order matching
//     function test_pairV2_limit_limit_order_matching() external {
//         uint32 pairMakerFee = 10; // 0.1%
//         uint32 pairTakerFee = 40; // 0.4%

//         vm.prank(OWNER);
//         PAIR.setPairFees(pairMakerFee, pairTakerFee);

//         uint256 initialQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);

//         // USER1 places limit sell order first (maker)
//         vm.prank(USER1);
//         ROUTER.submitSellLimit(
//             address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
//         );

//         // USER2 places limit buy order that matches (taker)
//         vm.prank(USER2);
//         ROUTER.submitBuyLimit(
//             address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
//         );

//         uint256 finalQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);
//         uint256 feeCollected = finalQuoteBalance - initialQuoteBalance;

//         // Fee should be based on maker fee (USER1's order was placed first)
//         uint256 expectedFee = Math.mulDiv(_toQuote(100), pairMakerFee, 10000);
//         assertEq(feeCollected, expectedFee, "Fee should use maker rate for first-placed order");
//     }

//     // Test zero fees
//     function test_pairV2_zero_fees() external {
//         vm.prank(OWNER);
//         PAIR.setPairFees(0, 0); // Zero fees for both maker and taker

//         uint256 initialQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);

//         // Execute a trade
//         vm.prank(USER1);
//         ROUTER.submitSellLimit(
//             address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
//         );

//         vm.prank(USER2);
//         ROUTER.submitBuyMarket(address(PAIR), _toQuote(100), type(uint256).max);

//         uint256 finalQuoteBalance = QUOTE.balanceOf(FEE_COLLECTOR);
//         uint256 feeCollected = finalQuoteBalance - initialQuoteBalance;

//         assertEq(feeCollected, 0, "No fees should be collected when fees are set to 0");
//     }

//     // Test backward compatibility with market fee updates
//     function test_pairV2_market_fee_propagation() external {
//         // Keep pair fees at default (NO_FEE_BPS = use market fees)
//         (uint32 makerFeeBps, uint32 takerFeeBps) = PAIR.getPairFees();
//         assertEq(makerFeeBps, NO_FEE_BPS, "Pair should use market fees");
//         assertEq(takerFeeBps, NO_FEE_BPS, "Pair should use market fees");

//         // Update market fees
//         uint32 newMakerFee = 50; // 0.5%
//         uint32 newTakerFee = 80; // 0.8%

//         vm.prank(OWNER);
//         MARKET.setMarketFees(newMakerFee, newTakerFee);

//         // Effective fees should reflect the new market fees
//         (uint32 effectiveMakerFee, uint32 effectiveTakerFee) = PAIR.getEffectiveFees();
//         assertEq(effectiveMakerFee, newMakerFee, "Effective maker fee should use updated market fee");
//         assertEq(effectiveTakerFee, newTakerFee, "Effective taker fee should use updated market fee");
//     }

//     // Test creating new pair with custom fees
//     function test_pairV2_create_pair_with_custom_fees() external {
//         uint32 customMakerFee = 25; // 0.25%
//         uint32 customTakerFee = 35; // 0.35%

//         IERC20 newBase = new T20("NEWBASE", "NEWBASE", 18);

//         vm.prank(OWNER);
//         address newPairAddr = MARKET.createPair(
//             address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, customMakerFee, customTakerFee
//         );

//         PairImplV2 newPair = PairImplV2(newPairAddr);

//         (uint32 pairMakerFee, uint32 pairTakerFee) = newPair.getPairFees();
//         assertEq(pairMakerFee, customMakerFee, "New pair should have custom maker fee");
//         assertEq(pairTakerFee, customTakerFee, "New pair should have custom taker fee");

//         (uint32 effectiveMakerFee, uint32 effectiveTakerFee) = newPair.getEffectiveFees();
//         assertEq(effectiveMakerFee, customMakerFee, "Effective maker fee should use custom fee");
//         assertEq(effectiveTakerFee, customTakerFee, "Effective taker fee should use custom fee");
//     }

//     // Test access control for fee setting functions
//     function test_pairV2_access_control() external {
//         vm.prank(USER1);
//         vm.expectRevert();
//         PAIR.setPairFees(50, 60);

//         vm.prank(USER1);
//         vm.expectRevert();
//         MARKET.setMarketFees(50, 60);
//     }

//     // Test events are emitted correctly
//     function test_pairV2_events() external {
//         uint32 newMakerFee = 15;
//         uint32 newTakerFee = 25;

//         vm.prank(OWNER);
//         vm.expectEmit(true, true, false, false);
//         emit PairFeesUpdated(newMakerFee, newTakerFee);
//         PAIR.setPairFees(newMakerFee, newTakerFee);

//         vm.prank(OWNER);
//         vm.expectEmit(true, true, false, false);
//         emit MarketFeesUpdated(newMakerFee, newTakerFee);
//         MARKET.setMarketFees(newMakerFee, newTakerFee);
//     }

//     // Events (to make the test compile)
//     event PairFeesUpdated(uint32 indexed makerFee, uint32 indexed takerFee);
//     event MarketFeesUpdated(uint32 indexed makerFee, uint32 indexed takerFee);
// }
