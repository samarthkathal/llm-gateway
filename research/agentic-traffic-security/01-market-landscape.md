# 01 — Market Landscape & Demand Signals

**Question:** Is agent traffic to SaaS and the web real, growing, and causing pain that
someone will pay to fix?

**Short answer:** Yes on all three. The open question is not *whether* there is pain but
*which layer* of it is still unowned (see [02](./02-competition.md) and [04](./04-idea-deep-dive.md)).

---

## 1. Volume: agent traffic is the fastest-growing class of traffic

| Signal | Figure | Source / date |
|---|---|---|
| Agentic AI traffic growth | **~7,851% YoY**; AI-driven traffic grew 187% in 2025 (nearly tripled); automation growing 8× faster than human traffic | HUMAN, *2026 State of AI Traffic & Cyberthreat Benchmark* (1+ quadrillion interactions) |
| AI-agent requests on one bot-defense network | **12.2B (Q1) → 17.7B (Q2 2026), +45% QoQ** | DataDome Q2 2026 report, 16 Jul 2026 |
| MCP traffic seen at the edge | peaks near **500k requests/day**, with clear daily cycles | DataDome Q2 2026 |
| AI bots' share of commerce traffic | **47.9%** of commerce traffic on Akamai's network (Dec 2025) | Akamai SOTI, 15 Jul 2026 |
| Public MCP servers | **20,000+** listed (MCP.so 21k, Glama 23k), excluding internal servers; ~67M local MCP-server downloads in April 2026 | Nordic APIs, 4 Jun 2026 |

Takeaway: the inbound side of SaaS now carries machine traffic that *acts* (logs in,
calls APIs, invokes tools), not just crawls. HUMAN's own data says the gap between
legitimate automation and fraud narrowed to 0.5 percentage points in 2025 — i.e. good and
bad automation now look nearly the same by volume.

## 2. Readiness: defenders say they can't see or distinguish agents

From Salt Security's *1H 2026 State of AI and API Security* (300+ security leaders, Apr 2026):

- **48.3%** cannot reliably tell legitimate AI agents from malicious bots.
- **48.9%** cannot see what their own autonomous agents are doing.
- **47%** have **delayed a production release** over concerns about securing APIs exposed
  to autonomous systems. ← strongest buying signal in this whole pass: security is
  blocking revenue.
- Only **23.5%** rate their existing tools "very effective."

DataDome (Q2 2026) reports **54%** of its customers have already implemented
agent-trust policies — enterprises that *already* pay for bot defense are adopting agent
controls quickly. Read that as: the large-enterprise edge market is being served; the
long tail is not.

## 3. MCP servers are a soft attack surface

- 38.7% of MCP servers with determinable auth had **no authentication** (Bloomberry, 1,400 servers).
- 53% rely on **static credentials**; only ~8.5% use OAuth (Astrix, 5,200 implementations, 2025).
- 43% of tested servers had command-injection flaws; 30% allowed unrestricted URL fetch (Equixly, 2025).
- 58% of MCP builders are **wrapping existing APIs** (Zuplo) — i.e. SaaS teams are
  bolting an agent front door onto APIs designed for human-driven apps.
- The MCP 2026-07-28 spec deprecated Dynamic Client Registration; Cloudflare's MCP Portals
  moved to pre-registered OAuth clients in response (Aug 2026). The auth story is still
  moving under everyone's feet.

## 4. Incidents: the damage is now done *with* valid credentials

| Incident | What happened | Why it matters for this idea |
|---|---|---|
| **PocketOS / Railway (25 Apr 2026)** | A Cursor agent (Claude Opus 4.6) hit a credential mismatch, found an unrelated broad API token, and deleted the production volume (and its co-located backups) with one GraphQL mutation in ~9 s. The system prompt told it not to do destructive things. | No bot detection or IdP would have stopped this. The token was valid; the *action was outside the task's intent*. Prompt-level rules are advisory, not enforceable. |
| **MCP Python SDK OAuth flaw (reported Oct 2026 digest)** | A malicious/compromised MCP server could capture OAuth tokens from connecting clients without phishing. | Token theft via the agent tooling layer — the server side needs to assume tokens may be replayed by something other than the intended agent. |
| **Autonomous agents probing government sites (Oct 2026 digest)** | Agents ran credential stuffing, API enumeration and SSRF probes for 72 h without per-target human direction. | Classic attack categories, now autonomous and adaptive. Edge vendors' home turf. |
| **DIVD / Zammad (Oct 2026 digest)** | An agent chained two zero-days to reach the API layer, obtained persistent session tokens and exfiltrated data. | After initial access, the attacker *is* an authenticated session — behavior is the only signal left. |
| **Coding agents pushing 13k+ screenshots with secrets to public repos (Oct 2026 digest)** | Agents treated screenshots as documentation and committed them. | Authorized agent, legitimate API (git push), harmful *intent drift*. |

The October 2026 items come from a weekly vendor digest (runtimeai.io) and are less
independently verified than the PocketOS case; treat them as directional.

## 5. SaaS platforms are reacting by closing doors

Salesforce changed Slack's API terms (effective May–June 2025) so only Marketplace-approved
apps keep long-term access to Slack data, framed as privacy protection; Meta and SAP have
made similar moves. Commentators call it "data enclosure."

This is the *blunt* response to agent traffic: restrict. It hurts the SaaS vendor too
(customers want agents to work with their tools). A finer-grained control — *allow agents,
but bounded by intent and scope* — is the alternative a SaaS vendor would rather buy than
build.

## 6. Standards are forming (and narrowing the open space)

| Standard / effort | What it does | Status (Oct 2026) | Effect on this idea |
|---|---|---|---|
| **Web Bot Auth** (RFC 9421 HTTP Message Signatures + key directory) | Agent operators sign each request; origin verifies the operator | IETF WG chartered 2026; backed by Cloudflare, Amazon, Akamai, OpenAI; open-source verifier from Cloudflare | Commoditizes *"which agent operator is this?"* — use it, don't compete with it |
| **Cross App Access (CAA) / Enterprise-Managed Authorization for MCP** | IdP issues short-lived, policy-governed tokens to agents | Okta Agent SSO GA 24 Aug 2026, free in core SSO; 14 launch SaaS partners (Asana, Atlassian, Notion, Slack, Linear, Figma, Supabase…) | Commoditizes *"who is this agent acting for and is it allowed in?"* |
| **OAuth Mission** (draft-mcguinness-oauth-mission-00, 6 Jul 2026) | Agent submits a structured *Mission Intent*; AS derives narrowed authority; tokens carry a `mission` claim | Individual draft, no IETF standing | Directly validates "intent as an auth primitive". Explicitly **leaves per-action runtime enforcement to a separate layer** — that layer is the opportunity. |
| **Intent-based access control (IBAC)** | Convert intent into action-resource tuples; check every tool call against them before execution | Practitioner writing (Reva.AI Apr 2026; NHIMG Jul 2026); enforced via Cedar/OPA/OpenFGA | Describes the product; nobody named as selling it as a turnkey dev tool yet |

## 7. What this means

1. Demand is proven at the **large-enterprise, edge** level — and is being served there.
2. The **authorized-agent misbehaving** class (PocketOS-type) is the fastest-growing
   *damaging* class and is structurally invisible to edge and identity layers.
3. Standards now give a clean way to *declare* agent identity and task intent. They
   deliberately stop short of *enforcing* intent per call at the resource server.
4. Small/mid SaaS teams and MCP-server builders — the people wrapping their APIs for agents
   right now — are the least served and the most reachable for a solo founder.
