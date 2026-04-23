package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	feral "github.com/wowsims/tbc/sim/druid/feral"
	"github.com/wowsims/tbc/sim/hunter"
	ret "github.com/wowsims/tbc/sim/paladin/retribution"
	"github.com/wowsims/tbc/sim/rogue"
	enh "github.com/wowsims/tbc/sim/shaman/enhancement"
	warrior "github.com/wowsims/tbc/sim/warrior/dps"
	googleProto "google.golang.org/protobuf/proto"
)

const (
	Iterations        int32 = 3000
	DST               int32 = 28830
	Brooch            int32 = 29383
	Abacus            int32 = 28288
	Hourglass         int32 = 28034
	Tsunami           int32 = 30627
	WarpSpring        int32 = 30450
	Swarmguard        int32 = 21670
	MarkChamp         int32 = 23206
	KissSpider        int32 = 22954
	Solarian          int32 = 30446
	Direbrew          int32 = 38287
	Madness           int32 = 32505
	AshtongueWarrior  int32 = 32485
	AshtongueHunter   int32 = 32487
	AshtonguePaladin  int32 = 32489
	AshtongueShaman   int32 = 32491
	AshtongueRogue    int32 = 32492
)

type selectedItem struct {
	Slot string `json:"slot"`
	ID   int32  `json:"id"`
	Note string `json:"note"`
}

type gearSet struct {
	Spec     string         `json:"spec"`
	Phase    string         `json:"phase"`
	URL      string         `json:"url"`
	Selected []selectedItem `json:"selected"`
}

type specCase struct {
	Key         string
	Label       string
	Class       proto.Class
	Race        proto.Race
	RaceLabel   string
	GemProfile  string
	Spec        interface{}
	RaidBuffs   *proto.RaidBuffs
	PartyBuffs  *proto.PartyBuffs
	PlayerBuffs *proto.IndividualBuffs
	Consumes    *proto.Consumes
	Debuffs     *proto.Debuffs
	Candidates  []int32
}

type pairResult struct {
	Spec      string
	Race      string
	Trinket1  int32
	Trinket2  int32
	DPS       float64
	HasDST    bool
	UptimePct float64
	URL       string
}

type summaryRow struct {
	Spec          string
	Race          string
	WithDST1      int32
	WithDST2      int32
	WithDSTDPS    float64
	WithoutDST1   int32
	WithoutDST2   int32
	WithoutDSTDPS float64
	Delta         float64
	Pct           float64
	UptimePct     float64
	URL           string
}

func main() {
	hunter.RegisterHunter()
	rogue.RegisterRogue()
	enh.RegisterEnhancementShaman()
	ret.RegisterRetributionPaladin()
	feral.RegisterFeralDruid()
	warrior.RegisterDpsWarrior()

	gearsetsPath := os.Getenv("DST_GEARSETS_JSON")
	if gearsetsPath == "" {
		gearsetsPath = "sources/wowhead-bis-gearsets.json"
	}
	gearsets := loadGearsets(gearsetsPath)

	var pairsOut []pairResult
	var summaries []summaryRow
	for caseIndex, c := range cases() {
		gearset, ok := gearsets[c.Key]
		if !ok {
			panic(fmt.Sprintf("missing gearset %s", c.Key))
		}

		var results []pairResult
		for _, pair := range trinketPairs(c.Candidates) {
			seed := int64(2026042300 + caseIndex)
			dps, uptime := run(c, gearset, pair[0], pair[1], seed)
			row := pairResult{
				Spec:      c.Label,
				Race:      c.RaceLabel,
				Trinket1:  pair[0],
				Trinket2:  pair[1],
				DPS:       dps,
				HasDST:    pair[0] == DST || pair[1] == DST,
				UptimePct: uptime / 300 * 100,
				URL:       gearset.URL,
			}
			results = append(results, row)
			pairsOut = append(pairsOut, row)
		}

		sort.SliceStable(results, func(i, j int) bool { return results[i].DPS > results[j].DPS })
		var withDST *pairResult
		var withoutDST *pairResult
		for i := range results {
			if results[i].HasDST && withDST == nil {
				withDST = &results[i]
			}
			if !results[i].HasDST && withoutDST == nil {
				withoutDST = &results[i]
			}
		}
		if withDST == nil || withoutDST == nil {
			panic(fmt.Sprintf("missing comparison rows for %s", c.Label))
		}
		summaries = append(summaries, summaryRow{
			Spec:          c.Label,
			Race:          c.RaceLabel,
			WithDST1:      withDST.Trinket1,
			WithDST2:      withDST.Trinket2,
			WithDSTDPS:    withDST.DPS,
			WithoutDST1:   withoutDST.Trinket1,
			WithoutDST2:   withoutDST.Trinket2,
			WithoutDSTDPS: withoutDST.DPS,
			Delta:         withDST.DPS - withoutDST.DPS,
			Pct:           (withDST.DPS - withoutDST.DPS) / withoutDST.DPS * 100,
			UptimePct:     withDST.UptimePct,
			URL:           gearset.URL,
		})
	}

	sort.SliceStable(summaries, func(i, j int) bool { return summaries[i].Delta > summaries[j].Delta })
	writeSummary(summaries)
	writePairs(pairsOut)
}

