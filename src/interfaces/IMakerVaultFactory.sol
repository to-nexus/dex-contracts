// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IMakerVaultFactory {
    function makerVaultAddress(address account) external view returns (address);
    function ensureMakerVault(address account) external returns (address vault);
}
