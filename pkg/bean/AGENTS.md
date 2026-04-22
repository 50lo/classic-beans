# pkg/bean

Scope: the canonical bean data model and pure helpers. This package owns the markdown/front-matter format, filename conventions, tag validation, ordering keys, and sort behavior.

## Entry Points

- `Parse` and `(*Bean).Render` define the on-disk markdown contract.
- `(*Bean).ETag` hashes rendered content, so any render change affects optimistic concurrency.
- `NewID`, `ParseFilename`, `BuildFilename`, and `Slugify` govern file naming.
- `OrderBetween` provides fractional order keys for manual ordering.
- `SortByStatusPriorityAndType` is the shared presentation sort used by CLI output.

## Contracts

- Preserve round-trip behavior between `Parse` and `Render` when changing front matter or body handling.
- File names are the source of truth for `ID` and optional `Slug`; keep the supported formats backward compatible.
- Tags are normalized to lowercase and must still pass `ValidateTag`; do not broaden the allowed alphabet casually.
- Empty priority sorts as `normal`; manual `Order` keys sort ahead of unordered beans.
- `MarshalJSON` exposes a computed `etag` field and should stay in sync with `Render`/`ETag`.

## Patterns

- Prefer small helpers in this package over reimplementing markdown or filename logic elsewhere.
- Keep body editing helpers strict: `ReplaceOnce` expects one unique match, and `AppendWithSeparator` preserves a single blank-line separator.
- Keep ordering helpers lexicographic and stable; callers rely on keys being comparable as plain strings.

## Pitfalls

- Do not change the filename grammar without preserving `ParseFilename` compatibility for old beans.
- Do not let parse/render tweaks silently change `ETag` unless you want the concurrency token to move.
- Avoid adding package dependencies here; this layer should remain reusable and low-level.
