// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin-contracts-5.2.0/access/Ownable.sol";
import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/utils/SafeERC20.sol";

import {IPair} from "./interfaces/IPair.sol";
import {IRouter} from "./interfaces/IRouter.sol";

/**
 * @title BuyBot
 * @notice A bot contract that automatically buys tokens at market price using its balance
 * @dev Only owner or authorized buyer can trigger the buy function
 */
contract BuyBot is Ownable {
    using SafeERC20 for IERC20;

    /// @notice Address representing native coin (ETH)
    address public constant NATIVE_COIN = address(1);

    error BuyBotInsufficientBalance(uint256 balance, uint256 minOrderAmount);
    error BuyBotInvalidRouter(address);
    error BuyBotInvalidPair(address);
    error BuyBotInvalidMinOrderAmount(uint256);
    error BuyBotIntervalNotPassed(uint256 timeSinceLastBuy, uint256 requiredInterval);
    error BuyBotUnauthorizedCaller(address caller);

    event MarketBuyExecuted(
        address indexed pair,
        address indexed quoteToken,
        address indexed baseToken,
        uint256 quoteAmount,
        address executor
    );
    event MinOrderAmountSet(uint256 indexed before, uint256 indexed current);
    event IntervalSet(uint256 indexed before, uint256 indexed current);
    event RecipientSet(address indexed before, address indexed current);
    event BuyerSet(address indexed before, address indexed current);
    event Withdrawn(address indexed token, address indexed to, uint256 amount);

    /// @notice CrossDexRouter address
    IRouter public router;

    /// @notice Minimum order amount required to execute market buy
    uint256 public minOrderAmount;

    /// @notice Minimum time interval (in seconds) between buy executions
    uint256 public interval;

    /// @notice Timestamp of the last buyMarket execution
    uint256 public lastBuyTime;

    /// @notice Address to receive purchased BASE tokens
    address public recipient;

    /// @notice Address authorized to execute buyMarket (in addition to owner)
    address public buyer;

    /**
     * @notice Contract constructor
     * @param _owner Owner address
     * @param _router CrossDexRouter address
     * @param _minOrderAmount Minimum order amount in QUOTE token
     * @param _interval Minimum time interval (in seconds) between buy executions
     * @param _recipient Address to receive purchased BASE tokens
     * @param _buyer Address authorized to execute buyMarket
     */
    constructor(
        address _owner,
        address _router,
        uint256 _minOrderAmount,
        uint256 _interval,
        address _recipient,
        address _buyer
    ) Ownable(_owner) {
        if (_router == address(0)) revert BuyBotInvalidRouter(_router);
        if (_minOrderAmount == 0) revert BuyBotInvalidMinOrderAmount(_minOrderAmount);

        router = IRouter(_router);
        minOrderAmount = _minOrderAmount;
        interval = _interval;
        recipient = _recipient;
        buyer = _buyer;

        emit MinOrderAmountSet(0, _minOrderAmount);
        emit IntervalSet(0, _interval);
        emit RecipientSet(address(0), _recipient);
        emit BuyerSet(address(0), _buyer);
    }

    // ===== MODIFIERS =====

    /**
     * @notice Only owner or authorized buyer can call
     */
    modifier onlyBuyer() {
        if (msg.sender != owner() && msg.sender != buyer) revert BuyBotUnauthorizedCaller(msg.sender);
        _;
    }

    // ===== PUBLIC FUNCTIONS =====

    /**
     * @notice Execute market buy order
     * @dev Can only be called by owner or authorized buyer
     * @param pair Trading pair address
     * @param amount Amount to spend (0 for entire balance)
     * @param maxMatchCount Maximum number of orders to match (0 for router default)
     */
    function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) external onlyBuyer {
        if (pair == address(0)) revert BuyBotInvalidPair(pair);

        // Check interval has passed since last buy
        if (interval > 0 && lastBuyTime > 0) {
            uint256 timeSinceLastBuy = block.timestamp - lastBuyTime;
            if (timeSinceLastBuy < interval) revert BuyBotIntervalNotPassed(timeSinceLastBuy, interval);
        }

        // Get pair configuration
        IPair.Config memory config = IPair(pair).getConfig();
        IERC20 quoteToken = config.QUOTE;
        IERC20 baseToken = config.BASE;

        // Get current balance
        uint256 balance = quoteToken.balanceOf(address(this));

        // Determine amount to spend (0 = entire balance)
        uint256 spendAmount = amount == 0 ? balance : amount;

        // Check minimum order amount
        if (spendAmount < minOrderAmount) revert BuyBotInsufficientBalance(spendAmount, minOrderAmount);

        // Check sufficient balance
        if (spendAmount > balance) revert BuyBotInsufficientBalance(balance, spendAmount);

        // Approve router if needed (only once per token)
        address routerAddress = address(router);
        if (quoteToken.allowance(address(this), routerAddress) < spendAmount) {
            quoteToken.forceApprove(routerAddress, type(uint256).max);
        }

        // Execute market buy
        router.submitBuyMarket(pair, spendAmount, maxMatchCount);

        emit MarketBuyExecuted(pair, address(quoteToken), address(baseToken), spendAmount, msg.sender);

        // Update last buy time
        lastBuyTime = block.timestamp;

        // Transfer BASE tokens to recipient if specified
        if (recipient != address(0)) {
            // Check if WETH was auto-unwrapped to ETH
            uint256 ethBalance = address(this).balance;
            if (ethBalance > 0) {
                // Transfer ETH to recipient
                (bool success,) = recipient.call{value: ethBalance}("");
                require(success, "ETH transfer failed");
            }

            // Transfer any remaining ERC20 BASE tokens
            uint256 baseBalance = baseToken.balanceOf(address(this));
            if (baseBalance > 0) baseToken.safeTransfer(recipient, baseBalance);
        }
    }

    // ===== OWNER ONLY FUNCTIONS =====

    /**
     * @notice Withdraw tokens from the contract
     * @dev Only owner can withdraw
     * @param token Token address to withdraw
     * @param amount Amount to withdraw (0 for entire balance)
     */
    function withdraw(address token, uint256 amount) external onlyOwner {
        IERC20 tokenContract = IERC20(token);
        uint256 balance = tokenContract.balanceOf(address(this));

        uint256 withdrawAmount = amount == 0 ? balance : amount;
        require(withdrawAmount <= balance, "Insufficient balance");

        tokenContract.safeTransfer(owner(), withdrawAmount);
        emit Withdrawn(token, owner(), withdrawAmount);
    }

    /**
     * @notice Withdraw native ETH from the contract
     * @dev Only owner can withdraw
     * @param amount Amount to withdraw (0 for entire balance)
     */
    function withdrawETH(uint256 amount) external onlyOwner {
        uint256 balance = address(this).balance;
        uint256 withdrawAmount = amount == 0 ? balance : amount;
        require(withdrawAmount <= balance, "Insufficient balance");

        (bool success,) = owner().call{value: withdrawAmount}("");
        require(success, "ETH transfer failed");
        emit Withdrawn(NATIVE_COIN, owner(), withdrawAmount);
    }

    /**
     * @notice Set minimum order amount
     * @dev Only owner can set
     * @param _minOrderAmount New minimum order amount
     */
    function setMinOrderAmount(uint256 _minOrderAmount) external onlyOwner {
        if (_minOrderAmount == 0) revert BuyBotInvalidMinOrderAmount(_minOrderAmount);
        emit MinOrderAmountSet(minOrderAmount, _minOrderAmount);
        minOrderAmount = _minOrderAmount;
    }

    /**
     * @notice Set time interval between buy executions
     * @dev Only owner can set
     * @param _interval New interval in seconds (0 to disable)
     */
    function setInterval(uint256 _interval) external onlyOwner {
        emit IntervalSet(interval, _interval);
        interval = _interval;
    }

    /**
     * @notice Set recipient address for purchased BASE tokens
     * @dev Only owner can set
     * @param _recipient New recipient address (address(0) to keep in contract)
     */
    function setRecipient(address _recipient) external onlyOwner {
        emit RecipientSet(recipient, _recipient);
        recipient = _recipient;
    }

    /**
     * @notice Set authorized buyer address
     * @dev Only owner can set
     * @param _buyer New buyer address
     */
    function setBuyer(address _buyer) external onlyOwner {
        emit BuyerSet(buyer, _buyer);
        buyer = _buyer;
    }

    // ===== VIEW FUNCTIONS =====

    /**
     * @notice Check if market buy can be executed
     * @param pair Trading pair address
     * @param caller Address of the caller to check authorization
     * @return canBuy Whether the buy can be executed (authorized, balance sufficient, and interval passed)
     * @return balance Current QUOTE token balance
     */
    function canBuyMarket(address pair, address caller) external view returns (bool canBuy, uint256 balance) {
        if (pair == address(0)) return (false, 0);

        // Check authorization
        if (caller != owner() && caller != buyer) return (false, 0);

        IPair.Config memory config = IPair(pair).getConfig();
        balance = config.QUOTE.balanceOf(address(this));

        // Check balance requirement
        if (balance < minOrderAmount) return (false, balance);

        // Check interval requirement
        if (interval > 0 && lastBuyTime > 0) {
            uint256 timeSinceLastBuy = block.timestamp - lastBuyTime;
            if (timeSinceLastBuy < interval) return (false, balance);
        }

        canBuy = true;
    }

    /**
     * @notice Get token balance
     * @param token Token address (NATIVE_COIN for ETH)
     * @return balance Token balance
     */
    function getBalance(address token) external view returns (uint256 balance) {
        if (token == NATIVE_COIN) return address(this).balance;
        return IERC20(token).balanceOf(address(this));
    }

    /**
     * @notice Receive ETH
     */
    receive() external payable {}
}
