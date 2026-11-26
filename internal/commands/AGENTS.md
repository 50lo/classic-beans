# `internal/commands`

CLI wiring for the `cbeans` command family.

What lives here:
- Root setup in `root.go`: shared persistent flags, config loading, and `beancore` initialization.
- Project commands: `init`, `prime`, `create`, `list`, `show`, `edit`, `update`, `delete`, `archive`, `check`, `roadmap`, and `version`.

Behavior to keep in mind:
- `PersistentPreRunE` skips core loading for `init`, `prime`, and `version`.
- `--beans-path` and `--config` override config discovery; `resolveBeansPath` rejects missing or non-directory overrides.
- `init` with an explicit beans path creates `.beans` only and intentionally skips writing `.beans.yml`.
- `prime` is quiet when no project config is discoverable; it is meant for agent priming, not user-facing output.

Common pitfalls:
- Do not move runtime initialization into commands that are supposed to stay cheap (`prime`, `version`).
- Keep command precedence consistent with config: flag > env > config/default.
