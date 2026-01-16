// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {IPairV3} from "./IPairV3.sol";

uint256 constant BPS_DENOMINATOR = 10000;

interface IFeeController {
    function initialize(bytes memory initData) external;

    // Buy 일 경우 Quote를 수량 + 수수료를 포함해야 하기때문에 미리 계산하는 로직 필요
    function calcBuyVolumeWithFee(bool isMaker, IPairV3.Order memory order) external view returns (uint256 buyVolume);

    // Buy 일 경우 Quote를 수량 + 수수료를 포함해야 하기때문에 미리 계산하는 로직 필요
    function calcBuyVolumeWithFee(bool isMaker, uint256 volume) external view returns (uint256 buyVolume);

    // maker, taker 기준으로 수수료 지불 로직
    // maker 의 수수료는 즉시 지불되고, taker 의 수수료는 정산 로직에서 지불됨 (maker 는 n 이지만 taker 는 1 이기 때문)
    function recodeMatch(IPairV3.Order memory makerOrder, IPairV3.Order memory sellOrder)
        external
        returns (uint256 makerFee);

    // 수수료 정산 로직
    function settleFees() external returns (uint256 takerFee);
}
