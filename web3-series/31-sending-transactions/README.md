# 31 — Sending Transactions from Go

> **Type:** Tutorial | **Language Focus:** Go

## Objective
Build, sign, and broadcast a transaction to transfer Ether or execute a contract.

## The Workflow
1. **Nonce**: Get next available nonce for account.
2. **Gas Price**: Suggested fee from node.
3. **Transaction**: Create `types.NewTransaction`.
4. **Sign**: Sign with private key using `types.SignTx`.
5. **Send**: Broadcast with `SendTransaction`.

## Sample (Transfer)
```go
nonce, _ := client.PendingNonceAt(ctx, fromAddress)
gasPrice, _ := client.SuggestGasPrice(ctx)
tx := types.NewTransaction(nonce, toAddress, value, 21000, gasPrice, nil)
signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
err := client.SendTransaction(ctx, signedTx)
```

