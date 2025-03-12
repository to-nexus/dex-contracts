// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IMarketInitializer {
    function QUOTE() external view returns (address);
    function initialize(address owner, address router, address feeCollector, address quote, address pairImpl)
        external;
}

interface IMarket {
    function checkTickSizeRoles(address account) external view;
}
