---
# beans-z4ry
title: 'Build web UI: Create/edit bean form'
status: scrapped
type: task
priority: normal
created_at: 2025-12-18T16:45:45Z
updated_at: 2026-11-26T10:52:04Z
order: zk
parent: beans-lbjp
---

Implement bean creation and editing forms.

## Tasks

- [x] Create BeanForm component for create/edit
- [x] Title input field
- [x] Type selector (dropdown with configured types)
- [x] Status selector
- [x] Priority selector
- [x] Tags input (comma-separated input)
- [x] Parent selector (dropdown with cycle prevention)
- [x] Blocking beans selector (deferred -- needs multi-select UI)
- [x] Markdown body editor (textarea, no preview yet)
- [x] Form validation
- [x] Submit via GraphQL mutation
- [x] Handle optimistic updates (via subscription)

## Design Notes

- Modal or slide-over for creation
- Could reuse for inline editing in detail view

## Summary of Changes

Implemented BeanForm component (create + edit) as a DaisyUI modal dialog:

- Title, type, status, priority fields
- Parent selector with cycle prevention
- Comma-separated tags input
- Markdown body textarea
- Form validation (title required)
- Create/update via GraphQL mutations
- Subscription handles reactive updates (no optimistic UI needed)
- Edit button in BeanDetail header
- "+ New Bean" button in the tab bar

Deferred: Blocking beans multi-select (would benefit from a proper multi-select component).



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
