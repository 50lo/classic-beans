---
# beans-tx7i
title: Add frontend unit tests for stores and key logic
status: scrapped
type: task
priority: normal
created_at: 2026-03-14T15:06:45Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-5txp
---

The frontend has essentially no unit tests (only nameGenerator.test.ts). The stores and drag-and-drop logic contain significant business logic that would benefit from fast, isolated tests.

## Scope

- BeansStore: subscription lifecycle, sorting, optimistic updates
- UIState: view routing, selection sync, URL persistence
- dragOrder: fractional indexing computation, reparent logic, edge cases
- AgentChatStore: message handling, subscription management

## Affected Files

- frontend/src/lib/stores/*.test.ts (new)
- frontend/src/lib/stores/dragOrder.test.ts (new)
- frontend/src/lib/uiState.test.ts (new)



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
