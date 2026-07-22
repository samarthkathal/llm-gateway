# Competitor Profiles (Deep Dives)

Ten profiles, ordered by how often we expect to meet them in a deal.

---

## 1. LiteLLM (BerriAI)
- **What:** Open-source LLM gateway; unified OpenAI-compatible API over 100+ providers.
- **Position:** Default choice for self-hosting engineering teams; free; huge community.
- **Strengths:** Zero cost, unlimited volume, provider breadth, dev mindshare.
- **Weaknesses:** Ops burden on customer, thin enterprise SLA story, no identity/org
  semantics, no compliance layer ([TrueFoundry alternatives review](https://www.truefoundry.com/blog/litellm-alternatives)).
- **vs. us:** We will lose any "just route my traffic" deal to free. We must sell what
  LiteLLM structurally lacks: tenancy, HRIS join, policy engine, insight layer. Consider
  *supporting* LiteLLM as an ingestion point rather than fighting it.

## 2. Portkey
- **What:** Managed AI gateway: routing, fallbacks, observability, guardrails, prompt mgmt.
- **Position:** Startup-to-midmarket managed convenience; fast integration.
- **Strengths:** Product polish, observability dashboards, quick start.
- **Weaknesses:** Per-request pricing contested at scale; governance is engineering-grade,
  not compliance-grade ([pricing analysis](https://www.truefoundry.com/blog/portkey-pricing-guide)).
- **vs. us:** Closest in "gateway + visibility" story. Differentiate on org-awareness and
  regulated-enterprise governance.

## 3. Kong AI Gateway
- **What:** AI plugins on the Kong API-gateway platform.
- **Strengths:** Enterprise install base, performance (their benchmark: 228% faster than
  Portkey, 859% vs LiteLLM — vendor-run, treat accordingly;
  [Kong](https://konghq.com/blog/engineering/ai-gateway-benchmark-kong-ai-gateway-portkey-litellm)),
  mature RBAC/SSO.
- **Weaknesses:** Requires Kong stack; AI semantics are plugin-deep only.
- **vs. us:** Threat in Kong-standardized accounts. Elsewhere, neutral.

## 4. TrueFoundry
- **What:** K8s-native AI gateway/control plane; VPC or air-gapped; RBAC, budgets, SOC2/HIPAA
  posture ([guide](https://www.truefoundry.com/blog/a-definitive-guide-to-ai-gateways-in-2026-competitive-landscape-comparison)).
- **Position:** The "enterprise governance gateway" — nearest neighbor in the infra quadrant.
- **vs. us:** They govern *access and spend*; we govern *meaning and policy over people data*.
  If they add HRIS-joined analytics, they become the direct competitor — watch closely.

## 5. Harmonic Security
- **What:** AI-usage DLP: monitors GenAI app usage, protects sensitive data in prompts.
- **Funding:** $26M total; $17.5M Series A (Next47, Oct 2024)
  ([Crunchbase](https://www.crunchbase.com/organization/harmonic-security), [PitchBook](https://pitchbook.com/profiles/company/538273-72)).
- **Position:** The last major independent in the capture-for-security quadrant.
- **vs. us:** Overlapping capture mechanics, disjoint output (risk events vs. intent graph).
  Also a signal: if Harmonic exits into a platform next, the whole quadrant is closed and
  our "insight layer" positioning becomes uncontested — or platform-bundled against us.

## 6. SentinelOne + Prompt Security
- **What:** GenAI firewall (prompt inspection, redaction, shadow-AI catalog) now inside
  Singularity ([AppSec Santa](https://appsecsanta.com/prompt-security)).
- **vs. us:** In security-led deals we will be asked "why not the platform module?" Answer:
  the module blocks threats; it does not tell the COO where the workforce's AI effort goes,
  and it has no tenant-facing insight product. Avoid feature-war on blocking.

## 7. Check Point + Lakera
- **What:** Prompt-injection/guardrail tech absorbed into Check Point's fabric
  ([Tracxn](https://tracxn.com/d/companies/lakera/__bJHYPbXKymeh9hjUrJceMqDjX9dz9WwohC0zfAtjCyw)).
- **vs. us:** Same shape as #6 from the network side.

## 8. Microsoft (Viva Insights + Copilot analytics + Purview)
- **What:** Bundled workplace analytics (calendar/collab signals), expanding Copilot usage
  reporting; Purview for DLP/compliance; Entra for identity.
- **Strengths:** Distribution, bundling price (≈$0 marginal), org data already in Graph.
- **Weaknesses:** Microsoft-ecosystem-only signal; Copilot-centric (blind to Claude, Gemini,
  internal apps); Viva's analytics are proxy-based; historically weak at cross-vendor AI.
- **vs. us:** The strategic threat. Moat = multi-provider coverage + neutral-vendor trust
  (an employer may not want Microsoft holding its intent graph) + depth of intent semantics.

## 9. ActivTrak / Insightful / Prodoscore
- **What:** Activity-monitoring heritage workforce analytics
  ([alternatives roundup](https://www.currentware.com/comparisons/activtrak-alternative/)).
- **vs. us:** Rarely in the same deal if we position correctly. Their existence is mainly a
  *category hazard*: procurement and works councils will pattern-match us to them unless the
  privacy architecture visibly differs.

## 10. Credo AI / Holistic AI / OneTrust (governance platforms)
- **What:** AI inventory, risk workflows, EU AI Act conformity tooling.
- **vs. us:** Complementary today — we are the live-evidence source their audits need.
  Partnership/integration candidates; long-term possible convergence.

---

## Summary scorecard

| Competitor | Gateway | Prompt semantics | Org/HRIS join | Insight product | Compliance layer | Buyer |
|---|---|---|---|---|---|---|
| LiteLLM | ●● | – | – | – | – | Eng |
| Portkey | ●● | ◐ (guardrails) | – | ◐ (usage) | ◐ | Eng |
| Kong AI GW | ●● | ◐ | – | – | ◐ | Eng/IT |
| TrueFoundry | ●● | ◐ | – | ◐ | ● | Eng/IT |
| Harmonic | ◐ (proxy) | ● (risk only) | – | – | ◐ | CISO |
| S1/Prompt Sec | ◐ | ● (risk only) | – | – | ◐ | CISO |
| Microsoft Viva | – | – | ●● (Graph) | ● (proxy-based) | ● (Purview) | CIO/CHRO |
| ActivTrak et al. | – | – | ◐ | ● (activity) | ◐ | COO/CHRO |
| Credo/Holistic | – | – | – | – | ●● (paperwork) | Compliance |
| **Intent Gateway (us)** | ●● | ●● (intent) | ●● | ●● (aggregate) | ●● (by design) | CIO+CISO→COO |

●● strong ● present ◐ partial – absent. The right three columns together are the whitespace.
