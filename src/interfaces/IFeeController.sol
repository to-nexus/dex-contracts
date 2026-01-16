// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {IPairV3} from "./IPairV3.sol";

uint256 constant BPS_DENOMINATOR = 10000;

/// @title IFeeController
/// @notice Interface for fee controller implementations (called via delegatecall from Pair).
interface IFeeController {
    // ─────────────────────────────────────────────────────────────────────────────
    // Errors
    // ─────────────────────────────────────────────────────────────────────────────

    error FeeControllerInvalidFeeBps();
    error FeeControllerInvalidFeeStructure(uint32 makerFee, uint32 takerFee);
    error FeeControllerInvalidFeeCollector();
    error FeeControllerNotDelegateCall();
    error FeeControllerInvalidPairConfig(address quote, uint256 denominator);

    /// @notice Initialize or update fee configuration.
    /// @dev Called via delegatecall. Implementation should decode initData for its specific config.
    /// @param initData Encoded fee configuration data
    function initialize(address quote, uint256 denominator, bytes memory initData) external;

    /// @notice Calculate total QUOTE volume including fee for a BUY order.
    /// @param isMaker true if calculating for maker (limit order on book), false for taker
    /// @param order The order to calculate for
    /// @return buyVolume The total volume including fee
    function calcBuyVolumeWithFeeByOrder(bool isMaker, IPairV3.Order memory order)
        external
        view
        returns (uint256 buyVolume);

    /// @notice Calculate total QUOTE volume including fee for a given base volume.
    /// @param isMaker true if calculating for maker, false for taker
    /// @param volume The base QUOTE volume (without fee)
    /// @return buyVolume The total volume including fee
    function calcBuyVolumeWithFee(bool isMaker, uint256 volume) external view returns (uint256 buyVolume);

    /// @notice Record a match and accumulate fees.
    /// @dev Called via delegatecall for each fill. Maker fee is returned, taker fee is accumulated.
    /// @param taker The taker order (uses taker.side to determine taker fee bps)
    /// @param maker The maker order (uses maker.feeBps which was set at order creation)
    /// @param tradeAmount The trade amount in BASE (for future extensibility)
    /// @param tradeQuoteAmount The trade volume in QUOTE (fee calculation base)
    /// @return makerFee The fee charged to the maker for this fill
    function recodeMatch(
        IPairV3.Order memory taker,
        IPairV3.Order memory maker,
        uint256 tradeAmount,
        uint256 tradeQuoteAmount
    ) external returns (uint256 makerFee);

    /// @notice Settle accumulated fees by transferring to fee recipient.
    /// @dev Called via delegatecall after all matches in a submit are done.
    /// @return takerFee The taker fee portion (for Pair's net calculation/event)
    function settleFees() external returns (uint256 takerFee);

    /// @notice Get seller maker fee bps (for Pair to set order.feeBps on SELL limit order)
    /// @dev Called via delegatecall to read from Pair's namespaced storage.
    function sellerMakerFeeBps() external view returns (uint32);

    /// @notice Get buyer maker fee bps (for Pair to set order.feeBps on BUY limit order)
    /// @dev Called via delegatecall to read from Pair's namespaced storage.
    function buyerMakerFeeBps() external view returns (uint32);
}
