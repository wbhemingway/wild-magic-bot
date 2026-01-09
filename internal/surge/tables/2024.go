package tables

import (
	"fmt"

	"github.com/wbhemingway/wild-magic-bot/internal/surge/effects"
)

func newEffect(f func() string) func() string {
	return func() string {
		return f()
	}
}

type effectRange struct {
	min, max int
	effect   func() string
}

func getTable2024Ranges() []effectRange {
	return []effectRange{
		{1, 4, func() string {
			return "Roll on this table at the start of each of your turns for the next minute, ignoring this result on subsequent rolls."
		}},
		{5, 8, newEffect(effects.HandleCreatureSummon)},
		{9, 12, func() string {
			return "For the next minute, you regain 5 Hit Points at the start of each of your turns."
		}},
		{13, 16, func() string {
			return "Creatures have Disadvantage on saving throws against the next spell you cast in the next minute that involves a saving throw."
		}},
		{17, 20, newEffect(effects.HandleSubjectedEffect)},
		{21, 24, func() string {
			return "For the next minute, all your spells with a casting time of an action have a casting time of a Bonus Action."
		}},
		{25, 28, func() string {
			return "You are transported to the Astral Plane until the end of your next turn. You then return to the space you previously occupied or the nearest unoccupied space if that space is occupied."
		}},
		{29, 32, func() string {
			return "The next time you cast a spell that deals damage within the next minute, don’t roll the spell’s damage dice for the damage. Instead use the highest number possible for each damage die."
		}},
		{33, 36, func() string { return "You have Resistance to all damage for the next minute." }},
		{37, 40, func() string {
			return "You turn into a potted plant until the start of your next turn. While you’re a plant, you have the Incapacitated condition and have Vulnerability to all damage. If you drop to 0 Hit Points, your pot breaks, and your form reverts."
		}},
		{41, 44, func() string {
			return "For the next minute, you can teleport up to 20 feet as a Bonus Action on each of your turns."
		}},
		{45, 48, func() string {
			return "You and up to three creatures you choose within 30 feet of you have the Invisible condition for 1 minute. This invisibility ends on a creature immediately after it makes an attack roll, deals damage, or casts a spell."
		}},
		{49, 52, func() string {
			return "A spectral shield hovers near you for the next minute, granting you a +2 bonus to AC and immunity to Magic Missile."
		}},
		{53, 56, func() string { return "You can take one extra action on this turn." }},
		{57, 60, newEffect(effects.HandleRandomSpell)},
		{61, 64, func() string {
			return fmt.Sprintf("For the next minute, any flammable, nonmagical object you touch that isn’t being worn or carried by another creature bursts into flame, takes %d Fire damage, and is burning.", effects.RollDie(4))
		}},
		{65, 68, func() string {
			return "If you die within the next hour, you immediately revive as if by the Reincarnate spell."
		}},
		{69, 72, func() string {
			return "You have the Frightened condition until the end of your next turn. The DM determines the source of your fear."
		}},
		{73, 76, func() string { return "You teleport up to 60 feet to an unoccupied space you can see." }},
		{77, 80, func() string {
			return fmt.Sprintf("A random creature within 60 feet of you has the Poisoned condition for %d hours.", effects.RollDie(4))
		}},
		{81, 84, func() string {
			return "You radiate Bright Light in a 30-foot radius for the next minute. Any creature that ends its turn within 5 feet of you has the Blinded condition until the end of its next turn."
		}},
		{85, 88, func() string {
			return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Necrotic damage. You regain Hit Points equal to the sum of the Necrotic damage dealt.", effects.RollDie(10))
		}},
		{89, 92, func() string {
			return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Lightning damage.", 4*effects.RollDie(10))
		}},
		{93, 96, func() string {
			return "You and all creatures within 30 feet of you have Vulnerability to Piercing damage for the next minute."
		}},
		{97, 100, newEffect(effects.HandleLuckyUnlucky)},
	}
}

// GetSurgeEffect takes a d100 roll as input and returns the formatted string
// for the corresponding wild magic surge effect.
func GetSurgeEffect2024(roll int) string {
	table := getTable2024Ranges()
	for _, entry := range table {
		if roll >= entry.min && roll <= entry.max {
			return entry.effect()
		}
	}
	return "Invalid roll. Please provide a roll between 1 and 100."
}
