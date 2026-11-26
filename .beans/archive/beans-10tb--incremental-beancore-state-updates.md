---
# beans-10tb
title: Incremental beancore state updates
status: scrapped
type: feature
priority: normal
created_at: 2025-12-20T09:50:14Z
updated_at: 2026-11-26T10:52:04Z
order: k
---

Update beancore state incrementally instead of full reload on file changes.

## Changes
- Accumulate fsnotify events during debounce window
- Process only affected files instead of reloading all beans
- Update search index incrementally per-bean
- Handle edge cases: rapid updates, create+delete, invalid files

## Files
- internal/beancore/watcher.go
- internal/beancore/core_test.go



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
