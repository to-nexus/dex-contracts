// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {Vm} from "forge-std/Vm.sol";

import {CrossDexImplV3} from "../src/CrossDexImplV3.sol";
import {FeeControllerV2Compat} from "../src/FeeControllerV2Compat.sol";
import {MarketImplV3} from "../src/MarketImplV3.sol";
import {PairImplV3} from "../src/PairImplV3.sol";
import {IOwnable} from "../src/interfaces/IOwnable.sol";
import {DEXV3BaseTest} from "./DEXV3Base.t.sol";
import {T20} from "./mock/T20.sol";

contract CrossDexImplV3Test is DEXV3BaseTest {
    function setUp() external {
        _deployV3(18, 18, 1e2, 1e6);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // View Functions Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_ROUTER() external view {
        assertEq(CROSS_DEX.ROUTER(), address(ROUTER));
    }

    function test_marketImpl() external view {
        assertTrue(CROSS_DEX.marketImpl() != address(0));
    }

    function test_pairImpl() external view {
        assertTrue(CROSS_DEX.pairImpl() != address(0));
    }

    function test_isMarket_valid() external view {
        assertTrue(CROSS_DEX.isMarket(address(MARKET)));
    }

    function test_isMarket_invalid() external view {
        assertFalse(CROSS_DEX.isMarket(address(0)));
        assertFalse(CROSS_DEX.isMarket(USER1));
    }

    function test_allMarkets() external view {
        (address[] memory markets, address[] memory quotes) = CROSS_DEX.allMarkets();
        assertEq(markets.length, 1);
        assertEq(quotes.length, 1);
        assertEq(markets[0], address(MARKET));
        assertEq(quotes[0], address(QUOTE));
    }

    function test_pairToMarket() external view {
        assertEq(CROSS_DEX.pairToMarket(address(PAIR)), address(MARKET));
    }

    function test_tickSizeSetter_initial() external view {
        assertEq(CROSS_DEX.tickSizeSetter(), address(0));
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Market Creation Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_createMarket_success() external {
        T20 newQuote = new T20("QUOTE2", "Q2", 18);

        vm.prank(OWNER);
        address market = CROSS_DEX.createMarket(OWNER, address(newQuote), address(FEE_CONTROLLER), "New Market");

        assertTrue(CROSS_DEX.isMarket(market));

        (address[] memory markets, address[] memory quotes) = CROSS_DEX.allMarkets();
        assertEq(markets.length, 2);
        assertEq(markets[1], market);
        assertEq(quotes[1], address(newQuote));
    }

    function test_createMarket_deterministic_create2() external {
        T20 newQuote = new T20("QUOTE2", "Q2", 18);
        string memory message = "TestMarket";

        // First deployment
        vm.prank(OWNER);
        address market1 = CROSS_DEX.createMarket(OWNER, address(newQuote), address(FEE_CONTROLLER), message);

        // Same salt should cause revert (address already used)
        vm.prank(OWNER);
        vm.expectRevert();
        CROSS_DEX.createMarket(OWNER, address(newQuote), address(FEE_CONTROLLER), message);

        // Different message should work
        vm.prank(OWNER);
        address market2 = CROSS_DEX.createMarket(OWNER, address(newQuote), address(FEE_CONTROLLER), "DifferentMessage");

        assertTrue(market1 != market2);
    }

    function test_createMarket_revert_unauthorized() external {
        T20 newQuote = new T20("QUOTE2", "Q2", 18);

        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(IOwnable.OwnableUnauthorizedAccount.selector, USER1));
        CROSS_DEX.createMarket(USER1, address(newQuote), address(FEE_CONTROLLER), "");
    }

    function test_createMarket_emitsEvent() external {
        T20 newQuote = new T20("QUOTE2", "Q2", 18);

        vm.expectEmit(true, false, true, false);
        emit CrossDexImplV3.MarketCreated(address(newQuote), address(0), OWNER, address(FEE_CONTROLLER), "test");

        vm.prank(OWNER);
        CROSS_DEX.createMarket(OWNER, address(newQuote), address(FEE_CONTROLLER), "test");
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Admin Functions Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_setTickSizeSetter_success() external {
        address newSetter = address(0x1234);

        vm.expectEmit(true, true, false, false);
        emit CrossDexImplV3.TickSizeSetterSet(address(0), newSetter);

        vm.prank(OWNER);
        CROSS_DEX.setTickSizeSetter(newSetter);

        assertEq(CROSS_DEX.tickSizeSetter(), newSetter);
    }

    function test_setTickSizeSetter_revert_zeroAddress() external {
        vm.prank(OWNER);
        vm.expectRevert(
            abi.encodeWithSelector(CrossDexImplV3.CrossDexInvalidTickSizeSetter.selector, address(0), address(0))
        );
        CROSS_DEX.setTickSizeSetter(address(0));
    }

    function test_setTickSizeSetter_revert_sameValue() external {
        address setter = address(0x1234);

        vm.prank(OWNER);
        CROSS_DEX.setTickSizeSetter(setter);

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(CrossDexImplV3.CrossDexInvalidTickSizeSetter.selector, setter, setter));
        CROSS_DEX.setTickSizeSetter(setter);
    }

    function test_setTickSizeSetter_revert_unauthorized() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(IOwnable.OwnableUnauthorizedAccount.selector, USER1));
        CROSS_DEX.setTickSizeSetter(address(0x1234));
    }

    function test_setPairImpl_success() external {
        address newPairImpl = address(new PairImplV3());
        address oldPairImpl = CROSS_DEX.pairImpl();

        vm.expectEmit(true, true, false, false);
        emit CrossDexImplV3.PairImplSet(oldPairImpl, newPairImpl);

        vm.prank(OWNER);
        CROSS_DEX.setPairImpl(newPairImpl);

        assertEq(CROSS_DEX.pairImpl(), newPairImpl);
    }

    function test_setPairImpl_revert_zeroAddress() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(CrossDexImplV3.CrossDexInitializeData.selector, bytes32("pairImpl")));
        CROSS_DEX.setPairImpl(address(0));
    }

    function test_setMarketImpl_success() external {
        address newMarketImpl = address(new MarketImplV3());
        address oldMarketImpl = CROSS_DEX.marketImpl();

        vm.expectEmit(true, true, false, false);
        emit CrossDexImplV3.MarketImplSet(oldMarketImpl, newMarketImpl);

        vm.prank(OWNER);
        CROSS_DEX.setMarketImpl(newMarketImpl);

        assertEq(CROSS_DEX.marketImpl(), newMarketImpl);
    }

    function test_setMarketImpl_revert_zeroAddress() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(CrossDexImplV3.CrossDexInitializeData.selector, bytes32("marketImpl")));
        CROSS_DEX.setMarketImpl(address(0));
    }

    function test_setFeeControllerAllow_add() external {
        address newFeeController = address(new FeeControllerV2Compat());

        vm.expectEmit(true, true, false, false);
        emit CrossDexImplV3.FeeControllerAllowed(newFeeController, true);

        vm.prank(OWNER);
        CROSS_DEX.setFeeControllerAllow(newFeeController, true);
    }

    function test_setFeeControllerAllow_remove() external {
        // FEE_CONTROLLER is already allowed
        vm.expectEmit(true, true, false, false);
        emit CrossDexImplV3.FeeControllerAllowed(address(FEE_CONTROLLER), false);

        vm.prank(OWNER);
        CROSS_DEX.setFeeControllerAllow(address(FEE_CONTROLLER), false);
    }

    function test_setFeeControllerAllow_addDuplicate_noEvent() external {
        // FEE_CONTROLLER is already allowed
        vm.recordLogs();
        vm.prank(OWNER);
        CROSS_DEX.setFeeControllerAllow(address(FEE_CONTROLLER), true);

        // Verify: No FeeControllerAllowed event should be emitted for duplicate add
        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(logs.length, 0, "No event should be emitted for duplicate add");
    }

    function test_setFeeControllerAllow_removeTwice_noEvent() external {
        // First remove
        vm.prank(OWNER);
        CROSS_DEX.setFeeControllerAllow(address(FEE_CONTROLLER), false);

        // Second remove - no event
        vm.recordLogs();
        vm.prank(OWNER);
        CROSS_DEX.setFeeControllerAllow(address(FEE_CONTROLLER), false);

        // Verify: No event should be emitted for duplicate remove
        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(logs.length, 0, "No event should be emitted for duplicate remove");
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Access Control Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_checkTickSizeRoles_revert_notSet() external {
        // tickSizeSetter is address(0), so should always revert
        vm.expectRevert(abi.encodeWithSelector(CrossDexImplV3.CrossDexUnauthorizedChangeTickSizes.selector, USER1));
        CROSS_DEX.checkTickSizeRoles(USER1);
    }

    function test_checkTickSizeRoles_revert_wrongAccount() external {
        address setter = address(0x1234);
        vm.prank(OWNER);
        CROSS_DEX.setTickSizeSetter(setter);

        vm.expectRevert(abi.encodeWithSelector(CrossDexImplV3.CrossDexUnauthorizedChangeTickSizes.selector, USER1));
        CROSS_DEX.checkTickSizeRoles(USER1);
    }

    function test_checkTickSizeRoles_success() external {
        address setter = address(0x1234);
        vm.prank(OWNER);
        CROSS_DEX.setTickSizeSetter(setter);

        // Should not revert
        CROSS_DEX.checkTickSizeRoles(setter);
    }

    function test_checkFeeControllerAllowed_revert_notAllowed() external {
        address unallowed = address(0x9999);
        vm.expectRevert(abi.encodeWithSelector(CrossDexImplV3.CrossDexInvalidFeeController.selector, unallowed));
        CROSS_DEX.checkFeeControllerAllowed(unallowed);
    }

    function test_checkFeeControllerAllowed_success() external view {
        // FEE_CONTROLLER is already allowed
        CROSS_DEX.checkFeeControllerAllowed(address(FEE_CONTROLLER));
    }

    function test_pairCreated_revert_notMarket() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(CrossDexImplV3.CrossDexInvalidMarketAddress.selector, USER1));
        CROSS_DEX.pairCreated(address(0x1234));
    }

    function test_pairCreated_success() external {
        // Create a fake pair address and call from MARKET
        address fakePair = address(0x1234);

        vm.prank(address(MARKET));
        CROSS_DEX.pairCreated(fakePair);

        assertEq(CROSS_DEX.pairToMarket(fakePair), address(MARKET));
    }
}
