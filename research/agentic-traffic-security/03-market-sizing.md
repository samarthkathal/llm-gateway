# 03 — Market Sizing

**Question:** How big is the pie, and how big is the slice a bootstrapped solo founder can
realistically take?

**Short answer:** Reference markets are large and growing ~30%/yr, but the honest
bottom-up number for the recommended wedge is a **$15–90M SAM** and a **$0.3–0.6M ARR
three-year SOM** — a good bootstrap business, not (yet) a venture-scale one on its own.

> All bottom-up figures below are **assumptions**, marked ⚠︎. Top-down figures are
> third-party estimates that vary widely by methodology.

---

## 1. Top-down reference markets

| Market | Size | Growth | Source | Relevance |
|---|---|---|---|---|
| API security | **$1.3B (2025) → $8.6B (2032)** | 31.2% CAGR 2026–32 | P&S Intelligence | Closest budget line: SaaS teams protecting APIs agents call |
| AI cybersecurity (broad) | ~$26B (2025) → **$172B by 2029** | 73.9% CAGR | Gartner, as cited by Software Strategies blog (secondary) | Ceiling, not addressable |
| Agentic-AI security sub-category | "hundreds of millions, not billions" today | — | Same blog's author estimate | Sanity check: the specific niche is young |
| Bot management | Multi-billion, ~14 vendors tracked; US growth "below average" vs emerging regions | — | QKS Group (paid report; figures not public) | Adjacent; edge players own it |
| Disclosed MCP-security funding | ~$40M across 4 startups (Mar 2026) | — | Software Strategies blog | Capital hasn't flooded the server-side MCP niche yet |

**Do not quote these as your TAM in a pitch.** They are context: the budget lines exist
and are growing; the specific sub-niche is early.

## 2. Bottom-up: who would buy the wedge

Target customer (from [04](./04-idea-deep-dive.md)): a B2B SaaS or dev-tool company,
roughly 10–500 employees, that **exposes an API and/or MCP server that agents call on
behalf of its users**, and that is below the threshold where Akamai/Salt/Wallarm
enterprise contracts make sense.

| Step | Estimate | Basis |
|---|---|---|
| Public MCP servers | 20,000+ | MCP.so / Glama counts (Nordic APIs, Jun 2026) |
| …that wrap an existing product API | ~58% → ~11,600 | Zuplo State of MCP |
| …that are run by a commercial SaaS/dev-tool company (not hobby) | ⚠︎ 10–20% → **~1,200–2,300** | assumption |
| B2B SaaS with public APIs that agents already call but have no MCP server yet | ⚠︎ ~3,000–4,000 in the size band | assumption (CLI/HTTP agents call REST directly) |
| **Addressable companies (SAM count)** | **⚠︎ ~5,000** | sum, rounded |
| ACV | ⚠︎ **$3k–$18k/yr** ($250–$1,500/mo, usage-based) | priced between Arcjet/WorkOS-style dev tools and enterprise API security |
| **SAM** | **⚠︎ $15M–$90M/yr** | 5,000 × $3k–$18k |

Expansion vectors that would lift SAM later (not counted):
- Agent-side companies (the *consumers*) wanting the same intent engine outbound → links
  back to the original Intent Gateway.
- Enterprises' internal APIs exposed to internal agents (larger ACV, slower sales).
- Per-request pricing growth as agent traffic compounds (DataDome saw +45% QoQ).

## 3. SOM: what a solo bootstrapper can capture

| Horizon | Customers | Blended ACV | ARR | Notes |
|---|---|---|---|---|
| Month 6 | 5–10 (design partners → paid) | ~$2k | $10–20k | mostly discounted pilots |
| Year 1 | 25–40 | ~$4k | $100–160k | self-serve + founder-led sales |
| Year 3 | 60–100 (1–2% of SAM) | ~$6k | **$360–600k** | sustainable solo/2-person business |

**Bootstrap break-even** (Bengaluru cost base, solo, infra included): ⚠︎ roughly
$60–100k ARR, i.e. 15–25 paying customers. That is the number to beat before considering
outside capital.

## 4. What would make this venture-scale

Only if one of these proves true in the first 12 months:
1. Intent-conformance becomes a **checkbox** SaaS vendors must show enterprise buyers
   ("how do you stop agents doing things users didn't approve?") — turns a nice-to-have
   into a sales requirement, like SOC 2.
2. The **product-insight** half (what agents try to do on my product) becomes the main
   value, sold to product/growth teams with bigger budgets than security.
3. An acquirer pattern repeats (see [02 §3](./02-competition.md)) — a $5–15M-ARR company
   in this niche is a plausible $100M+ acquisition.

## 5. Bottom line

- Market is **real and growing**, but the reachable slice is **small-to-mid**.
- The numbers support **bootstrapping** to ~$0.5M ARR; they don't yet support a
  VC-style "$1B TAM" story without the insight half or enterprise expansion.
- Validate the ACV assumption first: it's the most sensitive input (6× range).
