// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {Test, console} from "forge-std/Test.sol";

import {CrossDexImpl} from "../src/CrossDexImpl.sol";
import {CrossDexImplV2} from "../src/CrossDexImplV2.sol";
import {CrossDexRouter} from "../src/CrossDexRouter.sol";
import {MarketImpl} from "../src/MarketImpl.sol";
import {MarketImplV2} from "../src/MarketImplV2.sol";
import {PairImpl} from "../src/PairImpl.sol";
import {PairImplV2} from "../src/PairImplV2.sol";

import {ICrossDex} from "../src/interfaces/ICrossDex.sol";

import {T20} from "./mock/T20.sol";

/**
 * @title CrossDexUpgradeTest
 * @notice V1 (CrossDexImpl) → V2 (CrossDexImplV2) 업그레이드 테스트
 * @dev V1에서 마켓 생성 후 V2로 업그레이드하면서 reinitialize 함수가
 *      _allMarkets 매핑을 quote=>market에서 market=>quote로 변환하는지 검증
 */
contract CrossDexUpgradeTest is Test {
    address public constant OWNER = address(0x1);
    address public constant FEE_COLLECTOR = address(0x2);

    CrossDexImpl public crossDexV1;
    CrossDexImplV2 public crossDexV2;
    ERC1967Proxy public proxy;

    CrossDexRouter public router;
    MarketImpl public marketImpl;
    PairImpl public pairImpl;

    T20 public quoteToken1;
    T20 public quoteToken2;
    T20 public quoteToken3;

    uint256 public constant FIND_PREV_PRICE_COUNT = type(uint256).max;
    uint256 public constant MAX_MATCH_COUNT = type(uint256).max;
    uint256 public constant CANCEL_LIMIT = type(uint256).max;
    uint256 public constant FEE_BPS = 30; // 0.3%

    address public market1;
    address public market2;
    address public market3;

    function setUp() external {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");

        vm.startPrank(OWNER);

        // Deploy implementation contracts
        address routerImpl = address(new CrossDexRouter());
        marketImpl = new MarketImpl();
        pairImpl = new PairImpl();

        // Deploy V1 implementation
        address crossDexImplV1 = address(new CrossDexImpl());

        // Deploy proxy pointing to V1
        proxy = new ERC1967Proxy(crossDexImplV1, hex"");
        crossDexV1 = CrossDexImpl(address(proxy));

        // Initialize V1
        crossDexV1.initialize(
            OWNER,
            routerImpl,
            FIND_PREV_PRICE_COUNT,
            MAX_MATCH_COUNT,
            CANCEL_LIMIT,
            address(marketImpl),
            address(pairImpl)
        );

        router = CrossDexRouter(payable(crossDexV1.ROUTER()));

        // Deploy quote tokens
        quoteToken1 = new T20("Quote Token 1", "QUOTE1", 18);
        quoteToken2 = new T20("Quote Token 2", "QUOTE2", 6);
        quoteToken3 = new T20("Quote Token 3", "QUOTE3", 18);

        vm.label(address(quoteToken1), "QUOTE1");
        vm.label(address(quoteToken2), "QUOTE2");
        vm.label(address(quoteToken3), "QUOTE3");

        vm.stopPrank();
    }

    /**
     * @notice V1에서 마켓 생성 후 V2로 업그레이드하여 reinitialize 테스트
     */
    function test_upgrade_v1_to_v2_with_reinitialize() external {
        vm.startPrank(OWNER);

        // ===== V1: 3개의 마켓 생성 =====
        console.log("=== V1: Creating markets ===");
        market1 = crossDexV1.createMarket(OWNER, address(quoteToken1), FEE_COLLECTOR, FEE_BPS);
        market2 = crossDexV1.createMarket(OWNER, address(quoteToken2), FEE_COLLECTOR, FEE_BPS);
        market3 = crossDexV1.createMarket(OWNER, address(quoteToken3), FEE_COLLECTOR, FEE_BPS);

        vm.label(market1, "market1");
        vm.label(market2, "market2");
        vm.label(market3, "market3");

        console.log("Market1 (QUOTE1):", market1);
        console.log("Market2 (QUOTE2):", market2);
        console.log("Market3 (QUOTE3):", market3);

        // V1에서 quoteToMarket 함수로 조회 가능
        assertEq(crossDexV1.quoteToMarket(address(quoteToken1)), market1, "V1: quote1 should map to market1");
        assertEq(crossDexV1.quoteToMarket(address(quoteToken2)), market2, "V1: quote2 should map to market2");
        assertEq(crossDexV1.quoteToMarket(address(quoteToken3)), market3, "V1: quote3 should map to market3");

        // V1에서 allMarkets 조회
        (address[] memory quotesV1, address[] memory marketsV1) = crossDexV1.allMarkets();
        assertEq(quotesV1.length, 3, "V1: should have 3 markets");
        assertEq(marketsV1.length, 3, "V1: should have 3 markets");

        console.log("\n=== V1 allMarkets (quote => market) ===");
        for (uint256 i = 0; i < quotesV1.length; i++) {
            console.log("  Quote:", quotesV1[i], "-> Market:", marketsV1[i]);
        }

        // ===== V2로 업그레이드 =====
        console.log("\n=== Upgrading to V2 ===");
        address crossDexImplV2Addr = address(new CrossDexImplV2());
        crossDexV1.upgradeToAndCall(crossDexImplV2Addr, hex"");

        // Proxy를 V2로 캐스팅
        crossDexV2 = CrossDexImplV2(address(proxy));

        // ===== reinitialize 호출 전: isMarket은 false (매핑이 quote=>market이므로) =====
        console.log("\n=== Before reinitialize ===");
        assertFalse(crossDexV2.isMarket(market1), "Before reinit: market1 should not be recognized");
        assertFalse(crossDexV2.isMarket(market2), "Before reinit: market2 should not be recognized");
        assertFalse(crossDexV2.isMarket(market3), "Before reinit: market3 should not be recognized");

        // allMarkets는 여전히 V1 구조 (quote => market)
        (address[] memory beforeQuotes, address[] memory beforeMarkets) = crossDexV2.allMarkets();
        console.log("Before reinit - allMarkets returns:");
        for (uint256 i = 0; i < beforeQuotes.length; i++) {
            console.log("  Index:", i);
            console.log("    Quote:", beforeQuotes[i]);
            console.log("    Market:", beforeMarkets[i]);
        }

        // ===== reinitialize 호출 (V2 implementation도 함께 설정) =====
        console.log("\n=== Calling reinitialize with V2 implementations ===");
        address marketImplV2 = address(new MarketImplV2());
        address pairImplV2 = address(new PairImplV2());

        console.log("MarketImplV2:", marketImplV2);
        console.log("PairImplV2:", pairImplV2);

        crossDexV2.reinitialize(marketImplV2, pairImplV2);

        // implementation이 업데이트되었는지 확인
        assertEq(crossDexV2.marketImpl(), marketImplV2, "MarketImpl should be updated");
        assertEq(crossDexV2.pairImpl(), pairImplV2, "PairImpl should be updated");

        // ===== reinitialize 후: 매핑이 market=>quote로 변경됨 =====
        console.log("\n=== After reinitialize ===");

        // isMarket으로 마켓 확인 가능
        assertTrue(crossDexV2.isMarket(market1), "After reinit: market1 should be recognized");
        assertTrue(crossDexV2.isMarket(market2), "After reinit: market2 should be recognized");
        assertTrue(crossDexV2.isMarket(market3), "After reinit: market3 should be recognized");

        // allMarkets의 구조가 market => quote로 변경됨
        (address[] memory marketsV2, address[] memory quotesV2) = crossDexV2.allMarkets();
        assertEq(marketsV2.length, 3, "V2: should have 3 markets");
        assertEq(quotesV2.length, 3, "V2: should have 3 markets");

        console.log("After reinit - allMarkets returns (market => quote):");
        for (uint256 i = 0; i < marketsV2.length; i++) {
            console.log("  Market:", marketsV2[i], "-> Quote:", quotesV2[i]);
        }

        // 각 마켓이 올바른 quote와 매핑되어 있는지 확인
        bool found1 = false;
        bool found2 = false;
        bool found3 = false;

        for (uint256 i = 0; i < marketsV2.length; i++) {
            if (marketsV2[i] == market1 && quotesV2[i] == address(quoteToken1)) found1 = true;
            else if (marketsV2[i] == market2 && quotesV2[i] == address(quoteToken2)) found2 = true;
            else if (marketsV2[i] == market3 && quotesV2[i] == address(quoteToken3)) found3 = true;
        }

        assertTrue(found1, "Market1 should be mapped to QUOTE1");
        assertTrue(found2, "Market2 should be mapped to QUOTE2");
        assertTrue(found3, "Market3 should be mapped to QUOTE3");

        // ===== V2에서 새로운 마켓 생성 (같은 quote로 여러 마켓) =====
        console.log("\n=== V2: Creating additional markets with same quote ===");

        // QUOTE1으로 두 번째 마켓 생성
        address market1_v2 = crossDexV2.createMarket(
            OWNER,
            address(quoteToken1),
            FEE_COLLECTOR,
            abi.encode(uint32(20), uint32(30), uint32(0), uint32(0)), // fee data
            "Second market for QUOTE1"
        );

        vm.label(market1_v2, "market1_v2");
        console.log("Market1_V2 (QUOTE1):", market1_v2);

        // 두 마켓 모두 인식되어야 함
        assertTrue(crossDexV2.isMarket(market1), "Market1 should still be recognized");
        assertTrue(crossDexV2.isMarket(market1_v2), "Market1_V2 should be recognized");

        // allMarkets에 4개의 마켓이 있어야 함
        (address[] memory finalMarkets, address[] memory finalQuotes) = crossDexV2.allMarkets();
        assertEq(finalMarkets.length, 4, "V2: should have 4 markets after adding one more");

        console.log("\n=== Final state: 4 markets ===");
        for (uint256 i = 0; i < finalMarkets.length; i++) {
            console.log("  Market:", finalMarkets[i], "-> Quote:", finalQuotes[i]);
        }

        vm.stopPrank();
    }

    /**
     * @notice reinitialize를 두 번 호출하면 실패해야 함
     */
    function test_reinitialize_cannot_be_called_twice() external {
        vm.startPrank(OWNER);

        // V1에서 마켓 생성
        market1 = crossDexV1.createMarket(OWNER, address(quoteToken1), FEE_COLLECTOR, FEE_BPS);

        // V2로 업그레이드
        address crossDexImplV2Addr = address(new CrossDexImplV2());
        crossDexV1.upgradeToAndCall(crossDexImplV2Addr, hex"");
        crossDexV2 = CrossDexImplV2(address(proxy));

        // 첫 번째 reinitialize 호출 성공
        address marketImplV2 = address(new MarketImplV2());
        address pairImplV2 = address(new PairImplV2());
        crossDexV2.reinitialize(marketImplV2, pairImplV2);

        // 두 번째 reinitialize 호출 시 revert
        vm.expectRevert();
        crossDexV2.reinitialize(marketImplV2, pairImplV2);

        vm.stopPrank();
    }

    /**
     * @notice V1에서 마켓이 없어도 V2로 업그레이드 및 reinitialize 가능
     */
    function test_upgrade_v1_to_v2_with_no_markets() external {
        vm.startPrank(OWNER);

        // V1에서 마켓을 생성하지 않음
        (address[] memory quotesV1, address[] memory marketsV1) = crossDexV1.allMarkets();
        assertEq(quotesV1.length, 0, "V1: should have 0 markets");

        // V2로 업그레이드
        address crossDexImplV2Addr = address(new CrossDexImplV2());
        crossDexV1.upgradeToAndCall(crossDexImplV2Addr, hex"");
        crossDexV2 = CrossDexImplV2(address(proxy));

        // reinitialize 호출 (빈 배열 처리)
        address marketImplV2 = address(new MarketImplV2());
        address pairImplV2 = address(new PairImplV2());
        crossDexV2.reinitialize(marketImplV2, pairImplV2);

        // 여전히 마켓이 없어야 함
        (address[] memory marketsV2, address[] memory quotesV2) = crossDexV2.allMarkets();
        assertEq(marketsV2.length, 0, "V2: should still have 0 markets");

        vm.stopPrank();
    }

    /**
     * @notice V1의 데이터 구조를 검증
     */
    function test_v1_data_structure() external {
        vm.startPrank(OWNER);

        // 마켓 생성
        market1 = crossDexV1.createMarket(OWNER, address(quoteToken1), FEE_COLLECTOR, FEE_BPS);
        market2 = crossDexV1.createMarket(OWNER, address(quoteToken2), FEE_COLLECTOR, FEE_BPS);

        // V1에서는 같은 quote로 두 번째 마켓을 생성할 수 없음
        // Create2.deploy가 실패하므로 일반 revert 체크
        vm.expectRevert();
        crossDexV1.createMarket(OWNER, address(quoteToken1), FEE_COLLECTOR, FEE_BPS);

        // quoteToMarket 함수로 조회
        assertEq(crossDexV1.quoteToMarket(address(quoteToken1)), market1);
        assertEq(crossDexV1.quoteToMarket(address(quoteToken2)), market2);

        // allMarkets 반환 값 확인: (quotes[], markets[])
        (address[] memory quotes, address[] memory markets) = crossDexV1.allMarkets();
        assertEq(quotes.length, 2);
        assertEq(markets.length, 2);

        // V1 구조: 첫 번째 배열이 quote, 두 번째 배열이 market
        assertTrue(
            (quotes[0] == address(quoteToken1) && markets[0] == market1)
                || (quotes[1] == address(quoteToken1) && markets[1] == market1),
            "QUOTE1 should map to market1"
        );

        vm.stopPrank();
    }

    /**
     * @notice V2의 새로운 데이터 구조를 검증
     */
    function test_v2_data_structure() external {
        vm.startPrank(OWNER);

        // V2로 직접 업그레이드 (마켓 없이)
        address crossDexImplV2Addr = address(new CrossDexImplV2());
        crossDexV1.upgradeToAndCall(crossDexImplV2Addr, hex"");
        crossDexV2 = CrossDexImplV2(address(proxy));

        // V2 implementation과 함께 reinitialize
        address marketImplV2 = address(new MarketImplV2());
        address pairImplV2 = address(new PairImplV2());
        crossDexV2.reinitialize(marketImplV2, pairImplV2);

        // V2에서는 같은 quote로 여러 마켓 생성 가능 (message로 구분)
        address market1a = crossDexV2.createMarket(
            OWNER,
            address(quoteToken1),
            FEE_COLLECTOR,
            abi.encode(uint32(20), uint32(30), uint32(0), uint32(0)),
            "Market A"
        );

        address market1b = crossDexV2.createMarket(
            OWNER,
            address(quoteToken1),
            FEE_COLLECTOR,
            abi.encode(uint32(10), uint32(20), uint32(0), uint32(0)),
            "Market B"
        );

        // 두 마켓 모두 같은 quote를 사용하지만 다른 주소
        assertNotEq(market1a, market1b, "Markets should have different addresses");

        // 두 마켓 모두 isMarket으로 확인 가능
        assertTrue(crossDexV2.isMarket(market1a));
        assertTrue(crossDexV2.isMarket(market1b));

        // allMarkets 반환 값 확인: (markets[], quotes[])
        (address[] memory markets, address[] memory quotes) = crossDexV2.allMarkets();
        assertEq(markets.length, 2);
        assertEq(quotes.length, 2);

        // V2 구조: 첫 번째 배열이 market, 두 번째 배열이 quote
        uint256 foundCount = 0;
        for (uint256 i = 0; i < markets.length; i++) {
            if (quotes[i] == address(quoteToken1)) {
                assertTrue(
                    markets[i] == market1a || markets[i] == market1b, "Market should be one of the created markets"
                );
                foundCount++;
            }
        }
        assertEq(foundCount, 2, "Both markets should use QUOTE1");

        vm.stopPrank();
    }

    /**
     * @notice reinitialize 시 zero address를 전달하면 실패해야 함
     */
    function test_reinitialize_revert_on_zero_address() external {
        vm.startPrank(OWNER);

        // V2로 업그레이드
        address crossDexImplV2Addr = address(new CrossDexImplV2());
        crossDexV1.upgradeToAndCall(crossDexImplV2Addr, hex"");
        crossDexV2 = CrossDexImplV2(address(proxy));

        address marketImplV2 = address(new MarketImplV2());
        address pairImplV2 = address(new PairImplV2());

        // marketImpl이 zero address이면 revert
        vm.expectRevert();
        crossDexV2.reinitialize(address(0), pairImplV2);

        // pairImpl이 zero address이면 revert
        vm.expectRevert();
        crossDexV2.reinitialize(marketImplV2, address(0));

        // 둘 다 zero address이면 revert
        vm.expectRevert();
        crossDexV2.reinitialize(address(0), address(0));

        vm.stopPrank();
    }

    /**
     * @notice reinitialize는 owner만 호출 가능
     */
    function test_reinitialize_only_owner() external {
        // V2로 업그레이드
        vm.prank(OWNER);
        address crossDexImplV2Addr = address(new CrossDexImplV2());
        vm.prank(OWNER);
        crossDexV1.upgradeToAndCall(crossDexImplV2Addr, hex"");
        crossDexV2 = CrossDexImplV2(address(proxy));

        address marketImplV2 = address(new MarketImplV2());
        address pairImplV2 = address(new PairImplV2());

        // owner가 아닌 계정으로 호출 시 revert
        address nonOwner = address(0x999);
        vm.prank(nonOwner);
        vm.expectRevert();
        crossDexV2.reinitialize(marketImplV2, pairImplV2);

        // owner는 성공
        vm.prank(OWNER);
        crossDexV2.reinitialize(marketImplV2, pairImplV2);

        // implementation이 업데이트되었는지 확인
        assertEq(crossDexV2.marketImpl(), marketImplV2, "MarketImpl should be updated");
        assertEq(crossDexV2.pairImpl(), pairImplV2, "PairImpl should be updated");
    }
}
