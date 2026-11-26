---
# beans-27wn
title: PR status shows stale check state after agent push
status: scrapped
type: bug
priority: normal
created_at: 2026-03-17T17:53:50Z
updated_at: 2026-11-26T10:52:06Z
---

When an agent pushes to a PR and the turn ends, the immediate re-fetch of PR status returns stale data (checks still show 'pass' from before the push). This causes the UI to show 'Merge PR' instead of 'Checks Running'. Fix: add a delayed re-fetch after agent goes busy→idle.

## Summary of Changes

Added a delayed re-fetch (5s) after the agent transitions from busy to idle in `AgentActionsStore.notifyAgentStatus`. The immediate fetch still runs for actions that don't depend on external CI, but the follow-up fetch catches GitHub check status that wasn't registered yet at turn-end. Timer is cleaned up properly on component unmount via `stopPolling()`.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
