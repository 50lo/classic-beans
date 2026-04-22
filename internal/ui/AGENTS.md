# `internal/ui`

Shared rendering and tree helpers used by CLI output paths.

Scope:
- Builds ancestor-inclusive bean trees from filtered matches.
- Flattens and renders tree/list rows with width-aware column logic.
- Carries formatting concerns that are shared between list and detail views.

Behavior to preserve:
- `BuildTree` adds ancestors for context, not just the matched beans.
- Rendering assumes exact width accounting. Small changes here can cascade into broken list alignment.
- Tree nodes carry both match state and implicit status so the caller can dim context rows without losing meaning.

Use this package when:
- You need consistent tree/list presentation.
- You need a JSON-safe tree representation for non-interactive output.
