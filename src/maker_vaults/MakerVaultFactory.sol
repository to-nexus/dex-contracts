// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {Create2} from "@openzeppelin-contracts-5.2.0/utils/Create2.sol";

import {IMakerVaultFactory} from "../interfaces/IMakerVaultFactory.sol";
import {MakerVault} from "./MakerVault.sol";

contract MakerVaultFactory is IMakerVaultFactory {
    event Created(address indexed account, address vault);

    bytes public constant MAKER_VAULT_BYTECODE = type(MakerVault).creationCode;

    function makerVaultAddress(address account) public view override returns (address) {
        bytes32 salt = bytes32(bytes20(account));
        return Create2.computeAddress(salt, keccak256(_packCreationCode(account)));
    }

    function ensureMakerVault(address account) external returns (address vault) {
        bytes32 salt = bytes32(bytes20(account));
        vault = makerVaultAddress(account);
        if (vault.code.length == 0) {
            vault = Create2.deploy(0, salt, _packCreationCode(account));
            emit Created(account, vault);
        }
    }

    function _packCreationCode(address account) private pure returns (bytes memory) {
        return abi.encodePacked(MAKER_VAULT_BYTECODE, abi.encode(account));
    }
}
