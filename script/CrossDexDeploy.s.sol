// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";

import {Script, console} from "forge-std/Script.sol";

import {CrossDexImpl} from "../src/CrossDexImpl.sol";
import {MarketImpl} from "../src/MarketImpl.sol";
import {PairImpl} from "../src/PairImpl.sol";
import {RouterImpl} from "../src/RouterImpl.sol";
import {WCROSS} from "../src/WCROSS.sol";

contract CrossDexDeployScript is Script {
    /*
    forge script -vvv \
    --rpc-url http://127.0.0.1:8545 \
    --force \
    --broadcast \
    --private-key $PRIVATE_KEY \
    --sig "deploy(address,uint256)" \
    ./script/CrossDexDeploy.s.sol:CrossDexDeployScript $OWNER $MAX_MATCH_COUNT
    */
    function deploy(address owner, uint256 maxMatchCount) external {
        vm.startBroadcast();

        // deploy impl contracts
        address routerImpl = address(new RouterImpl());
        address marketImpl = address(new MarketImpl());
        address pairImpl = address(new PairImpl());

        // deploy cross dex
        address crossDexImpl = address(new CrossDexImpl());
        ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
        CrossDexImpl CROSS_DEX = CrossDexImpl(address(proxy));
        CROSS_DEX.initialize(owner, routerImpl, maxMatchCount, marketImpl, pairImpl);

        // get contracts from CROSS_DEX
        RouterImpl ROUTER = RouterImpl(CROSS_DEX.ROUTER());
        WCROSS WCross = WCROSS(ROUTER.WCross());
        vm.stopBroadcast();

        console.log("CROSS_DEX: ", address(CROSS_DEX));
        console.log("ROUTER: ", address(ROUTER));
        console.log("WCROSS: ", address(WCross));
    }
}
