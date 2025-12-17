# BuyBot 배포 가이드 (Cross Testnet)

## 네트워크 정보

| 항목 | 값 |
|------|-----|
| Network | Cross Testnet |
| RPC URL | https://testnet.crosstoken.io:22001 |

## 컨트랙트 주소

| 구분 | 주소 |
|------|------|
| DEX Router | 0xAa7B538655ce087116B1bfbb0e3C6057A405C391 |
| CROSSD/CROSS Pair | 0x8C46a3f37013cF587167d4c75d29b170cB5DAAbe |
| CROSSD Token | 0x9364ea6790f6E0EcFaa5164085f2a7de34EC55Fb |
| USDT (Cross Chain) | 0x9F85c7B5D7637E18f946cc8AF9C131318c6833d9 |
| Uniswap V3 SWAP_ROUTER | 0x61745D6Dc4342cF01bE2C145218CE1952260cd05 |
| V3_POOL (USDT/CROSSD, FEE=100) | 0x0fe6eBc8E3CD5Def9cdfC656Ec11C104e229741d |

## 배포 파라미터

| 파라미터 | 값 | 설명 |
|----------|-----|------|
| initialDelay | 0 | Admin 전송 딜레이 (초) |
| owner | 0xB777C937fa1afC99606aFa85c5b83cFe7f82BabD | 컨트랙트 소유자 |
| router | 0xAa7B538655ce087116B1bfbb0e3C6057A405C391 | CrossDex Router |
| minOrderAmount | 1000000000000000000 | 최소 주문량 (1e18 = 1 토큰) |
| interval | 0 | 거래 간격 (0 = 제한 없음) |
| recipient | 0xB777C937fa1afC99606aFa85c5b83cFe7f82BabD | BASE 토큰 수령 주소 |
| buyer | 0xB777C937fa1afC99606aFa85c5b83cFe7f82BabD | BUYER_ROLE 권한 주소 |
| manager | 0xB777C937fa1afC99606aFa85c5b83cFe7f82BabD | MANAGER_ROLE 권한 주소 |
| swapRouter | 0x61745D6Dc4342cF01bE2C145218CE1952260cd05 | Uniswap V3 SwapRouter |
| maxTickSlippage | 1 | 최대 tick 슬리피지 (1 tick = 0.01%) |

---

## 1. BuyBot 배포

```bash
cd /Users/luis/go/src/github.com/to-nexus/dex-contracts

forge script script/BuyBot.s.sol:BuyBotScript \
  --sig "deployBuyBot(uint48,address,address,uint256,uint256,address,address,address,address,uint24)" \
  0 \
  0xB777C937fa1afC99606aFa85c5b83cFe7f82BabD \
  0xAa7B538655ce087116B1bfbb0e3C6057A405C391 \
  1000000000000000000 \
  0 \
  0xB777C937fa1afC99606aFa85c5b83cFe7f82BabD \
  0xB777C937fa1afC99606aFa85c5b83cFe7f82BabD \
  0xB777C937fa1afC99606aFa85c5b83cFe7f82BabD \
  0x61745D6Dc4342cF01bE2C145218CE1952260cd05 \
  1 \
  --rpc-url https://testnet.crosstoken.io:22001 \
  --private-key YOUR_PRIVATE_KEY \
  --priority-gas-price 1gwei \
  --skip-simulation \
  --broadcast
```

> **주의**: `YOUR_PRIVATE_KEY`를 실제 private key로 교체하세요.

배포 성공 시 출력에서 **BuyBot 컨트랙트 주소**를 확인하세요.

**배포된 BuyBot 주소**: `0xdb66f3c9708AA74133Fb78dcDE252f585D7713a8`

---

## 2. SwapToken 설정 (USDT)

배포된 BuyBot 주소를 `BUYBOT_ADDRESS`에 입력하세요.

```bash
forge script script/BuyBot.s.sol:BuyBotScript \
  --sig "setSwapToken(address,address)" \
  BUYBOT_ADDRESS \
  0x9F85c7B5D7637E18f946cc8AF9C131318c6833d9 \
  --rpc-url https://testnet.crosstoken.io:22001 \
  --private-key YOUR_PRIVATE_KEY \
  --priority-gas-price 1gwei \
  --skip-simulation \
  --broadcast
```

