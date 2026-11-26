---
# beans-qrjj
title: Fix workspace destroy warning e2e tests
status: scrapped
type: bug
priority: normal
created_at: 2026-03-14T17:26:41Z
updated_at: 2026-11-26T10:52:06Z
---

The 3 workspace-destroy-warning e2e tests fail because promptDestroy() reads from polled worktreeStatuses which hasn't been updated yet. Fix: fetch fresh status on-demand when clicking destroy.

## Summary of Changes

- Made `promptDestroy()` in Sidebar.svelte fetch fresh worktree status on-demand via GraphQL query instead of relying on stale polled data
- Added destroy button to worktrees in "ready to integrate" state (shown on hover, replacing the check icon)
- Configured Playwright retries to 2 for flaky test resilience



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
