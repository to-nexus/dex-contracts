// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
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

    event MarketCreated(address indexed quote, address indexed market, address indexed owner, address fee_collector);

    address payable public ROUTER; // immutable

    address public marketImpl;
    address public pairImpl;

    EnumerableMap.AddressToAddressMap private _allMarkets; // quote => market
    mapping(address pair => address) public override pairToMarket;

    uint256[45] __gap;

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
        uint256 _maxMatchCount,
        address _marketImpl,
        address _pairImpl
    ) external initializer {
        if (_routerImpl == address(0)) revert CrossDexInitializeData("routerImpl");
        if (_marketImpl == address(0)) revert CrossDexInitializeData("marketImpl");
        if (_pairImpl == address(0)) revert CrossDexInitializeData("pairImpl");
        {
            // deploy router
            ERC1967Proxy proxy = new ERC1967Proxy(_routerImpl, hex"");
            ROUTER = payable(address(proxy));
            IRouterInitializer(ROUTER).initialize(_maxMatchCount);
        }
        {
            // deploy market & pair logic contracts
            marketImpl = _marketImpl;
            pairImpl = _pairImpl;
        }
        __Ownable_init(_owner);
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

    function isMarket(address market) public view returns (bool) {
        (bool success, bytes memory data) = market.staticcall(abi.encodeCall(IMarketInitializer.QUOTE, ()));
        if (!success) return false;
        address quote = abi.decode(data, (address));
        return _allMarkets.get(quote) == market;
    }

    function createMarket(address _owner, address _fee_collector, address _quote)
        external
        onlyOwner
        returns (address)
    {
        IMarketInitializer market = IMarketInitializer(address(new ERC1967Proxy(marketImpl, hex"")));
        market.initialize(_owner, ROUTER, _fee_collector, _quote, pairImpl);

        address _market = address(market);
        if (!_allMarkets.set(_quote, _market)) revert CrossDexAlreadyCreatedMarketQuote(_quote);

        emit MarketCreated(_quote, _market, _owner, _fee_collector);
        return _market;
    }

    function pairCreated(address pair) external override onlyMarket {
        pairToMarket[pair] = _msgSender();
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
