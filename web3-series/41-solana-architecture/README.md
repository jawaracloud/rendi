# 41 — Solana Architecture: Accounts and Programs

> **Type:** Explanation | **Language Focus:** Theory

## Objective
Understand Solana's 'Sealevel' runtime and why it scales differently than EVM.

## Key Differences

| Feature | Ethereum (EVM) | Solana |
|---------|----------------|---------|
| **Execution** | Sequential (mostly) | Parallel (Sealevel) |
| **State** | Global State Tree | Localized Accounts |
| **Code** | Smart Contracts (mutable?) | Programs (stateless/immutable) |
| **Data** | Stored in contract state | Stored in separate Accounts |

## Critical Concept: Accounts
In Solana, everything is an Account. Programs are just accounts marked 'executable'. Data is stored in 'Data Accounts' owned by programs.

