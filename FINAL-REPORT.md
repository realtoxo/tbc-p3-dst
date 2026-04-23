# Phase 3 Dragonspine Trophy Final LC Report

This is the officer-facing version of the Phase 3 DST research. It uses readable trinket pairs instead of "replacement" language.

## Assumptions

- Alliance raid.
- Full physical raid/group buffs.
- No Draenei hit aura.
- Wowhead Phase 3 BiS gear baselines.
- BM Hunter and Survival Hunter are separate.
- Fury Warrior and Arms Warrior are separate.
- Fury Warrior and Combat Rogue are assumed to have **no Warglaives**.
- Fury Warrior no-Warglaives baseline uses dual `Vengeful Gladiator's Slicer`.
- Combat Rogue no-Warglaives baseline uses `Vengeful Gladiator's Slicer` + `Blade of Savagery`.
- Warrior trinket candidates continue the prior no-Badge/no-Slayer assumption.
- The sim is a 3,000-iteration pair matrix, not a fixed one-for-one replacement pass.

Source CSVs:

- [Final P3 DST spreads](sources/final-p3-dst-spreads.csv)
- [Full P3 trinket-pair matrix](sources/wowsims-p3-trinket-pairs.csv)

Class-specialist corroboration:

- [Video Corroboration](docs/video-corroboration.md)

## TLDR Priority

Default DST lane under these assumptions:

1. **Combat Rogue**
2. **Fury Warrior**
3. **Survival Hunter**
4. **Retribution Paladin**, if core and high-execution
5. **Enhancement Shaman**
6. **Arms Warrior**
7. **BM Hunter**
8. **Feral Cat**

Ret has the highest measured personal DPS delta, but it remains execution-sensitive. Do not give DST to an inconsistent Ret over a core Rogue, Fury Warrior, or strong Survival Hunter.

Fury is a real high-priority Phase 3 DST candidate even without Warglaives. Rogue remains a premium claimant because its best no-DST pair is still far behind `DST + Warp-Spring`.

## Phase 3 DST Spread

This table compares the best tested pair containing DST against the best tested pair without DST.

| Rank | Spec | With DST pair | No DST pair | DST value |
|---:|---|---|---|---:|
| 1 | Retribution Paladin | DST + Brooch | Brooch + Abacus | +49.83 DPS |
| 2 | Combat Rogue | DST + Warp-Spring | Warp-Spring + Madness | +37.53 DPS |
| 3 | Fury Warrior | DST + Madness | Madness + Tsunami | +37.15 DPS |
| 4 | Survival Hunter | DST + Brooch | Brooch + Abacus | +31.90 DPS |
| 5 | Enhancement Shaman | DST + Madness | Madness + Ashtongue Vision | +28.59 DPS |
| 6 | Arms Warrior | DST + Brooch | Abacus + Direbrew | +27.93 DPS |
| 7 | BM Hunter | DST + Brooch | Brooch + Tsunami | +24.85 DPS |
| 8 | Feral Cat | DST + Tsunami | Tsunami + Brooch | +6.72 DPS |

## LC Interpretation

### Combat Rogue

P3 no-Warglaives Rogue still has a strong DST claim.

- Best DST pair: `DST + Warp-Spring`, 2220.64 DPS.
- Best no-DST pair: `Warp-Spring + Madness`, 2183.11 DPS.
- DST gap: **+37.53 DPS**.

Practical ruling: **Rogue remains a top DST target.** Madness is a strong no-DST fallback, but it does not close the DST gap.

### Fury Warrior

No-Warglaives Fury is nearly tied with Rogue in measured DST value.

- Best DST pair: `DST + Madness`, 1704.06 DPS.
- Best no-DST pair: `Madness + Tsunami`, 1666.91 DPS.
- DST gap: **+37.15 DPS**.

Practical ruling: **Fury should be in the first DST pool.** The no-Warglaives assumption does not push Fury out of priority.

### Survival Hunter

Survival remains the stronger Hunter DST claimant.

