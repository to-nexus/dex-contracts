// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {IUniswapV3Pool} from "../../src/interfaces/IUniswapV3Pool.sol";

/**
 * @title MockUniswapV3Pool
 * @notice Mock implementation of Uniswap V3 Pool for testing
 */
contract MockUniswapV3Pool is IUniswapV3Pool {
    address public override token0;
    address public override token1;

    uint160 public sqrtPriceX96;
    int24 public tick;

    constructor(address _token0, address _token1) {
        // Ensure tokens are sorted (token0 < token1)
        if (_token0 < _token1) {
            token0 = _token0;
            token1 = _token1;
        } else {
            token0 = _token1;
            token1 = _token0;
        }
        // Default to tick 0 (1:1 price)
        tick = 0;
        sqrtPriceX96 = 79228162514264337593543950336; // sqrt(1) * 2^96
    }

    function slot0()
        external
        view
        override
        returns (
            uint160 _sqrtPriceX96,
            int24 _tick,
            uint16 observationIndex,
            uint16 observationCardinality,
            uint16 observationCardinalityNext,
            uint8 feeProtocol,
            bool unlocked
        )
    {
        return (sqrtPriceX96, tick, 0, 1, 1, 0, true);
    }

    /**
     * @notice Set the current tick for testing
     * @param _tick New tick value
     */
    function setTick(int24 _tick) external {
        tick = _tick;
    }

    /**
     * @notice Set the current sqrtPriceX96 for testing
     * @param _sqrtPriceX96 New sqrtPriceX96 value
     */
    function setSqrtPriceX96(uint160 _sqrtPriceX96) external {
        sqrtPriceX96 = _sqrtPriceX96;
    }
}

