---
# beans-9nfz
title: Worktree error logging and UI feedback
status: scrapped
type: task
priority: normal
created_at: 2026-03-09T11:20:07Z
updated_at: 2026-11-26T10:52:04Z
order: zzy
---

Add server-side logging for worktree create/destroy operations, and show error messages in the web UI when worktree operations fail.

## Summary of Changes

- Added `log.Printf` calls to `internal/worktree/worktree.go` for all worktree create/remove operations, logging both success and failure cases (including the "path already exists" early exit)
- Fixed `Create` to reuse an existing branch when it already exists (e.g. from a previously removed worktree), instead of failing with "branch already exists"
- Added error display in `BeanDetail.svelte` — when a worktree create or remove operation fails, a dismissible error banner appears above the worktree section showing the error message from the server
- Added `TestCreateReusesExistingBranch` test covering the create-remove-recreate scenario



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
