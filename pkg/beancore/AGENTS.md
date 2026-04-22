# pkg/beancore

Scope: the authoritative runtime store for beans. This package owns disk I/O, in-memory state, and search indexing.

## Entry Points

- `New`, `Load`, `Close`, `All`, `Get`, `NormalizeID`, `Search`, `Create`, `Update`, and `Delete` are the main store operations.

## Contracts

- `Core` is thread-safe. Internal maps are protected by locks; preserve that boundary when changing helpers.
- `Load` rebuilds in-memory state from disk, reinitializes the search index if present, and skips dot-prefixed subtrees such as `.worktrees/`.
- Search is lazily initialized and backed by the in-memory index, not disk scans.
- `Get` and `NormalizeID` honor configured prefixes so short IDs can resolve to full IDs.

## Patterns

- Keep disk parsing in `loadBean`/`loadBeanFrom` and defaulting logic aligned.
- Preserve event fan-out semantics: slow consumers may drop batches instead of blocking the store.
- Be careful with search index updates on create/update/delete/load; stale indexing shows up as silent query drift.

## Pitfalls

- Do not make helpers traverse hidden directories; they are intentionally reserved for non-bean data.
- Avoid weakening ETag behavior. Concurrency checks depend on etags remaining accurate.
