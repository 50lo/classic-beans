---
# beans-edz5
title: Use default text size for plans in PendingInteraction
status: scrapped
type: bug
priority: normal
created_at: 2026-03-17T10:15:55Z
updated_at: 2026-11-26T10:52:06Z
---

The plan mode EXIT_PLAN section uses text-xs everywhere, making it hard to read. Should use default text size per the project's styling rules.

## Summary of Changes

Removed text-xs from the EXIT_PLAN section in PendingInteraction.svelte so plans render at default text size. Also increased the max-height of the plan content container from max-h-48 to max-h-96.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
