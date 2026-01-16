// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

uint32 constant NO_FEE_BPS = type(uint32).max; // Special value to indicate "use market fee"
uint32 constant BPS_DENOMINATOR = 10000; // Basis points denominator (100%)

interface IMarketV3 {
    function QUOTE() external view returns (address);
    function initialize(address owner, address router, address quote, address pairImpl, address feeController) external;

    function feeController() external view returns (address);
    function checkTickSizeRoles(address account) external view;
    function checkFeeControllerAllowed(address feeController) external view;

    function createPair(address base, uint256 tickSize, uint256 lotSize, bytes memory feeControllerInitData)
        external
        returns (address);
}
