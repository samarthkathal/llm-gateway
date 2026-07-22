# Market Sizing — TAM / SAM / SOM

> Method note: public market reports disagree by 2–3x depending on scope definitions. We cite
> ranges, state which report each figure comes from, and size bottom-up for SAM/SOM rather
> than trusting top-down numbers.

## Adjacent-market reference points (2025–2026)

| Market | Size (latest) | Growth | Source |
|---|---|---|---|
| Enterprise LLM (software) | $5.9–7.6B (2026) | ~28–30% CAGR → $48–91B by 2034–36 | [Fortune BI](https://www.fortunebusinessinsights.com/enterprise-llm-market-114178), [FMI](https://www.futuremarketinsights.com/reports/enterprise-llm-market) |
| LLM middleware/gateway layer | subset of above; ~42% of enterprises run one | ~49.6% CAGR through 2034 | [index.dev roundup](https://www.index.dev/blog/llm-enterprise-adoption-statistics) |
| Workforce analytics | $2.4–4.5B (2025–26) | ~12–13% CAGR → ~$7.1B by 2034 | [Fortune BI](https://www.fortunebusinessinsights.com/industry-reports/workforce-analytics-market-100299), [SkyQuest](https://www.skyquestt.com/report/workforce-analytics-market) |
| Employee monitoring software | $3.9–4.9B (2025) | ~18% CAGR → $7.6B by 2029 / $22.8B by 2034 | [TBRC](https://www.thebusinessresearchcompany.com/report/employee-monitoring-software-global-market-report), [Research & Markets via NLR](https://natlawreview.com/press-releases/employee-monitoring-market-track-hit-76-billion-2029) |
| AI governance platforms | $0.35–0.75B (2026) | ~35–45% CAGR; Gartner: >$1B by 2030 | [Grand View](https://www.grandviewresearch.com/industry-analysis/ai-governance-market-report), [Precedence](https://www.precedenceresearch.com/ai-governance-market), [Gartner via Nemko](https://digital.nemko.com/news/ai-governance-platforms-market-to-surpass-1-billion-by-2030) |

Takeaway: the two big, established budgets (workforce analytics, monitoring) grow steadily;
the two new budgets (LLM infra, AI governance) grow explosively from smaller bases. Our
product draws from all four, anchored in the new ones.

## TAM (top-down, blended)

Target buyer: organizations with ≥1,000 employees and material AI adoption.

- ~55–60k firms globally with 1,000+ employees; assume 70% will have governed AI programs
  by 2029 → ~40k addressable organizations.
- Blended ACV assumption at maturity: $60k–150k/yr (gateway + governance + insights tiers,
  scaling with seats).
- **TAM ≈ $2.4B–6B/yr** by decade end — consistent with the sum of the AI-governance and
  AI-usage-analytics slices of the adjacent markets above.

## SAM (what we can serve with this product in 3 years)

Constraints applied: English-first product; cloud-deployable; jurisdictions where employee
AI-usage analytics is lawful with notice + safeguards (US, UK, ANZ, Nordics, DACH-with-works-
council-playbook, India, SEA); mid-market-to-large enterprises (1k–20k employees) as the
realistic entry segment.

- ~12–15k organizations fit; assume 25% run inline AI gateways with governance budgets in the
  window → 3–3.75k orgs.
- Entry ACV $40k–80k.
- **SAM ≈ $150M–300M/yr.**

## SOM (credible 3-year capture)

Design-partner motion → 10–15 paying tenants in year 1–2; 40–80 tenants by year 3 at
$50–80k ACV.

- **SOM ≈ $3M–6M ARR by end of year 3.** That is a fundable seed→A trajectory in this
  category (see [`../07-financial/01-funding-landscape.md`](../07-financial/01-funding-landscape.md)),
  not a rocket ship — the bet is category timing (EU AI Act clock, AI-ROI pressure) expands
  the budget line underneath us.

## Sensitivity — what changes the numbers most

1. **Does "AI usage intelligence" become its own budget line?** If CFO-side AI-ROI reporting
   becomes standard (analyst pressure suggests yes), SAM roughly doubles.
2. **EU AI Act enforcement posture.** Strict enforcement from Dec 2027 accelerates EU SAM;
   further delay (already slipped once from Aug 2026) pushes EU revenue out a year.
3. **Bundling by platforms.** If Microsoft ships org-hierarchy-joined Copilot analytics
   with governance, the Microsoft-only segment of SAM (~30–40%) becomes hard to win —
   multi-provider enterprises remain open. See risk register.

## Segment prioritization (detail in [`03-target-segments.md`](./03-target-segments.md))

1. Tech-forward mid-market enterprises (1k–10k employees), US/UK — beachhead.
2. Regulated industries (finserv, healthcare, pharma) — highest willingness-to-pay for the
   governance layer; heavier sales.
3. EU enterprises with works councils — largest compliance pull, longest cycles; enter with
   a works-council-ready privacy posture as the differentiator, not despite it.
