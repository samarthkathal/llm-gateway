# Indicative Financial Model (Pre-Seed → Series A Horizon)

> Directional planning model, not a forecast. Assumptions are listed so they can be attacked
> individually; update quarterly against actuals.

## Pricing architecture (hypothesis to validate with design partners)

| Component | Price | Basis |
|---|---|---|
| Platform fee (Gateway+Governance) | $15–25k/yr | Gates policy engine, audit, SSO/SCIM |
| Seats (captured users) | $3–6/user/mo | Volume-tiered; the expansion engine |
| Insight tier | +30–50% on platform fee | Exec dashboards, exports, SEI integrations |
| Enterprise add-ons | $10–50k/yr each | VPC, EU residency, Workday, works-council mode, sector packs |

Reference blended ACVs: mid-market entry ~$45k; expanded mid-market ~$80k; regulated
enterprise $150k+.

## Cost structure peculiarities of this product

1. **Extraction compute is a real COGS line.** Every captured request costs an extraction
   call. Model: small/cheap model (Haiku-class) at ~$0.001–0.003 per request → at 50 requests
   /user/day: **$1–3/user/mo COGS** against $3–6/user/mo seat price. Gross margin on seats
   ~50–70% *unless* we: batch-extract, sample low-signal traffic, cache by similarity, or
   train a small classifier once volume allows (path back to 85%+). This is the #1 unit-
   economics lever and a deliberate MVP measurement goal.
2. **Gateway egress/compute is negligible** relative to extraction (<5% of COGS).
3. **Compliance is a fixed-cost staircase:** SOC 2 I (~$30–50k incl. tooling/audit), pen
   test (~$20k), SOC 2 II + ISO 27001 (~$60–100k cumulative), EU legal reviews (~$30–60k).
   Each unlocks a revenue tier; time them to demand, not to fear.

## Scenario model (36 months)

| | M6 | M12 | M18 | M24 | M36 |
|---|---|---|---|---|---|
| **Base case** tenants | 2 | 10 | 20 | 35 | 70 |
| ARR | $80k | $500k | $1.1M | $2.2M | $5M |
| Team | 3 | 6 | 9 | 14 | 25 |
| Net burn/mo | $60k | $110k | $160k | $220k | $350k |
| **Bear** (governance-only sells) ARR | $50k | $250k | $500k | $900k | $1.8M |
| **Bull** (AI-ROI budget line materializes) ARR | $120k | $800k | $2M | $4.5M | $10M |

Base case reaches default-alive around M30 on seed capital if Series A slips — the model is
deliberately survivable-in-bear: governance-tier revenue (security budget, annual prepay)
covers a 3–4 person core team by M18 even in bear.

## Key unit-economics targets

| Metric | Target by M12 | Why |
|---|---|---|
| Gross margin (blended) | ≥70% | Extraction-cost engineering proves out |
| CAC payback | <12 mo | Founder-led sales at entry ACVs makes this achievable |
| NRR | ≥115% | Seat growth + insight attach; the Series A story |
| Extraction cost/request | <$0.0015 | Sampling+caching+small-model roadmap |
| Logo retention | 100% | Tiny base; any churn is a product alarm |

## Use-of-funds discipline

Pre-seed: ~80% engineering/product, ~10% legal/compliance, ~10% infra.
Seed: ~55% eng, ~20% GTM, ~15% compliance staircase, ~10% ops.
No paid marketing until two tiers show organic pull (see
[`../04-effort/02-go-to-market-effort.md`](../04-effort/02-go-to-market-effort.md)).
