# 04 — Accounts, Wallets, and Key Pairs

> **Type:** Explanation | **Language Focus:** Theory

## Objective
Learn the math behind Web3 identities: Public/Private keys and Seed Phrases.

## The Hierarchy
1. **Entropy**: Randomness (128-256 bits).
2. **Mnemonic (BIP-39)**: 12-24 human-readable words derived from entropy.
3. **Seed**: A 512-bit binary string derived from the mnemonic.
4. **Private Key**: A secret 256-bit number.
5. **Public Key**: Derived via Elliptic Curve Cryptography (ECDSA/Ed25519).
6. **Address**: Hashed version of the public key (e.g., 0x... for ETH).

## Key Concepts
- **BIP-39**: Mnemonic standard.
- **BIP-44**: Hierarchical Deterministic (HD) Wallet derivation paths.
- **Non-custodial**: You own your keys; you are the only one who can sign transactions.

## Environment Check
Verify you have a wallet like MetaMask (EVM) or Phantom (Solana) ready for later lessons.

