// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {Create2} from "@openzeppelin-contracts-5.2.0/utils/Create2.sol";

import {CrossDexMakerVault} from "./CrossDexMakerVault.sol";
import {ICrossDexMakerVaultFactory} from "./interfaces/ICrossDexMakerVaultFactory.sol";

contract CrossDexMakerVaultFactory is ICrossDexMakerVaultFactory {
    event Created(address indexed account, address vault);

    bytes public constant MAKER_VAULT_BYTECODE = type(CrossDexMakerVault).creationCode;
    bytes32 public constant MAKER_VAULT_BYTECODE_HASH = keccak256(MAKER_VAULT_BYTECODE);

    function makerVaultAddress(address account) public view override returns (address) {
        bytes32 salt = bytes32(bytes20(account));
        return Create2.computeAddress(salt, MAKER_VAULT_BYTECODE_HASH);
    }

    function ensureMakerVault(address account) external returns (address vault) {
        bytes32 salt = bytes32(bytes20(account));
        vault = Create2.computeAddress(salt, MAKER_VAULT_BYTECODE_HASH);
        if (vault.code.length == 0) {
            vault = Create2.deploy(0, salt, MAKER_VAULT_BYTECODE);
            emit Created(account, vault);
        }
    }
}
