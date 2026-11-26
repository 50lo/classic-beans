---
# beans-r66y
title: Deduplicate ETagMismatchError type
status: scrapped
type: task
priority: normal
created_at: 2026-03-14T15:06:45Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-5txp
---

ETagMismatchError is defined identically in both pkg/beancore/ and internal/graph/. Should exist in one canonical location (pkg/beancore/ makes sense since that's where ETag validation happens) and be imported elsewhere.

## Affected Files

- pkg/beancore/core.go
- internal/graph/schema.resolvers.go



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
