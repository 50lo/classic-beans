# pkg

Scope: everything under `pkg/`. Keep this layer focused on reusable package contracts; avoid leaking UI-specific or command-specific behavior into lower packages.

## Map

- `pkg/bean/`: canonical bean model, markdown/front-matter parsing/rendering, IDs, tags, ordering, and sorting.
- `pkg/beancore/`: thread-safe in-memory store plus filesystem persistence and search indexing.
- `pkg/config/`: project config loading/saving, upward config discovery, and path resolution.
- `pkg/safepath/`: path traversal guards for user-influenced path components.

## Shared Rules

- Keep behavior contractual, not mechanical. Document and preserve invariants that callers depend on.
- Put shared facts at the least common ancestor. Child nodes should only add package-specific constraints.
- Preserve backward compatibility for on-disk formats, identifiers, and path conventions unless the migration is explicit.

## Dependencies

- `bean` is the lowest-level domain package. It should stay free of `beancore` and `config` dependencies.
- `beancore` depends on `bean`, `config`, and `internal/search`; it is the authoritative runtime store.

## Downlinks

- [bean](./bean/AGENTS.md)
- [beancore](./beancore/AGENTS.md)
