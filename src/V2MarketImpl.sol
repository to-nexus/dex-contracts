// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Create2} from "@openzeppelin-contracts-5.2.0/utils/Create2.sol";
import {EnumerableMap} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableMap.sol";

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/access/OwnableUpgradeable.sol";

import {V2PairImpl} from "./V2PairImpl.sol";
import {ICrossDexV2} from "./interfaces/ICrossDex.sol";
import {IMarket, IMarketInitializer} from "./interfaces/IMarket.sol";

contract V2MarketImpl is IMarket, IMarketInitializer, UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableMap for EnumerableMap.AddressToAddressMap;

    error MarketInvalidInitializeData(bytes32);
    error MarketInvalidBaseAddress(address);
    error MarketAlreadyCreatedBaseAddress(address);
    error MarketDeployPair();

    event PairCreated(address indexed pair, address indexed base, uint256 timestamp);
    event FeeCollectorChanged(address indexed before, address indexed current);
    event FeeBpsChanged(uint256 indexed before, uint256 indexed current);

    uint256 public deployed; // immutable
    ICrossDexV2 public CROSS_DEX; // immutable
    address public QUOTE; // immutable
    address public ROUTER; // immutable

    address public pairImpl;

    address public override feeCollector;
    uint32 public override feeBps; // BPS: basis point (1/10000)

    EnumerableMap.AddressToAddressMap private _allPairs; // base => pair

    uint256[41] private __gap;

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address _owner,
        address _router,
        address _quote,
        address _pairImpl,
        address _feeCollector,
        uint256 _feeBPS
    ) external override initializer {
        __Ownable_init(_owner);

        if (_owner == address(0)) revert MarketInvalidInitializeData("owner");
        if (_router == address(0)) revert MarketInvalidInitializeData("router");
        if (_quote == address(0)) revert MarketInvalidInitializeData("quote");
        if (_pairImpl == address(0)) revert MarketInvalidInitializeData("pairImpl");
        if (_feeCollector == address(0)) revert MarketInvalidInitializeData("feeCollector");
        if (_feeBPS > type(uint32).max) revert MarketInvalidInitializeData("feeBPS");

        deployed = block.number;
        CROSS_DEX = ICrossDexV2(_msgSender());
        QUOTE = _quote;
        ROUTER = _router;
        pairImpl = _pairImpl;

        feeCollector = _feeCollector;
        feeBps = uint32(_feeBPS);
    }

    function reinitialize() external reinitializer(2) onlyOwner {
        pairImpl = CROSS_DEX.pairImpl();
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
        // check account is tick size setter
        CROSS_DEX.checkTickSizeRoles(account);
    }

    function baseToPair(address base) external view returns (address) {
        return _allPairs.get(base);
    }

    function createPair(address base, uint256 tickSize, uint256 lotSize) external onlyOwner returns (address) {
        if (base == address(0) || base == address(QUOTE)) revert MarketInvalidBaseAddress(base);
        uint256 baseDecimals = IERC20Metadata(base).decimals();
        if (baseDecimals == 0) revert MarketInvalidBaseAddress(base);
        if (_allPairs.contains(base)) revert MarketAlreadyCreatedBaseAddress(base);

        bytes memory bytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(pairImpl, abi.encodeCall(V2PairImpl.initialize, (ROUTER, QUOTE, base, tickSize, lotSize)))
        );
        bytes32 salt = keccak256(abi.encodePacked(base));
        address pair = Create2.deploy(0, salt, bytecode);

        if (pair == address(0)) revert MarketDeployPair();
        if (!_allPairs.set(base, pair)) revert MarketAlreadyCreatedBaseAddress(base);

        CROSS_DEX.pairCreated(pair);
        emit PairCreated(pair, base, block.timestamp);
        return pair;
    }

    function setFeeCollector(address _feeCollector) external onlyOwner {
        if (_feeCollector == address(0)) revert MarketInvalidInitializeData("feeCollector");
        emit FeeCollectorChanged(feeCollector, _feeCollector);
        feeCollector = _feeCollector;
    }

    function setFeeBps(uint256 _feeBps) external onlyOwner {
        if (_feeBps > type(uint32).max) revert MarketInvalidInitializeData("feeBPS");
        emit FeeBpsChanged(feeBps, _feeBps);
        feeBps = uint32(_feeBps);
    }

    function updatePairImpl(address _newPairImpl) external onlyOwner {
        if (_newPairImpl == address(0)) revert MarketInvalidInitializeData("pairImpl");
        pairImpl = _newPairImpl;
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
