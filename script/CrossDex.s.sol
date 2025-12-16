// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {Script, console} from "forge-std/Script.sol";

import {CrossDexImpl} from "../src/CrossDexImpl.sol";
import {CrossDexRouter} from "../src/CrossDexRouter.sol";
import {MarketImpl} from "../src/MarketImpl.sol";
import {PairImpl} from "../src/PairImpl.sol";

contract CrossDexScript is Script {
    function setUp() external {}

    function deployImpls() external {
        vm.startBroadcast();
        new CrossDexImpl();
        new CrossDexRouter();
        new MarketImpl();
        new PairImpl();
        vm.stopBroadcast();
    }

    function deployCrossDex(
        address crossDexImpl,
        address owner,
        address crossDexRouter,
        uint256 findPrevPriceCount,
        uint256 maxMatchCount,
        uint256 cancelLimit,
        address marketImpl,
        address pairImpl
    ) external {
        vm.broadcast();
        new ERC1967Proxy(
            crossDexImpl,
            abi.encodeCall(
                CrossDexImpl.initialize,
                (owner, crossDexRouter, findPrevPriceCount, maxMatchCount, cancelLimit, marketImpl, pairImpl)
            )
        );
    }

    function createMarket(address _crossDex, address owner, address quote, address feeCollector, uint256 feeBps)
        external
    {
        vm.startBroadcast();
        CrossDexImpl crossDex = CrossDexImpl(_crossDex);
        crossDex.createMarket(owner, quote, feeCollector, feeBps);
        vm.stopBroadcast();
    }

    function createPair(address _market, address base, uint256 tickSize, uint256 lotSize) external {
        vm.startBroadcast();
        MarketImpl market = MarketImpl(_market);
        market.createPair(base, tickSize, lotSize);
        vm.stopBroadcast();
    }

    function printAllDexContracts(address _crossDex) external {
        vm.startBroadcast();
        CrossDexImpl crossDex = CrossDexImpl(_crossDex);
        {
            address router = crossDex.ROUTER();
            address wcross = address(CrossDexRouter(payable(router)).CROSS());
            console.log("Router", router);
            console.log("wCROSS", wcross);
        }
        {
            (address[] memory quotes, address[] memory markets) = crossDex.allMarkets();
            uint256 length = markets.length;

            for (uint256 i = 0; i < length; ++i) {
                console.log("==========================================");
                console.log("Market", markets[i], "Quote", quotes[i]);
                MarketImpl market = MarketImpl(markets[i]);
                (address[] memory bases, address[] memory pairs) = market.allPairs();
                uint256 pairLength = pairs.length;
                for (uint256 j = 0; j < pairLength; ++j) {
                    console.log("    Pair", pairs[j], "Base", bases[j]);
                }
            }
        }
        vm.stopBroadcast();
    }
}