func cases() []specCase {
	return []specCase{
		meleeCase("fury_warrior_p3", "Fury Warrior", proto.Class_ClassWarrior, proto.Race_RaceHuman, "Human", "strength", warrior.PlayerOptionsFury, warrior.FullConsumes, warrior.FullDebuffs, []int32{DST, Madness, Tsunami, Brooch, Abacus, Hourglass, Solarian, Direbrew, AshtongueWarrior}),
		meleeCase("arms_warrior_p3", "Arms Warrior", proto.Class_ClassWarrior, proto.Race_RaceHuman, "Human", "strength", warrior.PlayerOptionsArmsSlam, warrior.FullConsumes, warrior.FullDebuffs, []int32{DST, Madness, Tsunami, Brooch, Abacus, Hourglass, Solarian, Direbrew, AshtongueWarrior}),
		meleeCase("rogue_p3", "Combat Rogue", proto.Class_ClassRogue, proto.Race_RaceHuman, "Human", "agility", rogue.PlayerOptionsBasic, rogue.FullConsumes, rogue.FullDebuffs, []int32{DST, WarpSpring, Madness, AshtongueRogue, Brooch, Tsunami, Abacus, Hourglass, KissSpider}),
		hunterCase("bm_hunter_p3", "BM Hunter", hunter.PlayerOptionsBasic, []int32{DST, Madness, Brooch, AshtongueHunter, Tsunami, Abacus, Hourglass, Swarmguard}),
		hunterCase("survival_hunter_p3", "Survival Hunter", hunter.PlayerOptionsSV, []int32{DST, Madness, Brooch, AshtongueHunter, Tsunami, Abacus, Hourglass, Swarmguard}),
		meleeCase("feral_cat_p3", "Feral Cat", proto.Class_ClassDruid, proto.Race_RaceNightElf, "Night Elf", "agility", feral.PlayerOptionsBiteweave, feral.FullConsumes, feral.FullDebuffs, []int32{DST, Madness, Tsunami, Brooch, Direbrew, Abacus, Hourglass, MarkChamp}),
		meleeCase("enhancement_shaman_p3", "Enhancement Shaman", proto.Class_ClassShaman, proto.Race_RaceTauren, "Alliance Shaman no-hit proxy", "strength", enh.PlayerOptionsBasic, enh.FullConsumes, enh.FullDebuffs, []int32{DST, Madness, AshtongueShaman, Brooch, Tsunami, Abacus, Hourglass, Direbrew}),
		meleeCase("ret_paladin_p3", "Retribution Paladin", proto.Class_ClassPaladin, proto.Race_RaceHuman, "Human", "strength", ret.DefaultOptions, ret.FullConsumes, ret.FullDebuffs, []int32{DST, Brooch, Madness, AshtonguePaladin, Abacus, Tsunami, Direbrew, Hourglass, MarkChamp}),
	}
}

