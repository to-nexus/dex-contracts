// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IMarketInitializer {
    function initialize(address owner, address router, address feeCollector, address quote, address pairImpl)
        external;
}

interface IMarket {
    function baseToPair(address base) external view returns (address);
}
