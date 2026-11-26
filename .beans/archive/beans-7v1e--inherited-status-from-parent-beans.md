---
# beans-7v1e
title: Inherited status from parent beans
status: scrapped
type: feature
priority: normal
created_at: 2026-03-08T22:09:06Z
updated_at: 2026-11-26T10:52:05Z
order: zzzk
---

Add inherited/effective status that propagates terminal statuses (scrapped, completed) from parent beans to children transitively. Children keep their explicit status unchanged, but expose an inherited status for display and filtering purposes.

## Summary of Changes

Implemented inherited status: Core.InheritedStatus(), GraphQL fields, ready filter exclusion, list/show/TUI display with ↑status annotation



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
