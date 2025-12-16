// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {Script, console} from "forge-std/Script.sol";

import {CrossDexImplV2} from "../src/CrossDexImplV2.sol";
import {CrossDexRouterV2} from "../src/CrossDexRouterV2.sol";
import {IMarketV2, MarketImplV2} from "../src/MarketImplV2.sol";
import {PairImplV2} from "../src/PairImplV2.sol";

contract CreateCrossDMarket is Script {
    function copyMarket(address dex, address owner, address quote, address srcMarket, string memory message) external {
        vm.startBroadcast();
        CrossDexImplV2 DEX = CrossDexImplV2(dex);
        MarketImplV2 MARKET = MarketImplV2(srcMarket);

        address feeCollector = MARKET.feeCollector();
        IMarketV2.FeeConfig memory feeConfig = MARKET.getFeeConfig();
        bytes memory feeData = abi.encode(
            feeConfig.sellerMakerFeeBps,
            feeConfig.sellerTakerFeeBps,
            feeConfig.buyerMakerFeeBps,
            feeConfig.buyerTakerFeeBps
        );
        address market = DEX.createMarket(owner, quote, feeCollector, feeData, message);
        vm.stopBroadcast();

        console.log(string.concat(message, " Market created: "), market);
    }

    function copyMarketPairs(address srcMarket, address dstMarket) external {
        vm.startBroadcast();
        MarketImplV2 SRC_MARKET = MarketImplV2(srcMarket);
        MarketImplV2 DST_MARKET = MarketImplV2(dstMarket);

        (address[] memory bases, address[] memory pairs) = SRC_MARKET.allPairs();
        for (uint256 i = 0; i < bases.length; i++) {
            PairImplV2 PAIR = PairImplV2(pairs[i]);

            (uint256 tickSize, uint256 lotSize) = PAIR.tickSizes();
            (uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps) =
                PAIR.feeConfig();
            bytes memory feeData = abi.encode(sellerMakerFeeBps, sellerTakerFeeBps, buyerMakerFeeBps, buyerTakerFeeBps);

            DST_MARKET.createPair(bases[i], tickSize, lotSize, feeData);
        }
        vm.stopBroadcast();
    }

    function pauseAllPairs(address market) external {
        vm.startBroadcast();
        MarketImplV2 MARKET = MarketImplV2(market);
        (, address[] memory pairs) = MARKET.allPairs();
        for (uint256 i = 0; i < pairs.length; i++) {
            PairImplV2 PAIR = PairImplV2(pairs[i]);
            PAIR.setPause(true);
        }
        vm.stopBroadcast();
    }
}
