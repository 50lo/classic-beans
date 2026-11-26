---
# beans-04yk
title: Implement GraphQL HTTP endpoint
status: scrapped
type: task
priority: normal
created_at: 2025-12-18T16:44:35Z
updated_at: 2026-11-26T10:52:04Z
order: zy
parent: beans-lbjp
---

Expose the existing GraphQL schema over HTTP.

## Tasks

- [ ] Add HTTP handler using gqlgen handler package
- [ ] Set up CORS configuration for local development
- [ ] Mount GraphQL endpoint at /graphql
- [ ] Add GraphQL Playground at /graphql (GET requests) for development/debugging
- [ ] Write tests for the HTTP handler

## Dependencies

- Existing internal/graph/ resolver infrastructure

## Notes

- Use github.com/99designs/gqlgen/graphql/handler
- Consider using chi or stdlib for routing



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
