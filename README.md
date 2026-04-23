# TBC Phase 3 Dragonspine Trophy Research

This repository is a loot-council research pack for **Dragonspine Trophy** in TBC Phase 3, using Black Temple and Hyjal Summit gear assumptions.

The current baseline is:

- Alliance raid.
- Full physical raid/group buffs.
- No Draenei `Heroic Presence` hit aura.
- Phase 3 BiS gearsets sourced from Wowhead class/spec guides.
- No Warglaives for Fury Warrior or Combat Rogue.
- Warrior candidate lists continue the prior no-Badge/no-Slayer assumption.
- Separate treatment for BM Hunter vs Survival Hunter, Fury Warrior vs Arms Warrior, and every normal physical DPS DST claimant.

## Documents

- [LC Cheat Sheet](LC-CHEATSHEET.md): Quick officer-night DST decisions and DPS deltas.
- [Final LC Report](FINAL-REPORT.md): Human-readable Phase 3 DST award plan and spread interpretation.
- [DST Sim Results](docs/dst-sim-results.md): Parsed pair-matrix results and spec notes.
- [DST Methodology](docs/dst-methodology.md): Assumptions, source files, sim harness details, and caveats.
- [Video Corroboration](docs/video-corroboration.md): Transcript-backed class-specialist notes and conflicts with the sim matrix.
- [Sim vs Video Corroboration](docs/sim-video-corroboration.md): Where the sim numbers support or conflict with the video layer.
- [Sim Assumption Log](notes/sim-assumption-log.md): Decisions, overrides, and caveats captured while building the dataset.

## Bottom Line

The Phase 3 no-Warglaives numbers support **Combat Rogue, Fury Warrior, Survival Hunter, Enhancement Shaman, and core/high-execution Retribution Paladin** as the main DST priority pool.

Do not merge Warrior specs or Hunter specs together:

- Fury is materially ahead of Arms.
- Survival is materially ahead of BM.
- Rogue remains a premium DST claimant even with no Warglaives.
- Feral Cat remains the cleanest low-priority DST claimant; its measured DST gap is only **+6.72 DPS**.

## Key Data Sources

Primary source artifacts in this repo:

- [Final P3 DST spread CSV](sources/final-p3-dst-spreads.csv)
- [Full P3 trinket-pair matrix CSV](sources/wowsims-p3-trinket-pairs.csv)
- [Wowhead gearset JSON](sources/wowhead-bis-gearsets.json)
- [Wowhead extracted item CSV](sources/wowhead-bis-items.csv)
- [WoWSims pair harness](sources/wowsims-p3-dst-pairs.go)
- [Wowhead extraction script](scripts/extract-wowhead-bis.js)
- [YouTube video source manifest](sources/youtube-video-sources.csv)
- [Cleaned transcript directory](sources/youtube-transcripts-clean/)

External sources used:

- [WoWSims TBC](https://wowsims.github.io/tbc/)
- [WoWSims TBC GitHub](https://github.com/wowsims/tbc)
- [Fury Warrior Phase 3 BiS - Wowhead](https://www.wowhead.com/tbc/guide/fury-warrior-dps-bt-hyjal-phase-3-best-in-slot-gear-burning-crusade)
- [Arms Warrior Phase 3 BiS - Wowhead](https://www.wowhead.com/tbc/guide/arms-warrior-dps-bt-hyjal-phase-3-best-in-slot-gear-burning-crusade)
- [Rogue Phase 3 BiS - Wowhead](https://www.wowhead.com/tbc/guide/rogue-dps-bt-hyjal-phase-3-best-in-slot-gear-burning-crusade)
- [BM Hunter Phase 3 BiS - Wowhead](https://www.wowhead.com/tbc/guide/beast-mastery-hunter-dps-bt-hyjal-phase-3-best-in-slot-gear-burning-crusade)
- [Survival Hunter Phase 3 BiS - Wowhead](https://www.wowhead.com/tbc/guide/survival-hunter-dps-bt-hyjal-phase-3-best-in-slot-gear-burning-crusade)
- [Feral Druid DPS Phase 3 BiS - Wowhead](https://www.wowhead.com/tbc/guide/feral-druid-dps-bt-hyjal-phase-3-best-in-slot-gear-burning-crusade)
- [Enhancement Shaman Phase 3 BiS - Wowhead](https://www.wowhead.com/tbc/guide/enhancement-shaman-dps-bt-hyjal-phase-3-best-in-slot-gear-burning-crusade)
- [Retribution Paladin Phase 3 BiS - Wowhead](https://www.wowhead.com/tbc/guide/retribution-paladin-dps-bt-hyjal-phase-3-best-in-slot-gear-burning-crusade)
- [SNO x ROGUECRAFT: Rogue P3 BIS Gear & Prio Guide](https://www.youtube.com/watch?v=p-yTLhju3H4)
- [Ragebtw: Fury Warrior Phase 3 BIS LIST](https://www.youtube.com/watch?v=q_VdxqKmm7k)
- [Marrow: Phase 3 Fury Warrior BiS and Prio Guide](https://www.youtube.com/watch?v=STxvv_E7r-Y)
- [tomatosaucin: Arms Warrior BiS Phase 3](https://www.youtube.com/watch?v=yXE0bWWhz9o)
- [Veramos: Phase 3 Beast Mastery Hunter BiS and Prio Guide](https://www.youtube.com/watch?v=hsIsJVIyOCE)
- [Veramos: Phase 3 Survival Hunter BiS and Prio Guide](https://www.youtube.com/watch?v=yqQagrZrgos)
- [Darkest_TV: Enhancement Shaman Tier 6 / Phase 3 BiS List](https://www.youtube.com/watch?v=2d-1x_z7NdE)
- [Jambrosay: Enhancement Shaman Physical DPS Loot Guide](https://www.youtube.com/watch?v=hfmde-xAD4w)
- [TheRealBlayst: Retribution Paladin Phase 3 & 4 BIS](https://www.youtube.com/watch?v=N3aQ2v1KAxI)
- [Griftin: Phase 3 Loot Guide - Feral Druid](https://www.youtube.com/watch?v=Flgbx9eII7U)

## Current Research Status

The numeric baseline is complete for:

- Arms Warrior
- Fury Warrior
- Combat Rogue
- BM Hunter
- Survival Hunter
- Feral Cat
- Enhancement Shaman
- Retribution Paladin

This pass is a 3,000-iteration trinket-pair matrix. Rerun at 10,000 iterations before treating sub-3-DPS differences as meaningful.

The video corroboration layer is complete enough for all eight covered specs. It flags two follow-ups before final guild policy: Hunter/Madness ordering conflicts with the sim, and no-Warglaives Fury should be rerun with additional weapon baselines.
