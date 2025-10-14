# 🎯 Issue: Buyer 수수료 지불 시스템 구현

## 📋 Overview

현재 DEX 시스템에서는 Seller만 수수료를 지불하는 구조입니다. 이를 개선하여 Buyer도 수수료를 지불하도록 하여 보다 공정하고 균형잡힌 수수료 구조를 구현하고자 합니다.

## 🔄 Current State vs Target State

### 📌 Current State (현재 상태)
```
🔸 Sell Order:
  - Seller: BASE 토큰 제공 → QUOTE 토큰 수령 (수수료 차감됨)
  - 수수료: QUOTE 토큰에서 차감

🔸 Buy Order:  
  - Buyer: QUOTE 토큰 제공 → BASE 토큰 수령
  - 수수료: 없음 ❌
```

### 🎯 Target State (목표 상태)
```
🔸 Sell Order:
  - Seller: BASE 토큰 제공 → QUOTE 토큰 수령 (수수료 차감됨)
  - 수수료: QUOTE 토큰에서 차감 (기존과 동일)

🔸 Buy Order:
  - Buyer: QUOTE 토큰 제공 (주문금액 + 수수료) → BASE 토큰 수령
  - 수수료: QUOTE 토큰으로 별도 지불 ✅
```

## 🛠️ Implementation Requirements

### 1. **Router Layer Changes**

#### 1.1 Buy Order 수수료 계산 로직 추가
```solidity
// Before: 기존 로직
function submitBuyLimit(address pair, uint256 price, uint256 amount, ...) external {
    uint256 requiredQuote = price * amount / denominator;
    QUOTE.transferFrom(msg.sender, pair, requiredQuote);
}

// After: 수정된 로직
function submitBuyLimit(address pair, uint256 price, uint256 amount, ...) external {
    uint256 orderValue = price * amount / denominator;
    uint32 buyerFeeRate = _getBuyerFeeRate(pair, isTaker);
    uint256 feeAmount = orderValue * buyerFeeRate / BPS_DENOMINATOR;
    uint256 totalRequired = orderValue + feeAmount;
    
    QUOTE.transferFrom(msg.sender, pair, totalRequired);
}
```

#### 1.2 Market Order 수수료 계산
```solidity
function submitBuyMarket(address pair, uint256 maxQuoteAmount, ...) external {
    // 예상 매칭 금액 기준으로 수수료 계산
    uint256 estimatedFee = _calculateEstimatedBuyerFee(pair, maxQuoteAmount);
    // 최대 한도에 수수료 포함
}
```

### 2. **Pair Layer Changes**

#### 2.1 Buy Order 처리 시 수수료 분리 (단순화)
```solidity
function _executeBuyOrder(...) internal {
    // 1. 전달받은 QUOTE에서 수수료 분리
    uint256 feeAmount = _calculateBuyerFee(orderValue, feeBps);
    uint256 actualOrderValue = totalQuote - feeAmount;
    
    // 2. 수수료는 Pair 컨트랙트에서 보관 (매칭 시까지)
    // 매칭이 이루어질 때 _collectFee()에서 feeCollector에게 전송
    
    // 3. Order 구조체에 수수료율 저장 (취소 시 환불 계산용)
    _allOrders[orderId].feeBps = feeBps;
    
    // 4. 실제 주문 금액으로 매칭 진행
    // ... existing matching logic
}
```

#### 2.2 Cancel Order 시 수수료 환불 (단순화)
```solidity
function cancelOrder(uint256 orderId) external {
    Order storage order = _allOrders[orderId];
    require(order.owner == msg.sender, "Not owner");
    
    // Buy Order 취소 시 남은 주문금액 + 해당 수량에 대한 수수료 환불
    if (order.side == OrderSide.BUY) {
        uint256 remainingValue = Math.mulDiv(order.price, order.amount, DENOMINATOR);
        uint256 refundableFee = Math.mulDiv(remainingValue, order.feeBps, BPS_DENOMINATOR);
        uint256 totalRefund = remainingValue + refundableFee;
        
        // 주문금액 + 수수료를 한 번에 환불 (별도 이벤트 없음)
        QUOTE.safeTransfer(msg.sender, totalRefund);
    }
    
    // ... existing cancel logic
}
```

