// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IRouter {
    function isPair(address pair) external view returns (bool);
    function addPair(address pair) external;
}
