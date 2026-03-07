# 20 — ERC-721 Token Standard (NFTs)

> **Type:** Tutorial | **Language Focus:** Solidity

## Objective
Implementation of a Non-Fungible Token (NFT) using the OpenZeppelin library.

## The Code

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721, Ownable {
    uint256 private _nextTokenId;

    constructor() ERC721("MyNFT", "MNFT") Ownable(msg.sender) {}

    function safeMint(address to) public onlyOwner {
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
    }
}
```

## Key Concepts
- **_safeMint**: Ensures the receiving address is a contract that can handle ERC-721s (prevents 'lost' NFTs).
- **Ownable**: Restricts minting to the contract creator.

