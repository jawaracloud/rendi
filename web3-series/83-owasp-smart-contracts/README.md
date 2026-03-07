# 83 — Common Smart Contract Vulnerabilities

> **Type:** Explanation | **Language Focus:** Solidity

## Objective
Learn the 'Top 10' ways smart contracts get drained.

## The Hall of Shame
1. **Reentrancy**: Calling an external contract before updating your state.
2. **Arithmetic Over/Underflow**: (Fixed in Solidity 0.8+ but still relevant for low-level math).
3. **Access Control**: Missing `onlyOwner` or `onlyRole`.
4. **Front-running**: Competitors seeing your transaction in the mempool and outbidding your gas.
5. **Logic Errors**: Business logic bugs that bypass security.

## Mitigation
- Use OpenZeppelin's `ReentrancyGuard`.
- Always update state *before* external calls (Check-Effects-Interactions pattern).

