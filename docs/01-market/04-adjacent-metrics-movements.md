# Adjacent Metrics Movements — DORA, SPACE, Engineering Intelligence & Team Health

*Added in response to the question: do DORA metrics, velocity, team morale etc. fall under
AI governance, or can they ride the same train?*

## Short answer

They are a **separate, established category** — not AI governance — but they ride the
**same macro train** we do, and they are our most natural *integration surface and
expansion market*, not competition. One sub-area (inferred morale) is a hard red line.

## The category: Software Engineering Intelligence (SEI)

Gartner tracks this as "Software Engineering Intelligence Platforms": Jellyfish, LinearB,
DX (getdx.com), Swarmia, Faros AI, Haystack, Sleuth, plus Jira/Atlassian's own analytics
(Compass, Atlassian Analytics). What they measure:

| Framework | Metrics | Source systems |
|---|---|---|
| **DORA** | Deployment frequency, lead time for changes, change failure rate, MTTR | Git, CI/CD, incident tools |
| **SPACE** | Satisfaction & well-being, Performance, Activity, Communication, Efficiency | Above + surveys + calendar |
| **Flow/velocity** | Cycle time, WIP, throughput, sprint predictability | Jira/Linear/Azure Boards |
| **DevEx / morale** | Developer-experience scores, friction surveys, eNPS | Voluntary surveys (DX, Culture Amp) |

Buyer: **VP Engineering / CTO / engineering ops** — a different budget line from CISO
(governance) and CIO (gateway), and one that already survived its own "is this
surveillance?" debate by settling on team-level, survey-augmented metrics.

## Why they're on our train

The train is: **evidence-based leadership about how work actually happens, intensified by
"prove the AI ROI" pressure.** Every SEI vendor pivoted 2024–26 to "measure your AI coding
assistant impact" dashboards, because boards ask CTOs the exact question they ask COOs.
Structurally though, SEI platforms have a blind spot:

- **They see outcomes** (PRs merged, deploys, cycle time) and can correlate "Copilot seat
  active" with delivery stats.
- **They cannot see intent** — what work the AI was actually asked to do, on which project,
  at what depth. Correlation without mechanism.

**We hold the missing variable.** IntentRecords (see
[`../05-strategy/02-product-strategy.md`](../05-strategy/02-product-strategy.md)) joined
with DORA/flow outcomes answer the question SEI tools currently hand-wave: *which kinds of
AI-assisted work actually moved delivery metrics, and where is AI effort going that produces
nothing?*

## Strategic options (ranked)

1. **Integrate (do this):** export intent aggregates into SEI platforms (Jellyfish/DX/
   LinearB APIs, BI exports). We become the AI-signal source of record for a category that
   urgently needs it — distribution into engineering budgets without building dashboards
   they already own. Cheap: it's an export API on data we already have.
2. **Correlate (v2 module):** a lightweight "AI impact" view joining our intent aggregates
   with a tenant's own DORA data (via their existing Git/CI webhooks or SEI export).
   Sells to the CTO as a fourth budget door. Only after the core three tiers prove out.
3. **Compete (don't):** building full SEI (Jira/Git/CI connectors, flow analytics) is a
   different company with entrenched incumbents. Never lead here.

## The morale red line

Survey-based morale/DevEx (DX, Culture Amp model) is legitimate: voluntary self-report.
**Inferring morale, sentiment, or well-being from prompt content is prohibited territory** —
EU AI Act Art. 5 bans emotion inference in the workplace, and it is the single fastest way
to convert our privacy-proportionate positioning into a bossware story. If tenants ask for
"team sentiment from AI usage": the answer is no, permanently, and the refusal is written
into product principles ([non-goals](../05-strategy/02-product-strategy.md#explicit-non-goals-v1v2))
and the works-council pack — where it becomes a selling point.

## Verdict for positioning

- Taxonomy: DORA/SEI ≠ AI governance. Keep the categories distinct in all materials —
  conflating them muddies both stories in procurement.
- Narrative: same train, different cars. Pitch line: *"Your SEI tells you what shipped.
  Your intent gateway tells you what your people were actually trying to do — and whether
  AI helped. Together they close the AI-ROI loop."*
- Sequencing: integration API in the Insight tier roadmap (post-MVP), CTO-facing
  correlation module only after design partners validate demand.
