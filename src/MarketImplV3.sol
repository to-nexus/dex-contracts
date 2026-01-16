// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.5.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.5.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Create2} from "@openzeppelin-contracts-5.5.0/utils/Create2.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";
import {EnumerableMap} from "@openzeppelin-contracts-5.5.0/utils/structs/EnumerableMap.sol";

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.5.0/access/OwnableUpgradeable.sol";

import {PairImplV3} from "./PairImplV3.sol";
import {ICrossDexV3} from "./interfaces/ICrossDexV3.sol";
import {IMarketV3} from "./interfaces/IMarketV3.sol";

contract MarketImplV3 is IMarketV3, UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableMap for EnumerableMap.AddressToAddressMap;

    error MarketInvalidInitializeData(bytes32);
    error MarketInvalidBaseAddress(address);
    error MarketAlreadyCreatedBaseAddress(address);
    error MarketDeployPair();
    error MarketInvalidFeeStructure(uint32 makerFee, uint32 takerFee);

    event PairCreated(address indexed pair, address indexed base, uint256 timestamp);
    event MarketFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee);
    event PairImplSet(address indexed before, address indexed current);
    event FeeControllerUpdated(address indexed before, address indexed current);

    uint256 public deployed; // immutable
    ICrossDexV3 public CROSS_DEX; // immutable
    address public QUOTE; // immutable
    address public ROUTER; // immutable

    address public pairImpl;

    address public override feeController;
    uint32 private _emptySlot;

    EnumerableMap.AddressToAddressMap private _allPairs; // base => pair

    uint256[41] private __gap;

    constructor() {
        _disableInitializers();
    }

    // Initialize with 4 different fee rates encoded in bytes data
    function initialize(address _owner, address _router, address _quote, address _pairImpl, address _feeController)
        external
        override
        initializer
    {
        __Ownable_init(_owner);

        if (_owner == address(0)) revert MarketInvalidInitializeData("owner");
        if (_router == address(0)) revert MarketInvalidInitializeData("router");
        if (_quote == address(0)) revert MarketInvalidInitializeData("quote");
        if (_pairImpl == address(0)) revert MarketInvalidInitializeData("pairImpl");
        if (_feeController == address(0)) revert MarketInvalidInitializeData("feeController");

        deployed = block.number;
        CROSS_DEX = ICrossDexV3(_msgSender());
        QUOTE = _quote;
        ROUTER = _router;
        pairImpl = _pairImpl;
        feeController = _feeController;
    }

    function allPairs() external view returns (address[] memory bases, address[] memory pairs) {
        uint256 length = _allPairs.length();
        bases = new address[](length);
        pairs = new address[](length);
        for (uint256 i = 0; i < length; ++i) {
            (bases[i], pairs[i]) = _allPairs.at(i);
        }
    }

    function checkTickSizeRoles(address account) external view override {
        // check account is tick size setter
        CROSS_DEX.checkTickSizeRoles(account);
    }

    function checkFeeControllerAllowed(address _feeController) external view override {
        CROSS_DEX.checkFeeControllerAllowed(_feeController);
    }

    function baseToPair(address base) external view returns (address) {
        return _allPairs.get(base);
    }

    function createPair(address base, uint256 tickSize, uint256 lotSize, bytes memory feeControllerInitData)
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
                    PairImplV3.initialize,
                    (ROUTER, QUOTE, base, tickSize, lotSize, feeController, feeControllerInitData)
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

    function setPairImpl(address _pairImpl) external onlyOwner {
        if (_pairImpl == address(0)) revert MarketInvalidInitializeData("pairImpl");
        emit PairImplSet(pairImpl, _pairImpl);
        pairImpl = _pairImpl;
    }

    function setFeeController(
        uint256 startIndex,
        uint256 endIndex,
        bool isForce,
        address newFeeController,
        bytes memory feeControllerInitData
    ) external onlyOwner {
        CROSS_DEX.checkFeeControllerAllowed(newFeeController);
        if (feeController != newFeeController) {
            emit FeeControllerUpdated(feeController, newFeeController);
            feeController = newFeeController;
        }
        endIndex = Math.min(endIndex, _allPairs.length());

        for (uint256 i = startIndex; i < endIndex; ++i) {
            (, address pair) = _allPairs.at(i);
            PairImplV3 PAIR = PairImplV3(pair);
            if (!isForce) if (address(PAIR.feeController()) != newFeeController) continue;
            PAIR.setFeeController(newFeeController, feeControllerInitData);
        }
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
