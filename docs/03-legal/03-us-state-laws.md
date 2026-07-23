# US State Law Landscape

Federal baseline (ECPA business-purpose exception) makes employer monitoring of its own
systems broadly lawful. The action is in state law, and it is tightening yearly.

## Electronic-monitoring notice states (as of 2026)

| State | Requirement | Penalty/notes |
|---|---|---|
| **New York** | Written notice on hire + conspicuous workplace posting for phone/email/internet monitoring | $500 / $1,000 / $3,000 escalating per violation |
| **Connecticut** | Written notice before monitoring begins, specifying types of monitoring | First notice-law state; broad scope |
| **Delaware** | Written or electronic notice/acknowledgment before monitoring | Per-use acknowledgment option |
| **Colorado** | Notice requirements + AI-decision rules (Colorado AI Act) | Dual exposure: monitoring + ADS |

Sources: [AI Policy Desk state tracker](https://www.aipolicydesk.com/blog/ai-employee-monitoring-disclosure-laws-2026-state-tracker),
[Intelogos state-by-state guide](https://www.intelogos.com/blog/employee-monitoring-laws-by-state),
[WorkTime US monitoring-law FAQ](https://www.worktime.com/blog/legal-aspects/most-asked-questions-on-us-employee-monitoring-laws).

## California — the most demanding regime

- **CCPA/CPRA applies to employee data** (since 2023): notice at collection, access,
  deletion, correction rights for workers.
- **FEHA Automated-Decision Systems regulations:** disclosure and bias obligations when AI
  influences employment decisions.
- **2026 expansion:** data-minimization duties, employee access rights, and justification
  requirements for invasive monitoring methods
  ([2026 compliance guides](https://trick-tack.com/employee-monitoring-laws/),
  [Visual Timesheets overview](https://www.visualtimesheets.com/blog/employee-monitoring-laws-by-state)).

**Implication:** California compliance ≈ GDPR-lite. If the architecture passes GDPR
proportionality, California is covered; do not build a separate US-only data model.

## AI-specific employment-decision laws

- **Illinois** (HB 3773, eff. 2026): notice + anti-discrimination duties for AI in
  employment decisions.
- **Colorado AI Act:** developer/deployer duties for "consequential decisions" incl.
  employment (timeline shifted; monitor).
- **NYC Local Law 144:** bias audits for automated employment decision tools — only if a
  tool *makes/substantially assists* hiring/promotion decisions. Another reason our product
  must not feed such decisions.

## Wiretap/two-party-consent states (CA, IL, WA, etc.)

Low direct risk (employer's own systems, business purpose) but reinforces: notice +
acknowledgment workflow should be universal, not state-gated.

## Product requirements extracted

1. **Universal notice pack:** employee disclosure page + acknowledgment capture + NY-style
   poster template shipped per tenant; treat the strictest state (CA) as the default.
2. **No ADS features:** nothing in the product may make or substantially assist individual
   employment decisions → keeps us out of NYC LL144 / IL / CO ADS regimes.
3. **Data-minimization mode** satisfying CA 2026 rules = the same aggregate-first
   architecture GDPR requires. One architecture, all jurisdictions.
4. **Per-state config is a legal-content problem, not a code problem** — policy templates
   maintained as data (reviewable by counsel), not hardcoded logic.
