# 02 — Competition Analysis

**Question:** Who already protects SaaS (or the agents themselves) from agent traffic, and
where is the empty space?

**Short answer:** Five layers, all funded, three of them consolidating fast. The empty
space is a *developer-first, intent-aware enforcement point inside the SaaS/MCP server*
for teams below the six-figure-contract line.

---

## 1. The map

Plot every player by **where it sits on the request path** (agent side → network edge →
identity → API/MCP server) and **what it understands** (identity only → behavior → task intent).

```
                       understands TASK INTENT
                                 ▲
     Agent-side guardrails       │      ★ OPEN SPACE
     (Prompt Sec†, Lakera†,      │      intent-conformance at the
      Akto, Claude Code hooks,   │      SaaS API / MCP server,
      WitnessAI)                 │      dev-first, mid-market
                                 │      (partial: HUMAN, Wallarm)
  agent/user side ───────────────┼──────────────── SaaS / server side
                                 │
     Agent identity & authZ      │      Edge bot & agent trust
     (Okta Agent SSO, Auth0,     │      (Cloudflare, Akamai, DataDome,
      Arcade, Keycard, Descope,  │       HUMAN, Imperva, Kasada, Fastly)
      Stytch, WorkOS, Scalekit)  │      API security (Salt, Wallarm,
     MCP gateways (Runlayer,     │       Akamai/Noname, Cequence)
      MintMCP, Cloudflare Portals│
      Kong, TrueFoundry)         │
                                 ▼
                       understands IDENTITY only           † = acquired
```

## 2. Layer-by-layer

### Layer A — Edge bot & agent trust (SaaS side, network edge)

| Vendor | Agent-specific offering | Notes |
|---|---|---|
| **Cloudflare** | Web Bot Auth verification in Verified Bots; MCP traffic detection (`MCP-Protocol-Version` header → `is_mcp` selector), shadow-MCP detection, MCP Server Portals with OAuth (Aug 2026) | Authored Web Bot Auth. Free/cheap tiers reach the long tail. **Biggest bundling threat.** |
| **DataDome** | "Agent Trust" control plane; Priority Protect waiting room for agentic commerce (May 2026) | Forrester Wave Leader, Bot & Agent Trust Mgmt Q2 2026. 54% of customers run agent-trust policies. |
| **HUMAN Security** | AgenticTrust in HUMAN Sightline (Jul 2025): visibility into agents acting for users, pre/post-login; explicitly pitches *"intent behind traffic"*; open-source HUMAN Verified AI Agent; cryptographic verification for AWS Bedrock AgentCore Browser | **Closest to "intent" language** — but consumer-facing web/fraud, not B2B SaaS APIs. |
| **Akamai** | Bot Manager + API Security (via Noname acquisition, 2024); commerce-focused agentic-fraud research | Enterprise-only pricing. |
| Imperva (Thales), Kasada, F5, Fastly, Radware, Arkose, Netacea, Cequence | Varying agent features | Category defined by Forrester Q2 2026 as "managing trust across bots, AI agents, LLM crawlers…" |

**Why not compete here:** global network data is the moat; Cloudflare gives much of it
away. **Use** their signals (signature verification) as inputs.

### Layer B — Agent identity & authorization (between agent and SaaS)

| Vendor | What | Funding / status |
|---|---|---|
| **Okta** | Agent SSO (CAA) GA Aug 2026, **free in core SSO**; Okta for AI Agents (GA May 2026); bought **Permiso** (closed 26 Aug 2026) | Category-defining incumbent |
| **Arcade** | Agent authorization: which app features an agent may use, IdP-synced, action logs; 8,000+ MCP tools | $60M Series A (Jun 2026, SYN Ventures, Morgan Stanley, Wipro) |
| **Keycard** | Identity & access for AI agents | $38M at launch (Oct 2025) |
| **Descope** | Agentic identity hub, MCP auth | $35M seed extension (reported) |
| **Stytch** | Connected Apps / MCP auth; free **IsAgent** detector (UA + TLS fingerprint heuristics, Aug 2025) for analytics and routing agents to agent-friendly flows | — |
| WorkOS, Scalekit, Auth0 (Okta) | MCP-server auth kits for SaaS builders | Developer-first; **these are the closest go-to-market analogs** to the recommended wedge |

**Why not compete here:** identity is being bundled into IdPs at zero marginal price.
**Integrate** with CAA/OAuth tokens instead.

### Layer C — MCP gateways & MCP security

Runlayer ($11M seed, Nov 2025), Helmet Security ($9M seed), Operant AI (pivoted to MCP),
MintMCP (Okta CAA launch partner), Lunar.dev, Docker MCP Gateway, Kong, TrueFoundry,
Cloudflare MCP Portals. Total disclosed MCP-security funding was only **~$40M across four
startups** as of March 2026 — small versus the attack surface.

Mostly deployed **by the enterprise consuming MCP** (governing which servers employees'
agents can reach), not by the **SaaS vendor providing** the server.

