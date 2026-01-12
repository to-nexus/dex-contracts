// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {Script, console} from "forge-std/Script.sol";

import {CrossDexImplV2} from "../src/CrossDexImplV2.sol";
import {CrossDexRouterV2} from "../src/CrossDexRouterV2.sol";
import {MarketImplV2} from "../src/MarketImplV2.sol";
import {PairImplV2} from "../src/PairImplV2.sol";

contract CloseMarket is Script {
    // 1. transfer ownership to the market owner
    // gorge write --rpc-url $CROSS --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa --out transfer_ownership.json \
    // 0x4469a9879Ec02CE4313bef5e2F21b4F17f2B70e4 "execute(address,uint256,bytes)" 0xcb95777d0f8d2EfA5e836Cb65f814dF8C7261d83 0 \
    // $(cast calldata "transferOwnership(address)" 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa)

    // gorge sign --out transfer_ownership_signed.json transfer_ownership.json
    // gorge send --rpc-url $CROSS --out transfer_ownership_sent.json transfer_ownership_signed.json

    // 2. all pause
    // gorge script --rpc-url $CROSS --out all_pause.json \
    // --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa \
    // --sig "allPause(address)" \
    // --out ./script.json \
    // ./script/CloseMarket.s.sol:CloseMarket \
    // 0xcb95777d0f8d2EfA5e836Cb65f814dF8C7261d83

    // gorge sign --out all_pause_signed.json all_pause.json
    // gorge send --rpc-url $CROSS --out all_pause_sent.json all_pause_signed.json

    function allPause(address market) external {
        vm.startBroadcast();
        MarketImplV2 MARKET = MarketImplV2(payable(market));
        (, address[] memory pairs) = MARKET.allPairs();
        uint256 length = pairs.length;
        for (uint256 i = 0; i < length; i++) {
            PairImplV2 PAIR = PairImplV2(payable(pairs[i]));
            bool paused = PAIR.paused();
            if (paused) continue;
            PAIR.setPause(true);
        }
        vm.stopBroadcast();
    }

    function checkClosedMarket(address market) external {
        MarketImplV2 MARKET = MarketImplV2(payable(market));
        (, address[] memory pairs) = MARKET.allPairs();
        uint256 length = pairs.length;
        for (uint256 i = 0; i < length; i++) {
            uint256 baseReserve = PairImplV2(pairs[i]).baseReserve();
            uint256 quoteReserve = PairImplV2(pairs[i]).quoteReserve();
            if (baseReserve > 0 || quoteReserve > 0) console.log("pair", i, baseReserve, quoteReserve);
        }
    }

    // 3. emergency cancel all orders
    // SDODGE
    // gorge write --rpc-url $CROSS --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa --append --out .emergency_cancel_orders.json \
    // 0x07B1578B66F67AA02587D5B35AFC47D66FB39BBD "emergencyCancelOrder(uint256[])" "[1]"

    // NW
    // gorge write --rpc-url $CROSS --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa --append --out .emergency_cancel_orders.json \
    // 0x08A8132639D8E493DA3752C65F8DC695D6D03F64 "emergencyCancelOrder(uint256[])" "[7,11,12,15,16,18]"

    // PP
    // gorge write --rpc-url $CROSS --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa --append --out .emergency_cancel_orders.json \
    // 0x312F5C69ACE9B8FB1869D43AC7A5A6F065A5353D "emergencyCancelOrder(uint256[])" "[3,4,5]"

    // LADDER
    // gorge write --rpc-url $CROSS --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa --append --out .emergency_cancel_orders.json \
    // 0x831429C1197965FBF63A28CAAE8E91ECC00F55A8 "emergencyCancelOrder(uint256[])" "[1]"

    // ZEDDYZZANG
    // gorge write --rpc-url $CROSS --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa --append --out .emergency_cancel_orders.json \
    // 0x83CCF5064322EC3D0D7B2719C12125E8F8A09E96 "emergencyCancelOrder(uint256[])" "[2,3,8,10,11,12,13,15]"

    // PHRST
    // gorge write --rpc-url $CROSS --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa --append --out .emergency_cancel_orders.json \
    // 0x977AA9DE88D597168500C7F47F8EC0FB6815E977 "emergencyCancelOrder(uint256[])" "[2]"

    // TRUMP
    // gorge write --rpc-url $CROSS --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa --append --out .emergency_cancel_orders.json \
    // 0x9A2E6643D7326391E474EAAD62FEA530245343A8 "emergencyCancelOrder(uint256[])" "[1,3]"

    // KKUL789
    // gorge write --rpc-url $CROSS --sender 0xafcc9E7d739b03CC53e2152d368f365430CF3CCa --append --out .emergency_cancel_orders.json \
    // 0x9A52CC333CA7647FA8FDE626DCCBFFBCC2630B66 "emergencyCancelOrder(uint256[])" "[2,7,8]"

    // gorge sign --out emergency_cancel_orders_signed.json emergency_cancel_orders.json
    // gorge send --rpc-url $CROSS --out emergency_cancel_orders_sent.json emergency_cancel_orders_signed.json
}
