// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ICrossDex {
    function pairToMarket(address pair) external view returns (address);
    function pairCreated(address pair) external;
}
