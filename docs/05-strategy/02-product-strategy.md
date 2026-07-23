# Product Strategy

## Product principles (constraints that outrank features)

1. **Capture once, serve many.** One gateway interception feeds security policy, compliance
   evidence, and leadership insight. No second agent, no endpoint software.
2. **Aggregate by default, individual by exception.** Every insight surface enforces a
   k-anonymity floor (default k≥5). Individual-level access exists only for the data subject
   themself (transparency page) and DSAR flows. In works-council mode it doesn't exist at all.
3. **The employee is a user, not a target.** Transparency page ships in v1, not as a
   compliance retrofit. If we'd be uncomfortable showing an employee exactly what their
   employer sees, we don't build it.
4. **Policy is data.** Retention, redaction, jurisdiction modes, k-floors are per-tenant
   configuration with an immutable change log — because the tenant's config *is* their
   compliance posture (and our DPIA generator input).
5. **Meaning over payloads.** Derived structured intent is the asset; raw text is a
   liability with a TTL. Default: raw prompt/response discarded after extraction; tenants
   may opt into short-TTL retention windows.

## The core object: the intent record

```
IntentRecord {
  tenant_id, user_ref (pseudonymized), org_unit_path,   // from SCIM/Workday/local
  timestamp, provider, model,
  intent_category,        // taxonomy: build / analyze / write / research / support / admin ...
  topic_tags[],           // project/domain clusters, tenant-local
  work_context,           // inferred artifact type: code, contract, campaign, report ...
  effort_signal,          // tokens, turns, session depth
  outcome_signal,         // accepted/refined/abandoned (heuristic)
  policy_events[],        // redactions applied, blocks, escalations
}
```

Raw text → (redaction) → (extraction LLM, tenant-region endpoint) → IntentRecord →
(aggregation jobs) → insight surfaces. The taxonomy starts fixed (12–15 categories) +
tenant-local topic clustering; taxonomy quality is the product's kernel — instrument and
iterate on it before anything else.

## Tiering (maps to buyers)

| Tier | Contents | Buyer |
|---|---|---|
| **Gateway** (entry) | Multi-provider routing/failover, SSO, budgets, usage observability | CIO/platform |
| **Governance** (platform fee) | Policy engine, redaction, audit log, retention control, DSAR, transparency page | CISO/compliance |
| **Insight** (seat expansion) | Org-aware intent dashboards, adoption & effort maps, AI-ROI reporting | COO/CFO/CHRO |
| **Enterprise add-ons** | VPC deploy, EU residency, Workday connector, works-council mode, sector policy packs | Regulated/EU |

## Roadmap beyond MVP (thematic, not dated)

1. **Extraction quality flywheel:** eval harness on design-partner traffic; human-in-the-loop
   taxonomy corrections; per-tenant fine-tuned classifiers when volume justifies.
2. **Connector breadth:** SCIM → CSV → Workday → SuccessFactors/BambooHR; MCP/agent-traffic
   capture (agentic workloads will dwarf chat traffic; the gateway must speak MCP early).
3. **Benchmarks:** cross-tenant aggregate industry baselines (opt-in, k-protected).
4. **Deploy modes:** SaaS → EU region → VPC → (only if a lighthouse deal demands) air-gapped.
5. **Ecosystem:** export APIs into Credo/OneTrust (compliance), BI tools (insight), SIEM
   (security events) — be the source of record, let others own workflows.

## Explicit non-goals (v1–v2)

- Individual productivity scores or manager views below the k-floor — never (principle 2).
- Sentiment/emotion analysis — prohibited in EU workplaces, poison everywhere (Art. 5).
- Capturing non-AI channels (email/Slack bodies) — different product, different company.
- Building our own guardrail/injection-defense research capability — integrate instead.
- Model hosting/inference — we govern traffic, we don't serve models.

## Technical architecture notes (MVP-relevant decisions)

- **Tenancy:** single Postgres, `tenant_id` everywhere + RLS enforced at DB layer; isolation
  test suite in CI from first migration. Per-tenant envelope encryption keys (KMS) so tenant
  deletion = key destruction. Escape hatch to dedicated-DB for enterprise tier later.
- **Gateway:** evolve the existing Go service (`internal/llmgateway`) — Go is the right
  language for this layer. Real health-checked failover replaces the current random-split
  prototype; streaming passthrough; provider adapters behind the existing interface.
- **Async spine:** capture → queue → extraction workers, so gateway latency is unaffected by
  analysis (target: <10ms added p50 on passthrough).
- **Extraction:** provider-agnostic extraction calls through our own gateway (dogfood),
  pinned to tenant-region endpoints.
- **Identity:** OIDC for humans, API keys per workload; SCIM inbound for hierarchy; local
  users as fallback. Hierarchy stored as materialized path for fast rollups.
