// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

interface ICrossDex {
    function pairToMarket(address pair) external view returns (address);
    function checkTickSizeRoles(address account) external view;
    function pairCreated(address pair) external;
}
