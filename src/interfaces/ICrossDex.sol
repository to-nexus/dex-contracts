// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ICrossDex {
    function quoteToMarket(address quote) external view returns (address);
}
