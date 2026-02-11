# DEX Contracts V3

A decentralized exchange (DEX) that enables trading of tokens through an **on-chain order book**. V3 keeps the familiar V1/V2 trading UX (limit/market orders + native CROSS support) while modernizing fees via **pluggable Fee Controllers** and documenting a safe **V2 → V3 upgrade path**.

## 📋 Table of Contents

- [What is DEX V3?](#-what-is-dex-v3)
- [High-level Architecture](#-high-level-architecture)
- [Fees (V3 FeeController)](#-fees-v3-feecontroller)
- [Native CROSS Coin Support](#-native-cross-coin-support)
- [Documentation](#-documentation)
- [Development](#-development)
- [Security](#-security)
- [License](#-license)
- [Disclaimer](#-disclaimer)
- [Legacy V1/V2](#-legacy-v1v2)

## 🎯 What is DEX V3?

DEX V3 is an order-book DEX protocol composed of upgradeable, modular contracts:

- **CrossDex**: global registry + market creation
- **Router**: user entry point for submitting/canceling orders and handling native CROSS
- **Market**: per-quote market that deploys pairs
- **Pair**: the matching engine + order book
- **FeeController**: fee logic called via `delegatecall` from `Pair`

V3’s key design shift is that fee logic lives in a separate contract (FeeController) so the matching engine can stay stable while fee policies evolve.

## 🏗️ High-level Architecture

```text
User  ──>  CrossDexRouterV3  ──>  PairImplV3  ──>  Order Book + Matching
                 │                 │
                 │                 └─ delegatecall ──> FeeController (policy)
                 │
                 └─ isPair() / registry lookups ──> CrossDexImplV3 / MarketImplV3
```

The complete architecture, call flows, and invariants are documented in `docs/`.

## 💰 Fees (V3 FeeController)

In V3, each `PairImplV3` holds a `feeController` address and calls it via `delegatecall`.
Two fee controller implementations are included:

- **`FeeControllerV2Compat`**: V2-compatible 4-bps model (seller/buyer × maker/taker)
- **`FeeControllerV3Split`**: taker-only fee with a 3-way split (creator / maker rebate / system)

See `docs/FEES.md` for details.

## 🪙 Native CROSS Coin Support

The Router deploys a `WETH`-style wrapper for native CROSS and integrates it seamlessly:

- When trading with CROSS as BASE/QUOTE, users provide `msg.value` and the Router wraps it.
- When a transfer sends wrapped CROSS to a non-pair address, `WETH` auto-unwraps and transfers native CROSS.

## 📚 Documentation

Start here: `docs/README.md`.

Key documents:

- `docs/ARCHITECTURE.md`: system overview, diagrams, and call flows
- `docs/ROUTER_GUIDE.md`: how to submit/cancel orders (including adjacent hints, fees, and native CROSS)
- `docs/FEES.md`: how FeeController works and how to configure it
- `docs/UPGRADE_V2_TO_V3.md`: required upgrade steps + storage collision notes

## 🧰 Development

This is a Foundry project.

```bash
forge --version
forge build
forge test
```

### Running the V2 → V3 upgrade tests

The upgrade tests load V2 bytecode from `solc_0_8_28/out` via `vm.getCode(...)`.
Build legacy artifacts first:

```bash
cd solc_0_8_28
forge build
cd ..
forge test --match-path test/V2ToV3Upgrade.t.sol
```

## 🔒 Security

The contracts use established security patterns:

- **UUPS upgradeability** (`upgradeToAndCall`) with `onlyOwner` authorization
- **Deterministic deployments** via Create2 for Markets and Pairs
- **Reentrancy protection** in the Router (`ReentrancyGuardTransient`)
- **Contract-account restrictions** (EOA-only by default, with an owner-managed whitelist)
- **Native value accounting** to prevent leftover ETH and mitigate forced-ETH scenarios

Audit report: `audits/REP-final-20251103T123743Z.pdf`.

## 📄 License

This project is licensed under the Business Source License 1.1 (BUSL-1.1). See `LICENSE`.

## ⚠️ Disclaimer

This software is provided "as is" without warranty. Users should conduct their own audits and security reviews before using in production.

## 🕰️ Legacy V1/V2

Historical V1/V2 code and docs are available under `solc_0_8_28/`.
