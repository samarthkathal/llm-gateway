# Market Overview — What Market Is This, Actually?

## The idea in one paragraph

An enterprise-deployed, multi-tenant **LLM gateway** that sits inline on employees' AI traffic,
enriches every request with identity + org-hierarchy context (local users, SCIM, Workday),
extracts a structured **intent signal** ("what is this person actually working on?"), stores it
under per-tenant policy/governance controls, and rolls it up into **aggregate organizational
insights** — where effort is going, where AI is amplifying work, where hidden capacity sits.

## This product sits at the intersection of four existing markets

None of them alone is the product; the overlap is the opportunity.

### 1. LLM / AI gateways (infrastructure)
Routing, failover, cost control, observability for enterprise LLM traffic. Rapidly growing
(~42% of enterprises already run a middleware layer; the segment is forecast to grow at
~50% CAGR) but **rapidly commoditizing** — LiteLLM is open source, Portkey/Kong/TrueFoundry
compete on performance and price. The gateway is our *distribution point*, not our moat.

### 2. GenAI security / AI-usage DLP (security)
Prompt-level inspection, redaction, shadow-AI discovery. Validated hard by acquirers in
2025: Prompt Security → SentinelOne, Lakera → Check Point; Harmonic Security remains
independent ($26M raised). This proves enterprises will deploy inline prompt inspection —
the exact capture mechanism we need — but these products stop at *risk*; they throw away
the *meaning* of the traffic.

### 3. Workforce analytics / people analytics (HR tech)
Understanding how work happens: Viva Insights, ActivTrak, Insightful, Aware (→ Mimecast).
Market ~$2.4–4.5B in 2025–26 growing ~13% CAGR; employee-monitoring software specifically
~$4–4.9B growing ~18%. Established buyer (CHRO/COO) but reputationally loaded
("bossware") and legally the heaviest quadrant.

### 4. AI governance platforms (compliance)
Policy, audit, conformity with EU AI Act & friends: Credo AI, Holistic AI, OneTrust modules.
Small today (~$0.4–0.75B in 2026 depending on source) but the fastest-growing of the four
(~35–45% CAGR forecasts; Gartner sees >$1B by 2030). Regulation-driven budget that did not
exist three years ago.

## The synthesis: "Intent intelligence"

The unique claim: **AI traffic is the richest real-time signal of what a workforce is
actually doing that any enterprise has ever had.** Nobody currently joins that signal to the
org chart and turns it into decision-grade insight:

- Gateways see the traffic but treat payloads as opaque strings to route.
- AI-security tools parse the payloads but only to find secrets/PII, then discard semantics.
- Workforce analytics infers work from calendars/app-usage — weak proxies vs. actual prompts.
- Governance platforms manage model inventories and policies, not live usage meaning.

An intent gateway captures once at the choke point and serves four budget lines
(infra, security, HR insight, compliance) from one deployment.

## Why now (timing thesis)

1. **Enterprise LLM adoption crossed ~80% in 2026** — the traffic exists at scale for the
   first time. ([index.dev statistics roundup](https://www.index.dev/blog/llm-enterprise-adoption-statistics))
2. **Shadow AI is a named board-level risk** — CISOs are actively buying visibility into AI
   usage ([Fortinet](https://www.fortinet.com/blog/secure-networking/shadow-ai-the-invisible-risk-growing-inside-your-organization),
   [CIO.com](https://www.cio.com/article/4178708/shadow-ai-the-hidden-risk-expanding-across-the-enterprise.html)).
3. **EU AI Act high-risk (Annex III) obligations land ~Dec 2027** (postponed from Aug 2026) —
   a compliance clock that makes governed AI usage mandatory, with a ~17-month runway that is
   exactly an early-stage product cycle ([McCann FitzGerald](https://www.mccannfitzgerald.com/knowledge/technology/employment-spotlight-eu-ai-act-draft-guidelines-on-high-risk-ai-classification),
   [Crowell & Moring](https://www.crowell.com/en/insights/client-alerts/artificial-intelligence-and-human-resources-in-the-eu-a-2026-legal-overview)).
4. **CFOs are demanding AI ROI evidence** — after two years of spend, "what are we getting
   from AI?" has no good answer today; intent-level telemetry is that answer.
5. **The capture layer just got validated and absorbed** — security acquirers bought the
   inline-inspection startups, leaving the *insight* layer open.

## What this market is NOT

- **Not surveillance/productivity-scoring.** Individual-level "productivity scores" are the
  failure mode: legally high-risk (EU AI Act Annex III area 4, works-council veto territory)
  and the fastest way to lose employee trust and enterprise deals. See
  [`../03-legal/`](../03-legal/) and the positioning doc in
  [`../05-strategy/01-positioning-and-strategy.md`](../05-strategy/01-positioning-and-strategy.md).
- **Not another model router.** Routing is table stakes we inherit, not the pitch.

## Sources

- [Fortune Business Insights — Enterprise LLM market](https://www.fortunebusinessinsights.com/enterprise-llm-market-114178)
- [Future Market Insights — Enterprise LLM market](https://www.futuremarketinsights.com/reports/enterprise-llm-market)
- [index.dev — LLM enterprise adoption statistics 2026](https://www.index.dev/blog/llm-enterprise-adoption-statistics)
- [Fortune Business Insights — Workforce analytics market](https://www.fortunebusinessinsights.com/industry-reports/workforce-analytics-market-100299)
- [The Business Research Company — Employee monitoring software market](https://www.thebusinessresearchcompany.com/report/employee-monitoring-software-global-market-report)
- [Grand View Research — AI governance market](https://www.grandviewresearch.com/industry-analysis/ai-governance-market-report)
- [Nemko Digital — Gartner: AI governance platforms >$1B by 2030](https://digital.nemko.com/news/ai-governance-platforms-market-to-surpass-1-billion-by-2030)
- [AppSec Santa — Prompt Security acquired by SentinelOne](https://appsecsanta.com/prompt-security)
- [Tracxn — Lakera (acquired by Check Point)](https://tracxn.com/d/companies/lakera/__bJHYPbXKymeh9hjUrJceMqDjX9dz9WwohC0zfAtjCyw)
- [Mimecast — Aware acquisition](https://www.mimecast.com/resources/press-releases/mimecast-acquires-aware/)
