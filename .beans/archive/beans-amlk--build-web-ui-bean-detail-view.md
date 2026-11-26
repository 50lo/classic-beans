---
# beans-amlk
title: 'Build web UI: Bean detail view'
status: scrapped
type: task
priority: normal
created_at: 2025-12-18T16:45:35Z
updated_at: 2026-11-26T10:52:04Z
order: zV
parent: beans-lbjp
---

Implement the bean detail/edit interface.

## Tasks

- [ ] Create BeanDetail component
- [ ] Display full bean information (title, status, type, priority, tags, body)
- [ ] Render markdown body with syntax highlighting for code blocks
- [ ] Show parent and children beans with navigation
- [ ] Show blocking/blockedBy relationships
- [ ] Implement inline editing for all fields
- [ ] Add status change quick actions
- [ ] Subscribe to live updates for the current bean

## Design Notes

- Could be a slide-over panel or separate route
- Markdown editor with preview for body editing



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
