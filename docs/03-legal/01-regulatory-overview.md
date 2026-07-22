# Regulatory Overview — Can This Be Built Legally?

> **Verdict: yes, conditionally.** The product is lawful in every target market *if* it is
> architected as transparent, proportionate, aggregate-first governance tooling — and it is
> likely unlawful (or commercially undeployable) in key markets if built as individual-level
> monitoring. Legality is therefore an architecture input, not a compliance afterthought.
>
> Nothing in this folder is legal advice; engage employment/privacy counsel per jurisdiction
> before any production deployment.

## The four regulatory fronts

| Front | Instruments | Bite | Detail doc |
|---|---|---|---|
| EU AI systems law | EU AI Act (Annex III area 4: employment/worker management) | High-risk classification → deployer duties, worker notice; fines up to €15M / 3% turnover | [`02-eu-ai-act-and-gdpr.md`](./02-eu-ai-act-and-gdpr.md) |
| EU/UK data protection | GDPR / UK GDPR | DPIA required; lawful-basis problem (consent weak in employment); data-subject rights | [`02-eu-ai-act-and-gdpr.md`](./02-eu-ai-act-and-gdpr.md) |
| US state law | NY/CT/DE/CO monitoring-notice laws; CA FEHA ADS rules + 2026 privacy expansion; IL/CO AI-decision laws | Notice, disclosure, data-minimization duties; per-employee penalties | [`03-us-state-laws.md`](./03-us-state-laws.md) |
| Labor co-determination | Works councils (DE BetrVG §87, FR CSE, NL WOR); sector rules (FINRA/SEC, HIPAA) | Pre-deployment approval requirements; can veto or delay go-live for months | [`04-works-councils-and-global.md`](./04-works-councils-and-global.md) |

## Cross-cutting conclusions

### 1. The timeline gift: EU AI Act high-risk duties slipped to ~Dec 2027
Annex III high-risk obligations (which cover workplace AI systems) were postponed from
2 Aug 2026 to **2 Dec 2027**; systems already in use before that date are grandfathered
unless significantly changed (public authorities: 2030)
([McCann FitzGerald](https://www.mccannfitzgerald.com/knowledge/technology/employment-spotlight-eu-ai-act-draft-guidelines-on-high-risk-ai-classification),
[Crowell & Moring](https://www.crowell.com/en/insights/client-alerts/artificial-intelligence-and-human-resources-in-the-eu-a-2026-legal-overview)).
This gives ~17 months of runway to build conformity into the product and turn the deadline
into a *sales catalyst* (customers must be compliant; we ship the compliance).

### 2. Classification strategy: govern the AI, don't evaluate the worker
Annex III area 4 captures systems used for decisions on promotion/termination/task
allocation and for **monitoring and evaluating performance and behaviour** of workers
([Annex III text](https://artificialintelligenceact.eu/annex/3/)). Our exposure depends on
what the product *does and claims*:
- **Intent-governance mode** (aggregate insights, policy enforcement, no individual
  evaluation feeding employment decisions) → arguable non-high-risk or lower-obligation
  posture, though conservative counsel will still assume high-risk.
- **Any feature scoring individuals or feeding performance processes** → squarely high-risk,
  full deployer-duty stack (risk mgmt, logging, human oversight, worker notification under
  Art. 26(7); [Article 26](https://artificialintelligenceact.eu/article/26/)).
**Product rule derived:** no individual-evaluation features in core product; if a tenant
demands them, that's a separately-contracted high-risk module with the full conformity stack
— priced accordingly.

### 3. Consent won't carry this; design for legitimate interest + transparency
In employment relationships GDPR consent is presumptively invalid (power imbalance). The
viable basis is legitimate interest with a documented balancing test — which only survives if
the system is *necessary and proportionate*: minimal data, short retention, aggregate
outputs, real transparency. Every one of those is an architecture requirement.

### 4. The US is permissive-with-notice, and tightening
Monitoring is broadly lawful with advance written notice (NY/CT/DE/CO explicitly; CA now
adds data-minimization and access rights; IL/CO regulate AI-driven employment decisions)
([state tracker](https://www.aipolicydesk.com/blog/ai-employee-monitoring-disclosure-laws-2026-state-tracker),
[state-by-state guide](https://www.intelogos.com/blog/employee-monitoring-laws-by-state)).
Ship notice tooling (employee-facing disclosure page, poster/ack templates) as product.

### 5. Sector overlays are a revenue feature
FINRA/SEC recordkeeping, HIPAA, privilege, export control — regulated tenants *need* a
policy engine that can express their obligations. What is a burden for a monitoring product
is the core value proposition of a governance product.

## The legal posture, in product terms

1. Aggregate-first insights with k-anonymity floor (default k≥5, tenant-raisable).
2. Structured intent signal retained; raw prompt text minimized/short-TTL per tenant policy.
3. Employee transparency page: what is captured, what managers can see, rights & contacts.
4. Per-jurisdiction policy modes (e.g., "DE works-council mode" = aggregate-only, no
   individual drill-down at all).
5. DPIA-input generator, audit log, DSAR export/delete as first-class features.
6. Immutable record of tenant policy configuration (proves the tenant's compliance posture).
