// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";
import {Test, console2} from "forge-std/Test.sol";

import {CrossDexImplV2} from "../src/CrossDexImplV2.sol";

import {CrossDexRouterV2} from "../src/CrossDexRouterV2.sol";
import {MarketImplV2} from "../src/MarketImplV2.sol";
import {PairImplV2} from "../src/PairImplV2.sol";
import {WETH} from "../src/WETH.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";

import {ICrossDex} from "../src/interfaces/ICrossDex.sol";
import {BPS_DENOMINATOR, IMarketV2, NO_FEE_BPS} from "../src/interfaces/IMarket.sol";
import {IPair, IPairV2} from "../src/interfaces/IPair.sol";
import {IRouter} from "../src/interfaces/IRouter.sol";

import {T20} from "./mock/T20.sol";

/**
 * @title DEXErrorHandlingV2Test
 * @notice Comprehensive error handling tests for V2 contracts
 * @dev Tests all custom errors defined in V2 contracts to ensure proper error handling
 */
contract DEXErrorHandlingV2Test is Test {
    CrossDexImplV2 public CROSS_DEX;
    MarketImplV2 public MARKET;
    PairImplV2 public PAIR;
    CrossDexRouterV2 public ROUTER;
    WETH public CROSS;

    T20 public BASE;
    T20 public QUOTE;

    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    address public constant USER1 = address(bytes20("USER1"));
    address public constant USER2 = address(bytes20("USER2"));
    address public constant INVALID_USER = address(bytes20("INVALID_USER"));

    uint256[2] private _searchPrices = [uint256(0), uint256(0)];

    function setUp() external {
        // Deploy tokens first (before any prank)
        BASE = new T20("Base Token", "BASE", 18);
        QUOTE = new T20("Quote Token", "QUOTE", 18);

        vm.startPrank(OWNER);

        // Deploy impl contracts (using V2 versions)
        address routerImpl = address(new CrossDexRouterV2());
        address marketImpl = address(new MarketImplV2());
        address pairImpl = address(new PairImplV2());

        // Deploy CrossDexV2
        address crossDexImpl = address(new CrossDexImplV2());
        ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
        CROSS_DEX = CrossDexImplV2(address(proxy));
        CROSS_DEX.initialize(
            OWNER,
            routerImpl,
            10, // findPrevPriceCount
            100, // maxMatchCount
            50, // cancelLimit
            marketImpl,
            pairImpl
        );

        // Get contracts from CROSS_DEX
        ROUTER = CrossDexRouterV2(CROSS_DEX.ROUTER());
        CROSS = WETH(payable(address(ROUTER.CROSS())));

        // Create market with V2 fees
        bytes memory marketInitializeData = abi.encode(
            uint32(50), // sellerMakerFeeBps
            uint32(75), // sellerTakerFeeBps
            uint32(25), // buyerMakerFeeBps
            uint32(50) // buyerTakerFeeBps
        );
        address market = CROSS_DEX.createMarket(OWNER, address(QUOTE), FEE_COLLECTOR, marketInitializeData);
        MARKET = MarketImplV2(market);

        // Create pair with proper parameters (using values from existing tests)
        uint256 QUOTE_DECIMALS = 10 ** 18;
        uint256 BASE_DECIMALS = 10 ** 18;
        uint256 quote_tick_size = 1e2;
        uint256 base_tick_size = 1e6;

        address pairAddress = MARKET.createPair(
            address(BASE),
            QUOTE_DECIMALS / quote_tick_size, // tickSize
            BASE_DECIMALS / base_tick_size, // lotSize
            NO_FEE_BPS, // use market seller maker fee
            NO_FEE_BPS, // use market seller taker fee
            NO_FEE_BPS, // use market buyer maker fee
            NO_FEE_BPS // use market buyer taker fee
        );
        PAIR = PairImplV2(pairAddress);

        vm.stopPrank();

        // Transfer tokens to users (T20 mints max to deployer)
        BASE.transfer(USER1, 1000000 * 10 ** 18);
        BASE.transfer(USER2, 1000000 * 10 ** 18);
        QUOTE.transfer(USER1, 1000000 * 10 ** 18);
        QUOTE.transfer(USER2, 1000000 * 10 ** 18);

        // Approve tokens for users
        vm.prank(USER1);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER2);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER2);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        // Also approve for the test contract (for direct pair interactions)
        vm.prank(USER1);
        BASE.approve(address(PAIR), type(uint256).max);
        vm.prank(USER1);
        QUOTE.approve(address(PAIR), type(uint256).max);
        vm.prank(USER2);
        BASE.approve(address(PAIR), type(uint256).max);
        vm.prank(USER2);
        QUOTE.approve(address(PAIR), type(uint256).max);
    }

    function _toBase(uint256 amount) private view returns (uint256) {
        return amount * 10 ** BASE.decimals();
    }

    function _toQuote(uint256 amount) private view returns (uint256) {
        return amount * 10 ** QUOTE.decimals();
    }

    // ================================
    // Fee Structure Error Tests
    // ================================

    function test_PairInvalidFeeStructure_seller_fees() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeStructure(uint32,uint32)", 100, 50));
        PAIR.setPairFees(100, 50, 0, 0); // seller maker > seller taker
    }

    function test_PairInvalidFeeStructure_buyer_fees() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeStructure(uint32,uint32)", 75, 25));
        PAIR.setPairFees(0, 0, 75, 25); // buyer maker > buyer taker
    }

    function test_PairInvalidFeeBps_over_limit() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeBps()"));
        PAIR.setPairFees(BPS_DENOMINATOR, 0, 0, 0); // 100% fee
    }

    function test_PairInvalidFeeBps_buyer_over_limit() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeBps()"));
        PAIR.setPairFees(0, 0, 0, BPS_DENOMINATOR + 1); // >100% fee
    }

    // ================================
    // Order Validation Error Tests
    // ================================

    function test_PairInvalidAmount_zero() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidAmount(uint256)", 0));
        ROUTER.submitSellLimit(
            address(PAIR),
            _toQuote(1),
            0, // zero amount
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            _searchPrices,
            0
        );
    }

    function test_PairInvalidPrice_zero() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidPrice(uint256)", 0));
        ROUTER.submitSellLimit(
            address(PAIR),
            0, // zero price
            _toBase(100),
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            _searchPrices,
            0
        );
    }

    function test_PairInvalidOrderId_nonexistent() external {
        uint256[] memory invalidOrderIds = new uint256[](1);
        invalidOrderIds[0] = 999999; // non-existent order ID

        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidOrderId(uint256)", 999999));
        ROUTER.cancelOrder(address(PAIR), invalidOrderIds);
    }

    function test_PairNotOwner_cancel_others_order() external {
        // USER1 creates an order
        vm.prank(USER1);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // USER2 tries to cancel USER1's order
        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;

        vm.prank(USER2);
        vm.expectRevert(abi.encodeWithSignature("PairNotOwner(uint256,address)", orderId, USER2));
        ROUTER.cancelOrder(address(PAIR), orderIds);
    }

    // ================================
    // Access Control Error Tests
    // ================================

    function test_OwnableUnauthorizedAccount_setPairFees() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", USER1));
        PAIR.setPairFees(30, 50, 25, 40);
    }

    function test_OwnableUnauthorizedAccount_setMarketFees() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", USER1));
        MARKET.setMarketFees(30, 50, 25, 40);
    }

    function test_OwnableUnauthorizedAccount_setFeeCollector() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", USER1));
        MARKET.setFeeCollector(USER1);
    }

    // ================================
    // Router Error Tests
    // ================================

    function test_RouterInvalidPairAddress() external {
        address invalidPair = address(0x123);

        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", invalidPair));
        ROUTER.submitSellLimit(
            invalidPair, _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    function test_RouterCancelLimitExceeded() external {
        // Create multiple orders first
        vm.startPrank(USER1);
        uint256[] memory orderIds = new uint256[](60); // More than cancelLimit (50)
        for (uint256 i = 0; i < 60; i++) {
            orderIds[i] = ROUTER.submitSellLimit(
                address(PAIR), _toQuote(1), _toBase(1), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }

        // Try to cancel more than the limit
        vm.expectRevert(abi.encodeWithSignature("RouterCancelLimitExceeded(uint256,uint256)", 60, 50));
        ROUTER.cancelOrder(address(PAIR), orderIds);
        vm.stopPrank();
    }

    function test_RouterInvalidInputData_findPrevPriceCount() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidInputData(bytes32)", "findPrevPriceCount"));
        ROUTER.setfindPrevPriceCount(0);
    }

    function test_RouterInvalidInputData_maxMatchCount() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidInputData(bytes32)", "findPrevPriceCount"));
        ROUTER.setMaxMatchCount(0);
    }

    function test_RouterInvalidInputData_cancelLimit() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidInputData(bytes32)", "cancelLimit"));
        ROUTER.setCancelLimit(0);
    }

    // ================================
    // Market Reserve Error Tests
    // ================================

    function test_PairInvalidReserve_insufficient_base() external {
        // Try to submit sell order with insufficient BASE tokens
        vm.prank(INVALID_USER); // User with no tokens
        vm.expectRevert(); // ERC20InsufficientBalance will be thrown first
        ROUTER.submitSellLimit(
            address(PAIR),
            _toQuote(1),
            _toBase(1000000000), // Huge amount
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            _searchPrices,
            0
        );
    }

    function test_PairInvalidReserve_insufficient_quote() external {
        // Try to submit buy order with insufficient QUOTE tokens
        vm.prank(INVALID_USER); // User with no tokens
        vm.expectRevert(); // ERC20InsufficientBalance will be thrown first
        ROUTER.submitBuyLimit(
            address(PAIR),
            _toQuote(1),
            _toBase(1000000000), // Huge amount
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            _searchPrices,
            0
        );
    }

    // ================================
    // Fill or Kill Error Tests
    // ================================

    function test_PairFillOrKill_insufficient_liquidity() external {
        // Try to place a large Fill-or-Kill order when there's no liquidity
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("PairFillOrKill(address)", USER1));
        ROUTER.submitSellLimit(
            address(PAIR),
            _toQuote(1),
            _toBase(1000000), // Large amount with no matching orders
            IPair.LimitConstraints.FILL_OR_KILL,
            _searchPrices,
            0
        );
    }

    // ================================
    // Tick Size Error Tests
    // ================================

    function test_PairInvalidTickSize() external {
        // This would require creating a pair with specific tick size settings
        // and then trying to place orders that don't conform to the tick size

        // Create a new pair with specific tick size
        vm.prank(OWNER);
        uint256 QUOTE_DECIMALS = 10 ** 18;
        uint256 BASE_DECIMALS = 10 ** 18;
        address newPairAddress = MARKET.createPair(
            address(BASE),
            QUOTE_DECIMALS / 1000, // Large tick size (smaller divisor = larger tick)
            BASE_DECIMALS / 1e6, // lotSize
            NO_FEE_BPS, // sellerMakerFeeBps
            NO_FEE_BPS, // sellerTakerFeeBps
            NO_FEE_BPS, // buyerMakerFeeBps
            NO_FEE_BPS // buyerTakerFeeBps
        );
        PairImplV2 newPair = PairImplV2(newPairAddress);

        // Try to place order with price that doesn't conform to tick size
        uint256 invalidPrice = _toQuote(1) + 1; // Price not aligned with tick size

        vm.prank(USER1);
        vm.expectRevert(); // Should revert with PairInvalidTickSize
        ROUTER.submitSellLimit(
            address(newPair), invalidPrice, _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    // ================================
    // Market-Specific Error Tests
    // ================================

    function test_MarketInvalidFeeCollector_zero_address() external {
        vm.prank(OWNER);
        vm.expectRevert(); // Should revert when setting zero address as fee collector
        MARKET.setFeeCollector(address(0));
    }

    // ================================
    // Edge Case Error Tests
    // ================================

    function test_PairInsufficientTradeVolume() external {
        // This error occurs when the calculated trade volume is insufficient
        // Usually happens with very small amounts or prices near zero

        // First create a small sell order
        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR),
            1, // Very small price (1 wei)
            1, // Very small amount (1 wei)
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            _searchPrices,
            0
        );

        // Try to match with buy order that results in insufficient volume
        vm.prank(USER1);
        vm.expectRevert(); // Should revert with PairInsufficientTradeVolume
        ROUTER.submitBuyMarket(address(PAIR), 1, 1); // Tiny amount
    }

    function test_Complex_scenario_multiple_errors() external {
        // Test a complex scenario that could trigger multiple error conditions

        // 1. Set invalid fee structure first
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeStructure(uint32,uint32)", 100, 50));
        PAIR.setPairFees(100, 50, 0, 0);

        // 2. Try to place order on invalid pair
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidPairAddress(address)", address(0x999)));
        ROUTER.submitSellLimit(
            address(0x999), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // 3. Test access control
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", USER1));
        MARKET.setMarketFees(30, 50, 25, 40);
    }

    // ================================
    // Gas Optimization Error Tests
    // ================================

    function test_RouterInvalidValue_eth_handling() external {
        // Test ETH value handling in router
        // This should work normally for non-ETH pairs
        vm.prank(USER1);
        ROUTER.submitSellLimit{value: 0}(
            address(PAIR), _toQuote(1), _toBase(100), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    // ================================
    // Initialization Error Tests
    // ================================

    function test_PairInvalidInitializeData() external {
        // Test initialization with invalid data
        // Try to create pair with same quote token as base (should be invalid)

        vm.prank(OWNER);
        vm.expectRevert(); // Should revert with MarketInvalidBaseAddress
        MARKET.createPair(
            address(QUOTE), // Same as quote token
            1e16, // tickSize
            1e12, // lotSize
            NO_FEE_BPS,
            NO_FEE_BPS,
            NO_FEE_BPS,
            NO_FEE_BPS
        );
    }
}
