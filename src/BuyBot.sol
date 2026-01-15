// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {AccessControlDefaultAdminRules} from
    "@openzeppelin-contracts-5.2.0/access/extensions/AccessControlDefaultAdminRules.sol";
import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin-contracts-5.2.0/utils/ReentrancyGuard.sol";

import {IPair} from "./interfaces/IPair.sol";
import {IRouter} from "./interfaces/IRouter.sol";
import {ISwapRouter} from "./interfaces/ISwapRouter.sol";
import {IUniswapV3Pool} from "./interfaces/IUniswapV3Pool.sol";
import {TickMath} from "./lib/TickMath.sol";

/**
 * @title BuyBot
 * @notice A bot contract that automatically buys tokens at market price using its balance
 * @dev Uses AccessControlDefaultAdminRules for role-based permissions with admin transfer delay
 * @dev Roles: DEFAULT_ADMIN_ROLE (owner), BUYER_ROLE, MANAGER_ROLE
 * @dev Uses ReentrancyGuard to prevent reentrancy attacks during external calls
 */
contract BuyBot is AccessControlDefaultAdminRules, ReentrancyGuard {
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
    error BuyBotNotChanged(uint256 oldValue, uint256 newValue);
    error BuyBotAddressNotChanged(address oldValue, address newValue);
    error BuyBotInvalidAmount(uint256);
    error BuyBotIntervalNotPassed(uint256 timeSinceLastBuy, uint256 requiredInterval);
    error BuyBotInvalidBuyer(address buyer);
    error BuyBotInvalidManager(address manager);
    error BuyBotInvalidSwapRouter(address);
    error BuyBotInvalidTickSlippage(uint24);
    error BuyBotNoSwapToken();
    error BuyBotPoolNotFound(address tokenIn, address tokenOut);
    error BuyBotInsufficientSwapBalance(address token, uint256 balance);
    error BuyBotInvalidTokenAddresses();
    error BuyBotInvalidPool(address pool);
    error BuyBotETHTransferFailed();
    error BuyBotInsufficientWithdrawBalance();
    error BuyBotInvalidPoolTokens(address pool, address tokenIn, address tokenOut);

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
    event SwapExecuted(address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut);
    event SwapPoolSet(address indexed tokenIn, address indexed tokenOut, address indexed pool);
    event SwapTokenSet(address indexed before, address indexed current);
    event MaxTickSlippageSet(uint24 before, uint24 current);
    event SwapRouterSet(address before, address current);

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

    /// @notice Uniswap V3 SwapRouter address
    ISwapRouter public swapRouter;

    /// @notice Maximum tick slippage for swaps (1 tick = 0.01% price movement)
    uint24 public maxTickSlippage;

    /// @notice Token to swap from (e.g., WETH, USDT, etc.)
    address public swapToken;

    /// @notice Pool registry: tokenIn => tokenOut => poolAddress
    mapping(address => mapping(address => address)) public swapPools;

    /**
     * @notice Contract constructor
     * @param _initialDelay Delay (in seconds) for admin transfer (0 for no delay)
     * @param _owner Owner address (gets DEFAULT_ADMIN_ROLE)
     * @param _router CrossDexRouter address
     * @param _minOrderAmount Minimum order amount in QUOTE token
     * @param _interval Minimum time interval (in seconds) between buy executions
     * @param _recipient Address to receive purchased BASE tokens
     * @param _buyer Address authorized to execute buyMarket (gets BUYER_ROLE)
     * @param _manager Address authorized to set minOrderAmount and interval (gets MANAGER_ROLE)
     * @param _swapRouter Uniswap V3 SwapRouter address (address(0) to disable swapping)
     * @param _maxTickSlippage Maximum tick slippage for swaps (1 tick = 0.01% price movement)
     */
    constructor(
        uint48 _initialDelay,
        address _owner,
        address _router,
        uint256 _minOrderAmount,
        uint256 _interval,
        address _recipient,
        address _buyer,
        address _manager,
        address _swapRouter,
        uint24 _maxTickSlippage
    ) AccessControlDefaultAdminRules(_initialDelay, _owner) {
        if (_router == address(0)) revert BuyBotInvalidRouter(_router);
        if (_minOrderAmount == 0) revert BuyBotInvalidMinOrderAmount(_minOrderAmount);
        if (_buyer == address(0)) revert BuyBotInvalidBuyer(_buyer);
        if (_manager == address(0)) revert BuyBotInvalidManager(_manager);
        if (_maxTickSlippage == 0) revert BuyBotInvalidTickSlippage(_maxTickSlippage);

        router = IRouter(_router);
        minOrderAmount = _minOrderAmount;
        interval = _interval;
        recipient = _recipient;
        swapRouter = ISwapRouter(_swapRouter);
        maxTickSlippage = _maxTickSlippage;

        // Grant BUYER_ROLE to owner and buyer
        _grantRole(BUYER_ROLE, _owner);
        _grantRole(BUYER_ROLE, _buyer);

        // Grant MANAGER_ROLE to owner and manager
        _grantRole(MANAGER_ROLE, _owner);
        _grantRole(MANAGER_ROLE, _manager);

        emit MinOrderAmountSet(0, _minOrderAmount);
        emit IntervalSet(0, _interval);
        emit RecipientSet(address(0), _recipient);
        emit SwapRouterSet(address(0), _swapRouter);
        emit MaxTickSlippageSet(0, _maxTickSlippage);
    }

    // ===== PUBLIC FUNCTIONS =====

    /**
     * @notice Execute market buy order
     * @dev Can only be called by owner or authorized buyer
     * @dev Protected against reentrancy attacks
     * @param pair Trading pair address
     * @param amount Amount to spend (must be greater than 0)
     * @param maxMatchCount Maximum number of orders to match (0 for router default)
     */
    function buyMarket(address pair, uint256 amount, uint256 maxMatchCount)
        external
        nonReentrant
        onlyRole(BUYER_ROLE)
    {
        if (pair == address(0)) revert BuyBotInvalidPair(pair);
        // Get pair configuration
        IPair.Config memory config = IPair(pair).getConfig();
        IERC20 quoteToken = config.QUOTE;
        IERC20 baseToken = config.BASE;

        // Check minimum order amount
        if (amount < minOrderAmount) revert BuyBotInsufficientBalance(amount, minOrderAmount);

        // Get current balance
        uint256 balance = quoteToken.balanceOf(address(this));

        // Check sufficient balance
        if (amount > balance) revert BuyBotInsufficientBalance(balance, amount);

        // Check interval has passed since last buy
        if (interval > 0 && lastBuyTime > 0) {
            uint256 timeSinceLastBuy = block.timestamp - lastBuyTime;
            if (timeSinceLastBuy < interval) revert BuyBotIntervalNotPassed(timeSinceLastBuy, interval);
        }

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
                if (!success) revert BuyBotETHTransferFailed();
            }

            // Transfer any remaining ERC20 BASE tokens
            uint256 baseBalance = baseToken.balanceOf(address(this));
            if (baseBalance > 0) baseToken.safeTransfer(recipient, baseBalance);
        }
    }

    /**
     * @notice Swap configured swap token to the pair's quote token
     * @dev Can only be called by owner or authorized buyer
     * @dev Protected against reentrancy attacks
     * @dev Swaps all balance of swapToken to the pair's quote token
     * @dev Uses maxTickSlippage for precise tick-based price slippage control
     * @param pair Trading pair address (to determine quote token)
     * @param uniswapFee Uniswap V3 pool fee tier (500 = 0.05%, 3000 = 0.3%, 10000 = 1%)
     * @return amountOut Actual amount of tokens received from swap
     */
    function swapToQuote(address pair, uint24 uniswapFee)
        external
        nonReentrant
        onlyRole(BUYER_ROLE)
        returns (uint256 amountOut)
    {
        if (pair == address(0)) revert BuyBotInvalidPair(pair);
        if (address(swapRouter) == address(0)) revert BuyBotInvalidSwapRouter(address(0));
        if (swapToken == address(0)) revert BuyBotNoSwapToken();

        // Get quote token from pair
        IPair.Config memory config = IPair(pair).getConfig();
        address quoteToken = address(config.QUOTE);

        // Check balance
        uint256 balance = IERC20(swapToken).balanceOf(address(this));
        if (balance == 0) revert BuyBotInsufficientSwapBalance(swapToken, 0);

        // Verify pool exists
        address pool = swapPools[swapToken][quoteToken];
        if (pool == address(0)) revert BuyBotPoolNotFound(swapToken, quoteToken);

        // Approve swap router
        IERC20 tokenIn = IERC20(swapToken);
        address swapRouterAddress = address(swapRouter);
        if (tokenIn.allowance(address(this), swapRouterAddress) < balance) {
            tokenIn.forceApprove(swapRouterAddress, type(uint256).max);
        }

        // Calculate sqrtPriceLimitX96 based on maxTickSlippage
        uint160 sqrtPriceLimitX96 = _calculateSqrtPriceLimit(pool, swapToken);

        // Execute swap with tick-based price limit
        ISwapRouter.ExactInputSingleParams memory params = ISwapRouter.ExactInputSingleParams({
            tokenIn: swapToken,
            tokenOut: quoteToken,
            fee: uniswapFee,
            recipient: address(this),
            deadline: block.timestamp,
            amountIn: balance,
            amountOutMinimum: 0, // Price limit enforced by sqrtPriceLimitX96
            sqrtPriceLimitX96: sqrtPriceLimitX96
        });

        amountOut = swapRouter.exactInputSingle(params);

        emit SwapExecuted(swapToken, quoteToken, balance, amountOut);
    }

    /**
     * @notice Calculate sqrtPriceLimitX96 based on current pool price and maxTickSlippage
     * @param pool Uniswap V3 pool address
     * @param tokenIn Input token address
     * @return sqrtPriceLimitX96 The price limit for the swap
     */
    function _calculateSqrtPriceLimit(address pool, address tokenIn) internal view returns (uint160) {
        // Get current tick from pool
        (, int24 currentTick,,,,,) = IUniswapV3Pool(pool).slot0();

        // Determine swap direction
        address token0 = IUniswapV3Pool(pool).token0();
        bool zeroForOne = (tokenIn == token0);

        // Calculate target tick with slippage
        // zeroForOne: price decreases, so tick decreases
        // oneForZero: price increases, so tick increases
        int24 targetTick;
        if (zeroForOne) {
            targetTick = currentTick - int24(maxTickSlippage);
            // Ensure we don't go below MIN_TICK
            if (targetTick < TickMath.MIN_TICK) targetTick = TickMath.MIN_TICK;
        } else {
            targetTick = currentTick + int24(maxTickSlippage);
            // Ensure we don't exceed MAX_TICK
            if (targetTick > TickMath.MAX_TICK) targetTick = TickMath.MAX_TICK;
        }

        return TickMath.getSqrtRatioAtTick(targetTick);
    }

    // ===== OWNER ONLY FUNCTIONS =====

    /**
     * @notice Withdraw tokens from the contract
     * @dev Only owner can withdraw
     * @dev Protected against reentrancy attacks
     * @param token Token address to withdraw
     * @param amount Amount to withdraw (0 for entire balance)
     */
    function withdraw(address token, uint256 amount) external nonReentrant onlyRole(DEFAULT_ADMIN_ROLE) {
        IERC20 tokenContract = IERC20(token);
        uint256 balance = tokenContract.balanceOf(address(this));

        uint256 withdrawAmount = amount == 0 ? balance : amount;
        if (withdrawAmount > balance) revert BuyBotInsufficientWithdrawBalance();

        tokenContract.safeTransfer(defaultAdmin(), withdrawAmount);
        emit Withdrawn(token, defaultAdmin(), withdrawAmount);
    }

    /**
     * @notice Withdraw native ETH from the contract
     * @dev Only owner can withdraw
     * @dev Protected against reentrancy attacks
     * @param amount Amount to withdraw (0 for entire balance)
     */
    function withdrawETH(uint256 amount) external nonReentrant onlyRole(DEFAULT_ADMIN_ROLE) {
        uint256 balance = address(this).balance;
        uint256 withdrawAmount = amount == 0 ? balance : amount;
        if (withdrawAmount > balance) revert BuyBotInsufficientWithdrawBalance();

        (bool success,) = defaultAdmin().call{value: withdrawAmount}("");
        if (!success) revert BuyBotETHTransferFailed();
        emit Withdrawn(NATIVE_COIN, defaultAdmin(), withdrawAmount);
    }

    /**
     * @notice Set minimum order amount
     * @dev Only owner or manager can set
     * @param _minOrderAmount New minimum order amount
     */
    function setMinOrderAmount(uint256 _minOrderAmount) external onlyRole(MANAGER_ROLE) {
        if (_minOrderAmount == 0) revert BuyBotInvalidMinOrderAmount(_minOrderAmount);
        if (minOrderAmount == _minOrderAmount) revert BuyBotNotChanged(minOrderAmount, _minOrderAmount);
        uint256 oldValue = minOrderAmount;
        minOrderAmount = _minOrderAmount;
        emit MinOrderAmountSet(oldValue, _minOrderAmount);
    }

    /**
     * @notice Set time interval between buy executions
     * @dev Only owner or manager can set
     * @param _interval New interval in seconds (0 to disable)
     */
    function setInterval(uint256 _interval) external onlyRole(MANAGER_ROLE) {
        if (interval == _interval) revert BuyBotNotChanged(interval, _interval);
        uint256 oldValue = interval;
        interval = _interval;
        emit IntervalSet(oldValue, _interval);
    }

    /**
     * @notice Set recipient address for purchased BASE tokens
     * @dev Only owner can set. Setting to address(0) keeps purchased tokens in the contract.
     * @param _recipient New recipient address (address(0) to keep in contract)
     */
    function setRecipient(address _recipient) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (recipient == _recipient) revert BuyBotAddressNotChanged(recipient, _recipient);
        address oldValue = recipient;
        recipient = _recipient;
        emit RecipientSet(oldValue, _recipient);
    }

    // ===== SWAP CONFIGURATION FUNCTIONS =====

    /**
     * @notice Set swap pool for token pair
     * @dev Only manager can set pools. Validates that pool tokens match tokenIn/tokenOut.
     * @param tokenIn Input token address
     * @param tokenOut Output token address
     * @param pool Uniswap V3 pool address (address(0) to remove)
     */
    function setSwapPool(address tokenIn, address tokenOut, address pool) external onlyRole(MANAGER_ROLE) {
        if (tokenIn == address(0) || tokenOut == address(0)) revert BuyBotInvalidTokenAddresses();
        if (pool == address(0)) revert BuyBotInvalidPool(pool);

        // Verify pool tokens match tokenIn and tokenOut (if pool is not address(0))
        address poolToken0 = IUniswapV3Pool(pool).token0();
        address poolToken1 = IUniswapV3Pool(pool).token1();

        // Check if tokens match in either order (token0/token1 are sorted by address)
        bool isValid =
            (tokenIn == poolToken0 && tokenOut == poolToken1) || (tokenIn == poolToken1 && tokenOut == poolToken0);
        if (!isValid) revert BuyBotInvalidPoolTokens(pool, tokenIn, tokenOut);

        swapPools[tokenIn][tokenOut] = pool;
        emit SwapPoolSet(tokenIn, tokenOut, pool);
    }

    /**
     * @notice Set swap token address
     * @dev Only manager can set. Setting to address(0) disables swapping.
     * @param _swapToken Token address to swap from (address(0) to disable)
     */
    function setSwapToken(address _swapToken) external onlyRole(MANAGER_ROLE) {
        if (swapToken == _swapToken) revert BuyBotAddressNotChanged(swapToken, _swapToken);
        address oldValue = swapToken;
        swapToken = _swapToken;
        emit SwapTokenSet(oldValue, _swapToken);
    }

    /**
     * @notice Set maximum tick slippage for swaps
     * @dev Only manager can set. 1 tick = 0.01% price movement.
     * @param _maxTickSlippage New maximum tick slippage (1 = 1 tick = 0.01%)
     */
    function setMaxTickSlippage(uint24 _maxTickSlippage) external onlyRole(MANAGER_ROLE) {
        if (_maxTickSlippage == 0) revert BuyBotInvalidTickSlippage(_maxTickSlippage);
        if (maxTickSlippage == _maxTickSlippage) revert BuyBotNotChanged(maxTickSlippage, _maxTickSlippage);
        uint24 oldValue = maxTickSlippage;
        maxTickSlippage = _maxTickSlippage;
        emit MaxTickSlippageSet(oldValue, _maxTickSlippage);
    }

    /**
     * @notice Set Uniswap V3 SwapRouter address
     * @dev Only owner can set. Use address(0) to disable swapping functionality.
     * @param _swapRouter New SwapRouter address (address(0) to disable)
     */
    function setSwapRouter(address _swapRouter) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (address(swapRouter) == _swapRouter) revert BuyBotAddressNotChanged(address(swapRouter), _swapRouter);
        address oldValue = address(swapRouter);
        swapRouter = ISwapRouter(_swapRouter);
        emit SwapRouterSet(oldValue, _swapRouter);
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
        if (!hasRole(BUYER_ROLE, caller)) return (false, 0);

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
