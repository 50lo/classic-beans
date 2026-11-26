---
# beans-fv19
title: Worktree agents should see main repo's beans
status: scrapped
type: feature
priority: normal
created_at: 2026-03-10T09:17:26Z
updated_at: 2026-11-26T10:52:05Z
order: zzzzzk
---

When beans CLI runs in a secondary git worktree, auto-detect and redirect .beans/ path to the main worktree's .beans/ directory. This ensures agents in worktrees can see uncommitted beans from the main repo.

## Summary of Changes

- Created `internal/gitutil/worktree.go` with `MainWorktreeRoot()` function that detects secondary git worktrees by comparing `--git-common-dir` and `--git-dir`
- Created `internal/gitutil/worktree_test.go` with tests for main worktree, secondary worktree, and non-git directory cases
- Modified `internal/commands/root.go` to add worktree redirect fallback in `resolveBeansPath()` — when no `.beans/` is found locally and no explicit override is set, it checks the main worktree's `.beans/` path
- `--beans-path` flag and `BEANS_PATH` env var always take precedence over worktree detection



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
