// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IMarketInitializer {
    function QUOTE() external view returns (address);
    function initialize(
        address owner,
        address router,
        address quote,
        address pairImpl,
        address feeCollector,
        uint256 feeBps
    ) external;
}

interface IMarket {
    function feeBps() external view returns (uint32);
    function feeCollector() external view returns (address);
    function checkTickSizeRoles(address account) external view;
}

interface IMarketInitializerV2 is IMarketInitializer {
    function updatePairImpl(address) external;
}
