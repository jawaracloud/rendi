# 44 — Your First Solana Program with Anchor

> **Type:** Tutorial | **Language Focus:** Rust

## Objective
Write a basic Rust program that increments a counter on Solana.

## The Program (`lib.rs`)

```rust
use anchor_lang::prelude::*;

declare_id!("ProgramID...");

#[program]
pub mod counter_program {
    use super::*;

    pub func initialize(ctx: Context<Initialize>) -> Result<()> {
        let counter = &mut ctx.accounts.counter;
        counter.count = 0;
        Ok(())
    }

    pub func increment(ctx: Context<Update>) -> Result<()> {
        let counter = &mut ctx.accounts.counter;
        counter.count += 1;
        Ok(())
    }
}

#[account]
pub struct Counter {
    pub count: u64,
}
```

## Instructions
1. `anchor init my-counter`
2. Replace `programs/my-counter/src/lib.rs` with the code above.
3. `anchor build`
4. `solana-test-validator`
5. `anchor deploy`

