// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

// ═══════════════════════════════════════════════════════════════════════════════════════════════════
//
//  ██╗   ██╗██████╗     ████████╗ ██████╗     ██╗   ██╗██████╗      █████╗ ██╗   ██╗██████╗ ██╗████████╗
//  ██║   ██║╚════██╗    ╚══██╔══╝██╔═══██╗    ██║   ██║╚════██╗    ██╔══██╗██║   ██║██╔══██╗██║╚══██╔══╝
//  ██║   ██║ █████╔╝       ██║   ██║   ██║    ██║   ██║ █████╔╝    ███████║██║   ██║██║  ██║██║   ██║
//  ╚██╗ ██╔╝██╔═══╝        ██║   ██║   ██║    ╚██╗ ██╔╝ ╚═══██╗    ██╔══██║██║   ██║██║  ██║██║   ██║
//   ╚████╔╝ ███████╗       ██║   ╚██████╔╝     ╚████╔╝ ██████╔╝    ██║  ██║╚██████╔╝██████╔╝██║   ██║
//    ╚═══╝  ╚══════╝       ╚═╝    ╚═════╝       ╚═══╝  ╚═════╝     ╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚═╝   ╚═╝
//
// ═══════════════════════════════════════════════════════════════════════════════════════════════════
//
// V2 → V3 업그레이드 시 주의사항
// ================================
//
// 이 테스트는 DEX 컨트랙트를 V2에서 V3로 업그레이드할 때 발생할 수 있는 스토리지 충돌 문제를
// 검증하고, 안전한 업그레이드 절차를 문서화합니다.
//
// ┌───────────────────┬─────────────────────────────────────────────────────────────────────────────┐
// │  업그레이드 경로  │                              테스트 결과                                   │
// ├───────────────────┼─────────────────────────────────────────────────────────────────────────────┤
// │ CrossDexV2 → V3   │ ✅ 업그레이드 가능                                                         │
// │                   │    - 업그레이드 후 setFeeControllerAllow() 호출 필요                       │
// │                   │    - _allowedFeeControllers가 새로 추가되어 초기화 필요                    │
// ├───────────────────┼─────────────────────────────────────────────────────────────────────────────┤
// │ RouterV2 → V3     │ ✅ 업그레이드 가능                                                         │
// │                   │    - CrossDex 업그레이드와 함께 진행                                       │
// │                   │    - V3 Router는 calcBuyVolumeWithFee()를 사용 (V2는 getEffectiveFees)     │
// ├───────────────────┼─────────────────────────────────────────────────────────────────────────────┤
// │ MarketV2 → V3     │ ✅ 업그레이드 가능 (단, 즉시 setFeeController 호출 필수)                   │
// │                   │    - feeCollector(V2) → feeController(V3): 같은 슬롯, 값 덮어쓰기 필요     │
// │                   │    - _allPairs 매핑: 같은 슬롯 위치, 충돌 없음                             │
// │                   │    - _feeConfig(V2): V3에서 제거됨, 고아 데이터로 남음 (문제 없음)         │
// ├───────────────────┼─────────────────────────────────────────────────────────────────────────────┤
// │ PairV2 → V3       │ ⚠️ 업그레이드 가능하나 즉시 setFeeController() 호출 필수                   │
// │                   │    - feeConfig(V2, struct) → feeController(V3, address): 슬롯 충돌 발생!   │
// │                   │    - 업그레이드 직후 feeController가 손상된 주소값을 가짐                  │
// │                   │      예: FeeConfig(20,30,0,0) → 0x0000...001E00000014 (잘못된 주소)        │
// │                   │    - setFeeController() 미호출 시 모든 수수료 관련 함수 실패               │
// └───────────────────┴─────────────────────────────────────────────────────────────────────────────┘
//
// ═══════════════════════════════════════════════════════════════════════════════════════════════════
//
// 권장 업그레이드 순서
// ====================
//
// 1단계: CrossDex 업그레이드
//        ├─ upgradeToAndCall(crossDexV3Impl, "")
//        └─ setFeeControllerAllow(feeController, true)
//
// 2단계: Router 업그레이드
//        └─ upgradeToAndCall(routerV3Impl, "")
//
// 3단계: Market 업그레이드 (각 Market별로)
//        ├─ upgradeToAndCall(marketV3Impl, "")
//        └─ setFeeController(0, 0, false, feeController, initData)
//           ※ endIndex=0: V2 Pair들에 전파하지 않음 (아직 V2 상태)
//
// 4단계: Pair 업그레이드 (각 Pair별로)
//        ├─ upgradeToAndCall(pairV3Impl, "")
//        └─ setFeeController(feeController, initData)  ⚠️ 즉시 호출 필수!
//
// ═══════════════════════════════════════════════════════════════════════════════════════════════════
//
// Pair 스토리지 슬롯 분석 (핵심)
// ==============================
//
// V2 PairImplV2 (slot 25 기준):
// ┌─────────┬──────────────────────────────────────────────────────────────┐
// │ slot 25 │ FeeConfig feeConfig (struct: 4 x uint32 = 128 bits)         │
// │         │ ├─ sellerMakerFeeBps: uint32 (예: 20 = 0x14)                │
// │         │ ├─ sellerTakerFeeBps: uint32 (예: 30 = 0x1E)                │
// │         │ ├─ buyerMakerFeeBps:  uint32 (예: 0)                        │
// │         │ └─ buyerTakerFeeBps:  uint32 (예: 0)                        │
// └─────────┴──────────────────────────────────────────────────────────────┘
//
// V3 PairImplV3 (slot 25 기준):
// ┌─────────┬──────────────────────────────────────────────────────────────┐
// │ slot 25 │ IFeeController feeController (address: 160 bits)            │
// └─────────┴──────────────────────────────────────────────────────────────┘
//
// 업그레이드 후 문제:
// - V2의 FeeConfig 바이트가 V3에서 address로 해석됨
// - 예: FeeConfig(20, 30, 0, 0) → address(0x0000...001E00000014)
// - 이 주소는 유효한 컨트랙트가 아니므로 delegatecall 실패
//
// ═══════════════════════════════════════════════════════════════════════════════════════════════════
//
// setFeeController 초기화 데이터 형식
// ====================================
//
// FeeControllerV2Compat 사용 시:
// bytes memory initData = abi.encode(
//     feeCollector,      // address: 수수료 수취 주소
//     sellerMakerFeeBps, // uint32: 판매자 메이커 수수료 (BPS)
//     sellerTakerFeeBps, // uint32: 판매자 테이커 수수료 (BPS)
//     buyerMakerFeeBps,  // uint32: 구매자 메이커 수수료 (BPS)
//     buyerTakerFeeBps   // uint32: 구매자 테이커 수수료 (BPS)
// );
//
// 기존 V2 수수료 정책 유지 시, V2에서 사용하던 동일한 BPS 값을 그대로 사용하면 됩니다.
//
// ═══════════════════════════════════════════════════════════════════════════════════════════════════

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.5.0/proxy/utils/UUPSUpgradeable.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {CrossDexImplV3} from "../src/CrossDexImplV3.sol";
import {CrossDexRouterV3} from "../src/CrossDexRouterV3.sol";
import {FeeControllerV2Compat} from "../src/FeeControllerV2Compat.sol";
import {MarketImplV3} from "../src/MarketImplV3.sol";
import {PairImplV3} from "../src/PairImplV3.sol";

