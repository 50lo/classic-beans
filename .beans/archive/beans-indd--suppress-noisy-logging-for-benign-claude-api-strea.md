---
# beans-indd
title: Suppress noisy logging for benign Claude API stream events
status: scrapped
type: bug
priority: normal
created_at: 2026-03-10T14:47:39Z
updated_at: 2026-11-26T10:52:05Z
order: zzzy
---

parseInnerEvent only handles content_block_delta and content_block_start. Other benign lifecycle events (content_block_stop, message_delta, message_stop, message_start) fall through to eventUnknown and get logged as 'unhandled event', spamming the logs.

## Summary of Changes

- Added `eventIgnored` event type for recognized-but-not-actionable streaming events
- Handle `content_block_stop`, `message_start`, `message_delta`, `message_stop` inner events as ignored
- Handle top-level `user` events (tool results) as ignored
- Fixed pre-existing `NewManager` call in manager_test.go
- Added test cases for all newly handled event types



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
