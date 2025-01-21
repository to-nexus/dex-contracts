// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableSet.sol";

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/access/OwnableUpgradeable.sol";

import {PairImpl} from "./PairImpl.sol";
import {ICrossDex} from "./interfaces/ICrossDex.sol";

contract MarketImpl is UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error MarketInvalidInitializeData(bytes32);
    error MarketInvalidBaseAddress(address);
    error MarketAlreadyCreatedBaseAddress(address);
    error MarketDeployPair();

    event PairCreated(address indexed pair, address indexed base, uint256 timestamp);

    ICrossDex public CROSS_DEX; // immutable
    address public QUOTE; // immutable

    address public router;
    address public feeCollector;
    address public pairImpl;

    EnumerableSet.AddressSet private _allBases;
    mapping(address erc20 => address) private _allPairs;

    uint256[44] private __gap;

    function initialize(address _owner, address _router, address _feeCollector, address _quote, address _pairImpl)
        external
        initializer
    {
        if (_owner == address(0)) revert MarketInvalidInitializeData("owner");
        if (_router == address(0)) revert MarketInvalidInitializeData("router");
        if (_feeCollector == address(0)) revert MarketInvalidInitializeData("feeCollector");
        if (_quote == address(0)) revert MarketInvalidInitializeData("quote");
        if (_pairImpl == address(0)) revert MarketInvalidInitializeData("pairImpl");

        CROSS_DEX = ICrossDex(_msgSender());
        QUOTE = _quote;

        router = _router;
        feeCollector = _feeCollector;
        pairImpl = _pairImpl;

        __Ownable_init(_owner);
    }

    function allPairs() external view returns (address[] memory bases, address[] memory pairs) {
        bases = _allBases.values();
        uint256 length = bases.length;
        pairs = new address[](length);
        for (uint256 i = 0; i < length;) {
            pairs[i] = _allPairs[bases[i]];
            unchecked {
                ++i;
            }
        }
    }

    function pairByBase(address base) external view returns (address) {
        return _allPairs[base];
    }

    function createPair(
        address base,
        uint256 quoteTickSize,
        uint256 baseTickSize,
        uint256 makerFeePermil,
        uint256 takerFeePermil
    ) external onlyOwner returns (address pair) {
        if (base == address(0)) revert MarketInvalidBaseAddress(base);
        if (!_allBases.add(base)) revert MarketAlreadyCreatedBaseAddress(base);
        uint256 baseDecimals = IERC20Metadata(base).decimals();
        if (baseDecimals == 0) revert MarketInvalidBaseAddress(base);

        pair = address(
            new ERC1967Proxy(
                pairImpl,
                abi.encodeWithSelector(
                    PairImpl.initialize.selector,
                    owner(),
                    router,
                    QUOTE,
                    base,
                    quoteTickSize,
                    baseTickSize,
                    feeCollector,
                    makerFeePermil,
                    takerFeePermil
                )
            )
        );
        if (pair == address(0)) revert MarketDeployPair();
        _allPairs[base] = pair;

        CROSS_DEX.notifyPairCreated(pair);
        emit PairCreated(pair, base, block.timestamp);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
