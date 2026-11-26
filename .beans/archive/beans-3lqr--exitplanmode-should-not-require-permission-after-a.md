---
# beans-3lqr
title: ExitPlanMode should not require permission after accepting plan
status: scrapped
type: bug
priority: normal
created_at: 2026-03-11T13:40:48Z
updated_at: 2026-11-26T10:52:05Z
order: zzzk
---

When the user accepts a plan in plan mode, the ExitPlanMode tool fires and asks for permission. This is redundant and confusing — accepting the plan IS the user's explicit intent to exit plan mode. ExitPlanMode should be auto-approved (or skipped entirely) when triggered by plan acceptance.

## Summary of Changes

ExitPlanMode is now auto-approved in the backend via `autoApproveModeSwitch`. The mode is toggled and the process is immediately respawned — no permission prompt after plan acceptance.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
