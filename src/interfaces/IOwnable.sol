// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IOwnable {
    error OwnableUnauthorizedAccount(address account);

    function owner() external view returns (address);
}