func meleeCase(key, label string, class proto.Class, race proto.Race, raceLabel string, gemProfile string, spec interface{}, consumes *proto.Consumes, debuffs *proto.Debuffs, candidates []int32) specCase {
	return specCase{
		Key:         key,
		Label:       label,
		Class:       class,
		Race:        race,
		RaceLabel:   raceLabel,
		GemProfile:  gemProfile,
		Spec:        spec,
		RaidBuffs:   fullRaidBuffs(),
		PartyBuffs:  fullPhysicalPartyBuffs(),
		PlayerBuffs: fullIndividualBuffs(),
		Consumes:    googleProto.Clone(consumes).(*proto.Consumes),
		Debuffs:     googleProto.Clone(debuffs).(*proto.Debuffs),
		Candidates:  candidates,
	}
}

func hunterCase(key, label string, spec interface{}, candidates []int32) specCase {
	return specCase{
		Key:         key,
		Label:       label,
		Class:       proto.Class_ClassHunter,
		Race:        proto.Race_RaceNightElf,
		RaceLabel:   "Night Elf",
		GemProfile:  "agility",
		Spec:        spec,
		RaidBuffs:   fullRaidBuffs(),
		PartyBuffs:  fullHunterPartyBuffs(),
		PlayerBuffs: fullIndividualBuffs(),
		Consumes:    googleProto.Clone(hunter.FullConsumes).(*proto.Consumes),
		Debuffs:     googleProto.Clone(hunter.FullDebuffs).(*proto.Debuffs),
		Candidates:  candidates,
	}
}

func fullRaidBuffs() *proto.RaidBuffs {
	return &proto.RaidBuffs{
		ArcaneBrilliance: true,
		GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
		DivineSpirit:     proto.TristateEffect_TristateEffectImproved,
	}
}

func fullIndividualBuffs() *proto.IndividualBuffs {
	return &proto.IndividualBuffs{
		BlessingOfKings:     true,
		BlessingOfWisdom:    proto.TristateEffect_TristateEffectImproved,
		BlessingOfMight:     proto.TristateEffect_TristateEffectImproved,
		BlessingOfSalvation: true,
		UnleashedRage:       true,
	}
}

func fullPhysicalPartyBuffs() *proto.PartyBuffs {
	return &proto.PartyBuffs{
		Bloodlust:            1,
		Drums:                proto.Drums_DrumsOfBattle,
		BattleShout:          proto.TristateEffect_TristateEffectImproved,
		LeaderOfThePack:      proto.TristateEffect_TristateEffectImproved,
		GraceOfAirTotem:      proto.TristateEffect_TristateEffectImproved,
		ManaSpringTotem:      proto.TristateEffect_TristateEffectRegular,
		StrengthOfEarthTotem: proto.StrengthOfEarthType_EnhancingTotems,
		WindfuryTotemRank:    5,
		WindfuryTotemIwt:     2,
		SanctityAura:         proto.TristateEffect_TristateEffectImproved,
		TrueshotAura:         true,
		FerociousInspiration: 2,
		BraidedEterniumChain: true,
	}
}

func fullHunterPartyBuffs() *proto.PartyBuffs {
	return &proto.PartyBuffs{
		Bloodlust:            1,
		Drums:                proto.Drums_DrumsOfBattle,
		BattleShout:          proto.TristateEffect_TristateEffectImproved,
		LeaderOfThePack:      proto.TristateEffect_TristateEffectImproved,
		GraceOfAirTotem:      proto.TristateEffect_TristateEffectImproved,
		ManaSpringTotem:      proto.TristateEffect_TristateEffectRegular,
		StrengthOfEarthTotem: proto.StrengthOfEarthType_EnhancingTotems,
		TrueshotAura:         true,
		FerociousInspiration: 2,
	}
}

func loadGearsets(path string) map[string]gearSet {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var gearsets map[string]gearSet
	if err := json.Unmarshal(data, &gearsets); err != nil {
		panic(err)
	}
	return gearsets
}

