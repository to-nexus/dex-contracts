// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ITickSizeSetter {
    function updateInterval() external view returns (uint256);
    function calcTimestamp(uint256 timestamp) external view returns (uint256);
}

// Contract interfaces used by TickSizeSetter to configure tick sizes.

interface ICrossDexForTickSizeSetter {
    function ROUTER() external view returns (address);
    function allMarkets() external view returns (address[] memory, address[] memory);
    function pairToMarket(address pair) external view returns (address);
    function isMarket(address market) external view returns (bool);
}

interface IRouterForTickSizeSetter {
    function matchedPriceAt(address pair, uint256 timestamp) external view returns (uint256);
}

interface IMarketForTickSizeSetter {
    function QUOTE() external view returns (address);
    function allPairs() external view returns (address[] memory, address[] memory);
}

interface IPairForTickSizeSetter {
    function matchedPrice() external view returns (uint256);
    function paused() external view returns (bool);
    function tickSizes() external view returns (uint256 tick, uint256 lot);
    function setTickSize(uint256 _lotSize, uint256 _tickSize) external;
}
