// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Create2} from "@openzeppelin-contracts-5.2.0/utils/Create2.sol";
import {EnumerableMap} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableMap.sol";

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/access/OwnableUpgradeable.sol";

import {ICrossDex} from "./interfaces/ICrossDex.sol";
import {IMarketInitializer} from "./interfaces/IMarket.sol";
import {IRouter, IRouterInitializer} from "./interfaces/IRouter.sol";

import {WETH} from "./WETH.sol";

contract CrossDexImpl is ICrossDex, UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableMap for EnumerableMap.AddressToAddressMap;

    error CrossDexInitializeData(bytes32);
    error CrossDexAlreadyCreatedMarketQuote(address);
    error CrossDexInvalidMarketAddress(address);
    error CrossDexUnauthorizedChangeTickSizes(address);
    error CrossDexInvalidTickSizeSetter(address current, address input);

    event MarketCreated(address indexed quote, address indexed market, address indexed owner, address fee_collector);
    event TickSizeSetterSet(address indexed before, address indexed current);

    address payable public ROUTER; // immutable

    address public marketImpl;
    address public pairImpl;

    EnumerableMap.AddressToAddressMap private _allMarkets; // quote => market
    mapping(address pair => address) public override pairToMarket;

    address public tickSizeSetter;

    uint256[44] __gap;

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
        address _pairImpl
    ) external initializer {
        __Ownable_init(_owner);

        if (_routerImpl == address(0)) revert CrossDexInitializeData("routerImpl");
        if (_marketImpl == address(0)) revert CrossDexInitializeData("marketImpl");
        if (_pairImpl == address(0)) revert CrossDexInitializeData("pairImpl");
        {
            // deploy router
            ERC1967Proxy proxy = new ERC1967Proxy(_routerImpl, hex"");
            ROUTER = payable(address(proxy));
            IRouterInitializer(ROUTER).initialize(_findPrevPriceCount, _maxMatchCount, _cancelLimit);
        }
        {
            // deploy market & pair logic contracts
            marketImpl = _marketImpl;
            pairImpl = _pairImpl;
        }
    }

    function allMarkets() external view returns (address[] memory quotes, address[] memory markets) {
        uint256 length = _allMarkets.length();
        quotes = new address[](length);
        markets = new address[](length);
        for (uint256 i = 0; i < length;) {
            (quotes[i], markets[i]) = _allMarkets.at(i);
            unchecked {
                ++i;
            }
        }
    }

    function quoteToMarket(address quote) external view returns (address) {
        return _allMarkets.get(quote);
    }

    function checkTickSizeRoles(address account) external view override {
        // check account is tick size setter
        if (tickSizeSetter == address(0) || tickSizeSetter != account) {
            revert CrossDexUnauthorizedChangeTickSizes(account);
        }
    }

    function isMarket(address market) public view returns (bool) {
        (bool success, bytes memory data) = market.staticcall(abi.encodeCall(IMarketInitializer.QUOTE, ()));
        if (!success) return false;
        address quote = abi.decode(data, (address));
        return _allMarkets.get(quote) == market;
    }

    function createMarket(address _owner, address quote, address feeCollector, uint256 feeBps)
        external
        onlyOwner
        returns (address)
    {
        bytes memory bytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(
                marketImpl,
                abi.encodeCall(IMarketInitializer.initialize, (_owner, ROUTER, quote, pairImpl, feeCollector, feeBps))
            )
        );
        bytes32 salt = keccak256(abi.encodePacked(quote));
        address market = Create2.deploy(0, salt, bytecode);

        if (!_allMarkets.set(quote, market)) revert CrossDexAlreadyCreatedMarketQuote(quote);

        emit MarketCreated(quote, market, _owner, feeCollector);
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

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
