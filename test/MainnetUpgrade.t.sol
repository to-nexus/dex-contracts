// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {ERC1967Utils} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Utils.sol";
import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {CrossDexImplV3} from "../src/CrossDexImplV3.sol";
import {CrossDexRouterV3} from "../src/CrossDexRouterV3.sol";
import {FeeControllerV2Compat} from "../src/FeeControllerV2Compat.sol";
import {FeeControllerV3Dist} from "../src/FeeControllerV3Dist.sol";
import {FeeControllerV3Split} from "../src/FeeControllerV3Split.sol";
import {MarketImplV3} from "../src/MarketImplV3.sol";
import {PairImplV3} from "../src/PairImplV3.sol";
import {WETH} from "../src/WETH.sol";

import "../src/interfaces/IFeeController.sol";
import {IPairV3} from "../src/interfaces/IPairV3.sol";

import {T20} from "./mock/T20.sol";

/// @dev Minimal mock of Verse8MarketOwner — only `execute` with owner check
contract MockVerse8MarketOwner {
    address public owner;

    constructor(address _owner) {
        owner = _owner;
    }

    function execute(address to, uint256 value, bytes calldata data) external returns (bytes memory) {
        require(msg.sender == owner, "not owner");
        (bool success, bytes memory result) = to.call{value: value}(data);
        require(success, "execute failed");
        return result;
    }
}

