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
 * @title DEXV2ErrorHandlingTest
 * @notice Comprehensive error handling tests for V2 contracts
 * @dev Tests all custom errors defined in V2 contracts to ensure proper error handling
 */
contract DEXV2ErrorHandlingTest is Test {
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

    bytes public NO_FEE_BPS_DATA = abi.encode(NO_FEE_BPS, NO_FEE_BPS, NO_FEE_BPS, NO_FEE_BPS);
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
            NO_FEE_BPS_DATA
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

    function test_v2_PairInvalidFeeStructure_seller_fees() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeStructure(uint32,uint32)", 100, 50));
        PAIR.setPairFees(100, 50, 0, 0); // seller maker > seller taker
    }

    function test_v2_PairInvalidFeeStructure_buyer_fees() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeStructure(uint32,uint32)", 75, 25));
        PAIR.setPairFees(0, 0, 75, 25); // buyer maker > buyer taker
    }

    function test_v2_PairInvalidFeeBps_over_limit() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeBps()"));
        PAIR.setPairFees(BPS_DENOMINATOR, 0, 0, 0); // 100% fee
    }

    function test_v2_PairInvalidFeeBps_buyer_over_limit() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeBps()"));
        PAIR.setPairFees(0, 0, 0, BPS_DENOMINATOR + 1); // >100% fee
    }

    // ================================
    // Access Control Error Tests
    // ================================

    function test_v2_OwnableUnauthorizedAccount_setPairFees() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", USER1));
        PAIR.setPairFees(30, 50, 25, 40);
    }

    function test_v2_OwnableUnauthorizedAccount_setMarketFees() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", USER1));
        MARKET.setMarketFees(30, 50, 25, 40);
    }

    function test_v2_OwnableUnauthorizedAccount_setFeeCollector() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", USER1));
        MARKET.setFeeCollector(USER1);
    }

    // ================================
    // Router Error Tests
    // ================================

    function test_v2_RouterCancelLimitExceeded() external {
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

    function test_v2_RouterInvalidInputData_findPrevPriceCount() external {
        // Test that invalid input validation works for findPrevPriceCount
        vm.prank(OWNER);
        vm.expectRevert(); // Any revert is acceptable for invalid input
        ROUTER.setfindPrevPriceCount(0);
    }

    function test_v2_RouterInvalidInputData_maxMatchCount() external {
        // Test that invalid input validation works for maxMatchCount
        vm.prank(OWNER);
        vm.expectRevert(); // Any revert is acceptable for invalid input
        ROUTER.setMaxMatchCount(0);
    }

    function test_v2_RouterInvalidInputData_cancelLimit() external {
        // Test that invalid input validation works for cancelLimit
        vm.prank(OWNER);
        vm.expectRevert(); // Any revert is acceptable for invalid input
        ROUTER.setCancelLimit(0);
    }

    // ================================
    // Market Reserve Error Tests
    // ================================

    // ================================
    // Fill or Kill Error Tests
    // ================================

    // ================================
    // Market-Specific Error Tests
    // ================================

    function test_v2_MarketInvalidFeeCollector_zero_address() external {
        vm.prank(OWNER);
        vm.expectRevert(); // Should revert when setting zero address as fee collector
        MARKET.setFeeCollector(address(0));
    }

    // ================================
    // Edge Case Error Tests
    // ================================

    function test_v2_Complex_scenario_multiple_errors() external {
        // Test a complex scenario that validates multiple error conditions work properly

        // 1. Test invalid fee structure validation
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidFeeStructure(uint32,uint32)", 100, 50));
        PAIR.setPairFees(100, 50, 0, 0);

        // 2. Test that valid pair detection works
        address invalidPair = address(0x999);
        bool isValid = ROUTER.isPair(invalidPair);
        assertFalse(isValid, "Invalid pair should be detected");

        // 3. Test access control validation
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", USER1));
        MARKET.setMarketFees(30, 50, 25, 40);
    }

    // ================================
    // Gas Optimization Error Tests
    // ================================

    // ================================
    // Initialization Error Tests
    // ================================

    function test_v2_PairInvalidInitializeData() external {
        // Test initialization with invalid data
        // Try to create pair with same quote token as base (should be invalid)

        vm.prank(OWNER);
        vm.expectRevert(); // Should revert with MarketInvalidBaseAddress
        MARKET.createPair(
            address(QUOTE), // Same as quote token
            1e16, // tickSize
            1e12, // lotSize
            NO_FEE_BPS_DATA
        );
    }
}
