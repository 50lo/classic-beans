---
# beans-ocss
title: Persist workspace port assignments across restarts
status: scrapped
type: task
priority: normal
created_at: 2026-03-21T09:29:08Z
updated_at: 2026-11-26T10:52:08Z
---

Add Port field to worktreeMeta so port assignments survive beans serve restarts. Also document worktree metadata pattern in CLAUDE.md.


## Summary of Changes

- Added `AllocateSpecific(workspaceID, port)` to `portalloc.Allocator` — restores a specific port for a workspace, falling back to normal allocation on conflict
- Added `Port` field to `worktreeMeta` struct, with `SavePort`/`GetPort` methods on `worktree.Manager`
- Updated `serve.go` startup to restore persisted ports from worktree metadata before falling back to fresh allocation
- Updated `CreateWorktree` resolver to persist allocated port to metadata
- Added tests for `AllocateSpecific` (happy path, idempotency, conflict, nextIndex advancement)
- Documented the worktree metadata file pattern in CLAUDE.md



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
