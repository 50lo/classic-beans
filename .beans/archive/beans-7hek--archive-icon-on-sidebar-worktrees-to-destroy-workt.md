---
# beans-7hek
title: Archive icon on sidebar worktrees to destroy worktree
status: scrapped
type: feature
priority: normal
created_at: 2026-03-11T22:13:21Z
updated_at: 2026-11-26T10:52:07Z
---

Add an archive/destroy icon to worktree items in the sidebar. When clicked, it should confirm and then destroy the worktree using the existing removeWorktree/stopWork mutation.

## Summary of Changes

Added an archive icon button to each worktree item in the sidebar. The icon appears on hover, and clicking it opens a confirmation modal before destroying the worktree via `removeWorktree`. If the destroyed worktree was the active view, navigation redirects to Planning.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
