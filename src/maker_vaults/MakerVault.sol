// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {IERC20, SafeERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin-contracts-5.2.0/utils/Address.sol";

contract MakerVault {
    event Claimed(address indexed account, address[] erc20s, bool withCoin);

    address payable public immutable ACCOUNT;

    constructor(address account_) {
        ACCOUNT = payable(account_);
    }

    receive() external payable {}

    function balanceOf(address[] calldata erc20s) external view returns (uint256[] memory, uint256) {
        uint256 length = erc20s.length;
        uint256[] memory balances = new uint256[](length);
        address vault = address(this);
        unchecked {
            for (uint256 i = 0; i < length; ++i) {
                IERC20 erc20 = IERC20(erc20s[i]);
                balances[i] = erc20.balanceOf(vault);
            }
        }
        return (balances, vault.balance);
    }

    function claim(address[] calldata erc20s, bool withCoin) external {
        // send all ERC20 tokens to the account
        uint256 length = erc20s.length;
        address vault = address(this);
        unchecked {
            for (uint256 i = 0; i < length; ++i) {
                IERC20 erc20 = IERC20(erc20s[i]);
                uint256 _balance = erc20.balanceOf(vault);
                if (_balance != 0) SafeERC20.safeTransfer(erc20, ACCOUNT, _balance);
            }
        }
        if (withCoin) {
            // send all ETH to the account
            uint256 _balance = vault.balance;
            if (_balance != 0) Address.sendValue(ACCOUNT, _balance);
        }
        emit Claimed(ACCOUNT, erc20s, withCoin);
    }
}
