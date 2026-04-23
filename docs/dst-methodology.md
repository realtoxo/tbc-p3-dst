# Phase 3 DST Methodology

This document describes how the Phase 3 Dragonspine Trophy priority numbers were produced.

## Item Mechanics

`Dragonspine Trophy` is item `28830`.

Wowhead lists it as:

- +40 Attack Power.
- Chance on melee/ranged attack to gain +325 Haste Rating for 10 seconds.
- 20 second cooldown.

WoWSims models it as:

- Static `AttackPower: 40` and `RangedAttackPower: 40`.
- A `Dragonspine Trophy Proc` temporary aura with `MeleeHaste: 325`.
- 10 second aura duration.
- 20 second internal cooldown.
- 1.0 PPM manager on melee or ranged landed hits.

## Gear Baseline

Gearsets are extracted from Wowhead Phase 3 class/spec BiS guides:

- Arms Warrior
- Fury Warrior
- Combat Rogue
- BM Hunter
- Survival Hunter
- Feral Cat
- Enhancement Shaman
- Retribution Paladin

The extractor writes:

- [sources/wowhead-bis-gearsets.json](../sources/wowhead-bis-gearsets.json)
- [sources/wowhead-bis-items.csv](../sources/wowhead-bis-items.csv)

Manual corrections are recorded in [scripts/extract-wowhead-bis.js](../scripts/extract-wowhead-bis.js). The important corrections are:

- Fury Warrior: no Warglaives, dual `Vengeful Gladiator's Slicer`.
- Combat Rogue: no Warglaives, `Vengeful Gladiator's Slicer` + `Blade of Savagery`.
- Arms Warrior, Ret Paladin, and Feral Cat: manual two-hand weapon selection where the Wowhead section title does not fit the generic extractor perfectly.
- BM Hunter, Survival Hunter, and Enhancement Shaman: manual mainhand/offhand splits where the Wowhead guide lists multiple weapon roles in one table.

## Sim Baseline

Harness: [sources/wowsims-p3-dst-pairs.go](../sources/wowsims-p3-dst-pairs.go)

Outputs:

- [sources/wowsims-p3-dst-summary.csv](../sources/wowsims-p3-dst-summary.csv)
- [sources/wowsims-p3-trinket-pairs.csv](../sources/wowsims-p3-trinket-pairs.csv)
- [sources/final-p3-dst-spreads.csv](../sources/final-p3-dst-spreads.csv)

The harness:

- Loads the Wowhead Phase 3 gearsets.
- Tests trinket pairs rather than fixed one-for-one replacements.
- Chooses the best pair containing DST and compares it to the best pair without DST.
- Runs 3,000 iterations per trinket pair.
- Uses a 5 minute single-target encounter.
- Sets target mob type to neutral/unknown so `Mark of the Champion` is not accidentally active.
- Uses rare-quality gems and normal physical DPS enchants.

## Simulation Assumptions

Use these assumptions when presenting the numbers:

| Assumption | Value |
|---|---|
| Simulator | Local WoWSims TBC Go engine |
| Encounter | 5 minute single-target fight |
| Target type | Neutral/unknown mob type |
| Iterations | 3,000 per trinket pair |
| Output metric | Personal DPS, not raid DPS |
| Raid style | Normalized full physical-DPS support environment |
| Faction | Alliance assumptions |
| Draenei hit aura | Excluded |
| Bloodlust | Enabled in party buffs |
| Drums | Drums of Battle enabled |
| Consumables | WoWSims full-consume preset per class/spec |
| Gear source | Wowhead Phase 3 BiS guide extraction plus documented manual overrides |
| Gems | Rare-quality gems selected by agility or strength profile |
| Enchants | Standard physical DPS enchants selected by slot/profile |
| Warglaives | Excluded by assumption |
| Trinket method | Best pair containing DST vs best pair without DST |
| Survival Hunter raid value | Expose Weakness raid DPS is not credited |
| Support-spec raid value | Ret/Enh/Feral utility is not credited |

The goal is not to reconstruct a legal five-player party. The goal is a consistent trinket opportunity-cost comparison across specs.

Race handling:

- Warriors, Rogue, and Ret use Human.
- Hunters and Feral use Night Elf.
- Enhancement uses a Tauren no-hit proxy because WoWSims automatically grants Draenei hit to Draenei characters. This is a sim workaround, not a Horde assumption.

Buff handling:

- Full raid buffs are enabled.
- Physical support buffs include Bloodlust, Drums, Improved Battle Shout, Improved Leader of the Pack, Grace of Air, Strength of Earth, Trueshot Aura, Ferocious Inspiration, and related physical support where applicable.
- No Draenei hit aura is included.

This is a normalized physical DPS environment, not an exact legal five-player party reconstruction. The goal is trinket opportunity cost, not full raid composition modeling.

## Commands

Regenerate Wowhead gearsets:

```bash
node scripts/extract-wowhead-bis.js
```

Run the sim matrix from the local WoWSims checkout:

```bash
export RESEARCH_REPO="$(pwd)"
export WOWSIMS_TBC="/path/to/wowsims-tbc"

mkdir -p "$WOWSIMS_TBC/cmd/p3dstpairs"
cp "$RESEARCH_REPO/sources/wowsims-p3-dst-pairs.go" "$WOWSIMS_TBC/cmd/p3dstpairs/main.go"
cd "$WOWSIMS_TBC"

DST_GEARSETS_JSON="$RESEARCH_REPO/sources/wowhead-bis-gearsets.json" \
DST_PAIR_OUTPUT="$RESEARCH_REPO/sources/wowsims-p3-trinket-pairs.csv" \
  go run ./cmd/p3dstpairs > "$RESEARCH_REPO/sources/wowsims-p3-dst-summary.csv"

cp "$RESEARCH_REPO/sources/wowsims-p3-dst-summary.csv" "$RESEARCH_REPO/sources/final-p3-dst-spreads.csv"
```

## Caveats

- These are personal DPS deltas, not raid DPS deltas.
- Results use 3,000 iterations, so rerun at 10,000 iterations before arguing over small pair differences.
- Warglaives are excluded by assumption; if a Rogue or Fury Warrior gets Warglaives, rerun the matrix.
- `Mark of the Champion` needs a separate Demon/Undead target pass if LC wants encounter-specific swap guidance.
- Exact raid comp, kill times, player execution, and trinket availability can move individual decisions.
