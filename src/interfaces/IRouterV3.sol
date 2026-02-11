// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

interface IRouterV3 {
    function initialize(uint256 findPrevPriceCount, uint256 maxMatchCount, uint256 cancelLimit) external;
    function isPair(address pair) external view returns (bool);
}
