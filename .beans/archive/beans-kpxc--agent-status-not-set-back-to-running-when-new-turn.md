---
# beans-kpxc
title: Agent status not set back to Running when new turn starts within same process
status: scrapped
type: bug
priority: normal
created_at: 2026-03-12T15:22:19Z
updated_at: 2026-11-26T10:52:05Z
---

When Claude Code starts a new turn within the same process (e.g. after a Stop hook), readOutput keeps processing events but never sets session.Status back to StatusRunning. The UI shows the agent as idle while it's clearly still working.

## Summary of Changes

Added `ensureRunning()` helper in `readOutput()` that transitions the session status from Idle back to Running when a new turn starts within the same Claude Code process. Called at the start of `eventAssistantMessage`, `eventToolUse`, `eventNewTextBlock`, and `eventTextDelta` handlers. Added test `TestReadOutputMultiTurnResetsStatus` that verifies the Idle → Running → Idle transition across two turns in a single process.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
