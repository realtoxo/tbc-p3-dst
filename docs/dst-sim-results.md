# Phase 3 DST Sim Results

This document summarizes the Phase 3 trinket-pair matrix.

## Final Spread

The spread compares the best pair containing DST against the best pair without DST.

| Rank | Spec | With DST pair | No DST pair | With DST DPS | No DST DPS | DST value |
|---:|---|---|---|---:|---:|---:|
| 1 | Ret Paladin | DST + Brooch | Brooch + Abacus | 2346.33 | 2296.50 | +49.83 |
| 2 | Combat Rogue | DST + Warp-Spring | Warp-Spring + Madness | 2220.64 | 2183.11 | +37.53 |
| 3 | Fury Warrior | DST + Madness | Madness + Tsunami | 1704.06 | 1666.91 | +37.15 |
| 4 | Survival Hunter | DST + Brooch | Brooch + Abacus | 2050.36 | 2018.46 | +31.90 |
| 5 | Enhancement Shaman | DST + Madness | Madness + Ashtongue Vision | 2180.47 | 2151.88 | +28.59 |
| 6 | Arms Warrior | DST + Brooch | Abacus + Direbrew | 1504.25 | 1476.32 | +27.93 |
| 7 | BM Hunter | DST + Brooch | Brooch + Tsunami | 2110.25 | 2085.41 | +24.85 |
| 8 | Feral Cat | DST + Tsunami | Tsunami + Brooch | 2615.30 | 2608.59 | +6.72 |

## Raw Outputs

- [Final spread CSV](../sources/final-p3-dst-spreads.csv)
- [Full trinket-pair matrix CSV](../sources/wowsims-p3-trinket-pairs.csv)
- [WoWSims harness](../sources/wowsims-p3-dst-pairs.go)

## Read Notes

- This is personal DPS only.
- Survival Hunter raid value from `Expose Weakness` is not credited as a raid-wide gain here.
- Ret, Arms, Enhancement, and Feral support value is not converted into raid DPS.
- The target is neutral/unknown, so `Mark of the Champion` is not modeled as a Demon/Undead encounter item.
- Results use 3,000 iterations per trinket pair. Treat tiny differences cautiously.
