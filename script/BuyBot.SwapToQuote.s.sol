// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {BuyBot} from "../src/BuyBot.sol";
import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title SwapToQuoteScript
 * @notice Script to execute swapToQuote on deployed BuyBot
 */
contract SwapToQuoteScript is Script {
    // Cross Testnet addresses
    address constant BUYBOT = 0xdb66f3c9708AA74133Fb78dcDE252f585D7713a8; // Deployed BuyBot address
    address constant PAIR = 0x8C46a3f37013cF587167d4c75d29b170cB5DAAbe; // CROSSD/CROSS pair
    address constant USDT = 0x9F85c7B5D7637E18f946cc8AF9C131318c6833d9;
    address constant CROSSD = 0x9364ea6790f6E0EcFaa5164085f2a7de34EC55Fb;
    uint24 constant FEE_TIER = 100; // 0.01%
    uint256 constant MIN_AMOUNT_OUT = 0; // Minimum amount out (0 for no slippage protection)

    function run() external {
        // Check BuyBot address is set
        require(BUYBOT != address(0), "Set BUYBOT address first");

        BuyBot buyBot = BuyBot(payable(BUYBOT));

        // Log current state
        console.log("========== SwapToQuote Execution ==========");
        console.log("BuyBot:", BUYBOT);
        console.log("Pair:", PAIR);
        console.log("SwapToken:", buyBot.swapToken());

        // Check USDT balance
        uint256 usdtBalance = IERC20(USDT).balanceOf(BUYBOT);
        console.log("USDT Balance:", usdtBalance);

        if (usdtBalance == 0) {
            console.log("ERROR: No USDT balance in BuyBot");
            return;
        }

        // Execute swap
        vm.broadcast();
        uint256 amountOut = buyBot.swapToQuote(PAIR, FEE_TIER, MIN_AMOUNT_OUT);

        console.log("========== Swap Result ==========");
        console.log("Amount Out (CROSSD):", amountOut);
        console.log("New CROSSD Balance:", IERC20(CROSSD).balanceOf(BUYBOT));
    }

    /**
     * @notice Execute swapToQuote with custom parameters
     * @param buyBotAddress BuyBot contract address
     * @param pairAddress Trading pair address
     * @param feeTier Uniswap V3 fee tier
     * @param minAmountOut Minimum amount of quote tokens to receive (for slippage protection)
     */
    function swapToQuote(address buyBotAddress, address pairAddress, uint24 feeTier, uint256 minAmountOut) external {
        BuyBot buyBot = BuyBot(payable(buyBotAddress));

        console.log("========== SwapToQuote Execution ==========");
        console.log("BuyBot:", buyBotAddress);
        console.log("Pair:", pairAddress);
        console.log("Fee Tier:", feeTier);
        console.log("Min Amount Out:", minAmountOut);
        console.log("SwapToken:", buyBot.swapToken());

        vm.broadcast();
        uint256 amountOut = buyBot.swapToQuote(pairAddress, feeTier, minAmountOut);

        console.log("========== Swap Result ==========");
        console.log("Amount Out:", amountOut);
    }
}
