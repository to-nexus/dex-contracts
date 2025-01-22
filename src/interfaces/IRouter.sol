// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IRouterInitializer {
    function initialize(address owner, uint256 maxMatchCount) external;
}

interface IRouter {
    function isPair(address pair) external view returns (bool);
}
