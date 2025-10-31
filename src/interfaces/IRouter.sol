// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

interface IRouterInitializer {
    function initialize(uint256 findPrevPriceCount, uint256 maxMatchCount, uint256 cancelLimit) external;
}

interface IRouter {
    function isPair(address pair) external view returns (bool);
    function submitBuyMarket(address pair, uint256 amount, uint256 maxMatchCount) external payable;
    function submitSellMarket(address pair, uint256 amount, uint256 maxMatchCount) external payable;
}
