// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {CrossDexImplV3} from "../src/CrossDexImplV3.sol";
import {FeeControllerV2Compat} from "../src/FeeControllerV2Compat.sol";
import {MarketImplV3} from "../src/MarketImplV3.sol";
import {PairImplV3} from "../src/PairImplV3.sol";
import {IOwnable} from "../src/interfaces/IOwnable.sol";
import {DEXV3BaseTest} from "./DEXV3Base.t.sol";
import {T20} from "./mock/T20.sol";

contract MarketImplV3Test is DEXV3BaseTest {
    function setUp() external {
        _deployV3(18, 18, 1e2, 1e6);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // View Functions Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_deployed() external view {
        assertGt(MARKET.deployed(), 0);
    }

    function test_CROSS_DEX() external view {
        assertEq(address(MARKET.CROSS_DEX()), address(CROSS_DEX));
    }

    function test_QUOTE() external view {
        assertEq(MARKET.QUOTE(), address(QUOTE));
    }

    function test_ROUTER() external view {
        assertEq(MARKET.ROUTER(), address(ROUTER));
    }

    function test_pairImpl() external view {
        assertTrue(MARKET.pairImpl() != address(0));
    }

    function test_feeController() external view {
        assertEq(MARKET.feeController(), address(FEE_CONTROLLER));
    }

    function test_allPairs() external view {
        (address[] memory bases, address[] memory pairs) = MARKET.allPairs();
        assertEq(bases.length, 1);
        assertEq(pairs.length, 1);
        assertEq(bases[0], address(BASE));
        assertEq(pairs[0], address(PAIR));
    }

    function test_baseToPair() external view {
        assertEq(MARKET.baseToPair(address(BASE)), address(PAIR));
    }

    function test_baseToPair_nonExistent() external {
        vm.expectRevert();
        MARKET.baseToPair(address(0x1234));
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Pair Creation Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_createPair_success() external {
        T20 newBase = new T20("BASE2", "B2", 18);

        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        vm.prank(OWNER);
        address pair = MARKET.createPair(address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);

        assertTrue(pair != address(0));
        assertEq(MARKET.baseToPair(address(newBase)), pair);

        (address[] memory bases, address[] memory pairs) = MARKET.allPairs();
        assertEq(bases.length, 2);
        assertEq(pairs.length, 2);
    }

    function test_createPair_deterministic_create2() external {
        T20 newBase = new T20("BASE2", "B2", 18);

        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        vm.prank(OWNER);
        address pair1 = MARKET.createPair(address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);

        // Same base should cause revert
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(MarketImplV3.MarketAlreadyCreatedBaseAddress.selector, address(newBase)));
        MARKET.createPair(address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);

        // Different base should work
        T20 newBase2 = new T20("BASE3", "B3", 18);
        vm.prank(OWNER);
        address pair2 = MARKET.createPair(address(newBase2), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);

        assertTrue(pair1 != pair2);
    }

    function test_createPair_revert_zeroBase() external {
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(MarketImplV3.MarketInvalidBaseAddress.selector, address(0)));
        MARKET.createPair(address(0), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
    }

    function test_createPair_revert_baseEqualsQuote() external {
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(MarketImplV3.MarketInvalidBaseAddress.selector, address(QUOTE)));
        MARKET.createPair(address(QUOTE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
    }

    function test_createPair_revert_duplicateBase() external {
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        // BASE is already used in setUp
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(MarketImplV3.MarketAlreadyCreatedBaseAddress.selector, address(BASE)));
        MARKET.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
    }

    function test_createPair_revert_unauthorized() external {
        T20 newBase = new T20("BASE2", "B2", 18);
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(IOwnable.OwnableUnauthorizedAccount.selector, USER1));
        MARKET.createPair(address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
    }

    function test_createPair_emitsEvent() external {
        T20 newBase = new T20("BASE2", "B2", 18);
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        vm.expectEmit(false, true, false, true);
        emit MarketImplV3.PairCreated(address(0), address(newBase), block.timestamp);

        vm.prank(OWNER);
        MARKET.createPair(address(newBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Admin Functions Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_setPairImpl_success() external {
        address newPairImpl = address(new PairImplV3());
        address oldPairImpl = MARKET.pairImpl();

        vm.expectEmit(true, true, false, false);
        emit MarketImplV3.PairImplSet(oldPairImpl, newPairImpl);

        vm.prank(OWNER);
        MARKET.setPairImpl(newPairImpl);

        assertEq(MARKET.pairImpl(), newPairImpl);
    }

    function test_setPairImpl_revert_zeroAddress() external {
        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(MarketImplV3.MarketInvalidInitializeData.selector, bytes32("pairImpl")));
        MARKET.setPairImpl(address(0));
    }

    function test_setPairImpl_revert_unauthorized() external {
        vm.prank(USER1);
        vm.expectRevert(abi.encodeWithSelector(IOwnable.OwnableUnauthorizedAccount.selector, USER1));
        MARKET.setPairImpl(address(0x1234));
    }

    function test_setFeeController_updateFeeController() external {
        // Create a new fee controller
        FeeControllerV2Compat newFeeController = new FeeControllerV2Compat();

        // Allow the new fee controller
        vm.prank(OWNER);
        CROSS_DEX.setFeeControllerAllow(address(newFeeController), true);

        bytes memory feeData = abi.encode(FEE_COLLECTOR, uint32(10), uint32(20), uint32(5), uint32(10));

        address[] memory pairs = new address[](1);
        pairs[0] = address(PAIR);

        vm.expectEmit(true, true, false, false);
        emit MarketImplV3.FeeControllerUpdated(address(FEE_CONTROLLER), address(newFeeController));

        vm.prank(OWNER);
        MARKET.setFeeController(pairs, address(newFeeController), feeData);

        assertEq(MARKET.feeController(), address(newFeeController));
    }

    function test_setFeeController_updatePairs() external {
        bytes memory feeData = abi.encode(FEE_COLLECTOR, uint32(10), uint32(20), uint32(5), uint32(10));

        address[] memory pairs = new address[](1);
        pairs[0] = address(PAIR);

        // Update the pair
        vm.prank(OWNER);
        MARKET.setFeeController(pairs, address(FEE_CONTROLLER), feeData);

        // Verify PAIR still has fee controller
        assertEq(address(PAIR.feeController()), address(FEE_CONTROLLER));
    }

    function test_setFeeController_selectivePairUpdate() external {
        // Create multiple pairs
        T20 base2 = new T20("BASE2", "B2", 18);
        T20 base3 = new T20("BASE3", "B3", 18);

        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        vm.startPrank(OWNER);
        address pair2 = MARKET.createPair(address(base2), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
        address pair3 = MARKET.createPair(address(base3), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
        vm.stopPrank();

        // Update only specific pairs (pair1 and pair2)
        address[] memory pairsToUpdate = new address[](2);
        pairsToUpdate[0] = address(PAIR);
        pairsToUpdate[1] = pair2;

        bytes memory newFeeData = abi.encode(FEE_COLLECTOR, uint32(10), uint32(20), uint32(5), uint32(10));
        vm.prank(OWNER);
        MARKET.setFeeController(pairsToUpdate, address(FEE_CONTROLLER), newFeeData);

        // All 3 pairs should exist
        (, address[] memory allPairs) = MARKET.allPairs();
        assertEq(allPairs.length, 3);

        // Verify pair3 was not included in the update (but still has same controller)
        assertEq(address(PairImplV3(pair3).feeController()), address(FEE_CONTROLLER));
    }

    function test_setFeeController_emptyPairsArray() external {
        // Create a new fee controller
        FeeControllerV2Compat newFeeController = new FeeControllerV2Compat();

        vm.prank(OWNER);
        CROSS_DEX.setFeeControllerAllow(address(newFeeController), true);

        bytes memory feeData = abi.encode(FEE_COLLECTOR, uint32(10), uint32(20), uint32(5), uint32(10));

        address[] memory emptyPairs = new address[](0);

        vm.expectEmit(true, true, false, false);
        emit MarketImplV3.FeeControllerUpdated(address(FEE_CONTROLLER), address(newFeeController));

        // Should update market feeController but not iterate any pairs
        vm.prank(OWNER);
        MARKET.setFeeController(emptyPairs, address(newFeeController), feeData);

        // Market feeController updated
        assertEq(MARKET.feeController(), address(newFeeController));
        // But PAIR still has old feeController
        assertEq(address(PAIR.feeController()), address(FEE_CONTROLLER));
    }

    function test_setFeeController_revert_notAllowed() external {
        address unallowedFeeController = address(0x9999);
        bytes memory feeData = abi.encode(FEE_COLLECTOR, uint32(10), uint32(20), uint32(5), uint32(10));

        address[] memory pairs = new address[](1);
        pairs[0] = address(PAIR);

        vm.prank(OWNER);
        vm.expectRevert(
            abi.encodeWithSelector(CrossDexImplV3.CrossDexInvalidFeeController.selector, unallowedFeeController)
        );
        MARKET.setFeeController(pairs, unallowedFeeController, feeData);
    }

    function test_setFeeController_revert_invalidPair() external {
        bytes memory feeData = abi.encode(FEE_COLLECTOR, uint32(10), uint32(20), uint32(5), uint32(10));

        address fakePair = address(0x1234);
        address[] memory pairs = new address[](1);
        pairs[0] = fakePair;

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(MarketImplV3.MarketInvalidPairAddress.selector, fakePair));
        MARKET.setFeeController(pairs, address(FEE_CONTROLLER), feeData);
    }

    function test_setFeeController_revert_pairFromOtherMarket() external {
        // Create another market
        T20 otherQuote = new T20("OTHERQUOTE", "OQ", 18);
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        vm.prank(OWNER);
        address otherMarket = CROSS_DEX.createMarket(OWNER, address(otherQuote), address(FEE_CONTROLLER), "other");

        // Create a pair in the other market
        T20 otherBase = new T20("OTHERBASE", "OB", 18);
        vm.prank(OWNER);
        address otherPair = MarketImplV3(otherMarket)
            .createPair(address(otherBase), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);

        // Try to update the other market's pair from this market
        address[] memory pairs = new address[](1);
        pairs[0] = otherPair;

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSelector(MarketImplV3.MarketInvalidPairAddress.selector, otherPair));
        MARKET.setFeeController(pairs, address(FEE_CONTROLLER), feeData);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Access Control Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_checkTickSizeRoles_delegatesToCrossDex() external {
        // Without tickSizeSetter, should revert
        vm.expectRevert(abi.encodeWithSelector(CrossDexImplV3.CrossDexUnauthorizedChangeTickSizes.selector, USER1));
        MARKET.checkTickSizeRoles(USER1);
    }

    function test_checkFeeControllerAllowed_delegatesToCrossDex() external {
        address unallowed = address(0x9999);
        vm.expectRevert(abi.encodeWithSelector(CrossDexImplV3.CrossDexInvalidFeeController.selector, unallowed));
        MARKET.checkFeeControllerAllowed(unallowed);
    }

    function test_checkFeeControllerAllowed_success() external view {
        // Should not revert for allowed fee controller
        MARKET.checkFeeControllerAllowed(address(FEE_CONTROLLER));
    }
}
