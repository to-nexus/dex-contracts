// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

interface IMarketInitializer {
    function QUOTE() external view returns (address);
    function initialize(
        address owner,
        address router,
        address quote,
        address pairImpl,
        address feeCollector,
        uint256 feeBps
    ) external;
}

interface IMarket {
    function feeBps() external view returns (uint32);
    function feeCollector() external view returns (address);
    function checkTickSizeRoles(address account) external view;
    function createPair(address base, uint256 tickSize, uint256 lotSize) external returns (address);
    function setFeeCollector(address _feeCollector) external;
    function setFeeBps(uint256 _feeBps) external;
}

uint32 constant NO_FEE_BPS = type(uint32).max; // Special value to indicate "use market fee"
uint32 constant BPS_DENOMINATOR = 10000; // Basis points denominator (100%)

interface IMarketV2 {
    function QUOTE() external view returns (address);
    function initialize(
        address owner,
        address router,
        address quote,
        address pairImpl,
        address feeCollector,
        bytes memory feeData
    ) external;

    struct FeeConfig {
        uint32 sellerMakerFeeBps; // Seller Maker fee (BPS)
        uint32 sellerTakerFeeBps; // Seller Taker fee (BPS)
        uint32 buyerMakerFeeBps; // Buyer Maker fee (BPS)
        uint32 buyerTakerFeeBps; // Buyer Taker fee (BPS)
    }

    function feeCollector() external view returns (address);
    function checkTickSizeRoles(address account) external view;
    function getFeeConfig() external view returns (FeeConfig memory);

    function createPair(address base, uint256 tickSize, uint256 lotSize, bytes memory feeData)
        external
        returns (address);
}
