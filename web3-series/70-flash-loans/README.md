# 70 — Flash Loans: Concept and Go Simulation

> **Type:** Tutorial | **Language Focus:** Go + Solidity

## Objective
Understand atomic transactions where you borrow millions without collateral—provided you pay it back in the same block.

## How it works
1. **Borrow**: Your contract asks Aave/Uniswap for $1M.
2. **Execute**: You perform an arbitrage or liquidation.
3. **Repay**: You pay back $1M + small fee.
4. **Verification**: If step 3 fails, the entire transaction (including step 1) is reverted as if it never happened.

## Instructions
Review the `FlashLoanReceiver.sol` pattern in the reference material.

