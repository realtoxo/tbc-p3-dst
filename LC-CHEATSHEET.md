# Phase 3 DST LC Cheat Sheet

Use this as the quick officer-night reference. The full reasoning lives in [Final LC Report](FINAL-REPORT.md).

## Fixed Assumptions

- Alliance raid, full physical support, no Draenei hit aura.
- Phase 3 Wowhead BiS baselines.
- No Warglaives.
- Fury baseline: dual `Vengeful Gladiator's Slicer`.
- Rogue baseline: `Vengeful Gladiator's Slicer` + `Blade of Savagery`.
- Warrior trinket candidates exclude Badge of the Swarmguard and Slayer's Crest.
- 3,000-iteration WoWSims trinket-pair matrix.
- Video corroboration flags Hunter/Madness and no-Glaives Fury weapon baselines as follow-up checks.

## Fast Ruling

Default DST lane:

1. Combat Rogue
2. Fury Warrior
3. Survival Hunter
4. Core/high-execution Ret Paladin
5. Enhancement Shaman
6. Arms Warrior
7. BM Hunter
8. Feral Cat

Ret sims highest, but do not blindly award it to Ret. Treat it as a top claim only for a stable, core, high-execution player.

## Officer Table

| Rank | Spec | With DST pair | No DST pair | DST value |
|---:|---|---|---|---:|
| 1 | Ret Paladin | DST + Brooch | Brooch + Abacus | +49.83 DPS |
| 2 | Combat Rogue | DST + Warp-Spring | Warp-Spring + Madness | +37.53 DPS |
| 3 | Fury Warrior | DST + Madness | Madness + Tsunami | +37.15 DPS |
| 4 | Survival Hunter | DST + Brooch | Brooch + Abacus | +31.90 DPS |
| 5 | Enhancement Shaman | DST + Madness | Madness + Ashtongue Vision | +28.59 DPS |
| 6 | Arms Warrior | DST + Brooch | Abacus + Direbrew | +27.93 DPS |
| 7 | BM Hunter | DST + Brooch | Brooch + Tsunami | +24.85 DPS |
| 8 | Feral Cat | DST + Tsunami | Tsunami + Brooch | +6.72 DPS |

## Spec Calls

- Combat Rogue: top DST target even without Warglaives.
- Fury Warrior: top DST target even without Warglaives; pairs DST with Madness.
- Survival Hunter: strongest Hunter DST claim.
- Ret Paladin: highest measured gap, but award only to core/high-execution Ret.
- Enhancement Shaman: real middle-priority DST claim.
- Arms Warrior: valid, but below Fury and the stronger claims.
- BM Hunter: valid, but behind Survival.
- Feral Cat: do not early-DST; the gap is tiny.

## Practical Script

1. Award early DST to core Rogue/Fury/Survival candidates.
2. Award early DST to Ret only when player quality justifies it.
3. Move Enhancement into the next tier.
4. Let Arms and BM qualify after higher-value claims.
5. Avoid early Feral DST.
6. Rerun at 10,000 iterations before fighting over sub-3-DPS trinket-pair differences.
7. Do not treat Hunter second-trinket ordering as settled until the Madness mismatch is rerun.

Source data:

- [Final P3 DST spread CSV](sources/final-p3-dst-spreads.csv)
- [Full P3 trinket-pair matrix CSV](sources/wowsims-p3-trinket-pairs.csv)
