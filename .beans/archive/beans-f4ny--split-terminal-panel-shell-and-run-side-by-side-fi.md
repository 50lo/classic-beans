---
# beans-f4ny
title: 'Split terminal panel: Shell and Run side-by-side, fix run WebSocket'
status: scrapped
type: bug
priority: normal
created_at: 2026-03-18T10:10:09Z
updated_at: 2026-11-26T10:52:06Z
---

Two issues:
1. Terminal panel should show Shell and Run as side-by-side panels (horizontal split) instead of tabs
2. Run doesn't work because resolveTerminalWorkDir doesn't strip __run suffix from session IDs, so WebSocket connections for run sessions fail immediately

## Tasks
- [x] Fix resolveTerminalWorkDir to strip __run suffix
- [x] Replace Shell/Run tabs with horizontal split layout
- [x] Update tests (frontend build clean, Go vet clean)



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
