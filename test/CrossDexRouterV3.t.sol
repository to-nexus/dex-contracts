// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {Vm} from "forge-std/Vm.sol";

import {CrossDexRouterV3} from "../src/CrossDexRouterV3.sol";
import {IOwnable} from "../src/interfaces/IOwnable.sol";
import {IPairV3} from "../src/interfaces/IPairV3.sol";
import {DEXV3BaseTest} from "./DEXV3Base.t.sol";

contract ContractCaller {
    function callRouter(CrossDexRouterV3 router, address pair, uint256 price, uint256 amount) external {
        router.submitBuyLimit(
            pair, price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, [uint256(0), uint256(0)], 0
        );
    }
}

/// @dev Contract to force send ETH via selfdestruct (simulates H-01 attack vector)
contract ForceSendEth {
    constructor() payable {}

    function boom(address payable to) external {
        selfdestruct(to);
    }
}

contract CrossDexRouterV3Test is DEXV3BaseTest {
    function setUp() external {
        _deployV3(18, 18, 1e2, 1e6);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // View Functions Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_owner() external view {
        assertEq(ROUTER.owner(), OWNER);
    }

    function test_isPair_valid() external view {
        assertTrue(ROUTER.isPair(address(PAIR)));
    }

    function test_isPair_invalid() external view {
        assertFalse(ROUTER.isPair(address(0)));
        assertFalse(ROUTER.isPair(USER1));
    }

    function test_getRequiredBuyVolume() external {
        uint256 volume = _toQuote(1000);
        uint256 requiredVolume = ROUTER.getRequiredBuyVolume(address(PAIR), volume);
        // With 0 buyer taker fee, should return same volume
        assertEq(requiredVolume, volume);
    }

    function test_findPrevPriceCount() external view {
        assertEq(ROUTER.findPrevPriceCount(), FIND_PREV_PRICE_COUNT);
    }

    function test_maxMatchCount() external view {
        assertEq(ROUTER.maxMatchCount(), MAX_MATCH_COUNT);
    }

    function test_cancelLimit() external view {
        assertEq(ROUTER.cancelLimit(), CANCEL_LIMIT);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Submit Order Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_submitSellLimit_transfersBaseFromUser() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        uint256 user2BaseBefore = BASE.balanceOf(USER2);
        uint256 pairBaseBefore = BASE.balanceOf(address(PAIR));

        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        assertEq(BASE.balanceOf(USER2), user2BaseBefore - amount);
        assertEq(BASE.balanceOf(address(PAIR)), pairBaseBefore + amount);
    }

    function test_submitBuyLimit_transfersQuoteFromUser() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);
        uint256 pairQuoteBefore = QUOTE.balanceOf(address(PAIR));

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        assertEq(QUOTE.balanceOf(USER1), user1QuoteBefore - quoteVolume);
        assertEq(QUOTE.balanceOf(address(PAIR)), pairQuoteBefore + quoteVolume);
    }

    function test_submit_revert_invalidPair() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(CrossDexRouterV3.RouterInvalidPairAddress.selector, address(0)));
        ROUTER.submitBuyLimit(
            address(0), _toQuote(100), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
    }

    function test_submit_revert_contractAccountBlocked() external {
        ContractCaller caller = new ContractCaller();

        // Fund the caller
        vm.prank(OWNER);
        QUOTE.transfer(address(caller), _toQuote(10000));

        vm.prank(address(caller));
        QUOTE.approve(address(ROUTER), type(uint256).max);

        vm.expectRevert(abi.encodeWithSelector(CrossDexRouterV3.RouterContractAccountBlocked.selector, address(caller)));
        caller.callRouter(ROUTER, address(PAIR), _toQuote(100), _toBase(10));
    }

    function test_submit_contractWhitelisted_succeeds() external {
        ContractCaller caller = new ContractCaller();

        // Fund the caller
        vm.prank(OWNER);
        QUOTE.transfer(address(caller), _toQuote(10000));

        vm.prank(address(caller));
        QUOTE.approve(address(ROUTER), type(uint256).max);

        // Whitelist the caller
        address[] memory accounts = new address[](1);
        accounts[0] = address(caller);
        vm.prank(OWNER);
        ROUTER.setWhitelistedCodeAccount(accounts, true);

        // Should succeed now
        caller.callRouter(ROUTER, address(PAIR), _toQuote(100), _toBase(10));

        // Verify order was created
        assertGt(PAIR.quoteReserve(), 0);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Cancel Order Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_cancelOrder_success() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = orderId;

        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), orderIds);

        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, address(0));
    }

    function test_cancelOrder_revert_exceedsCancelLimit() external {
        // Set cancel limit to 2
        vm.prank(OWNER);
        ROUTER.setCancelLimit(2);

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(1);

        // Create 3 orders
        vm.startPrank(USER1);
        uint256[] memory orderIds = new uint256[](3);
        for (uint256 i = 0; i < 3; i++) {
            orderIds[i] = ROUTER.submitBuyLimit(
                address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }

        // Try to cancel all 3 at once
        vm.expectRevert(abi.encodeWithSelector(CrossDexRouterV3.RouterCancelLimitExceeded.selector, 3, 2));
        ROUTER.cancelOrder(address(PAIR), orderIds);
        vm.stopPrank();
    }

    function test_cancelOrder_emptyArray_noOp() external {
        uint256[] memory orderIds = new uint256[](0);
        vm.prank(USER1);
        ROUTER.cancelOrder(address(PAIR), orderIds);
        // Should not revert
    }

    function test_cancelOrder_revert_invalidPair() external {
        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = 1;

        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(CrossDexRouterV3.RouterInvalidPairAddress.selector, address(0)));
        ROUTER.cancelOrder(address(0), orderIds);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Admin Functions Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_setFindPrevPriceCount_success() external {
        uint256 newCount = 500;

        vm.expectEmit(true, true, false, false);
        emit CrossDexRouterV3.FindPrevPriceCountChanged(FIND_PREV_PRICE_COUNT, newCount);

        vm.prank(OWNER);
        ROUTER.setFindPrevPriceCount(newCount);

        assertEq(ROUTER.findPrevPriceCount(), newCount);
    }

    function test_setFindPrevPriceCount_revert_zero() external {
        vm.prank(OWNER);
        vm.expectRevert(
            abi.encodeWithSelector(CrossDexRouterV3.RouterInvalidInputData.selector, bytes32("findPrevPriceCount"))
        );
        ROUTER.setFindPrevPriceCount(0);
    }

    function test_setFindPrevPriceCount_revert_unauthorized() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(IOwnable.OwnableUnauthorizedAccount.selector, USER1));
        ROUTER.setFindPrevPriceCount(500);
    }

    function test_setMaxMatchCount_success() external {
        uint256 newCount = 50;

        vm.expectEmit(true, true, false, false);
        emit CrossDexRouterV3.MaxMatchCountChanged(MAX_MATCH_COUNT, newCount);

        vm.prank(OWNER);
        ROUTER.setMaxMatchCount(newCount);

        assertEq(ROUTER.maxMatchCount(), newCount);
    }

    function test_setMaxMatchCount_revert_zero() external {
        vm.prank(OWNER);
        vm.expectRevert(
            abi.encodeWithSelector(CrossDexRouterV3.RouterInvalidInputData.selector, bytes32("maxMatchCount"))
        );
        ROUTER.setMaxMatchCount(0);
    }

    function test_setCancelLimit_success() external {
        uint256 newLimit = 10;

        vm.expectEmit(true, true, false, false);
        emit CrossDexRouterV3.CancelLimitChanged(CANCEL_LIMIT, newLimit);

        vm.prank(OWNER);
        ROUTER.setCancelLimit(newLimit);

        assertEq(ROUTER.cancelLimit(), newLimit);
    }

    function test_setCancelLimit_revert_zero() external {
        vm.prank(OWNER);
        vm.expectRevert(
            abi.encodeWithSelector(CrossDexRouterV3.RouterInvalidInputData.selector, bytes32("cancelLimit"))
        );
        ROUTER.setCancelLimit(0);
    }

    function test_setWhitelistedCodeAccount_add() external {
        address[] memory accounts = new address[](2);
        accounts[0] = address(0x1234);
        accounts[1] = address(0x5678);

        vm.expectEmit(true, false, false, true);
        emit CrossDexRouterV3.WhitelistedCodeAccountSet(accounts[0], true);
        vm.expectEmit(true, false, false, true);
        emit CrossDexRouterV3.WhitelistedCodeAccountSet(accounts[1], true);

        vm.prank(OWNER);
        ROUTER.setWhitelistedCodeAccount(accounts, true);
    }

    function test_setWhitelistedCodeAccount_remove() external {
        // First add
        address[] memory accounts = new address[](1);
        accounts[0] = address(0x1234);
        vm.prank(OWNER);
        ROUTER.setWhitelistedCodeAccount(accounts, true);

        // Then remove
        vm.expectEmit(true, false, false, true);
        emit CrossDexRouterV3.WhitelistedCodeAccountSet(accounts[0], false);

        vm.prank(OWNER);
        ROUTER.setWhitelistedCodeAccount(accounts, false);
    }

    function test_setWhitelistedCodeAccount_addDuplicate_noEvent() external {
        address[] memory accounts = new address[](1);
        accounts[0] = address(0x1234);

        // First add
        vm.prank(OWNER);
        ROUTER.setWhitelistedCodeAccount(accounts, true);

        // Add again - should not emit event
        vm.recordLogs();
        vm.prank(OWNER);
        ROUTER.setWhitelistedCodeAccount(accounts, true);

        // Verify: No event should be emitted for duplicate add
        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(logs.length, 0, "No event should be emitted for duplicate add");
    }

    function test_setWhitelistedCodeAccount_revert_unauthorized() external {
        address[] memory accounts = new address[](1);
        accounts[0] = address(0x1234);

        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(IOwnable.OwnableUnauthorizedAccount.selector, USER1));
        ROUTER.setWhitelistedCodeAccount(accounts, true);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // MaxMatchCount Clamping Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_maxMatchCount_clampingBehavior() external {
        // Set maxMatchCount to 5
        vm.prank(OWNER);
        ROUTER.setMaxMatchCount(5);

        // Create many buy orders
        for (uint256 i = 0; i < 10; i++) {
            vm.prank(USER1);
            ROUTER.submitBuyLimit(
                address(PAIR), _toQuote(100), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }

        // Sell with maxMatchCount = 0 (should use global maxMatchCount = 5)
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(10), 0);

        // Should have matched only 5 orders
        uint256[] memory prices = new uint256[](1);
        prices[0] = _toQuote(100);
        uint256[][] memory orderIds = PAIR.ordersByPrices(IPairV3.OrderSide.BUY, prices);
        assertEq(orderIds[0].length, 5);
    }

    function test_maxMatchCount_userSpecifiedLower() external {
        // Set maxMatchCount to 10
        vm.prank(OWNER);
        ROUTER.setMaxMatchCount(10);

        // Create 10 buy orders
        for (uint256 i = 0; i < 10; i++) {
            vm.prank(USER1);
            ROUTER.submitBuyLimit(
                address(PAIR), _toQuote(100), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }

        // Sell with user-specified maxMatchCount = 3
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(10), 3);

        // Should have matched only 3 orders
        uint256[] memory prices = new uint256[](1);
        prices[0] = _toQuote(100);
        uint256[][] memory orderIds = PAIR.ordersByPrices(IPairV3.OrderSide.BUY, prices);
        assertEq(orderIds[0].length, 7); // 10 - 3 = 7 remaining
    }

    function test_maxMatchCount_userSpecifiedHigher_useGlobal() external {
        // Set maxMatchCount to 3
        vm.prank(OWNER);
        ROUTER.setMaxMatchCount(3);

        // Create 10 buy orders
        for (uint256 i = 0; i < 10; i++) {
            vm.prank(USER1);
            ROUTER.submitBuyLimit(
                address(PAIR), _toQuote(100), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }

        // Sell with user-specified maxMatchCount = 100 (higher than global)
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(10), 100);

        // Should have matched only 3 orders (global limit)
        uint256[] memory prices = new uint256[](1);
        prices[0] = _toQuote(100);
        uint256[][] memory orderIds = PAIR.ordersByPrices(IPairV3.OrderSide.BUY, prices);
        assertEq(orderIds[0].length, 7);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // getRequiredBuyVolume Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_getRequiredBuyVolume_revert_invalidPair() external {
        vm.expectRevert(abi.encodeWithSelector(CrossDexRouterV3.RouterInvalidPairAddress.selector, address(0)));
        ROUTER.getRequiredBuyVolume(address(0), _toQuote(1000));
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // H-01 Fix: Forced ETH Injection (selfdestruct) DoS Prevention Tests
    // ═══════════════════════════════════════════════════════════════════════════

    /// @notice Test that forced ETH injection via selfdestruct does not DoS the router
    /// @dev This tests the fix for H-01: Router should still work after forced ETH injection
    function test_forcedEthInjection_doesNotDoSRouter() external {
        // 1. Force send ETH to router via selfdestruct
        uint256 forcedAmount = 1 ether;
        ForceSendEth forceContract = new ForceSendEth{value: forcedAmount}();
        forceContract.boom(payable(address(ROUTER)));

        // Verify router has ETH balance (from forced injection)
        assertEq(address(ROUTER).balance, forcedAmount);

        // 2. Submit order should still work (this would revert before the fix)
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        vm.prank(USER1);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // 3. Verify order was created successfully
        IPairV3.Order memory order = PAIR.orderById(orderId);
        assertEq(order.owner, USER1);
        assertEq(order.price, price);
        assertEq(order.amount, amount);

        // Router balance should still be the forced amount (not touched)
        assertEq(address(ROUTER).balance, forcedAmount);
    }

    /// @notice Test that router balance delta is correctly validated
    /// @dev Ensures that msg.value is properly accounted for and not left in the router
    function test_forcedEthInjection_balanceDeltaValidation() external {
        // Force send some ETH first
        uint256 forcedAmount = 0.5 ether;
        ForceSendEth forceContract = new ForceSendEth{value: forcedAmount}();
        forceContract.boom(payable(address(ROUTER)));

        // Multiple orders should work
        vm.startPrank(USER1);
        for (uint256 i = 0; i < 3; i++) {
            ROUTER.submitBuyLimit(
                address(PAIR), _toQuote(100), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }
        vm.stopPrank();

        vm.startPrank(USER2);
        for (uint256 i = 0; i < 3; i++) {
            ROUTER.submitSellLimit(
                address(PAIR), _toQuote(100), _toBase(1), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }
        vm.stopPrank();

        // Router balance should still be exactly the forced amount
        assertEq(address(ROUTER).balance, forcedAmount);
    }

    /// @notice Test skim function recovers forced ETH
    function test_skim_success() external {
        // Force send ETH to router
        uint256 forcedAmount = 2 ether;
        ForceSendEth forceContract = new ForceSendEth{value: forcedAmount}();
        forceContract.boom(payable(address(ROUTER)));

        assertEq(address(ROUTER).balance, forcedAmount);

        // Owner skims the ETH
        address payable recipient = payable(address(0xBEEF));
        uint256 recipientBalanceBefore = recipient.balance;

        vm.expectEmit(true, false, false, true);
        emit CrossDexRouterV3.Skim(recipient, forcedAmount);

        vm.prank(OWNER);
        ROUTER.skim(recipient);

        // Verify ETH was transferred
        assertEq(address(ROUTER).balance, 0);
        assertEq(recipient.balance, recipientBalanceBefore + forcedAmount);
    }

    /// @notice Test skim with zero balance does nothing
    function test_skim_zeroBalance_noOp() external {
        assertEq(address(ROUTER).balance, 0);

        address payable recipient = payable(address(0xBEEF));

        // Should not revert, just no-op
        vm.recordLogs();
        vm.prank(OWNER);
        ROUTER.skim(recipient);

        // No event should be emitted
        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(logs.length, 0);
    }

    /// @notice Test skim reverts for non-owner
    function test_skim_revert_unauthorized() external {
        // Force send ETH to router
        ForceSendEth forceContract = new ForceSendEth{value: 1 ether}();
        forceContract.boom(payable(address(ROUTER)));

        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(IOwnable.OwnableUnauthorizedAccount.selector, USER1));
        ROUTER.skim(payable(USER1));
    }

    /// @notice Test full scenario: inject, submit orders, then skim
    function test_fullScenario_injectSubmitSkim() external {
        // 1. Attacker injects ETH
        uint256 forcedAmount = 1 ether;
        ForceSendEth forceContract = new ForceSendEth{value: forcedAmount}();
        forceContract.boom(payable(address(ROUTER)));

        // 2. Users can still trade normally
        vm.prank(USER1);
        uint256 buyOrderId = ROUTER.submitBuyLimit(
            address(PAIR), _toQuote(100), _toBase(10), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(5), 0);

        // Verify partial fill occurred
        IPairV3.Order memory order = PAIR.orderById(buyOrderId);
        assertEq(order.amount, _toBase(5)); // 10 - 5 = 5 remaining

        // 3. Admin recovers the forced ETH
        address payable treasury = payable(makeAddr("treasury"));
        vm.prank(OWNER);
        ROUTER.skim(treasury);

        assertEq(address(ROUTER).balance, 0);
        assertEq(treasury.balance, forcedAmount);

        // 4. Trading continues normally after skim
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), _toBase(5), 0);

        order = PAIR.orderById(buyOrderId);
        assertEq(order.amount, 0); // Fully filled
    }
}
