/**
 * Router Order Script Example (English)
 * 
 * TypeScript example for placing orders using CrossDex Router
 * Low-level example code using ethers.js
 * 
 * Usage:
 * 1. npm install ethers
 * 2. Set PRIVATE_KEY environment variable (export PRIVATE_KEY=your_private_key)
 * 3. ts-node router-order-typescript-en.ts
 */

import { ethers } from 'ethers';

// ===== Contract Addresses (Mainnet) =====
const ROUTER = '0x6690844Aac584AcA982E195B7BDeBd48740fbcb1';
const GAME_MARKET = '0xa0f50f79615247530fABcC3efd79B8e5b961b966';
const USDTX_MARKET = '0xB7811907b2839d6b5CCF908D6B58dE944D8AfbA7';
const VERSE8_MARKET = '0xcb95777d0f8d2EfA5e836Cb65f814dF8C7261d83';
const WETH = '0x52D3256c7d6C7522C6D593b2aC662dBF610E6813';

// RPC URL
const RPC_URL = 'https://mainnet.crosstoken.io:22001';

// ===== Order Constraints =====
const GOOD_TILL_CANCEL = 0;
const IMMEDIATE_OR_CANCEL = 1;
const FILL_OR_KILL = 2;

// ===== Interface ABIs (Only required functions) =====

// IRouter interface
const IRouterABI = [
    'function submitBuyLimit(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] calldata adjacent, uint256 maxMatchCount) external payable returns (uint256)',
    'function submitSellLimit(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] memory adjacent, uint256 maxMatchCount) external payable returns (uint256)',
    'function submitBuyMarket(address pair, uint256 quoteVolume, uint256 maxMatchCount) external payable',
    'function submitSellMarket(address pair, uint256 amount, uint256 maxMatchCount) external payable',
    'function cancelOrder(address pair, uint256[] calldata orderIds) external',
    'function getRequiredBuyVolume(address pair, uint256 quoteVolume) external view returns (uint256)',
    'function isPair(address pair) external view returns (bool)',
];

// IMarket interface
const IMarketABI = [
    'function baseToPair(address base) external view returns (address)',
    'function allPairs() external view returns (address[] memory bases, address[] memory pairs)',
    'function QUOTE() external view returns (address)',
];

// IERC20 interface
const IERC20ABI = [
    'function approve(address spender, uint256 amount) external returns (bool)',
    'function allowance(address owner, address spender) external view returns (uint256)',
    'function balanceOf(address account) external view returns (uint256)',
    'function transfer(address to, uint256 amount) external returns (bool)',
];

// ===== Utility Functions =====

/**
 * Get private key
 * From environment variable or direct input (dev only)
 */
function getPrivateKey(): string {
    // Method 1: From environment variable (recommended)
    const privateKey = process.env.PRIVATE_KEY;
    if (privateKey) {
        return privateKey;
    }
    
    // Method 2: Direct input (dev only)
    // ⚠️ Warning: Never use in production!
    // return 'your_private_key_here';
    
    throw new Error('PRIVATE_KEY environment variable is not set');
}

/**
 * Create Provider and Wallet
 */
function createWallet(): { provider: ethers.JsonRpcProvider; wallet: ethers.Wallet } {
    const provider = new ethers.JsonRpcProvider(RPC_URL);
    const privateKey = getPrivateKey();
    const wallet = new ethers.Wallet(privateKey, provider);
    
    console.log('Wallet address:', wallet.address);
    return { provider, wallet };
}

/**
 * Approve ERC20 token for Router
 * @param tokenAddress Token address
 * @param amount Amount to approve (0 for unlimited)
 */
async function approveToken(
    wallet: ethers.Wallet,
    tokenAddress: string,
    amount: bigint = BigInt(0)
): Promise<void> {
    const tokenContract = new ethers.Contract(tokenAddress, IERC20ABI, wallet);
    
    const approveAmount = amount === BigInt(0) 
        ? ethers.MaxUint256 
        : amount;
    
    // Check current allowance
    const currentAllowance = await tokenContract.allowance(
        wallet.address,
        ROUTER
    );
    
    if (currentAllowance < approveAmount) {
        console.log(`Approving ${tokenAddress}...`);
        const tx = await tokenContract.approve(ROUTER, approveAmount);
        console.log('Approve transaction hash:', tx.hash);
        
        const receipt = await tx.wait();
        console.log('Approve confirmed in block:', receipt.blockNumber);
    } else {
        console.log('Token already approved with sufficient amount');
    }
}

/**
 * Get Pair address from Market
 * @param marketAddress Market address
 * @param baseToken BASE token address
 */
