# Target Segments & Buyer Analysis

## Who actually signs the check

The product touches four budget owners. Ranked by likelihood of being the *economic buyer*:

| Rank | Buyer | Budget line | What they buy from us | Deal risk |
|---|---|---|---|---|
| 1 | **CISO / security** | Shadow-AI visibility, DLP | Governed gateway + policy enforcement + audit | Sees us vs. Netskope/Zscaler/Prompt-class tools |
| 2 | **CIO / platform eng** | LLM infra, cost control | Multi-provider gateway + spend/usage observability | Sees us vs. LiteLLM (free) — must win on governance |
| 3 | **COO / CFO (transformation)** | AI ROI measurement | Org-level intent insights, adoption dashboards | Newest budget; evangelical sale |
| 4 | **CHRO / people analytics** | Workforce insight | Aggregate work-intent analytics | Most legally sensitive; never lead with this |

**Playbook implication:** land with CISO/CIO on governance+gateway (defensible, urgent,
budgeted), expand to COO/CHRO insights once trust and data are in place. Leading with
workforce insight inverts the risk profile and stalls in legal review.

## Beachhead: tech-forward mid-market (1,000–10,000 employees), US/UK/ANZ

- Already have 3+ AI providers in use, no consolidated gateway, growing shadow-AI anxiety.
- Procurement is fast (<90 days), works councils absent, notice-based monitoring lawful.
- Need: single governed AI entry point + "what is our AI actually being used for" reporting.
- Entry ACV $40–80k; expansion via seats and insight tiers.

## Second: regulated industries (financial services, healthcare, pharma, legal)

- Highest willingness-to-pay for policy enforcement, audit trails, retention control
  (FINRA/SEC recordkeeping, HIPAA, privilege).
- Buy the *governance* story hardest; insights are secondary.
- Longer cycles (6–12 months), security reviews, likely demand VPC/on-prem deploy —
  schedule this in the roadmap before entering the segment seriously.

## Third: EU enterprises (DACH/FR/NL/Nordics)

- Strongest regulatory pull (EU AI Act deployer duties by ~Dec 2027) and the strongest
  regulatory friction (works councils, GDPR DPIAs).
- Enter only with: works-council briefing pack, aggregate-only default mode, EU data
  residency, and a DPIA template as a product artifact
  (see [`../03-legal/04-works-councils-and-global.md`](../03-legal/04-works-councils-and-global.md)).
- Counter-intuitive advantage: a privacy-proportionate-by-design product can *win* here
  precisely because US bossware-style competitors cannot pass co-determination review.

## Anti-segments (do not pursue early)

- **<500-employee companies** — no governance budget; LiteLLM/direct API is fine for them.
- **Call-center/BPO productivity monitoring** — pulls the product into per-agent scoring,
  i.e., the bossware quadrant we must avoid (legal + brand contamination).
- **Public-sector EU** — AI Act obligations for public authorities are even tighter, and
  procurement timelines exceed early-stage runway.

## User personas (not buyers)

- **AI platform admin** — configures providers, policies, tenants; needs clean multi-tenant ops.
- **Compliance officer / DPO** — needs audit logs, retention proof, DSAR export, DPIA inputs.
- **Team lead / executive** — consumes aggregate insight dashboards (k-anonymized).
- **Employee** — the subject; needs a transparency page showing exactly what is captured and
  what their employer can/cannot see. Treat the employee as a persona with rights, not just a
  data source — this is both an ethical stance and the thing that unblocks works councils.