/// @title MainnetUpgradeTest
/// @notice Simulates the full post-audit upgrade scenario described in .cursor/upgrade-commands.md
/// @dev Phases:
///   1-1. Deploy new implementations
///   1-2. Upgrade CrossDex + Router, swap fee controllers
///   2.   Upgrade wCROSS/CROSSD markets (Owner A)
///   3-0. Transfer FORGE market ownership from Verse8MarketOwner → EOA
///   3-1. Upgrade FORGE market (Verse8 Admin)
///   4-1. Migrate fee controllers (same type, preserving existing config)
///   4-1-4. Transfer FORGE market ownership back to Verse8MarketOwner
contract MainnetUpgradeTest is Test {
    // ─────────────────────────────────────────────────────────────────────────
    // Actors
    // ─────────────────────────────────────────────────────────────────────────

    address constant CROSSDEX_OWNER = address(bytes20("CROSSDEX_OWNER"));
    address constant MARKET_OWNER_A = address(bytes20("MARKET_OWNER_A"));
    address constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    address constant CREATOR = address(bytes20("CREATOR"));
    address constant USER1 = address(bytes20("USER1"));
    address constant USER2 = address(bytes20("USER2"));

    // Verse8 Admin = MARKET_OWNER_A (same key as mainnet)
    address constant VERSE8_ADMIN = MARKET_OWNER_A;

    // ─────────────────────────────────────────────────────────────────────────
    // Contracts — "current mainnet" state
    // ─────────────────────────────────────────────────────────────────────────

    CrossDexImplV3 crossDex;
    CrossDexRouterV3 router;

    MarketImplV3 marketWCross;
    MarketImplV3 marketCrossD;
    MarketImplV3 marketForge;

    MockVerse8MarketOwner forgeMarketOwner;

    PairImplV3[] pairsWCross;
    PairImplV3[] pairsCrossD;
    PairImplV3[] pairsForge;

    FeeControllerV2Compat oldFcV2Compat;
    FeeControllerV3Split oldFcV3Split;

    IERC20 quoteWCross;
    IERC20 quoteCrossD;
    IERC20[] basesWCross;
    IERC20[] basesCrossD;
    IERC20[] basesForge;

    uint256 constant QUOTE_DECIMALS = 1e18;
    uint256 constant BASE_DECIMALS = 1e18;
    uint256 constant TICK_SIZE = QUOTE_DECIMALS;
    uint256 constant LOT_SIZE = BASE_DECIMALS;

    uint32 constant V2_SELLER_MAKER = 500;
    uint32 constant V2_SELLER_TAKER = 500;

    uint32 constant V3SPLIT_TAKER = 100;
    uint32 constant V3SPLIT_CREATOR_SHARE = 3000;
    uint32 constant V3SPLIT_MAKER_REBATE = 2000;

    uint256[2] internal _searchPrices;

    // ─────────────────────────────────────────────────────────────────────────
    // New implementations (deployed in Phase 1-1)
    // ─────────────────────────────────────────────────────────────────────────

    address newCrossDexImpl;
    address newRouterImpl;
    address newMarketImpl;
    address newPairImpl;
    FeeControllerV2Compat newFcV2Compat;
    FeeControllerV3Dist newFcV3Dist;
    FeeControllerV3Split newFcV3Split;

    // ─────────────────────────────────────────────────────────────────────────
    // Setup — simulate current mainnet deployment
    // ─────────────────────────────────────────────────────────────────────────

    function setUp() public {
        vm.label(CROSSDEX_OWNER, "crossDexOwner");
        vm.label(MARKET_OWNER_A, "marketOwnerA / verse8Admin");
        vm.label(FEE_COLLECTOR, "feeCollector");
        vm.label(CREATOR, "creator");
        vm.label(USER1, "user1");
        vm.label(USER2, "user2");

        _deployCurrentMainnetState();
        _seedUsersAndApprovals();
    }

    function _deployCurrentMainnetState() internal {
        vm.startPrank(CROSSDEX_OWNER);

        oldFcV2Compat = new FeeControllerV2Compat();
        oldFcV3Split = new FeeControllerV3Split();

        address routerImpl = address(new CrossDexRouterV3());
        address marketImpl = address(new MarketImplV3());
        address pairImpl = address(new PairImplV3());
        address crossDexImpl = address(new CrossDexImplV3());

        ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
        crossDex = CrossDexImplV3(address(proxy));
        crossDex.initialize(
            CROSSDEX_OWNER,
            routerImpl,
            type(uint256).max,
            type(uint256).max,
            type(uint256).max,
            marketImpl,
            pairImpl,
            address(0)
        );

        router = CrossDexRouterV3(crossDex.ROUTER());
        crossDex.setFeeControllerAllow(address(oldFcV2Compat), true);
        crossDex.setFeeControllerAllow(address(oldFcV3Split), true);

        // ── wCROSS Market (Owner A, V2Compat, 3 pairs) ──
        quoteWCross = new T20("wCROSS", "CROSS", 18);
        address mWCross = crossDex.createMarket(MARKET_OWNER_A, address(quoteWCross), address(oldFcV2Compat), "wCROSS");
        marketWCross = MarketImplV3(mWCross);

        vm.stopPrank();
        vm.startPrank(MARKET_OWNER_A);

        bytes memory v2InitData = abi.encode(FEE_COLLECTOR, V2_SELLER_MAKER, V2_SELLER_TAKER, uint32(0), uint32(0));
        for (uint256 i = 0; i < 3; ++i) {
            T20 base = new T20(string(abi.encodePacked("BASE_W", vm.toString(i))), "BW", 18);
            basesWCross.push(base);
            address pair = marketWCross.createPair(address(base), TICK_SIZE, LOT_SIZE, v2InitData);
            pairsWCross.push(PairImplV3(pair));
        }

        // ── CROSSD Market (Owner A, V2Compat, 3 pairs) ──
        vm.stopPrank();
        vm.startPrank(CROSSDEX_OWNER);
        quoteCrossD = new T20("CROSSD", "CROSSD", 18);
        address mCrossD = crossDex.createMarket(MARKET_OWNER_A, address(quoteCrossD), address(oldFcV2Compat), "CROSSD");
        marketCrossD = MarketImplV3(mCrossD);
        vm.stopPrank();
        vm.startPrank(MARKET_OWNER_A);

        for (uint256 i = 0; i < 3; ++i) {
            T20 base = new T20(string(abi.encodePacked("BASE_D", vm.toString(i))), "BD", 18);
            basesCrossD.push(base);
            address pair = marketCrossD.createPair(address(base), TICK_SIZE, LOT_SIZE, v2InitData);
            pairsCrossD.push(PairImplV3(pair));
        }
        vm.stopPrank();

        // ── FORGE Market (owner = Verse8MarketOwner contract, V3Split, 3 pairs) ──
        forgeMarketOwner = new MockVerse8MarketOwner(VERSE8_ADMIN);
        vm.label(address(forgeMarketOwner), "forgeMarketOwner");

        vm.startPrank(CROSSDEX_OWNER);
        address mForge =
            crossDex.createMarket(address(forgeMarketOwner), address(quoteCrossD), address(oldFcV3Split), "FORGE");
        marketForge = MarketImplV3(mForge);
        vm.stopPrank();

        // Verse8MarketOwner.execute → Market.createPair (simulates real flow)
        vm.startPrank(VERSE8_ADMIN);
        for (uint256 i = 0; i < 3; ++i) {
            T20 base = new T20(string(abi.encodePacked("BASE_F", vm.toString(i))), "BF", 18);
            basesForge.push(base);
            address pairCreator = address(uint160(0xC0DE0000 + i));
            bytes memory splitInitData =
                abi.encode(FEE_COLLECTOR, pairCreator, V3SPLIT_TAKER, V3SPLIT_CREATOR_SHARE, V3SPLIT_MAKER_REBATE);

            bytes memory callData =
                abi.encodeCall(MarketImplV3.createPair, (address(base), TICK_SIZE, LOT_SIZE, splitInitData));
            bytes memory result = forgeMarketOwner.execute(address(marketForge), 0, callData);
            address pair = abi.decode(result, (address));
            pairsForge.push(PairImplV3(pair));
        }
        vm.stopPrank();
    }

    function _seedUsersAndApprovals() internal {
        _fundAndApprove(USER1, quoteWCross, basesWCross);
        _fundAndApprove(USER2, quoteWCross, basesWCross);
        _fundAndApprove(USER1, quoteCrossD, basesCrossD);
        _fundAndApprove(USER2, quoteCrossD, basesCrossD);
        _fundAndApprove(USER1, quoteCrossD, basesForge);
        _fundAndApprove(USER2, quoteCrossD, basesForge);
    }

    function _fundAndApprove(address user, IERC20 quote, IERC20[] storage bases) internal {
        vm.startPrank(CROSSDEX_OWNER);
        quote.transfer(user, 100_000 * QUOTE_DECIMALS);
        vm.stopPrank();

        for (uint256 i = 0; i < bases.length; ++i) {
            vm.prank(MARKET_OWNER_A);
            try bases[i].transfer(user, 100_000 * BASE_DECIMALS) {} catch {}
            vm.prank(VERSE8_ADMIN);
            try bases[i].transfer(user, 100_000 * BASE_DECIMALS) {} catch {}
        }

        vm.startPrank(user);
        quote.approve(address(router), type(uint256).max);
        for (uint256 i = 0; i < bases.length; ++i) {
            bases[i].approve(address(router), type(uint256).max);
        }
        vm.stopPrank();
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Main test: full upgrade scenario
    // ─────────────────────────────────────────────────────────────────────────

    function test_full_upgrade_scenario() public {
        // ── Pre-upgrade: place orders to verify state preservation ──
        _placeOrdersOnAllPairs();

        // Snapshot pre-upgrade state
        _SnapshotAll memory before = _snapshotAll();

        // ── Phase 1-1: Deploy new implementations ──
        _phase1_1_deployImplementations();

        // ── Phase 1-2: Upgrade CrossDex + Router ──
        _phase1_2_upgradeCrossDexAndRouter();

        // ── Phase 2: Upgrade wCROSS + CROSSD markets (Owner A) ──
        _phase2_upgradeMarketsOwnerA();

        // ── Phase 3: Upgrade FORGE market (Owner B) ──
        _phase3_upgradeMarketOwnerB();

        // ── Phase 4-1: Migrate fee controllers (same type) ──
        _phase4_1_migrateFeeControllers();

        // ── Post-upgrade verification ──
        _verifyPostUpgrade(before);
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Phase implementations
    // ─────────────────────────────────────────────────────────────────────────

    function _phase1_1_deployImplementations() internal {
        vm.startPrank(CROSSDEX_OWNER);
        newCrossDexImpl = address(new CrossDexImplV3());
        newRouterImpl = address(new CrossDexRouterV3());
        newMarketImpl = address(new MarketImplV3());
        newPairImpl = address(new PairImplV3());
        newFcV2Compat = new FeeControllerV2Compat();
        newFcV3Dist = new FeeControllerV3Dist();
        newFcV3Split = new FeeControllerV3Split();
        vm.stopPrank();

        assertTrue(newCrossDexImpl != address(0), "crossDexImpl deployed");
        assertTrue(newPairImpl != address(0), "pairImpl deployed");
        assertTrue(address(newFcV3Dist) != address(0), "fcV3Dist deployed");
    }

    function _phase1_2_upgradeCrossDexAndRouter() internal {
        address oldMarketImpl = crossDex.marketImpl();
        address oldPairImpl = crossDex.pairImpl();

        vm.startPrank(CROSSDEX_OWNER);

        // Upgrade CrossDex proxy
        crossDex.upgradeToAndCall(newCrossDexImpl, hex"");
        assertEq(crossDex.version(), 3);

        // Upgrade Router proxy
        router.upgradeToAndCall(newRouterImpl, hex"");
        assertEq(router.version(), 3);

        // Update impl references
        crossDex.setMarketImpl(newMarketImpl);
        crossDex.setPairImpl(newPairImpl);
        assertEq(crossDex.marketImpl(), newMarketImpl);
        assertEq(crossDex.pairImpl(), newPairImpl);
        assertTrue(crossDex.marketImpl() != oldMarketImpl, "marketImpl changed");
        assertTrue(crossDex.pairImpl() != oldPairImpl, "pairImpl changed");

        // Remove old fee controllers
        crossDex.setFeeControllerAllow(address(oldFcV2Compat), false);
        crossDex.setFeeControllerAllow(address(oldFcV3Split), false);

        // Allow new fee controllers
        crossDex.setFeeControllerAllow(address(newFcV2Compat), true);
        crossDex.setFeeControllerAllow(address(newFcV3Dist), true);
        crossDex.setFeeControllerAllow(address(newFcV3Split), true);

        vm.stopPrank();

        // Old FCs should be disallowed
        vm.expectRevert();
        crossDex.checkFeeControllerAllowed(address(oldFcV2Compat));

        vm.expectRevert();
        crossDex.checkFeeControllerAllowed(address(oldFcV3Split));

        // New FCs should be allowed
        crossDex.checkFeeControllerAllowed(address(newFcV2Compat));
        crossDex.checkFeeControllerAllowed(address(newFcV3Dist));
        crossDex.checkFeeControllerAllowed(address(newFcV3Split));
    }

    function _phase2_upgradeMarketsOwnerA() internal {
        vm.startPrank(MARKET_OWNER_A);

        // ── wCROSS Market ──
        marketWCross.upgradeToAndCall(newMarketImpl, hex"");
        marketWCross.setPairImpl(newPairImpl);
        assertEq(marketWCross.version(), 3);
        assertEq(marketWCross.pairImpl(), newPairImpl);

        for (uint256 i = 0; i < pairsWCross.length; ++i) {
            pairsWCross[i].upgradeToAndCall(newPairImpl, hex"");
            assertEq(pairsWCross[i].version(), 3);
        }

        // ── CROSSD Market ──
        marketCrossD.upgradeToAndCall(newMarketImpl, hex"");
        marketCrossD.setPairImpl(newPairImpl);
        assertEq(marketCrossD.version(), 3);

        for (uint256 i = 0; i < pairsCrossD.length; ++i) {
            pairsCrossD[i].upgradeToAndCall(newPairImpl, hex"");
            assertEq(pairsCrossD[i].version(), 3);
        }

        vm.stopPrank();
    }

    function _phase3_upgradeMarketOwnerB() internal {
        // Phase 3-0: Transfer FORGE ownership from Verse8MarketOwner → VERSE8_ADMIN EOA
        assertEq(marketForge.owner(), address(forgeMarketOwner), "FORGE owner is Verse8MarketOwner");

        vm.prank(VERSE8_ADMIN);
        forgeMarketOwner.execute(
            address(marketForge), 0, abi.encodeWithSignature("transferOwnership(address)", VERSE8_ADMIN)
        );
        assertEq(marketForge.owner(), VERSE8_ADMIN, "FORGE owner transferred to EOA");

        // Phase 3-1: Upgrade FORGE market + pairs (now as direct EOA owner)
        vm.startPrank(VERSE8_ADMIN);

        marketForge.upgradeToAndCall(newMarketImpl, hex"");
        marketForge.setPairImpl(newPairImpl);
        assertEq(marketForge.version(), 3);

        for (uint256 i = 0; i < pairsForge.length; ++i) {
            pairsForge[i].upgradeToAndCall(newPairImpl, hex"");
            assertEq(pairsForge[i].version(), 3);
        }

        vm.stopPrank();
    }

    function _phase4_1_migrateFeeControllers() internal {
        // ── 4-1-1. wCROSS: V2Compat → new V2Compat ──
        _migrateFeeControllerAs(MARKET_OWNER_A, marketWCross, pairsWCross, address(newFcV2Compat));

        // ── 4-1-2. CROSSD: V2Compat → new V2Compat ──
        _migrateFeeControllerAs(MARKET_OWNER_A, marketCrossD, pairsCrossD, address(newFcV2Compat));

        // ── 4-1-3. FORGE: V3Split → new V3Split (VERSE8_ADMIN still owns the market) ──
        _migrateFeeControllerAs(VERSE8_ADMIN, marketForge, pairsForge, address(newFcV3Split));

        // ── 4-1-4. Return FORGE ownership to Verse8MarketOwner ──
        vm.prank(VERSE8_ADMIN);
        marketForge.transferOwnership(address(forgeMarketOwner));
        assertEq(marketForge.owner(), address(forgeMarketOwner), "FORGE owner returned to Verse8MarketOwner");
    }

    function _migrateFeeControllerAs(address owner, MarketImplV3 market, PairImplV3[] storage pairs, address newFc)
        internal
    {
        bytes[] memory configs = new bytes[](pairs.length);
        for (uint256 i = 0; i < pairs.length; ++i) {
            (, configs[i]) = pairs[i].getFeeControllerConfig();
        }

        vm.startPrank(owner);

        market.setFeeController(new address[](0), newFc, hex"");
        assertEq(market.feeController(), newFc);

        for (uint256 i = 0; i < pairs.length; ++i) {
            pairs[i].setFeeController(newFc, configs[i]);
            assertEq(address(pairs[i].feeController()), newFc);
        }

        vm.stopPrank();
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Order helpers
    // ─────────────────────────────────────────────────────────────────────────

    function _placeOrdersOnAllPairs() internal {
        uint256 price = 100 * QUOTE_DECIMALS;
        uint256 amount = 10 * BASE_DECIMALS;

        for (uint256 i = 0; i < pairsWCross.length; ++i) {
            _placeSellOrder(USER1, address(pairsWCross[i]), price, amount);
        }
        for (uint256 i = 0; i < pairsCrossD.length; ++i) {
            _placeSellOrder(USER1, address(pairsCrossD[i]), price, amount);
        }
        for (uint256 i = 0; i < pairsForge.length; ++i) {
            _placeSellOrder(USER1, address(pairsForge[i]), price, amount);
        }
    }

    function _placeSellOrder(address user, address pair, uint256 price, uint256 amount) internal {
        vm.prank(user);
        router.submitSellLimit(pair, price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Snapshot & verification
    // ─────────────────────────────────────────────────────────────────────────

    struct _PairSnapshot {
        uint256 baseReserve;
        uint256 quoteReserve;
        uint32 sMk;
        uint32 sTk;
        uint32 bMk;
        uint32 bTk;
    }

    struct _SnapshotAll {
        _PairSnapshot[] wCross;
        _PairSnapshot[] crossD;
        _PairSnapshot[] forge;
    }

    function _snapshotPair(PairImplV3 pair) internal returns (_PairSnapshot memory s) {
        (s.sMk, s.sTk, s.bMk, s.bTk) = pair.getEffectiveFees();
        s.baseReserve = pair.baseReserve();
        s.quoteReserve = pair.quoteReserve();
    }

    function _snapshotAll() internal returns (_SnapshotAll memory s) {
        s.wCross = new _PairSnapshot[](pairsWCross.length);
        s.crossD = new _PairSnapshot[](pairsCrossD.length);
        s.forge = new _PairSnapshot[](pairsForge.length);

        for (uint256 i = 0; i < pairsWCross.length; ++i) {
            s.wCross[i] = _snapshotPair(pairsWCross[i]);
        }
        for (uint256 i = 0; i < pairsCrossD.length; ++i) {
            s.crossD[i] = _snapshotPair(pairsCrossD[i]);
        }
        for (uint256 i = 0; i < pairsForge.length; ++i) {
            s.forge[i] = _snapshotPair(pairsForge[i]);
        }
    }

    function _verifyPostUpgrade(_SnapshotAll memory before) internal {
        // Verify fee configs preserved
        for (uint256 i = 0; i < pairsWCross.length; ++i) {
            _PairSnapshot memory after_ = _snapshotPair(pairsWCross[i]);
            assertEq(after_.sMk, before.wCross[i].sMk, "wCross sMk preserved");
            assertEq(after_.sTk, before.wCross[i].sTk, "wCross sTk preserved");
            assertEq(after_.bMk, before.wCross[i].bMk, "wCross bMk preserved");
            assertEq(after_.bTk, before.wCross[i].bTk, "wCross bTk preserved");
            assertEq(after_.baseReserve, before.wCross[i].baseReserve, "wCross baseReserve preserved");
        }

        for (uint256 i = 0; i < pairsCrossD.length; ++i) {
            _PairSnapshot memory after_ = _snapshotPair(pairsCrossD[i]);
            assertEq(after_.sMk, before.crossD[i].sMk, "crossD sMk preserved");
            assertEq(after_.sTk, before.crossD[i].sTk, "crossD sTk preserved");
            assertEq(after_.baseReserve, before.crossD[i].baseReserve, "crossD baseReserve preserved");
        }

        for (uint256 i = 0; i < pairsForge.length; ++i) {
            _PairSnapshot memory after_ = _snapshotPair(pairsForge[i]);
            assertEq(after_.sMk, before.forge[i].sMk, "forge sMk preserved");
            assertEq(after_.sTk, before.forge[i].sTk, "forge sTk preserved");
            assertEq(after_.baseReserve, before.forge[i].baseReserve, "forge baseReserve preserved");
        }

        // Verify market-level fee controller updated
        assertEq(marketWCross.feeController(), address(newFcV2Compat), "wCross market FC updated");
        assertEq(marketCrossD.feeController(), address(newFcV2Compat), "crossD market FC updated");
        assertEq(marketForge.feeController(), address(newFcV3Split), "forge market FC updated");

        // Verify every wCROSS pair points to new V2Compat + feeCollector preserved
        for (uint256 i = 0; i < pairsWCross.length; ++i) {
            assertEq(address(pairsWCross[i].feeController()), address(newFcV2Compat), "wCross pair FC updated");
            (, bytes memory data) = pairsWCross[i].getFeeControllerConfig();
            (address fc,,,,) = abi.decode(data, (address, uint32, uint32, uint32, uint32));
            assertEq(fc, FEE_COLLECTOR, "wCross pair feeCollector preserved");
        }

        // Verify every CROSSD pair points to new V2Compat + feeCollector preserved
        for (uint256 i = 0; i < pairsCrossD.length; ++i) {
            assertEq(address(pairsCrossD[i].feeController()), address(newFcV2Compat), "crossD pair FC updated");
            (, bytes memory data) = pairsCrossD[i].getFeeControllerConfig();
            (address fc,,,,) = abi.decode(data, (address, uint32, uint32, uint32, uint32));
            assertEq(fc, FEE_COLLECTOR, "crossD pair feeCollector preserved");
        }

        // Verify every FORGE pair points to new V3Split + per-pair creator preserved
        for (uint256 i = 0; i < pairsForge.length; ++i) {
            assertEq(address(pairsForge[i].feeController()), address(newFcV3Split), "forge pair FC updated");
            (, bytes memory data) = pairsForge[i].getFeeControllerConfig();
            (address fc, address creator,,,) = abi.decode(data, (address, address, uint32, uint32, uint32));
            address expectedCreator = address(uint160(0xC0DE0000 + i));
            assertEq(creator, expectedCreator, "forge pair creator preserved");
            assertEq(fc, FEE_COLLECTOR, "forge pair feeCollector preserved");
        }

        // Verify can still place orders after upgrade
        uint256 price = 200 * QUOTE_DECIMALS;
        uint256 amount = 5 * BASE_DECIMALS;
        _placeSellOrder(USER2, address(pairsWCross[0]), price, amount);
        _placeSellOrder(USER2, address(pairsCrossD[0]), price, amount);
        _placeSellOrder(USER2, address(pairsForge[0]), price, amount);
    }
}
