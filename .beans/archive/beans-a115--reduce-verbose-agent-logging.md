---
# beans-a115
title: Reduce verbose agent logging
status: scrapped
type: task
priority: normal
created_at: 2026-03-11T22:07:58Z
updated_at: 2026-11-26T10:52:08Z
---

Server logs excessive agent message data. Key issues: unhandled events logged as raw JSON, all stderr from Claude Code logged, spawn args include session IDs. Fix by truncating/suppressing verbose output and using proper log levels.

## Summary of Changes

Reduced verbose agent logging in internal/agent/claude.go:
- Silenced stderr drain (Claude Code writes verbose progress info that overwhelmed logs)
- Unhandled events now log only the event type, not the full JSON payload
- Removed CLI args from process spawn log (session IDs, permission flags)



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
