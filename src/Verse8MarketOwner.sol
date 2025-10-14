// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin-contracts-5.2.0/access/Ownable.sol";

import {IMarket} from "./interfaces/IMarket.sol";
import {IUpgradeable} from "./interfaces/IUpgradeable.sol";

/**
 * @title Verse8MarketOwner
 * @notice A contract that manages MarketImpl ownership with separated permissions
 * @dev Delegates createPair authority to a separate pairCreator,
 *      while the owner retains all other permissions.
 *      This contract is non-upgradeable.
 */
contract Verse8MarketOwner is Ownable {
    error Verse8MarketOwnerInvalidMarketAddress(address);
    error Verse8MarketOwnerInvalidPairCreatorAddress(address);
    error Verse8MarketOwnerUnauthorizedPairCreator(address);

    event MarketSet(address indexed before, address indexed current);
    event PairCreatorSet(address indexed before, address indexed current);
    event PairCreated(
        address indexed market, address indexed base, address indexed pair, uint256 tickSize, uint256 lotSize
    );

    /// @notice MarketImpl contract address
    address public market;

    /// @notice Address with pair creation authority
    address public pairCreator;

    modifier onlyPairCreatorOrOwner() {
        address sender = _msgSender();
        if (sender != owner() && sender != pairCreator) revert Verse8MarketOwnerUnauthorizedPairCreator(sender);
        _;
    }

    /**
     * @notice Contract constructor
     * @param _owner Owner address (full permissions)
     * @param _market MarketImpl contract address
     * @param _pairCreator Address with pair creation authority
     */
    constructor(address _owner, address _market, address _pairCreator) Ownable(_owner) {
        if (_market == address(0)) revert Verse8MarketOwnerInvalidMarketAddress(_market);
        if (_pairCreator == address(0)) revert Verse8MarketOwnerInvalidPairCreatorAddress(_pairCreator);

        market = _market;
        pairCreator = _pairCreator;

        emit MarketSet(address(0), _market);
        emit PairCreatorSet(address(0), _pairCreator);
    }

    // ===== PAIR CREATOR OR OWNER FUNCTIONS =====

    /**
     * @notice Create a new trading pair
     * @dev Can only be called by pairCreator or owner
     * @param base BASE token address
     * @param tickSize Minimum price unit for QUOTE token
     * @param lotSize Minimum quantity unit for BASE token
     * @return pair Address of the created Pair contract
     */
    function createPair(address base, uint256 tickSize, uint256 lotSize)
        external
        onlyPairCreatorOrOwner
        returns (address pair)
    {
        pair = IMarket(market).createPair(base, tickSize, lotSize);
        emit PairCreated(market, base, pair, tickSize, lotSize);
    }

    // ===== OWNER ONLY FUNCTIONS =====

    /**
     * @notice Change the fee collector address
     * @dev Can only be called by owner
     * @param _feeCollector New fee collector address
     */
    function setFeeCollector(address _feeCollector) external onlyOwner {
        IMarket(market).setFeeCollector(_feeCollector);
    }

    /**
     * @notice Set the trading fee rate
     * @dev Can only be called by owner
     * @param _feeBps New fee rate (in BPS, 1 BPS = 0.01%)
     */
    function setFeeBps(uint256 _feeBps) external onlyOwner {
        IMarket(market).setFeeBps(_feeBps);
    }

    /**
     * @notice Upgrade MarketImpl contract to a new implementation
     * @dev Can only be called by owner
     * @param newImplementation New implementation contract address
     */
    function upgradeMarket(address newImplementation) external onlyOwner {
        IUpgradeable(market).upgradeToAndCall(newImplementation, "");
    }

    /**
     * @notice Upgrade MarketImpl contract to a new implementation and call initializer
     * @dev Can only be called by owner
     * @param newImplementation New implementation contract address
     * @param data Initialization function call data
     */
    function upgradeMarketAndCall(address newImplementation, bytes memory data) external onlyOwner {
        IUpgradeable(market).upgradeToAndCall(newImplementation, data);
    }

    /**
     * @notice Transfer MarketImpl ownership to another address
     * @dev Can only be called by owner, use for emergency or governance transition
     * @param newOwner New MarketImpl owner address
     */
    function transferMarketOwnership(address newOwner) external onlyOwner {
        IUpgradeable(market).transferOwnership(newOwner);
    }

    /**
     * @notice Change the pairCreator address
     * @dev Can only be called by owner
     * @param _pairCreator New pairCreator address
     */
    function setPairCreator(address _pairCreator) external onlyOwner {
        if (_pairCreator == address(0)) revert Verse8MarketOwnerInvalidPairCreatorAddress(_pairCreator);
        emit PairCreatorSet(pairCreator, _pairCreator);
        pairCreator = _pairCreator;
    }

    /**
     * @notice Change the market address
     * @dev Can only be called by owner, use when switching to a new MarketImpl
     * @param _market New MarketImpl address
     */
    function setMarket(address _market) external onlyOwner {
        if (_market == address(0)) revert Verse8MarketOwnerInvalidMarketAddress(_market);
        emit MarketSet(market, _market);
        market = _market;
    }

    // ===== VIEW FUNCTIONS =====

    /**
     * @notice Check if the given address has pair creation authority
     * @param account Address to check
     * @return Whether the address has authority
     */
    function canCreatePair(address account) external view returns (bool) {
        return account == owner() || account == pairCreator;
    }
}
