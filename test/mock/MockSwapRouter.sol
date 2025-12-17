// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {ISwapRouter} from "../../src/interfaces/ISwapRouter.sol";
import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";

/**
 * @title MockSwapRouter
 * @notice Mock implementation of Uniswap V3 SwapRouter for testing
 */
contract MockSwapRouter is ISwapRouter {
    // Exchange rate: 1 tokenIn = exchangeRate tokenOut (in 18 decimals)
    mapping(address => mapping(address => uint256)) public exchangeRates;

    /**
     * @notice Set exchange rate for token pair
     * @param tokenIn Input token address
     * @param tokenOut Output token address
     * @param rate Exchange rate (1e18 = 1:1)
     */
    function setExchangeRate(address tokenIn, address tokenOut, uint256 rate) external {
        exchangeRates[tokenIn][tokenOut] = rate;
    }

    /**
     * @notice Mock implementation of exactInputSingle
     * @param params Swap parameters
     * @return amountOut Amount of output tokens
     */
    function exactInputSingle(ExactInputSingleParams calldata params) external payable override returns (uint256) {
        require(params.tokenIn != address(0) && params.tokenOut != address(0), "Invalid tokens");
        require(params.amountIn > 0, "Invalid amountIn");

        // Get exchange rate (default to 1:1 if not set)
        uint256 rate = exchangeRates[params.tokenIn][params.tokenOut];
        if (rate == 0) rate = 1e18;

        // Calculate output amount
        uint256 amountOut = (params.amountIn * rate) / 1e18;

        require(amountOut >= params.amountOutMinimum, "Insufficient output amount");

        // Transfer tokens
        IERC20(params.tokenIn).transferFrom(msg.sender, address(this), params.amountIn);
        IERC20(params.tokenOut).transfer(params.recipient, amountOut);

        return amountOut;
    }

    /**
     * @notice Fund the router with tokens for swapping
     * @param token Token address
     * @param amount Amount to fund
     */
    function fundRouter(address token, uint256 amount) external {
        IERC20(token).transferFrom(msg.sender, address(this), amount);
    }
}