---

## 3. SwapPool 설정 (USDT → CROSSD)

```bash
forge script script/BuyBot.s.sol:BuyBotScript \
  --sig "setSwapPool(address,address,address,address)" \
  BUYBOT_ADDRESS \
  0x9F85c7B5D7637E18f946cc8AF9C131318c6833d9 \
  0x9364ea6790f6E0EcFaa5164085f2a7de34EC55Fb \
  0x0fe6eBc8E3CD5Def9cdfC656Ec11C104e229741d \
  --rpc-url https://testnet.crosstoken.io:22001 \
  --private-key YOUR_PRIVATE_KEY \
  --priority-gas-price 1gwei \
  --skip-simulation \
  --broadcast
```

---

## 4. 배포 확인

```bash
# BuyBot 정보 출력
forge script script/BuyBot.s.sol:BuyBotScript \
  --sig "printBuyBotInfo(address)" \
  BUYBOT_ADDRESS \
  --rpc-url https://testnet.crosstoken.io:22001

# 특정 pair에 대한 buyMarket 가능 여부 확인
forge script script/BuyBot.s.sol:BuyBotScript \
  --sig "checkCanBuyMarket(address,address,address)" \
  BUYBOT_ADDRESS \
  0x8C46a3f37013cF587167d4c75d29b170cB5DAAbe \
  0xB777C937fa1afC99606aFa85c5b83cFe7f82BabD \
  --rpc-url https://testnet.crosstoken.io:22001
```

---

## 5. 사용 예시

### 5.1 USDT를 CROSSD로 스왑

BuyBot에 USDT를 입금한 후 실행:

**방법 1: BuyBot.s.sol 사용**

```bash
forge script script/BuyBot.s.sol:BuyBotScript \
  --sig "swapToQuote(address,address,uint24)" \
  BUYBOT_ADDRESS \
  0x8C46a3f37013cF587167d4c75d29b170cB5DAAbe \
  100 \
  --rpc-url https://testnet.crosstoken.io:22001 \
  --private-key YOUR_PRIVATE_KEY \
  --priority-gas-price 1gwei \
  --skip-simulation \
  --broadcast
```

**방법 2: 전용 스크립트 사용 (BuyBot.SwapToQuote.s.sol)**

```bash
forge script script/BuyBot.SwapToQuote.s.sol:SwapToQuoteScript \
  --sig "swapToQuote(address,address,uint24)" \
  BUYBOT_ADDRESS \
  0x8C46a3f37013cF587167d4c75d29b170cB5DAAbe \
  100 \
  --rpc-url https://testnet.crosstoken.io:22001 \
  --private-key YOUR_PRIVATE_KEY \
  --priority-gas-price 1gwei \
  --skip-simulation \
  --broadcast
```

> **참고**: `100`은 Uniswap V3 fee tier (0.01%)

### 5.2 시장가 매수 실행

BuyBot에 CROSSD(quote token)가 있을 때:

```bash
forge script script/BuyBot.s.sol:BuyBotScript \
  --sig "buyMarket(address,address,uint256,uint256)" \
  BUYBOT_ADDRESS \
  0x8C46a3f37013cF587167d4c75d29b170cB5DAAbe \
  1000000000000000000 \
  0 \
  --rpc-url https://testnet.crosstoken.io:22001 \
  --private-key YOUR_PRIVATE_KEY \
  --priority-gas-price 1gwei \
  --skip-simulation \
  --broadcast
```

---

## 6. DEX 화이트리스팅

BuyBot이 DEX에서 거래하려면 화이트리스팅이 필요합니다.

**대상**: 배포된 BuyBot 컨트랙트 주소

---

## 체크리스트

- [ ] BuyBot 배포 완료
- [ ] setSwapToken (USDT) 설정 완료
- [ ] setSwapPool (USDT → CROSSD) 설정 완료
- [ ] DEX 화이트리스팅 완료
- [ ] BuyBot에 USDT 또는 CROSSD 입금
- [ ] 테스트 트랜잭션 실행

