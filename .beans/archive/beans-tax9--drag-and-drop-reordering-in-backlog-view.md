---
# beans-tax9
title: Drag and drop reordering in backlog view
status: scrapped
type: feature
priority: normal
created_at: 2026-03-09T17:10:51Z
updated_at: 2026-11-26T10:52:04Z
order: zs
---

Add drag and drop support to the backlog view so users can manually reorder beans by dragging them. This should update the bean's order field to persist the new position. Note: there is already in-progress work on this — see the unstaged files backlogDrag.svelte.ts and dragOrder.ts in the working tree.

## Summary of Changes

Implemented in commit 0380e71. See beans-424n for full details.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
