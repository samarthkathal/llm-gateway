# Differentiation Thesis

## The one-sentence claim

> We are the only system that turns an enterprise's AI traffic into an **org-aware,
> policy-governed intent graph** — telling leadership where work and AI effort actually go,
> without becoming surveillance.

## Three stacked differentiators (each requires the previous)

### 1. Position: the authenticated choke point
Inline gateway on sanctioned AI traffic → we see *authenticated, per-user, full-fidelity*
requests, unlike network SSE tools (flow-level) or workforce analytics (proxy signals).
Not defensible alone — gateways are commodity — but it is the required vantage point.

### 2. Data model: the intent graph (the actual moat)
Every request → structured intent record `{user, org-unit, manager-chain, intent category,
project/topic, tooling, outcome quality}` joined to the HRIS/Workday/SCIM hierarchy.
Accumulated over time this becomes a longitudinal map of how the organization actually
works — a dataset that (a) no point competitor collects, (b) compounds with tenure,
(c) switching cost grows with history. **This is what we own that the acquirers of
Prompt Security and Lakera did not buy.**

### 3. Contract: governance-native by design
Aggregate-first (k-anonymity floor), per-tenant policy engine (retention, redaction,
jurisdiction modes), employee transparency surface, DPIA/audit outputs as product artifacts.
This is simultaneously our legal survival requirement and a *sales weapon* in EU/works-council
accounts where bossware-pattern competitors structurally cannot pass review.

## Why incumbents don't just copy it

| Incumbent | Why copying is hard for them |
|---|---|
| Gateway vendors | Would need HRIS integrations, people-data compliance posture, and a business-buyer motion — none in their DNA or margin model. |
| Security platforms (S1, Check Point, Netskope) | Their product contract is *block and discard*; retaining semantic people-data creates liability their model doesn't price, and their buyer (CISO) didn't ask for COO insights. |
| Microsoft | Can copy the most; but is Copilot-ecosystem-biased, and enterprises increasingly resist Microsoft holding yet another sensitive graph. Multi-provider neutrality is a real wedge. |
| Workforce analytics | No gateway position, no engineering trust, and a brand that works councils reject. |
| Governance platforms | No inline infrastructure; adding it is a company-rebuild. |

## What we deliberately do NOT differentiate on

- **Routing performance / latency** — inherit table stakes (or embed LiteLLM); never the pitch.
- **Prompt-injection defense depth** — partner or integrate; Check Point/S1 won that race.
- **Individual productivity scoring** — a differentiator we *refuse*: it would be the fastest
  path to revenue with bossware buyers and the fastest path to category death. Codify the
  refusal (k-floor, no per-person manager dashboards by default) so it's credible.

## Proof points the MVP must demonstrate

1. Same deployment satisfies a CISO checklist (policy, redaction, audit) **and** produces a
   COO-grade insight ("Org X spent 34% of AI-assisted effort on migration project Y").
2. Works-council pack: show a German customer's counsel the aggregate-only mode, retention
   proof, and employee transparency page — and get a written "would pass co-determination"
   opinion. That artifact is worth more than any feature.
3. Tenant isolation demonstrable under audit (RLS + per-tenant keys), because "intent data"
   leakage across tenants is existential.
