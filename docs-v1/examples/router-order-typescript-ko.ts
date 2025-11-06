/**
 * Router 주문 스크립트 예제 (한국어)
 * 
 * CrossDex Router를 사용하여 주문을 넣는 TypeScript 예제
 * ethers.js를 사용한 로우레벨 예제 코드
 * 
 * 사용법:
 * 1. npm install ethers
 * 2. 환경 변수에 PRIVATE_KEY 설정 (export PRIVATE_KEY=your_private_key)
 * 3. ts-node router-order-typescript-ko.ts
 */

import { ethers } from 'ethers';

// ===== 컨트랙트 주소 (Mainnet) =====
const ROUTER = '0x6690844Aac584AcA982E195B7BDeBd48740fbcb1';
const GAME_MARKET = '0xa0f50f79615247530fABcC3efd79B8e5b961b966';
const USDTX_MARKET = '0xB7811907b2839d6b5CCF908D6B58dE944D8AfbA7';
const VERSE8_MARKET = '0xcb95777d0f8d2EfA5e836Cb65f814dF8C7261d83';
const WETH = '0x52D3256c7d6C7522C6D593b2aC662dBF610E6813';

// RPC URL
const RPC_URL = 'https://mainnet.crosstoken.io:22001';

// ===== 주문 제약 조건 =====
const GOOD_TILL_CANCEL = 0;
const IMMEDIATE_OR_CANCEL = 1;
const FILL_OR_KILL = 2;

// ===== 인터페이스 ABI (필요한 함수만 선언) =====

// IRouter 인터페이스
const IRouterABI = [
    'function submitBuyLimit(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] calldata adjacent, uint256 maxMatchCount) external payable returns (uint256)',
    'function submitSellLimit(address pair, uint256 price, uint256 amount, uint8 constraints, uint256[2] memory adjacent, uint256 maxMatchCount) external payable returns (uint256)',
    'function submitBuyMarket(address pair, uint256 quoteVolume, uint256 maxMatchCount) external payable',
    'function submitSellMarket(address pair, uint256 amount, uint256 maxMatchCount) external payable',
    'function cancelOrder(address pair, uint256[] calldata orderIds) external',
    'function getRequiredBuyVolume(address pair, uint256 quoteVolume) external view returns (uint256)',
    'function isPair(address pair) external view returns (bool)',
];

// IMarket 인터페이스
const IMarketABI = [
    'function baseToPair(address base) external view returns (address)',
    'function allPairs() external view returns (address[] memory bases, address[] memory pairs)',
    'function QUOTE() external view returns (address)',
];

// IERC20 인터페이스
const IERC20ABI = [
    'function approve(address spender, uint256 amount) external returns (bool)',
    'function allowance(address owner, address spender) external view returns (uint256)',
    'function balanceOf(address account) external view returns (uint256)',
    'function transfer(address to, uint256 amount) external returns (bool)',
];

// ===== 유틸리티 함수 =====

/**
 * 프라이빗키 가져오기
 * 환경 변수에서 가져오거나 직접 입력 (개발 환경)
 */
function getPrivateKey(): string {
    // 방법 1: 환경 변수에서 가져오기 (권장)
    const privateKey = process.env.PRIVATE_KEY;
    if (privateKey) {
        return privateKey;
    }
    
    // 방법 2: 직접 입력 (개발 환경만)
    // ⚠️ 주의: 프로덕션에서는 절대 사용하지 마세요!
    // return 'your_private_key_here';
    
    throw new Error('PRIVATE_KEY environment variable is not set');
}

/**
 * Provider 및 Wallet 생성
 */
function createWallet(): { provider: ethers.JsonRpcProvider; wallet: ethers.Wallet } {
    const provider = new ethers.JsonRpcProvider(RPC_URL);
    const privateKey = getPrivateKey();
    const wallet = new ethers.Wallet(privateKey, provider);
    
    console.log('Wallet address:', wallet.address);
    return { provider, wallet };
}

/**
 * ERC20 토큰 Approve 처리
 * @param tokenAddress 토큰 주소
 * @param amount Approve할 양 (0이면 무제한)
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
    
    // 현재 allowance 확인
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
 * Market에서 Pair 주소 조회
 * @param marketAddress Market 주소
 * @param baseToken BASE 토큰 주소
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
 * 매수 지정가 주문
 * @param wallet Wallet 인스턴스
 * @param pair Pair 주소
 * @param price 지정 가격
 * @param amount 주문 수량
 * @param constraints 주문 제약 조건 (0: GTC, 1: IOC, 2: FOK)
 * @param maxMatchCount 최대 매칭 횟수 (0이면 Router 기본값 사용)
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
    
    // Pair 유효성 검증
    const isValidPair = await routerContract.isPair(pair);
    if (!isValidPair) {
        throw new Error('Invalid pair address');
    }
    
    // 필요한 수수료 포함 QUOTE 양 계산
    const quoteVolume = (price * amount) / BigInt(10 ** 18);
    const requiredQuote = await routerContract.getRequiredBuyVolume(pair, quoteVolume);
    console.log('Required QUOTE (including fee):', ethers.formatEther(requiredQuote));
    
    // adjacent 파라미터 (이전 가격 검색 범위)
    const adjacent: [bigint, bigint] = [BigInt(0), BigInt(0)];
    
    // 트랜잭션 전송
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
    
    // Order ID 추출 (트랜잭션 로그에서)
    // 실제로는 이벤트를 파싱해야 하지만, 간단히 예제로만 표시
    const orderId = receipt.logs.length > 0 ? BigInt(receipt.logs[0].topics[1]) : BigInt(0);
    console.log('Order ID:', orderId.toString());
    
    return orderId;
}

/**
 * 매도 지정가 주문
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
 * 매수 시장가 주문
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
    
    // 필요한 수수료 포함 QUOTE 양 계산
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
 * 매도 시장가 주문
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
 * 주문 취소
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
 * CROSS 네이티브 코인으로 매수 주문
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
    
    // 필요한 수수료 포함 QUOTE 양 계산
    const quoteVolume = (price * amount) / BigInt(10 ** 18);
    const requiredQuote = await routerContract.getRequiredBuyVolume(pair, quoteVolume);
    
    // Wallet의 CROSS 잔액 확인
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
 * 완전한 예제: 토큰 Approve 후 매수 지정가 주문
 */
async function exampleBuyLimitOrder() {
    try {
        // 1. Wallet 생성
        const { provider, wallet } = createWallet();
        
        // 2. 파라미터 설정 (예제)
        const marketAddress = GAME_MARKET;
        const baseToken = '0x...'; // BASE 토큰 주소 입력
        const quoteToken = '0x...'; // QUOTE 토큰 주소 입력
        const price = ethers.parseEther('100'); // 가격 (예: 100 CROSS)
        const amount = ethers.parseEther('1'); // 수량 (예: 1 BASE 토큰)
        
        // 3. Pair 주소 조회
        const pair = await getPair(provider, marketAddress, baseToken);
        
        // 4. QUOTE 토큰 Approve (무제한)
        await approveToken(wallet, quoteToken, BigInt(0));
        
        // 5. 매수 지정가 주문
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

// 메인 실행
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

