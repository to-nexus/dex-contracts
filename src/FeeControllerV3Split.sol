// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/utils/SafeERC20.sol";
import {ERC165} from "@openzeppelin-contracts-5.5.0/utils/introspection/ERC165.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";

import {BPS_DENOMINATOR, IFeeController} from "./interfaces/IFeeController.sol";
import {IPairV3} from "./interfaces/IPairV3.sol";

/// @title FeeControllerV3Split
/// @notice V3 fee controller implementing taker-only fee with 3-way split (creator/maker rebate/system).
///         MakerFee is always 0. TakerFee is split into: creator, maker rebate (paid immediately), and feeCollector.
///         Designed to be called via delegatecall from PairImplV3.
/// @dev Persistent config stored in Pair's storage via ERC-7201 namespaced slot.
///      Per-transaction data stored in transient storage for gas efficiency.
contract FeeControllerV3Split is IFeeController, ERC165 {
    using SafeERC20 for IERC20;
    using Math for uint256;

    // ─────────────────────────────────────────────────────────────────────────────
    // Events
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Emitted when fees are settled and distributed.
    /// @param takerId The taker order ID that initiated this fee settlement
    /// @param totalTakerFee Total taker fee collected (before split)
    /// @param creatorFee Amount sent to creator
    /// @param feeCollectorFee Amount sent to feeCollector (system fee)
    /// @param makerRebatePaidTotal Amount already paid as maker rebates during matching
    /// @param creator Address of the creator receiving creatorFee
    /// @param feeCollector Address of the system fee collector
    event FeeControllerV3FeesSettled(
        uint256 indexed takerId,
        uint256 totalTakerFee,
        uint256 creatorFee,
        uint256 feeCollectorFee,
        uint256 makerRebatePaidTotal,
        address indexed creator,
        address indexed feeCollector
    );

    /// @notice Emitted when maker rebate is paid.
    /// @param orderId The order ID that initiated this fee settlement
    /// @param maker The maker order (maker.owner receives rebate)
    /// @param rebate The amount of maker rebate paid
    event FeeControllerV3MakerRebatePaid(uint256 indexed orderId, address indexed maker, uint256 rebate);

    // ─────────────────────────────────────────────────────────────────────────────
    // ERC-7201 Namespaced Persistent Storage
    // ─────────────────────────────────────────────────────────────────────────────

    /// @custom:storage-location erc7201:cross.storage.FeeControllerV3Split
    struct FeeControllerV3SplitStorage {
        // --- Persistent config (set via initialize) ---
        address feeCollector;
        address creator;
        uint32 takerFeeBps;
        uint32 creatorShareBps;
        uint32 makerRebateShareBps;
        // Cached from Pair for gas savings
        IERC20 quote;
        uint256 denominator;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.FeeControllerV3Split")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant FeeControllerV3SplitStorageLocation =
        0xc063cf2dfd170dd7c3dd72990f0424c6080fb35c03fb56e3a8e6c5b294381000;

    function _getFeeControllerV3SplitStorage() private pure returns (FeeControllerV3SplitStorage storage $) {
        assembly {
            $.slot := FeeControllerV3SplitStorageLocation
        }
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // ERC-7201 Namespaced Transient Storage (EIP-1153)
    // ─────────────────────────────────────────────────────────────────────────────

    /// @custom:storage-location erc7201:cross.storage.FeeControllerV3Split.transient
    /// Slot offsets from FeeControllerV3SplitTransientLocation:
    ///   +0: currentTakerId (uint256) - validates same taker across matches
    ///   +1: takerFeeAccTotal (uint256) - accumulated total taker fees
    ///   +2: makerRebatePaidTotal (uint256) - accumulated maker rebates already paid

    // keccak256(abi.encode(uint256(keccak256("cross.storage.FeeControllerV3Split.transient")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant FeeControllerV3SplitTransientLocation =
        0xc402bc4465edd18493ad1ac2cd9f5e2cd78832a5c76906ac70c20b838757df00;

    // ─────────────────────────────────────────────────────────────────────────────
    // Delegatecall enforcement
    // ─────────────────────────────────────────────────────────────────────────────

    /// @dev Store the original address at deployment for delegatecall check.
    address private immutable _SELF;

    constructor() {
        _SELF = address(this);
    }

    /// @dev Ensures the function is called via delegatecall.
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
    /// @dev Called via delegatecall from Pair.
    /// @param quote The QUOTE token address from Pair
    /// @param denominator The BASE token denominator from Pair
    /// @param initData abi.encode(feeCollector, creator, takerFeeBps, creatorShareBps, makerRebateShareBps)
    function initialize(address quote, uint256 denominator, bytes memory initData) external override onlyDelegateCall {
        (
            address _feeCollector,
            address _creator,
            uint32 _takerFeeBps,
            uint32 _creatorShareBps,
            uint32 _makerRebateShareBps
        ) = abi.decode(initData, (address, address, uint32, uint32, uint32));

        // Validation
        if (_feeCollector == address(0)) revert FeeControllerInvalidFeeCollector();
        if (_creator == address(0)) revert FeeControllerInvalidFeeCollector(); // reuse error for creator
        if (_takerFeeBps >= BPS_DENOMINATOR) revert FeeControllerInvalidFeeBps();
        if (_creatorShareBps + _makerRebateShareBps > BPS_DENOMINATOR) revert FeeControllerInvalidFeeBps();

        FeeControllerV3SplitStorage storage $ = _getFeeControllerV3SplitStorage();
        $.feeCollector = _feeCollector;
        $.creator = _creator;
        $.takerFeeBps = _takerFeeBps;
        $.creatorShareBps = _creatorShareBps;
        $.makerRebateShareBps = _makerRebateShareBps;

        // Cache quote/denominator from Pair (immutable after first init)
        if (address($.quote) == address(0)) {
            $.quote = IERC20(quote);
            $.denominator = denominator;
        } else {
            if (address($.quote) != quote) revert FeeControllerInvalidPairConfig(quote, denominator);
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
        FeeControllerV3SplitStorage storage $ = _getFeeControllerV3SplitStorage();
        uint256 baseVolume = Math.mulDiv(order.price, order.amount, $.denominator);
        // MakerFee is always 0 in V3Split
        if (isMaker) return baseVolume;
        // TakerFee applies
        return baseVolume + Math.mulDiv(baseVolume, $.takerFeeBps, BPS_DENOMINATOR);
    }

    /// @notice Calculate total QUOTE volume including fee for a given base volume.
    /// @param isMaker true if calculating for maker, false for taker
    /// @param volume The base QUOTE volume (without fee)
    /// @return buyVolume The total volume including fee
    function calcBuyVolumeWithFee(bool isMaker, uint256 volume) external view override returns (uint256 buyVolume) {
        // MakerFee is always 0 in V3Split
        if (isMaker) return volume;
        FeeControllerV3SplitStorage storage $ = _getFeeControllerV3SplitStorage();
        return volume + Math.mulDiv(volume, $.takerFeeBps, BPS_DENOMINATOR);
    }

    /// @notice Record a match, calculate taker fee, and immediately pay maker rebate.
    /// @dev Called via delegatecall for each fill during matching.
    ///      - MakerFee is always 0 (returned value)
    ///      - TakerFee is accumulated in transient storage
    ///      - Maker rebate is calculated and immediately transferred to maker
    /// @param takerId The taker order ID (for transient storage validation)
    /// @param makerId The maker order ID (for event emission)
    /// @param maker The maker order (maker.owner receives rebate)
    /// @param tradeQuoteAmount The trade volume in QUOTE (fee calculation base)
    /// @return makerFee Always returns 0 (maker pays no fee in V3Split)
    function recordMatch(
        uint256 takerId,
        uint256 makerId,
        IPairV3.Order memory, /* taker - unused */
        IPairV3.Order memory maker,
        uint256, /* tradeAmount - unused */
        uint256 tradeQuoteAmount
    ) external override onlyDelegateCall returns (uint256 makerFee) {
        // Check if this is the first recordMatch call in this transaction
        uint256 currentTakerId = _tloadTakerId();
        if (currentTakerId == 0) {
            // First call: cache takerId
            _tstoreTakerId(takerId);
        } else {
            // Subsequent call: validate takerId matches
            if (currentTakerId != takerId) revert FeeControllerTakerIdMismatch(currentTakerId, takerId);
        }

        FeeControllerV3SplitStorage storage $ = _getFeeControllerV3SplitStorage();

        // Calculate taker fee and accumulate
        uint256 takerFee = 0;
        if ($.takerFeeBps != 0) {
            takerFee = Math.mulDiv(tradeQuoteAmount, $.takerFeeBps, BPS_DENOMINATOR);
            _tstoreTakerFeeAcc(_tloadTakerFeeAcc() + takerFee);
        }

        // Calculate and immediately pay maker rebate
        if (takerFee != 0 && $.makerRebateShareBps != 0) {
            uint256 rebate = Math.mulDiv(takerFee, $.makerRebateShareBps, BPS_DENOMINATOR);
            if (rebate != 0) {
                _tstoreMakerRebatePaid(_tloadMakerRebatePaid() + rebate);
                $.quote.safeTransfer(maker.owner, rebate);
                emit FeeControllerV3MakerRebatePaid(makerId, maker.owner, rebate);
            }
        }

        // MakerFee is always 0 in V3Split
        return 0;
    }

    /// @notice Settle accumulated fees by transferring to creator and feeCollector.
    /// @dev Called via delegatecall after all matches in a submit are done.
    ///      Distribution: creatorFee from total, feeCollector gets remainder after rebates.
    ///      Follows CEI pattern: Effects before Interactions.
    /// @return takerFeeTotal The total taker fee (for Pair's net calculation/event)
    function settleFees() external override onlyDelegateCall returns (uint256 takerFeeTotal) {
        // Read from transient storage
        uint256 takerId = _tloadTakerId();
        takerFeeTotal = _tloadTakerFeeAcc();
        uint256 makerRebatePaid = _tloadMakerRebatePaid();

        // Effects: Reset transient storage BEFORE external calls (CEI pattern)
        // Always reset even if takerFeeTotal == 0 to allow subsequent trades in same tx
        _tstoreTakerId(0);
        _tstoreTakerFeeAcc(0);
        _tstoreMakerRebatePaid(0);

        // Calculate fee distribution
        FeeControllerV3SplitStorage storage $ = _getFeeControllerV3SplitStorage();

        if (takerFeeTotal > 0) {
            // Calculate creator's share from total taker fee
            uint256 creatorFee = Math.mulDiv(takerFeeTotal, $.creatorShareBps, BPS_DENOMINATOR);

            // Collector gets: total - creatorFee - already paid rebates
            // This ensures no dust is lost
            uint256 collectorFee = takerFeeTotal - creatorFee - makerRebatePaid;

            // Interactions: Transfer fees LAST
            if (creatorFee != 0) $.quote.safeTransfer($.creator, creatorFee);
            if (collectorFee != 0) $.quote.safeTransfer($.feeCollector, collectorFee);

            // Emit settlement event
            emit FeeControllerV3FeesSettled(
                takerId, takerFeeTotal, creatorFee, collectorFee, makerRebatePaid, $.creator, $.feeCollector
            );
        }
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // View helpers for Pair (returns fee bps for order.feeBps storage)
    // MakerFee is always 0 in V3Split
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Get seller maker fee bps (always 0 for V3Split)
    function sellerMakerFeeBps() external pure returns (uint32) {
        return 0;
    }

    /// @notice Get buyer maker fee bps (always 0 for V3Split)
    function buyerMakerFeeBps() external pure returns (uint32) {
        return 0;
    }

    /// @notice Get current fee collector address
    function feeCollector() external view returns (address) {
        return _getFeeControllerV3SplitStorage().feeCollector;
    }

    /// @notice Get current creator address
    function creator() external view returns (address) {
        return _getFeeControllerV3SplitStorage().creator;
    }

    /// @notice Get taker fee bps
    function takerFeeBps() external view returns (uint32) {
        return _getFeeControllerV3SplitStorage().takerFeeBps;
    }

    /// @notice Get creator share bps (percentage of taker fee)
    function creatorShareBps() external view returns (uint32) {
        return _getFeeControllerV3SplitStorage().creatorShareBps;
    }

    /// @notice Get maker rebate share bps (percentage of taker fee)
    function makerRebateShareBps() external view returns (uint32) {
        return _getFeeControllerV3SplitStorage().makerRebateShareBps;
    }

    function supportsInterface(bytes4 interfaceId) public view override returns (bool) {
        return interfaceId == type(IFeeController).interfaceId || super.supportsInterface(interfaceId);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Transient Storage Helpers
    // ─────────────────────────────────────────────────────────────────────────────

    function _tloadTakerId() private view returns (uint256 value) {
        bytes32 slot = FeeControllerV3SplitTransientLocation;
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreTakerId(uint256 value) private {
        bytes32 slot = FeeControllerV3SplitTransientLocation;
        assembly {
            tstore(slot, value)
        }
    }

    function _tloadTakerFeeAcc() private view returns (uint256 value) {
        bytes32 slot = bytes32(uint256(FeeControllerV3SplitTransientLocation) + 1);
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreTakerFeeAcc(uint256 value) private {
        bytes32 slot = bytes32(uint256(FeeControllerV3SplitTransientLocation) + 1);
        assembly {
            tstore(slot, value)
        }
    }

    function _tloadMakerRebatePaid() private view returns (uint256 value) {
        bytes32 slot = bytes32(uint256(FeeControllerV3SplitTransientLocation) + 2);
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreMakerRebatePaid(uint256 value) private {
        bytes32 slot = bytes32(uint256(FeeControllerV3SplitTransientLocation) + 2);
        assembly {
            tstore(slot, value)
        }
    }
}
