---
# beans-xf78
title: Implement GraphQL subscriptions for live updates
status: scrapped
type: task
priority: normal
created_at: 2025-12-18T16:45:17Z
updated_at: 2026-11-26T10:52:04Z
order: zz
parent: beans-lbjp
---

Add WebSocket support for GraphQL subscriptions to enable real-time updates.

## Tasks

- [ ] Add subscription type to GraphQL schema (beanCreated, beanUpdated, beanDeleted)
- [ ] Implement file watcher on .beans/ directory using fsnotify
- [ ] Set up WebSocket transport using gorilla/websocket
- [ ] Integrate with gqlgen subscription handler
- [ ] Broadcast file change events to connected subscribers
- [ ] Handle WebSocket connection lifecycle (connect, disconnect, keepalive)
- [ ] Write tests for subscription functionality

## Schema Addition

type Subscription {
  beanChanged: BeanChangeEvent!
}

type BeanChangeEvent {
  type: ChangeType!
  bean: Bean
  beanId: ID!
}

enum ChangeType {
  CREATED
  UPDATED
  DELETED
}

## Notes

- Use graphql-ws protocol (modern standard)
- Debounce rapid file changes to avoid flooding clients



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
