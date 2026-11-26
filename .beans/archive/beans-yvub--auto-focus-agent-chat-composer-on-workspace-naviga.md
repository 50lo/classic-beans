---
# beans-yvub
title: Auto-focus agent chat composer on workspace navigation
status: scrapped
type: feature
priority: normal
created_at: 2026-03-12T15:13:12Z
updated_at: 2026-11-26T10:52:07Z
---

After creating a new workspace and opening it, the agent chat input composer should be focused automatically.

## Summary of Changes

Added auto-focus to the agent chat composer textarea when navigating to a workspace. A `$effect` tracks `beanId` changes and calls `focus()` on the textarea element. Also added an e2e assertion to verify the behavior.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
