package tables

import (
	"fmt"
	"math/rand"

	"github.com/wbhemingway/wild-magic-bot/internal/surge/effects"
)

func newEffect(r *rand.Rand, f func(*rand.Rand) string) func() string {
	return func() string {
		return f(r)
	}
}

func GetTable2024(r *rand.Rand) map[int]func() string {
	return map[int]func() string{
		1:  func() string { return "Roll on this table at the start of each of your turns for the next minute, ignoring this result on subsequent rolls." },
		2:  func() string { return "Roll on this table at the start of each of your turns for the next minute, ignoring this result on subsequent rolls." },
		3:  func() string { return "Roll on this table at the start of each of your turns for the next minute, ignoring this result on subsequent rolls." },
		4:  func() string { return "Roll on this table at the start of each of your turns for the next minute, ignoring this result on subsequent rolls." },
		5:  newEffect(r, effects.HandleCreatureSummon),
		6:  newEffect(r, effects.HandleCreatureSummon),
		7:  newEffect(r, effects.HandleCreatureSummon),
		8:  newEffect(r, effects.HandleCreatureSummon),
		9:  func() string { return "For the next minute, you regain 5 Hit Points at the start of each of your turns." },
		10: func() string { return "For the next minute, you regain 5 Hit Points at the start of each of your turns." },
		11: func() string { return "For the next minute, you regain 5 Hit Points at the start of each of your turns." },
		12: func() string { return "For the next minute, you regain 5 Hit Points at the start of each of your turns." },
		13: func() string { return "Creatures have Disadvantage on saving throws against the next spell you cast in the next minute that involves a saving throw." },
		14: func() string { return "Creatures have Disadvantage on saving throws against the next spell you cast in the next minute that involves a saving throw." },
		15: func() string { return "Creatures have Disadvantage on saving throws against the next spell you cast in the next minute that involves a saving throw." },
		16: func() string { return "Creatures have Disadvantage on saving throws against the next spell you cast in the next minute that involves a saving throw." },
		17: newEffect(r, effects.HandleSubjectedEffect),
		18: newEffect(r, effects.HandleSubjectedEffect),
		19: newEffect(r, effects.HandleSubjectedEffect),
		20: newEffect(r, effects.HandleSubjectedEffect),
		21: func() string { return "For the next minute, all your spells with a casting time of an action have a casting time of a Bonus Action." },
		22: func() string { return "For the next minute, all your spells with a casting time of an action have a casting time of a Bonus Action." },
		23: func() string { return "For the next minute, all your spells with a casting time of an action have a casting time of a Bonus Action." },
		24: func() string { return "For the next minute, all your spells with a casting time of an action have a casting time of a Bonus Action." },
		25: func() string { return "You are transported to the Astral Plane until the end of your next turn. You then return to the space you previously occupied or the nearest unoccupied space if that space is occupied." },
		26: func() string { return "You are transported to the Astral Plane until the end of your next turn. You then return to the space you previously occupied or the nearest unoccupied space if that space is occupied." },
		27: func() string { return "You are transported to the Astral Plane until the end of your next turn. You then return to the space you previously occupied or the nearest unoccupied space if that space is occupied." },
		28: func() string { return "You are transported to the Astral Plane until the end of your next turn. You then return to the space you previously occupied or the nearest unoccupied space if that space is occupied." },
		29: func() string { return "The next time you cast a spell that deals damage within the next minute, don’t roll the spell’s damage dice for the damage. Instead use the highest number possible for each damage die." },
		30: func() string { return "The next time you cast a spell that deals damage within the next minute, don’t roll the spell’s damage dice for the damage. Instead use the highest number possible for each damage die." },
		31: func() string { return "The next time you cast a spell that deals damage within the next minute, don’t roll the spell’s damage dice for the damage. Instead use the highest number possible for each damage die." },
		32: func() string { return "The next time you cast a spell that deals damage within the next minute, don’t roll the spell’s damage dice for the damage. Instead use the highest number possible for each damage die." },
		33: func() string { return "You have Resistance to all damage for the next minute." },
		34: func() string { return "You have Resistance to all damage for the next minute." },
		35: func() string { return "You have Resistance to all damage for the next minute." },
		36: func() string { return "You have Resistance to all damage for the next minute." },
		37: func() string { return "You turn into a potted plant until the start of your next turn. While you’re a plant, you have the Incapacitated condition and have Vulnerability to all damage. If you drop to 0 Hit Points, your pot breaks, and your form reverts." },
		38: func() string { return "You turn into a potted plant until the start of your next turn. While you’re a plant, you have the Incapacitated condition and have Vulnerability to all damage. If you drop to 0 Hit Points, your pot breaks, and your form reverts." },
		39: func() string { return "You turn into a potted plant until the start of your next turn. While you’re a plant, you have the Incapacitated condition and have Vulnerability to all damage. If you drop to 0 Hit Points, your pot breaks, and your form reverts." },
		40: func() string { return "You turn into a potted plant until the start of your next turn. While you’re a plant, you have the Incapacitated condition and have Vulnerability to all damage. If you drop to 0 Hit Points, your pot breaks, and your form reverts." },
		41: func() string { return "For the next minute, you can teleport up to 20 feet as a Bonus Action on each of your turns." },
		42: func() string { return "For the next minute, you can teleport up to 20 feet as a Bonus Action on each of your turns." },
		43: func() string { return "For the next minute, you can teleport up to 20 feet as a Bonus Action on each of your turns." },
		44: func() string { return "For the next minute, you can teleport up to 20 feet as a Bonus Action on each of your turns." },
		45: func() string { return "You and up to three creatures you choose within 30 feet of you have the Invisible condition for 1 minute. This invisibility ends on a creature immediately after it makes an attack roll, deals damage, or casts a spell." },
		46: func() string { return "You and up to three creatures you choose within 30 feet of you have the Invisible condition for 1 minute. This invisibility ends on a creature immediately after it makes an attack roll, deals damage, or casts a spell." },
		47: func() string { return "You and up to three creatures you choose within 30 feet of you have the Invisible condition for 1 minute. This invisibility ends on a creature immediately after it makes an attack roll, deals damage, or casts a spell." },
		48: func() string { return "You and up to three creatures you choose within 30 feet of you have the Invisible condition for 1 minute. This invisibility ends on a creature immediately after it makes an attack roll, deals damage, or casts a spell." },
		49: func() string { return "A spectral shield hovers near you for the next minute, granting you a +2 bonus to AC and immunity to Magic Missile." },
		50: func() string { return "A spectral shield hovers near you for the next minute, granting you a +2 bonus to AC and immunity to Magic Missile." },
		51: func() string { return "A spectral shield hovers near you for the next minute, granting you a +2 bonus to AC and immunity to Magic Missile." },
		52: func() string { return "A spectral shield hovers near you for the next minute, granting you a +2 bonus to AC and immunity to Magic Missile." },
		53: func() string { return "You can take one extra action on this turn." },
		54: func() string { return "You can take one extra action on this turn." },
		55: func() string { return "You can take one extra action on this turn." },
		56: func() string { return "You can take one extra action on this turn." },
		57: newEffect(r, effects.HandleRandomSpell),
		58: newEffect(r, effects.HandleRandomSpell),
		59: newEffect(r, effects.HandleRandomSpell),
		60: newEffect(r, effects.HandleRandomSpell),
		61: func() string { return fmt.Sprintf("For the next minute, any flammable, nonmagical object you touch that isn’t being worn or carried by another creature bursts into flame, takes %d Fire damage, and is burning.", effects.RollDie(r, 4)) },
		62: func() string { return fmt.Sprintf("For the next minute, any flammable, nonmagical object you touch that isn’t being worn or carried by another creature bursts into flame, takes %d Fire damage, and is burning.", effects.RollDie(r, 4)) },
		63: func() string { return fmt.Sprintf("For the next minute, any flammable, nonmagical object you touch that isn’t being worn or carried by another creature bursts into flame, takes %d Fire damage, and is burning.", effects.RollDie(r, 4)) },
		64: func() string { return fmt.Sprintf("For the next minute, any flammable, nonmagical object you touch that isn’t being worn or carried by another creature bursts into flame, takes %d Fire damage, and is burning.", effects.RollDie(r, 4)) },
		65: func() string { return "If you die within the next hour, you immediately revive as if by the Reincarnate spell." },
		66: func() string { return "If you die within the next hour, you immediately revive as if by the Reincarnate spell." },
		67: func() string { return "If you die within the next hour, you immediately revive as if by the Reincarnate spell." },
		68: func() string { return "If you die within the next hour, you immediately revive as if by the Reincarnate spell." },
		69: func() string { return "You have the Frightened condition until the end of your next turn. The DM determines the source of your fear." },
		70: func() string { return "You have the Frightened condition until the end of your next turn. The DM determines the source of your fear." },
		71: func() string { return "You have the Frightened condition until the end of your next turn. The DM determines the source of your fear." },
		72: func() string { return "You have the Frightened condition until the end of your next turn. The DM determines the source of your fear." },
		73: func() string { return "You teleport up to 60 feet to an unoccupied space you can see." },
		74: func() string { return "You teleport up to 60 feet to an unoccupied space you can see." },
		75: func() string { return "You teleport up to 60 feet to an unoccupied space you can see." },
		76: func() string { return "You teleport up to 60 feet to an unoccupied space you can see." },
		77: func() string { return fmt.Sprintf("A random creature within 60 feet of you has the Poisoned condition for %d hours.", effects.RollDie(r, 4)) },
		78: func() string { return fmt.Sprintf("A random creature within 60 feet of you has the Poisoned condition for %d hours.", effects.RollDie(r, 4)) },
		79: func() string { return fmt.Sprintf("A random creature within 60 feet of you has the Poisoned condition for %d hours.", effects.RollDie(r, 4)) },
		80: func() string { return fmt.Sprintf("A random creature within 60 feet of you has the Poisoned condition for %d hours.", effects.RollDie(r, 4)) },
		81: func() string { return "You radiate Bright Light in a 30-foot radius for the next minute. Any creature that ends its turn within 5 feet of you has the Blinded condition until the end of its next turn." },
		82: func() string { return "You radiate Bright Light in a 30-foot radius for the next minute. Any creature that ends its turn within 5 feet of you has the Blinded condition until the end of its next turn." },
		83: func() string { return "You radiate Bright Light in a 30-foot radius for the next minute. Any creature that ends its turn within 5 feet of you has the Blinded condition until the end of its next turn." },
		84: func() string { return "You radiate Bright Light in a 30-foot radius for the next minute. Any creature that ends its turn within 5 feet of you has the Blinded condition until the end of its next turn." },
		85: func() string { return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Necrotic damage. You regain Hit Points equal to the sum of the Necrotic damage dealt.", effects.RollDie(r, 10)) },
		86: func() string { return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Necrotic damage. You regain Hit Points equal to the sum of the Necrotic damage dealt.", effects.RollDie(r, 10)) },
		87: func() string { return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Necrotic damage. You regain Hit Points equal to the sum of the Necrotic damage dealt.", effects.RollDie(r, 10)) },
		88: func() string { return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Necrotic damage. You regain Hit Points equal to the sum of the Necrotic damage dealt.", effects.RollDie(r, 10)) },
		89: func() string { return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Lightning damage.", 4*effects.RollDie(r, 10)) },
		90: func() string { return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Lightning damage.", 4*effects.RollDie(r, 10)) },
		91: func() string { return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Lightning damage.", 4*effects.RollDie(r, 10)) },
		92: func() string { return fmt.Sprintf("Up to three creatures of your choice that you can see within 30 feet of you take %d Lightning damage.", 4*effects.RollDie(r, 10)) },
		93: func() string { return "You and all creatures within 30 feet of you have Vulnerability to Piercing damage for the next minute." },
		94: func() string { return "You and all creatures within 30 feet of you have Vulnerability to Piercing damage for the next minute." },
		95: func() string { return "You and all creatures within 30 feet of you have Vulnerability to Piercing damage for the next minute." },
		96: func() string { return "You and all creatures within 30 feet of you have Vulnerability to Piercing damage for the next minute." },
		97: newEffect(r, effects.HandleLuckyUnlucky),
		98: newEffect(r, effects.HandleLuckyUnlucky),
		99: newEffect(r, effects.HandleLuckyUnlucky),
		100: newEffect(r, effects.HandleLuckyUnlucky),
	}
}

// GetSurgeEffect takes a d100 roll as input and returns the formatted string
// for the corresponding wild magic surge effect.
func GetSurgeEffect(r *rand.Rand, roll int) string {
	table := GetTable2024(r)
	if effectFunc, ok := table[roll]; ok {
		return effectFunc()
	}
	return "Invalid roll. Please provide a roll between 1 and 100."
}
