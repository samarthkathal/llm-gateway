# Competitive Landscape

## Map: four quadrants around the choke point

Every competitor touches enterprise AI traffic somewhere. Plotting by *what they extract from
the traffic* (nothing/risk-signals/meaning) vs. *who they serve* (engineering/security/business):

```
                        serves business/HR
                              ▲
      Workforce analytics    │     ★ INTENT GATEWAY
      (Viva, ActivTrak,      │     (org-aware meaning from
       Insightful, Aware†)   │      AI traffic, governed)
                             │
  weak proxies ──────────────┼────────────── rich signal
  (calendar, app usage)      │              (actual prompts)
                             │
      LLM gateways           │     GenAI security / DLP
      (LiteLLM, Portkey,     │     (Prompt Sec†, Lakera†,
       Kong, TrueFoundry)    │      Harmonic, Netskope,
                             │      Zscaler, Nightfall)
                              ▼
                        serves engineering/security
                                        † = acquired
```

## Quadrant 1 — LLM gateways (our substrate, their commodity)

LiteLLM (OSS, self-host, high volume), Portkey (managed, fast start, observability),
Kong AI Gateway (enterprise plugin ecosystem, fastest in Kong's own benchmark),
TrueFoundry (K8s-native, VPC, RBAC/governance), Helicone, OpenRouter.

- **Their game:** latency, provider coverage, cost dashboards, guardrail plugins.
- **Why they don't do what we do:** payloads are opaque; no identity-graph/org-hierarchy
  join; "governance" means rate limits and budgets, not intent policy.
- **Threat level: medium.** Any of them could add an "insights" tab; none owns HRIS-joined
  semantics. TrueFoundry/Kong are closest on enterprise governance features.
- Sources: [TrueFoundry landscape guide](https://www.truefoundry.com/blog/a-definitive-guide-to-ai-gateways-in-2026-competitive-landscape-comparison),
  [Kong benchmark](https://konghq.com/blog/engineering/ai-gateway-benchmark-kong-ai-gateway-portkey-litellm),
  [Particula decision framework](https://particula.tech/blog/ai-gateway-decision-litellm-portkey-kong-ai-gateway).

## Quadrant 2 — GenAI security / AI-usage DLP (validated, then absorbed)

- **Prompt Security** — $18M Series A (Nov 2024), **acquired by SentinelOne (Aug/Sep 2025)**
  ([AppSec Santa](https://appsecsanta.com/prompt-security), [Wikipedia](https://en.wikipedia.org/wiki/Prompt_Security)).
- **Lakera** — ~$30M raised, **acquired by Check Point (Sep 2025)**
  ([Tracxn](https://tracxn.com/d/companies/lakera/__bJHYPbXKymeh9hjUrJceMqDjX9dz9WwohC0zfAtjCyw)).
- **Harmonic Security** — independent, $26M total ($17.5M A led by Next47)
  ([PitchBook](https://pitchbook.com/profiles/company/538273-72), [Crunchbase](https://www.crunchbase.com/organization/harmonic-security)).
- Plus platform SSE incumbents: Netskope, Zscaler, Palo Alto adding GenAI app control.
- **Their game:** block/redact risky prompts, shadow-AI discovery, per-app policy.
- **Why they don't do what we do:** they classify traffic for *risk* and discard *meaning*;
  no org-hierarchy insight layer; buyer stops at CISO.
- **Threat level: high at the capture point.** SSE vendors can force-proxy all AI traffic
  network-wide. Our counter: they see network flows, not authenticated per-user API traffic
  with org context — and their DNA is blocking, not analytics.

## Quadrant 3 — Workforce analytics (the cautionary tale + the budget)

- **Microsoft Viva Insights** — bundled, calendar/collab-signal based; Copilot analytics
  expanding. The bundling threat.
- **ActivTrak / Insightful / Prodoscore** — activity-tracking heritage ("bossware" gravity)
  ([comparison roundup](https://reclaim.ai/blog/microsoft-viva-insights-alternatives)).
- **Aware** — NLP on internal comms (Slack/Teams), $60M Series C led by Goldman (2021),
  **acquired by Mimecast (Aug 2024)**
  ([PR Newswire](https://www.prnewswire.com/news-releases/aware-raises-60-million-in-series-c-funding-led-by-goldman-sachs-asset-management-301398993.html),
  [Mimecast](https://www.mimecast.com/resources/press-releases/mimecast-acquires-aware/)).
  Aware is the closest historical analog: semantic analysis of employee communications,
  sold for governance/insight — and it exited into a security platform.
- **Their game:** productivity/engagement metrics from behavioral proxies.
- **Why they don't do what we do:** no gateway position, no AI-traffic signal, no
  multi-provider LLM control plane.
- **Threat level: medium-low** as competitors; **high** as brand contamination — we must
  never be categorized with them. See positioning doc.

## Quadrant 4 — AI governance platforms (policy without traffic)

Credo AI, Holistic AI, OneTrust AI Governance, IBM watsonx.governance; observability-adjacent:
Arize, Arthur, Fiddler, Galileo.

- **Their game:** model inventories, risk assessments, EU AI Act conformity workflows.
- **Why they don't do what we do:** paperwork-first, not inline; no live usage capture.
- **Threat level: low-medium; high partnership potential** — we generate the live evidence
  their conformity workflows need.

## Structural read

1. **The capture layer consolidates into security platforms** (three acquisitions in ~14
   months). Standing still as "a smarter capture proxy" means competing with SentinelOne and
   Check Point features. The durable position is the **intent data model + org graph +
   insight layer** — the thing acquirers did *not* get in those deals.
2. **No incumbent owns the CIO+CISO+COO trifecta** from one deployment; each quadrant serves
   one buyer. The cross-quadrant product is the whitespace.
3. **Microsoft is the systemic risk** (Copilot analytics + Purview + Entra). The defensible
   segment is multi-provider enterprises — which is most large enterprises in practice
   (80%+ LLM adoption spans multiple providers;
   [index.dev](https://www.index.dev/blog/llm-enterprise-adoption-statistics)).