async function getPair(
    provider: ethers.JsonRpcProvider,
    marketAddress: string,
    baseToken: string
): Promise<string> {
    const marketContract = new ethers.Contract(marketAddress, IMarketABI, provider);
    const pair = await marketContract.baseToPair(baseToken);
    
    if (pair === ethers.ZeroAddress) {
        throw new Error('Pair not found');
    }
    
    console.log('Pair address:', pair);
    return pair;
}

/**
 * Submit buy limit order
 * @param wallet Wallet instance
 * @param pair Pair address
 * @param price Limit price
 * @param amount Order amount
 * @param constraints Order constraints (0: GTC, 1: IOC, 2: FOK)
 * @param maxMatchCount Maximum match count (0 for router default)
 */
async function submitBuyLimitOrder(
    wallet: ethers.Wallet,
    pair: string,
    price: bigint,
    amount: bigint,
    constraints: number = GOOD_TILL_CANCEL,
    maxMatchCount: bigint = BigInt(0)
): Promise<bigint> {
    const routerContract = new ethers.Contract(ROUTER, IRouterABI, wallet);
    
    // Validate pair
    const isValidPair = await routerContract.isPair(pair);
    if (!isValidPair) {
        throw new Error('Invalid pair address');
    }
    
    // Calculate required QUOTE amount including fee
    const quoteVolume = (price * amount) / BigInt(10 ** 18);
    const requiredQuote = await routerContract.getRequiredBuyVolume(pair, quoteVolume);
    console.log('Required QUOTE (including fee):', ethers.formatEther(requiredQuote));
    
    // adjacent parameter (previous price search range)
    const adjacent: [bigint, bigint] = [BigInt(0), BigInt(0)];
    
    // Send transaction
    console.log('Submitting buy limit order...');
    const tx = await routerContract.submitBuyLimit(
        pair,
        price,
        amount,
        constraints,
        adjacent,
        maxMatchCount,
        { value: 0 }
    );
    
    console.log('Transaction hash:', tx.hash);
    const receipt = await tx.wait();
    console.log('Transaction confirmed in block:', receipt.blockNumber);
    
    // Extract Order ID (from transaction logs)
    // In practice, you need to parse events, but shown as example only
    const orderId = receipt.logs.length > 0 ? BigInt(receipt.logs[0].topics[1]) : BigInt(0);
    console.log('Order ID:', orderId.toString());
    
    return orderId;
}

/**
 * Submit sell limit order
 */
async function submitSellLimitOrder(
    wallet: ethers.Wallet,
    pair: string,
    price: bigint,
    amount: bigint,
    constraints: number = GOOD_TILL_CANCEL,
    maxMatchCount: bigint = BigInt(0)
): Promise<bigint> {
    const routerContract = new ethers.Contract(ROUTER, IRouterABI, wallet);
    
    const isValidPair = await routerContract.isPair(pair);
    if (!isValidPair) {
        throw new Error('Invalid pair address');
    }
    
    const adjacent: [bigint, bigint] = [BigInt(0), BigInt(0)];
    
    console.log('Submitting sell limit order...');
    const tx = await routerContract.submitSellLimit(
        pair,
        price,
        amount,
        constraints,
        adjacent,
        maxMatchCount,
        { value: 0 }
    );
    
    console.log('Transaction hash:', tx.hash);
    const receipt = await tx.wait();
    console.log('Transaction confirmed in block:', receipt.blockNumber);
    
    const orderId = receipt.logs.length > 0 ? BigInt(receipt.logs[0].topics[1]) : BigInt(0);
    console.log('Order ID:', orderId.toString());
    
    return orderId;
}

/**
 * Submit buy market order
 */
async function submitBuyMarketOrder(
    wallet: ethers.Wallet,
    pair: string,
    quoteVolume: bigint,
    maxMatchCount: bigint = BigInt(0)
): Promise<void> {
    const routerContract = new ethers.Contract(ROUTER, IRouterABI, wallet);
    
    const isValidPair = await routerContract.isPair(pair);
    if (!isValidPair) {
        throw new Error('Invalid pair address');
    }
    
    // Calculate required QUOTE amount including fee
    const requiredQuote = await routerContract.getRequiredBuyVolume(pair, quoteVolume);
    console.log('Required QUOTE (including fee):', ethers.formatEther(requiredQuote));
    
    console.log('Submitting buy market order...');
    const tx = await routerContract.submitBuyMarket(
        pair,
        quoteVolume,
        maxMatchCount,
        { value: 0 }
    );
    
    console.log('Transaction hash:', tx.hash);
    const receipt = await tx.wait();
    console.log('Transaction confirmed in block:', receipt.blockNumber);
}

/**
 * Submit sell market order
 */
