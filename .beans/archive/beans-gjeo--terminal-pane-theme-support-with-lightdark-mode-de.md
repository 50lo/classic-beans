---
# beans-gjeo
title: Terminal pane theme support with light/dark mode detection
status: scrapped
type: task
priority: normal
created_at: 2026-03-12T19:23:23Z
updated_at: 2026-11-26T10:52:08Z
---

Add light and dark color themes to the xterm.js terminal pane, with automatic switching based on prefers-color-scheme. The terminal background syncs with the app's --color-surface CSS variable.

## Summary of Changes

- Added dark and light xterm.js ANSI color palettes
- Background color resolved from Tailwind's --color-surface CSS variable
- MediaQueryList listener switches theme dynamically when OS color scheme changes
- Proper cleanup of event listener in onDestroy
- Replaced hardcoded bg color with bg-surface Tailwind class



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
