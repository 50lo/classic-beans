---
# beans-kj41
title: GraphQL subscriptions for bean changes
status: scrapped
type: feature
priority: normal
created_at: 2025-12-20T09:50:21Z
updated_at: 2026-11-26T10:52:04Z
order: Vz
---

Add real-time GraphQL subscriptions for bean change events.

## Changes
- Add Subscription type to GraphQL schema with beanChanged field
- Add BeanChangeEvent type and ChangeType enum (CREATED/UPDATED/DELETED)
- Implement subscription resolver using beancore.Subscribe()
- Add WebSocket transport to serve command
- Start file watcher automatically when server starts

## Files
- internal/graph/schema.graphqls
- internal/graph/schema.resolvers.go
- internal/graph/model/models_gen.go (generated)
- internal/graph/generated.go (generated)
- cmd/serve.go



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
