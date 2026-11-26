---
# beans-eo0m
title: Filter inappropriate words from generated workspace names!!
status: scrapped
type: feature
priority: normal
created_at: 2026-03-13T15:19:14Z
updated_at: 2026-11-26T10:52:04Z
order: c
---

Add a blocklist to the workspace name generator to prevent inappropriate adjectives (e.g. 'sexual', 'violent', 'naked') from appearing in randomly generated workspace/worktree names. Pre-filters the unique-names-generator dictionaries at module load time rather than using a retry loop.

## Tasks

- [x] Add blocklist and pre-filter adjectives/animals dictionaries
- [x] Update generateWorkspaceName to use filtered dictionaries
- [x] Write tests for filtered dictionaries and name generation
- [x] Code review

## Summary of Changes

- Added a blocklist of inappropriate adjectives (aggressive, sexual, violent, naked, etc.) to the workspace name generator
- Pre-filtered the `unique-names-generator` dictionaries at module load time
- Added vitest tests verifying blocked words are excluded and dictionary size remains reasonable



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
