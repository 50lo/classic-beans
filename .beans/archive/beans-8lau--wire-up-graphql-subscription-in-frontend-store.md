---
# beans-8lau
title: Wire up GraphQL subscription in frontend store
status: scrapped
type: feature
priority: normal
created_at: 2025-12-20T10:04:15Z
updated_at: 2026-11-26T10:52:04Z
order: z
---

Add real-time updates to BeansStore via GraphQL subscription.

## Changes
- Subscribe to beanChanged events on load
- Handle CREATED/UPDATED/DELETED events
- Update store state reactively

## Files
- frontend/src/lib/beans.svelte.ts



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
