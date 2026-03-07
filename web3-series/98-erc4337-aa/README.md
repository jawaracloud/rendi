# 98 — Account Abstraction (ERC-4337)

> **Type:** Explanation | **Language Focus:** Theory

## Objective
Understand how 'Smart Accounts' replace 'Private Key Accounts' (EOAs).

## Features
- **Social Recovery**: Gain access without a seed phrase.
- **Gas Abstraction**: Pay gas in USDC or have it sponsored by a 'Paymaster'.
- **Bundling**: Approve and Swap in a single click.
- **Security**: Custom logic like daily spending limits.

## How it works
Users send 'UserOperations' to a separate mempool. 'Bundlers' package these and send them to an 'EntryPoint' contract on-chain.

