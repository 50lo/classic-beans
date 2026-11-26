---
# beans-uts2
title: Fix frontend svelte-check warnings and errors
status: scrapped
type: bug
priority: normal
created_at: 2026-03-18T14:44:40Z
updated_at: 2026-11-26T10:52:06Z
---

svelte-check reports 3 errors and 1 warning:
- WorkspaceView.svelte: worktreeId used before declaration
- +page.svelte: type mismatch string|undefined vs string
- Sidebar.svelte: a11y missing label on button

## Summary of Changes

Fixed 3 errors and 1 warning reported by svelte-check:

- **WorkspaceView.svelte**: Moved props declaration above derived state that references `worktreeId`, fixing "used before declaration" errors
- **+page.svelte**: Added non-null assertion to `page.params.beanId` since this param always exists on the `[beanId]` route
- **Sidebar.svelte**: Added `aria-label` to worktree bean buttons whose text content is set via a Svelte action



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
