# V1 to V2 Migration Guide

Complete guide for migrating from DEX Contracts V1 to V2.

## 📋 Table of Contents

- [Overview](#overview)
- [Breaking Changes](#breaking-changes)
- [Migration Steps](#migration-steps)
- [Upgrade Process](#upgrade-process)
- [Common Issues](#common-issues)

## 🎯 Overview

V2 introduces several improvements over V1 while maintaining backward compatibility for existing markets. This guide will help you understand the changes and migrate your code.

### Key Improvements in V2

1. **Multiple Markets per Quote**: Create multiple markets with the same quote token
2. **Granular Fee Structure**: Separate fees for maker/taker and buyer/seller
3. **Contract Account Whitelisting**: Allow specific contracts to interact with router
4. **Improved Market Management**: Better market tracking and discovery

## ⚠️ Breaking Changes

### 1. Market Creation API

**V1:**
```solidity
function createMarket(
    address _owner,
    address quote,
    address feeCollector,
    uint256 feeBps  // Single fee rate
) external onlyOwner returns (address);
```

**V2:**
```solidity
function createMarket(
    address _owner,
    address quote,
    address feeCollector,
    bytes memory feeData,  // Encoded 4 fee rates
    string memory message  // Additional identifier
) external onlyOwner returns (address);
```

**Changes:**
- `feeBps` (uint256) → `feeData` (bytes) + `message` (string)
- Fee data must encode: `(sellerMakerFeeBps, sellerTakerFeeBps, buyerMakerFeeBps, buyerTakerFeeBps)`

### 2. Market Storage Mapping

**V1:**
```solidity
EnumerableMap.AddressToAddressMap private _allMarkets; // quote => market
function quoteToMarket(address quote) external view returns (address);
```

**V2:**
```solidity
EnumerableMap.AddressToAddressMap private _allMarkets; // market => quote
// quoteToMarket() removed
function allMarkets() external view returns (address[] memory markets, address[] memory quotes);
```

**Changes:**
- Mapping direction reversed: `quote => market` → `market => quote`
- `quoteToMarket()` function removed
- `allMarkets()` now returns `(markets[], quotes[])` instead of `(quotes[], markets[])`

### 3. Fee Structure

**V1:**
```solidity
uint32 public feeBps; // Single fee for all trades
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

**Changes:**
- Single fee rate → 4 separate fee rates
- Fees can now differ by order side and type

### 4. Router Modifier

**V1:**
```solidity
modifier checkValue() {
    _;
    if (address(this).balance != 0) revert RouterInvalidValue();
}
// All contract accounts blocked
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

**Changes:**
- `checkValue` → `checkSubmit`
- Whitelist system for contract accounts
- `setWhitelistedCodeAccount()` function added

## 🔄 Migration Steps

### Step 0: CROSS (Native Coin) Handling

Both V1 and V2 use the same WETH wrapper for Cross Chain native coin (CROSS). No migration needed for CROSS handling:

- **Pair Internal**: CROSS is treated as ERC20 within pairs (same in both versions)
- **External Transfers**: CROSS automatically unwraps to native coin when transferred to non-pair addresses (same in both versions)
- **User Experience**: Users can continue sending native CROSS via `payable` functions (same in both versions)

The WETH wrapper contract behavior is identical in V1 and V2, so existing CROSS integration code should work without changes.

### Step 1: Update Market Creation Code

**Before (V1):**
```solidity
uint256 feeBps = 30; // 0.3%
address market = crossDex.createMarket(owner, USDT, feeCollector, feeBps);
```

**After (V2):**
```solidity
// Define 4 fee rates
uint32 sellerMakerFee = 25;  // 0.25% for sellers (limit orders)
uint32 sellerTakerFee = 30;  // 0.30% for sellers (market orders)
uint32 buyerMakerFee = 25;    // 0.25% for buyers (limit orders)
uint32 buyerTakerFee = 30;    // 0.30% for buyers (market orders)

// Encode fee data
bytes memory feeData = abi.encode(sellerMakerFee, sellerTakerFee, buyerMakerFee, buyerTakerFee);

// Create market with message identifier
string memory message = "Main USDT Market";
address market = crossDex.createMarket(owner, USDT, feeCollector, feeData, message);
```

### Step 2: Update Market Lookup Code

**Before (V1):**
```solidity
address market = crossDex.quoteToMarket(USDT);
```

**After (V2):**
```solidity
// Option 1: Get all markets and filter
(address[] memory markets, address[] memory quotes) = crossDex.allMarkets();
address market;
for (uint i = 0; i < markets.length; i++) {
    if (quotes[i] == USDT && /* additional filter */) {
        market = markets[i];
        break;
    }
}

// Option 2: Store market address from creation event
// Recommended: Track market addresses from MarketCreated events
```

### Step 3: Update Fee Configuration Code

**Before (V1):**
```solidity
market.setFeeBps(30); // Single fee update
```

**After (V2):**
```solidity
// Update all 4 fee rates
market.setMarketFees(
    25,  // sellerMakerFeeBps
    30,  // sellerTakerFeeBps
    25,  // buyerMakerFeeBps
    30   // buyerTakerFeeBps
);

// Or query current fees
IMarketV2.FeeConfig memory fees = market.getFeeConfig();
```

### Step 4: Update Router Interaction Code

**Before (V1):**
```solidity
// Contract accounts cannot interact
// All contract calls blocked
```

**After (V2):**
```solidity
// Whitelist contract accounts that need access
address[] memory contracts = new address[](1);
contracts[0] = myContract;
router.setWhitelistedCodeAccount(contracts, true);

// Now myContract can interact with router
```

## 🔧 Upgrade Process

### For Existing Deployments

If you have existing V1 deployments:

1. **Deploy V2 Implementation Contracts**
   ```bash
   # Deploy new implementations
   forge script script/UpgradeCrossDexV2.s.sol:UpgradeCrossDexV2 --rpc-url <network>
   ```

2. **Upgrade CrossDex Proxy**
   ```solidity
   // Upgrade the proxy to point to V2 implementation
   crossDex.upgradeTo(v2Implementation);
   ```

3. **Reinitialize if Needed**
   ```solidity
   // V2 adds reinitialize() for storage migration
   crossDex.reinitialize(newMarketImpl, newPairImpl);
   ```

4. **Migrate Existing Markets** (if needed)
   - Existing markets remain functional
   - New markets should use V2 API
   - Consider migrating to new fee structure

### Storage Migration

V2 includes a `reinitialize()` function for storage migration:

```solidity
function reinitialize(address _marketImpl, address _pairImpl) external onlyOwner reinitializer(2) {
    // Migrates _allMarkets from quote=>market to market=>quote
    // Updates implementation addresses
}
```

## 🐛 Common Issues

### Issue 1: Market Not Found

**Problem:**
```solidity
// V1 code trying to use V2 contract
address market = crossDex.quoteToMarket(USDT); // ❌ Function doesn't exist
```

**Solution:**
```solidity
// Use allMarkets() and filter
(address[] memory markets, address[] memory quotes) = crossDex.allMarkets();
// Find market with matching quote
```

### Issue 2: Fee Encoding Error

**Problem:**
```solidity
// Incorrect fee encoding
bytes memory feeData = abi.encode(30); // ❌ Wrong: single value
```

**Solution:**
```solidity
// Correct: encode all 4 values
bytes memory feeData = abi.encode(
    uint32(25), // sellerMakerFeeBps
    uint32(30), // sellerTakerFeeBps
    uint32(25), // buyerMakerFeeBps
    uint32(30)  // buyerTakerFeeBps
);
```

### Issue 3: Contract Account Blocked

**Problem:**
```solidity
// Contract trying to call router (V2)
contract MyBot {
    function trade() external {
        router.submitBuyLimit(...); // ❌ Reverts: contract account blocked
    }
}
```

**Solution:**
```solidity
// Whitelist the contract first
router.setWhitelistedCodeAccount([address(myBot)], true);
// Then it can trade
```

### Issue 4: Storage Slot Collision

**Problem:** Upgrading without proper migration may cause storage issues.

**Solution:**
- Always use `reinitialize()` for storage migration
- Verify storage layout compatibility
- Test thoroughly on testnet first

## 📚 Additional Resources

- [V1 README](./README.md) - V1 overview and features
- [Main README](../README.md) - V2 documentation

---

**Need Help?** If you encounter issues during migration, please review the test files in `test/DEXV2CrossDexUpgrade.t.sol` for reference implementations.

