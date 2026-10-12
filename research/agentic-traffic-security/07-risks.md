# 07 — Risk Register

Scored 1–5 for likelihood (L) and impact (I). Score = L × I.

| # | Risk | L | I | Score | Mitigation | Early warning |
|---|---|---|---|---|---|---|
| R1 | **Platform bundling**: Cloudflare adds intent/step-up policies to MCP Portals or Workers; Okta pushes action-level policy via CAA | 4 | 5 | **20** | Live *inside* the app (SDK), consume their signals, be the best integration; own the cross-tenant insights data they don't collect | Cloudflare/Okta changelog mentions "intent", "action policy", "step-up for agents" |
| R2 | **Intent declaration never catches on** — agents don't send intent, so conformance is guesswork | 3 | 4 | 12 | Product works in policy-only and inferred modes; declared intent is an upgrade, not a requirement; carrot (higher limits) for declarers | <5% of sessions declare intent after 3 months |
| R3 | **False positives break customers' agent UX** (legit agent blocked mid-task) | 4 | 3 | 12 | Log-only default; step-up instead of block; per-route rollout; one-click allowlist | Partners disabling enforcement |
| R4 | **Wallarm/Salt move down-market** with self-serve agent tiers | 3 | 4 | 12 | Speed + developer experience; open spec; price | Self-serve pricing pages from API-security vendors |
| R5 | **Model vendors own the "agent safety check"** (Jev, OpenAI Decisions) and SaaS teams DIY with 20 lines | 3 | 3 | 9 | Value is placement + policy + approvals + insights, not the classifier; make the `Decider` swappable | Blog posts showing DIY recipes gaining traction |
| R6 | **Pain not yet felt in mid-market** — "we'll deal with it after an incident" | 3 | 4 | 12 | Lead with incidents; target infra/dev-tool SaaS with destructive APIs; sell the *launch unblocker* (47% delayed releases) | Interview gate in doc 06 |
| R7 | **Latency/availability**: middleware in the request path becomes a blamed dependency | 2 | 4 | 8 | Fail-open by default (configurable), local rules, async model checks, no network call on hot path | p99 added latency >5 ms |
| R8 | **Data protection**: payloads carry end-user personal data; insights could leak it | 2 | 4 | 8 | Reuse original repo's constraints: redact before persistence, aggregate-first, short TTL, per-tenant keys | Customer DPA questions in pilots |
| R9 | **Solo-founder bandwidth** across SDKs, dashboard, sales | 4 | 3 | 12 | Two languages max (Go, TS); no enterprise features pre-PMF; ruthless kill criteria | >50% of time on non-customer work |
| R10 | **Standards churn** (MCP auth spec changed in Jul 2026; drafts may die) | 3 | 2 | 6 | Adapter layer per intent source; track IETF WG | — |

**Top three to watch:** R1 (bundling), R6 (pain not yet felt), R2/R3 (intent adoption and
false positives). If R1 materialises before 10 paying customers, pivot per doc 06 §6.
