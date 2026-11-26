---
# beans-541b
title: Consolidate duplicate ETag error types
status: scrapped
type: task
priority: normal
created_at: 2026-03-13T17:49:02Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-oyic
---

ETagMismatchError and ETagRequiredError are defined in both internal/graph/resolver.go and pkg/beancore/core.go. Consolidate to one location (beancore, since it's the domain layer).



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
