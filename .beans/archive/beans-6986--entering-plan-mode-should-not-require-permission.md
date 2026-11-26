---
# beans-6986
title: Entering plan mode should not require permission
status: scrapped
type: bug
priority: normal
created_at: 2026-03-11T13:40:51Z
updated_at: 2026-11-26T10:52:05Z
order: zzzV
---

When a user or agent enters plan mode, it triggers a permission prompt. Entering plan mode is a read-only/planning action and should not require approval — it's a mode switch, not a mutation.

## Summary of Changes

EnterPlanMode is now auto-approved in the backend via `autoApproveModeSwitch`. The mode is toggled and the process is immediately respawned with `--resume` — no pending interaction is set and no user prompt is shown.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
