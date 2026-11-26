---
# beans-rren
title: Avoid disk I/O while holding write lock in Core.Update()
status: scrapped
type: task
priority: normal
created_at: 2026-03-13T17:49:05Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-oyic
---

In core.go Update(), the ETag calculation reads from disk while holding the write lock. If the filesystem is slow, this blocks all other goroutines. Consider reading the file before acquiring the lock, or caching the on-disk ETag.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