func run(c specCase, gearset gearSet, trinket1 int32, trinket2 int32, seed int64) (float64, float64) {
	gear := buildGear(c, gearset, trinket1, trinket2)
	raid := core.SinglePlayerRaidProto(
		core.WithSpec(&proto.Player{
			Race:      c.Race,
			Class:     c.Class,
			Equipment: gear,
			Consumes:  c.Consumes,
			Buffs:     c.PlayerBuffs,
		}, c.Spec),
		googleProto.Clone(c.PartyBuffs).(*proto.PartyBuffs),
		googleProto.Clone(c.RaidBuffs).(*proto.RaidBuffs),
		googleProto.Clone(c.Debuffs).(*proto.Debuffs),
	)

	encounter := core.MakeSingleTargetEncounter(5)
	encounter.Targets[0].MobType = proto.MobType_MobTypeUnknown

	result := core.RunRaidSim(&proto.RaidSimRequest{
		Raid:      raid,
		Encounter: encounter,
		SimOptions: &proto.SimOptions{
			Iterations: Iterations,
			IsTest:     true,
			RandomSeed: seed,
		},
	})

	player := result.RaidMetrics.Parties[0].Players[0]
	uptime := 0.0
	for _, aura := range player.Auras {
		if aura.Id != nil && aura.Id.GetItemId() == DST {
			uptime = aura.UptimeSecondsAvg
		}
	}
	return player.Dps.Avg, uptime
}

func buildGear(c specCase, gearset gearSet, trinket1 int32, trinket2 int32) *proto.EquipmentSpec {
	var eq items.EquipmentSpec
	for _, selected := range gearset.Selected {
		slot, ok := slotFor(selected.Slot)
		if !ok {
			continue
		}
		id := selected.ID
		if selected.Slot == "trinket1" {
			id = trinket1
		}
		if selected.Slot == "trinket2" {
			id = trinket2
		}
		eq[slot] = itemSpecFor(c, slot, id)
	}
	eq[items.ItemSlotTrinket1] = itemSpecFor(c, items.ItemSlotTrinket1, trinket1)
	eq[items.ItemSlotTrinket2] = itemSpecFor(c, items.ItemSlotTrinket2, trinket2)
	equipment := items.NewEquipmentSet(eq)
	return equipment.ToEquipmentSpecProto()
}

func itemSpecFor(c specCase, slot items.ItemSlot, id int32) items.ItemSpec {
	item, ok := items.ByID[id]
	if !ok {
		panic(fmt.Sprintf("unknown item id %d in %s", id, c.Key))
	}
	return items.ItemSpec{
		ID:      id,
		Enchant: enchantFor(c, slot, item),
		Gems:    gemsFor(c.GemProfile, item),
	}
}

func slotFor(slot string) (items.ItemSlot, bool) {
	switch slot {
	case "head":
		return items.ItemSlotHead, true
	case "neck":
		return items.ItemSlotNeck, true
	case "shoulder":
		return items.ItemSlotShoulder, true
	case "back":
		return items.ItemSlotBack, true
	case "chest":
		return items.ItemSlotChest, true
	case "wrist":
		return items.ItemSlotWrist, true
	case "hands":
		return items.ItemSlotHands, true
	case "waist":
		return items.ItemSlotWaist, true
	case "legs":
		return items.ItemSlotLegs, true
	case "feet":
		return items.ItemSlotFeet, true
	case "ring1":
		return items.ItemSlotFinger1, true
	case "ring2":
		return items.ItemSlotFinger2, true
	case "trinket1":
		return items.ItemSlotTrinket1, true
	case "trinket2":
		return items.ItemSlotTrinket2, true
	case "main_hand", "two_hand":
		return items.ItemSlotMainHand, true
	case "off_hand":
		return items.ItemSlotOffHand, true
	case "ranged":
		return items.ItemSlotRanged, true
	default:
		return 0, false
	}
}

