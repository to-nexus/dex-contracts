// // SPDX-License-Identifier: BUSL-1.1
// pragma solidity ^0.8.13;

// import {Test, console} from "forge-std/Test.sol";

// import {Verse8MarketOwner} from "../../src/Verse8MarketOwner.sol";
// import {MarketImplV2} from "../../src/v2/MarketImplV2.sol";
// import {PairImplV2} from "../../src/v2/PairImplV2.sol";
// import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";

// contract RemoteTest is Test {
//     PairImplV2 public crossRubyx = PairImplV2(0x77F322F2e7913193B0eF6beaa356aC06B8448DA6);
//     PairImplV2 public usdtRubyx = PairImplV2(0x44AC60d1b6AfBb79beC09F4ebB7dE020F5267aBA);

//     function setUp() external {}

//     function test_emergencyCancelOrder() external {
//         address target = 0x6B5fE3cc89B26C7D5eFb722079Db29D434654239;

//         vm.startPrank(0xafcc9E7d739b03CC53e2152d368f365430CF3CCa);
//         crossRubyx.setPause(true);
//         uint256[] memory orderIds = new uint256[](1);
//         orderIds[0] = 52769;
//         crossRubyx.emergencyCancelOrder(orderIds);
//         crossRubyx.setPause(false);

//         usdtRubyx.setPause(true);
//         orderIds = new uint256[](3);
//         orderIds[0] = 13568;
//         orderIds[1] = 13589;
//         orderIds[2] = 13613;
//         usdtRubyx.emergencyCancelOrder(orderIds);
//         usdtRubyx.setPause(false);

//         {
//             (uint256 a, uint256 b) = crossRubyx.accountReserves(target);
//             console.log("crossRubyx", a, b);
//         }
//         {
//             (uint256 a, uint256 b) = usdtRubyx.accountReserves(target);
//             console.log("usdtRubyx", a, b);
//         }

//         vm.stopPrank();
//     }

//     function test_emergencyCancelOrder_2() external {
//         Verse8MarketOwner verse8Owner = Verse8MarketOwner(0x4469a9879Ec02CE4313bef5e2F21b4F17f2B70e4);
//         MarketImplV2 market = MarketImplV2(0xcb95777d0f8d2EfA5e836Cb65f814dF8C7261d83);
//         /*
//                 vm.startPrank(0xafcc9E7d739b03CC53e2152d368f365430CF3CCa);
//                 verse8Owner.execute(
//                     address(market), 0, abi.encodeCall(market.transferOwnership, (0xafcc9E7d739b03CC53e2152d368f365430CF3CCa))
//                 );

//                 (, address[] memory pairs) = market.allPairs();
//                 for (uint256 i = 0; i < pairs.length; i++) {
//                     PairImplV2(pairs[i]).setPause(true);
//                 }
//                 vm.roll(block.number + 1);

//                 { // SDODGE
//                     uint256[] memory orderIds = new uint256[](1);
//                     orderIds[0] = 52769;
//                     PairImplV2(0x07B1578b66f67aA02587D5b35AFc47d66FB39bBd).emergencyCancelOrder{gas: 4e6}(orderIds);
//                 }
//                 vm.roll(block.number + 1);
//                 { // NW
//                     uint256[] memory orderIds = new uint256[](7);
//                     orderIds[0] = 7;
//                     orderIds[1] = 11;
//                     orderIds[2] = 12;
//                     orderIds[3] = 15;
//                     orderIds[4] = 16;
//                     orderIds[5] = 18;
//                     PairImplV2(0x08A8132639d8E493dA3752C65F8Dc695D6D03f64).emergencyCancelOrder(orderIds);
//                 }
//                 vm.roll(block.number + 1);
//                 { // PP
//                     uint256[] memory orderIds = new uint256[](3);
//                     orderIds[0] = 3;
//                     orderIds[1] = 4;
//                     orderIds[2] = 5;
//                     PairImplV2(0x312F5C69aCE9B8fb1869D43Ac7a5a6F065A5353d).emergencyCancelOrder(orderIds);
//                 }
//                 vm.roll(block.number + 1);
//                 { // LADDER
//                     uint256[] memory orderIds = new uint256[](1);
//                     orderIds[0] = 1;
//                     PairImplV2(0x831429C1197965fBF63a28CAaE8E91ECc00F55a8).emergencyCancelOrder(orderIds);
//                 }
//                 vm.roll(block.number + 1);
//                 { // ZEDDYZZANG
//                     uint256[] memory orderIds = new uint256[](8);
//                     orderIds[0] = 2;
//                     orderIds[1] = 3;
//                     orderIds[2] = 8;
//                     orderIds[3] = 10;
//                     orderIds[4] = 11;
//                     orderIds[5] = 12;
//                     orderIds[6] = 13;
//                     orderIds[7] = 15;
//                     PairImplV2(0x83CcF5064322EC3d0d7b2719c12125E8F8A09e96).emergencyCancelOrder(orderIds);
//                 }
//                 vm.roll(block.number + 1);
//                 { // PHRST
//                     uint256[] memory orderIds = new uint256[](2);
//                     orderIds[0] = 2;
//                     orderIds[1] = 2;
//                     PairImplV2(0x977AA9dE88D597168500c7F47F8ec0Fb6815E977).emergencyCancelOrder(orderIds);
//                 }
//                 vm.roll(block.number + 1);
//                 { // TRUMP
//                     uint256[] memory orderIds = new uint256[](2);
//                     orderIds[0] = 1;
//                     orderIds[1] = 3;
//                     PairImplV2(0x9A2E6643D7326391e474eaAd62feA530245343a8).emergencyCancelOrder(orderIds);
//                 }
//                 vm.roll(block.number + 1);
//                 { // KKUL789
//                     uint256[] memory orderIds = new uint256[](3);
//                     orderIds[0] = 2;
//                     orderIds[1] = 7;
//                     orderIds[2] = 8;
//                     PairImplV2(0x9a52cC333Ca7647FA8fDE626DCcBffBCc2630b66).emergencyCancelOrder(orderIds);
//                 }
//                 vm.roll(block.number + 1);
//         */
//         (, address[] memory pairs) = market.allPairs();
//         for (uint256 i = 0; i < pairs.length; i++) {
//             uint256 baseReserve = PairImplV2(pairs[i]).baseReserve();
//             uint256 quoteReserve = PairImplV2(pairs[i]).quoteReserve();
//             if (baseReserve > 0 || quoteReserve > 0) console.log("pair", i, baseReserve, quoteReserve);
//         }

//         // vm.stopPrank();
//     }
// }
