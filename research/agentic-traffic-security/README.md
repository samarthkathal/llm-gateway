# Agentic Traffic Security — the security angle on intent capture

> Research pass compiled **13 October 2026** by Claude, as a follow-on to the
> [Intent Gateway research in `docs/`](../../docs/README.md) (July 2026).
> This folder is additive: nothing under `docs/` or the Go code was changed.

## The question

The original Intent Gateway captures *outbound* intent — what employees are trying to do
with LLMs. This pass asks the mirror question:

> SaaS products were hardened for humans and for classic bots (DDoS, scraping, credential
> stuffing) by Cloudflare/Akamai/CrowdStrike-style layers. Now **agents** — Claude Code,
> Codex, Cursor, browser agents, MCP clients, CLIs — are calling SaaS APIs and web apps
> on users' behalf. **Is there a real, fundable need for a security engine that
> understands agent traffic and its intent, and where should it sit — user end,
> middleware, or SaaS/enterprise end?**

## Documents

| # | Doc | What it answers |
|---|---|---|
| 01 | [Market landscape & demand signals](./01-market-landscape.md) | Is agent traffic real, growing, and hurting anyone? What standards are forming? |
| 02 | [Competition](./02-competition.md) | Who already plays here — five layers, ~40 companies, funding and M&A |
| 03 | [Market sizing](./03-market-sizing.md) | Top-down reference markets + a bottom-up, bootstrap-scale SOM |
| 04 | [Idea deep-dive & positioning](./04-idea-deep-dive.md) | Threat model, the actual gap, user vs middleware vs SaaS end, recommended wedge |
| 05 | [Intent-model economics](./05-intent-model-economics.md) | Jev, OpenAI Decisions API, open models — is per-request intent inference affordable, and should you fine-tune? |
| 06 | [MVP & PMF plan](./06-mvp-and-pmf-plan.md) | What to build in 8 weeks solo, how to test fit, pricing, kill criteria |
| 07 | [Risk register](./07-risks.md) | What kills this, scored |
| 08 | [Sources](./08-sources.md) | Every link used |

## TL;DR

1. **The problem is real and measured.** Agentic traffic is growing in the thousands of
   percent YoY (HUMAN), 45% quarter-on-quarter on DataDome's network, and roughly half of
   security leaders say they can't tell good agents from bad bots or see what their agents
   are doing (Salt, 300+ leaders). Real incidents are now *authorized* agents doing damage
   with valid credentials (PocketOS/Railway, April 2026), not just bots breaking in.
2. **Your thesis is directionally right, but the market noticed first.** Forrester created
   a "Bot and Agent Trust Management" Wave in Q2 2026. Okta made Agent SSO free inside core
   SSO (Aug 2026). Cloudflare detects MCP traffic natively. Wallarm ships MCP schema and
   session-sequence enforcement. Acquirers (Cisco, Okta, Cyera, Palo Alto, CrowdStrike,
   SailPoint) spent billions on agent-identity and agent-security startups in 2026.
3. **The edge (is this an agent?) and identity (who is it, what token?) layers are taken.**
   A solo founder cannot win either.
4. **The gap that remains is narrow but real:** *authenticated agents behaving outside the
   task the user approved*, seen from the SaaS/MCP-server side. Identity vendors issue the
   token and step away; edge vendors don't see inside authenticated sessions; API-security
   vendors sell to large enterprises with six-figure contracts. Small and mid-market SaaS
   and MCP-server builders have nothing developer-friendly that answers *"does this
   sequence of calls match what the agent said it was doing?"*
5. **This is where "intent" stops being a buzzword and becomes a security primitive.**
   IETF drafts (OAuth Mission, July 2026) and intent-based access control writing are
   converging on "declare the task, then enforce against it." Cheap decision models (Jev,
   OpenAI Decisions API, open guard models) make a per-request intent-conformance check
   affordable for the first time.
6. **Recommended wedge:** a developer-first **agent-aware middleware/SDK for SaaS APIs and
   MCP servers** ("Arcjet for agents"): identify the agent, bind each session to a declared
   intent, check calls against it, step-up destructive actions to a human, and give the
   SaaS product team an anonymized view of *what agents are trying to do on their
   product*. That last piece reuses the Intent Gateway's insight DNA.
7. **Verdict: conditional go.** Worth an 8-week validation sprint, not a full build.
   Proceed only if ≥5 of 15 SaaS/MCP builders interviewed describe agent-caused incidents
   or blocked launches *and* ≥2 agree to run a design-partner pilot. See doc 06 for kill
   criteria. Don't fine-tune your own model until you have labeled production traffic
   (doc 05).

## How this relates to the original Intent Gateway

| | Intent Gateway (docs/) | Agentic Traffic Security (this folder) |
|---|---|---|
| Traffic direction | Outbound: employees → LLMs | Inbound: agents → SaaS APIs / MCP servers |
| Buyer | Enterprise CISO/CIO, then COO/CFO | SaaS CTO / platform / security engineer; MCP-server builders |
| Intent used for | Org-level insight, governance | Authorization conformance + product insight |
| Legal weight | Heavy (employee data, works councils) | Lighter (machine traffic; still personal data in payloads) |
| Solo-founder fit | Low–medium (enterprise sales, compliance artifacts) | Medium–high (bottom-up, self-serve, usage pricing) |

The shared asset is the **intent extraction engine**. Built once, it can serve both
directions later — but start with one.

> Nothing here is legal or investment advice. Third-party market figures vary widely by
> methodology and several 2026 sources are vendor reports; each claim links its source in
> [08-sources.md](./08-sources.md).
