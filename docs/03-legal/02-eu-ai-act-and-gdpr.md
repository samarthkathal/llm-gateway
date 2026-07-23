# EU AI Act & GDPR — Deep Dive

## EU AI Act

### Where we fall in the Act

Annex III, Area 4 ("Employment, workers management and access to self-employment") lists as
high-risk, among others, AI systems intended to be used:

> "to make decisions affecting terms of work-related relationships, the promotion or
> termination …, to allocate tasks based on individual behaviour or personal traits or
> characteristics, **or to monitor and evaluate the performance and behaviour of persons in
> such relationships**."
> — [Annex III](https://artificialintelligenceact.eu/annex/3/), [EU AI Act Service Desk](https://ai-act-service-desk.ec.europa.eu/en/ai-act/annex-3)

An intent-capture system *can* be argued outside this (governance/insight, not evaluation of
individuals), but the conservative planning assumption is: **if individual-level data about
work behaviour is processed and surfaceable, expect high-risk treatment.** Draft Commission
guidelines on high-risk classification are still evolving
([McCann FitzGerald analysis](https://www.mccannfitzgerald.com/knowledge/technology/employment-spotlight-eu-ai-act-draft-guidelines-on-high-risk-ai-classification)).

### Timeline (as of July 2026)

| Date | What applies |
|---|---|
| Feb 2025 | Prohibited practices (e.g., emotion recognition at work — relevant: **never** ship sentiment-on-employees features in EU) |
| Aug 2025 | GPAI obligations |
| **2 Dec 2027** | **Annex III high-risk obligations** (postponed from 2 Aug 2026); pre-existing systems grandfathered unless significantly changed |
| Aug 2030 | Public-authority systems |

Sources: [Crowell & Moring 2026 overview](https://www.crowell.com/en/insights/client-alerts/artificial-intelligence-and-human-resources-in-the-eu-a-2026-legal-overview),
[CSA readiness note](https://labs.cloudsecurityalliance.org/research/csa-research-note-eu-ai-act-high-risk-compliance-deadline-20/).

### Roles: we are provider; the customer is deployer

- **Us (provider):** risk-management system, data governance, technical documentation,
  logging capability, transparency/instructions for use, human-oversight design, accuracy &
  robustness, conformity assessment + CE marking if high-risk.
- **Customer (deployer):** use per instructions, human oversight, input-data relevance,
  monitoring, record-keeping — and **Art. 26(7): inform workers and their representatives
  before putting a high-risk workplace system into service**
  ([Article 26](https://artificialintelligenceact.eu/article/26/)).

**Product implication:** the deployer's duties are our onboarding checklist. Ship the worker
notification templates, oversight console, and log exports so the customer's Art. 26 duties
are a wizard, not a legal project. Penalties for non-compliance reach €15M / 3% global
turnover ([deployer guide](https://euaicompass.com/eu-ai-act-high-risk-deployer-guide.html)).

### Prohibited-practice red lines (never build, any tier)

- Emotion inference about employees (Art. 5) — no sentiment scoring of individuals.
- Social-scoring-like aggregation that affects individuals' treatment.
- Anything inferring protected characteristics from prompts.

## GDPR

### Lawful basis
- **Consent: presumptively invalid** in employment (imbalance of power). Do not architect
  around consent banners.
- **Legitimate interest (Art. 6(1)(f)):** the workable basis for intent governance — requires
  a documented Legitimate Interest Assessment; survives only with strict proportionality
  (minimization, aggregation, transparency).
- **Legal obligation (Art. 6(1)(c)):** available for the *policy-enforcement* functions in
  regulated sectors (e.g., recordkeeping duties) — a reason to keep governance and analytics
  as separable processing purposes.

### DPIA is effectively mandatory
Systematic monitoring + innovative tech + employee data = Art. 35 DPIA triggers on multiple
axes. **Productize it:** generate a pre-filled DPIA from the tenant's actual policy
configuration (data categories, retention, access model, k-floor). This converts a sales
blocker into an accelerator.

### Data-subject rights we must engineer for
- Access/portability (Art. 15/20): per-employee export of their intent records.
- Erasure (Art. 17): hard-delete pipelines incl. derived aggregates recomputation.
- Art. 22: no solely-automated decisions with significant effect — another reason the
  product must never feed automated employment decisions.

### Transfers & residency
EU tenants will demand EU processing. Plan EU region deployment + per-tenant data residency
from the architecture stage (cheap now, migration-project later). LLM sub-processing (our
intent-extraction calls) must run on EU-resident endpoints for EU tenants and appear in the
sub-processor list.

## Engineering requirements extracted (traceability to MVP)

| Legal driver | Product requirement | MVP phase |
|---|---|---|
| Art. 26(7) worker notice | Employee transparency page + notification templates | P1 |
| GDPR minimization | Structured-signal-first storage; raw-text TTL policy | P2 |
| GDPR Art. 15/17 | DSAR export & erasure endpoints | P3 |
| DPIA duty | Policy-config → DPIA-input generator | P3 |
| High-risk logging | Immutable audit log, exportable | P2 |
| Proportionality | k-anonymity floor on all insight surfaces | P2 |
| Residency | Tenant-pinned region/keys | Architecture now, EU region post-MVP |
