// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Create2} from "@openzeppelin-contracts-5.2.0/utils/Create2.sol";
import {EnumerableMap} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableMap.sol";

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/access/OwnableUpgradeable.sol";

import {PairImplV2} from "./PairImplV2.sol";
import {ICrossDex} from "./interfaces/ICrossDex.sol";
import {IMarketInitializer, IMarketV2, NO_FEE_BPS} from "./interfaces/IMarket.sol";

contract MarketImplV2 is IMarketV2, IMarketInitializer, UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableMap for EnumerableMap.AddressToAddressMap;

    error MarketInvalidInitializeData(bytes32);
    error MarketInvalidBaseAddress(address);
    error MarketAlreadyCreatedBaseAddress(address);
    error MarketDeployPair();

    event PairCreated(address indexed pair, address indexed base, uint256 timestamp);
    event FeeCollectorChanged(address indexed before, address indexed current);
    event MarketFeesUpdated(uint32 indexed makerFee, uint32 indexed takerFee);

    uint256 public deployed; // immutable
    ICrossDex public CROSS_DEX; // immutable
    address public QUOTE; // immutable
    address public ROUTER; // immutable

    address public pairImpl;

    address public override feeCollector;
    uint32 public override makerFeeBps; // Maker fee (backward compatibility)

    EnumerableMap.AddressToAddressMap private _allPairs; // base => pair

    uint32 public override takerFeeBps; // Taker fee (new)

    uint256[40] private __gap;

    constructor() {
        _disableInitializers();
    }

    // 기존의 interface 를 변경하지 않기 위해 feeBps 를 maker, taker fee 로 동일하게 설정
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
        if (_feeBPS > 10000) revert MarketInvalidInitializeData("feeBps");

        deployed = block.number;
        CROSS_DEX = ICrossDex(_msgSender());
        QUOTE = _quote;
        ROUTER = _router;
        pairImpl = _pairImpl;

        feeCollector = _feeCollector;
        makerFeeBps = takerFeeBps = uint32(_feeBPS);
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

    function createPair(address base, uint256 tickSize, uint256 lotSize, uint32 _makerFeeBps, uint32 _takerFeeBps)
        external
        onlyOwner
        returns (address)
    {
        if (base == address(0) || base == address(QUOTE)) revert MarketInvalidBaseAddress(base);
        uint256 baseDecimals = IERC20Metadata(base).decimals();
        if (baseDecimals == 0) revert MarketInvalidBaseAddress(base);
        if (_allPairs.contains(base)) revert MarketAlreadyCreatedBaseAddress(base);

        bytes memory bytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(
                pairImpl,
                abi.encodeCall(
                    PairImplV2.initialize, (ROUTER, QUOTE, base, tickSize, lotSize, _makerFeeBps, _takerFeeBps)
                )
            )
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

    function setMarketFees(uint32 _makerFeeBps, uint32 _takerFeeBps) external onlyOwner {
        if (_makerFeeBps != NO_FEE_BPS && _makerFeeBps >= 10000) revert MarketInvalidInitializeData("makerFeeBPS");
        if (_takerFeeBps != NO_FEE_BPS && _takerFeeBps >= 10000) revert MarketInvalidInitializeData("takerFeeBPS");

        emit MarketFeesUpdated(_makerFeeBps, _takerFeeBps);
        makerFeeBps = _makerFeeBps; // feeBps represents maker fee
        takerFeeBps = _takerFeeBps;
    }

    // New getter functions for fee information
    function getMarketFees() external view returns (uint32 makerFee, uint32 takerFee) {
        return (makerFeeBps, takerFeeBps);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
