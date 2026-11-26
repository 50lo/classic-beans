---
# beans-r7oc
title: Improve subscription backpressure handling
status: scrapped
type: task
priority: low
created_at: 2026-03-13T17:49:10Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-oyic
---

Non-blocking fanOut silently drops events when a subscriber's channel is full. For a UI that depends on subscription accuracy, this could cause state drift. Consider larger buffers, a catch-up mechanism, or full-state resync on overflow.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
