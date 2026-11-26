---
# beans-l4ag
title: Auto-reload config file on changes and push updates to frontend
status: scrapped
type: feature
priority: normal
created_at: 2026-03-11T18:12:28Z
updated_at: 2026-11-26T10:52:02Z
order: zzzV
blocked_by:
    - beans-bbjk
---

Make beans-serve (or beanscore) watch the .beans.yml config file for changes and automatically push updated config (like agent actions) to the frontend via GraphQL subscriptions. Currently config is only read at startup.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
