// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.1.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.1.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.1.0/token/ERC20/extensions/IERC20Metadata.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.1.0/utils/structs/EnumerableSet.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.1.0/access/OwnableUpgradeable.sol";

import {Pair, PairImpl} from "./Pair.sol";

contract Factory is ERC1967Proxy {
    constructor(address implementation, address router, address feeCollector, address quote, address pairImpl)
        ERC1967Proxy(
            implementation,
            abi.encodeWithSelector(FactoryImpl.initialize.selector, router, feeCollector, quote, pairImpl)
        )
    {}
}

contract FactoryImpl is UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error FactoryInvalidInitializeData(bytes32);
    error FactoryInvalidBaseAddress(address);
    error FactoryAlreadCreatedBaseAddress(address);
    error FactoryDeployPair();

    address public QUOTE;

    address public router;
    address public feeCollector;
    address public pairImpl;

    EnumerableSet.AddressSet private _allBases;
    mapping(address erc20 => address) private _allPairs;

    function initialize(address _router, address _feeCollector, address _quote, address _pairImpl)
        external
        initializer
    {
        if (_router == address(0)) revert FactoryInvalidInitializeData("router");
        if (_feeCollector == address(0)) revert FactoryInvalidInitializeData("feeCollector");
        if (_quote == address(0)) revert FactoryInvalidInitializeData("quote");
        if (_pairImpl == address(0)) revert FactoryInvalidInitializeData("pairImpl");

        pairImpl = _pairImpl;
        router = _router;
        feeCollector = _feeCollector;
        QUOTE = _quote;

        __Ownable_init(_msgSender());
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
        uint256 quoteFeePermile,
        uint256 baseFeePermile
    ) external onlyOwner returns (address pair) {
        if (base == address(0)) revert FactoryInvalidBaseAddress(base);
        if (!_allBases.add(base)) revert FactoryAlreadCreatedBaseAddress(base);
        uint256 baseDecimals = IERC20Metadata(base).decimals();
        if (baseDecimals == 0) revert FactoryInvalidBaseAddress(base);

        pair = address(
            new Pair(
                pairImpl,
                owner(),
                router,
                QUOTE,
                base,
                quoteTickSize,
                baseTickSize,
                feeCollector,
                quoteFeePermile,
                baseFeePermile
            )
        );
        if (pair == address(0)) revert FactoryDeployPair();
        _allPairs[base] = pair;
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}

    uint256[44] __gap;
}
