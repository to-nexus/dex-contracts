// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {Script, console} from "forge-std/Script.sol";

import {CrossDexImplV2} from "../src/CrossDexImplV2.sol";
import {CrossDexRouterV2} from "../src/CrossDexRouterV2.sol";
import {MarketImplV2} from "../src/MarketImplV2.sol";
import {PairImplV2} from "../src/PairImplV2.sol";
import {Verse8MarketOwner} from "../src/Verse8MarketOwner.sol";

contract DeployCrossDexV2 is Script {
    function setUp() external {}

    function deployV2Impls() external {
        vm.startBroadcast();
        address crossDexImpl = address(new CrossDexImplV2());
        address crossDexRouter = address(new CrossDexRouterV2());
        address marketImpl = address(new MarketImplV2());
        address pairImpl = address(new PairImplV2());
        vm.stopBroadcast();

        console.log("CrossDexImplV2:", crossDexImpl);
        console.log("CrossDexRouterV2:", crossDexRouter);
        console.log("MarketImplV2:", marketImpl);
        console.log("PairImplV2:", pairImpl);
    }

    function deployVerse8MarketOwner(address owner, address[] memory creators) external {
        vm.startBroadcast();
        address verse8MarketOwner = address(new Verse8MarketOwner(owner, creators));
        vm.stopBroadcast();
        console.log("Verse8MarketOwner:", verse8MarketOwner);
    }

    function deployCrossDexV2(
        address crossDexImpl,
        address owner,
        address crossDexRouter,
        uint256 findPrevPriceCount,
        uint256 maxMatchCount,
        uint256 cancelLimit,
        address marketImpl,
        address pairImpl,
        address tickSizeSetter
    ) external {
        vm.broadcast();
        address crossDex = address(
            new ERC1967Proxy(
                crossDexImpl,
                abi.encodeCall(
                    CrossDexImplV2.initialize,
                    (
                        owner,
                        crossDexRouter,
                        findPrevPriceCount,
                        maxMatchCount,
                        cancelLimit,
                        marketImpl,
                        pairImpl,
                        tickSizeSetter
                    )
                )
            )
        );
        console.log("CrossDexV2:", crossDex);
    }
}
