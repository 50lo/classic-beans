---
# beans-kpvs
title: Centralize ID normalization logic
status: scrapped
type: task
priority: normal
created_at: 2026-03-14T15:06:45Z
updated_at: 2026-11-26T10:52:03Z
parent: beans-5txp
---

NormalizeID() is called everywhere — resolvers, validators, helpers — with repeated patterns. A single RequireFullID(id string) (string, error) method on Core, used consistently, would reduce noise and potential bugs.

## Current Pattern (repeated ~20+ times)

fullID, ok := c.NormalizeID(id)
if !ok {
    return fmt.Errorf("bean not found: %s", id)
}

## Proposed Pattern

fullID, err := c.RequireFullID(id)
if err != nil {
    return err
}

## Affected Files

- pkg/beancore/core.go
- internal/graph/schema.resolvers.go
- internal/graph/helpers.go



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
