// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";
import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";

import {CrossDexImplV3} from "../src/CrossDexImplV3.sol";
import {CrossDexRouterV3} from "../src/CrossDexRouterV3.sol";
import {FeeControllerV4Dist} from "../src/FeeControllerV4Dist.sol";
import {MarketImplV3} from "../src/MarketImplV3.sol";
import {PairImplV3} from "../src/PairImplV3.sol";
import {BPS_DENOMINATOR} from "../src/interfaces/IFeeController.sol";
import {IPairV3} from "../src/interfaces/IPairV3.sol";
import {T20} from "./mock/T20.sol";

/// @title FeeControllerV4DistTest
/// @notice Tests for FeeControllerV4Dist: 4-way fees with N-way distribution to multiple recipients.
contract FeeControllerV4DistTest is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant RECIPIENT_A = address(bytes20("RECIPIENT_A"));
    address public constant RECIPIENT_B = address(bytes20("RECIPIENT_B"));
    address public constant RECIPIENT_C = address(bytes20("RECIPIENT_C"));
    address public constant USER1 = address(bytes20("USER1"));
    address public constant USER2 = address(bytes20("USER2"));

    CrossDexImplV3 public CROSS_DEX;
    CrossDexRouterV3 public ROUTER;
    IERC20 public QUOTE;
    IERC20 public BASE;
    MarketImplV3 public MARKET;
    PairImplV3 public PAIR;
    FeeControllerV4Dist public FEE_CONTROLLER;

    uint256 public QUOTE_DECIMALS = 1e18;
    uint256 public BASE_DECIMALS = 1e18;
    uint32 public SELLER_MAKER_FEE = 20;
    uint32 public SELLER_TAKER_FEE = 30;
    uint32 public BUYER_MAKER_FEE = 10;
    uint32 public BUYER_TAKER_FEE = 10;

    uint256[2] internal _searchPrices;

    function setUp() external {
        vm.label(OWNER, "owner");
        vm.label(RECIPIENT_A, "recipientA");
        vm.label(RECIPIENT_B, "recipientB");
        vm.label(RECIPIENT_C, "recipientC");
        vm.label(USER1, "user1");
        vm.label(USER2, "user2");

        vm.startPrank(OWNER);

        FEE_CONTROLLER = new FeeControllerV4Dist();

        address routerImpl = address(new CrossDexRouterV3());
        address marketImpl = address(new MarketImplV3());
        address pairImpl = address(new PairImplV3());
        address crossDexImpl = address(new CrossDexImplV3());
        ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
        CROSS_DEX = CrossDexImplV3(address(proxy));
        CROSS_DEX.initialize(
            OWNER, routerImpl, type(uint256).max, type(uint256).max, type(uint256).max, marketImpl, pairImpl, address(0)
        );
        CROSS_DEX.setFeeControllerAllow(address(FEE_CONTROLLER), true);

        ROUTER = CrossDexRouterV3(CROSS_DEX.ROUTER());

        QUOTE = new T20("QUOTE", "QUOTE", 18);
        BASE = new T20("BASE", "BASE", 18);

        address[] memory recipients = new address[](3);
        recipients[0] = RECIPIENT_A;
        recipients[1] = RECIPIENT_B;
        recipients[2] = RECIPIENT_C;
        uint32[] memory ratios = new uint32[](3);
        ratios[0] = 3000;
        ratios[1] = 3000;
        ratios[2] = 4000;
        bytes32[] memory labels = new bytes32[](3);
        labels[0] = bytes32("CREATOR");
        labels[1] = bytes32("PLATFORM");
        labels[2] = bytes32("SYSTEM");

        bytes memory feeControllerInitData = abi.encode(
            SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE, recipients, ratios, labels
        );

        address market = CROSS_DEX.createMarket(OWNER, address(QUOTE), address(FEE_CONTROLLER), "");
        MARKET = MarketImplV3(market);
        address pair = MARKET.createPair(address(BASE), 1e16, 1e12, feeControllerInitData);
        PAIR = PairImplV3(pair);

        QUOTE.transfer(USER1, 50000e18);
        QUOTE.transfer(USER2, 50000e18);
        BASE.transfer(USER1, 50000e18);
        BASE.transfer(USER2, 50000e18);

        vm.stopPrank();

        vm.prank(USER1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER1);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER2);
        BASE.approve(address(ROUTER), type(uint256).max);
    }

    function _toBase(uint256 x) internal view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) internal view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }

    function _toTradeVolume(uint256 price, uint256 amount) internal view returns (uint256) {
        return Math.mulDiv(price, amount, BASE_DECIMALS);
    }

    function _calcFee(uint256 amount, uint32 bps) internal pure returns (uint256) {
        return Math.mulDiv(amount, bps, BPS_DENOMINATOR);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Initialize validation
    // ─────────────────────────────────────────────────────────────────────────────

    function test_initialize_revert_ratios_sum_not_10000() external {
        address[] memory recipients = new address[](3);
        recipients[0] = RECIPIENT_A;
        recipients[1] = RECIPIENT_B;
        recipients[2] = RECIPIENT_C;
        uint32[] memory ratios = new uint32[](3);
        ratios[0] = 3000;
        ratios[1] = 3000;
        ratios[2] = 3999; // sum = 9999
        bytes32[] memory labels = new bytes32[](3);
        labels[0] = bytes32("A");
        labels[1] = bytes32("B");
        labels[2] = bytes32("C");

        bytes memory badInitData = abi.encode(
            SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE, recipients, ratios, labels
        );

        vm.expectRevert(
            abi.encodeWithSelector(FeeControllerV4Dist.FeeControllerV4DistRatiosBpsSumNot10000.selector, 9999)
        );
        vm.prank(OWNER);
        PAIR.setFeeController(address(FEE_CONTROLLER), badInitData);
    }

    function test_initialize_revert_zero_recipient() external {
        address[] memory recipients = new address[](2);
        recipients[0] = RECIPIENT_A;
        recipients[1] = address(0);
        uint32[] memory ratios = new uint32[](2);
        ratios[0] = 5000;
        ratios[1] = 5000;
        bytes32[] memory labels = new bytes32[](2);
        labels[0] = bytes32("A");
        labels[1] = bytes32("B");

        bytes memory badInitData = abi.encode(
            SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE, recipients, ratios, labels
        );

        vm.expectRevert(FeeControllerV4Dist.FeeControllerV4DistInvalidRecipients.selector);
        vm.prank(OWNER);
        PAIR.setFeeController(address(FEE_CONTROLLER), badInitData);
    }

    function test_initialize_revert_too_many_recipients() external {
        address[] memory recipients = new address[](11);
        uint32[] memory ratios = new uint32[](11);
        bytes32[] memory labels = new bytes32[](11);
        for (uint256 i = 0; i < 11; ++i) {
            recipients[i] = address(uint160(100 + i));
            ratios[i] = i < 10 ? 909 : uint32(10000 - 909 * 10);
            labels[i] = bytes32(uint256(i));
        }

        bytes memory badInitData = abi.encode(
            SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE, recipients, ratios, labels
        );

        vm.expectRevert(FeeControllerV4Dist.FeeControllerV4DistInvalidRecipients.selector);
        vm.prank(OWNER);
        PAIR.setFeeController(address(FEE_CONTROLLER), badInitData);
    }

    function test_get_effective_fees() external {
        (uint32 sMk, uint32 sTk, uint32 bMk, uint32 bTk) = PAIR.getEffectiveFees();
        assertEq(sMk, SELLER_MAKER_FEE);
        assertEq(sTk, SELLER_TAKER_FEE);
        assertEq(bMk, BUYER_MAKER_FEE);
        assertEq(bTk, BUYER_TAKER_FEE);
    }

    function test_get_config_id() external {
        (bytes32 configId,) = PAIR.getFeeControllerConfig();
        assertEq(configId, keccak256("FeeControllerV4Dist.v1"));
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // settleFees N-way distribution and dust
    // ─────────────────────────────────────────────────────────────────────────────

    function test_settle_fees_distributes_to_three_recipients() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 aBefore = QUOTE.balanceOf(RECIPIENT_A);
        uint256 bBefore = QUOTE.balanceOf(RECIPIENT_B);
        uint256 cBefore = QUOTE.balanceOf(RECIPIENT_C);

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        uint256 makerFee = _calcFee(quoteVolume, BUYER_MAKER_FEE);
        uint256 takerFee = _calcFee(quoteVolume, SELLER_TAKER_FEE);
        uint256 totalFee = makerFee + takerFee;
        uint256 expectedA = Math.mulDiv(totalFee, 3000, BPS_DENOMINATOR);
        uint256 expectedB = Math.mulDiv(totalFee, 3000, BPS_DENOMINATOR);
        uint256 expectedC = totalFee - expectedA - expectedB;

        assertEq(QUOTE.balanceOf(RECIPIENT_A) - aBefore, expectedA, "RECIPIENT_A");
        assertEq(QUOTE.balanceOf(RECIPIENT_B) - bBefore, expectedB, "RECIPIENT_B");
        assertEq(QUOTE.balanceOf(RECIPIENT_C) - cBefore, expectedC, "RECIPIENT_C (with remainder)");
    }

    function test_settle_fees_dust_to_last_recipient() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(7);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        uint256 makerFee = _calcFee(quoteVolume, BUYER_MAKER_FEE);
        uint256 takerFee = _calcFee(quoteVolume, SELLER_TAKER_FEE);
        uint256 totalFee = makerFee + takerFee;
        uint256 aGot = Math.mulDiv(totalFee, 3000, BPS_DENOMINATOR);
        uint256 bGot = Math.mulDiv(totalFee, 3000, BPS_DENOMINATOR);
        uint256 cExpected = totalFee - aGot - bGot;

        assertEq(QUOTE.balanceOf(RECIPIENT_A), aGot, "A exact share");
        assertEq(QUOTE.balanceOf(RECIPIENT_B), bGot, "B exact share");
        assertEq(QUOTE.balanceOf(RECIPIENT_C), cExpected, "C gets remainder (dust)");
        assertEq(aGot + bGot + cExpected, totalFee, "sum equals totalFee");
    }

    function test_fees_settled_event_includes_labels() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        vm.recordLogs();
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        Vm.Log[] memory entries = vm.getRecordedLogs();
        bytes32 expectedTopic =
            keccak256("FeeControllerV4DistFeesSettled(uint256,uint256,address[],uint256[],bytes32[])");
        bool found = false;
        for (uint256 i = 0; i < entries.length; ++i) {
            if (entries[i].topics.length >= 1 && entries[i].topics[0] == expectedTopic) {
                found = true;
                break;
            }
        }
        assertTrue(found, "FeeControllerV4DistFeesSettled event should be emitted");
    }

    function test_record_match_accumulates_maker_and_taker() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);

        vm.prank(USER2);
        ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        uint256 totalMakerFee = _calcFee(quoteVolume, SELLER_MAKER_FEE);
        uint256 totalTakerFee = _calcFee(quoteVolume, BUYER_TAKER_FEE);
        uint256 totalFee = totalMakerFee + totalTakerFee;

        uint256 aBefore = QUOTE.balanceOf(RECIPIENT_A);
        uint256 bBefore = QUOTE.balanceOf(RECIPIENT_B);
        uint256 cBefore = QUOTE.balanceOf(RECIPIENT_C);

        vm.prank(USER1);
        ROUTER.submitBuyMarket(address(PAIR), quoteVolume + totalTakerFee, 0);

        uint256 expectedA = Math.mulDiv(totalFee, 3000, BPS_DENOMINATOR);
        uint256 expectedB = Math.mulDiv(totalFee, 3000, BPS_DENOMINATOR);
        uint256 expectedC = totalFee - expectedA - expectedB;

        assertEq(QUOTE.balanceOf(RECIPIENT_A) - aBefore, expectedA, "A");
        assertEq(QUOTE.balanceOf(RECIPIENT_B) - bBefore, expectedB, "B");
        assertEq(QUOTE.balanceOf(RECIPIENT_C) - cBefore, expectedC, "C");
    }

    /// @notice Verifies: (1) fee payers pay exactly the configured BPS, (2) recipients receive by ratio,
    ///         (3) Pair holds no excess QUOTE/BASE after settlement (no dust left in contract).
    function test_fee_payer_pays_exact_bps_and_pair_has_no_excess_balance() external {
        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);
        uint256 expectedMakerFee = _calcFee(quoteVolume, BUYER_MAKER_FEE);
        uint256 expectedTakerFee = _calcFee(quoteVolume, SELLER_TAKER_FEE);
        uint256 expectedTotalFee = expectedMakerFee + expectedTakerFee;

        uint256 user1QuoteBefore = QUOTE.balanceOf(USER1);
        uint256 user1BaseBefore = BASE.balanceOf(USER1);
        uint256 user2QuoteBefore = QUOTE.balanceOf(USER2);
        uint256 user2BaseBefore = BASE.balanceOf(USER2);
        uint256 aBefore = QUOTE.balanceOf(RECIPIENT_A);
        uint256 bBefore = QUOTE.balanceOf(RECIPIENT_B);
        uint256 cBefore = QUOTE.balanceOf(RECIPIENT_C);

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);

        uint256 user1QuoteAfter = QUOTE.balanceOf(USER1);
        uint256 user1BaseAfter = BASE.balanceOf(USER1);
        uint256 user2QuoteAfter = QUOTE.balanceOf(USER2);
        uint256 user2BaseAfter = BASE.balanceOf(USER2);

        uint256 makerPaid = user1QuoteBefore - user1QuoteAfter - quoteVolume;
        assertEq(makerPaid, expectedMakerFee, "maker pays exactly buyerMakerFeeBps of quote volume");

        uint256 takerReceived = user2QuoteAfter - user2QuoteBefore;
        assertEq(takerReceived, quoteVolume - expectedTakerFee, "taker receives quote minus sellerTakerFeeBps");

        uint256 totalReceivedByRecipients = (QUOTE.balanceOf(RECIPIENT_A) - aBefore)
            + (QUOTE.balanceOf(RECIPIENT_B) - bBefore) + (QUOTE.balanceOf(RECIPIENT_C) - cBefore);
        assertEq(totalReceivedByRecipients, expectedTotalFee, "recipients receive total fee");

        assertEq(PAIR.quoteReserve(), 0, "quoteReserve zero after full match");
        assertEq(PAIR.baseReserve(), 0, "baseReserve zero after full match");
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "Pair holds no excess QUOTE after settlement");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "Pair holds no excess BASE after settlement");

        assertEq(user1BaseAfter - user1BaseBefore, amount, "buyer receives exact BASE");
        assertEq(user2BaseBefore - user2BaseAfter, amount, "seller sent exact BASE");
    }

    function test_single_recipient_10000_bps() external {
        address[] memory pairs = new address[](1);
        pairs[0] = address(PAIR);
        address[] memory oneRecipient = new address[](1);
        oneRecipient[0] = RECIPIENT_A;
        uint32[] memory oneRatio = new uint32[](1);
        oneRatio[0] = 10000;
        bytes32[] memory oneLabel = new bytes32[](1);
        oneLabel[0] = bytes32("ONLY");

        bytes memory initData = abi.encode(
            SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE, oneRecipient, oneRatio, oneLabel
        );

        vm.prank(OWNER);
        MARKET.setFeeController(pairs, address(FEE_CONTROLLER), initData);

        uint256 price = _toQuote(100);
        uint256 amount = _toBase(10);
        uint256 quoteVolume = _toTradeVolume(price, amount);
        uint256 makerFee = _calcFee(quoteVolume, BUYER_MAKER_FEE);
        uint256 takerFee = _calcFee(quoteVolume, SELLER_TAKER_FEE);
        uint256 expectedFee = makerFee + takerFee;

        vm.prank(USER1);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPairV3.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);

        uint256 aBefore = QUOTE.balanceOf(RECIPIENT_A);
        vm.prank(USER2);
        ROUTER.submitSellMarket(address(PAIR), amount, 0);
        assertEq(QUOTE.balanceOf(RECIPIENT_A) - aBefore, expectedFee, "single recipient gets full fee");
    }
}
