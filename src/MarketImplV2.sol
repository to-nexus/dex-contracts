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
import {BPS_DENOMINATOR, IMarketInitializer, IMarketV2, NO_FEE_BPS} from "./interfaces/IMarket.sol";

contract MarketImplV2 is IMarketV2, UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableMap for EnumerableMap.AddressToAddressMap;

    error MarketInvalidInitializeData(bytes32);
    error MarketInvalidBaseAddress(address);
    error MarketAlreadyCreatedBaseAddress(address);
    error MarketDeployPair();

    event PairCreated(address indexed pair, address indexed base, uint256 timestamp);
    event FeeCollectorChanged(address indexed before, address indexed current);
    event MarketFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee);

    uint256 public deployed; // immutable
    ICrossDex public CROSS_DEX; // immutable
    address public QUOTE; // immutable
    address public ROUTER; // immutable

    address public pairImpl;

    address public override feeCollector;
    FeeConfig private _feeConfig;

    EnumerableMap.AddressToAddressMap private _allPairs; // base => pair

    uint256[38] private __gap;

    constructor() {
        _disableInitializers();
    }

    // Initialize with 4 different fee rates encoded in bytes data
    function initialize(
        address _owner,
        address _router,
        address _quote,
        address _pairImpl,
        address _feeCollector,
        bytes memory feeData
    ) external override initializer {
        __Ownable_init(_owner);

        if (_owner == address(0)) revert MarketInvalidInitializeData("owner");
        if (_router == address(0)) revert MarketInvalidInitializeData("router");
        if (_quote == address(0)) revert MarketInvalidInitializeData("quote");
        if (_pairImpl == address(0)) revert MarketInvalidInitializeData("pairImpl");
        if (_feeCollector == address(0)) revert MarketInvalidInitializeData("feeCollector");

        // Decode 4 different fee rates
        (uint32 _sellerMakerFeeBps, uint32 _sellerTakerFeeBps, uint32 _buyerMakerFeeBps, uint32 _buyerTakerFeeBps) =
            abi.decode(feeData, (uint32, uint32, uint32, uint32));

        // Validate fee rates
        if (_sellerMakerFeeBps != NO_FEE_BPS && _sellerMakerFeeBps >= BPS_DENOMINATOR) {
            revert MarketInvalidInitializeData("sellerMakerFeeBps");
        }
        if (_sellerTakerFeeBps != NO_FEE_BPS && _sellerTakerFeeBps >= BPS_DENOMINATOR) {
            revert MarketInvalidInitializeData("sellerTakerFeeBps");
        }
        if (_buyerMakerFeeBps != NO_FEE_BPS && _buyerMakerFeeBps >= BPS_DENOMINATOR) {
            revert MarketInvalidInitializeData("buyerMakerFeeBps");
        }
        if (_buyerTakerFeeBps != NO_FEE_BPS && _buyerTakerFeeBps >= BPS_DENOMINATOR) {
            revert MarketInvalidInitializeData("buyerTakerFeeBps");
        }

        deployed = block.number;
        CROSS_DEX = ICrossDex(_msgSender());
        QUOTE = _quote;
        ROUTER = _router;
        pairImpl = _pairImpl;

        feeCollector = _feeCollector;
        _feeConfig.sellerMakerFeeBps = _sellerMakerFeeBps; // Seller Maker fee
        _feeConfig.sellerTakerFeeBps = _sellerTakerFeeBps; // Seller Taker fee
        _feeConfig.buyerMakerFeeBps = _buyerMakerFeeBps; // Buyer Maker fee
        _feeConfig.buyerTakerFeeBps = _buyerTakerFeeBps; // Buyer Taker fee
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

    function getFeeConfig() external view override returns (FeeConfig memory) {
        return _feeConfig;
    }

    function baseToPair(address base) external view returns (address) {
        return _allPairs.get(base);
    }

    function createPair(
        address base,
        uint256 tickSize,
        uint256 lotSize,
        uint32 _sellerMakerFeeBps,
        uint32 _sellerTakerFeeBps,
        uint32 _buyerMakerFeeBps,
        uint32 _buyerTakerFeeBps
    ) external onlyOwner returns (address) {
        if (base == address(0) || base == address(QUOTE)) revert MarketInvalidBaseAddress(base);
        uint256 baseDecimals = IERC20Metadata(base).decimals();
        if (baseDecimals == 0) revert MarketInvalidBaseAddress(base);
        if (_allPairs.contains(base)) revert MarketAlreadyCreatedBaseAddress(base);

        // Encode 4 different fee rates
        bytes memory feeData = abi.encode(_sellerMakerFeeBps, _sellerTakerFeeBps, _buyerMakerFeeBps, _buyerTakerFeeBps);

        bytes memory bytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(
                pairImpl, abi.encodeCall(PairImplV2.initialize, (ROUTER, QUOTE, base, tickSize, lotSize, feeData))
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

    function setMarketFees(
        uint32 _sellerMakerFeeBps,
        uint32 _sellerTakerFeeBps,
        uint32 _buyerMakerFeeBps,
        uint32 _buyerTakerFeeBps
    ) external onlyOwner {
        if (_sellerMakerFeeBps != NO_FEE_BPS && _sellerMakerFeeBps >= BPS_DENOMINATOR) {
            revert MarketInvalidInitializeData("sellerMakerFeeBps");
        }
        if (_sellerTakerFeeBps != NO_FEE_BPS && _sellerTakerFeeBps >= BPS_DENOMINATOR) {
            revert MarketInvalidInitializeData("sellerTakerFeeBps");
        }
        if (_buyerMakerFeeBps != NO_FEE_BPS && _buyerMakerFeeBps >= BPS_DENOMINATOR) {
            revert MarketInvalidInitializeData("buyerMakerFeeBps");
        }
        if (_buyerTakerFeeBps != NO_FEE_BPS && _buyerTakerFeeBps >= BPS_DENOMINATOR) {
            revert MarketInvalidInitializeData("buyerTakerFeeBps");
        }

        emit MarketFeesUpdated(_sellerMakerFeeBps, _sellerTakerFeeBps, _buyerMakerFeeBps, _buyerTakerFeeBps);
        _feeConfig.sellerMakerFeeBps = _sellerMakerFeeBps; // feeBps represents seller maker fee
        _feeConfig.sellerTakerFeeBps = _sellerTakerFeeBps; // seller taker fee
        _feeConfig.buyerMakerFeeBps = _buyerMakerFeeBps; // buyer maker fee
        _feeConfig.buyerTakerFeeBps = _buyerTakerFeeBps; // buyer taker fee
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
