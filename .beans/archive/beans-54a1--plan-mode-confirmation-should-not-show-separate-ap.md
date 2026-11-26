---
# beans-54a1
title: Plan mode confirmation should not show separate Approve + Act button
status: scrapped
type: bug
priority: normal
created_at: 2026-03-11T13:40:49Z
updated_at: 2026-11-26T10:52:05Z
order: zzzs
---

The plan mode confirmation dialog shows both 'Approve' and 'Approve + Act' buttons, but they do the same thing — there's no distinction between approving a plan and approving + acting on it, since accepting a plan always leads to action. Remove the redundant button.

## Summary of Changes

Removed the entire plan mode approval UI block from AgentChat.svelte (including the 'Approve + Act' button). Since both EnterPlanMode and ExitPlanMode are now auto-approved on the backend, the frontend never receives these interaction types, making the UI dead code.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
