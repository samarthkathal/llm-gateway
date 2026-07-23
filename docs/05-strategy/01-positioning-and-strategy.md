# Strategy Analysis — Positioning, Moat, Sequencing

## Positioning statement

> **For** enterprises scaling AI across their workforce, **who** need to govern AI usage and
> understand what it's actually doing for the organization, **the Intent Gateway** is an
> AI-traffic control plane **that** turns every sanctioned AI interaction into governed,
> org-aware insight. **Unlike** AI security tools that only block risk, gateways that only
> route traffic, or workforce-monitoring tools that surveil individuals, **we** capture
> meaning once — at the gateway — and serve security, compliance, and leadership from it,
> with employee-visible privacy guarantees built in.

### Category naming

Do not create a category name prematurely. Attach to funded, searched categories in this
order: **"AI usage governance"** (CISO budget, exists now) → **"AI gateway"** (CIO budget,
exists now) → introduce **"intent intelligence"** as our differentiated layer inside those,
letting it graduate to a category if the market picks it up.

### Words we never use

"Monitoring", "surveillance", "productivity tracking", "employee scoring". Not cosmetic —
each maps to a legal regime and a procurement red flag (see [`../03-legal/`](../03-legal/)).
The product language is *governance, visibility, insight, potential*.

## Strategic frame (Porter, abbreviated)

- **Supplier power:** LLM providers — mitigated by multi-provider abstraction (our gateway
  heritage is the hedge itself).
- **Buyer power:** high in beachhead (mid-market has options); mitigated by compounding
  data value — the intent graph gets more valuable with tenure, raising switching costs.
- **Substitutes:** "do nothing + Excel", Copilot dashboards, SSE add-ons. Countered by the
  double story (governance *and* insight from one deployment).
- **New entrants:** low barrier to the gateway, **high barrier to the trust position** —
  compliance artifacts, works-council pack, k-anonymity architecture, SOC 2 take time that
  can't be compressed by capital alone. Trust is the moat capital can't shortcut.
- **Rivalry:** intense in each adjacent quadrant, empty at the intersection (see
  [`../02-competition/01-competitive-landscape.md`](../02-competition/01-competitive-landscape.md)).

## The three strategic bets (make them explicit, revisit quarterly)

1. **The intersection bet:** enterprises prefer one governed capture point serving four
   budgets over four point tools. *Falsifier:* design partners buy governance but ignore
   insights (or vice versa) → collapse to the half that sells.
2. **The proportionality bet:** privacy-proportionate architecture wins deals rather than
   limiting them (esp. EU). *Falsifier:* repeated losses to bossware-grade competitors on
   feature depth → the market wants what we refuse to build; exit or re-scope.
3. **The timing bet:** EU AI Act Dec-2027 + CFO AI-ROI pressure create the budget line
   within 24 months. *Falsifier:* further AI Act delay + AI-spend retrenchment → extend
   runway, lean on security budget only.

## Moat construction sequence

| Horizon | Moat layer | Why it compounds |
|---|---|---|
| 0–6 mo | Compliance artifacts (SOC 2 start, DPIA generator, transparency UX) | Time-gated; competitors must serve the same time |
| 6–18 mo | Intent taxonomy + extraction quality tuned on real traffic | Data-network effect within each tenant |
| 18–36 mo | Cross-tenant benchmark insights ("your AI adoption vs. industry", aggregate-only) | True network effect; needs tenant count |
| 36+ mo | Works-council/regulator trust record in EU | Reputation moat; near-impossible to fast-follow |

## Build / partner / avoid

- **Build:** tenancy, identity+hierarchy graph, intent extraction, policy engine, insights.
- **Partner/adopt:** routing internals (embed or interop with LiteLLM), prompt-injection
  defense (integrate, don't compete), conformity workflows (integrate with Credo/OneTrust).
- **Avoid:** individual performance analytics, sentiment analysis, content of *non-AI*
  employee communications (email/Slack bodies) — Aware's turf and the surveillance trap.

## Exit landscape (informing, not driving, strategy)

Acquirer classes evidenced by 2024–25 deals: security platforms (SentinelOne, Check Point,
Mimecast pattern), enterprise-software platforms (Microsoft/Salesforce/ServiceNow), HCM
vendors (Workday — note: also our integration partner). Building the org-aware insight layer
maximizes optionality across all three; building only capture repeats the already-consolidated
pattern at lower multiples.
