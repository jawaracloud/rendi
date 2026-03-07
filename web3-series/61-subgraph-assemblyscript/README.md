# 61 — Writing a Subgraph (Schema + Mappings)

> **Type:** Tutorial | **Language Focus:** AssemblyScript

## Objective
Define how to transform raw Ethereum events into a searchable GraphQL API using The Graph.

## Workflow
1. **Schema**: Define entities in GraphQL.
2. **Manifest (`subgraph.yaml`)**: Point to the contract and event handlers.
3. **Mappings**: Write AssemblyScript code that handles the event and saves to the database.

```typescript
export function handleTransfer(event: TransferEvent): void {
  let entity = new Transfer(event.transaction.hash.toHex());
  entity.from = event.params.from;
  entity.to = event.params.to;
  entity.value = event.params.value;
  entity.save();
}
```

