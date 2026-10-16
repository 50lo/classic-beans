---
# beans-5mpo
title: Remove Web-UI, TUI, GraphQL, worktrees, and forge
status: completed
type: task
priority: critical
created_at: 2026-04-22T22:09:11Z
updated_at: 2026-04-22T22:23:08Z
---

Simplify beans to CLI-only issue tracker. Remove Web-UI, TUI, GraphQL, worktree management, and forge integration.

## Todo
- [ ] Delete removable directories/packages
- [ ] Clean internal/commands/ (remove serve/tui/query/graphql/prime commands)
- [ ] Strip pkg/beancore/ (subscriptions, worktree merging)
- [ ] Update cmd/beans/main.go (single binary)
- [ ] Update mise.toml and .goreleaser.yaml
- [ ] Update AGENTS.md files
- [ ] go mod tidy, build, and test

## Summary of Changes

Removed Web UI (SvelteKit frontend), TUI (Bubble Tea), GraphQL API, worktree management, forge integration, and agent management. The project is now a single CLI binary () backed by  with full-text search.

**Deleted packages:** frontend, cmd/beans-serve, cmd/beans-tui, internal/web, internal/cors, internal/graph, internal/tui, internal/terminal, internal/agent, internal/portalloc, internal/worktree, internal/gitutil, pkg/beangraph, pkg/forge

**Rewritten commands:** create, list, show, update, delete, roadmap, prime - all now use beancore directly instead of GraphQL resolver layer.

**Stripped from beancore:** subscriptions, file watching, worktree bean merging, dirty bean save helpers (SaveDirty/SaveBean).

**Added to beancore:** Filter struct and FilteredBeans() method to replace GraphQL filtering.

**Updated:** mise.toml (single binary build), .goreleaser.yaml (single binary release), prompt.tmpl (no GraphQL), all AGENTS.md files.
