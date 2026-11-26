---
# beans-scm4
title: Workspace remembers currently open bean
status: scrapped
type: feature
priority: normal
created_at: 2026-03-13T19:07:31Z
updated_at: 2026-11-26T10:52:07Z
---

Each workspace should remember which bean it has open. When switching between workspaces, the previously selected bean should be restored. Currently selectedBeanId is a single global value that gets lost on workspace switch.

## Summary of Changes

Changed `selectedBeanId` in `UIState` from a single global value to a per-view map (`selectedBeanByView`), keyed by `activeView` ('planning' or worktree ID). When switching between workspaces, the previously selected bean is automatically restored. The URL `?bean=` query param is synced to reflect the current view's selection.

**Files changed:**
- `frontend/src/lib/uiState.svelte.ts` — replaced single `selectedBeanId` state with `selectedBeanByView` record and getter/setter that reads from the current active view
- `frontend/src/routes/+layout.svelte` — call `syncFromUrl` before setting initial `selectedBeanId` so it's stored under the correct view key



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
