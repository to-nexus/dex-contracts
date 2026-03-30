// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";
import {Test} from "forge-std/Test.sol";

import {CrossDexImplV3} from "../src/CrossDexImplV3.sol";
import {CrossDexRouterV3} from "../src/CrossDexRouterV3.sol";
import {FeeControllerV2Compat} from "../src/FeeControllerV2Compat.sol";
import {FeeControllerV3Dist} from "../src/FeeControllerV3Dist.sol";
import {FeeControllerV3Split} from "../src/FeeControllerV3Split.sol";
import {MarketImplV3} from "../src/MarketImplV3.sol";
import {PairImplV3} from "../src/PairImplV3.sol";

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

/// @title TestnetUpgradeTest
/// @notice Simulates the full testnet upgrade scenario:
///   1-1. Deploy new implementations
///   1-2. Upgrade CrossDex + Router, swap fee controllers
///   2.   Upgrade Game/CROSSD markets (Owner A, same key)
///   3-0. Transfer FORGE market ownership from Verse8MarketOwner → EOA
///   3-1. Upgrade FORGE market (Verse8 Admin)
///   4-1. Migrate fee controllers (same type, preserving existing config)
///   4-1-4. Transfer FORGE market ownership back to Verse8MarketOwner
///   4-2. Convert selected pairs from V2Compat → V3Dist (3-way distribution)
///   +    Create SHILTZx pair with V3Dist and verify fee distribution
contract TestnetUpgradeTest is Test {
    // ─────────────────────────────────────────────────────────────────────────
    // Actors
    // ─────────────────────────────────────────────────────────────────────────

    address constant CROSSDEX_OWNER = address(bytes20("CROSSDEX_OWNER"));
    address constant MARKET_OWNER_A = address(bytes20("MARKET_OWNER_A"));

    // Verse8 Admin = MARKET_OWNER_A (same key as testnet)
    address constant VERSE8_ADMIN = MARKET_OWNER_A;
    address constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    address constant USER1 = address(bytes20("USER1"));
    address constant USER2 = address(bytes20("USER2"));

    // V3Dist fee recipients
    address constant FOUNDATION = address(bytes20("FOUNDATION"));
    address constant NEXUS = address(bytes20("NEXUS"));
    address constant DEV_A = address(bytes20("DEV_A"));
    address constant DEV_B = address(bytes20("DEV_B"));
    address constant DEV_SEALM = address(bytes20("DEV_SEALM"));

    // ─────────────────────────────────────────────────────────────────────────
    // V3Dist constants
    // ─────────────────────────────────────────────────────────────────────────

    uint32 constant RATIO_FOUNDATION = 1000;
    uint32 constant RATIO_NEXUS = 4500;
    uint32 constant RATIO_DEVELOPER = 4500;

    uint32 constant V3DIST_SELLER_MAKER = 500;
    uint32 constant V3DIST_SELLER_TAKER = 500;

    bytes32 constant LABEL_FOUNDATION = "FOUNDATION";
    bytes32 constant LABEL_NEXUS = "NEXUS";
    bytes32 constant LABEL_DEVELOPER = "DEVELOPER";

    // ─────────────────────────────────────────────────────────────────────────
    // Contracts — "current testnet" state
    // ─────────────────────────────────────────────────────────────────────────

    CrossDexImplV3 crossDex;
    CrossDexRouterV3 router;

    MarketImplV3 marketGame;
    MarketImplV3 marketCrossD;
    MarketImplV3 marketForge;

    MockVerse8MarketOwner forgeMarketOwner;

    // Game/CROSSD: pairs[0,1] → V3Dist, pairs[2] stays V2Compat
    PairImplV3[] pairsGame;
    PairImplV3[] pairsCrossD;
    PairImplV3[] pairsForge;

    FeeControllerV2Compat oldFcV2Compat;
    FeeControllerV3Split oldFcV3Split;

    IERC20 quoteWCross;
    IERC20 quoteCrossD;
    IERC20[] basesGame;
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

    // SHILTZx pair (created after Phase 4-2)
    IERC20 shiltzxToken;
    PairImplV3 pairShiltzxGame;
    PairImplV3 pairShiltzxCrossD;

    // ─────────────────────────────────────────────────────────────────────────
    // Setup — simulate current testnet deployment
    // ─────────────────────────────────────────────────────────────────────────

    function setUp() public {
        vm.label(CROSSDEX_OWNER, "crossDexOwner");
        vm.label(MARKET_OWNER_A, "marketOwnerA / verse8Admin");
        vm.label(FEE_COLLECTOR, "feeCollector");
        vm.label(USER1, "user1");
        vm.label(USER2, "user2");
        vm.label(FOUNDATION, "foundation");
        vm.label(NEXUS, "nexus");
        vm.label(DEV_A, "devA");
        vm.label(DEV_B, "devB");
        vm.label(DEV_SEALM, "devSealM");

        _deployCurrentTestnetState();
        _seedUsersAndApprovals();
    }

    function _deployCurrentTestnetState() internal {
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

        // ── Game Market (Owner A, V2Compat, 3 pairs) ──
        quoteWCross = new T20("wCROSS", "CROSS", 18);
        address mGame = crossDex.createMarket(MARKET_OWNER_A, address(quoteWCross), address(oldFcV2Compat), "Game");
        marketGame = MarketImplV3(mGame);

        vm.stopPrank();
        vm.startPrank(MARKET_OWNER_A);

        bytes memory v2InitData = abi.encode(FEE_COLLECTOR, V2_SELLER_MAKER, V2_SELLER_TAKER, uint32(0), uint32(0));
        for (uint256 i = 0; i < 3; ++i) {
            T20 base = new T20(string(abi.encodePacked("BASE_G", vm.toString(i))), "BG", 18);
            basesGame.push(base);
            address pair = marketGame.createPair(address(base), TICK_SIZE, LOT_SIZE, v2InitData);
            pairsGame.push(PairImplV3(pair));
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
        _fundAndApprove(USER1, quoteWCross, basesGame);
        _fundAndApprove(USER2, quoteWCross, basesGame);
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
    // Main test: full upgrade + fee upgrade + SHILTZx creation
    // ─────────────────────────────────────────────────────────────────────────

    function test_full_upgrade_scenario() public {
        // ── Pre-upgrade: place orders to verify state preservation ──
        _placeOrdersOnAllPairs();

        _SnapshotAll memory before = _snapshotAll();

        // ── Phase 1-1: Deploy new implementations ──
        _phase1_1_deployImplementations();

        // ── Phase 1-2: Upgrade CrossDex + Router ──
        _phase1_2_upgradeCrossDexAndRouter();

        // ── Phase 2: Upgrade Game + CROSSD markets (Owner A) ──
        _phase2_upgradeMarketsOwnerA();

        // ── Phase 3: Upgrade FORGE market (Verse8MarketOwner → EOA → upgrade) ──
        _phase3_upgradeMarketForge();

        // ── Phase 4-1: Migrate fee controllers (same type) ──
        _phase4_1_migrateFeeControllers();

        // ── Post Phase 4-1: verify state preservation ──
        _verifyPostUpgrade(before);

        // ── Phase 4-2: V2Compat → V3Dist for selected pairs ──
        _phase4_2_convertToV3Dist();

        // ── Verify V3Dist configuration ──
        _verifyV3DistConfiguration();

        // ── Create SHILTZx pairs ──
        _createShiltzxPairs();

        // ── Verify V3Dist fee distribution via trade ──
        _verifyV3DistFeeDistribution();

        // ── Verify post-conversion trading on all pair types ──
        _verifyTradingAfterAllChanges();
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Phase 1-1
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

    // ─────────────────────────────────────────────────────────────────────────
    // Phase 1-2
    // ─────────────────────────────────────────────────────────────────────────

    function _phase1_2_upgradeCrossDexAndRouter() internal {
        address oldMarketImpl = crossDex.marketImpl();
        address oldPairImpl = crossDex.pairImpl();

        vm.startPrank(CROSSDEX_OWNER);

        crossDex.upgradeToAndCall(newCrossDexImpl, hex"");
        assertEq(crossDex.version(), 3);

        router.upgradeToAndCall(newRouterImpl, hex"");
        assertEq(router.version(), 3);

        crossDex.setMarketImpl(newMarketImpl);
        crossDex.setPairImpl(newPairImpl);
        assertEq(crossDex.marketImpl(), newMarketImpl);
        assertEq(crossDex.pairImpl(), newPairImpl);
        assertTrue(crossDex.marketImpl() != oldMarketImpl, "marketImpl changed");
        assertTrue(crossDex.pairImpl() != oldPairImpl, "pairImpl changed");

        crossDex.setFeeControllerAllow(address(oldFcV2Compat), false);
        crossDex.setFeeControllerAllow(address(oldFcV3Split), false);
        crossDex.setFeeControllerAllow(address(newFcV2Compat), true);
        crossDex.setFeeControllerAllow(address(newFcV3Dist), true);
        crossDex.setFeeControllerAllow(address(newFcV3Split), true);

        vm.stopPrank();

        vm.expectRevert();
        crossDex.checkFeeControllerAllowed(address(oldFcV2Compat));
        vm.expectRevert();
        crossDex.checkFeeControllerAllowed(address(oldFcV3Split));

        crossDex.checkFeeControllerAllowed(address(newFcV2Compat));
        crossDex.checkFeeControllerAllowed(address(newFcV3Dist));
        crossDex.checkFeeControllerAllowed(address(newFcV3Split));
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Phase 2
    // ─────────────────────────────────────────────────────────────────────────

    function _phase2_upgradeMarketsOwnerA() internal {
        vm.startPrank(MARKET_OWNER_A);

        marketGame.upgradeToAndCall(newMarketImpl, hex"");
        marketGame.setPairImpl(newPairImpl);
        assertEq(marketGame.version(), 3);
        assertEq(marketGame.pairImpl(), newPairImpl);

        for (uint256 i = 0; i < pairsGame.length; ++i) {
            pairsGame[i].upgradeToAndCall(newPairImpl, hex"");
            assertEq(pairsGame[i].version(), 3);
        }

        marketCrossD.upgradeToAndCall(newMarketImpl, hex"");
        marketCrossD.setPairImpl(newPairImpl);
        assertEq(marketCrossD.version(), 3);

        for (uint256 i = 0; i < pairsCrossD.length; ++i) {
            pairsCrossD[i].upgradeToAndCall(newPairImpl, hex"");
            assertEq(pairsCrossD[i].version(), 3);
        }

        vm.stopPrank();
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Phase 3 — FORGE: Verse8MarketOwner → EOA → upgrade → (ownership returned in 4-1-4)
    // ─────────────────────────────────────────────────────────────────────────

    function _phase3_upgradeMarketForge() internal {
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

    // ─────────────────────────────────────────────────────────────────────────
    // Phase 4-1 — FC migration (same type, config preserved)
    // ─────────────────────────────────────────────────────────────────────────

    function _phase4_1_migrateFeeControllers() internal {
        // 4-1-1. Game Market (V2Compat → new V2Compat)
        _migrateFeeControllerAs(MARKET_OWNER_A, marketGame, pairsGame, address(newFcV2Compat));

        // 4-1-2. CROSSD Market (V2Compat → new V2Compat)
        _migrateFeeControllerAs(MARKET_OWNER_A, marketCrossD, pairsCrossD, address(newFcV2Compat));

        // 4-1-3. FORGE Market (V3Split → new V3Split, VERSE8_ADMIN still owns the market)
        _migrateFeeControllerAs(VERSE8_ADMIN, marketForge, pairsForge, address(newFcV3Split));

        // 4-1-4. Return FORGE ownership to Verse8MarketOwner
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
    // Phase 4-2 — V2Compat → V3Dist for selected pairs
    // ─────────────────────────────────────────────────────────────────────────

    function _phase4_2_convertToV3Dist() internal {
        vm.startPrank(MARKET_OWNER_A);

        // ── Game Market: pairs[0] → V3Dist (DEV_A), pairs[1] → V3Dist (DEV_B) ──
        address[] memory batch = new address[](1);

        batch[0] = address(pairsGame[0]);
        marketGame.setFeeController(batch, address(newFcV3Dist), _buildV3DistInitData(DEV_A));

        batch[0] = address(pairsGame[1]);
        marketGame.setFeeController(batch, address(newFcV3Dist), _buildV3DistInitData(DEV_B));

        // ── CROSSD Market: pairs[0] → V3Dist (DEV_A), pairs[1] → V3Dist (DEV_B) ──
        batch[0] = address(pairsCrossD[0]);
        marketCrossD.setFeeController(batch, address(newFcV3Dist), _buildV3DistInitData(DEV_A));

        batch[0] = address(pairsCrossD[1]);
        marketCrossD.setFeeController(batch, address(newFcV3Dist), _buildV3DistInitData(DEV_B));

        vm.stopPrank();
    }

    // ─────────────────────────────────────────────────────────────────────────
    // SHILTZx pair creation
    // ─────────────────────────────────────────────────────────────────────────

    function _createShiltzxPairs() internal {
        shiltzxToken = new T20("SHILTZx", "SHILTZX", 18);

        shiltzxToken.transfer(USER1, 100_000 * BASE_DECIMALS);
        shiltzxToken.transfer(USER2, 100_000 * BASE_DECIMALS);

        vm.prank(USER1);
        shiltzxToken.approve(address(router), type(uint256).max);
        vm.prank(USER2);
        shiltzxToken.approve(address(router), type(uint256).max);

        bytes memory initSealM = _buildV3DistInitData(DEV_SEALM);

        vm.startPrank(MARKET_OWNER_A);

        address pGame = marketGame.createPair(address(shiltzxToken), TICK_SIZE, LOT_SIZE, initSealM);
        pairShiltzxGame = PairImplV3(pGame);

        address pCrossD = marketCrossD.createPair(address(shiltzxToken), TICK_SIZE, LOT_SIZE, initSealM);
        pairShiltzxCrossD = PairImplV3(pCrossD);

        vm.stopPrank();

        assertEq(address(pairShiltzxGame.feeController()), address(newFcV3Dist), "SHILTZx Game FC is V3Dist");
        assertEq(address(pairShiltzxCrossD.feeController()), address(newFcV3Dist), "SHILTZx CROSSD FC is V3Dist");
    }

    // ─────────────────────────────────────────────────────────────────────────
    // V3Dist init data builder
    // ─────────────────────────────────────────────────────────────────────────

    function _buildV3DistInitData(address developer) internal pure returns (bytes memory) {
        address[] memory recipients = new address[](3);
        recipients[0] = FOUNDATION;
        recipients[1] = NEXUS;
        recipients[2] = developer;

        uint32[] memory ratios = new uint32[](3);
        ratios[0] = RATIO_FOUNDATION;
        ratios[1] = RATIO_NEXUS;
        ratios[2] = RATIO_DEVELOPER;

        bytes32[] memory labels = new bytes32[](3);
        labels[0] = LABEL_FOUNDATION;
        labels[1] = LABEL_NEXUS;
        labels[2] = LABEL_DEVELOPER;

        return abi.encode(V3DIST_SELLER_MAKER, V3DIST_SELLER_TAKER, uint32(0), uint32(0), recipients, ratios, labels);
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Order helpers
    // ─────────────────────────────────────────────────────────────────────────

    function _placeOrdersOnAllPairs() internal {
        uint256 price = 100 * QUOTE_DECIMALS;
        uint256 amount = 10 * BASE_DECIMALS;

        for (uint256 i = 0; i < pairsGame.length; ++i) {
            _placeSellOrder(USER1, address(pairsGame[i]), price, amount);
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

    function _placeBuyOrder(address user, address pair, uint256 price, uint256 amount) internal {
        vm.prank(user);
        router.submitBuyLimit(pair, price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Snapshot & post-upgrade verification (Phase 4-1 state preservation)
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
        _PairSnapshot[] game;
        _PairSnapshot[] crossD;
        _PairSnapshot[] forge;
    }

    function _snapshotPair(PairImplV3 pair) internal returns (_PairSnapshot memory s) {
        (s.sMk, s.sTk, s.bMk, s.bTk) = pair.getEffectiveFees();
        s.baseReserve = pair.baseReserve();
        s.quoteReserve = pair.quoteReserve();
    }

    function _snapshotAll() internal returns (_SnapshotAll memory s) {
        s.game = new _PairSnapshot[](pairsGame.length);
        s.crossD = new _PairSnapshot[](pairsCrossD.length);
        s.forge = new _PairSnapshot[](pairsForge.length);

        for (uint256 i = 0; i < pairsGame.length; ++i) {
            s.game[i] = _snapshotPair(pairsGame[i]);
        }
        for (uint256 i = 0; i < pairsCrossD.length; ++i) {
            s.crossD[i] = _snapshotPair(pairsCrossD[i]);
        }
        for (uint256 i = 0; i < pairsForge.length; ++i) {
            s.forge[i] = _snapshotPair(pairsForge[i]);
        }
    }

    function _verifyPostUpgrade(_SnapshotAll memory before) internal {
        // Fee configs preserved
        for (uint256 i = 0; i < pairsGame.length; ++i) {
            _PairSnapshot memory a = _snapshotPair(pairsGame[i]);
            assertEq(a.sMk, before.game[i].sMk, "game sMk preserved");
            assertEq(a.sTk, before.game[i].sTk, "game sTk preserved");
            assertEq(a.bMk, before.game[i].bMk, "game bMk preserved");
            assertEq(a.bTk, before.game[i].bTk, "game bTk preserved");
            assertEq(a.baseReserve, before.game[i].baseReserve, "game baseReserve preserved");
        }

        for (uint256 i = 0; i < pairsCrossD.length; ++i) {
            _PairSnapshot memory a = _snapshotPair(pairsCrossD[i]);
            assertEq(a.sMk, before.crossD[i].sMk, "crossD sMk preserved");
            assertEq(a.sTk, before.crossD[i].sTk, "crossD sTk preserved");
            assertEq(a.baseReserve, before.crossD[i].baseReserve, "crossD baseReserve preserved");
        }

        for (uint256 i = 0; i < pairsForge.length; ++i) {
            _PairSnapshot memory a = _snapshotPair(pairsForge[i]);
            assertEq(a.sMk, before.forge[i].sMk, "forge sMk preserved");
            assertEq(a.sTk, before.forge[i].sTk, "forge sTk preserved");
            assertEq(a.baseReserve, before.forge[i].baseReserve, "forge baseReserve preserved");
        }

        // Market-level FC updated
        assertEq(marketGame.feeController(), address(newFcV2Compat), "game market FC updated");
        assertEq(marketCrossD.feeController(), address(newFcV2Compat), "crossD market FC updated");
        assertEq(marketForge.feeController(), address(newFcV3Split), "forge market FC updated");

        // All Game/CROSSD pairs on new V2Compat, feeCollector preserved
        for (uint256 i = 0; i < pairsGame.length; ++i) {
            assertEq(address(pairsGame[i].feeController()), address(newFcV2Compat), "game pair FC updated");
            (, bytes memory data) = pairsGame[i].getFeeControllerConfig();
            (address fc,,,,) = abi.decode(data, (address, uint32, uint32, uint32, uint32));
            assertEq(fc, FEE_COLLECTOR, "game pair feeCollector preserved");
        }

        for (uint256 i = 0; i < pairsCrossD.length; ++i) {
            assertEq(address(pairsCrossD[i].feeController()), address(newFcV2Compat), "crossD pair FC updated");
            (, bytes memory data) = pairsCrossD[i].getFeeControllerConfig();
            (address fc,,,,) = abi.decode(data, (address, uint32, uint32, uint32, uint32));
            assertEq(fc, FEE_COLLECTOR, "crossD pair feeCollector preserved");
        }

        // FORGE pairs on new V3Split, per-pair creator preserved
        for (uint256 i = 0; i < pairsForge.length; ++i) {
            assertEq(address(pairsForge[i].feeController()), address(newFcV3Split), "forge pair FC updated");
            (, bytes memory data) = pairsForge[i].getFeeControllerConfig();
            (address fc, address creator,,,) = abi.decode(data, (address, address, uint32, uint32, uint32));
            address expectedCreator = address(uint160(0xC0DE0000 + i));
            assertEq(creator, expectedCreator, "forge pair creator preserved");
            assertEq(fc, FEE_COLLECTOR, "forge pair feeCollector preserved");
        }
    }

    // ─────────────────────────────────────────────────────────────────────────
    // V3Dist configuration verification (after Phase 4-2)
    // ─────────────────────────────────────────────────────────────────────────

    function _verifyV3DistConfiguration() internal {
        // Market-level FC should now be V3Dist
        assertEq(marketGame.feeController(), address(newFcV3Dist), "game market FC is V3Dist");
        assertEq(marketCrossD.feeController(), address(newFcV3Dist), "crossD market FC is V3Dist");

        // Converted pairs → V3Dist with correct config
        _assertV3DistConfig(pairsGame[0], DEV_A, "game[0]");
        _assertV3DistConfig(pairsGame[1], DEV_B, "game[1]");
        _assertV3DistConfig(pairsCrossD[0], DEV_A, "crossD[0]");
        _assertV3DistConfig(pairsCrossD[1], DEV_B, "crossD[1]");

        // Non-converted pairs → still V2Compat
        assertEq(address(pairsGame[2].feeController()), address(newFcV2Compat), "game[2] stays V2Compat");
        assertEq(address(pairsCrossD[2].feeController()), address(newFcV2Compat), "crossD[2] stays V2Compat");

        // V2Compat pairs: fee rates unchanged
        (uint32 sMk, uint32 sTk, uint32 bMk, uint32 bTk) = pairsGame[2].getEffectiveFees();
        assertEq(sMk, V2_SELLER_MAKER, "game[2] V2Compat sMk unchanged");
        assertEq(sTk, V2_SELLER_TAKER, "game[2] V2Compat sTk unchanged");
        assertEq(bMk, 0, "game[2] V2Compat bMk unchanged");
        assertEq(bTk, 0, "game[2] V2Compat bTk unchanged");
    }

    function _assertV3DistConfig(PairImplV3 pair, address expectedDev, string memory label) internal {
        assertEq(address(pair.feeController()), address(newFcV3Dist), string.concat(label, " FC is V3Dist"));

        (uint32 sMk, uint32 sTk, uint32 bMk, uint32 bTk) = pair.getEffectiveFees();
        assertEq(sMk, V3DIST_SELLER_MAKER, string.concat(label, " sMk"));
        assertEq(sTk, V3DIST_SELLER_TAKER, string.concat(label, " sTk"));
        assertEq(bMk, 0, string.concat(label, " bMk"));
        assertEq(bTk, 0, string.concat(label, " bTk"));

        (, bytes memory data) = pair.getFeeControllerConfig();
        (
            uint32 dsMk,
            uint32 dsTk,
            uint32 dbMk,
            uint32 dbTk,
            address[] memory recipients,
            uint32[] memory ratios,
            bytes32[] memory labels
        ) = abi.decode(data, (uint32, uint32, uint32, uint32, address[], uint32[], bytes32[]));

        assertEq(dsMk, V3DIST_SELLER_MAKER, string.concat(label, " storage sMk"));
        assertEq(dsTk, V3DIST_SELLER_TAKER, string.concat(label, " storage sTk"));
        assertEq(dbMk, 0, string.concat(label, " storage bMk"));
        assertEq(dbTk, 0, string.concat(label, " storage bTk"));

        assertEq(recipients.length, 3, string.concat(label, " 3 recipients"));
        assertEq(recipients[0], FOUNDATION, string.concat(label, " recipient[0]"));
        assertEq(recipients[1], NEXUS, string.concat(label, " recipient[1]"));
        assertEq(recipients[2], expectedDev, string.concat(label, " recipient[2]"));

        assertEq(ratios[0], RATIO_FOUNDATION, string.concat(label, " ratio[0]"));
        assertEq(ratios[1], RATIO_NEXUS, string.concat(label, " ratio[1]"));
        assertEq(ratios[2], RATIO_DEVELOPER, string.concat(label, " ratio[2]"));

        assertEq(labels[0], LABEL_FOUNDATION, string.concat(label, " label[0]"));
        assertEq(labels[1], LABEL_NEXUS, string.concat(label, " label[1]"));
        assertEq(labels[2], LABEL_DEVELOPER, string.concat(label, " label[2]"));
    }

    // ─────────────────────────────────────────────────────────────────────────
    // V3Dist fee distribution verification via actual trade
    // ─────────────────────────────────────────────────────────────────────────

    function _verifyV3DistFeeDistribution() internal {
        // Trade on SHILTZx/wCROSS pair (V3Dist, DEV_SEALM)
        uint256 price = 200 * QUOTE_DECIMALS;
        uint256 amount = 5 * BASE_DECIMALS;
        // tradeQuoteAmount = 200 * 5 = 1000 QUOTE_DECIMALS
        // makerFee (seller) = 1000e18 * 500 / 10000 = 50e18
        // takerFee (buyer)  = 0
        // Distribution: FOUNDATION 10% = 5e18, NEXUS 45% = 22.5e18, DEV_SEALM 45% = 22.5e18

        uint256 foundationBefore = quoteWCross.balanceOf(FOUNDATION);
        uint256 nexusBefore = quoteWCross.balanceOf(NEXUS);
        uint256 devSealmBefore = quoteWCross.balanceOf(DEV_SEALM);

        // USER1 places sell limit (becomes maker)
        _placeSellOrder(USER1, address(pairShiltzxGame), price, amount);

        // USER2 places buy limit at same price (crosses → match)
        _placeBuyOrder(USER2, address(pairShiltzxGame), price, amount);

        uint256 tradeQuoteAmount = Math.mulDiv(price, amount, BASE_DECIMALS);
        uint256 totalFee = Math.mulDiv(tradeQuoteAmount, V3DIST_SELLER_MAKER, BPS_DENOMINATOR);
        uint256 expectedFoundation = Math.mulDiv(totalFee, RATIO_FOUNDATION, BPS_DENOMINATOR);
        uint256 expectedNexus = Math.mulDiv(totalFee, RATIO_NEXUS, BPS_DENOMINATOR);
        uint256 expectedDev = totalFee - expectedFoundation - expectedNexus;

        assertEq(quoteWCross.balanceOf(FOUNDATION) - foundationBefore, expectedFoundation, "FOUNDATION received 10%");
        assertEq(quoteWCross.balanceOf(NEXUS) - nexusBefore, expectedNexus, "NEXUS received 45%");
        assertEq(quoteWCross.balanceOf(DEV_SEALM) - devSealmBefore, expectedDev, "DEV_SEALM received 45%");

        assertTrue(totalFee > 0, "fee was collected");
        assertEq(expectedFoundation + expectedNexus + expectedDev, totalFee, "distribution sums to total");
    }

    // ─────────────────────────────────────────────────────────────────────────
    // Post-all-changes: verify trading on every pair type
    // ─────────────────────────────────────────────────────────────────────────

    function _verifyTradingAfterAllChanges() internal {
        uint256 price = 300 * QUOTE_DECIMALS;
        uint256 amount = 2 * BASE_DECIMALS;

        // V3Dist pairs (converted)
        _placeSellOrder(USER2, address(pairsGame[0]), price, amount);
        _placeSellOrder(USER2, address(pairsCrossD[0]), price, amount);

        // V2Compat pairs (not converted)
        _placeSellOrder(USER2, address(pairsGame[2]), price, amount);
        _placeSellOrder(USER2, address(pairsCrossD[2]), price, amount);

        // V3Split pairs (FORGE)
        _placeSellOrder(USER2, address(pairsForge[0]), price, amount);

        // SHILTZx pairs
        _placeSellOrder(USER2, address(pairShiltzxGame), price, amount);
        _placeSellOrder(USER2, address(pairShiltzxCrossD), price, amount);
    }
}
