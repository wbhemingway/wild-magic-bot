package tables

import (
	"github.com/wbhemingway/wild-magic-bot/internal/surge/effects"
)

// GetSurgeEffect2014 takes a d100 roll as input and returns the formatted string
// for the corresponding wild magic surge effect from the 2014 table.
func GetSurgeEffect2014(roll int) string {
	table := getTable2014Ranges()
	for _, entry := range table {
		if roll >= entry.min && roll <= entry.max {
			return entry.effect()
		}
	}
	return "Invalid roll. Please provide a roll between 1 and 100."
}

func getTable2014Ranges() []effectRange {
	return []effectRange{
		{1, 2, func() string {
			return "Roll on this table at the start of each of your turns for the next minute, ignoring this result on subsequent rolls."
		}},
		{3, 4, func() string {
			return "For the next minute, you can see any invisible creature if you have line of sight to it."
		}},
		{5, 6, func() string {
			return "A modron chosen and controlled by the DM appears in an unoccupied space within 5 feet of you, then disappears 1 minute later."
		}},
		{7, 8, func() string { return "You cast Fireball as a 3rd-level spell centered on yourself." }},
		{9, 10, func() string { return "You cast Magic Missile as a 5th-level spell." }},
		{11, 12, effects.HandleHeightChange()},
		{13, 14, func() string { return "You cast Confusion centered on yourself." }},
		{15, 16, func() string {
			return "For the next minute, you regain 5 hit points at the start of each of your turns."
		}},
		{17, 18, func() string {
			return "You grow a long beard made of feathers that remains until you sneeze, at which point the feathers explode out from your face."
		}},
		{19, 20, func() string { return "You cast Grease centered on yourself." }},
		{21, 22, func() string {
			return "Creatures have disadvantage on saving throws against the next spell you cast in the next minute that involves a saving throw."
		}},
		{23, 24, func() string {
			return "Your skin turns a vibrant shade of blue. A remove curse spell can end this effect."
		}},
		{25, 26, func() string {
			return "An eye appears on your forehead for the next minute. During that time, you have advantage on Wisdom (Perception) checks that rely on sight."
		}},
		{27, 28, func() string {
			return "For the next minute, all your spells with a casting time of 1 action have a casting time of 1 bonus action."
		}},
		{29, 30, func() string {
			return "You teleport up to 60 feet to an unoccupied space of your choice that you can see."
		}},
		{31, 32, func() string {
			return "You are transported to the Astral Plane until the end of your next turn, after which time you return to the space you previously occupied or the nearest unoccupied space if that space is occupied."
		}},
		{33, 34, func() string {
			return "Maximize the damage of the next damaging spell you cast within the next minute."
		}},
		{35, 36, effects.HandleAgeChange()},
		{37, 38, func() string {
			return "For the next minute, you regain 1d6 hit points at the start of each of your turns."
		}},
		{39, 40, func() string {
			return "You grow uncontrollably, increasing your size by one category for the next minute."
		}},
		{41, 42, func() string {
			return "You turn into a potted plant until the start of your next turn. While a plant, you are incapacitated and have vulnerability to all damage."
		}},
		{43, 44, func() string {
			return "For the next minute, you can teleport up to 20 feet as a bonus action on each of your turns."
		}},
		{45, 46, func() string { return "You cast Levitate on yourself." }},
		{47, 48, func() string {
			return "A unicorn controlled by the DM appears in a space within 5 feet of you, then disappears 1 minute later."
		}},
		{49, 50, func() string {
			return "You can't speak for the next minute. Whenever you try, pink bubbles float out of your mouth."
		}},
		{51, 52, func() string {
			return "A spectral shield hovers near you for the next minute, granting you a +2 bonus to AC and immunity to magic missile."
		}},
		{53, 54, func() string { return "You are immune to being intoxicated by alcohol for the next 5d6 days." }},
		{55, 56, func() string { return "Your hair falls out but grows back within 24 hours." }},
		{57, 58, func() string {
			return "For the next minute, any flammable object you touch that isn't being worn or carried by another creature bursts into flame."
		}},
		{59, 60, func() string { return "You regain your lowest-level expended spell slot." }},
		{61, 62, func() string { return "For the next minute, you must shout when you speak." }},
		{63, 64, func() string { return "You cast Fog Cloud centered on yourself." }},
		{65, 66, func() string {
			return "Up to three creatures you choose within 30 feet of you take 4d10 lightning damage."
		}},
		{67, 68, func() string { return "You are frightened by the nearest creature until the end of your next turn." }},
		{69, 70, func() string {
			return "Each creature within 30 feet of you becomes invisible for the next minute. The invisibility ends on a creature when it attacks or casts a spell."
		}},
		{71, 72, func() string { return "You gain resistance to all damage for the next minute." }},
		{73, 74, func() string { return "A random creature within 60 feet of you is poisoned for 1d4 hours." }},
		{75, 76, func() string {
			return "You glow with bright light in a 30-foot radius for the next minute. Any creature that ends its turn within 5 feet of you is blinded until the end of its next turn."
		}},
		{77, 78, func() string {
			return "You cast Polymorph on yourself. If you fail the saving throw, you turn into a sheep for the spell's duration."
		}},
		{79, 80, func() string {
			return "Illusory butterflies and flower petals flutter in the air within 10 feet of you for the next minute."
		}},
		{81, 82, func() string { return "You can take one additional action immediately." }},
		{83, 84, func() string {
			return "Each creature within 30 feet of you takes 1d10 necrotic damage. You regain hit points equal to the sum of the damage dealt."
		}},
		{85, 86, func() string { return "You cast Mirror Image." }},
		{87, 88, func() string { return "You cast Fly on a random creature within 60 feet of you." }},
		{89, 90, func() string {
			return "You become invisible for the next minute. During that time, other creatures can't hear you. The invisibility ends if you attack or cast a spell."
		}},
		{91, 92, func() string {
			return "If you die within the next minute, you immediately come back to life as if by the Reincarnate spell."
		}},
		{93, 94, func() string { return "Your size increases by one size category for the next minute." }},
		{95, 96, func() string {
			return "You and all creatures within 30 feet of you gain vulnerability to piercing damage for the next minute."
		}},
		{97, 98, func() string { return "You are surrounded by faint, ethereal music for the next minute." }},
		{99, 100, func() string { return "You regain all expended sorcery points." }},
	}
}
