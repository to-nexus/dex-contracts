# Router Guide (CrossDexRouterV3)

This guide explains how to interact with the V3 order book via `CrossDexRouterV3`.

It is written for:

- frontends
- bots
- integrators
- power users (advanced EOA users)

## 📋 Table of Contents

- [Router Responsibilities](#-router-responsibilities)
- [Finding Market / Pair Addresses](#-finding-market--pair-addresses)
- [ERC20 Approvals](#-erc20-approvals)
- [Adjacent Hint (Order Book Insertion Optimization)](#-adjacent-hint-order-book-insertion-optimization)
- [Limit Orders](#-limit-orders)
- [Market Orders](#-market-orders)
- [Cancel Orders](#-cancel-orders)
- [Native CROSS Handling](#-native-cross-handling)
- [Contract Account Restriction (EOA-only default)](#-contract-account-restriction-eoa-only-default)
- [Common Reverts / Troubleshooting](#-common-reverts--troubleshooting)

## 🎛️ Router Responsibilities

The Router is the **only supported entry point** for user trading actions:

- validates that a `pair` is registered (via CrossDex registry)
- transfers required tokens from user to the Pair
- wraps native CROSS into the Router-owned `WETH` contract when needed
- enforces non-reentrancy and native value accounting
- forwards the order to the Pair matching engine

## 🔎 Finding Market / Pair Addresses

### CrossDex → Router

The Router address is stored in CrossDex:

- `CrossDexImplV3.ROUTER()` returns the Router proxy address.

### CrossDex → Markets

CrossDex tracks all markets and their quote tokens:

- `CrossDexImplV3.allMarkets()` returns `(markets[], quotes[])`.

Markets are created by the CrossDex owner:

- `CrossDexImplV3.createMarket(owner, quote, feeController, message)`

### Market → Pairs

Within a Market:

- `MarketImplV3.baseToPair(base)` returns the Pair address for a BASE token.
- `MarketImplV3.allPairs()` returns `(bases[], pairs[])`.

## ✅ ERC20 Approvals

If the order requires ERC20 transfers, approve the Router as spender:

- SELL orders: approve **BASE** token to Router
- BUY orders: approve **QUOTE** token to Router

Native CROSS does **not** require approvals (see [Native CROSS Handling](#-native-cross-handling)).

## 🧭 Adjacent Hint (Order Book Insertion Optimization)

The Router forwards `adjacent` (`uint256[2]`) to the Pair to optimize insertion/search.

Concept:

- Price levels are stored in a sorted linked-list.
- Finding where to insert a new limit order can be expensive.
- `adjacent` provides “nearby” price hints that the Pair can start searching from.

Recommended usage:

- If you don’t have order-book context: use `[0, 0]`
- If you know a nearby existing price: set `adjacent[0] = thatPrice`
- If you know two nearby prices: set `adjacent = [priceA, priceB]`

The Pair will use the first hint that exists; otherwise it falls back to the list head.

Gas bounding:

- the Router has `findPrevPriceCount`, a maximum number of list steps to search.

## 🧾 Limit Orders

Limit orders support the following “time in force” constraints:

- `GOOD_TILL_CANCEL` (GTC): stays on the book
- `IMMEDIATE_OR_CANCEL` (IOC): match what you can immediately; cancel the rest
- `FILL_OR_KILL` (FOK): must fill entirely immediately, or revert

### Submit SELL limit

Function:

- `submitSellLimit(pair, price, amount, constraints, adjacent, maxMatchCount)`

Interpretation:

- You are selling `amount` of BASE at `price` (QUOTE per BASE), or better.
- Your BASE is transferred to the Pair first.
- The Pair matches against existing BUY orders (highest price first).

Validation highlights (Pair-side):

- `price % tickSize == 0`
- `amount % lotSize == 0`

### Submit BUY limit

Function:

- `submitBuyLimit(pair, price, amount, constraints, adjacent, maxMatchCount)`

Interpretation:

- You are buying `amount` of BASE and willing to pay up to `price`.
- Router calculates the QUOTE deposit as:
  - `baseVolume = price * amount / DENOMINATOR`
  - **then** adds the required fee using `Pair.calcBuyVolumeWithFee(baseVolume)`
- That fee logic comes from the Pair’s FeeController (see `docs/FEES.md`).

Important nuance:

- Even though your order is a *limit* order, Router pre-funds BUY with **taker-fee** assumptions
  because a BUY limit can immediately match and become a taker.

### `maxMatchCount`

Router and Pair support a match-count bound:

- If `maxMatchCount == 0`, Router applies its configured default.
- This prevents a single order from consuming too much gas by matching too many fills.

## 🏃 Market Orders

Market orders execute immediately against the best available prices.

### Submit SELL market

Function:

- `submitSellMarket(pair, amount, maxMatchCount)`

Meaning:

- sell exactly `amount` of BASE, consuming the BUY book from best price downward

### Submit BUY market

Function:

- `submitBuyMarket(pair, quoteVolume, maxMatchCount)`

Meaning:

- spend up to `quoteVolume` (in QUOTE units) to buy as much BASE as possible
- Router will pre-fund fee requirements:
  - `required = Pair.calcBuyVolumeWithFee(quoteVolume)`
  - transfers `required` into the Pair

Note:

- `quoteVolume` is the intended “spend” budget *excluding* fee.
- The fee policy decides how much extra must be transferred.

## 🧹 Cancel Orders

Function:

- `cancelOrder(pair, orderIds[])`

Safety:

- The Router enforces an upper bound `cancelLimit`.
- If `orderIds.length > cancelLimit`, it reverts.

Cancellation is processed by the Pair and refunds any remaining reserved tokens back to the order owner.

## 🪙 Native CROSS Handling

V3 supports native CROSS via a Router-owned `WETH` wrapper:

- Router deploys `WETH` during its initialization.
- When a BASE or QUOTE token equals the Router’s `CROSS` address, Router uses `msg.value` and mints wrapped CROSS directly to the Pair (`WETH.mintTo{value: ...}(pair)`).

Behavior when receiving wrapped CROSS:

- If wrapped CROSS is transferred to a **non-Pair** address, the wrapper auto-burns and sends native CROSS.
- If transferred to a **Pair** address, it remains wrapped (Pairs must hold ERC20 balances for accounting).

## 🚫 Contract Account Restriction (EOA-only default)

The Router blocks contract accounts by default:

- if `msg.sender.code.length != 0`, it reverts
- unless the account is explicitly whitelisted by the owner

Whitelisting:

- `setWhitelistedCodeAccount(accounts[], whitelisted)`

Important caveat:

- Contracts can bypass this check while in construction (`code.length == 0` during constructor).

## 🧯 Common Reverts / Troubleshooting

### Router-level reverts

- `RouterInvalidPairAddress(pair)`
  - the `pair` is not registered in CrossDex (`pairToMarket(pair) == 0`)
- `RouterContractAccountBlocked(account)`
  - caller is a contract and not whitelisted
- `RouterCancelLimitExceeded(length, limit)`
  - too many order IDs in one cancel call
- `RouterInvalidValue()`
  - leftover native value detected (usually means wrong `msg.value` usage)

### Pair-level reverts (most common)

- `PairInvalidPrice(price)`
  - not divisible by tick size, or zero
- `PairInvalidAmount(amount)`
  - not divisible by lot size, or zero
- `PairInvalidTickSize(tickSize, lotSize, denominator)`
  - tick/lot produce a volume that cannot be represented precisely
- FeeController-related errors
  - indicates fee controller is misconfigured or not set (common after upgrades if not initialized)

