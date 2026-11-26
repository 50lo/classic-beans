---
# beans-mx9n
title: Show progress animation on BeanCard when agent is running
status: scrapped
type: feature
priority: normal
created_at: 2026-03-10T15:03:10Z
updated_at: 2026-11-26T10:52:05Z
order: zzzzzV
---

Add a visual indicator (animated glow/pulse) on the BeanCard green corner when the agent is actively working in that bean's worktree. Requires a new global agent status subscription, a frontend store, and animation in BeanCard.

## Summary of Changes

Added a pulsing animation on the BeanCard green corner indicator when the agent is actively running in that bean's worktree. This involved:

- New `activeAgentStatuses` GraphQL subscription broadcasting running agent status globally
- Global pub/sub in the agent manager (`SubscribeGlobal`/`UnsubscribeGlobal`/`ListRunningSessions`)
- New `AgentStatusesStore` frontend store tracking running bean IDs
- CSS pulse animation on the worktree corner triangle in `BeanCard`



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
