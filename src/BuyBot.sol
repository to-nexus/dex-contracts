// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {AccessControl} from "@openzeppelin-contracts-5.2.0/access/AccessControl.sol";
import {Ownable} from "@openzeppelin-contracts-5.2.0/access/Ownable.sol";
import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/utils/SafeERC20.sol";

import {IPair} from "./interfaces/IPair.sol";
import {IRouter} from "./interfaces/IRouter.sol";

/**
 * @title BuyBot
 * @notice A bot contract that automatically buys tokens at market price using its balance
 * @dev Uses AccessControl for role-based permissions (BUYER_ROLE, MANAGER_ROLE)
 */
contract BuyBot is Ownable, AccessControl {
    using SafeERC20 for IERC20;

    /// @notice Role for executing buyMarket function
    bytes32 public constant BUYER_ROLE = keccak256("BUYER_ROLE");

    /// @notice Role for setting minOrderAmount and interval
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");

    /// @notice Address representing native coin (ETH)
    address public constant NATIVE_COIN = address(1);

    error BuyBotInsufficientBalance(uint256 balance, uint256 minOrderAmount);
    error BuyBotInvalidRouter(address);
    error BuyBotInvalidPair(address);
    error BuyBotInvalidMinOrderAmount(uint256);
    error BuyBotInvalidAmount(uint256);
    error BuyBotIntervalNotPassed(uint256 timeSinceLastBuy, uint256 requiredInterval);
    error BuyBotUnauthorizedCaller(address caller);
    error BuyBotInvalidBuyer(address buyer);
    error BuyBotInvalidManager(address manager);

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

    /**
     * @notice Contract constructor
     * @param _owner Owner address (also gets DEFAULT_ADMIN_ROLE)
     * @param _router CrossDexRouter address
     * @param _minOrderAmount Minimum order amount in QUOTE token
     * @param _interval Minimum time interval (in seconds) between buy executions
     * @param _recipient Address to receive purchased BASE tokens
     * @param _buyer Address authorized to execute buyMarket (gets BUYER_ROLE)
     * @param _manager Address authorized to set minOrderAmount and interval (gets MANAGER_ROLE)
     */
    constructor(
        address _owner,
        address _router,
        uint256 _minOrderAmount,
        uint256 _interval,
        address _recipient,
        address _buyer,
        address _manager
    ) Ownable(_owner) {
        if (_router == address(0)) revert BuyBotInvalidRouter(_router);
        if (_minOrderAmount == 0) revert BuyBotInvalidMinOrderAmount(_minOrderAmount);
        if (_buyer == address(0)) revert BuyBotInvalidBuyer(_buyer);
        if (_manager == address(0)) revert BuyBotInvalidManager(_manager);

        router = IRouter(_router);
        minOrderAmount = _minOrderAmount;
        interval = _interval;
        recipient = _recipient;

        // Grant DEFAULT_ADMIN_ROLE to owner
        _grantRole(DEFAULT_ADMIN_ROLE, _owner);

        // Grant BUYER_ROLE
        _grantRole(BUYER_ROLE, _buyer);

        // Grant MANAGER_ROLE
        _grantRole(MANAGER_ROLE, _manager);

        emit MinOrderAmountSet(0, _minOrderAmount);
        emit IntervalSet(0, _interval);
        emit RecipientSet(address(0), _recipient);
    }

    // ===== MODIFIERS =====

    /**
     * @notice Only addresses with BUYER_ROLE or owner can call
     */
    modifier onlyBuyer() {
        if (!hasRole(BUYER_ROLE, msg.sender) && msg.sender != owner()) revert BuyBotUnauthorizedCaller(msg.sender);
        _;
    }

    /**
     * @notice Only addresses with MANAGER_ROLE or owner can call
     */
    modifier onlyManager() {
        if (!hasRole(MANAGER_ROLE, msg.sender) && msg.sender != owner()) revert BuyBotUnauthorizedCaller(msg.sender);
        _;
    }

    // ===== PUBLIC FUNCTIONS =====

    /**
     * @notice Execute market buy order
     * @dev Can only be called by owner or authorized buyer
     * @param pair Trading pair address
     * @param amount Amount to spend (must be greater than 0)
     * @param maxMatchCount Maximum number of orders to match (0 for router default)
     */
    function buyMarket(address pair, uint256 amount, uint256 maxMatchCount) external onlyBuyer {
        if (pair == address(0)) revert BuyBotInvalidPair(pair);
        if (amount == 0) revert BuyBotInvalidAmount(amount);

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

        // Check minimum order amount
        if (amount < minOrderAmount) revert BuyBotInsufficientBalance(amount, minOrderAmount);

        // Check sufficient balance
        if (amount > balance) revert BuyBotInsufficientBalance(balance, amount);

        // Approve router if needed (only once per token)
        address routerAddress = address(router);
        if (quoteToken.allowance(address(this), routerAddress) < amount) {
            quoteToken.forceApprove(routerAddress, type(uint256).max);
        }

        // Execute market buy
        router.submitBuyMarket(pair, amount, maxMatchCount);

        emit MarketBuyExecuted(pair, address(quoteToken), address(baseToken), amount, msg.sender);

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
     * @dev Only owner or manager can set
     * @param _minOrderAmount New minimum order amount
     */
    function setMinOrderAmount(uint256 _minOrderAmount) external onlyManager {
        if (_minOrderAmount == 0) revert BuyBotInvalidMinOrderAmount(_minOrderAmount);
        emit MinOrderAmountSet(minOrderAmount, _minOrderAmount);
        minOrderAmount = _minOrderAmount;
    }

    /**
     * @notice Set time interval between buy executions
     * @dev Only owner or manager can set
     * @param _interval New interval in seconds (0 to disable)
     */
    function setInterval(uint256 _interval) external onlyManager {
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
     * @notice Grant BUYER_ROLE to an address
     * @dev Only owner or admin can grant roles
     * @param account Address to grant the role
     */
    function grantBuyerRole(address account) external onlyOwner {
        grantRole(BUYER_ROLE, account);
    }

    /**
     * @notice Revoke BUYER_ROLE from an address
     * @dev Only owner or admin can revoke roles
     * @param account Address to revoke the role
     */
    function revokeBuyerRole(address account) external onlyOwner {
        revokeRole(BUYER_ROLE, account);
    }

    /**
     * @notice Grant MANAGER_ROLE to an address
     * @dev Only owner or admin can grant roles
     * @param account Address to grant the role
     */
    function grantManagerRole(address account) external onlyOwner {
        grantRole(MANAGER_ROLE, account);
    }

    /**
     * @notice Revoke MANAGER_ROLE from an address
     * @dev Only owner or admin can revoke roles
     * @param account Address to revoke the role
     */
    function revokeManagerRole(address account) external onlyOwner {
        revokeRole(MANAGER_ROLE, account);
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
        if (!hasRole(BUYER_ROLE, caller) && caller != owner()) return (false, 0);

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
