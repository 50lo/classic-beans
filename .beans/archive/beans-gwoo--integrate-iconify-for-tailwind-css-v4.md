---
# beans-gwoo
title: Integrate Iconify for Tailwind CSS v4
status: scrapped
type: task
priority: normal
created_at: 2026-03-10T15:14:04Z
updated_at: 2026-11-26T10:52:05Z
order: zzzc
---

Integrate the @iconify/tailwind4 plugin into the frontend for icon support.

## Steps

- [x] Install `@iconify/tailwind4` and `@iconify/json` as dev dependencies
- [x] Add `@plugin "@iconify/tailwind4"` directive to the CSS file
- [ ] Verify icons render correctly (e.g. `<span class="icon-[mdi-light--home]"></span>`)
- [ ] Consider configuring prefix/scale options and clean selectors if needed

## Reference

- Docs: https://iconify.design/docs/usage/css/tailwind/tailwind4/
- Usage syntax: `icon-[prefix--name]` (e.g. `icon-[mdi-light--home]`)
- Clean selector alternative requires specifying icon set prefixes in config



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
