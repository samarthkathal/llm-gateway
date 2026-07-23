# Risk Register

Scoring: Likelihood (L) and Impact (I) 1–5; Exposure = L×I. Reviewed quarterly; owners
assigned once the team exists. Top 5 exposures in bold.

| # | Risk | L | I | Exp | Mitigation | Early-warning signal |
|---|---|---|---|---|---|---|
| R1 | **Category contamination: product perceived as "bossware", poisoning deals and press** | 4 | 5 | **20** | Aggregate-first architecture as *visible* design (transparency page, k-floor); never use monitoring vocabulary; publish privacy architecture openly | Prospects route us to HR-tech procurement track; "monitoring" appears in our own sales decks |
| R2 | **Platform bundling: Microsoft (Copilot analytics+Purview) or SSE vendors ship "good-enough" AI-usage insight free** | 4 | 4 | **16** | Multi-provider neutrality; depth of intent taxonomy; move fast on org-graph join they lack; SEI integrations widen the moat | Copilot dashboard roadmap announcements; deals stalling on "we'll wait for Microsoft" |
| R3 | **Intent-extraction quality insufficient — insights feel generic, insight tier doesn't attach** | 3 | 5 | **15** | Treat taxonomy as the kernel: eval harness from week 1, design-partner feedback loop, human-in-the-loop corrections | Design partners view dashboards <1×/week; execs can't quote a finding |
| R4 | **Legal misstep in one deployment (works council, GDPR complaint) becomes public** | 2 | 5 | 10 | Jurisdiction modes enforced in code, not contract; counsel review per new country; refuse out-of-posture tenant requests | Tenant asks to disable k-floor "just for one team" |
| R5 | **Cross-tenant data leak** (existential for a trust product) | 2 | 5 | 10 | DB-layer RLS + per-tenant keys + isolation tests in CI + pen test before GA | Any RLS bypass found internally |
| R6 | **The intersection bet fails — buyers want point tools, not the double story** | 3 | 4 | 12 | Tiering lets each budget buy alone; watch attach-rate metric; be ready to collapse to governance-only | Insight-tier attach <20% at month 12 |
| R7 | **EU AI Act slips again / enforcement is toothless → compliance urgency evaporates** | 3 | 3 | 9 | Don't build the company on the deadline; security + ROI stories carry US revenue regardless | Another postponement announcement |
| R8 | **LLM providers add native enterprise usage-analytics** (OpenAI/Anthropic admin consoles) | 3 | 3 | 9 | Cross-provider + org-graph + policy is beyond a provider console's scope/neutrality | Provider console ships org-chart integration |
| R9 | **Gateway operational failure (latency/outage) burns infra credibility** | 2 | 4 | 8 | Async analysis off hot path; passthrough mode on component failure; SLOs from day 1 | p50 overhead >15ms; any full outage |
| R10 | **Funding winter for non-AI-core infra; seed round takes too long** | 2 | 4 | 8 | Keep burn ≤2–3 people until revenue; design-partner LOIs before raise; security-budget revenue is countercyclical | <20% investor-meeting → diligence conversion |
| R11 | **Workday/HRIS integration harder than planned, delaying org-graph value** | 3 | 2 | 6 | SCIM+CSV first (already the plan); Workday only for a paying enterprise | Connector eng-weeks >2× estimate |
| R12 | **Key-person dependency (solo-founder-shaped project)** | 3 | 3 | 9 | Documentation-first culture (this repo); early technical hire; scope MVP honestly | Bus-factor complaints from design partners |
| R13 | **Employee backlash at a tenant (walkouts/press over AI monitoring)** | 2 | 4 | 8 | Employee transparency page + genuinely aggregate insights are the defense; provide tenant comms kit for rollout | Tenant deploys silently without employee comms |
| R14 | **Prompt data poisoning / gaming once employees know intent is measured** | 2 | 2 | 4 | Aggregate use dulls incentive to game; no individual stakes attached (by design) | Anomalous taxonomy distributions per team |

## Kill criteria (agreed in advance, so pivots are decisions, not drifts)

- Month 6: <3 design partners deployed → re-examine wedge (governance-only pivot).
- Month 12: <$250k ARR **and** insight attach <20% → collapse to best-selling tier or wind down.
- Any confirmed cross-tenant leak of intent data → full-stop incident: disclose, remediate,
  and reassess viability honestly — trust products rarely survive a second one.

## Insurance & structural mitigations

- Tech E&O + cyber liability before first production tenant.
- Contractual: liability caps, tenant indemnity for *their* deployment-law compliance
  (we ship the tools; the employer owns employment-law duties as deployer).
- Data minimization as blast-radius control: what we don't store can't leak (raw-text TTL).
