# Intent Gateway

**A multi-tenant enterprise AI gateway that turns workforce AI usage into governed,
org-aware insight.**

Enterprises route their employees' LLM traffic (Claude, OpenAI, and others) through the
gateway. Every sanctioned request is enriched with identity and org-hierarchy context,
passed through per-tenant policy (redaction, retention, jurisdiction modes), distilled into
a structured **intent signal** — *what is this person actually working on?* — and rolled up
into aggregate, k-anonymized organizational insights. One deployment serves four budget
owners: platform (routing/failover/cost), security (AI-usage governance), compliance
(audit/DPIA/DSAR), and leadership (where AI effort actually goes).

**Deliberately not surveillance:** aggregate-first insights with a hard k-anonymity floor,
no individual scoring, raw text discarded after extraction by default, and an
employee-facing transparency page in v1. That posture is a legal requirement in our target
markets and the core of the product's differentiation — see the research below.

## Status

🚧 **Pre-MVP.** Pre-ideation research is complete; the build plan is set. The current code
is a prototype Go LLM gateway (provider abstraction + failover skeleton) that becomes the
capture layer in the MVP plan.

## Research & planning docs

The full pre-ideation knowledge base lives in [`docs/`](./docs/README.md):

| | |
|---|---|
| [Market analysis](./docs/01-market/) | What market this is, TAM/SAM/SOM, target segments, adjacent metrics movements (DORA/SEI) |
| [Competition](./docs/02-competition/) | Landscape map, ten competitor profiles, differentiation thesis |
| [Legal & regulatory](./docs/03-legal/) | EU AI Act + GDPR, US state laws, works councils, sector rules |
| [Effort analysis](./docs/04-effort/) | Region-entry effort matrix, GTM effort plan |
| [Strategy](./docs/05-strategy/) | Positioning, moat, product strategy & principles |
| [Risk](./docs/06-risk/) | Scored risk register with kill criteria |
| [Financial](./docs/07-financial/) | Funding landscape/comparables, indicative financial model |
| [MVP plan](./docs/08-mvp/) | The 12-week build plan and acceptance criteria |

Start with the [docs index](./docs/README.md) for the TL;DR.

## Repository layout

```
main.go                      # prototype entry point (failover simulation)
internal/
  llmgateway/                # gateway core: provider selection & failover (prototype)
  llmprovider/               # provider interface + Claude/OpenAI mock adapters
  common/                    # shared request/response types
docs/                        # pre-ideation research & MVP plan (see above)
```

The MVP restructures this into `cmd/gateway`, `cmd/api`, `cmd/worker` per
[`docs/08-mvp/01-mvp-plan.md`](./docs/08-mvp/01-mvp-plan.md).

## Quick start (prototype)

```bash
go run .   # runs the failover simulation against mock providers
```

## License & disclaimer

Private, all rights reserved (pre-incorporation).
The legal research in `docs/03-legal/` is planning material, not legal advice.