import {BPS_DENOMINATOR} from "../src/interfaces/IFeeController.sol";
import {IPairV3} from "../src/interfaces/IPairV3.sol";

import {T20} from "./mock/T20.sol";

/// @title V2ToV3UpgradeTest
/// @notice V2 → V3 업그레이드 경로 및 스토리지 충돌 검증 테스트
/// @dev vm.getCode()를 사용하여 legacy_v1_v2/out/에서 V2 바이트코드를 로드합니다.
///      테스트 실행 전 `cd legacy_v1_v2 && forge build` 필요.
contract V2ToV3UpgradeTest is Test {
    // ─────────────────────────────────────────────────────────────────────────────
    // Constants
    // ─────────────────────────────────────────────────────────────────────────────

    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    address public constant USER1 = address(bytes20("USER1"));
    address public constant USER2 = address(bytes20("USER2"));

    uint256 public constant FIND_PREV_PRICE_COUNT = type(uint256).max;
    uint256 public constant MAX_MATCH_COUNT = type(uint256).max;
    uint256 public constant CANCEL_LIMIT = type(uint256).max;

    // Default fees
    uint32 public constant SELLER_MAKER_FEE = 20; // 0.2%
    uint32 public constant SELLER_TAKER_FEE = 30; // 0.3%
    uint32 public constant BUYER_MAKER_FEE = 0;
    uint32 public constant BUYER_TAKER_FEE = 0;

    // ─────────────────────────────────────────────────────────────────────────────
    // V2 Contracts (loaded via vm.getCode)
    // ─────────────────────────────────────────────────────────────────────────────

    address public crossDexV2;
    address public routerV2;
    address public marketV2;
    address public pairV2;

    // ─────────────────────────────────────────────────────────────────────────────
    // V3 Implementation Contracts
    // ─────────────────────────────────────────────────────────────────────────────

    CrossDexImplV3 public crossDexV3Impl;
    CrossDexRouterV3 public routerV3Impl;
    MarketImplV3 public marketV3Impl;
    PairImplV3 public pairV3Impl;
    FeeControllerV2Compat public feeController;

    // ─────────────────────────────────────────────────────────────────────────────
    // Tokens
    // ─────────────────────────────────────────────────────────────────────────────

    T20 public QUOTE;
    T20 public BASE;

    uint256 public QUOTE_DECIMALS;
    uint256 public BASE_DECIMALS;

    // ─────────────────────────────────────────────────────────────────────────────
    // V2 Interface types (for casting)
    // ─────────────────────────────────────────────────────────────────────────────

    // V2 FeeConfig struct layout (same as IPairV2.FeeConfig)
    struct FeeConfigV2 {
        uint32 sellerMakerFeeBps;
        uint32 sellerTakerFeeBps;
        uint32 buyerMakerFeeBps;
        uint32 buyerTakerFeeBps;
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Setup
    // ─────────────────────────────────────────────────────────────────────────────

    function setUp() public {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");
        vm.label(USER1, "user1");
        vm.label(USER2, "user2");

        QUOTE_DECIMALS = 10 ** 6; // 6 decimals like USDT
        BASE_DECIMALS = 10 ** 18; // 18 decimals

        // Deploy tokens
        QUOTE = new T20("QUOTE", "QUOTE", 6);
        BASE = new T20("BASE", "BASE", 18);

        // Deploy V3 implementation contracts (for upgrade targets)
        vm.startPrank(OWNER);
        _deployV3Implementations();
        vm.stopPrank();
    }

    function _deployV3Implementations() internal {
        crossDexV3Impl = new CrossDexImplV3();
        routerV3Impl = new CrossDexRouterV3();
        marketV3Impl = new MarketImplV3();
        pairV3Impl = new PairImplV3();
        feeController = new FeeControllerV2Compat();
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // V2 Deployment Helpers
    // ─────────────────────────────────────────────────────────────────────────────

    function _deployV2CrossDexImpl() internal returns (address) {
        bytes memory bytecode = vm.getCode("legacy_v1_v2/out/CrossDexImplV2.sol/CrossDexImplV2.json");
        address deployed;
        assembly {
            deployed := create(0, add(bytecode, 0x20), mload(bytecode))
        }
        require(deployed != address(0), "CrossDexImplV2 deployment failed");
        return deployed;
    }

    function _deployV2RouterImpl() internal returns (address) {
        bytes memory bytecode = vm.getCode("legacy_v1_v2/out/CrossDexRouterV2.sol/CrossDexRouterV2.json");
        address deployed;
        assembly {
            deployed := create(0, add(bytecode, 0x20), mload(bytecode))
        }
        require(deployed != address(0), "CrossDexRouterV2 deployment failed");
        return deployed;
    }

    function _deployV2MarketImpl() internal returns (address) {
        bytes memory bytecode = vm.getCode("legacy_v1_v2/out/MarketImplV2.sol/MarketImplV2.json");
        address deployed;
        assembly {
            deployed := create(0, add(bytecode, 0x20), mload(bytecode))
        }
        require(deployed != address(0), "MarketImplV2 deployment failed");
        return deployed;
    }

    function _deployV2PairImpl() internal returns (address) {
        bytes memory bytecode = vm.getCode("legacy_v1_v2/out/PairImplV2.sol/PairImplV2.json");
        address deployed;
        assembly {
            deployed := create(0, add(bytecode, 0x20), mload(bytecode))
        }
        require(deployed != address(0), "PairImplV2 deployment failed");
        return deployed;
    }

    /// @notice Deploy complete V2 system
    function _deployV2System() internal {
        vm.startPrank(OWNER);

        // Deploy V2 implementation contracts
        address crossDexImpl = _deployV2CrossDexImpl();
        address routerImpl = _deployV2RouterImpl();
        address marketImpl = _deployV2MarketImpl();
        address pairImpl = _deployV2PairImpl();

        // Deploy CrossDex proxy
        ERC1967Proxy crossDexProxy = new ERC1967Proxy(crossDexImpl, hex"");
        crossDexV2 = address(crossDexProxy);

        // Initialize CrossDex V2
        (bool success,) = crossDexV2.call(
            abi.encodeWithSignature(
                "initialize(address,address,uint256,uint256,uint256,address,address,address)",
                OWNER,
                routerImpl,
                FIND_PREV_PRICE_COUNT,
                MAX_MATCH_COUNT,
                CANCEL_LIMIT,
                marketImpl,
                pairImpl,
                address(0) // tickSizeSetter
            )
        );
        require(success, "CrossDex V2 initialize failed");

        // Get router address
        (success,) = crossDexV2.staticcall(abi.encodeWithSignature("ROUTER()"));
        require(success, "Failed to get ROUTER");
        (, bytes memory routerData) = crossDexV2.staticcall(abi.encodeWithSignature("ROUTER()"));
        routerV2 = abi.decode(routerData, (address));

        // Create market with fees
        bytes memory feeData = abi.encode(SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);
        (success,) = crossDexV2.call(
            abi.encodeWithSignature(
                "createMarket(address,address,address,bytes,string)",
                OWNER,
                address(QUOTE),
                FEE_COLLECTOR,
                feeData,
                "TestMarket"
            )
        );
        require(success, "Market V2 creation failed");

        // Get market address (V2 uses quote => market mapping)
        (, bytes memory marketsData) = crossDexV2.staticcall(abi.encodeWithSignature("allMarkets()"));
        (address[] memory markets,) = abi.decode(marketsData, (address[], address[]));
        require(markets.length > 0, "No markets found");
        marketV2 = markets[0];

        // Create pair in market
        uint256 tickSize = QUOTE_DECIMALS / 100; // 0.01 QUOTE
        uint256 lotSize = BASE_DECIMALS / 1000; // 0.001 BASE
        bytes memory pairFeeData = abi.encode(SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

        (success,) = marketV2.call(
            abi.encodeWithSignature(
                "createPair(address,uint256,uint256,bytes)", address(BASE), tickSize, lotSize, pairFeeData
            )
        );
        require(success, "Pair V2 creation failed");

        // Get pair address
        (, bytes memory pairData) = marketV2.staticcall(abi.encodeWithSignature("baseToPair(address)", address(BASE)));
        pairV2 = abi.decode(pairData, (address));

        vm.stopPrank();

        // Setup token balances and approvals
        _setupTokensAndApprovals();
    }

    function _setupTokensAndApprovals() internal {
        // Transfer tokens to users
        QUOTE.transfer(USER1, 50000 * QUOTE_DECIMALS);
        QUOTE.transfer(USER2, 50000 * QUOTE_DECIMALS);
        BASE.transfer(USER1, 50000 * BASE_DECIMALS);
        BASE.transfer(USER2, 50000 * BASE_DECIMALS);

        // Approvals
        vm.prank(USER1);
        QUOTE.approve(routerV2, type(uint256).max);
        vm.prank(USER1);
        BASE.approve(routerV2, type(uint256).max);

        vm.prank(USER2);
        QUOTE.approve(routerV2, type(uint256).max);
        vm.prank(USER2);
        BASE.approve(routerV2, type(uint256).max);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Helper Functions
    // ─────────────────────────────────────────────────────────────────────────────

    function _toBase(uint256 x) internal view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) internal view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }

    function _submitSellOrderV2(address user, uint256 price, uint256 amount) internal returns (uint256 orderId) {
        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        vm.prank(user);
        (bool success, bytes memory returnData) = routerV2.call(
            abi.encodeWithSignature(
                "submitSellLimit(address,uint256,uint256,uint8,uint256[2],uint256)",
                pairV2,
                price,
                amount,
                0, // GOOD_TILL_CANCEL
                adjacent,
                0 // maxMatchCount (0 = use default)
            )
        );
        require(success, "submitSellLimit failed");
        orderId = abi.decode(returnData, (uint256));
    }

    function _submitBuyOrderV2(address user, uint256 price, uint256 amount) internal returns (uint256 orderId) {
        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        vm.prank(user);
        (bool success, bytes memory returnData) = routerV2.call(
            abi.encodeWithSignature(
                "submitBuyLimit(address,uint256,uint256,uint8,uint256[2],uint256)",
                pairV2,
                price,
                amount,
                0, // GOOD_TILL_CANCEL
                adjacent,
                0 // maxMatchCount (0 = use default)
            )
        );
        require(success, "submitBuyLimit failed");
        orderId = abi.decode(returnData, (uint256));
    }

    function _getV2PairFeeConfig() internal view returns (FeeConfigV2 memory config) {
        (, bytes memory data) = pairV2.staticcall(abi.encodeWithSignature("feeConfig()"));
        (config.sellerMakerFeeBps, config.sellerTakerFeeBps, config.buyerMakerFeeBps, config.buyerTakerFeeBps) =
            abi.decode(data, (uint32, uint32, uint32, uint32));
    }

    function _getPairBaseReserve() internal view returns (uint256) {
        (, bytes memory data) = pairV2.staticcall(abi.encodeWithSignature("baseReserve()"));
        return abi.decode(data, (uint256));
    }

    function _getPairQuoteReserve() internal view returns (uint256) {
        (, bytes memory data) = pairV2.staticcall(abi.encodeWithSignature("quoteReserve()"));
        return abi.decode(data, (uint256));
    }

    function _getPairMatchedPrice() internal view returns (uint256) {
        (, bytes memory data) = pairV2.staticcall(abi.encodeWithSignature("matchedPrice()"));
        return abi.decode(data, (uint256));
    }

    function _upgradePairToV3() internal {
        vm.prank(OWNER);
        UUPSUpgradeable(pairV2).upgradeToAndCall(address(pairV3Impl), hex"");
    }

    function _upgradeCrossDexToV3() internal {
        vm.prank(OWNER);
        UUPSUpgradeable(crossDexV2).upgradeToAndCall(address(crossDexV3Impl), hex"");
    }

    function _upgradeMarketToV3() internal {
        vm.prank(OWNER);
        UUPSUpgradeable(marketV2).upgradeToAndCall(address(marketV3Impl), hex"");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Test: CrossDex V2 → V3 Upgrade
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice CrossDex V2 → V3 upgrade is LOW risk
    /// @dev New variable `_allowedFeeControllers` is added after `tickSizeSetter`
    ///      The __gap size remains the same, so it occupies V2's gap space
    function test_crossDex_v2_to_v3_upgrade() public {
        _deployV2System();

        // Record V2 state before upgrade
        (, bytes memory routerData) = crossDexV2.staticcall(abi.encodeWithSignature("ROUTER()"));
        address routerBefore = abi.decode(routerData, (address));

        (, bytes memory marketImplData) = crossDexV2.staticcall(abi.encodeWithSignature("marketImpl()"));
        address marketImplBefore = abi.decode(marketImplData, (address));

        (, bytes memory pairImplData) = crossDexV2.staticcall(abi.encodeWithSignature("pairImpl()"));
        address pairImplBefore = abi.decode(pairImplData, (address));

        // Upgrade CrossDex to V3
        _upgradeCrossDexToV3();

        // Verify state is preserved
        CrossDexImplV3 crossDexV3 = CrossDexImplV3(crossDexV2);
        assertEq(address(crossDexV3.ROUTER()), routerBefore, "ROUTER should be preserved");
        assertEq(crossDexV3.marketImpl(), marketImplBefore, "marketImpl should be preserved");
        assertEq(crossDexV3.pairImpl(), pairImplBefore, "pairImpl should be preserved");

        // V3-specific: _allowedFeeControllers should be empty after upgrade
        // This needs to be initialized after upgrade
        vm.prank(OWNER);
        crossDexV3.setFeeControllerAllow(address(feeController), true);

        // Verify fee controller is now allowed
        // This would revert if not allowed
        vm.prank(OWNER);
        crossDexV3.checkFeeControllerAllowed(address(feeController));
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Test: Pair V2 → V3 Storage Collision
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Demonstrates storage collision when upgrading Pair V2 → V3 without mitigation
    /// @dev V2's `FeeConfig feeConfig` (128 bits) occupies the same slot as V3's `IFeeController feeController` (160 bits)
    ///      After upgrade without setFeeController, the feeController will read corrupted data
    function test_pair_v2_to_v3_storage_collision() public {
        _deployV2System();

        // Submit some orders to create active state
        uint256 price = _toQuote(100); // 100 QUOTE per BASE
        uint256 amount = _toBase(10); // 10 BASE
        _submitSellOrderV2(USER1, price, amount);

        // Record V2 state
        uint256 baseReserveBefore = _getPairBaseReserve();
        FeeConfigV2 memory feeConfigBefore = _getV2PairFeeConfig();

        // Verify V2 fee config is set
        assertEq(feeConfigBefore.sellerMakerFeeBps, SELLER_MAKER_FEE, "V2 sellerMakerFeeBps");
        assertEq(feeConfigBefore.sellerTakerFeeBps, SELLER_TAKER_FEE, "V2 sellerTakerFeeBps");

        // Read raw storage slot for feeConfig (slot index varies by contract layout)
        // FeeConfig is after _accountReserves mapping
        // In V2: slot 100 is feeConfig (FeeConfig struct)
        // In V3: slot 100 is feeController (address)
        bytes32 rawFeeConfigSlot = vm.load(pairV2, bytes32(uint256(100)));
        console.log("Raw feeConfig slot before upgrade:");
        console.logBytes32(rawFeeConfigSlot);

        // Upgrade Pair to V3 WITHOUT calling setFeeController
        _upgradePairToV3();

        // After upgrade, the same slot is now read as feeController address
        PairImplV3 pairV3 = PairImplV3(pairV2);

        // The feeController address will be corrupted (V2's FeeConfig bytes interpreted as address)
        address corruptedFeeController = address(pairV3.feeController());
        console.log("Corrupted feeController address after upgrade:", corruptedFeeController);

        // The corrupted address is derived from the FeeConfig struct bytes
        // FeeConfig = (sellerMaker, sellerTaker, buyerMaker, buyerTaker) = (20, 30, 0, 0)
        // As bytes: 0x00000014 00000001e 00000000 00000000 (in little-endian order in slot)
        // This will NOT be a valid contract address

        // Verify that calling feeController-related functions will fail
        // calcBuyVolumeWithFee requires delegatecall to feeController
        // This should revert because the corrupted address is not a valid contract
        uint256 testVolume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.expectRevert();
        pairV3.calcBuyVolumeWithFee(testVolume);

        // Verify reserves are preserved (non-mapping storage is safe)
        assertEq(pairV3.baseReserve(), baseReserveBefore, "baseReserve should be preserved");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Test: Pair V2 → V3 Upgrade with Mitigation
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Demonstrates correct upgrade path with immediate setFeeController call
    function test_pair_v2_to_v3_upgrade_with_mitigation() public {
        _deployV2System();

        // Submit orders to create active state BEFORE any upgrades
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        _submitSellOrderV2(USER1, price, amount);
        _submitBuyOrderV2(USER2, price / 2, amount); // Lower price, won't match

        // Record state before upgrade
        uint256 baseReserveBefore = _getPairBaseReserve();
        uint256 quoteReserveBefore = _getPairQuoteReserve();

        // Step 1: Upgrade CrossDex and allow fee controller
        _upgradeCrossDexToV3();
        vm.prank(OWNER);
        CrossDexImplV3(crossDexV2).setFeeControllerAllow(address(feeController), true);

        // Step 2: Upgrade Market (set feeController on Market only, not propagating to pairs yet)
        _upgradeMarketToV3();
        // Set market's feeController without propagating to V2 pairs (startIndex=0, endIndex=0)
        vm.prank(OWNER);
        MarketImplV3(marketV2)
            .setFeeController(
                0,
                0, // endIndex=0 means don't touch any pairs
                false,
                address(feeController),
                abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE)
            );

        // Step 3: Upgrade Pair to V3
        _upgradePairToV3();
        PairImplV3 pairV3 = PairImplV3(pairV2);

        // Step 4: CRITICAL - Immediately call setFeeController after Pair upgrade
        bytes memory feeControllerInitData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);
        vm.prank(OWNER);
        pairV3.setFeeController(address(feeController), feeControllerInitData);

        // Verify feeController is set correctly
        assertEq(address(pairV3.feeController()), address(feeController), "feeController should be set");

        // Verify state preservation
        assertEq(pairV3.baseReserve(), baseReserveBefore, "baseReserve preserved");
        assertEq(pairV3.quoteReserve(), quoteReserveBefore, "quoteReserve preserved");

        // Verify fee calculation works - this should NOT revert now
        uint256 expectedVolume = Math.mulDiv(price, amount, BASE_DECIMALS);
        uint256 volumeWithFee = pairV3.calcBuyVolumeWithFee(expectedVolume);
        uint256 expectedFee = Math.mulDiv(expectedVolume, BUYER_TAKER_FEE, BPS_DENOMINATOR);
        assertEq(volumeWithFee, expectedVolume + expectedFee, "Fee calculation should work");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Test: Market V2 → V3 NOT SUPPORTED
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Demonstrates that Market V2 → V3 direct upgrade is NOT supported
    /// @dev V2 has `_feeConfig` after `_allPairs`, V3 removes it
    ///      V2's `feeCollector` becomes V3's `feeController` (same slot, different purpose)
    ///      Direct upgrade will leave orphaned data and non-functional feeController
    function test_market_v2_to_v3_NOT_SUPPORTED() public {
        _deployV2System();

        // First upgrade CrossDex and allow fee controller
        _upgradeCrossDexToV3();
        vm.prank(OWNER);
        CrossDexImplV3(crossDexV2).setFeeControllerAllow(address(feeController), true);

        // Record V2 Market state
        (, bytes memory feeCollectorData) = marketV2.staticcall(abi.encodeWithSignature("feeCollector()"));
        address feeCollectorBefore = abi.decode(feeCollectorData, (address));
        assertEq(feeCollectorBefore, FEE_COLLECTOR, "V2 feeCollector should be set");

        // Upgrade Market to V3
        _upgradeMarketToV3();
        MarketImplV3 marketV3 = MarketImplV3(marketV2);

        // In V3, the same slot is now read as feeController
        // Since V2's feeCollector was FEE_COLLECTOR address, V3's feeController will be the same
        // BUT FEE_COLLECTOR is an EOA, not a valid FeeController contract!
        address upgradedFeeController = marketV3.feeController();
        assertEq(upgradedFeeController, FEE_COLLECTOR, "feeController reads V2's feeCollector value");

        // The fee controller is not allowed in CrossDex (and it's not even a valid contract)
        vm.expectRevert();
        marketV3.checkFeeControllerAllowed(upgradedFeeController);

        // MITIGATION: Must call setFeeController immediately after upgrade
        vm.prank(OWNER);
        marketV3.setFeeController(
            0,
            0,
            true,
            address(feeController),
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE)
        );

        // Now feeController is correctly set
        assertEq(marketV3.feeController(), address(feeController), "feeController should be set after mitigation");

        // V3's getFeeConfig() no longer exists - functionality is now in FeeController
        // This is why direct upgrade without understanding the architecture change is dangerous
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Test: Recommended Migration Path
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Full migration path from V2 to V3
    /// @dev 1. Upgrade CrossDex → allow FeeController
    ///      2. Upgrade Market → setFeeController
    ///      3. Upgrade each Pair → setFeeController
    ///      4. Verify full system functionality
    function test_recommended_migration_path() public {
        _deployV2System();

        // Create active trading state
        uint256 price = _toQuote(100);
        uint256 sellOrderId = _submitSellOrderV2(USER1, price, _toBase(10));
        uint256 buyOrderId = _submitBuyOrderV2(USER2, price / 2, _toBase(5)); // Won't match

        // Record pre-migration state
        uint256 baseReserveBefore = _getPairBaseReserve();
        uint256 quoteReserveBefore = _getPairQuoteReserve();

        console.log("=== Pre-Migration State ===");
        console.log("baseReserve:", baseReserveBefore);
        console.log("quoteReserve:", quoteReserveBefore);

        // Step 1-3: Perform upgrades
        _performV2ToV3Upgrades();

        // Step 4: Verify State Preservation
        console.log("\n=== Step 4: Verify State ===");
        PairImplV3 pairV3 = PairImplV3(pairV2);
        assertEq(pairV3.baseReserve(), baseReserveBefore, "baseReserve preserved");
        assertEq(pairV3.quoteReserve(), quoteReserveBefore, "quoteReserve preserved");

        // Verify orders are still intact
        assertEq(pairV3.orderById(sellOrderId).amount, _toBase(10), "Sell order amount preserved");
        assertEq(pairV3.orderById(buyOrderId).amount, _toBase(5), "Buy order amount preserved");

        // Step 5: Test post-migration trading
        _testPostMigrationTrading(price, sellOrderId, buyOrderId);

        console.log("\n=== Migration Complete ===");
    }

    function _performV2ToV3Upgrades() internal {
        // Step 1: Upgrade CrossDex V2 → V3
        console.log("\n=== Step 1: Upgrade CrossDex ===");
        _upgradeCrossDexToV3();
        vm.prank(OWNER);
        CrossDexImplV3(crossDexV2).setFeeControllerAllow(address(feeController), true);
        console.log("CrossDex upgraded and FeeController allowed");

        // Step 1b: Upgrade Router V2 → V3
        console.log("\n=== Step 1b: Upgrade Router ===");
        vm.prank(OWNER);
        UUPSUpgradeable(routerV2).upgradeToAndCall(address(routerV3Impl), hex"");
        console.log("Router upgraded to V3");

        // Step 2: Upgrade Market V2 → V3 (without propagating to V2 pairs)
        console.log("\n=== Step 2: Upgrade Market ===");
        _upgradeMarketToV3();
        bytes memory feeInitData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);
        vm.prank(OWNER);
        // endIndex=0 means don't propagate to any pairs (they're still V2)
        MarketImplV3(marketV2).setFeeController(0, 0, false, address(feeController), feeInitData);
        console.log("Market upgraded and FeeController set");

        // Step 3: Upgrade Pair V2 → V3
        console.log("\n=== Step 3: Upgrade Pair ===");
        _upgradePairToV3();
        vm.prank(OWNER);
        PairImplV3(pairV2).setFeeController(address(feeController), feeInitData);
        console.log("Pair upgraded and FeeController set");
    }

    function _testPostMigrationTrading(uint256 price, uint256 sellOrderId, uint256 buyOrderId) internal {
        console.log("\n=== Step 5: Test Post-Migration Trading ===");

        // Get V3 contracts
        CrossDexRouterV3 routerV3 = CrossDexRouterV3(CrossDexImplV3(crossDexV2).ROUTER());
        PairImplV3 pairV3 = PairImplV3(pairV2);

        // Approve V3 router
        _approveRouterV3(routerV3);

        // Cancel existing orders (they use V2 fee structure)
        _cancelOrdersV3(routerV3, pairV3, sellOrderId, buyOrderId);

        // Submit new V3 orders
        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        vm.prank(USER1);
        routerV3.submitSellLimit(
            address(pairV3), price, _toBase(5), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, adjacent, 0
        );

        vm.prank(USER2);
        routerV3.submitBuyLimit(
            address(pairV3), price, _toBase(3), IPairV3.LimitConstraints.GOOD_TILL_CANCEL, adjacent, 0
        );

        // Check matched price was updated (3 BASE matched at price 100)
        assertEq(pairV3.matchedPrice(), price, "Matched price should be updated");
    }

    function _approveRouterV3(CrossDexRouterV3 routerV3) internal {
        vm.prank(USER1);
        QUOTE.approve(address(routerV3), type(uint256).max);
        vm.prank(USER1);
        BASE.approve(address(routerV3), type(uint256).max);
        vm.prank(USER2);
        QUOTE.approve(address(routerV3), type(uint256).max);
        vm.prank(USER2);
        BASE.approve(address(routerV3), type(uint256).max);
    }

    function _cancelOrdersV3(CrossDexRouterV3 routerV3, PairImplV3 pairV3, uint256 sellOrderId, uint256 buyOrderId)
        internal
    {
        uint256[] memory orderIds = new uint256[](1);
        orderIds[0] = sellOrderId;
        vm.prank(USER1);
        routerV3.cancelOrder(address(pairV3), orderIds);

        orderIds[0] = buyOrderId;
        vm.prank(USER2);
        routerV3.cancelOrder(address(pairV3), orderIds);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Test: Storage Slot Verification
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Verify exact storage slot positions for Pair feeConfig/feeController
    function test_pair_storage_slot_analysis() public {
        _deployV2System();

        // Set specific fee values to identify in storage
        // V2 FeeConfig: (20, 30, 0, 0) = 0x...0000001e00000014 (packed)

        // The feeConfig slot position in PairImplV2:
        // - MARKET: slot 0 (address)
        // - ROUTER: slot 1 (address)
        // - BASE: slot 2 (IERC20)
        // - QUOTE: slot 3 (IERC20)
        // - DENOMINATOR: slot 4 (uint256)
        // - baseReserve: slot 5 (uint256)
        // - quoteReserve: slot 6 (uint256)
        // - matchedPrice: slot 7 (uint256)
        // - matchedAt: slot 8 (uint256)
        // - tickSize: slot 9 (uint256)
        // - lotSize: slot 10 (uint256)
        // - minTradeVolume: slot 11 (uint256)
        // - _orderIdCounter: slot 12 (uint256)
        // - _prices: slot 13-14 (List.U256[2])
        // - _sellOrders: slot 15 (mapping)
        // - _buyOrders: slot 16 (mapping)
        // - _allOrders: slot 17 (mapping)
        // - _accountReserves: slot 18 (mapping)
        // - feeConfig: slot 19 (FeeConfig struct - 4 uint32 = 1 slot)

        // Note: The actual slot number depends on inherited contract layout
        // PausableUpgradeable adds slots before our variables

        // Read slot 19 (approximate - may need adjustment based on inheritance)
        bytes32 slot19 = vm.load(pairV2, bytes32(uint256(19)));
        console.log("Slot 19:");
        console.logBytes32(slot19);

        // Let's find the actual feeConfig slot by searching
        for (uint256 i = 0; i < 150; i++) {
            bytes32 slotData = vm.load(pairV2, bytes32(i));
            // Look for our fee values: 20 (0x14) and 30 (0x1e)
            if (uint32(uint256(slotData)) == 20 && uint32(uint256(slotData) >> 32) == 30) {
                console.log("Found feeConfig at slot:", i);
                console.logBytes32(slotData);
                break;
            }
        }
    }
}
