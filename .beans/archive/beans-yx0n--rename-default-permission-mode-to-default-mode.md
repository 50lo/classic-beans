---
# beans-yx0n
title: Rename default_permission_mode to default_mode
status: scrapped
type: task
priority: normal
created_at: 2026-03-11T12:24:53Z
updated_at: 2026-11-26T10:52:05Z
order: zzzzy
---

Rename the `default_permission_mode` config key/field to `default_mode` since the permissions system was removed and it now just controls Plan vs Act mode.

## Files to update
- `pkg/config/config.go` — struct tag + save logic
- `pkg/config/config_test.go` — test assertions
- `.beans.yml` — project config

## Tasks
- [x] Rename Go struct tag and save key
- [x] Update tests
- [x] Update .beans.yml
- [x] Run tests

## Summary of Changes

Renamed `default_permission_mode` to `default_mode` across the codebase:
- `pkg/config/config.go`: struct field `DefaultPermissionMode` → `DefaultMode`, YAML tag, save key, and getter method `GetDefaultPermissionMode()` → `GetDefaultMode()`
- `pkg/config/config_test.go`: all test references updated
- `internal/agent/manager.go`: type `DefaultPermissionMode` → `DefaultMode`, struct field, constructor param
- `internal/commands/serve.go`: call site updated
- `.beans.yml`: config key updated



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