### Layer D — API security vendors adding agent features

| Vendor | Agentic feature |
|---|---|
| **Wallarm** | Agentic AI Protection: MCP tool-call schema conformance, per-caller access policies, **session analysis of call sequences** to catch legitimate tools used out of sequence, shadow-agent discovery, AI payload inspection |
| Salt Security | Agentic posture & discovery; published the readiness survey |
| Akamai API Security (Noname), Cequence, Harness (Traceable) | API discovery + abuse detection extended to agents |

**Wallarm is the single most direct competitor** to the recommended wedge: session-sequence
analysis is a behavioral proxy for intent. Difference: it infers anomalies; it does not
take a *declared* intent and check conformance, and it sells into enterprise security teams
via a filtering-node deployment.

### Layer E — Agent-side runtime / posture (enterprise deploying agents)

Zenity, Noma ($100M Series B, Jul 2025), Geordie AI (RSAC 2026 Innovation Sandbox winner),
WitnessAI ($58M Series B, Jan 2026), Akto (Claude Code guardrails), Koi (agent/extension
inventory → Palo Alto). Coding-agent vendors ship their own hooks/permission systems.

This is the **"user end"** option in your question. It is crowded and increasingly built
into the agents themselves.

## 3. Consolidation scoreboard (signals what acquirers value)

| Date | Acquirer → Target | Layer | Price (where reported) |
|---|---|---|---|
| Jul 2025 | Palo Alto → Protect AI | Model security | $634.5M (10-K) |
| 2025 | SentinelOne → Prompt Security | Agent-side runtime | undisclosed |
| 2025 | Check Point → Lakera | Agent-side guardrails | undisclosed |
| 2025 | Snyk → Invariant Labs | MCP runtime | undisclosed |
| 2025 | F5 → CalypsoAI | Guardrails at traffic layer | ~$180M reported |
| Sep 2025 | CrowdStrike → Pangea | AI guardrails | undisclosed |
| Sep 2025 | Cato → Aim Security | AI usage security | undisclosed |
| — | CrowdStrike → SGNL | Continuous authZ | $627.9M (10-Q) |
| — | Palo Alto → **Portkey** | **AI gateway** (was a competitor in your original `docs/`) | $117M (10-K) |
| — | Palo Alto → Koi | Agent/extension inventory | $231M (10-K) |
| Feb 2026 | Palo Alto → CyberArk | Machine identity | ~$21–25B |
| 2026 | Proofpoint → Acuvity | Agent posture | undisclosed |
| Jun 2026 | Cisco → Astrix | NHI / agent identity | ~$400M reported |
| Jun 2026 | SailPoint → Entro | NHI | ~$122.6M |
| Aug 2026 | Okta → Permiso | Identity threat detection | undisclosed |
| Sep 2026 | Cyera → Oasis Security | NHI / agentic access | "valued at $1B" |
| 2026 | OpenAI → Promptfoo | Evals | pending |

**Reading:** small, focused agent-security teams get bought at $100M–$1B within 1–3 years.
Exit path exists. It also means every capability you build has a large buyer *and* a large
incumbent that may ship it as a feature.

## 4. Gap analysis

| Capability | Edge (A) | Identity (B) | MCP GW (C) | API sec (D) | Agent-side (E) | **Wedge** |
|---|---|---|---|---|---|---|
| Is it an agent? Which operator? | ✅ | partial | partial | partial | n/a | consume Web Bot Auth / IsAgent-style signals |
| Who does it act for; token valid? | ❌ | ✅ | ✅ | partial | n/a | consume CAA/OAuth claims |
| Per-call check against **declared** task intent | ❌ | ❌ (stops at token) | partial (tool allowlists) | ❌ (anomaly only) | partial (client side, bypassable) | **✅ core** |
| Step-up/human approval for destructive ops at the server | ❌ | partial | partial | ❌ | client-side only | **✅** |
| Product insight: what agents try to do on *my* SaaS | ❌ | ❌ | ❌ | ❌ | ❌ | **✅ (Intent Gateway DNA)** |
| Self-serve, usage-priced, SDK in app code | partial (Cloudflare) | ✅ (WorkOS/Stytch) | partial | ❌ | ❌ | **✅** |

## 5. Competitive verdict

- **Do not build:** bot detection, agent identity, MCP gateway, or client-side agent firewall.
  Each is either free-in-a-bundle or owned by a funded leader.
- **Can build:** a thin, server-side *intent-conformance and step-up layer* that consumes
  everyone else's signals, plus an analytics view SaaS product teams actually want.
- **Watch closely:** Wallarm (session sequences), HUMAN (intent language), Cloudflare
  (could add "MCP intent policies" to Portals), Arcade/Okta (could push action-level policy
  deeper into the resource server via CAA). Any of them shipping *declared-intent*
  enforcement for the mid-market is a kill trigger (see [07](./07-risks.md)).
