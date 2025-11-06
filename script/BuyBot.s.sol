// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";

import {BuyBot} from "../src/BuyBot.sol";

contract BuyBotScript is Script {
    function setUp() external {}

    /**
     * @notice Deploy a new BuyBot instance
     * @param initialDelay Delay (in seconds) for admin transfer (0 for no delay)
     * @param owner Owner address (gets DEFAULT_ADMIN_ROLE)
     * @param router CrossDexRouter address
     * @param minOrderAmount Minimum order amount in QUOTE token
     * @param interval Minimum time interval (in seconds) between buy executions (0 to disable)
     * @param recipient Address to receive purchased BASE tokens (address(0) to keep in contract)
     * @param buyer Address authorized to execute buyMarket (gets BUYER_ROLE)
     * @param manager Address authorized to set minOrderAmount and interval (gets MANAGER_ROLE)
     */
    function deployBuyBot(
        uint48 initialDelay,
        address owner,
        address router,
        uint256 minOrderAmount,
        uint256 interval,
        address recipient,
        address buyer,
        address manager
    ) external returns (address) {
        vm.broadcast();
        BuyBot bot = new BuyBot(initialDelay, owner, router, minOrderAmount, interval, recipient, buyer, manager);
        return address(bot);
    }

    /**
     * @notice Execute market buy
     * @param buyBot BuyBot contract address
     * @param pair Trading pair address
     * @param amount Amount to spend (must be > 0)
     * @param maxMatchCount Maximum number of orders to match (0 for router default)
     */
    function buyMarket(address buyBot, address pair, uint256 amount, uint256 maxMatchCount) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).buyMarket(pair, amount, maxMatchCount);
    }

    /**
     * @notice Withdraw tokens from BuyBot
     * @param buyBot BuyBot contract address
     * @param token Token address to withdraw
     * @param amount Amount to withdraw (0 for entire balance)
     */
    function withdraw(address buyBot, address token, uint256 amount) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).withdraw(token, amount);
    }

    /**
     * @notice Withdraw ETH from BuyBot
     * @param buyBot BuyBot contract address
     * @param amount Amount to withdraw (0 for entire balance)
     */
    function withdrawETH(address buyBot, uint256 amount) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).withdrawETH(amount);
    }

    /**
     * @notice Set minimum order amount
     * @param buyBot BuyBot contract address
     * @param minOrderAmount New minimum order amount
     */
    function setMinOrderAmount(address buyBot, uint256 minOrderAmount) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).setMinOrderAmount(minOrderAmount);
    }

    /**
     * @notice Set time interval between buy executions
     * @param buyBot BuyBot contract address
     * @param interval New interval in seconds (0 to disable)
     */
    function setInterval(address buyBot, uint256 interval) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).setInterval(interval);
    }

    /**
     * @notice Set recipient address for purchased BASE tokens
     * @param buyBot BuyBot contract address
     * @param recipient New recipient address (address(0) to keep in contract)
     */
    function setRecipient(address buyBot, address recipient) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).setRecipient(recipient);
    }

    /**
     * @notice Grant a role to an address
     * @param buyBot BuyBot contract address
     * @param role Role to grant (use BUYER_ROLE or MANAGER_ROLE constants)
     * @param account Address to grant the role
     */
    function grantRole(address buyBot, bytes32 role, address account) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).grantRole(role, account);
    }

    /**
     * @notice Revoke a role from an address
     * @param buyBot BuyBot contract address
     * @param role Role to revoke (use BUYER_ROLE or MANAGER_ROLE constants)
     * @param account Address to revoke the role
     */
    function revokeRole(address buyBot, bytes32 role, address account) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).revokeRole(role, account);
    }

    /**
     * @notice Begin admin transfer (if delay is set)
     * @param buyBot BuyBot contract address
     * @param newAdmin New admin address
     */
    function beginDefaultAdminTransfer(address buyBot, address newAdmin) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).beginDefaultAdminTransfer(newAdmin);
    }

    /**
     * @notice Accept admin transfer (if delay is set)
     * @param buyBot BuyBot contract address
     */
    function acceptDefaultAdminTransfer(address buyBot) external {
        vm.broadcast();
        BuyBot(payable(buyBot)).acceptDefaultAdminTransfer();
    }

    /**
     * @notice Print all BuyBot information
     * @param buyBot BuyBot contract address
     */
    function printBuyBotInfo(address buyBot) external view {
        BuyBot bot = BuyBot(payable(buyBot));

        console.log("==========================================");
        console.log("BuyBot Address:", buyBot);
        console.log("------------------------------------------");
        console.log("Owner:", bot.owner());
        console.log("Router:", address(bot.router()));
        console.log("Recipient:", bot.recipient());
        console.log("------------------------------------------");
        console.log("Min Order Amount:", bot.minOrderAmount());
        console.log("Interval:", bot.interval(), "seconds");
        console.log("Last Buy Time:", bot.lastBuyTime());
        console.log("------------------------------------------");
        console.log("Note: Check BUYER_ROLE and MANAGER_ROLE via hasRole()");
        console.log("==========================================");
    }

    /**
     * @notice Check if BuyBot can execute buy for a specific pair
     * @param buyBot BuyBot contract address
     * @param pair Trading pair address
     * @param caller Address to check authorization
     */
    function checkCanBuyMarket(address buyBot, address pair, address caller) external view {
        BuyBot bot = BuyBot(payable(buyBot));
        (bool canBuy, uint256 balance) = bot.canBuyMarket(pair, caller);

        console.log("==========================================");
        console.log("Can Buy Market Check");
        console.log("------------------------------------------");
        console.log("Pair:", pair);
        console.log("Caller:", caller);
        console.log("Can Buy:", canBuy);
        console.log("Balance:", balance);
        console.log("==========================================");
    }

    /**
     * @notice Get token balance in BuyBot
     * @param buyBot BuyBot contract address
     * @param token Token address (use NATIVE_COIN constant for ETH)
     */
    function getBalance(address buyBot, address token) external view {
        BuyBot bot = BuyBot(payable(buyBot));
        uint256 balance = bot.getBalance(token);

        console.log("==========================================");
        console.log("BuyBot Balance");
        console.log("------------------------------------------");
        console.log("Token:", token);
        console.log("Balance:", balance);
        console.log("==========================================");
    }

    /**
     * @notice Deploy BuyBot with default settings for testing
     * @dev Default: 0 delay, 100e18 min order, 0 interval, owner as recipient/buyer/manager
     */
    function deployBuyBotDefault(address owner, address router) external returns (address) {
        vm.broadcast();
        BuyBot bot = new BuyBot(
            0, // initialDelay: 0 (no delay)
            owner, // owner
            router, // router
            100e18, // minOrderAmount: 100 tokens
            0, // interval: 0 (disabled)
            owner, // recipient: owner
            owner, // buyer: owner
            owner // manager: owner
        );
        return address(bot);
    }
}
