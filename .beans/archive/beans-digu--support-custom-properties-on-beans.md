---
# beans-digu
title: Support custom properties on beans
status: scrapped
type: feature
priority: normal
created_at: 2025-12-13T00:52:24Z
updated_at: 2026-11-26T10:52:02Z
order: s
---

Allow users to attach custom key-value properties to beans. Custom properties should live under a dedicated `properties` key in the frontmatter to keep them separate from built-in fields.

## Example

```yaml
---
title: Fix authentication bug
status: in-progress
type: bug
properties:
  github_issue: "#135"
  author: alice@bob.com
  estimate: 3
  reviewed: true
---
```

## Considerations

- Properties can be any YAML-supported type (string, number, boolean, etc.)
- Should be exposed via GraphQL (probably as JSON scalar or key-value pairs)
- Could support filtering/searching by property values in the future
- CLI: `beans update <id> --set key=value` or similar



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
