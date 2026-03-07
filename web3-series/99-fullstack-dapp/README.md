# 99 — Building a Full-Stack DApp (Go Backend)

> **Type:** Tutorial | **Language Focus:** Go + Solidity

## Objective
The Capstone project: Connect a Solidity contract, a Go indexer, and an IPFS storage layer.

## The Architecture
1. **Foundry**: Smart contracts for voting.
2. **Kubo (Go)**: Poll for IPFS CIDs where candidate bios are stored.
3. **Go Backend**:
    - Polls for `Voted` events via `ethclient`.
    - Stores aggregated results in a cache.
    - Serves a JSON API for the final results.

## Final Instructions
Run the integrated script in `scripts/deploy-capstone.sh` to see all components work together.

