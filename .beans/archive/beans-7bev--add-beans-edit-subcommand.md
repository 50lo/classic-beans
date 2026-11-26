---
# beans-7bev
title: Add beans edit subcommand
status: scrapped
type: feature
priority: normal
created_at: 2026-05-05T07:06:18Z
updated_at: 2026-11-26T10:52:06Z
---

Add a `beans edit <bean-id>` subcommand that opens the bean's markdown file in the user's editor. Editor resolution: $EDITOR > $VISUAL > nano.

## Summary of Changes

- Added `internal/commands/edit.go` with the `beans edit <id>` subcommand
- Editor resolution: $EDITOR > $VISUAL > `nano` fallback
- Looks up bean via `core.Get()` (supports prefix normalization), opens file at `core.FullPath()` in editor
- Registered in `internal/commands/register.go`
- Tests in `internal/commands/edit_test.go` covering editor resolution and argument validation
- Refs: beans-7bev



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
