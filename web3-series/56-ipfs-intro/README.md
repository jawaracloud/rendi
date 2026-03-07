# 56 — IPFS: Content-Addressable Storage

> **Type:** Explanation | **Language Focus:** Theory

## Objective
Understand how Content IDs (CIDs) differ from location-based URLs.

## HTTP vs IPFS
- **HTTP**: "Get me the file at servers.com/cat.jpg". If the server moves, the link breaks.
- **IPFS**: "Get me the file with hash Qm...cat". If anyone in the network has it, you get it.

## Key Concepts
- **Hashing**: SHA-256 (multihash).
- **DAG (Directed Acyclic Graph)**: How larger files are split into chunks.
- **Pinning**: Ensuring a node keeps a copy and doesn't garbage collect it.

