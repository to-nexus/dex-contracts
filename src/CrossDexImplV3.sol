// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.5.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.5.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Create2} from "@openzeppelin-contracts-5.5.0/utils/Create2.sol";
import {EnumerableMap} from "@openzeppelin-contracts-5.5.0/utils/structs/EnumerableMap.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.5.0/utils/structs/EnumerableSet.sol";

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.5.0/access/OwnableUpgradeable.sol";

import {ICrossDexV3} from "./interfaces/ICrossDexV3.sol";
import {IMarketV3} from "./interfaces/IMarketV3.sol";
import {IRouterV3} from "./interfaces/IRouterV3.sol";

import {WETH} from "./WETH.sol";

contract CrossDexImplV3 is ICrossDexV3, UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableMap for EnumerableMap.AddressToAddressMap;
    using EnumerableSet for EnumerableSet.AddressSet;

    error CrossDexInitializeData(bytes32);
    error CrossDexInvalidMarketAddress(address);
    error CrossDexUnauthorizedChangeTickSizes(address);
    error CrossDexInvalidTickSizeSetter(address current, address input);
    error CrossDexInvalidFeeController(address feeController);

    event MarketCreated(
        address indexed quote, address indexed market, address indexed owner, address fee_collector, string message
    );
    event TickSizeSetterSet(address indexed before, address indexed current);
    event PairImplSet(address indexed before, address indexed current);
    event MarketImplSet(address indexed before, address indexed current);
    event FeeControllerAllowed(address indexed feeController, bool indexed allowed);

    address payable public ROUTER; // immutable

    address public marketImpl;
    address public pairImpl;

    EnumerableMap.AddressToAddressMap private _allMarkets; // market => quote (update v2)
    mapping(address pair => address) public override pairToMarket;

    address public tickSizeSetter;
    EnumerableSet.AddressSet private _allowedFeeControllers;

    uint256[42] __gap;

    modifier onlyMarket() {
        if (!isMarket(_msgSender())) revert CrossDexInvalidMarketAddress(_msgSender());
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address _owner,
        address _routerImpl,
        uint256 _findPrevPriceCount,
        uint256 _maxMatchCount,
        uint256 _cancelLimit,
        address _marketImpl,
        address _pairImpl,
        address _tickSizeSetter
    ) external initializer {
        __Ownable_init(_owner);

        if (_routerImpl == address(0)) revert CrossDexInitializeData("routerImpl");
        if (_marketImpl == address(0)) revert CrossDexInitializeData("marketImpl");
        if (_pairImpl == address(0)) revert CrossDexInitializeData("pairImpl");
        {
            // deploy router
            ERC1967Proxy proxy = new ERC1967Proxy(_routerImpl, hex"");
            ROUTER = payable(address(proxy));
            IRouterV3(ROUTER).initialize(_findPrevPriceCount, _maxMatchCount, _cancelLimit);
        }
        {
            // deploy market & pair logic contracts
            marketImpl = _marketImpl;
            pairImpl = _pairImpl;
            if (_tickSizeSetter != address(0)) tickSizeSetter = _tickSizeSetter;
        }
    }

    function allMarkets() external view returns (address[] memory markets, address[] memory quotes) {
        uint256 length = _allMarkets.length();
        markets = new address[](length);
        quotes = new address[](length);
        for (uint256 i = 0; i < length;) {
            (markets[i], quotes[i]) = _allMarkets.at(i);
            unchecked {
                ++i;
            }
        }
    }

    function checkTickSizeRoles(address account) external view override {
        // check account is tick size setter
        if (tickSizeSetter == address(0) || tickSizeSetter != account) {
            revert CrossDexUnauthorizedChangeTickSizes(account);
        }
    }

    function checkFeeControllerAllowed(address feeController) external view override {
        if (!_allowedFeeControllers.contains(feeController)) revert CrossDexInvalidFeeController(feeController);
    }

    function isMarket(address market) public view returns (bool) {
        return _allMarkets.contains(market);
    }

    function createMarket(address _owner, address quote, address feeController, string memory message)
        external
        onlyOwner
        returns (address)
    {
        bytes memory bytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(
                marketImpl, abi.encodeCall(IMarketV3.initialize, (_owner, ROUTER, quote, pairImpl, feeController))
            )
        );
        bytes32 salt = keccak256(abi.encode(quote, message));
        address market = Create2.deploy(0, salt, bytecode);
        _allMarkets.set(market, quote);

        emit MarketCreated(quote, market, _owner, feeController, message);
        return market;
    }

    function setTickSizeSetter(address setter) external onlyOwner {
        if (setter == address(0) || tickSizeSetter == setter) {
            revert CrossDexInvalidTickSizeSetter(tickSizeSetter, setter);
        }

        emit TickSizeSetterSet(tickSizeSetter, setter);
        tickSizeSetter = setter;
    }

    function pairCreated(address pair) external override onlyMarket {
        pairToMarket[pair] = _msgSender();
    }

    function setPairImpl(address _pairImpl) external onlyOwner {
        if (_pairImpl == address(0)) revert CrossDexInitializeData("pairImpl");
        emit PairImplSet(pairImpl, _pairImpl);
        pairImpl = _pairImpl;
    }

    function setMarketImpl(address _marketImpl) external onlyOwner {
        if (_marketImpl == address(0)) revert CrossDexInitializeData("marketImpl");
        emit MarketImplSet(marketImpl, _marketImpl);
        marketImpl = _marketImpl;
    }

    function setFeeControllerAllow(address feeController, bool allowed) external onlyOwner {
        bool ok;
        if (allowed) ok = _allowedFeeControllers.add(feeController);
        else ok = _allowedFeeControllers.remove(feeController);
        if (ok) emit FeeControllerAllowed(feeController, allowed);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
