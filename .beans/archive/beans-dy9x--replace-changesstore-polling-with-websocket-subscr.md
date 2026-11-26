---
# beans-dy9x
title: Replace ChangesStore polling with WebSocket subscription
status: scrapped
type: feature
priority: normal
created_at: 2026-03-14T15:06:45Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-5txp
---

ChangesStore uses a 3-second setInterval to poll for file changes, which is inconsistent with the rest of the codebase that uses WebSocket subscriptions. This adds unnecessary network overhead and can show stale data for up to 3 seconds.

## Proposed Fix

Add a GraphQL subscription for file changes (similar to beanChanged) and use it in ChangesStore instead of polling.

## Affected Files

- frontend/src/lib/stores/changesStore.svelte.ts
- internal/graph/schema.graphqls (new subscription)
- internal/graph/schema.resolvers.go (subscription resolver)



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
