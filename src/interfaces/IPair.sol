// SPDX-License-Identifier: UNLICENSED
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

    function limit(Order memory order, LimitConstraints constraints, uint256[2] memory adjacent, uint256 maxMatchCount)
        external
        returns (uint256 orderId);
    function market(Order memory order, uint256 spendAmount, uint256 maxMatchCount) external;
    function cancel(address caller, uint256[] memory orderIds) external;
}
