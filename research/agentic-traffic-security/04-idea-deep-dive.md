# 04 — Idea Deep-Dive & Positioning

**Question:** What exactly is broken, where should the security engine sit (user end,
middleware, or SaaS/enterprise end), and what is the defensible wedge?

---

## 1. Testing the macro thesis

Your thesis: *the "SaaS apocalypse" didn't happen; enterprises keep buying SaaS rather
than vibe-coding replacements (security and liability), and instead point agents at
existing SaaS. So SaaS becomes the agent's toolbox — and SaaS isn't hardened for that.*

| Claim | Evidence | Verdict |
|---|---|---|
| Agents increasingly operate *through* existing SaaS | 58% of MCP builders wrap existing APIs; Okta's CAA launch partners are mainstream SaaS (Asana, Atlassian, Notion, Slack, Linear, Figma); 20k+ public MCP servers | **Supported** |
| SaaS vendors aren't ready | 48% can't tell good agents from bad; 47% delayed releases over agent-exposed APIs; most MCP servers use static creds or none | **Supported** |
| Enterprises prefer buying over self-building | Plausible and consistent with incident record (PocketOS-type failures), but no hard data found in this pass | **Plausible, unproven** — don't lean on it in pitches |

The part of the thesis that matters for the product holds: **SaaS APIs are becoming the
agent front door, and their security model assumes a human is driving the client.**

## 2. Threat model: four classes of agent traffic

| Class | Example | Who catches it today |
|---|---|---|
| **T1. Unauthenticated hostile automation** | Scrapers, credential stuffing, API enumeration by autonomous agents | Edge vendors (Cloudflare, DataDome, Akamai, HUMAN). **Solved enough.** |
| **T2. Impersonation** | Bot claiming to be ChatGPT-User / Claude | Web Bot Auth signatures, HUMAN Verified Agent. **Being solved by standards.** |
| **T3. Unauthorized token use** | Stolen/replayed OAuth token; malicious MCP server capturing tokens | IdPs (short-lived tokens via CAA), ITDR (Permiso→Okta). **Partially solved.** |
| **T4. Authorized agent, out-of-intent action** | PocketOS: valid token, agent deletes prod volume while "fixing credentials"; agents committing screenshots with secrets; prompt-injected agent exfiltrating via legit API calls | **Nobody, at the server, for the mid-market.** ← the gap |

T4 is the important one: it grows with agent *adoption*, not with attacker activity. Every
customer that successfully onboards agents increases their T4 exposure. Identity says
"yes, allowed"; the edge says "yes, a known agent"; the damage happens anyway.

## 3. Where should the engine sit?

| | **User end** (agent-side firewall: CLI wrapper, local proxy, IDE plugin) | **Middleware** (gateway between agents and SaaS, run by enterprise) | **SaaS end** (SDK/middleware inside the SaaS API / MCP server) |
|---|---|---|---|
| Sees user's actual prompt/intent | ✅ best | partial | ❌ must be declared or inferred |
| Can enforce (not bypassable) | ❌ user/agent can bypass | ✅ for that enterprise only | ✅ for all callers of that SaaS |
| Who pays | Individual devs / enterprise IT | Enterprise security | SaaS vendor (security/platform/product) |
| Competitive density | Very high (built into Claude Code/Codex/Cursor; Prompt Sec, Akto, WitnessAI) | High (MCP gateways, SSE vendors, Okta, Cloudflare) | **Low for mid-market** (Wallarm is enterprise) |
| Distribution for solo dev | Hard to monetize (devs expect free) | Enterprise sales cycles | **Bottom-up: SDK, npm/Go module, free tier** |
| Network effect | Weak | Per-enterprise | **Cross-tenant agent behavior patterns** |
| **Fit** | ❌ | ⚠︎ later | ✅ **start here** |

**Recommendation: SaaS end, as middleware/SDK.** It is the only position that is
enforceable, under-served in the mid-market, and sellable bottom-up by one person.

The one real weakness of the SaaS end — it doesn't see the user's prompt — is handled in §5.

## 4. The product: "agent-aware request layer" for SaaS APIs and MCP servers

Working name placeholder: **Intent Guard** (rename later).

