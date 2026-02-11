// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";

interface IPair {
    enum OrderSide {
        SELL,
        BUY
    }

    enum LimitConstraints {
        GOOD_TILL_CANCEL,
        IMMEDIATE_OR_CANCEL,
        FILL_OR_KILL
    }

    enum CloseType {
        COMPLETED,
        IMMEDIATE_OR_CANCEL,
        CANCEL,
        EMERGENCY
    }

    struct Order {
        OrderSide side;
        address owner;
        uint32 feeBps;
        uint256 price;
        uint256 amount;
    }

    struct Config {
        IERC20 QUOTE;
        IERC20 BASE;
        uint256 DENOMINATOR;
    }

    function getConfig() external view returns (Config memory);
    function findPrevPrice(OrderSide side, uint256 price, uint256[2] calldata adjacent, uint256 findMaxCount)
        external
        view
        returns (uint256);
    function submitLimitOrder(
        Order memory order,
        LimitConstraints constraints,
        uint256 prevPrice,
        uint256 maxMatchCount
    ) external returns (uint256 orderId);
    function submitMarketOrder(Order memory order, uint256 spendAmount, uint256 maxMatchCount) external;
    function cancelOrder(address caller, uint256[] memory orderIds) external;
}

interface IPairV2 is IPair {
    struct FeeConfig {
        uint32 sellerMakerFeeBps; // Seller Maker fee (BPS)
        uint32 sellerTakerFeeBps; // Seller Taker fee (BPS)
        uint32 buyerMakerFeeBps; // Buyer Maker fee (BPS)
        uint32 buyerTakerFeeBps; // Buyer Taker fee (BPS)
    }

    function feeConfig()
        external
        view
        returns (uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps);

    function getEffectiveFees()
        external
        view
        returns (uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps);
}
