# Architecture (DEX Contracts V3)

This document explains the **V3 system architecture**, component responsibilities, and the most important call flows and invariants.

## 📋 Table of Contents

- [System Layers](#-system-layers)
- [Contract Relationship Diagram](#-contract-relationship-diagram)
- [Deployment Model (Proxies + Create2)](#-deployment-model-proxies--create2)
- [Core Trading Flow](#-core-trading-flow)
- [Order Book Data Structures](#-order-book-data-structures)
- [Access Control Model](#-access-control-model)
- [Key Invariants](#-key-invariants)

## 🧱 System Layers

V3 is built as 4 “layers” + a pluggable fee layer:

1. **CrossDex (global)**: deploys Router proxy, creates Markets, tracks (Market ↔ Quote), tracks (Pair → Market), and controls allowed FeeControllers.
2. **Router (entry point)**: validates pair addresses, transfers tokens, wraps native CROSS, submits/cancels orders, applies anti-reentrancy and value-accounting checks.
3. **Market (per quote token)**: creates Pairs (BASE/QUOTE) via Create2 and registers them in CrossDex.
4. **Pair (matching engine)**: stores the order book, matches orders, tracks reserves, and executes fee logic via FeeController.
5. **FeeController (policy)**: a separate contract that is executed via `delegatecall` from Pair and can implement different fee policies.

## 🧩 Contract Relationship Diagram

```mermaid
graph TB
    User[User / Integrator]
    CrossDex[CrossDexImplV3<br/>Global Registry]
    Router[CrossDexRouterV3<br/>Entry Point]
    Market[MarketImplV3<br/>Per-Quote Market]
    Pair[PairImplV3<br/>Order Book + Matching]
    Fee[FeeController<br/>(delegatecall)]
    WETH[WETH (wrapped CROSS)]

    User -->|submit/cancel| Router
    Router -->|isPair() check| CrossDex
    CrossDex -->|createMarket (Create2)| Market
    Market -->|createPair (Create2)| Pair
    Market -->|pairCreated| CrossDex
    Pair -->|delegatecall: fees| Fee
    Router -->|wrap/unwrap native CROSS| WETH

    style CrossDex fill:#e1f5ff
    style Router fill:#fff4e1
    style Market fill:#e8f5e9
    style Pair fill:#fce4ec
    style Fee fill:#ede7f6
    style WETH fill:#f3e5f5
```

## 🏗️ Deployment Model (Proxies + Create2)

V3 uses **UUPS** implementation contracts behind `ERC1967Proxy`.

- **CrossDexImplV3** is deployed as an implementation, and a proxy is created externally (or in tests).
- **Router** is deployed as an `ERC1967Proxy` *from inside CrossDex.initialize()*, and then initialized.
- **Markets** are deployed via **Create2** as `ERC1967Proxy` instances pointing to `MarketImplV3`.
- **Pairs** are deployed via **Create2** as `ERC1967Proxy` instances pointing to `PairImplV3`.

Deterministic salts:

- **Market** salt: `keccak256(abi.encode(quote, message))`
- **Pair** salt: `keccak256(abi.encodePacked(base))` (per Market)

This enables deterministic addresses while still allowing multiple markets per quote token in V2/V3 (because the salt includes `message`).

## 🔄 Core Trading Flow

### Limit order submission

Limit order submission is always routed through the Router:

1. **Router validates**
   - `isPair(pair)` via CrossDex registry (`pairToMarket(pair) != 0`)
   - caller restrictions (EOA-only by default, unless whitelisted)
2. **Router transfers funds**
   - For SELL: transfers BASE (or wraps native CROSS and mints to the Pair)
   - For BUY: transfers QUOTE *including required fee pre-funding* (policy-dependent)
3. **Router calls Pair**
   - `Pair.findPrevPrice(...)` (off-chain or on-chain via Router) is used to compute insertion hints
   - `Pair.submitLimitOrder(...)` executes matching immediately if possible
4. **Pair matches and settles**
   - matches against opposite book until filled / no liquidity / `maxMatchCount` reached
   - uses FeeController hooks during matching and settles fees at the end of the match set
5. **Router post-check**
   - ensures no leftover native value remains (prevents accidental ETH retention and mitigates forced-ETH DoS patterns)

### Market order submission

Market orders are also routed through the Router. The Pair treats them as:

- **SELL market**: “price goes down until amount is sold”
- **BUY market**: “price goes up until quote budget is exhausted”

The same fee settlement mechanism applies, but the inputs differ (BUY uses quote budget, SELL uses base amount).

## 🧾 Order Book Data Structures

`PairImplV3` maintains:

- **Price levels** using `List.U256`:
  - SELL prices: **ascending** (head = cheapest)
  - BUY prices: **descending** (head = most expensive)
- **FIFO order queues per price**
  - `_sellOrders[price]` and `_buyOrders[price]` store order IDs in chronological order
- **Orders mapping**
  - `_allOrders[orderId] -> Order { side, owner, feeBps, price, amount }`

The key optimization is that the Router can provide an **adjacent hint** (`uint256[2] adjacent`) that helps the Pair find an insertion location without scanning from the head every time. This bounds gas via `findPrevPriceCount`.

## 🔐 Access Control Model

At a high level:

- **CrossDex owner**
  - creates markets
  - updates implementation addresses (`marketImpl`, `pairImpl`)
  - sets tick size setter
  - controls the FeeController allow-list
- **Market owner**
  - creates pairs
  - can update market-wide `feeController` and propagate to selected pairs
- **Pair “owner”**
  - the Market owner (Pair delegates ownership to Market.owner())
  - can pause and emergency-cancel orders, and update tick sizes (if authorized)
- **Router owner**
  - CrossDex owner (Router delegates ownership to CrossDex.owner())

## ✅ Key Invariants

These invariants are relied upon throughout the system:

- **Pair validity**
  - A Pair is valid if `CrossDex.pairToMarket(pair) != address(0)`.
- **Tick/lot divisibility**
  - `tickSize > 0`, `lotSize > 0`
  - `tickSize * lotSize % DENOMINATOR == 0` (where `DENOMINATOR = 10 ** BASE.decimals()`).
- **Reserve safety**
  - Pair checks that on-chain token balances cover reserves and in-flight deposits before matching.
- **Fee settlement model**
  - FeeController functions are executed via `delegatecall` and must assume Pair storage context.
  - Per-transaction fee accumulators use transient storage (EIP-1153) to avoid persistent writes.
- **Router value accounting**
  - The Router enforces that native balance does not change unexpectedly during submission calls.

