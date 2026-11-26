---
# beans-4mua
title: Fix suppressed a11y warnings in frontend components
status: scrapped
type: bug
priority: normal
created_at: 2026-03-13T17:49:14Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-oyic
---

RenderedMarkdown.svelte, AgentMessages.svelte, and PlanningView.svelte use svelte-ignore to suppress accessibility warnings. Project rules say never suppress a11y warnings. Fix with proper keyboard handlers or semantic elements (e.g. use <a> tags in rendered markdown, use <button> instead of div with role=button).



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