#### 2.3 수수료 이벤트 (기존 이벤트 유지)
```solidity
// 기존 FeeCollect 이벤트는 수정하지 않음
// 선지불 수수료 추적/환불에 대한 별도 이벤트는 발생시키지 않음
// 필요시 기존 FeeCollect 이벤트만 활용 (Seller 수수료 징수 시에만)
```

### 3. **Fee Calculation Logic**

#### 3.1 확장된 수수료율 결정 함수
```solidity
// Buyer 수수료율 조회
function _getBuyerFeeRate(address pair, bool isTaker) internal view returns (uint32) {
    if (isTaker) {
        return IPairV2(pair).getEffectiveFees().buyerTaker;
    } else {
        return IPairV2(pair).getEffectiveFees().buyerMaker;
    }
}

// Seller 수수료율 조회 (기존 로직 확장)
function _getSellerFeeRate(address pair, bool isTaker) internal view returns (uint32) {
    if (isTaker) {
        return IPairV2(pair).getEffectiveFees().sellerTaker;
    } else {
        return IPairV2(pair).getEffectiveFees().sellerMaker;
    }
}
```

#### 3.2 MarketImplV2 수수료 설정 함수 확장
```solidity
function setMarketFees(
    uint32 _sellerMakerFeeBps,
    uint32 _sellerTakerFeeBps,
    uint32 _buyerMakerFeeBps,  // 신규
    uint32 _buyerTakerFeeBps   // 신규
) external onlyOwner {
    // 유효성 검증
    if (_sellerMakerFeeBps != NO_FEE_BPS && _sellerMakerFeeBps >= BPS_DENOMINATOR) revert MarketInvalidInitializeData("sellerMakerFeeBPS");
    if (_sellerTakerFeeBps != NO_FEE_BPS && _sellerTakerFeeBps >= BPS_DENOMINATOR) revert MarketInvalidInitializeData("sellerTakerFeeBPS");
    if (_buyerMakerFeeBps != NO_FEE_BPS && _buyerMakerFeeBps >= BPS_DENOMINATOR) revert MarketInvalidInitializeData("buyerMakerFeeBPS");
    if (_buyerTakerFeeBps != NO_FEE_BPS && _buyerTakerFeeBps >= BPS_DENOMINATOR) revert MarketInvalidInitializeData("buyerTakerFeeBPS");

    emit MarketFeesUpdated(_sellerMakerFeeBps, _sellerTakerFeeBps, _buyerMakerFeeBps, _buyerTakerFeeBps);
    
    sellerMakerFeeBps = _sellerMakerFeeBps;
    sellerTakerFeeBps = _sellerTakerFeeBps;
    buyerMakerFeeBps = _buyerMakerFeeBps;
    buyerTakerFeeBps = _buyerTakerFeeBps;
}
```

#### 3.3 PairImplV2 수수료 설정 함수 확장
```solidity
function setPairFees(
    uint32 _sellerMakerFeeBps,
    uint32 _sellerTakerFeeBps,
    uint32 _buyerMakerFeeBps,  // 신규
    uint32 _buyerTakerFeeBps   // 신규
) external onlyOwner {
    // 유효성 검증 및 설정
    feeConfig = FeeConfig({
        sellerMakerFeeBps: _sellerMakerFeeBps,
        sellerTakerFeeBps: _sellerTakerFeeBps,
        buyerMakerFeeBps: _buyerMakerFeeBps,
        buyerTakerFeeBps: _buyerTakerFeeBps
    });
    
    emit PairFeesUpdated(_sellerMakerFeeBps, _sellerTakerFeeBps, _buyerMakerFeeBps, _buyerTakerFeeBps);
}
```

## 📊 Technical Specifications

### 0. **상수 정의**
```solidity
// IMarket.sol에서 정의된 상수들 활용
uint32 constant NO_FEE_BPS = type(uint32).max;  // "마켓 수수료 사용" 특수값
uint32 constant BPS_DENOMINATOR = 10000;        // 베이시스 포인트 분모 (100%)

// 수수료 유효성 검증
if (_feeBps != NO_FEE_BPS && _feeBps >= BPS_DENOMINATOR) revert InvalidFee();
```

