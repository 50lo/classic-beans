---
# beans-fydu
title: Replace findMainBeanFile directory scan with in-memory lookup
status: scrapped
type: task
priority: normal
created_at: 2026-03-13T17:49:07Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-oyic
---

findMainBeanFile() does os.ReadDir() and linearly scans filenames on every worktree change. Core already has a beans map indexed by ID — use it instead of hitting the filesystem.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