```
 Agent (Claude Code / Codex / browser agent / MCP client)
   │  HTTP or MCP request
   │  + optional: Web Bot Auth signature, CAA/OAuth token, mission claim / intent header
   ▼
┌──────────────────────── SaaS API / MCP server ────────────────────────┐
│  Intent Guard middleware (Go / Node / Python SDK, or sidecar)          │
│   1. Classify caller  → human | known agent (verified) | unknown agent │
│   2. Resolve intent   → declared (mission claim, header, MCP session)  │
│                         or inferred (first N calls + endpoint semantics)│
│   3. Check conformance→ fast decision model: in-intent | drift | unsafe │
│   4. Act              → allow | rate-limit | require step-up (human     │
│                         approval link) | block | log-only               │
│   5. Emit event       → redacted, aggregate-first                       │
└────────────────────────────────────────────────────────────────────────┘
   ▼                                  ▼
 App handler                    Intent Guard cloud
                                 • policy (data, versioned)
                                 • approvals inbox (Slack/email)
                                 • "Agent insights": what agents try to do
                                   on your product, success/failure by intent
```

### Core capabilities (MVP → later)

1. **Agent classification** — consume, don't build: Web Bot Auth verification, known
   agent UA/TLS fingerprints, MCP headers, token claims (agent vs user client).
2. **Intent binding** — accept a declared intent from any of: OAuth `mission` claim
   (draft-mcguinness-oauth-mission), RFC 9396 authorization details, an `X-Agent-Intent`
   header, or MCP session metadata. If absent, infer a coarse intent from the first calls.
3. **Destructive-action step-up** — the PocketOS fix: endpoints tagged destructive
   (DELETE, bulk export, permission changes, billing) require a human approval when
   called by an agent, regardless of token scope. Works even with *zero* intent model —
   this is the MVP's "aha".
4. **Intent conformance** — per-call (or per-window) decision: does this call plausibly
   serve the declared intent? Uses a constrained-output model (doc [05](./05-intent-model-economics.md)).
5. **Agent insights for the SaaS product team** — anonymized, aggregate view of which
   intents agents pursue, where they fail, which endpoints they misuse. This is the Intent
   Gateway's insight engine pointed the other way, and it gives the product a second buyer
   (product/growth) beyond security.

## 5. The hard problem: where does the intent come from?

The SaaS server never sees the user's prompt. Three sources, in order of strength:

| Source | Strength | Availability today | Plan |
|---|---|---|---|
| **Signed/declared intent** (mission claim, RAR, intent header) | Strong, auditable | Low — drafts only | Support from day 1; publish an open spec + client helpers so agent builders can send it. Give declared-intent callers higher limits ("carrot"). |
| **Session-inferred intent** (first N calls + endpoint semantics → coarse task label) | Medium | Always | Default mode. Drift = calls that don't fit the label. |
| **Policy-only** (no intent; destructive endpoints step-up for agents) | Weak but deterministic | Always | MVP baseline; zero ML risk. |

This layering means the product **works on day 1 without anyone adopting a new standard**,
and gets better as standards land. It also turns "intent" from a marketing word into a
measurable control: % of agent sessions with declared intent, drift rate, step-ups
approved/denied.

## 6. Positioning

> **For** SaaS and dev-tool teams whose APIs and MCP servers are now called by AI agents,
> **who** can't tell what those agents are trying to do or stop them doing what users
> never approved, **Intent Guard** is an agent-aware middleware that binds every agent
> session to its intent, checks each call against it, and requires a human for
> destructive actions. **Unlike** edge bot defense (sees agents, not their tasks) and
> agent identity (issues tokens, then steps away), **it enforces intent at the API itself**
> — and shows your product team what agents are really doing with your product.

**Words to use:** agent-aware, intent-bound, step-up, conformance, agent insights.
**Avoid:** "AI firewall" (crowded), "bot management" (Cloudflare's word), and the original
repo's banned workforce vocabulary.

## 7. Defensibility (honest)

| Moat candidate | Strength | Comment |
|---|---|---|
| Intent model accuracy | Low early, medium later | Generic decision models are cheap and good; advantage only after proprietary labeled traffic (doc 05). |
| Cross-tenant agent behavior patterns | **Medium–high** | "Claude Code sessions that call X then Y usually precede Z" — network effect across SaaS customers that no single SaaS has. |
| Developer adoption / SDK embed | Medium | Switching cost once wired into handlers and approval flows. |
| Open intent-declaration spec | Medium | If your header/claim format gets adopted by agent builders, you're the reference implementation. |
| Integrations (Okta CAA, Cloudflare, Web Bot Auth) | Low | Everyone will integrate. |

## 8. Connection to the original Intent Gateway

Same core asset — an **intent extraction and aggregation engine** with privacy-first
aggregation — two directions:
- **Outbound** (original): enterprise workforce → LLMs. Heavy legal, enterprise sales.
- **Inbound** (this): agents → SaaS. Lighter legal, bottom-up sales, faster feedback.

Suggested sequencing: build the engine inside the inbound product first (faster PMF
signal), keep the k-anonymity/aggregation code shareable, and revisit the outbound
enterprise product once the engine is proven on real traffic.
