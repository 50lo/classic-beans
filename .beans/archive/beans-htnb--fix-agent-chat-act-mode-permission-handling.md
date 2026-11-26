---
# beans-htnb
title: Fix agent chat Act mode permission handling
status: scrapped
type: bug
priority: normal
created_at: 2026-03-10T11:57:44Z
updated_at: 2026-11-26T10:52:05Z
order: zzzs
---

Use acceptEdits permission mode and add Bash(beans:*) allowed tool for Act mode. Add logging for permission denials.

## Summary of Changes

- **`internal/agent/claude.go` — `buildClaudeArgs()`**: Act mode now uses `--permission-mode acceptEdits` and `--allowedTools Bash(beans:*)` so file operations and beans CLI work without permission prompts.
- **`internal/agent/claude.go` — `readOutput()`**: Permission denials now log individual tool names (`permission denied: tool=X`) instead of just a count.
- **`internal/agent/manager_test.go`**: Updated `TestBuildClaudeArgs_ActMode`, `TestBuildClaudeArgs_NoPlanMode`, and `TestBuildClaudeArgs_AllowedTools` to match the new behavior.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
