---
# beans-x38m
title: Sandboxed agent execution
status: scrapped
type: feature
priority: low
created_at: 2026-03-10T10:01:00Z
updated_at: 2026-11-26T10:52:03Z
order: zy
---

Investigate and implement lightweight sandboxing for agent processes running in worktrees. Docker is the most pragmatic starting point (cross-platform, familiar). Could also explore macOS sandbox-exec, Linux bwrap/unshare. Make it opt-in via config (e.g. agent.sandbox: docker). This becomes more important as we move towards yolo mode for all agent sessions.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
