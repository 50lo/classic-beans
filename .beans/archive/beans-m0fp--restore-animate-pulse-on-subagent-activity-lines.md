---
# beans-m0fp
title: Restore animate-pulse on subagent activity lines
status: scrapped
type: bug
priority: normal
created_at: 2026-03-17T13:55:03Z
updated_at: 2026-11-26T10:52:06Z
---

The subagent activity output lines in the agent chat no longer pulse. They should have animate-pulse while active.

## Summary of Changes

Added `animate-pulse` class to subagent activity lines in AgentMessages.svelte to restore the pulsing animation while subagents are working.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
