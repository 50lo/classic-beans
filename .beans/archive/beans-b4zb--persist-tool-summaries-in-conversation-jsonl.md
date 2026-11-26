---
# beans-b4zb
title: Persist tool summaries in conversation JSONL
status: scrapped
type: bug
priority: normal
created_at: 2026-03-09T21:07:01Z
updated_at: 2026-11-26T10:52:05Z
order: zzzV
---

Tool messages are persisted to JSONL immediately with just the tool name (e.g. 'Bash'), before the summary is extracted from tool input deltas. After server restart, the agent chat shows bare tool names instead of rich descriptions like 'Bash: Build beans binary'. Fix by deferring tool message persistence until the summary is available.

## Summary of Changes

Deferred tool message persistence in `internal/agent/claude.go` so that tool summaries (e.g. "Bash: Build beans binary") are included in the JSONL log instead of just bare tool names ("Bash").

- Added `pendingToolPersist` flag and `flushToolMsg()` closure in `readOutput`
- Tool messages are no longer persisted immediately on `eventToolUse`
- Instead, they're flushed (with accumulated summary) when the next significant event arrives
- Added `TestToolMessagePersistsWithSummary` test



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
