// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {EnumerableMap} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableMap.sol";

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/access/OwnableUpgradeable.sol";

import {PairImpl} from "./PairImpl.sol";
import {ICrossDex} from "./interfaces/ICrossDex.sol";
import {IMarket, IMarketInitializer} from "./interfaces/IMarket.sol";

contract MarketImpl is IMarket, IMarketInitializer, UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableMap for EnumerableMap.AddressToAddressMap;

    error MarketInvalidInitializeData(bytes32);
    error MarketInvalidBaseAddress(address);
    error MarketAlreadyCreatedBaseAddress(address);
    error MarketDeployPair();

    event PairCreated(address indexed pair, address indexed base, uint256 timestamp);

    uint256 public deployed; // immutable
    ICrossDex public CROSS_DEX; // immutable
    address public QUOTE; // immutable

    address public router;
    address public feeCollector;
    address public pairImpl;

    EnumerableMap.AddressToAddressMap private _allPairs; // base => pair

    uint256[43] private __gap;

    constructor() {
        _disableInitializers();
    }

    function initialize(address _owner, address _router, address _feeCollector, address _quote, address _pairImpl)
        external
        override
        initializer
    {
        if (_owner == address(0)) revert MarketInvalidInitializeData("owner");
        if (_router == address(0)) revert MarketInvalidInitializeData("router");
        if (_feeCollector == address(0)) revert MarketInvalidInitializeData("feeCollector");
        if (_quote == address(0)) revert MarketInvalidInitializeData("quote");
        if (_pairImpl == address(0)) revert MarketInvalidInitializeData("pairImpl");

        deployed = block.number;
        CROSS_DEX = ICrossDex(_msgSender());
        QUOTE = _quote;

        router = _router;
        feeCollector = _feeCollector;
        pairImpl = _pairImpl;

        __Ownable_init(_owner);
    }

    function allPairs() external view returns (address[] memory bases, address[] memory pairs) {
        uint256 length = _allPairs.length();
        bases = new address[](length);
        pairs = new address[](length);
        for (uint256 i = 0; i < length;) {
            (bases[i], pairs[i]) = _allPairs.at(i);
            unchecked {
                ++i;
            }
        }
    }

    function checkTickSizeRoles(address account) external view override {
        // check account is owner or tick size setter
        if (account != owner()) CROSS_DEX.checkTickSizeRoles(account);
    }

    function baseToPair(address base) external view returns (address) {
        return _allPairs.get(base);
    }

    function createPair(address base, uint256 quoteTickSize, uint256 baseTickSize, uint256 feeBps)
        external
        onlyOwner
        returns (address pair)
    {
        if (base == address(0) || base == address(QUOTE)) revert MarketInvalidBaseAddress(base);
        uint256 baseDecimals = IERC20Metadata(base).decimals();
        if (baseDecimals == 0) revert MarketInvalidBaseAddress(base);
        if (_allPairs.contains(base)) revert MarketAlreadyCreatedBaseAddress(base);

        pair = address(
            new ERC1967Proxy(
                pairImpl,
                abi.encodeWithSelector(
                    PairImpl.initialize.selector, router, QUOTE, base, quoteTickSize, baseTickSize, feeCollector, feeBps
                )
            )
        );
        if (pair == address(0)) revert MarketDeployPair();
        if (!_allPairs.set(base, pair)) revert MarketAlreadyCreatedBaseAddress(base);

        CROSS_DEX.pairCreated(pair);
        emit PairCreated(pair, base, block.timestamp);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
