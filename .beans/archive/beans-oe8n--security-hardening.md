---
# beans-oe8n
title: Security hardening
status: scrapped
type: epic
priority: normal
created_at: 2026-03-09T17:01:23Z
updated_at: 2026-11-26T10:52:02Z
order: zzk
---

Harden beans against common web security vulnerabilities identified in security review. Priority-ordered by impact and ease of fix. Covers XSS via unsanitized markdown, missing CORS/origin restrictions, no input validation on GraphQL mutations, path traversal risks, missing CSP headers, and no request/connection limits. Out of scope for now: authentication (separate effort for multi-user), rate limiting (low priority for local use), HTTPS/TLS (reverse proxy concern).



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants.
