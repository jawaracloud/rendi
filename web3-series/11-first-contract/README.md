# 11 — What Is Solidity? Your First Smart Contract

> **Type:** Tutorial | **Language Focus:** Solidity

## Objective
Write, compile, and understand a basic 'Hello World' storage contract.

## Prerequisites
- Foundry installed (or use the dev container).

## The Contract (`SimpleStorage.sol`)

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

contract SimpleStorage {
    uint256 private number;

    event NumberUpdated(uint256 newValue);

    function set(uint256 _number) public {
        number = _number;
        emit NumberUpdated(_number);
    }

    function get() public view returns (uint256) {
        return number;
    }
}
```

## Instructions
1. Initialize a new Foundry project: `forge init`
2. Save the code in `src/SimpleStorage.sol`.
3. Compile: `forge build`
4. Run a local test node: `anvil`
5. Deploy (local): `forge create --rpc-url http://127.0.0.1:8545 --private-key <ANVIL_KEY> src/SimpleStorage.sol:SimpleStorage`

