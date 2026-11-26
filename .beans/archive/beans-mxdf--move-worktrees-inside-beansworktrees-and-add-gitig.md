---
# beans-mxdf
title: Move worktrees inside .beans/worktrees/ and add .gitignore
status: scrapped
type: task
priority: normal
created_at: 2026-03-10T09:33:09Z
updated_at: 2026-11-26T10:52:05Z
order: zzzzzw
---

## Summary of Changes

- Changed worktree path from sibling directories (`<repo>-<beanID>`) to `.beans/worktrees/<beanID>`
- Added `.gitignore` creation to `beans init` that excludes `worktrees/` and `conversations/`
- Updated `NewManager` to accept both `repoRoot` and `beansDir` parameters
- Added `.gitignore` to this project's own `.beans/` directory



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