async function submitSellMarketOrder(
    wallet: ethers.Wallet,
    pair: string,
    amount: bigint,
    maxMatchCount: bigint = BigInt(0)
): Promise<void> {
    const routerContract = new ethers.Contract(ROUTER, IRouterABI, wallet);
    
    const isValidPair = await routerContract.isPair(pair);
    if (!isValidPair) {
        throw new Error('Invalid pair address');
    }
    
    console.log('Submitting sell market order...');
    const tx = await routerContract.submitSellMarket(
        pair,
        amount,
        maxMatchCount,
        { value: 0 }
    );
    
    console.log('Transaction hash:', tx.hash);
    const receipt = await tx.wait();
    console.log('Transaction confirmed in block:', receipt.blockNumber);
}

/**
 * Cancel orders
 */
async function cancelOrders(
    wallet: ethers.Wallet,
    pair: string,
    orderIds: bigint[]
): Promise<void> {
    const routerContract = new ethers.Contract(ROUTER, IRouterABI, wallet);
    
    const isValidPair = await routerContract.isPair(pair);
    if (!isValidPair) {
        throw new Error('Invalid pair address');
    }
    
    if (orderIds.length === 0) {
        throw new Error('No order IDs provided');
    }
    
    console.log(`Cancelling ${orderIds.length} orders...`);
    const tx = await routerContract.cancelOrder(pair, orderIds);
    
    console.log('Transaction hash:', tx.hash);
    const receipt = await tx.wait();
    console.log('Transaction confirmed in block:', receipt.blockNumber);
}

/**
 * Submit buy order with native CROSS
 */
async function submitBuyLimitWithNativeCROSS(
    wallet: ethers.Wallet,
    pair: string,
    price: bigint,
    amount: bigint,
    maxMatchCount: bigint = BigInt(0)
): Promise<bigint> {
    const routerContract = new ethers.Contract(ROUTER, IRouterABI, wallet);
    
    const isValidPair = await routerContract.isPair(pair);
    if (!isValidPair) {
        throw new Error('Invalid pair address');
    }
    
    // Calculate required QUOTE amount including fee
    const quoteVolume = (price * amount) / BigInt(10 ** 18);
    const requiredQuote = await routerContract.getRequiredBuyVolume(pair, quoteVolume);
    
    // Check wallet CROSS balance
    const balance = await wallet.provider.getBalance(wallet.address);
    if (balance < requiredQuote) {
        throw new Error(`Insufficient CROSS balance. Required: ${ethers.formatEther(requiredQuote)}, Available: ${ethers.formatEther(balance)}`);
    }
    
    const adjacent: [bigint, bigint] = [BigInt(0), BigInt(0)];
    
    console.log('Submitting buy limit order with native CROSS...');
    const tx = await routerContract.submitBuyLimit(
        pair,
        price,
        amount,
        GOOD_TILL_CANCEL,
        adjacent,
        maxMatchCount,
        { value: requiredQuote }
    );
    
    console.log('Transaction hash:', tx.hash);
    const receipt = await tx.wait();
    console.log('Transaction confirmed in block:', receipt.blockNumber);
    
    const orderId = receipt.logs.length > 0 ? BigInt(receipt.logs[0].topics[1]) : BigInt(0);
    console.log('Order ID:', orderId.toString());
    
    return orderId;
}

/**
 * Complete example: Approve token and submit buy limit order
 */
async function exampleBuyLimitOrder() {
    try {
        // 1. Create wallet
        const { provider, wallet } = createWallet();
        
        // 2. Set parameters (example)
        const marketAddress = GAME_MARKET;
        const baseToken = '0x...'; // BASE token address
        const quoteToken = '0x...'; // QUOTE token address
        const price = ethers.parseEther('100'); // Price (e.g., 100 CROSS)
        const amount = ethers.parseEther('1'); // Amount (e.g., 1 BASE token)
        
        // 3. Get Pair address
        const pair = await getPair(provider, marketAddress, baseToken);
        
        // 4. Approve QUOTE token (unlimited)
        await approveToken(wallet, quoteToken, BigInt(0));
        
        // 5. Submit buy limit order
        const orderId = await submitBuyLimitOrder(
            wallet,
            pair,
            price,
            amount,
            GOOD_TILL_CANCEL,
            BigInt(0)
        );
        
        console.log('✅ Buy limit order submitted successfully!');
        console.log('Order ID:', orderId.toString());
        
    } catch (error) {
        console.error('❌ Error:', error);
        throw error;
    }
}

// Main execution
if (require.main === module) {
    exampleBuyLimitOrder()
        .then(() => process.exit(0))
        .catch((error) => {
            console.error(error);
            process.exit(1);
        });
}

export {
    approveToken,
    getPair,
    submitBuyLimitOrder,
    submitSellLimitOrder,
    submitBuyMarketOrder,
    submitSellMarketOrder,
    cancelOrders,
    submitBuyLimitWithNativeCROSS,
};