### 1. **수수료 데이터 구조**

#### 1.1 MarketImplV2 수수료 저장소 확장
```solidity
contract MarketImplV2 {
    // 기존 필드들...
    uint32 public override sellerMakerFeeBps;  // Seller Maker 수수료
    uint32 public override sellerTakerFeeBps;  // Seller Taker 수수료
    uint32 public override buyerMakerFeeBps;   // Buyer Maker 수수료 (신규)
    uint32 public override buyerTakerFeeBps;   // Buyer Taker 수수료 (신규)
}
```

#### 1.2 PairImplV2 수수료 구조 확장
```solidity
struct FeeConfig {
    uint32 sellerMakerFeeBps;  // Seller Maker 수수료
    uint32 sellerTakerFeeBps;  // Seller Taker 수수료
    uint32 buyerMakerFeeBps;   // Buyer Maker 수수료 (신규)
    uint32 buyerTakerFeeBps;   // Buyer Taker 수수료 (신규)
}
```

#### 1.3 Buy Order 수수료 처리 (단순화)
```solidity
// buyOrderPrepaidFees 매핑 불필요!
// Cancel 시 남은 주문 수량에 대해 수수료를 다시 계산해서 환불
// Order 구조체의 기존 feeBps 필드 활용하여 수수료율 저장
```

#### 1.4 초기화 함수 수수료 구조 개선
```solidity
// Before: 기존 2개 수수료 매개변수
function initialize(
    address router,
    address quote,
    address base,
    uint256 _tickSize,
    uint256 _lotSize,
    uint32 _makerFeeBps,
    uint32 _takerFeeBps
) external initializer { ... }

// After: bytes로 수수료 데이터 전달 (4개 수수료)
function initialize(
    address router,
    address quote,
    address base,
    uint256 _tickSize,
    uint256 _lotSize,
    bytes memory feeData  // 4개 수수료를 인코딩한 데이터
) external initializer {
    // 수수료 데이터 디코딩
    (
        uint32 _sellerMakerFeeBps,
        uint32 _sellerTakerFeeBps,
        uint32 _buyerMakerFeeBps,
        uint32 _buyerTakerFeeBps
    ) = abi.decode(feeData, (uint32, uint32, uint32, uint32));
    
    // 기존 초기화 로직...
    // 4개 수수료 설정
}

// 호출 시 수수료 데이터 인코딩 예시
bytes memory feeData = abi.encode(
    uint32(20),  // sellerMakerFeeBps
    uint32(30),  // sellerTakerFeeBps
    uint32(15),  // buyerMakerFeeBps
    uint32(25)   // buyerTakerFeeBps
);
```

### 2. **수수료 지불 방식**
- **통화**: 모든 수수료는 QUOTE 토큰으로만 지불
- **계산 시점**: Router에서 주문 제출 시
- **수취자**: MarketImpl의 feeCollector 주소
- **이벤트**: 기존 `FeeCollect` 이벤트는 수정하지 않음
- **환불**: Buy Order 취소 시 주문금액과 수수료를 합쳐서 한 번에 환불

### 3. **수수료율 적용 매트릭스**

|        | Maker Fee            | Taker Fee            |
|--------|----------------------|----------------------|
| Seller | `sellerMakerFeeBps`  | `sellerTakerFeeBps`  |
| Buyer  | `buyerMakerFeeBps`   | `buyerTakerFeeBps`   |

### 4. **가스 최적화**
```solidity
// 한 번의 transferFrom으로 주문금액 + 수수료 전송
QUOTE.transferFrom(buyer, pair, orderValue + feeAmount);

// Pair에서 수수료와 주문금액 분리 처리
```

## 🔄 Migration Strategy

### Phase 1: 데이터 구조 확장
- [ ] MarketImplV2: 4개 수수료 필드 추가
- [ ] PairImplV2: FeeConfig 구조체 확장 및 initialize 함수 수정
- [ ] Order 구조체의 기존 feeBps 필드 활용
- [ ] initialize 함수에서 bytes memory feeData → abi.decode 로직 구현

