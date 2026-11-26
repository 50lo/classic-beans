---
# beans-jmic
title: Fix integrate button leaving main worktree out of sync
status: scrapped
type: bug
priority: normal
created_at: 2026-03-14T15:18:17Z
updated_at: 2026-11-26T10:52:06Z
---

git update-ref moves the ref but does not update main's working directory or index, causing phantom uncommitted changes after integration. Fix: use git merge --ff-only from main's directory instead.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
