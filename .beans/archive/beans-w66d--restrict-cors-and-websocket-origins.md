---
# beans-w66d
title: Restrict CORS and WebSocket origins
status: scrapped
type: task
priority: high
created_at: 2026-03-09T17:01:38Z
updated_at: 2026-11-26T10:52:04Z
order: zz
parent: beans-oe8n
---

The server currently sets Access-Control-Allow-Origin: * (serve.go line 65) and the WebSocket upgrader's CheckOrigin always returns true (line 97). This allows any website to make authenticated requests to the beans API and establish WebSocket subscriptions. Fix: replace the wildcard CORS origin with a configurable allowlist defaulting to localhost origins (http://localhost:*, http://127.0.0.1:*). The WebSocket CheckOrigin should validate the Origin header against the same allowlist. Consider adding a --cors-origin flag to beans serve for custom origins. Also review Access-Control-Allow-Headers and Access-Control-Allow-Methods to be specific rather than wildcard.

## Summary of Changes

- Added `internal/cors` package with origin checker supporting exact matches and port wildcards
- Replaced `Access-Control-Allow-Origin: *` with per-origin validation defaulting to localhost
- WebSocket `CheckOrigin` now validates against the same allowlist
- Added `--cors-origin` CLI flag and `server.cors_origins` config option
- Narrowed `Access-Control-Allow-Headers` to `Content-Type` only
- Added `Vary: Origin` header for correct caching behavior



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
