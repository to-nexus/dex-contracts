// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/IERC20.sol";

interface IPair {
    enum OrderType {
        NONE,
        SELL,
        BUY
    }

    enum CloseType {
        ALL_MATCH,
        MARKET,
        CANCEL
    }

    struct Order {
        OrderType _type;
        address owner;
        uint32 feePermile;
        uint256 price;
        uint256 amount;
    }

    function QUOTE() external view returns (IERC20);
    function BASE() external view returns (IERC20);
    function DENOMINATOR() external view returns (uint256);

    function limit(Order memory order, uint256 searchPrice, uint256 maxMatchCount) external returns (uint256 orderId);
    function market(Order memory order, uint256 spendAmount, uint256 maxMatchCount) external;
    function cancel(address caller, uint256[] memory orderIds) external;
}
