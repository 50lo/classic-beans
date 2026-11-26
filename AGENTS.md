# Beans Agent Guide

Classic Beans is a CLI issue tracker. Issues are Markdown files under `.beans/`, managed via the `cbeans` CLI.

## Intent Layer

This repository uses `AGENTS.md` files as sparse coding-agent context. Read this root file first, then the nearest child node for the area you are editing:

- [pkg](./pkg/AGENTS.md): reusable domain, storage, config, and path-safety packages.
- [internal](./internal/AGENTS.md): CLI commands, search, output formatting, and tree/list rendering.

Keep new guidance at the least common ancestor that covers it. Avoid duplicating rules between parent and child nodes.

## Cross-Cutting Rules

- Always write or update tests for behavior changes.
- Preserve backward compatibility for `.beans/` markdown, `.beans.yml`, IDs, paths, and on-disk formats unless a migration is explicit.

## Commands

- Build: `mise build`
- Run all tests: `mise test`
- Run Go tests directly: `go test ./...`
- Run the CLI through mise: `mise cbeans`

For read-only CLI checks, this repo's own `.beans/` data is fine. For write tests, create a separate temporary project or use `--beans-path`.

## Commits and PRs

- Use conventional commit messages.
- Include relevant bean IDs in commit bodies when available, e.g. `Refs: beans-abcd`.
- Mark breaking changes with `!`.
- On PR branches, keep commits focused and update the PR description with the actual changes.
