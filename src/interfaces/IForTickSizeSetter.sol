// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// Contract interfaces used by TickSizeSetter to configure tick sizes.

interface ICrossDex {
    function allMarkets() external view returns (address[] memory, address[] memory);
}

interface IMarket {
    function QUOTE() external view returns (address);
    function allPairs() external view returns (address[] memory, address[] memory);
}

interface IPair {
    function matchedPrice() external view returns (uint256);
    function setTickSize(uint256 _lotSize, uint256 _tickSize) external;
}
