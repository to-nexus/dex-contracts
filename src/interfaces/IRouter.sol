// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IRouterInitializer {
    function initialize(uint256 findPrevPriceCount, uint256 maxMatchCount, uint256 cancelLimit) external;
}

interface IRouter {
    function isPair(address pair) external view returns (bool);
}

interface IRouterInitializerV2 {
    function initialize(
        uint256 findPrevPriceCount,
        uint256 maxMatchCount,
        uint256 cancelLimit,
        address makerVaultFactory
    ) external;
}

interface IRouterV2 is IRouter {
    function makerVaultAddress(address account) external view returns (address);
}
