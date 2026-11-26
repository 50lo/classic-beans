---
# beans-sa92
title: Remove automatic reordering of workspaces in sidebar
status: scrapped
type: task
priority: normal
created_at: 2026-03-18T10:10:19Z
updated_at: 2026-11-26T10:52:08Z
---

Keep workspaces in creation order (oldest first) instead of reordering by last activity

## Summary of Changes

Removed the LastActiveAt-based sort in `worktree.Manager.List()`, so workspaces now stay in git's native creation order (oldest first) instead of reordering by last activity. Updated the corresponding test.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
