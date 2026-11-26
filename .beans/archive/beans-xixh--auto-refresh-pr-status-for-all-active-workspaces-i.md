---
# beans-xixh
title: Auto-refresh PR status for all active workspaces in sidebar
status: scrapped
type: bug
priority: normal
created_at: 2026-03-17T19:19:24Z
updated_at: 2026-11-26T10:52:05Z
---

PR status icons on sidebar workspace cards don't auto-refresh. The WorktreesChanged subscription only re-emits PR data on worktree filesystem events, so sidebar icons go stale. Need periodic PR status polling for all active worktrees.

## Summary of Changes

Added a 30-second periodic ticker to the `WorktreesChanged` subscription resolver in `schema.resolvers.go`. On each tick, the resolver re-fetches the worktree list and populates PR data asynchronously, then re-emits the updated list to all subscribers. This ensures sidebar PR status icons stay current even when no worktree filesystem events occur.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