func enchantFor(c specCase, slot items.ItemSlot, item items.Item) int32 {
	switch slot {
	case items.ItemSlotHead:
		return 29192
	case items.ItemSlotShoulder:
		return 28888
	case items.ItemSlotBack:
		return 34004
	case items.ItemSlotChest:
		return 24003
	case items.ItemSlotWrist:
		if c.GemProfile == "agility" {
			return 34002
		}
		return 27899
	case items.ItemSlotHands:
		if c.GemProfile == "agility" {
			return 33152
		}
		return 33995
	case items.ItemSlotLegs:
		return 29535
	case items.ItemSlotFeet:
		return 28279
	case items.ItemSlotMainHand, items.ItemSlotOffHand:
		if c.Class == proto.Class_ClassHunter {
			if item.HandType == proto.HandType_HandTypeTwoHand {
				return 22556
			}
			return 33165
		}
		if c.Class == proto.Class_ClassDruid {
			return 22556
		}
		return 22559
	case items.ItemSlotRanged:
		if c.Class == proto.Class_ClassHunter {
			return 23766
		}
	}
	return 0
}

func gemsFor(profile string, item items.Item) []int32 {
	gems := make([]int32, 0, len(item.GemSockets))
	for _, socket := range item.GemSockets {
		switch socket {
		case proto.GemColor_GemColorMeta:
			gems = append(gems, 32409)
		case proto.GemColor_GemColorBlue:
			if profile == "agility" {
				gems = append(gems, 24055)
			} else {
				gems = append(gems, 24054)
			}
		case proto.GemColor_GemColorYellow:
			if profile == "agility" {
				gems = append(gems, 24061)
			} else {
				gems = append(gems, 24058)
			}
		default:
			if profile == "agility" {
				gems = append(gems, 24028)
			} else {
				gems = append(gems, 24027)
			}
		}
	}
	return gems
}

func trinketPairs(ids []int32) [][2]int32 {
	ids = uniqueIDs(ids)
	var out [][2]int32
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			out = append(out, [2]int32{ids[i], ids[j]})
		}
	}
	return out
}

func uniqueIDs(ids []int32) []int32 {
	seen := map[int32]bool{}
	var out []int32
	for _, id := range ids {
		if seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, id)
	}
	return out
}

func itemName(id int32) string {
	if item, ok := items.ByID[id]; ok {
		return fmt.Sprintf("%s (%d)", item.Name, id)
	}
	return fmt.Sprintf("Item %d", id)
}

func pairName(a, b int32) string {
	return fmt.Sprintf("%s + %s", itemName(a), itemName(b))
}

func writeSummary(rows []summaryRow) {
	w := csv.NewWriter(os.Stdout)
	_ = w.Write([]string{"spec", "race", "with_dst_pair", "no_dst_pair", "with_dst_dps", "no_dst_dps", "dst_value_dps", "dst_value_pct", "dst_uptime_pct", "wowhead_gear_url"})
	for _, row := range rows {
		_ = w.Write([]string{
			row.Spec,
			row.Race,
			pairName(row.WithDST1, row.WithDST2),
			pairName(row.WithoutDST1, row.WithoutDST2),
			fmt.Sprintf("%.2f", row.WithDSTDPS),
			fmt.Sprintf("%.2f", row.WithoutDSTDPS),
			fmt.Sprintf("%.2f", row.Delta),
			fmt.Sprintf("%.2f", row.Pct),
			fmt.Sprintf("%.2f", row.UptimePct),
			row.URL,
		})
	}
	w.Flush()
}

func writePairs(rows []pairResult) {
	outputPath := os.Getenv("DST_PAIR_OUTPUT")
	if outputPath == "" {
		outputPath = "sources/wowsims-p3-trinket-pairs.csv"
	}
	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	_ = w.Write([]string{"spec", "race", "trinket_pair", "has_dst", "dps", "dst_uptime_pct", "wowhead_gear_url"})
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].Spec == rows[j].Spec {
			return rows[i].DPS > rows[j].DPS
		}
		return rows[i].Spec < rows[j].Spec
	})
	for _, row := range rows {
		_ = w.Write([]string{
			row.Spec,
			row.Race,
			pairName(row.Trinket1, row.Trinket2),
			fmt.Sprintf("%t", row.HasDST),
			fmt.Sprintf("%.2f", row.DPS),
			fmt.Sprintf("%.2f", row.UptimePct),
			row.URL,
		})
	}
	w.Flush()
}
