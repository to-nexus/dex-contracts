# V1에서 V2로의 마이그레이션 가이드

DEX 컨트랙트 V1에서 V2로의 완전한 마이그레이션 가이드입니다.

## 📋 목차

- [개요](#개요)
- [Breaking Changes](#breaking-changes)
- [마이그레이션 단계](#마이그레이션-단계)
- [업그레이드 프로세스](#업그레이드-프로세스)
- [일반적인 문제](#일반적인-문제)

## 🎯 개요

V2는 기존 마켓의 하위 호환성을 유지하면서 V1에 비해 여러 개선사항을 도입했습니다. 이 가이드는 변경사항을 이해하고 코드를 마이그레이션하는 데 도움이 됩니다.

### V2의 주요 개선사항

1. **Quote당 여러 마켓**: 동일한 quote 토큰으로 여러 마켓 생성
2. **세분화된 수수료 구조**: maker/taker 및 buyer/seller에 대한 별도 수수료
3. **컨트랙트 계정 화이트리스트**: 특정 컨트랙트가 라우터와 상호작용할 수 있도록 허용
4. **개선된 마켓 관리**: 더 나은 마켓 추적 및 발견

## ⚠️ Breaking Changes

### 1. 마켓 생성 API

**V1:**
```solidity
function createMarket(
    address _owner,
    address quote,
    address feeCollector,
    uint256 feeBps  // 단일 수수료율
) external onlyOwner returns (address);
```

**V2:**
```solidity
function createMarket(
    address _owner,
    address quote,
    address feeCollector,
    bytes memory feeData,  // 인코딩된 4개의 수수료율
    string memory message  // 추가 식별자
) external onlyOwner returns (address);
```

**변경사항:**
- `feeBps` (uint256) → `feeData` (bytes) + `message` (string)
- 수수료 데이터는 다음을 인코딩해야 합니다: `(sellerMakerFeeBps, sellerTakerFeeBps, buyerMakerFeeBps, buyerTakerFeeBps)`

### 2. 마켓 스토리지 매핑

**V1:**
```solidity
EnumerableMap.AddressToAddressMap private _allMarkets; // quote => market
function quoteToMarket(address quote) external view returns (address);
```

**V2:**
```solidity
EnumerableMap.AddressToAddressMap private _allMarkets; // market => quote
// quoteToMarket() 제거됨
function allMarkets() external view returns (address[] memory markets, address[] memory quotes);
```

**변경사항:**
- 매핑 방향 반전: `quote => market` → `market => quote`
- `quoteToMarket()` 함수 제거
- `allMarkets()`는 이제 `(quotes[], markets[])` 대신 `(markets[], quotes[])`를 반환합니다

### 3. 수수료 구조

**V1:**
```solidity
uint32 public feeBps; // 모든 거래에 대한 단일 수수료
```

**V2:**
```solidity
struct FeeConfig {
    uint32 sellerMakerFeeBps;
    uint32 sellerTakerFeeBps;
    uint32 buyerMakerFeeBps;
    uint32 buyerTakerFeeBps;
}
FeeConfig private _feeConfig;
```

**변경사항:**
- 단일 수수료율 → 4개의 별도 수수료율
- 수수료는 이제 주문 측면 및 유형별로 다를 수 있습니다

### 4. 라우터 수정자

**V1:**
```solidity
modifier checkValue() {
    _;
    if (address(this).balance != 0) revert RouterInvalidValue();
}
// 모든 컨트랙트 계정 차단
```

**V2:**
```solidity
modifier checkSubmit() {
    _checkAccountCode(_msgSender());
    _;
    if (address(this).balance != 0) revert RouterInvalidValue();
}

function _checkAccountCode(address account) private view {
    if (whitelistedCodeAccounts.contains(account)) return;
    if (account.code.length != 0) revert RouterContractAccountBlocked(account);
}
```

**변경사항:**
- `checkValue` → `checkSubmit`
- 컨트랙트 계정에 대한 화이트리스트 시스템
- `setWhitelistedCodeAccount()` 함수 추가

## 🔄 마이그레이션 단계

### 0단계: CROSS (네이티브 코인) 처리

V1과 V2 모두 크로스 체인 네이티브 코인(CROSS)에 대해 동일한 WETH 래퍼를 사용합니다. CROSS 처리를 위한 마이그레이션이 필요하지 않습니다:

- **페어 내부**: CROSS는 페어 내에서 ERC20으로 처리됩니다(두 버전 모두 동일)
- **외부 전송**: CROSS가 페어가 아닌 주소로 전송되면 자동으로 네이티브 코인으로 언래핑됩니다(두 버전 모두 동일)
- **사용자 경험**: 사용자는 `payable` 함수를 통해 네이티브 CROSS를 계속 전송할 수 있습니다(두 버전 모두 동일)

WETH 래퍼 컨트랙트 동작은 V1과 V2에서 동일하므로 기존 CROSS 통합 코드는 변경 없이 작동해야 합니다.

### 1단계: 마켓 생성 코드 업데이트

**이전 (V1):**
```solidity
uint256 feeBps = 30; // 0.3%
address market = crossDex.createMarket(owner, USDT, feeCollector, feeBps);
```

**이후 (V2):**
```solidity
// 4개의 수수료율 정의
uint32 sellerMakerFee = 25;  // 판매자 지정가 주문 0.25%
uint32 sellerTakerFee = 30;  // 판매자 시장가 주문 0.30%
uint32 buyerMakerFee = 25;    // 구매자 지정가 주문 0.25%
uint32 buyerTakerFee = 30;    // 구매자 시장가 주문 0.30%

// 수수료 데이터 인코딩
bytes memory feeData = abi.encode(sellerMakerFee, sellerTakerFee, buyerMakerFee, buyerTakerFee);

// 메시지 식별자와 함께 마켓 생성
string memory message = "Main USDT Market";
address market = crossDex.createMarket(owner, USDT, feeCollector, feeData, message);
```

### 2단계: 마켓 조회 코드 업데이트

**이전 (V1):**
```solidity
address market = crossDex.quoteToMarket(USDT);
```

**이후 (V2):**
```solidity
// 옵션 1: 모든 마켓을 가져와서 필터링
(address[] memory markets, address[] memory quotes) = crossDex.allMarkets();
address market;
for (uint i = 0; i < markets.length; i++) {
    if (quotes[i] == USDT && /* 추가 필터 */) {
        market = markets[i];
        break;
    }
}

// 옵션 2: 생성 이벤트에서 마켓 주소 저장
// 권장: MarketCreated 이벤트에서 마켓 주소 추적
```

### 3단계: 수수료 구성 코드 업데이트

**이전 (V1):**
```solidity
market.setFeeBps(30); // 단일 수수료 업데이트
```

**이후 (V2):**
```solidity
// 모든 4개의 수수료율 업데이트
market.setMarketFees(
    25,  // sellerMakerFeeBps
    30,  // sellerTakerFeeBps
    25,  // buyerMakerFeeBps
    30   // buyerTakerFeeBps
);

// 또는 현재 수수료 조회
IMarketV2.FeeConfig memory fees = market.getFeeConfig();
```

### 4단계: 라우터 상호작용 코드 업데이트

**이전 (V1):**
```solidity
// 컨트랙트 계정은 상호작용할 수 없음
// 모든 컨트랙트 호출 차단
```

**이후 (V2):**
```solidity
// 접근이 필요한 컨트랙트 계정을 화이트리스트에 추가
address[] memory contracts = new address[](1);
contracts[0] = myContract;
router.setWhitelistedCodeAccount(contracts, true);

// 이제 myContract가 라우터와 상호작용할 수 있음
```

## 🔧 업그레이드 프로세스

### 기존 배포의 경우

기존 V1 배포가 있는 경우:

1. **V2 구현 컨트랙트 배포**
   ```bash
   # 새로운 구현 배포
   forge script script/UpgradeCrossDexV2.s.sol:UpgradeCrossDexV2 --rpc-url <network>
   ```

2. **CrossDex 프록시 업그레이드**
   ```solidity
   // 프록시가 V2 구현을 가리키도록 업그레이드
   crossDex.upgradeTo(v2Implementation);
   ```

3. **필요시 재초기화**
   ```solidity
   // V2는 스토리지 마이그레이션을 위한 reinitialize() 추가
   crossDex.reinitialize(newMarketImpl, newPairImpl);
   ```

4. **기존 마켓 마이그레이션** (필요한 경우)
   - 기존 마켓은 계속 작동합니다
   - 새 마켓은 V2 API를 사용해야 합니다
   - 새로운 수수료 구조로 마이그레이션을 고려하세요

### 스토리지 마이그레이션

V2는 스토리지 마이그레이션을 위한 `reinitialize()` 함수를 포함합니다:

```solidity
function reinitialize(address _marketImpl, address _pairImpl) external onlyOwner reinitializer(2) {
    // _allMarkets를 quote=>market에서 market=>quote로 마이그레이션
    // 구현 주소 업데이트
}
```

## 🐛 일반적인 문제

### 문제 1: 마켓을 찾을 수 없음

**문제:**
```solidity
// V1 코드가 V2 컨트랙트를 사용하려고 시도
address market = crossDex.quoteToMarket(USDT); // ❌ 함수가 존재하지 않음
```

**해결책:**
```solidity
// allMarkets()를 사용하고 필터링
(address[] memory markets, address[] memory quotes) = crossDex.allMarkets();
// 일치하는 quote로 마켓 찾기
```

### 문제 2: 수수료 인코딩 오류

**문제:**
```solidity
// 잘못된 수수료 인코딩
bytes memory feeData = abi.encode(30); // ❌ 잘못됨: 단일 값
```

**해결책:**
```solidity
// 올바름: 모든 4개 값 인코딩
bytes memory feeData = abi.encode(
    uint32(25), // sellerMakerFeeBps
    uint32(30), // sellerTakerFeeBps
    uint32(25), // buyerMakerFeeBps
    uint32(30)  // buyerTakerFeeBps
);
```

### 문제 3: 컨트랙트 계정 차단

**문제:**
```solidity
// 라우터를 호출하려는 컨트랙트 (V2)
contract MyBot {
    function trade() external {
        router.submitBuyLimit(...); // ❌ Reverts: 컨트랙트 계정 차단
    }
}
```

**해결책:**
```solidity
// 먼저 컨트랙트를 화이트리스트에 추가
router.setWhitelistedCodeAccount([address(myBot)], true);
// 그 다음 거래 가능
```

### 문제 4: 스토리지 슬롯 충돌

**문제:** 적절한 마이그레이션 없이 업그레이드하면 스토리지 문제가 발생할 수 있습니다.

**해결책:**
- 항상 스토리지 마이그레이션을 위해 `reinitialize()` 사용
- 스토리지 레이아웃 호환성 확인
- 먼저 테스트넷에서 철저히 테스트

## 📚 추가 리소스

- [V1 README](./README.ko.md) - V1 개요 및 기능
- [메인 README](../README.ko.md) - V2 문서

---

**도움이 필요하신가요?** 마이그레이션 중 문제가 발생하면 참조 구현을 위해 `test/DEXV2CrossDexUpgrade.t.sol`의 테스트 파일을 검토하세요.

