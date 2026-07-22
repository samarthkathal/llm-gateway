# MVP Plan — 12 Weeks to Design-Partner Deployable

## MVP thesis

Prove, with 3–5 design partners, that **one gateway deployment can (a) pass a security
team's governance checklist and (b) produce an org-level insight an executive repeats in a
meeting** — the "double story" defined in
[`../04-effort/02-go-to-market-effort.md`](../04-effort/02-go-to-market-effort.md).

Everything below is scoped to that proof. Anything not on the critical path of that
sentence is out.

## Current state of this repo

The repo contains a prototype Go LLM gateway (`internal/llmgateway`, `internal/llmprovider`)
with a provider interface, Claude/OpenAI mock providers, and a naive failover loop in
`main.go`. It is the seed of Phase 3 below; everything else is greenfield.

## Build phases

### P0 (week 1): foundations & scaffolding
- Repo restructure: `cmd/gateway`, `cmd/api`, `cmd/worker`, `internal/...`; CI (build, test,
  lint); Postgres + migrations tooling; config layer replacing the hardcoded values noted in
  `gateway.go`.

### P1 (weeks 2–4): multi-tenant identity core
- Tenants, local users, org units; RBAC (admin / manager / member / auditor).
- **Postgres RLS on every tenant-scoped table from the first migration + isolation test
  suite in CI** (non-negotiable; see risk R5).
- OIDC login; API keys per tenant/workload.
- SCIM v2 inbound (users + manager attribute → hierarchy as materialized path); CSV import
  fallback. *(Workday connector explicitly deferred — see effort matrix.)*
- Employee transparency page v0 (static per-tenant disclosure: what is captured, what
  managers can/cannot see) — shipping this in P1 is the legal posture made visible.

### P2 (weeks 4–7): governed capture pipeline
- Harden gateway: real health-checked failover (replace random-split), streaming
  passthrough, tenant/user context threaded through `common.Request`, <10ms added p50,
  passthrough-on-failure mode.
- Capture → queue → worker: redaction pass (secrets/PII patterns) **before** persistence;
  intent extraction via small-model call producing the IntentRecord
  (schema in [`../05-strategy/02-product-strategy.md`](../05-strategy/02-product-strategy.md));
  fixed 12-category taxonomy v1 + tenant-local topic tags.
- **Default: raw text discarded post-extraction**; per-tenant TTL opt-in (0/7/30 days).
- Immutable audit log (policy events, admin actions, config changes).
- Extraction eval harness with a labeled seed set — taxonomy quality is the product kernel
  (risk R3); measure from the first week it runs.

### P3 (weeks 7–10): policy engine & compliance surface
- Per-tenant policy config as data: retention TTLs, redaction rule classes, provider
  allow/deny, jurisdiction mode (v1: `standard` / `strict-aggregate` for works-council
  posture), k-floor value (default 5, raisable only).
- Policy-config change log (immutable) + DPIA-input export generated from live config.
- DSAR endpoints: per-user export + erasure (incl. queued recompute of aggregates).

### P4 (weeks 9–12, overlapping): insight surface v1
- Aggregation jobs: intent-category and topic rollups by org unit, k-floor enforced at
  query layer (not UI layer).
- One dashboard that earns the exec quote: org-level "where AI effort goes" (categories ×
  org units × trend), adoption depth, top emergent topics. Plus CSV/API export.
- No individual views anywhere. Manager views start at k≥5 rollups.

### Continuous
- Weekly design-partner deployments from week 6 (behind flags); overhead/latency SLO
  dashboards; extraction-cost metering per tenant (feeds the unit-economics model in
  [`../07-financial/02-financial-model.md`](../07-financial/02-financial-model.md)).

## Team & effort

2 engineers × 12 weeks (~21–27 eng-weeks estimated in
[`../04-effort/01-effort-matrix.md`](../04-effort/01-effort-matrix.md)) + founder on
design-partner recruitment from week 0. Solo-founder variant: 16–18 weeks with P4 cut to a
single hardcoded dashboard.

## MVP acceptance criteria

1. Two tenants on one deployment; isolation test suite green; pen-test-ready RLS.
2. A design partner routes real traffic through SSO-authenticated gateway with <10ms p50 overhead.
3. Security checklist demo: policy config → enforced redaction → audit trail export, end-to-end.
4. Intent dashboard shows a *true, non-obvious* statement about the partner org's AI usage,
   validated by the partner ("yes, that's real, and we didn't know it").
5. Transparency page live for partner employees; DSAR export works.
6. Extraction cost measured at <$0.003/request with the sampling knob functional.

## Explicitly deferred (and where the plan for each lives)

| Deferred | Trigger to build | Doc |
|---|---|---|
| Workday connector | First enterprise deal requiring it | [`../04-effort/01-effort-matrix.md`](../04-effort/01-effort-matrix.md) |
| EU residency + works-council pack | EU pipeline ≥2 qualified | [`../03-legal/04-works-councils-and-global.md`](../03-legal/04-works-councils-and-global.md) |
| SEI/DORA integrations & CTO module | Insight tier validated | [`../01-market/04-adjacent-metrics-movements.md`](../01-market/04-adjacent-metrics-movements.md) |
| VPC/self-host deploy | Regulated-segment entry | [`../01-market/03-target-segments.md`](../01-market/03-target-segments.md) |
| SOC 2 program | First paying tenant | [`../03-legal/04-works-councils-and-global.md`](../03-legal/04-works-councils-and-global.md) |
| MCP/agent-traffic capture | Agent traffic at partners >10% | [`../05-strategy/02-product-strategy.md`](../05-strategy/02-product-strategy.md) |
