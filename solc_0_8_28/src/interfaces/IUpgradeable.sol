// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

interface IUpgradeable {
    function upgradeToAndCall(address newImplementation, bytes memory data) external payable;
    function transferOwnership(address newOwner) external;
}
