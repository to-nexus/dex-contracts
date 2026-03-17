// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/utils/SafeERC20.sol";
import {ERC165} from "@openzeppelin-contracts-5.5.0/utils/introspection/ERC165.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";

import {BPS_DENOMINATOR, IFeeController} from "./interfaces/IFeeController.sol";
import {IPairV3} from "./interfaces/IPairV3.sol";

/// @title FeeControllerV3Dist
/// @notice 4-way fee controller (seller/buyer x maker/taker) with N-way distribution to multiple recipients by BPS ratio.
///         Designed to be called via delegatecall from PairImplV3.
/// @dev Persistent config stored in Pair's storage via ERC-7201 namespaced slot.
///      Per-transaction data stored in transient storage for gas efficiency.
contract FeeControllerV3Dist is IFeeController, ERC165 {
    using SafeERC20 for IERC20;
    using Math for uint256;

    uint8 public constant MAX_RECIPIENTS = 10;

    error FeeControllerV3DistInvalidRecipients();
    error FeeControllerV3DistRatiosBpsSumNot10000(uint256 bpsSum);

    /// @notice Emitted when fees are settled and distributed to multiple recipients.
    /// @param takerId The taker order ID that initiated this fee settlement
    /// @param totalFee The total fee amount distributed (maker + taker)
    /// @param recipients Addresses receiving the fees
    /// @param amounts Amounts transferred to each recipient
    /// @param labels bytes32 label per recipient (e.g. CREATOR, PLATFORM)
    event FeeControllerV3DistFeesSettled(
        uint256 indexed takerId, uint256 totalFee, address[] recipients, uint256[] amounts, bytes32[] labels
    );

    /// @custom:storage-location erc7201:cross.storage.FeeControllerV3Dist
    struct FeeControllerV3DistStorage {
        uint32 sellerMakerFeeBps;
        uint32 sellerTakerFeeBps;
        uint32 buyerMakerFeeBps;
        uint32 buyerTakerFeeBps;
        IERC20 quote;
        uint256 denominator;
        uint8 recipientCount;
        address[] recipients;
        uint32[] ratios;
        bytes32[] labels;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.FeeControllerV3Dist")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant FeeControllerV3DistStorageLocation =
        0xcf8197069ed4914f85bcb17827b19530499b47ac931f1345bd9e11975128a100;

    function _getFeeControllerV3DistStorage() private pure returns (FeeControllerV3DistStorage storage $) {
        assembly {
            $.slot := FeeControllerV3DistStorageLocation
        }
    }

    /// @custom:storage-location erc7201:cross.storage.FeeControllerV3Dist.transient
    /// Slot offsets: +0 currentTakerId, +1 takerFeeBps, +2 makerFeeAcc, +3 takerFeeAcc
    // keccak256(abi.encode(uint256(keccak256("cross.storage.FeeControllerV3Dist.transient")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant FeeControllerV3DistTransientLocation =
        0xa65ece664cc680a5cef1d3092a2777f9d43ea55cc7a3dfff71e7ffe246389c00;

    address private immutable _SELF;

    constructor() {
        _SELF = address(this);
    }

    modifier onlyDelegateCall() {
        if (address(this) == _SELF) revert FeeControllerNotDelegateCall();
        _;
    }

    /// @notice Initialize or update fee configuration.
    /// @param initData abi.encode(sellerMakerFeeBps, sellerTakerFeeBps, buyerMakerFeeBps, buyerTakerFeeBps, recipients[], ratios[], labels[])
    function initialize(address quote, uint256 denominator, bytes memory initData) external override onlyDelegateCall {
        (
            uint32 sMk,
            uint32 sTk,
            uint32 bMk,
            uint32 bTk,
            address[] memory recipients,
            uint32[] memory ratios,
            bytes32[] memory labels
        ) = abi.decode(initData, (uint32, uint32, uint32, uint32, address[], uint32[], bytes32[]));

        if (sMk >= BPS_DENOMINATOR || sTk >= BPS_DENOMINATOR) revert FeeControllerInvalidFeeBps();
        if (bMk >= BPS_DENOMINATOR || bTk >= BPS_DENOMINATOR) revert FeeControllerInvalidFeeBps();
        if (sTk < sMk) revert FeeControllerInvalidFeeStructure(sMk, sTk);
        if (bTk < bMk) revert FeeControllerInvalidFeeStructure(bMk, bTk);

        uint256 n = recipients.length;
        if (n == 0 || n != ratios.length || n != labels.length || n > MAX_RECIPIENTS) {
            revert FeeControllerV3DistInvalidRecipients();
        }

        uint256 ratioSum = 0;
        for (uint256 i = 0; i < n; ++i) {
            if (recipients[i] == address(0)) revert FeeControllerV3DistInvalidRecipients();
            if (ratios[i] == 0) revert FeeControllerV3DistInvalidRecipients();
            ratioSum += ratios[i];
        }
        if (ratioSum != BPS_DENOMINATOR) revert FeeControllerV3DistRatiosBpsSumNot10000(ratioSum);

        FeeControllerV3DistStorage storage $ = _getFeeControllerV3DistStorage();
        $.sellerMakerFeeBps = sMk;
        $.sellerTakerFeeBps = sTk;
        $.buyerMakerFeeBps = bMk;
        $.buyerTakerFeeBps = bTk;
        $.recipientCount = uint8(n);

        delete $.recipients;
        delete $.ratios;
        delete $.labels;
        for (uint256 i = 0; i < n; ++i) {
            $.recipients.push(recipients[i]);
            $.ratios.push(ratios[i]);
            $.labels.push(labels[i]);
        }

        if (address($.quote) == address(0)) {
            if (quote == address(0)) revert FeeControllerInvalidPairConfig(quote, denominator);
            $.quote = IERC20(quote);
            $.denominator = denominator;
        } else {
            if (address($.quote) != quote) revert FeeControllerInvalidPairConfig(quote, denominator);
        }
    }

    function calcBuyVolumeWithFeeByOrder(bool isMaker, IPairV3.Order memory order)
        external
        view
        override
        returns (uint256 buyVolume)
    {
        FeeControllerV3DistStorage storage $ = _getFeeControllerV3DistStorage();
        uint256 baseVolume = Math.mulDiv(order.price, order.amount, $.denominator);
        uint32 bps = isMaker ? $.buyerMakerFeeBps : $.buyerTakerFeeBps;
        return baseVolume + Math.mulDiv(baseVolume, bps, BPS_DENOMINATOR);
    }

    function calcBuyVolumeWithFee(bool isMaker, uint256 volume) external view override returns (uint256 buyVolume) {
        FeeControllerV3DistStorage storage $ = _getFeeControllerV3DistStorage();
        uint32 bps = isMaker ? $.buyerMakerFeeBps : $.buyerTakerFeeBps;
        return volume + Math.mulDiv(volume, bps, BPS_DENOMINATOR);
    }

    function recordMatch(
        uint256 takerId,
        uint256,
        IPairV3.Order memory taker,
        IPairV3.Order memory maker,
        uint256,
        uint256 tradeQuoteAmount
    ) external override onlyDelegateCall returns (uint256 makerFee) {
        uint32 takerBps;
        uint256 currentTakerId = _tloadTakerId();
        if (currentTakerId == 0) {
            _tstoreTakerId(takerId);
            FeeControllerV3DistStorage storage $ = _getFeeControllerV3DistStorage();
            takerBps = taker.side == IPairV3.OrderSide.SELL ? $.sellerTakerFeeBps : $.buyerTakerFeeBps;
            _tstoreTakerFeeBps(takerBps);
        } else {
            if (currentTakerId != takerId) revert FeeControllerTakerIdMismatch(currentTakerId, takerId);
            takerBps = _tloadTakerFeeBps();
        }

        uint32 makerBps = maker.feeBps;
        if (makerBps != 0) {
            makerFee = Math.mulDiv(tradeQuoteAmount, makerBps, BPS_DENOMINATOR);
            _tstoreMakerFeeAcc(_tloadMakerFeeAcc() + makerFee);
        }

        if (takerBps != 0) {
            uint256 takerFee = Math.mulDiv(tradeQuoteAmount, takerBps, BPS_DENOMINATOR);
            _tstoreTakerFeeAcc(_tloadTakerFeeAcc() + takerFee);
        }
    }

    function settleFees() external override onlyDelegateCall returns (uint256 takerFeeTotal) {
        uint256 takerId = _tloadTakerId();
        uint256 makerFeeTotal = _tloadMakerFeeAcc();
        takerFeeTotal = _tloadTakerFeeAcc();
        uint256 totalFee = makerFeeTotal + takerFeeTotal;

        _tstoreTakerId(0);
        _tstoreTakerFeeBps(0);
        _tstoreMakerFeeAcc(0);
        _tstoreTakerFeeAcc(0);

        if (totalFee > 0) {
            FeeControllerV3DistStorage storage $ = _getFeeControllerV3DistStorage();
            uint8 n = $.recipientCount;
            address[] memory recipients = new address[](n);
            uint256[] memory amounts = new uint256[](n);
            bytes32[] memory labels = new bytes32[](n);
            uint256 distributed = 0;

            for (uint256 i = 0; i < n; ++i) {
                recipients[i] = $.recipients[i];
                labels[i] = $.labels[i];
                if (i < n - 1) {
                    amounts[i] = Math.mulDiv(totalFee, $.ratios[i], BPS_DENOMINATOR);
                    distributed += amounts[i];
                } else {
                    amounts[i] = totalFee - distributed;
                }
            }

            for (uint256 i = 0; i < n; ++i) {
                if (amounts[i] != 0) $.quote.safeTransfer(recipients[i], amounts[i]);
            }

            emit FeeControllerV3DistFeesSettled(takerId, totalFee, recipients, amounts, labels);
        }
    }

    function sellerMakerFeeBps() external view returns (uint32) {
        return _getFeeControllerV3DistStorage().sellerMakerFeeBps;
    }

    function buyerMakerFeeBps() external view returns (uint32) {
        return _getFeeControllerV3DistStorage().buyerMakerFeeBps;
    }

    function getEffectiveFees() external view returns (uint32, uint32, uint32, uint32) {
        FeeControllerV3DistStorage storage $ = _getFeeControllerV3DistStorage();
        return ($.sellerMakerFeeBps, $.sellerTakerFeeBps, $.buyerMakerFeeBps, $.buyerTakerFeeBps);
    }

    function getConfigId() external pure returns (bytes32) {
        return keccak256("FeeControllerV3Dist.v1");
    }

    function getStorage() external view returns (bytes memory) {
        FeeControllerV3DistStorage storage $ = _getFeeControllerV3DistStorage();
        return abi.encode(
            $.sellerMakerFeeBps,
            $.sellerTakerFeeBps,
            $.buyerMakerFeeBps,
            $.buyerTakerFeeBps,
            $.recipients,
            $.ratios,
            $.labels
        );
    }

    function supportsInterface(bytes4 interfaceId) public view override returns (bool) {
        return interfaceId == type(IFeeController).interfaceId || super.supportsInterface(interfaceId);
    }

    function _tloadTakerId() private view returns (uint256 value) {
        bytes32 slot = FeeControllerV3DistTransientLocation;
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreTakerId(uint256 value) private {
        bytes32 slot = FeeControllerV3DistTransientLocation;
        assembly {
            tstore(slot, value)
        }
    }

    function _tloadTakerFeeBps() private view returns (uint32 value) {
        bytes32 slot = bytes32(uint256(FeeControllerV3DistTransientLocation) + 1);
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreTakerFeeBps(uint32 value) private {
        bytes32 slot = bytes32(uint256(FeeControllerV3DistTransientLocation) + 1);
        assembly {
            tstore(slot, value)
        }
    }

    function _tloadMakerFeeAcc() private view returns (uint256 value) {
        bytes32 slot = bytes32(uint256(FeeControllerV3DistTransientLocation) + 2);
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreMakerFeeAcc(uint256 value) private {
        bytes32 slot = bytes32(uint256(FeeControllerV3DistTransientLocation) + 2);
        assembly {
            tstore(slot, value)
        }
    }

    function _tloadTakerFeeAcc() private view returns (uint256 value) {
        bytes32 slot = bytes32(uint256(FeeControllerV3DistTransientLocation) + 3);
        assembly {
            value := tload(slot)
        }
    }

    function _tstoreTakerFeeAcc(uint256 value) private {
        bytes32 slot = bytes32(uint256(FeeControllerV3DistTransientLocation) + 3);
        assembly {
            tstore(slot, value)
        }
    }
}
