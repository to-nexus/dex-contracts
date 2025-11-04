# DEX Contracts V2

A decentralized exchange (DEX) that enables trading of tokens through an order book system. Trade limit and market orders with flexible fee structures and native CROSS coin support.

## 📋 Table of Contents

- [What is DEX V2?](#what-is-dex-v2)
- [Key Features](#key-features)
- [Order Types](#order-types)
- [Fees](#fees)
- [Native CROSS Coin Support](#native-cross-coin-support)
- [Security](#security)
- [License](#license)
- [Disclaimer](#disclaimer)
- [Technical Documentation](#technical-documentation)

## 🎯 What is DEX V2?

DEX V2 is a decentralized exchange protocol that allows users to trade tokens through an order book. Unlike automated market makers (AMMs), this DEX uses a traditional order book model where buy and sell orders are matched at specified prices.

### Key Highlights

- **Order Book Trading**: Buy and sell orders are matched through an order book, giving you control over your trade prices
- **Flexible Fee Structure**: Each market can have different fee rates for makers and takers
- **Multiple Markets**: Multiple markets can exist for the same quote token, each with different fee policies
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

DEX V2 uses a flexible fee structure with four separate fee rates:

1. **Seller Maker Fee**: Fee for sellers who place limit orders
2. **Seller Taker Fee**: Fee for sellers who execute market orders
3. **Buyer Maker Fee**: Fee for buyers who place limit orders
4. **Buyer Taker Fee**: Fee for buyers who execute market orders

### How Fees Work

- Each market sets its own fee rates
- **Maker orders** (limit orders): Typically lower fees since you provide liquidity to the order book
- **Taker orders** (market orders): Typically higher fees since you consume liquidity
- Fees can be different for buy vs sell orders

The fee structure allows markets to incentivize certain trading behaviors, such as encouraging limit orders by offering lower maker fees.

## 🔒 Security

The DEX contracts have been professionally audited to ensure security and reliability.

### Security Measures

- All contracts follow security best practices
- Comprehensive input validation
- Protection against reentrancy attacks
- Access control for administrative functions

For detailed security information, see the audit report in [`audits/REP-final-20251103T123743Z.pdf`](audits/REP-final-20251103T123743Z.pdf).

## 📄 License

This project is licensed under the Business Source License 1.1 (BUSL-1.1). See the [LICENSE](./LICENSE) file for details.

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

## 📚 Technical Documentation

For detailed technical documentation about the contract architecture, implementation details, and V2 changes, see [Technical Documentation](docs-v1/TECHNICAL.md).
