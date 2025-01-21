// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ICrossDex {
    function notifyPairCreated(address pair) external;
}
