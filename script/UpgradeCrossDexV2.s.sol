// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {Script, console} from "forge-std/Script.sol";

import {CrossDexImplV2} from "../src/CrossDexImplV2.sol";
import {CrossDexRouterV2} from "../src/CrossDexRouterV2.sol";
import {MarketImplV2} from "../src/MarketImplV2.sol";
import {PairImplV2} from "../src/PairImplV2.sol";

contract UpgradeCrossDexV2 is Script {
    uint32 constant NO_FEE_BPS = type(uint32).max;

    function setUp() external {}

    function upgradeCrossDexV2(address dex, address crossDexImpl, address marketImpl, address pairImpl) external {
        CrossDexImplV2 DEX = CrossDexImplV2(dex);

        vm.startBroadcast();
        DEX.upgradeToAndCall(crossDexImpl, abi.encodeCall(CrossDexImplV2.reinitialize, (marketImpl, pairImpl)));
        vm.stopBroadcast();
    }

    function upgradeCrossRouterV2(address router, address crossRouterImpl) external {
        vm.startBroadcast();
        CrossDexRouterV2(router).upgradeToAndCall(crossRouterImpl, "");
        vm.stopBroadcast();
    }

    function upgradeMarkets(
        address[] memory markets,
        address marketImpl,
        address pairImpl,
        uint32 sellerMakerFeeBps,
        uint32 sellerTakerFeeBps,
        uint32 buyerMakerFeeBps,
        uint32 buyerTakerFeeBps
    ) external {
        vm.startBroadcast();
        for (uint256 i = 0; i < markets.length; i++) {
            MarketImplV2 market = MarketImplV2(markets[i]);
            market.upgradeToAndCall(marketImpl, "");
            market.setMarketFees(sellerMakerFeeBps, sellerTakerFeeBps, buyerMakerFeeBps, buyerTakerFeeBps);
            market.setPairImpl(pairImpl);
        }
        vm.stopBroadcast();
    }

    function upgradePairs(address[] memory pairs, address pairImpl) external {
        vm.startBroadcast();
        for (uint256 i = 0; i < pairs.length; i++) {
            PairImplV2 pair = PairImplV2(pairs[i]);
            pair.upgradeToAndCall(pairImpl, "");
            pair.setPairFees(NO_FEE_BPS, NO_FEE_BPS, NO_FEE_BPS, NO_FEE_BPS);
        }
        vm.stopBroadcast();
    }
}
