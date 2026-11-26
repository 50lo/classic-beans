---
# beans-wn9j
title: 'Fix panic: send on closed channel in WorktreesChanged subscription'
status: scrapped
type: bug
priority: normal
created_at: 2026-03-20T15:09:01Z
updated_at: 2026-11-26T10:52:06Z
---

populatePRsAsync goroutines can race with close(out) when context cancels, causing panic: send on closed channel

## Summary of Changes

Added a `sync.WaitGroup` to the `WorktreesChanged` subscription resolver to track in-flight `populatePRsAsync` goroutines. The main goroutine now waits for all async PR-fetch goroutines to complete before closing the output channel, preventing the "send on closed channel" panic.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
