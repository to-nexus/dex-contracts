// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";

import {BPS_DENOMINATOR, IFeeController} from "./interfaces/IFeeController.sol";
import {IPairV3} from "./interfaces/IPairV3.sol";

/// @title FeeControllerV2Compat
/// @notice V2-compatible fee controller implementing 4 fee bps (seller/buyer × maker/taker).
///         Designed to be called via delegatecall from PairImplV3.
/// @dev All state is stored in Pair's storage via ERC-7201 namespaced slot.
contract FeeControllerV2Compat is IFeeController {
    using SafeERC20 for IERC20;
    using Math for uint256;

    /// @dev Reserved slot for taker fee bps direct storage access (future extensibility)
    /// keccak256(abi.encode(uint256(keccak256("crossdex.feecontroller.v2compat.takerfeebps")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant _takerFeeBpsSlot = 0x5813c8d80096621638b3ed51bf4250962c0ac17965191c985906fbb2f45de500;

    // ─────────────────────────────────────────────────────────────────────────────
    // ERC-7201 Namespaced Storage
    // ─────────────────────────────────────────────────────────────────────────────

    /// @custom:storage-location erc7201:crossdex.feecontroller.v2compat
    struct Layout {
        // --- Persistent config (set via initialize) ---
        address feeCollector;
        uint32 sellerMakerFeeBps;
        uint32 sellerTakerFeeBps;
        uint32 buyerMakerFeeBps;
        uint32 buyerTakerFeeBps;
        // Cached from Pair for gas savings
        IERC20 quote;
        uint256 denominator;
        // --- Per-submit accumulation (reset after settleFees) ---
        uint256 makerFeeAcc;
        uint256 takerFeeAcc;
    }

    // keccak256(abi.encode(uint256(keccak256("crossdex.feecontroller.v2compat")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant STORAGE_SLOT = 0x8a0c9d8ec1d9f8b3f4e5a6b7c8d9e0f1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c600;

    function _layout() private pure returns (Layout storage $) {
        assembly {
            $.slot := STORAGE_SLOT
        }
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Delegatecall enforcement
    // ─────────────────────────────────────────────────────────────────────────────

    /// @dev Store the original address at deployment for delegatecall check.
    address private immutable _SELF;

    constructor() {
        _SELF = address(this);
    }

    /// @dev Ensures the function is called via delegatecall (msg.sender != address(this) in original context).
    modifier onlyDelegateCall() {
        if (address(this) == _SELF) revert FeeControllerNotDelegateCall();
        _;
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // IFeeController Implementation
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Initialize or update fee configuration.
    /// @dev Called via delegatecall from Pair. Reads QUOTE/DENOMINATOR from Pair's storage directly.
    /// @param initData abi.encode(feeCollector, sellerMakerBps, sellerTakerBps, buyerMakerBps, buyerTakerBps)
    function initialize(address quote, uint256 denominator, bytes memory initData) external override onlyDelegateCall {
        (address _feeCollector, uint32 sMk, uint32 sTk, uint32 bMk, uint32 bTk) =
            abi.decode(initData, (address, uint32, uint32, uint32, uint32));

        // Validation
        if (_feeCollector == address(0)) revert FeeControllerInvalidFeeCollector();
        if (sMk >= BPS_DENOMINATOR || sTk >= BPS_DENOMINATOR) revert FeeControllerInvalidFeeBps();
        if (bMk >= BPS_DENOMINATOR || bTk >= BPS_DENOMINATOR) revert FeeControllerInvalidFeeBps();
        if (sTk < sMk) revert FeeControllerInvalidFeeStructure(sMk, sTk);
        if (bTk < bMk) revert FeeControllerInvalidFeeStructure(bMk, bTk);

        Layout storage $ = _layout();
        $.feeCollector = _feeCollector;
        $.sellerMakerFeeBps = sMk;
        $.sellerTakerFeeBps = sTk;
        $.buyerMakerFeeBps = bMk;
        $.buyerTakerFeeBps = bTk;
        // Reset accumulators on re-initialize
        $.makerFeeAcc = 0;
        $.takerFeeAcc = 0;
        if (address($.quote) == address(0)) {
            $.quote = IERC20(quote);
            $.denominator = denominator;
        } else {
            if (address($.quote) != address(quote)) revert FeeControllerInvalidPairConfig(quote, denominator);
        }
    }

    /// @notice Calculate total QUOTE volume including fee for a BUY order.
    /// @param isMaker true if calculating for maker (limit order on book), false for taker
    /// @param order The order to calculate for
    /// @return buyVolume The total volume including fee
    function calcBuyVolumeWithFeeByOrder(bool isMaker, IPairV3.Order memory order)
        external
        view
        override
        returns (uint256 buyVolume)
    {
        // Note: view function - onlyDelegateCall not needed since it doesn't modify state
        // and will read from caller's storage in delegatecall context anyway
        Layout storage $ = _layout();
        uint256 baseVolume = Math.mulDiv(order.price, order.amount, $.denominator);
        uint32 bps = isMaker ? $.buyerMakerFeeBps : $.buyerTakerFeeBps;
        return baseVolume + Math.mulDiv(baseVolume, bps, BPS_DENOMINATOR);
    }

    /// @notice Calculate total QUOTE volume including fee for a given base volume.
    /// @param isMaker true if calculating for maker, false for taker
    /// @param volume The base QUOTE volume (without fee)
    /// @return buyVolume The total volume including fee
    function calcBuyVolumeWithFee(bool isMaker, uint256 volume) external view override returns (uint256 buyVolume) {
        Layout storage $ = _layout();
        uint32 bps = isMaker ? $.buyerMakerFeeBps : $.buyerTakerFeeBps;
        return volume + Math.mulDiv(volume, bps, BPS_DENOMINATOR);
    }

    /// @notice Record a match and accumulate fees.
    /// @dev Called via delegatecall for each fill during matching.
    ///      - Maker fee uses maker.feeBps (set at order creation time for V2 compatibility)
    ///      - Taker fee uses current config based on taker.side
    /// @param taker The taker order (uses taker.side to determine taker fee bps)
    /// @param maker The maker order (uses maker.feeBps which was set at order creation)
    /// @param tradeQuoteAmount The trade volume in QUOTE (fee calculation base)
    /// @return makerFee The fee charged to the maker for this fill
    function recodeMatch(
        IPairV3.Order memory taker,
        IPairV3.Order memory maker,
        uint256, /* tradeAmount - unused in V2Compat, reserved for extensibility */
        uint256 tradeQuoteAmount
    ) external override onlyDelegateCall returns (uint256 makerFee) {
        Layout storage $ = _layout();

        // Maker fee: use feeBps stored in maker order at creation time (V2 compatibility)
        // Taker fee: use current config based on taker side
        uint32 makerBps = maker.feeBps;
        uint32 takerBps = taker.side == IPairV3.OrderSide.SELL ? $.sellerTakerFeeBps : $.buyerTakerFeeBps;

        // Accumulate fees
        if (makerBps != 0) makerFee = Math.mulDiv(tradeQuoteAmount, makerBps, BPS_DENOMINATOR);
        $.makerFeeAcc += makerFee;
        if (takerBps != 0) {
            uint256 takerFee = Math.mulDiv(tradeQuoteAmount, takerBps, BPS_DENOMINATOR);
            $.takerFeeAcc += takerFee;
        }
    }

    /// @notice Settle accumulated fees by transferring to feeCollector.
    /// @dev Called via delegatecall after all matches in a submit are done.
    /// @return takerFeeTotal The taker fee portion (for Pair's net calculation/event)
    function settleFees() external override onlyDelegateCall returns (uint256 takerFeeTotal) {
        Layout storage $ = _layout();

        uint256 makerFeeTotal = $.makerFeeAcc;
        takerFeeTotal = $.takerFeeAcc;
        uint256 totalFee = makerFeeTotal + takerFeeTotal;

        // Reset accumulators
        $.makerFeeAcc = 0;
        $.takerFeeAcc = 0;

        // Transfer total fee to feeCollector
        if (totalFee > 0) $.quote.safeTransfer($.feeCollector, totalFee);
        // TODO emit event
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // View helpers for Pair (returns fee bps for order.feeBps storage)
    // These are called via delegatecall from Pair when storing maker orders.
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Get seller maker fee bps (for Pair to set order.feeBps on SELL limit order)
    function sellerMakerFeeBps() external view returns (uint32) {
        return _layout().sellerMakerFeeBps;
    }

    /// @notice Get buyer maker fee bps (for Pair to set order.feeBps on BUY limit order)
    function buyerMakerFeeBps() external view returns (uint32) {
        return _layout().buyerMakerFeeBps;
    }

    /// @notice Get current fee collector address
    function feeCollector() external view returns (address) {
        return _layout().feeCollector;
    }
}
