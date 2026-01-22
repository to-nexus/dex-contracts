// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/utils/SafeERC20.sol";
import {ERC165} from "@openzeppelin-contracts-5.5.0/utils/introspection/ERC165.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";

import {BPS_DENOMINATOR, IFeeController} from "./interfaces/IFeeController.sol";
import {IPairV3} from "./interfaces/IPairV3.sol";

/// @title FeeControllerV2Compat
/// @notice V2-compatible fee controller implementing 4 fee bps (seller/buyer × maker/taker).
///         Designed to be called via delegatecall from PairImplV3.
/// @dev Persistent config stored in Pair's storage via ERC-7201 namespaced slot.
///      Per-transaction data stored in transient storage for gas efficiency.
contract FeeControllerV2Compat is IFeeController, ERC165 {
    using SafeERC20 for IERC20;
    using Math for uint256;

    // ─────────────────────────────────────────────────────────────────────────────
    // Events
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Emitted when fees are settled and transferred to feeCollector.
    /// @param takerId The taker order ID that initiated this fee settlement
    /// @param feeCollector The address receiving the fees
    /// @param totalFee The total fee amount transferred (maker + taker fees)
    event FeeControllerFeesSettled(uint256 indexed takerId, address indexed feeCollector, uint256 totalFee);

    // ─────────────────────────────────────────────────────────────────────────────
    // ERC-7201 Namespaced Persistent Storage
    // ─────────────────────────────────────────────────────────────────────────────

    /// @custom:storage-location erc7201:cross.storage.FeeControllerV2Compat
    struct FeeControllerV2CompatStorage {
        // --- Persistent config (set via initialize) ---
        address feeCollector;
        uint32 sellerMakerFeeBps;
        uint32 sellerTakerFeeBps;
        uint32 buyerMakerFeeBps;
        uint32 buyerTakerFeeBps;
        // Cached from Pair for gas savings
        IERC20 quote;
        uint256 denominator;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.FeeControllerV2Compat")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant FeeControllerV2CompatStorageLocation =
        0xda974a09e7b8542215f909d6fdf2a86dab7afcfd32616d508be15f49de00f500;

    function _getFeeControllerV2CompatStorage() private pure returns (FeeControllerV2CompatStorage storage $) {
        assembly {
            $.slot := FeeControllerV2CompatStorageLocation
        }
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // ERC-7201 Namespaced Transient Storage (EIP-1153)
    // ─────────────────────────────────────────────────────────────────────────────

    /// @custom:storage-location erc7201:cross.storage.FeeControllerV2Compat.transient
    /// Slot offsets from FeeControllerV2CompatTransientLocation:
    ///   +0: currentTakerId (uint256) - validates same taker across matches
    ///   +1: takerFeeBps (uint32) - cached taker fee bps for gas optimization
    ///   +2: makerFeeAcc (uint256) - accumulated maker fees
    ///   +3: takerFeeAcc (uint256) - accumulated taker fees

    // keccak256(abi.encode(uint256(keccak256("cross.storage.FeeControllerV2Compat.transient")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant FeeControllerV2CompatTransientLocation =
        0xaf6c89c994254c39a619ba26987a4abab224fa2af9a4311d25848a45ac164700;

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
        _checkDelegateCall();
        _;
    }

    function _checkDelegateCall() private view {
        if (address(this) == _SELF) revert FeeControllerNotDelegateCall();
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

        FeeControllerV2CompatStorage storage $ = _getFeeControllerV2CompatStorage();
        $.feeCollector = _feeCollector;
        $.sellerMakerFeeBps = sMk;
        $.sellerTakerFeeBps = sTk;
        $.buyerMakerFeeBps = bMk;
        $.buyerTakerFeeBps = bTk;
        // Note: fee accumulators are in transient storage, auto-reset per transaction
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
        FeeControllerV2CompatStorage storage $ = _getFeeControllerV2CompatStorage();
        uint256 baseVolume = Math.mulDiv(order.price, order.amount, $.denominator);
        uint32 bps = isMaker ? $.buyerMakerFeeBps : $.buyerTakerFeeBps;
        return baseVolume + Math.mulDiv(baseVolume, bps, BPS_DENOMINATOR);
    }

    /// @notice Calculate total QUOTE volume including fee for a given base volume.
    /// @param isMaker true if calculating for maker, false for taker
    /// @param volume The base QUOTE volume (without fee)
    /// @return buyVolume The total volume including fee
    function calcBuyVolumeWithFee(bool isMaker, uint256 volume) external view override returns (uint256 buyVolume) {
        FeeControllerV2CompatStorage storage $ = _getFeeControllerV2CompatStorage();
        uint32 bps = isMaker ? $.buyerMakerFeeBps : $.buyerTakerFeeBps;
        return volume + Math.mulDiv(volume, bps, BPS_DENOMINATOR);
    }

    /// @notice Record a match and accumulate fees.
    /// @dev Called via delegatecall for each fill during matching.
    ///      - On first call: caches takerId and takerFeeBps in transient storage
    ///      - On subsequent calls: validates takerId and uses cached takerFeeBps
    ///      - Maker fee uses maker.feeBps (set at order creation time for V2 compatibility)
    /// @param takerId The taker order ID (for transient storage validation)
    /// @param taker The taker order (uses taker.side to determine taker fee bps on first call)
    /// @param maker The maker order (uses maker.feeBps which was set at order creation)
    /// @param tradeQuoteAmount The trade volume in QUOTE (fee calculation base)
    /// @return makerFee The fee charged to the maker for this fill
    function recordMatch(
        uint256 takerId,
        uint256, /* makerId - unused in V2Compat */
        IPairV3.Order memory taker,
        IPairV3.Order memory maker,
        uint256, /* tradeAmount - unused in V2Compat, reserved for extensibility */
        uint256 tradeQuoteAmount
    ) external override onlyDelegateCall returns (uint256 makerFee) {
        uint32 takerBps;

        // Check if this is the first recordMatch call in this transaction
        uint256 currentTakerId = _tloadTakerId();
        if (currentTakerId == 0) {
            // First call: cache takerId and takerFeeBps
            _tstoreTakerId(takerId);
            FeeControllerV2CompatStorage storage $ = _getFeeControllerV2CompatStorage();
            takerBps = taker.side == IPairV3.OrderSide.SELL ? $.sellerTakerFeeBps : $.buyerTakerFeeBps;
            _tstoreTakerFeeBps(takerBps);
        } else {
            // Subsequent call: validate takerId matches
            if (currentTakerId != takerId) revert FeeControllerTakerIdMismatch(currentTakerId, takerId);
            // Use cached takerFeeBps
            takerBps = _tloadTakerFeeBps();
        }

        // Maker fee: use feeBps stored in maker order at creation time (V2 compatibility)
        uint32 makerBps = maker.feeBps;

        // Accumulate fees in transient storage
        if (makerBps != 0) makerFee = Math.mulDiv(tradeQuoteAmount, makerBps, BPS_DENOMINATOR);
        _tstoreMakerFeeAcc(_tloadMakerFeeAcc() + makerFee);

        if (takerBps != 0) {
            uint256 takerFee = Math.mulDiv(tradeQuoteAmount, takerBps, BPS_DENOMINATOR);
            _tstoreTakerFeeAcc(_tloadTakerFeeAcc() + takerFee);
        }
    }

    /// @notice Settle accumulated fees by transferring to feeCollector.
    /// @dev Called via delegatecall after all matches in a submit are done.
    ///      Reads accumulated fees from transient storage (auto-reset at tx end).
    ///      Follows CEI pattern: Effects before Interactions to prevent reentrancy.
    /// @return takerFeeTotal The taker fee portion (for Pair's net calculation/event)
    function settleFees() external override onlyDelegateCall returns (uint256 takerFeeTotal) {
        // Read from transient storage
        uint256 takerId = _tloadTakerId();
        uint256 makerFeeTotal = _tloadMakerFeeAcc();
        takerFeeTotal = _tloadTakerFeeAcc();
        uint256 totalFee = makerFeeTotal + takerFeeTotal;

        // Effects: Reset transient storage BEFORE external call (CEI pattern)
        // Always reset even if totalFee == 0 to allow subsequent trades in same tx
        _tstoreTakerId(0);
        _tstoreMakerFeeAcc(0);
        _tstoreTakerFeeAcc(0);

        // Interactions: Transfer fee and emit event LAST
        if (totalFee > 0) {
            FeeControllerV2CompatStorage storage $ = _getFeeControllerV2CompatStorage();
            $.quote.safeTransfer($.feeCollector, totalFee);
            emit FeeControllerFeesSettled(takerId, $.feeCollector, totalFee);
        }
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // View helpers for Pair (returns fee bps for order.feeBps storage)
    // These are called via delegatecall from Pair when storing maker orders.
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Get seller maker fee bps (for Pair to set order.feeBps on SELL limit order)
    function sellerMakerFeeBps() external view returns (uint32) {
        return _getFeeControllerV2CompatStorage().sellerMakerFeeBps;
    }

    /// @notice Get buyer maker fee bps (for Pair to set order.feeBps on BUY limit order)
    function buyerMakerFeeBps() external view returns (uint32) {
        return _getFeeControllerV2CompatStorage().buyerMakerFeeBps;
    }

    /// @notice Get all effective fee bps in a single call.
    function getEffectiveFees() external view returns (uint32, uint32, uint32, uint32) {
        FeeControllerV2CompatStorage storage $ = _getFeeControllerV2CompatStorage();
        return ($.sellerMakerFeeBps, $.sellerTakerFeeBps, $.buyerMakerFeeBps, $.buyerTakerFeeBps);
    }

    /// @notice Get the configuration identifier for this FeeController implementation.
    /// @return configId keccak256("FeeControllerV2Compat.v1")
    function getConfigId() external pure returns (bytes32) {
        return keccak256("FeeControllerV2Compat.v1");
    }

    /// @notice Get the entire fee configuration storage as encoded bytes.
    /// @dev Returns: abi.encode(feeCollector, sellerMakerFeeBps, sellerTakerFeeBps, buyerMakerFeeBps, buyerTakerFeeBps)
    ///      Note: quote and denominator are excluded as they are Pair-level config, not fee config.
    /// @return data ABI-encoded fee configuration
    function getStorage() external view returns (bytes memory) {
        FeeControllerV2CompatStorage storage $ = _getFeeControllerV2CompatStorage();
        return
            abi.encode($.feeCollector, $.sellerMakerFeeBps, $.sellerTakerFeeBps, $.buyerMakerFeeBps, $.buyerTakerFeeBps);
    }

    function supportsInterface(bytes4 interfaceId) public view override returns (bool) {
        return interfaceId == type(IFeeController).interfaceId || super.supportsInterface(interfaceId);
    }

    function _tloadTakerId() private view returns (uint256 value) {
        bytes32 slot = FeeControllerV2CompatTransientLocation;
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreTakerId(uint256 value) private {
        bytes32 slot = FeeControllerV2CompatTransientLocation;
        assembly {
            tstore(slot, value)
        }
    }

    function _tloadTakerFeeBps() private view returns (uint32 value) {
        bytes32 slot = bytes32(uint256(FeeControllerV2CompatTransientLocation) + 1);
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreTakerFeeBps(uint32 value) private {
        bytes32 slot = bytes32(uint256(FeeControllerV2CompatTransientLocation) + 1);
        assembly {
            tstore(slot, value)
        }
    }

    function _tloadMakerFeeAcc() private view returns (uint256 value) {
        bytes32 slot = bytes32(uint256(FeeControllerV2CompatTransientLocation) + 2);
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreMakerFeeAcc(uint256 value) private {
        bytes32 slot = bytes32(uint256(FeeControllerV2CompatTransientLocation) + 2);
        assembly {
            tstore(slot, value)
        }
    }

    function _tloadTakerFeeAcc() private view returns (uint256 value) {
        bytes32 slot = bytes32(uint256(FeeControllerV2CompatTransientLocation) + 3);
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreTakerFeeAcc(uint256 value) private {
        bytes32 slot = bytes32(uint256(FeeControllerV2CompatTransientLocation) + 3);
        assembly {
            tstore(slot, value)
        }
    }
}
