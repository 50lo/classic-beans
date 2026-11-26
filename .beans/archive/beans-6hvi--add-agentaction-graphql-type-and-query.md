---
# beans-6hvi
title: Add AgentAction GraphQL type and query
status: scrapped
type: task
priority: normal
created_at: 2026-03-11T19:11:46Z
updated_at: 2026-11-26T10:52:07Z
parent: beans-x7bd
---

Define the backend types and GraphQL schema for agent actions.

## Tasks
- [x] Add `AgentAction` type to GraphQL schema: `{ id: ID!, label: String!, description: String }`
- [x] Add query field `agentActions(beanId: ID!): [AgentAction!]!`
- [x] Implement resolver that returns available actions (hardcode Commit and Review for now)
- [x] Run codegen (`mise codegen`)
- [x] Add tests for the resolver

## Summary of Changes

Added AgentAction GraphQL type, agentActions query, and consolidated action registry in agent_helpers.go. Merged into main repo alongside beans-tk4x.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
