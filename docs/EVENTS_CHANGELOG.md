# Events Changelog: V2 → V3

This document details all event changes between DEX V2 and V3 contracts. Use this as a reference when migrating indexers, subgraphs, or off-chain event listeners.

---

## 📋 Table of Contents

- [Summary](#summary)
- [CrossDexImpl](#crossdeximpl-v2--v3)
- [CrossDexRouter](#crossdexrouter-v2--v3)
- [MarketImpl](#marketimpl-v2--v3)
- [PairImpl](#pairimpl-v2--v3)
- [FeeController (V3 New)](#feecontroller-v3-new-contracts)

---

## Summary

| Contract | Added | Modified | Removed |
|----------|-------|----------|---------|
| CrossDexImpl | 1 | 1 | 0 |
| CrossDexRouter | 1 | 0 | 0 |
| MarketImpl | 1 | 0 | 1 |
| PairImpl | 1 | 1 | 1 |
| FeeController | 2 (new) | - | - |

---

## CrossDexImpl (V2 → V3)

### Unchanged Events

| Event | Signature | Description |
|-------|-----------|-------------|
| `TickSizeSetterSet` | `(address indexed before, address indexed current)` | Emitted when tick size setter address changes |
| `PairImplSet` | `(address indexed before, address indexed current)` | Emitted when pair implementation address changes |
| `MarketImplSet` | `(address indexed before, address indexed current)` | Emitted when market implementation address changes |

### Modified Events

#### `MarketCreated`

**V2:**
```solidity
event MarketCreated(
    address indexed quote,
    address indexed market,
    address indexed owner,
    address fee_collector,  // snake_case
    string message
);
```

**V3:**
```solidity
event MarketCreated(
    address indexed quote,
    address indexed market,
    address indexed owner,
    address feeCollector,   // camelCase
    string message
);
```

| Change | Detail |
|--------|--------|
| Parameter rename | `fee_collector` → `feeCollector` (naming convention) |
| Semantic | No change - still represents the fee recipient address |

### New Events (V3)

#### `FeeControllerAllowed`

```solidity
event FeeControllerAllowed(address indexed feeController, bool indexed allowed);
```

| Field | Type | Indexed | Description |
|-------|------|---------|-------------|
| `feeController` | `address` | ✅ | FeeController contract address |
| `allowed` | `bool` | ✅ | Whether the controller is allowed |

**Purpose:** Emitted when a FeeController implementation is added to or removed from the allowlist. Only allowed FeeControllers can be used by Markets and Pairs.

**When emitted:** `CrossDexImplV3.setFeeControllerAllow()`

---

## CrossDexRouter (V2 → V3)

### Unchanged Events

| Event | Signature | Description |
|-------|-----------|-------------|
| `FindPrevPriceCountChanged` | `(uint256 indexed before, uint256 indexed current)` | Max iterations for price discovery |
| `MaxMatchCountChanged` | `(uint256 indexed before, uint256 indexed current)` | Max matches per order submission |
| `CancelLimitChanged` | `(uint256 indexed before, uint256 indexed current)` | Max orders per cancel call |
| `WhitelistedCodeAccountSet` | `(address indexed account, bool whitelisted)` | Contract account whitelist change |

### New Events (V3)

#### `Skim`

```solidity
event Skim(address indexed to, uint256 amount);
```

| Field | Type | Indexed | Description |
|-------|------|---------|-------------|
| `to` | `address` | ✅ | Recipient of recovered ETH |
| `amount` | `uint256` | ❌ | Amount of ETH transferred |

**Purpose:** Emitted when the owner recovers ETH that was forcibly sent to the Router (e.g., via `selfdestruct`). This prevents DoS attacks where an attacker force-sends ETH to brick the router's balance checks.

**When emitted:** `CrossDexRouterV3.skim()`

---

## MarketImpl (V2 → V3)

### Unchanged Events

| Event | Signature | Description |
|-------|-----------|-------------|
| `PairCreated` | `(address indexed pair, address indexed base, uint256 timestamp)` | New trading pair deployed |
| `PairImplSet` | `(address indexed before, address indexed current)` | Pair implementation address change |

### Removed Events (V3)

#### `FeeCollectorChanged` ❌

**V2:**
```solidity
event FeeCollectorChanged(address indexed before, address indexed current);
```

**Reason for removal:** In V3, fee collection is delegated to FeeController contracts. The `feeCollector` address is now part of FeeController's configuration, not Market's.

**Replacement:** Use `FeeControllerUpdated` to track when the Market's FeeController changes.

#### `MarketFeesUpdated` ❌

**V2:**
```solidity
event MarketFeesUpdated(
    uint32 sellerMakerFee,
    uint32 sellerTakerFee,
    uint32 buyerMakerFee,
    uint32 buyerTakerFee
);
```

**Reason for removal:** Fee configuration is now managed by FeeController. The event definition exists in V3 code but is never emitted.

**Replacement:** Listen to FeeController initialization events or query FeeController state directly.

### New Events (V3)

#### `FeeControllerUpdated`

```solidity
event FeeControllerUpdated(address indexed before, address indexed current);
```

| Field | Type | Indexed | Description |
|-------|------|---------|-------------|
| `before` | `address` | ✅ | Previous FeeController address |
| `current` | `address` | ✅ | New FeeController address |

**Purpose:** Emitted when the Market's default FeeController is changed. Pairs created after this change will use the new FeeController.

**When emitted:** `MarketImplV3.setFeeController()`

---

## PairImpl (V2 → V3)

### Unchanged Events

| Event | Signature | Description |
|-------|-----------|-------------|
| `OrderCreated` | `(address indexed owner, uint256 indexed orderId, OrderSide indexed side, uint256 price, uint256 amount, uint256 timestamp)` | New order placed |
| `OrderMatched` | `(uint256 indexed sellId, uint256 indexed buyId, uint256 indexed price, uint256 amount, uint256 timestamp)` | Two orders matched |
| `OrderClosed` | `(uint256 indexed orderId, CloseType indexed closeType, uint256 timestamp)` | Order fully filled/cancelled |
| `TickSizeUpdated` | `(uint256 beforeLotSize, uint256 newLotSize, uint256 beforeTickSize, uint256 newTickSize)` | Tick/lot size changed |
| `Skim` | `(address indexed caller, address indexed erc20, address indexed to, uint256 amount)` | Excess tokens recovered |

### Modified Events

#### `FeeCollect`

**V2:**
```solidity
event FeeCollect(
    uint256 indexed orderId,
    address indexed owner,
    uint256 amount,
    address indexed recipient,  // ❌ Removed in V3
    uint256 feeBps,             // ❌ Removed in V3
    uint256 fee,
    uint256 value
);
```

**V3:**
```solidity
event FeeCollect(
    uint256 indexed orderId,
    address indexed owner,
    uint256 amount,
    uint256 fee,
    uint256 value
);
```

| Change | Detail |
|--------|--------|
| Removed `recipient` | Fee recipient is now determined by FeeController |
| Removed `feeBps` | Fee rate is encapsulated in FeeController |
| Remaining fields | Same meaning: `amount` (gross), `fee` (deducted), `value` (net received) |

**Migration note:** If you need `recipient` or `feeBps`, listen to `FeeControllerFeesSettled` from the FeeController instead.

### Removed Events (V3)

#### `PairFeesUpdated` ❌

**V2:**
```solidity
event PairFeesUpdated(
    uint32 sellerMakerFee,
    uint32 sellerTakerFee,
    uint32 buyerMakerFee,
    uint32 buyerTakerFee
);
```

**Reason for removal:** Pair-level fee configuration is now managed through FeeController. Use `FeeControllerUpdated` to detect when a Pair's fee logic changes.

### New Events (V3)

#### `FeeControllerUpdated`

```solidity
event FeeControllerUpdated(address indexed before, address indexed current);
```

| Field | Type | Indexed | Description |
|-------|------|---------|-------------|
| `before` | `address` | ✅ | Previous FeeController address |
| `current` | `address` | ✅ | New FeeController address |

**Purpose:** Emitted when a Pair's FeeController is changed. The new controller's `initialize()` is called immediately after.

**When emitted:** `PairImplV3.setFeeController()`

---

## FeeController (V3 New Contracts)

V3 introduces a pluggable fee system via FeeController contracts. These are called via `delegatecall` from PairImplV3 and emit their own events.

### FeeControllerV2Compat

Provides V2-compatible maker/taker fee structure (4 separate fee rates).

#### `FeeControllerFeesSettled`

```solidity
event FeeControllerFeesSettled(
    uint256 indexed takerId,
    address indexed feeCollector,
    uint256 totalFee
);
```

| Field | Type | Indexed | Description |
|-------|------|---------|-------------|
| `takerId` | `uint256` | ✅ | Order ID of the taker that triggered settlement |
| `feeCollector` | `address` | ✅ | Address receiving the accumulated fees |
| `totalFee` | `uint256` | ❌ | Sum of all maker + taker fees in this match batch |

**When emitted:** After all matches in a single `submitLimitOrder()` or `submitMarketOrder()` call are processed.

**Note:** This event is emitted from the Pair's address (due to `delegatecall`), not from the FeeController's address.

---

### FeeControllerV3Split

Provides taker-only fee with 3-way split (creator / maker rebate / system).

#### `FeeControllerV3FeesSettled`

```solidity
event FeeControllerV3FeesSettled(
    uint256 indexed takerId,
    uint256 totalTakerFee,
    uint256 creatorFee,
    uint256 feeCollectorFee,
    uint256 makerRebatePaidTotal,
    address indexed creator,
    address indexed feeCollector
);
```

| Field | Type | Indexed | Description |
|-------|------|---------|-------------|
| `takerId` | `uint256` | ✅ | Order ID of the taker |
| `totalTakerFee` | `uint256` | ❌ | Total fee collected from taker |
| `creatorFee` | `uint256` | ❌ | Portion sent to creator |
| `feeCollectorFee` | `uint256` | ❌ | Portion sent to system fee collector |
| `makerRebatePaidTotal` | `uint256` | ❌ | Sum of rebates paid to makers (paid during matching, not at settlement) |
| `creator` | `address` | ✅ | Creator address |
| `feeCollector` | `address` | ✅ | System fee collector address |

**Fee split formula:**
```
totalTakerFee = creatorFee + makerRebatePaidTotal + feeCollectorFee
```

**Note:** Maker rebates are transferred immediately during each match (not batched), but the total is reported in this event for reconciliation.

---

## 🔄 Migration Checklist for Indexers

### Event Topic Changes

| Event | V2 Topic0 | V3 Topic0 | Notes |
|-------|-----------|-----------|-------|
| `FeeCollect` | Different | Different | Signature changed (fewer params) |
| Others | Same | Same | Topic0 unchanged |

### Recommended Actions

1. **Update `FeeCollect` parser**
   - Remove `recipient` and `feeBps` field parsing
   - Add handler for `FeeControllerFeesSettled` if fee details needed

2. **Add new event handlers**
   - `FeeControllerAllowed` (CrossDex)
   - `Skim` (Router)
   - `FeeControllerUpdated` (Market, Pair)
   - `FeeControllerFeesSettled` / `FeeControllerV3FeesSettled`

3. **Remove deprecated handlers**
   - `FeeCollectorChanged` (Market)
   - `PairFeesUpdated` (Pair)
   - `MarketFeesUpdated` (Market) - was already rarely used

4. **Note delegatecall behavior**
   - FeeController events are emitted from **Pair address**, not FeeController address
   - Filter by `log.address == pairAddress` when indexing FeeController events

---

## 📊 Event Comparison Table

| Contract | Event Name | V2 | V3 | Change Type |
|----------|------------|----|----|-------------|
| CrossDexImpl | `MarketCreated` | ✅ | ✅ | Modified (param rename) |
| CrossDexImpl | `TickSizeSetterSet` | ✅ | ✅ | Unchanged |
| CrossDexImpl | `PairImplSet` | ✅ | ✅ | Unchanged |
| CrossDexImpl | `MarketImplSet` | ✅ | ✅ | Unchanged |
| CrossDexImpl | `FeeControllerAllowed` | ❌ | ✅ | **New** |
| Router | `FindPrevPriceCountChanged` | ✅ | ✅ | Unchanged |
| Router | `MaxMatchCountChanged` | ✅ | ✅ | Unchanged |
| Router | `CancelLimitChanged` | ✅ | ✅ | Unchanged |
| Router | `WhitelistedCodeAccountSet` | ✅ | ✅ | Unchanged |
| Router | `Skim` | ❌ | ✅ | **New** |
| Market | `PairCreated` | ✅ | ✅ | Unchanged |
| Market | `PairImplSet` | ✅ | ✅ | Unchanged |
| Market | `FeeCollectorChanged` | ✅ | ❌ | **Removed** |
| Market | `MarketFeesUpdated` | ✅ | ❌ | **Removed** |
| Market | `FeeControllerUpdated` | ❌ | ✅ | **New** |
| Pair | `OrderCreated` | ✅ | ✅ | Unchanged |
| Pair | `OrderMatched` | ✅ | ✅ | Unchanged |
| Pair | `OrderClosed` | ✅ | ✅ | Unchanged |
| Pair | `FeeCollect` | ✅ | ✅ | Modified (params removed) |
| Pair | `TickSizeUpdated` | ✅ | ✅ | Unchanged |
| Pair | `Skim` | ✅ | ✅ | Unchanged |
| Pair | `PairFeesUpdated` | ✅ | ❌ | **Removed** |
| Pair | `FeeControllerUpdated` | ❌ | ✅ | **New** |
| FeeController | `FeeControllerFeesSettled` | ❌ | ✅ | **New** (V2Compat) |
| FeeController | `FeeControllerV3FeesSettled` | ❌ | ✅ | **New** (V3Split) |
