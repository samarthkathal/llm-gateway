# CLAUDE.md — Project Context for AI Assistants

## What this project is

**Intent Gateway**: a multi-tenant enterprise LLM gateway that captures workforce *intent*
signals from AI usage, stores them under per-tenant policy/governance, and produces
aggregate, org-aware insights. Target buyers: CISO/CIO first (governance + gateway),
COO/CFO second (insight tier). Currently **pre-MVP**: research is done, the build follows
[`docs/08-mvp/01-mvp-plan.md`](./docs/08-mvp/01-mvp-plan.md).

Read [`docs/README.md`](./docs/README.md) for the full research base (market, competition,
legal, effort, strategy, risk, financial, MVP plan). When product questions arise, those
docs are the source of truth; update them when decisions change.

## Non-negotiable product constraints

These come from the legal/strategy research and outrank any feature request. Do not write
code that violates them; if a request conflicts, flag it and point to the relevant doc.

1. **Aggregate-first, k-anonymity floor (default k≥5)** on every insight surface, enforced
   at the query layer, not the UI. No individual-level manager views.
   → `docs/05-strategy/02-product-strategy.md`
2. **No individual evaluation/scoring features** — nothing that makes or substantially
   assists employment decisions (EU AI Act Annex III, NYC LL144, IL/CO ADS laws).
   → `docs/03-legal/`
3. **No sentiment/emotion inference on employees, ever** (EU AI Act Art. 5 prohibition).
   → `docs/01-market/04-adjacent-metrics-movements.md` (morale red line)
4. **Raw prompt/response text is a liability**: redact before persistence, discard after
   extraction by default, tenant-configurable short TTL only.
5. **Tenant isolation is existential**: Postgres RLS on every tenant-scoped table + the
   isolation test suite must stay green in CI. Per-tenant envelope keys.
6. **Policy is data**: retention/redaction/jurisdiction/k-floor live in per-tenant config
   with an immutable change log (feeds the DPIA generator).
7. **Employee transparency page** ships with capture, never after it.

## Codebase

- **Language:** Go (module `emergent`, Go 1.25). Standard library bias; add dependencies
  deliberately.
- **Current layout:** `main.go` (prototype failover simulation),
  `internal/llmgateway` (gateway core), `internal/llmprovider` (provider interface +
  Claude/OpenAI mocks), `internal/common` (shared types).
- **MVP target layout:** `cmd/gateway`, `cmd/api`, `cmd/worker`, expanded `internal/` —
  see the MVP plan P0.
- **Known prototype flaws** (intentional, to be replaced, not patched): hardcoded
  primary/secondary providers, random-split failover, `println` debugging, empty
  Request/Response types, no tests, no persistence, no auth.

## Working conventions

- Run `go build ./... && go vet ./...` before committing; add tests with any non-trivial
  new package (the prototype has none — new code should not follow that example).
- Keep docs and code in sync: a change that alters pricing, tiering, policy semantics, or
  legal posture must update the corresponding file under `docs/`.
- Vocabulary matters (legal + positioning): use *governance, visibility, insight*; never
  *monitoring, surveillance, tracking, scoring* in code identifiers, UI strings, or docs.
  → `docs/05-strategy/01-positioning-and-strategy.md`
- Development branch for the current workstream: `claude/enterprise-intent-gateway-mvp-91dtjy`.