### Phase 2: Router 수정
- [ ] Buy order 수수료 계산 로직 추가
- [ ] Market order 수수료 예상 계산
- [ ] 4가지 수수료율 조회 로직

### Phase 3: Pair 수정  
- [ ] Buy order 처리 시 수수료 분리 및 즉시 징수
- [ ] Order.feeBps 필드에 수수료율 저장
- [ ] Cancel 시 남은 수량 기준 수수료 계산 후 환불
- [ ] 기존 FeeCollect 이벤트 활용

### Phase 4: 통합 테스트
- [ ] 4가지 수수료 조합 테스트
- [ ] Buy Order 취소 시 수수료 환불 테스트
- [ ] End-to-end 시나리오 테스트
- [ ] 가스 비용 최적화
- [ ] 보안 검증

## 🧪 Test Cases

### 1. **기본 시나리오**
```solidity
function test_buyer_fee_limit_order() external {
    // 1. Buyer가 limit order 제출 (주문금액 + Buyer Maker/Taker 수수료)
    // 2. 4가지 수수료율이 올바르게 계산되는지 확인
    // 3. Buyer/Seller 수수료가 모두 feeCollector에게 전송되는지 확인
    // 4. 기존 FeeCollect 이벤트가 각각 발생하는지 확인
}

function test_buyer_fee_market_order() external {
    // Market order에서 Buyer Taker 수수료 계산 테스트
}

function test_seller_fee_compatibility() external {
    // 기존 Seller 수수료 로직이 그대로 작동하는지 확인
}
```

### 2. **수수료 설정 테스트**
```solidity
function test_market_fees_four_parameters() external {
    // MarketImplV2에서 4가지 수수료 설정 테스트
    marketV2.setMarketFees(
        20,  // sellerMakerFeeBps
        30,  // sellerTakerFeeBps  
        15,  // buyerMakerFeeBps
        25   // buyerTakerFeeBps
    );
}

function test_pair_fees_four_parameters() external {
    // PairImplV2에서 4가지 수수료 설정 테스트
    pairV2.setPairFees(
        10,  // sellerMakerFeeBps
        20,  // sellerTakerFeeBps
        8,   // buyerMakerFeeBps
        18   // buyerTakerFeeBps
    );
}

function test_initialize_with_fee_data_bytes() external {
    // bytes로 인코딩된 수수료 데이터로 초기화 테스트
    bytes memory feeData = abi.encode(
        uint32(25),  // sellerMakerFeeBps
        uint32(35),  // sellerTakerFeeBps
        uint32(20),  // buyerMakerFeeBps
        uint32(30)   // buyerTakerFeeBps
    );
    
    // PairImplV2 초기화
    pairV2.initialize(
        router,
        quoteToken,
        baseToken,
        tickSize,
        lotSize,
        feeData
    );
    
    // 수수료가 올바르게 설정되었는지 확인
    FeeConfig memory fees = pairV2.getEffectiveFees();
    assertEq(fees.sellerMakerFeeBps, 25);
    assertEq(fees.sellerTakerFeeBps, 35);
    assertEq(fees.buyerMakerFeeBps, 20);
    assertEq(fees.buyerTakerFeeBps, 30);
}

function test_initialize_invalid_fee_data() external {
    // 잘못된 수수료 데이터로 초기화 시 실패 테스트
    bytes memory invalidFeeData = abi.encode(uint32(15000)); // 단일 값
    
    vm.expectRevert();
    pairV2.initialize(router, quoteToken, baseToken, tickSize, lotSize, invalidFeeData);
}
```

### 3. **Buy Order 수수료 관리 테스트 (단순화)**
```solidity
function test_buy_order_fee_collection() external {
    // 1. Buy Order 등록 시 수수료 즉시 징수 확인
    // 2. Order.feeBps 필드에 수수료율 저장 확인
}

function test_buy_order_cancel_fee_refund() external {
    // 1. Buy Order 취소 시 남은 수량 기준 수수료 계산
    // 2. 주문금액 + 계산된 수수료를 합쳐서 환불 확인
    // 3. 부분 체결 후 취소 시 남은 수량에 대해서만 환불되는지 검증
}

function test_buy_order_full_vs_partial_cancel() external {
    // 1. 완전 미체결 주문 취소: 전체 수량에 대한 수수료 환불
    // 2. 부분 체결 후 취소: 남은 수량에 대한 수수료만 환불
}
```

