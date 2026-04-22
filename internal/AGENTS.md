# Internal

This tree holds the CLI glue for beans: command wiring, search, output formatting, and tree/list rendering.

Use the nearest child node first:
- `internal/commands/AGENTS.md` for Cobra entrypoints and CLI-specific behavior.
- `internal/ui/AGENTS.md` for shared tree/list rendering helpers.

Shared rules:
- Keep package docs short and operational. Capture invariants, not code listings.
- Treat `pkg/*` as the lower-level domain/config layer; `internal/*` wires it into CLI behavior.
- Avoid duplicating the same rule in child files. Put cross-cutting facts here and only package-specific gotchas below.
