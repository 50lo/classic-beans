---
# beans-1usq
title: Replace sidebar polling with workspaceStatuses subscription
status: scrapped
type: task
priority: normal
created_at: 2026-03-18T09:31:48Z
updated_at: 2026-11-26T10:52:08Z
---

Replace the 3s polling of MainChanges and WorktreeStatuses queries in Sidebar.svelte with a new workspaceStatuses GraphQL subscription. The backend will check git status every 10s and emit via WebSocket, eliminating HTTP request spam.

## Summary of Changes

- Added `WorkspaceStatus` type and `workspaceStatuses` subscription to GraphQL schema
- Implemented server-side subscription resolver that checks git status every 10s and only emits on change
- Replaced 3s HTTP polling in Sidebar.svelte with WebSocket subscription
- Updated `promptDestroy` to use `worktreeStore.getWorktreeStatus()` for on-demand fresh data
- Removed unused `MainChanges` and `WorktreeStatuses` queries from operations.graphql



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
