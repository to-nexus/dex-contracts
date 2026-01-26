// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

/**
 * @title IUniswapV3Pool
 * @notice Minimal interface for Uniswap V3 Pool
 */
interface IUniswapV3Pool {
    /**
     * @notice The 0th storage slot in the pool stores many values
     * @return sqrtPriceX96 The current price of the pool as a sqrt(token1/token0) Q64.96 value
     * @return tick The current tick of the pool
     * @return observationIndex The index of the last oracle observation
     * @return observationCardinality The current maximum number of observations stored
     * @return observationCardinalityNext The next maximum number of observations
     * @return feeProtocol The protocol fee for both tokens
     * @return unlocked Whether the pool is currently locked to reentrancy
     */
    function slot0()
        external
        view
        returns (
            uint160 sqrtPriceX96,
            int24 tick,
            uint16 observationIndex,
            uint16 observationCardinality,
            uint16 observationCardinalityNext,
            uint8 feeProtocol,
            bool unlocked
        );

    /// @notice The first of the two tokens of the pool, sorted by address
    function token0() external view returns (address);

    /// @notice The second of the two tokens of the pool, sorted by address
    function token1() external view returns (address);
}
