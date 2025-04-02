// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IRouterInitializer {
    function initialize(uint256 findPrevPriceCount, uint256 maxMatchCount) external;
}

interface IRouter {
    function isPair(address pair) external view returns (bool);
}
