# 26 — go-ethereum (Geth) Overview and ethclient

> **Type:** Explanation | **Language Focus:** Go

## Objective
Learn how to use the `ethclient` package to speak to Ethereum nodes from a Go application.

## Simple Connection

```go
package main

import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Latest Block: %d\n", header.Number.Uint64())
}
```

## Key Concepts
- **JSON-RPC**: The transport protocol.
- **Dial**: Establishing a connection (supports HTTP, WS, and IPC).

