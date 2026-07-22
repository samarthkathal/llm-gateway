# Intent Gateway — Pre-Ideation Research & Planning Docs

This folder is the pre-ideation knowledge base for the **Enterprise Intent Gateway**: a
multi-tenant LLM gateway that captures workforce *intent* signals from AI usage, stores them
securely, enforces per-tenant policy and governance, and surfaces organization-level insights.

The research was compiled in July 2026 as a market-analysis step **before** committing to the
MVP build. Read it top-to-bottom for the full picture, or jump to the section you need.

## Structure

| Section | Contents |
|---|---|
| [`01-market/`](./01-market/) | What market this actually is, market sizing (TAM/SAM/SOM), target segments |
| [`02-competition/`](./02-competition/) | Competitive landscape map, competitor profiles, differentiation thesis |
| [`03-legal/`](./03-legal/) | Regulatory overview, EU AI Act + GDPR deep dive, US state laws, works councils & global |
| [`04-effort/`](./04-effort/) | Metrics-based effort matrices (marketing, market-entry, legal per region), GTM effort |
| [`05-strategy/`](./05-strategy/) | Positioning & strategic analysis, product strategy |
| [`06-risk/`](./06-risk/) | Risk register with likelihood/impact scoring and mitigations |
| [`07-financial/`](./07-financial/) | Funding landscape & comparables, indicative financial model |
| [`08-mvp/`](./08-mvp/) | The MVP build plan derived from all of the above |

## How to read this

1. **If you have 10 minutes:** read [`01-market/01-market-overview.md`](./01-market/01-market-overview.md),
   then [`02-competition/03-differentiation.md`](./02-competition/03-differentiation.md),
   then [`08-mvp/01-mvp-plan.md`](./08-mvp/01-mvp-plan.md).
2. **If you are evaluating whether to invest time/money:** add
   [`06-risk/01-risk-register.md`](./06-risk/01-risk-register.md) and
   [`07-financial/01-funding-landscape.md`](./07-financial/01-funding-landscape.md).
3. **If you are about to talk to a prospect or lawyer:** read all of [`03-legal/`](./03-legal/).

## Headline conclusions (TL;DR)

- **The wedge is real but the framing matters.** "Intent capture" sold as employee monitoring
  is legally heavy and commercially toxic; sold as *AI-usage governance + enterprise insight*
  it rides three funded tailwinds: shadow-AI risk, EU AI Act compliance, and AI-ROI
  measurement.
- **The gateway layer alone is commoditized.** LiteLLM/Portkey/Kong/TrueFoundry own
  routing/observability. The defensible layer is the *org-aware intent graph* on top —
  hierarchy-joined, policy-governed, aggregate-first insights.
- **Legality is a product feature, not a checkbox.** EU AI Act Annex III (workplace) rules now
  land ~Dec 2027; US state notice laws (NY/CT/DE/CO + CA's 2026 rules) and works-council
  co-determination in DE/FR/NL make "privacy-proportionate by default" the only viable design.
- **Consolidation is the exit pattern.** Prompt Security → SentinelOne, Lakera → Check Point,
  Aware → Mimecast: the capture layer gets bought by platforms. Build the data/insight layer
  they can't easily replicate.
- **MVP verdict: feasible** with a small team in ~12 weeks against a design-partner motion —
  see [`08-mvp/01-mvp-plan.md`](./08-mvp/01-mvp-plan.md).

> **Caveat:** figures cited from third-party market reports vary widely by methodology; each
> doc links its sources. Nothing here is legal advice — validate the legal sections with
> counsel before any deployment that touches employee data.
