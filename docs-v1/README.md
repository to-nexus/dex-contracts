# DEX Contracts V1

A decentralized exchange (DEX) that enables trading of tokens through an order book system. Trade limit and market orders with a simple fee structure and native CROSS coin support.

## 📋 Table of Contents

- [What is DEX V1?](#what-is-dex-v1)
- [Key Features](#key-features)
- [Order Types](#order-types)
- [Fees](#fees)
- [Native CROSS Coin Support](#native-cross-coin-support)
- [Security](#security)
- [License](#license)
- [Disclaimer](#disclaimer)

## 🎯 What is DEX V1?

DEX V1 is the original version of a decentralized exchange protocol that allows users to trade tokens through an order book. It provides a simple and straightforward trading experience with a single fee rate per market.

### Key Highlights

- **Order Book Trading**: Buy and sell orders are matched through an order book
- **Simple Fee Structure**: Each market uses a single fee rate for all trades
- **One Market per Quote Token**: Each quote token can have exactly one market
- **Native CROSS Support**: Use Cross Chain's native CROSS coin directly without manual wrapping

## ✨ Key Features

### Order Types

#### Limit Orders

Place orders at a specific price you choose. Your order will wait in the order book until someone matches it at your price.

**Time-in-Force Options:**
- **Good Till Cancel (GTC)**: Order stays active until you cancel it or it gets filled
- **Immediate Or Cancel (IOC)**: Order executes immediately at your price or better, unfilled portion is cancelled
- **Fill Or Kill (FOK)**: Order must execute completely at your price or better, otherwise cancelled

#### Market Orders

Execute your order immediately at the best available price in the order book. Perfect for when you want to trade right away without waiting.

### Native CROSS Coin Support

The DEX seamlessly supports Cross Chain's native CROSS coin:

- **Send CROSS directly**: You can send native CROSS coins with your transaction (no need to wrap manually)
- **Automatic handling**: The system automatically wraps CROSS for trading and unwraps it when transferring to your wallet
- **Transparent experience**: Use CROSS just like any other token without extra steps

## 💰 Fees

DEX V1 uses a simple fee structure:

- **Single Fee Rate**: Each market has one fee rate that applies to all trades
- **Same for All**: The same fee percentage applies regardless of:
  - Whether it's a buy or sell order
  - Whether it's a limit order (maker) or market order (taker)
- **Market-Level Configuration**: Each market sets its own fee rate

### How Fees Work

When you trade, a percentage of the trade amount is taken as a fee. For example, if a market has a 0.3% fee rate and you trade 1000 tokens, 3 tokens will be collected as a fee.

## 🔒 Security

The DEX contracts follow security best practices to ensure safe trading.

### Security Measures

- All contracts follow security best practices
- Comprehensive input validation
- Protection against reentrancy attacks
- Access control for administrative functions

## 📄 License

This project is licensed under the Business Source License 1.1 (BUSL-1.1). See the [LICENSE](../LICENSE) file for details.

**License Terms:**
- **Licensor**: Nexus Co., Ltd.
- **Change Date**: 2029-10-29
- **Change License**: MIT License (after Change Date)

Until the Change Date, this license permits:
- Copying and modification
- Creating derivative works
- Redistribution
- Non-production use

## ⚠️ Disclaimer

This software is provided "as is" without warranty. Users should conduct their own audits and security reviews before using in production.

---

**Note**: For detailed technical documentation about the contract architecture and implementation, please refer to the contract source code and interfaces.
