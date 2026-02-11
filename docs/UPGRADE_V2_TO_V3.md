# Upgrade Guide: V2 → V3

This document describes a **safe upgrade path** from the legacy V2 contracts to V3.

V3 introduces a major architectural change: fee logic moves from “fee config structs” to an external
**FeeController** executed via `delegatecall` from `PairImplV3`.

As a result, **some V2 storage slots will be reinterpreted in V3** and must be corrected immediately
after upgrading.

## 📋 Table of Contents

- [What Changes in V3](#-what-changes-in-v3)
- [High-Risk Storage Collisions](#-high-risk-storage-collisions)
- [Recommended Upgrade Order](#-recommended-upgrade-order)
- [FeeController Initialization Data](#-feecontroller-initialization-data)
- [Post-upgrade Verification Checklist](#-post-upgrade-verification-checklist)

## 🧠 What Changes in V3

### Fees: FeeConfig → FeeController

In V2:

- Market and Pair store fee configuration (structs / fields) directly in their own storage.

In V3:

- Market stores a single `feeController` address (the market default).
- Pair stores a single `feeController` address (the pair policy).
- Fee parameters are stored in **namespaced storage (ERC-7201)** inside the Pair, written by the FeeController via `delegatecall`.

### CrossDex adds an allow-list

`CrossDexImplV3` introduces an allow-list of FeeController addresses.
After upgrading CrossDex to V3 this allow-list is **empty** and must be initialized.

## ⚠️ High-Risk Storage Collisions

The V2 → V3 upgrade is not “fire and forget”. Some V2 slots are reinterpreted in V3.

### 1) Pair V2 → Pair V3 (CRITICAL)

In V2:

- Pair has `FeeConfig feeConfig` (packed uint32 fields) in a storage slot.

In V3:

- the same slot becomes `IFeeController feeController` (an address).

After upgrade without mitigation, the old bytes are read as an **invalid address** (e.g. `0x...001E00000014`)
and any call that requires fee logic (e.g. `calcBuyVolumeWithFee`) will revert.

✅ **Mitigation (mandatory):**

- Immediately after upgrading each Pair to `PairImplV3`, call:

```solidity
pairV3.setFeeController(feeControllerAddress, feeControllerInitData);
```

### 2) Market V2 → Market V3 (mandatory mitigation)

In V2:

- Market stores `feeCollector` (an address).

In V3:

- the same slot becomes `feeController` (an address).

After upgrade, you will typically have `feeController == old feeCollector`, which is usually an EOA,
not a FeeController contract (and likely not allowed by CrossDex).

✅ **Mitigation (mandatory):**

- Immediately after upgrading each Market to `MarketImplV3`, call:

```solidity
marketV3.setFeeController(
    /* pairs */ new address[](0),
    /* newFeeController */ feeControllerAddress,
    /* initData */ feeControllerInitData
);
```

Use an empty `pairs` array while your Pairs are still V2 (do not attempt to propagate to V2 Pairs).

### 3) CrossDex V2 → CrossDex V3 (mandatory post-step)

CrossDex V3 introduces the FeeController allow-list.

✅ **Mitigation (mandatory):**

- After upgrading CrossDex to V3, allow your FeeController(s):

```solidity
crossDexV3.setFeeControllerAllow(feeControllerAddress, true);
```

### 4) Router V2 → Router V3 (low risk)

Router V3 changes internal accounting and adds features (e.g., stricter native-value checks),
but the storage layout upgrade is designed to be safe under UUPS.

## ✅ Recommended Upgrade Order

This sequence is validated in `test/V2ToV3Upgrade.t.sol`.

### Step 0: Prepare FeeController implementation(s)

Decide which fee policy you want after upgrade:

- Keep V2-like behavior: deploy and use `FeeControllerV2Compat`
- Switch to new model: deploy and use `FeeControllerV3Split`

### Step 1: Upgrade CrossDex (V2 → V3)

```solidity
UUPSUpgradeable(crossDexProxy).upgradeToAndCall(crossDexV3Impl, hex"");
```

Immediately initialize the allow-list:

```solidity
CrossDexImplV3(crossDexProxy).setFeeControllerAllow(feeControllerAddress, true);
```

### Step 2: Upgrade Router (V2 → V3)

```solidity
UUPSUpgradeable(routerProxy).upgradeToAndCall(routerV3Impl, hex"");
```

### Step 3: Upgrade each Market (V2 → V3)

```solidity
UUPSUpgradeable(marketProxy).upgradeToAndCall(marketV3Impl, hex"");
```

Then **immediately** set Market fee controller (do not propagate to Pairs yet):

```solidity
address[] memory emptyPairs = new address[](0);
MarketImplV3(marketProxy).setFeeController(emptyPairs, feeControllerAddress, feeControllerInitData);
```

### Step 4: Upgrade each Pair (V2 → V3)

```solidity
UUPSUpgradeable(pairProxy).upgradeToAndCall(pairV3Impl, hex"");
```

Then **immediately** fix the storage collision:

```solidity
PairImplV3(pairProxy).setFeeController(feeControllerAddress, feeControllerInitData);
```

### Step 5: Verify and resume operations

After all components are upgraded and FeeControllers are set, perform a small trading test:

- submit a SELL limit order
- submit a BUY limit order that partially matches
- cancel remaining orders
- verify reserves return to 0

See:

- `test/V2ToV3Upgrade.t.sol`
- `test/DEXV3E2EScenarioV2Compat.t.sol`
- `test/DEXV3E2EScenarioV3Split.t.sol`

## 🧾 FeeController Initialization Data

### FeeControllerV2Compat initData

Use the same BPS values you used in V2.

```solidity
bytes memory feeControllerInitData = abi.encode(
    feeCollector,         // address
    sellerMakerFeeBps,    // uint32
    sellerTakerFeeBps,    // uint32
    buyerMakerFeeBps,     // uint32
    buyerTakerFeeBps      // uint32
);
```

### FeeControllerV3Split initData

```solidity
bytes memory feeControllerInitData = abi.encode(
    feeCollector,         // address
    creator,              // address
    takerFeeBps,          // uint32
    creatorShareBps,      // uint32 (share of taker fee)
    makerRebateShareBps   // uint32 (share of taker fee)
);
```

## ✅ Post-upgrade Verification Checklist

After upgrading:

- CrossDex:
  - `ROUTER()` unchanged and points to Router proxy
  - `marketImpl()` / `pairImpl()` unchanged unless intentionally updated
  - `checkFeeControllerAllowed(feeController)` succeeds

- Market (each):
  - `feeController()` equals the intended FeeController contract address (not the old feeCollector EOA)

- Pair (each):
  - `feeController()` equals the intended FeeController contract address
  - calling `calcBuyVolumeWithFee(volume)` does **not** revert

- Router:
  - submitting a minimal order does not revert with `RouterInvalidValue()`

If any of the above fail, the most common cause is that `setFeeController(...)` was not called
immediately after upgrading Market/Pair, or the FeeController was not allow-listed in CrossDex.

