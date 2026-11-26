---
# beans-p2pw
title: Fix navigator.platform access at module scope in FilterInput
status: scrapped
type: bug
priority: normal
created_at: 2026-03-13T17:49:18Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-oyic
---

FilterInput.svelte accesses navigator.platform at module scope, which will break during SSR/build. Needs a browser guard from $app/environment.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