### 4. **엣지 케이스**
```solidity
function test_buyer_insufficient_balance_with_fee() external {
    // Buyer가 수수료를 포함한 잔액이 부족한 경우
}

function test_zero_buyer_fees() external {
    // Buyer 수수료가 0인 경우 (선지불/환불 로직 없음)
}

function test_buy_order_complete_fill_no_refund() external {
    // Buy Order가 완전 체결된 경우 환불할 수수료 없음 확인
}
```

### 5. **Maker/Taker × Buyer/Seller 매트릭스 테스트**
```solidity
function test_seller_maker_vs_seller_taker_fee() external {
    // Seller의 Maker vs Taker 수수료 차이 테스트
}

function test_buyer_maker_vs_buyer_taker_fee() external {
    // Buyer의 Maker vs Taker 수수료 차이 테스트
}

function test_all_four_fee_combinations() external {
    // 4가지 수수료 조합 모두 테스트
    // (SellerMaker, SellerTaker, BuyerMaker, BuyerTaker)
}
```

## ⚠️ Potential Issues & Considerations

### 1. **사용자 경험**
- **영향**: Buyer가 더 많은 QUOTE 토큰을 준비해야 함
- **대응**: 프론트엔드에서 4가지 수수료율과 필요 금액 명확히 표시

### 2. **데이터 구조 확장 (단순화)**
- **Storage 증가**: MarketImpl과 PairImpl 모두 4개 수수료 필드 저장
- **추가 Storage 없음**: `buyOrderPrepaidFees` 매핑 불필요, 기존 Order.feeBps 활용
- **가스 비용**: 설정/조회 함수만 증가, 복잡한 수수료 추적 로직 없음
- **이벤트**: 기존 FeeCollect 이벤트는 수정하지 않음, 별도 이벤트 없음

### 3. **호환성**
- **기존 컨트랙트**: 2개 → 4개 수수료 매개변수로 인터페이스 변경
- **initialize 함수**: uint32 매개변수 → bytes memory feeData로 변경
- **마이그레이션**: 기존 수수료를 Seller 수수료로, Buyer 수수료는 별도 설정
- **Factory 컨트랙트**: 새로운 initialize 시그니처에 맞게 호출 로직 수정 필요

### 4. **수수료 환불 단순화**
- **간단한 계산**: 취소 시 남은 수량에 대해 수수료 재계산하여 환불
- **Storage 효율성**: 추가 매핑 없이 기존 Order 구조체 활용
- **정확성**: 부분 체결 후에도 남은 수량에 정확한 수수료 적용

### 5. **보안**
- **수수료 계산 오류**: Overflow/underflow 방지
- **Re-entrancy**: 수수료 전송/환불 시 보안 고려
- **권한 검증**: Cancel 시 주문 소유자 검증 강화

## 🎯 Expected Benefits

### 1. **공정성 및 유연성 향상**
- Buyer와 Seller 모두 수수료 부담으로 공정성 증대
- 4가지 수수료율 조합으로 최대 유연성 제공
- 역할별(Maker/Taker) × 거래방향별(Buy/Sell) 세밀한 인센티브 설계

### 2. **수익 증대**
- 모든 거래에서 Buyer/Seller 양방향 수수료 수취
- 프로토콜 수익성 대폭 향상

### 3. **시장 효율성 개선**
- 세분화된 수수료 구조로 유동성 공급 최적화
- Maker 인센티브 강화로 시장 깊이 증가
- Buyer/Seller 균형잡힌 참여 유도

---

**Priority**: High  
**Complexity**: Very High (4가지 수수료 + 선지불 수수료 추적/환불 시스템)  
**Estimated Timeline**: 4-5 weeks  
**Dependencies**: Taker/Maker 차등 수수료 시스템 (완료)  
**Breaking Changes**: 인터페이스 변경 (2개 → 4개 수수료 매개변수)  
**Critical Feature**: Buy Order 선지불 수수료 환불 시스템 (업그레이드 호환성 고려)