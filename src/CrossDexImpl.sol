// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableSet.sol";

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/access/OwnableUpgradeable.sol";

import {MarketImpl} from "./MarketImpl.sol";
import {PairImpl} from "./PairImpl.sol";
import {RouterImpl} from "./RouterImpl.sol";
import {WCROSS} from "./WCROSS.sol";

import {ICrossDex} from "./interfaces/ICrossDex.sol";
import {IRouter} from "./interfaces/IRouter.sol";

contract CrossDexImpl is ICrossDex, UUPSUpgradeable, OwnableUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error CrossDexAlreadyCreatedMarketQuote(address);
    error CrossDexInvalidMarketAddress(address);

    event CreatedMarket(address indexed quote, address indexed market, address indexed owner, address fee_collector);

    address payable public ROUTER;

    address public marketImpl;
    address public pairImpl;

    EnumerableSet.AddressSet private _allQuotes;
    EnumerableSet.AddressSet private _allMarkets;
    mapping(address quote => address) public quoteToMarket;

    modifier onlyMarket() {
        if (!_allMarkets.contains(_msgSender())) revert CrossDexInvalidMarketAddress(_msgSender());
        _;
    }

    function initialize(address _owner, uint256 _maxMatchCount) external initializer {
        {
            // deploy router
            RouterImpl routerImpl = new RouterImpl();
            ERC1967Proxy proxy = new ERC1967Proxy(address(routerImpl), hex"");
            RouterImpl router = RouterImpl(payable(address(proxy)));
            router.initialize(_owner, _maxMatchCount);
        }
        {
            // deploy market & pair logic contracts
            marketImpl = address(new MarketImpl());
            pairImpl = address(new PairImpl());
        }
        __Ownable_init(_owner);
    }

    function allMarkets() external view returns (address[] memory, address[] memory) {
        return (_allQuotes.values(), _allMarkets.values());
    }

    function createMarket(address _owner, address _fee_collector, address _quote) external onlyOwner {
        if (!_allQuotes.add(_quote)) revert CrossDexAlreadyCreatedMarketQuote(_quote);

        MarketImpl _market = MarketImpl(address(new ERC1967Proxy(marketImpl, hex"")));
        _market.initialize(_owner, ROUTER, _fee_collector, _quote, pairImpl);

        address market = address(_market);
        _allMarkets.add(market);
        quoteToMarket[_quote] = market;

        emit CreatedMarket(_quote, market, _owner, _fee_collector);
    }

    function notifyPairCreated(address pair) external onlyMarket {
        IRouter(ROUTER).addPair(pair);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
