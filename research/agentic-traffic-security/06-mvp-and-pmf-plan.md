# 06 — MVP & Product-Market-Fit Plan (solo, bootstrapped)

**Goal:** In 8 weeks, learn whether SaaS/MCP builders will install and pay for
server-side, intent-aware agent controls — spending as little build time as possible
before that answer is known.

---

## 1. Ideal design partner (who to talk to first)

- B2B SaaS or dev-tool company, 10–500 people, **already exposing an MCP server or seeing
  agent traffic on its API** (Claude Code / Codex / Cursor users scripting against it).
- Has at least one **destructive or sensitive endpoint** (delete, bulk export, billing,
  permissions, deploy).
- A technical founder/CTO or platform lead who can install middleware without procurement.
- Bonus: infra/dev-tools categories (hosting, databases, CI/CD, observability) — the
  PocketOS/Railway pattern lives here and your platform/DevOps background gives instant
  credibility.

Where to find them: Okta CAA / MintMCP partner lists, MCP server directories (Glama,
MCP.so, PulseMCP) filtered to companies, YC/Peak XV dev-tool portfolios, Indian SaaS
(Bengaluru) founders for in-person access.

## 2. Weeks 0–2: problem validation (no code)

15 interviews, 30 minutes each. Questions:

1. How do you know when a request comes from an agent vs your UI vs a script?
2. Have agents done anything on your platform users didn't intend? Walk me through the last one.
3. Has agent risk ever blocked or delayed a launch (MCP server, public API, write scopes)?
4. What would you do today if an agent bulk-deleted a customer's data with a valid token?
5. Who would own fixing this — security, platform, or product? What have you tried?
6. Would knowing *what agents try to do on your product* change your roadmap?

**Pass gate (end of week 2):** ≥5/15 describe a concrete agent incident or blocked launch,
**and** ≥2 agree to a free pilot with real traffic. Otherwise → see kill criteria.

Cheap parallel test: a landing page + one technical blog post ("How a 9-second GraphQL
mutation wiped prod — and the 40 lines of middleware that would have stopped it") posted
to HN / r/devops / MCP Discord. Target: ≥100 waitlist sign-ups from company domains.

## 3. Weeks 2–8: build the smallest thing that proves value

Go fits your stack and the existing repo; ship **Go + Node (Express/Hono) middleware**
first since most MCP servers are TS.

| Week | Deliverable | Proves |
|---|---|---|
| 2–3 | **Agent classifier middleware**: Web Bot Auth verify (reuse Cloudflare's OSS), known agent fingerprints, MCP header detection, token-claim parsing → tags each request human / verified agent / unknown agent; log-only dashboard | Customers can *see* their agent traffic (aha #1) |
| 3–4 | **Destructive-endpoint step-up**: annotate routes; agent calls require one-click approval via Slack/email link; default deny on timeout | Prevents PocketOS-class incidents with zero ML (aha #2) |
| 4–5 | **Intent binding**: accept `X-Agent-Intent` header / OAuth `mission` claim / MCP session metadata; open-source a tiny client helper + spec | Declared intent is practical |
| 5–6 | **Conformance check** via `Decider` interface (Jev or Luna); async on sampled calls, inline on sensitive routes; drift → tighten session | Intent adds value beyond rules |
| 6–7 | **Agent insights page**: top intents, drift rate, step-ups approved/denied, endpoints agents fail on (aggregate-only) | Second buyer (product team) cares |
| 7–8 | Pricing page, self-serve signup, docs; convert pilots | Willingness to pay |

**Explicitly out of scope for MVP:** own model, on-prem, SOC 2, enterprise SSO for the
dashboard, browser-agent (web UI) protection, outbound/workforce intent.

## 4. PMF metrics

| Metric | Target by week 12 | Why |
|---|---|---|
| Design partners with middleware in production | ≥3 | Real install is the strongest signal |
| Step-ups triggered that a human **denied** | ≥1 per partner | Proof it caught something real |
| Weekly active dashboard viewers per partner | ≥2 people | Value beyond install-and-forget |
| Paid conversions from pilots | ≥2 at ≥$250/mo | Willingness to pay |
| Sean Ellis "very disappointed" | ≥40% of active users | Classic PMF bar |
| Share of agent sessions with declared intent | trend ↑ | Whether the intent standard is catching on |

## 5. Pricing hypothesis (to test, not to trust)

| Tier | Price | Includes |
|---|---|---|
| Free | $0 | Agent classification + log-only, 100k requests/mo |
| Team | $250/mo | Step-up approvals, intent binding, 2M requests/mo, 30-day insights |
| Growth | $1,000/mo | Conformance model on sensitive routes, 20M requests, 1-year insights, SLA |
| Enterprise | custom | Self-hosted model, data residency, SSO, custom policies |

Benchmark positioning: above free IsAgent-style detectors, well below enterprise API
security contracts.

## 6. Kill / pivot criteria

| Signal | Action |
|---|---|
| <5/15 interviewees have a concrete agent pain by week 2 | Stop. Problem not felt yet at this segment; revisit in 6 months. |
| Pain exists but "Cloudflare/Okta already covers it" is the common answer | Pivot to the **insights** half (agent analytics for product teams) |
| Pilots install but never trigger a denied step-up in 4 weeks | Value too theoretical → narrow to infra/dev-tool SaaS with destructive APIs, or stop |
| Cloudflare, Okta or Wallarm ships declared-intent enforcement for the mid-market | Become an integration/plugin for them, or exit to the outbound Intent Gateway |
| Strong pull from *enterprises* consuming SaaS rather than vendors | Revisit the middleware option (doc 04 §3) with the original gateway code |

## 7. Bootstrap budget (8 weeks)

| Item | ⚠︎ Est. cost |
|---|---|
| Hosting (one small VM / Fly.io / Railway) + DB | ~$50/mo |
| Decision-model API calls during pilots | <$50 total |
| Domain, email, landing page | ~$30 |
| **Total cash** | **< $300** — the real cost is your time |
