# Effort Matrix — Metrics-Based Analysis

Scoring model: each cell rates **effort required** (1 = trivial → 5 = extreme) for each
market-entry dimension, per region. "Score" = weighted attractiveness
(revenue potential ÷ blended effort), higher = better. Weights: legal 0.35, GTM/marketing
0.25, product localization 0.15, compliance ops 0.25 — legality and compliance dominate
because they gate deployment entirely in this category.

## Region entry matrix

| Region | Legal effort | GTM/marketing effort | Product localization | Compliance ops | Revenue potential (1–5) | Attractiveness |
|---|---|---|---|---|---|---|
| **US mid-market** | 2 (notice regimes) | 2 (direct + PLG-ish) | 1 | 2 (SOC 2) | 4 | **★ 1.9 — enter first** |
| **UK** | 2.5 (UK GDPR) | 2 | 1 | 2.5 | 3 | 1.3 — fast follow |
| **ANZ** | 2 (surveillance acts, notice) | 2.5 (presence needed) | 1 | 2 | 2 | 1.0 — opportunistic |
| **US enterprise/regulated** | 3 (sector rules) | 4 (long cycles, security review) | 2 (VPC deploy) | 4 (SOC2 II, HIPAA, FINRA) | 5 | 1.5 — year 2 |
| **DACH** | 5 (works councils + GDPR + AI Act) | 4 (trust-building, local presence) | 3 (DE language, EU residency) | 5 | 4 | 0.9 — year 2–3, with pack |
| **France/Benelux/Nordics** | 4.5 | 4 | 3 | 4.5 | 3 | 0.7 — follows DACH pack |
| **India/SEA** | 2 (DPDP permissive) | 3 (price sensitivity) | 1 | 2 | 2 | 1.0 — opportunistic inbound |
| **Japan** | 3.5 | 4.5 (channel-dependent) | 4 | 3.5 | 3 | 0.5 — not before Series A |

## Marketing effort breakdown (beachhead: US mid-market)

| Channel | Effort (1–5) | Expected yield | Notes |
|---|---|---|---|
| Design-partner outbound (founder-led) | 3 | **Highest** | 10 warm CISO/CIO intros > any campaign at this stage |
| Technical content/SEO ("shadow AI", "AI usage governance", "EU AI Act deployer") | 2 | High, compounding | Regulatory-deadline content has built-in urgency |
| Open-source gateway shim (OSS capture layer, paid brain) | 4 | Medium-high | Counters LiteLLM gravity; dev mindshare funnel |
| Analyst relations (Gartner AI TRiSM / AI governance categories) | 3 | Medium (year 2 payoff) | Category placement matters for enterprise procurement |
| Security-conference presence (BSides→RSA arc) | 3 | Medium | CISO trust channel |
| Paid acquisition | 2 | **Low** | Category too new; CAC unjustifiable pre-A |

## Legality effort ranking (summary of [`../03-legal/`](../03-legal/))

Cheapest → most expensive to become deployable:
1. India (DPDP) → 2. US notice states → 3. ANZ → 4. UK → 5. California-grade US → 6.
Canada/Québec → 7. EU non-works-council → 8. EU works-council countries (DE hardest).

## Engineering effort calibration (for the MVP plan)

| Workstream | Effort (eng-weeks, 2 engineers) | Confidence |
|---|---|---|
| Multi-tenant core (tenants, users, RBAC, RLS) | 3–4 | High |
| SSO (OIDC) + SCIM ingest | 3 | High |
| Gateway hardening (real failover, streaming, provider adapters) | 3 | High — seed code exists |
| Capture pipeline + redaction + intent extraction | 4–5 | Medium — extraction quality is the unknown |
| Policy engine v1 (retention, redaction rules, k-floor) | 3 | Medium |
| Audit log + DSAR export | 2 | High |
| Insights aggregation + minimal dashboard | 3–4 | Medium |
| Workday connector | 3 (defer; SCIM+CSV first) | Low — API idiosyncrasy risk |
| **Total MVP** | **~21–27 eng-weeks** ≈ 12 calendar weeks for 2 engineers | |

## Decision rule this matrix implies

Spend year-1 effort where legal effort ≤2.5 (US mid-market, UK, ANZ, India inbound); spend
year-1 *preparation* (packs, residency, certifications) for the high-effort/high-revenue EU
cells so year-2 entry is unblocked; never spend effort where revenue ≤2 and legal ≥3.