- Best DST pair: `DST + Brooch`, 2050.36 DPS.
- Best no-DST pair: `Brooch + Abacus`, 2018.46 DPS.
- DST gap: **+31.90 DPS**.

Practical ruling: **Survival is a strong DST target**, especially when the player is core and the raid values the spec's broader support.

### Retribution Paladin

Ret has the highest measured personal gap, but it should still be filtered by player quality.

- Best DST pair: `DST + Brooch`, 2346.33 DPS.
- Best no-DST pair: `Brooch + Abacus`, 2296.50 DPS.
- DST gap: **+49.83 DPS**.

Practical ruling: **Core/high-execution Ret is a valid early DST award.** Do not use this number to justify awarding DST to an inconsistent Ret ahead of more reliable premium claimants.

### Enhancement Shaman

Enhancement remains a meaningful DST user in P3.

- Best DST pair: `DST + Madness`, 2180.47 DPS.
- Best no-DST pair: `Madness + Ashtongue Vision`, 2151.88 DPS.
- DST gap: **+28.59 DPS**.

Practical ruling: **Enhancement belongs in the middle DST pool.** It is below Rogue/Fury/Survival on this pass, but the gap is still real.

### Arms Warrior

Arms is a legitimate mid-pack claimant, but Fury is the higher Warrior DST target.

- Best DST pair: `DST + Brooch`, 1504.25 DPS.
- Best no-DST pair: `Abacus + Direbrew`, 1476.32 DPS.
- DST gap: **+27.93 DPS**.

Practical ruling: **Award Arms after stronger Rogue/Fury/Survival/Ret claims are handled.**

### BM Hunter

BM is valid but lower than Survival.

- Best DST pair: `DST + Brooch`, 2110.25 DPS.
- Best no-DST pair: `Brooch + Tsunami`, 2085.41 DPS.
- DST gap: **+24.85 DPS**.

Practical ruling: **BM can receive DST, but should generally wait behind Survival.**

### Feral Cat

Feral remains the clean low-priority DST claimant.

- Best DST pair: `DST + Tsunami`, 2615.30 DPS.
- Best no-DST pair: `Tsunami + Brooch`, 2608.59 DPS.
- DST gap: **+6.72 DPS**.

Practical ruling: **Do not spend early DST on Feral.** Tsunami/Brooch or other non-DST setups cover almost all of the personal value.

## Madness of the Betrayer Notes

Madness is important in Phase 3, but it does not affect every class the same way in this sim:

- Fury and Enhancement pair DST with Madness.
- Rogue's best DST pair is still `DST + Warp-Spring`; Madness is the best no-DST partner with Warp-Spring.
- Ret, Survival, BM, and Arms do not use Madness in their best DST pair under this 5-minute full-buff profile.
- Several Madness/Brooch/Tsunami differences are small enough that a 10,000-iteration rerun is appropriate before arguing over margins under 3 DPS.

The video pass adds one important warning: Veramos' BM and Survival guides both say Madness replaces Brooch for Hunters, while this 3,000-iteration sim ranks `DST + Brooch` slightly higher. Treat Hunter second-trinket ordering as unresolved until the Hunter rows are rerun and inspected. The DST claim itself remains supported.

The video pass also found mixed no-Warglaives Fury weapon assumptions. The current report uses the Wowhead/sim dual-`Vengeful Gladiator's Slicer` baseline, while Fury videos commonly discuss `Dragon Strike` plus a non-Glaive offhand. Rerun Fury with additional no-Glaives weapon baselines before treating the exact Fury number as final.

## Final Practical Award Plan

Use this if you want a simple loot-council policy:

1. Put **Combat Rogue** and **Fury Warrior** in the first non-Ret DST pool.
2. Include **Survival Hunter** as a high-priority claimant.
3. Include **core/high-execution Ret Paladin** as an early award candidate, but screen hard for player quality.
4. Put **Enhancement Shaman** in the middle priority pool.
5. Treat **Arms Warrior** as valid but below Fury.
6. Treat **BM Hunter** as valid but below Survival.
7. Keep **Feral Cat** off early DST unless the larger claims are already covered.
