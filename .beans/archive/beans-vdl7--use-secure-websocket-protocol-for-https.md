---
# beans-vdl7
title: Use secure WebSocket protocol for HTTPS
status: scrapped
type: task
priority: low
created_at: 2026-03-09T17:02:08Z
updated_at: 2026-11-26T10:52:04Z
order: zzz
parent: beans-oe8n
---

graphqlClient.ts line 9 hardcodes ws:// for WebSocket connections. If beans is ever served behind HTTPS (via reverse proxy), this will fail or expose data in plaintext. Fix: detect the page protocol and use wss:// when on HTTPS. Something like: const wsProtocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'. Low priority since beans currently runs localhost HTTP only, but trivial to fix and future-proofs the code.

## Summary of Changes

Updated `frontend/src/lib/graphqlClient.ts` to detect the page protocol (`window.location.protocol`) and use `wss://` for WebSocket connections when on HTTPS, falling back to `ws://` for HTTP. This future-proofs the WebSocket connection for deployments behind HTTPS reverse proxies.



## Reasons for Scrapping

Out of scope for the Classic Beans backlog, which now tracks the Workflow CLI commands epic (`beans-mmyp`) and its descendants. This item was previously marked completed.
