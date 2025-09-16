// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ICrossDex {
    function pairToMarket(address pair) external view returns (address);
    function checkTickSizeRoles(address account) external view;
    function pairCreated(address pair) external;
}

interface ICrossDexV2 is ICrossDex {
    function pairImpl() external view returns (address);
    function MAKER_VAULT_FACTORY() external view returns (address);
}
