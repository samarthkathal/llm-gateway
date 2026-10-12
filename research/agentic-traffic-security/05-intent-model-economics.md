# 05 — Intent-Model Economics: is per-request intent inference affordable, and should you fine-tune?

**Your hypothesis (13 Oct 2026):** OpenAI's new decision model, Typesafe's Jev, and open
alternatives make an intent-capturing engine financially feasible; fine-tuning our own
model would be more accurate and more valuable.

**Verdict:** **Feasible — yes, clearly.** **Fine-tune — yes, but not first.** Start on a
hosted decision model, design the product so it *generates labels*, and distill into your
own small model once you have real traffic. The payoff of your own model is mostly
**latency, cost and self-hosting**, not raw accuracy on day one.

---

## 1. The new class: constrained-output "decision" models

| Model | What it is | Released | Key facts |
|---|---|---|---|
| **Typesafe Jev** (v1.13, 17 Sep 2026 on OpenRouter) | Non-generative transformer: returns only options you define (labels, scores, yes/no); several questions per input in one pass; no per-task retraining | ~15 Sep 2026 | Fastest-adopted model in Vercel AI Gateway's history within 24 h; reported pricing ~$42 per billion input tokens with free output (secondary source) |
| **OpenAI Decisions API** (GPT-6 Luna) | Give a question + fixed answer list + text/image context → one answer from the list; for routing, classification, choosing an agent's next step from an allowlist | DevDay 2026 (~29 Sep 2026), limited preview → broader | Pricing and latency not published at launch. Note: one benchmark calls the model "GPT-5.6 Luna" — naming varies across sources |
| **gpt-oss-safeguard** (120B / 20B, open weights) | Reasoning classifier that takes *your policy as a prompt* at inference | 29 Oct 2025 | Self-hostable; good for "does this action violate policy P?" |
| Llama Guard, small fine-tuned encoders (ModernBERT-class) | Open classifiers | various | Cheapest/fastest once trained on your labels |

These map almost exactly onto the conformance question in [04 §4](./04-idea-deep-dive.md):
*"Given declared intent I and call C, answer one of {in-intent, drift, unsafe}."*
No free text is needed — which is why this class is cheap.

## 2. Benchmark reality check

Independent comparison (LessWrong, 23 Sep 2026; 2,200 reasoning traces, harm-level labels):

| Model | Mean latency | Cost / 1,000 decisions | Exact-match | Macro-F1 |
|---|---|---|---|---|
| **Jev** | **542 ms** | **$0.056** | 71.5% | 58.2% |
| GPT Luna (OpenAI) | 2,007 ms | $0.23 | **79.4%** | **70.3%** |
| Claude Sonnet 5 | 3,348 ms | $3.17 | 71.4% | 56.4% |

Caveats from the post and comments: Jev was **weaker at detecting the harmful cases** (the
ones that matter most for security); the author's headline multipliers don't match the
listed prices (~57× and ~4× cheaper, not 566×/41×); and a commenter argued a small
fine-tuned transformer could match it. One benchmark, one task — treat as directional.

## 3. Unit economics for the wedge

Assume a mid-market SaaS customer with **1M agent calls/month** and a price of
**$500/month** (doc 03 ACV midpoint).

| Strategy | Model cost / month | % of revenue | Viable? |
|---|---|---|---|
| Frontier LLM on every call (Sonnet-class) | ~$3,170 | 634% | ❌ |
| Luna on every call | ~$230 | 46% | ⚠︎ thin margin |
| Jev on every call | ~$56 | 11% | ✅ |
| Rules on every call + Jev/Luna only on **sensitive endpoints or ~10% of calls** | ~$6–25 | 1–5% | ✅✅ |
| Own distilled small model, self-hosted | ⚠︎ ~$5–20 (CPU/GPU share) | ~1–4% | ✅✅ + no third-party data egress |

So your hypothesis holds: **the economics flipped in Sep–Oct 2026.** A year earlier the
only option was a general LLM at ~$3/1k, which kills a per-request security product.

## 4. Latency is the real constraint, not cost

A 0.5–2 s hosted call **cannot sit inline on every API request** of a SaaS product.
Design around it:

1. **Inline, deterministic, <5 ms:** agent classification, destructive-endpoint step-up,
   allowlists, rate limits. No model.
2. **Inline, model, only where it matters:** destructive/sensitive endpoints (where a
   human step-up is already acceptable latency) and the *first* calls of a session.
3. **Async, per window:** conformance on sampled/batched calls; on drift, tighten the
   session (rate-limit, require step-up next call, revoke).
4. **Later, local model:** a distilled ⚠︎ ~100–300M-parameter encoder in the SDK/sidecar
   could plausibly run in tens of milliseconds on CPU — this is what makes fine-tuning
   worthwhile.

## 5. Should you fine-tune your own model?

| Argument | For fine-tuning now | Against fine-tuning now |
|---|---|---|
| Accuracy | Domain-specific (API calls + intents) beats general | **No labeled data yet**; label set will change weekly during discovery |
| Cost | Cheapest at scale | Hosted decision models already ~$0.06/1k |
| Latency | Only way to go inline everywhere | Not needed for MVP (§4 layers 1–3) |
| Data control | Customer payloads never leave their infra — big enterprise selling point | Matters after PMF, not before |
| Moat | Model + labels become proprietary | A generic fine-tune with no unique data is not a moat |
| Effort | — | Weeks of ML work a solo founder should spend on customers |

**Recommended path**

1. **Weeks 0–8:** Policy-only + Jev (or Luna) behind a `Decider` interface; log every
   decision with redacted features.
2. **Make the product generate labels.** Every **step-up approval or denial by a human is a
   ground-truth label** ("this agent action was / was not what I intended"). Design
   partners' admins correcting false positives are labels too. This is the most important
   design decision in the whole product.
3. **At ~5–20k labeled decisions** (⚠︎ estimate) across ≥3 customers: distill the hosted
   model's outputs + human labels into a small open model; A/B against the hosted model.
4. **Ship self-hosted model** as an enterprise/on-prem tier feature.

Your instinct is right about where the value ends up — a model trained on cross-tenant
agent-behavior labels is the defensible asset (see moat table in doc 04) — but the order
is **data first, model second**.

## 6. Risk to note

Model vendors are moving into exactly this slot: OpenAI pitches Decisions for choosing an
agent's next step from an allowlist, and Jev markets "agent safety checks" (allow / flag /
block a risky command). Your product must not be "a wrapper around Jev." The value is in
**where you sit** (inside the SaaS API), **what you bind to** (declared intent + identity),
**what you do** (step-up, policy), and **what you learn** (cross-tenant labels).
