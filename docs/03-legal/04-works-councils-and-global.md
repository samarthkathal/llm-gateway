# Works Councils, Sector Rules & Rest-of-World

## Works councils — the real EU gatekeeper

Before any EU AI Act question arises, technical systems capable of monitoring employee
behaviour or performance typically require **co-determination approval** from employee
representative bodies. This is a *pre-deployment veto*, not a disclosure duty.

| Country | Instrument | Effect on us |
|---|---|---|
| Germany | BetrVG §87(1)(6) — works council must approve introduction of technical systems "designed to monitor" behaviour/performance (case law: *capable of* monitoring suffices) | Works Council agreement (Betriebsvereinbarung) required per customer before go-live; typically 2–6 months |
| France | CSE consultation before introducing means of employee activity monitoring; CNIL proportionality doctrine | Consultation + DPIA; individual continuous monitoring presumed disproportionate |
| Netherlands | WOR Art. 27 — works council consent for personnel-monitoring systems | Consent required; aggregate-only posture materially helps |
| Nordics | Strong union consultation norms; co-determination acts (e.g., Swedish MBL) | Union negotiation expected |
| Austria/Belgium | Similar co-determination regimes | Same pattern |

### Strategy: make co-determination a wedge, not a wall

1. **Ship a "works-council mode":** aggregate-only insights (k≥5 hard floor), no individual
   drill-down capability *at all* (absent, not permissioned-off — councils and courts
   distinguish "cannot" from "does not"), short raw-text TTL, on-record processing purposes.
2. **Ship the Betriebsvereinbarung starter pack:** template agreement + technical annex
   describing exactly what the system can and cannot do. Selling into Germany without this
   adds ~1–2 quarters per deal; with it, we become the vendor counsel recommends.
3. **Sequence EU entry after the pack exists** — see effort matrix
   ([`../04-effort/01-effort-matrix.md`](../04-effort/01-effort-matrix.md)).

## Sector overlays (revenue-relevant, not just risk)

| Sector | Rule | Product hook |
|---|---|---|
| Financial services | FINRA/SEC books-and-records; communications surveillance duties | Retention policies + immutable archive per tenant — they *must* buy something like this |
| Healthcare | HIPAA — PHI in prompts | PHI detection/redaction policy class; BAA readiness |
| Legal | Privilege | Matter-aware exclusion policies |
| Defense/export | ITAR/EAR | Data-residency + provider-restriction policies (e.g., "US-only endpoints") |

## Rest-of-world quick scan

| Region | Posture | Note |
|---|---|---|
| UK | UK GDPR + ICO employment guidance; no works councils | Near-US ease with GDPR-grade duties; good second market |
| Canada | PIPEDA + provincial (Québec Law 25 strict) | Notice + proportionality; similar to CA-US |
| Australia/NZ | Workplace surveillance acts (NSW/ACT/VIC: notice requirements) | Notice-based; straightforward |
| India | DPDP Act 2023 — employer processing allowed for employment purposes | Permissive today; large English-speaking SAM |
| Japan | APPI + labor practice | Consultation culture; localization needed |
| Brazil | LGPD (GDPR-like) | Same architecture carries |
| China | PIPL + data-localization | Out of scope indefinitely (localization + state-access complexity) |

## Compliance operations roadmap (company-level, not product)

| When | Milestone |
|---|---|
| Pre-revenue | Privacy policy, DPA template, sub-processor list, security whitepaper |
| First paying tenant | SOC 2 Type I program start; pen test |
| ~Month 12 | SOC 2 Type II; ISO 27001 decision (EU sales gate) |
| Pre-EU entry | EU hosting region, EU rep (GDPR Art. 27), works-council pack, DPIA generator live |
| Pre-Dec-2027 | EU AI Act provider conformity file for any feature that could be classed high-risk |
