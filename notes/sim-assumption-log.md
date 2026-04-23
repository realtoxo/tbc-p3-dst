# Phase 3 Sim Assumption Log

## Accepted Assumptions

- Use Wowhead Phase 3 BiS guide gearsets as the starting point.
- Keep Alliance, full physical buffs, and no Draenei hit aura.
- Split BM Hunter from Survival Hunter.
- Split Fury Warrior from Arms Warrior.
- Exclude Warglaives for Fury Warrior and Combat Rogue.
- Continue no-Badge/no-Slayer handling for Warrior trinket candidates.
- Use a trinket-pair matrix rather than a fixed replacement matrix.

## Manual Gear Overrides

- Fury Warrior: `Vengeful Gladiator's Slicer` mainhand and offhand.
- Combat Rogue: `Vengeful Gladiator's Slicer` mainhand and `Blade of Savagery` offhand.
- Arms Warrior: `Cataclysm's Edge`.
- Ret Paladin: `Torch of the Damned`.
- Feral Cat: `Vengeful Gladiator's Staff`.
- BM Hunter: `Boundless Agony` mainhand and `Blade of Infamy` offhand.
- Survival Hunter: dual `Blade of Infamy`.
- Enhancement Shaman: `Vengeful Gladiator's Cleaver` mainhand and `Syphon of the Nathrezim` offhand.

## Follow-Up Work

- Rerun at 10,000 iterations if the report will be used as a final guild policy artifact.
- Class-specialist video corroboration added in [docs/video-corroboration.md](../docs/video-corroboration.md).
- Resolve the Hunter/Madness mismatch: Veramos says Madness replaces Brooch, while the current 3,000-iteration sim ranks Brooch/Tsunami/Abacus slightly ahead in some Hunter pairings.
- Rerun no-Warglaives Fury with `Dragon Strike + Vengeful Slicer` and `Dragon Strike + Blade of Infamy` baselines.
- Add Demon/Undead target passes for `Mark of the Champion` encounter-specific swaps.
