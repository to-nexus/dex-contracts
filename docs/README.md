# DEX Contracts V3 Documentation

This directory contains **detailed technical documentation** for the V3 CrossDEX order-book contracts.

## 📋 Table of Contents

- [Start Here](#-start-here)
- [Documentation Map](#-documentation-map)
- [Key Concepts](#-key-concepts)
- [Versioning and Legacy](#-versioning-and-legacy)

## 🚀 Start Here

Recommended reading order:

1. **Architecture overview**: [`ARCHITECTURE.md`](./ARCHITECTURE.md)
2. **How to trade via Router**: [`ROUTER_GUIDE.md`](./ROUTER_GUIDE.md)
3. **Fees and FeeController**: [`FEES.md`](./FEES.md)
4. **Upgrade (V2 → V3)**: [`UPGRADE_V2_TO_V3.md`](./UPGRADE_V2_TO_V3.md)
5. **Events Changelog (V2 → V3)**: [`EVENTS_CHANGELOG.md`](./EVENTS_CHANGELOG.md)

## 🗺️ Documentation Map

- `ARCHITECTURE.md`: components, responsibilities, call flows, and invariants
- `FEES.md`: FeeController design, configuration, and settlement flow
- `ROUTER_GUIDE.md`: end-user / integrator guide for submitting and canceling orders
- `UPGRADE_V2_TO_V3.md`: safe migration plan and storage collision notes
- `EVENTS_CHANGELOG.md`: all event changes between V2 and V3 (for indexer migration)

## 🧠 Key Concepts

- **Upgradeability**: the system uses **UUPS** implementations behind `ERC1967Proxy`.
- **Deterministic deployments**: Markets and Pairs are deployed via **Create2**.
- **FeeController**: fee logic is **pluggable** and executed in Pair context via `delegatecall`.
- **Native CROSS**: Router-owned `WETH` wrapper provides a seamless native coin experience.
- **EOA-only default**: Router blocks contract accounts by default (whitelist available).

## 🕰️ Versioning and Legacy

- **V1/V2 legacy** lives under `solc_0_8_28/`.
- The V3 docs are written to be self-contained, but legacy docs can be used for historical context.

