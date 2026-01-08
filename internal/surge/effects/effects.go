package effects

import (
	"fmt"
	"math/rand"
)

// rollDie simulates rolling a single die with 'sides' number of faces.
func RollDie(r *rand.Rand, sides int) int {
	return r.Intn(sides) + 1
}

// HandleCreatureSummon rolls 1d4 and returns a formatted string with the determined creature.
func HandleCreatureSummon(r *rand.Rand) string {
	roll := RollDie(r, 4)
	creature := ""
	switch roll {
	case 1:
		creature = "Modron Duodrone"
	case 2:
		creature = "Flumph"
	case 3:
		creature = "Modron Monodrone"
	case 4:
		creature = "Unicorn"
	}
	return fmt.Sprintf("A creature that is Friendly toward you appears in a random unoccupied space within 60 feet of you. The creature is under the DM’s control and disappears 1 minute later. You rolled a %d, summoning a %s. See the Monster Manual for the creature’s stat block.", roll, creature)
}

// HandleSubjectedEffect rolls 1d8 and returns a formatted string with the determined effect.
func HandleSubjectedEffect(r *rand.Rand) string {
	roll := RollDie(r, 8)
	effect := ""
	switch roll {
	case 1:
		effect = "you’re surrounded by faint, ethereal music only you and creatures within 5 feet of you can hear"
	case 2:
		effect = "your size increases by one size category"
	case 3:
		effect = "you grow a long beard made of feathers that remains until you sneeze, at which point the feathers explode from your face and vanish"
	case 4:
		effect = "you must shout when you speak"
	case 5:
		effect = "illusory butterflies flutter in the air within 10 feet of you"
	case 6:
		effect = "an eye appears on your forehead, granting you Advantage on Wisdom (Perception) checks"
	case 7:
		effect = "pink bubbles float out of your mouth whenever you speak"
	case 8:
		effect = "your skin turns a vibrant shade of blue for 24 hours or until the effect is ended by a Remove Curse spell"
	}
	return fmt.Sprintf("You are subjected to an effect that lasts for 1 minute unless its description says otherwise. You rolled a %d: %s.", roll, effect)
}

// HandleRandomSpell rolls 1d10 and returns a formatted string with the determined spell.
func HandleRandomSpell(r *rand.Rand) string {
	roll := RollDie(r, 10)
	spell := ""
	switch roll {
	case 1:
		spell = "Confusion"
	case 2:
		spell = "Fireball"
	case 3:
		spell = "Fog Cloud"
	case 4:
		spell = "Fly (cast on a random creature within 60 feet of you)"
	case 5:
		spell = "Grease"
	case 6:
		spell = "Levitate (cast on yourself)"
	case 7:
		spell = "Magic Missile (cast as a level 5 spell)"
	case 8:
		spell = "Mirror Image"
	case 9:
		spell = "Polymorph (cast on yourself), and if you fail the saving throw, you turn into a Goat (see appendix B)"
	case 10:
		spell = "See Invisibility"
	}
	return fmt.Sprintf("You cast a random spell. If the spell normally requires Concentration, it doesn’t require Concentration in this case; the spell lasts for its full duration. You rolled a %d, casting %s.", roll, spell)
}

// HandleLuckyUnlucky rolls 1d6 and returns a formatted string with the determined effect.
func HandleLuckyUnlucky(r *rand.Rand) string {
	roll := RollDie(r, 6)
	effect := ""
	switch roll {
	case 1:
		effect = fmt.Sprintf("you regain %d Hit Points", RollDie(r, 10)+RollDie(r, 10))
	case 2:
		effect = fmt.Sprintf("one ally of your choice within 300 feet of you regains %d Hit Points", RollDie(r, 10)+RollDie(r, 10))
	case 3:
		effect = "you regain your lowest-level expended spell slot"
	case 4:
		effect = "one ally of your choice within 300 feet of you regains their lowest-level expended spell slot"
	case 5:
		effect = "you regain all your expended Sorcery Points"
	case 6:
		effect = "all the effects of row 17–20 affect you simultaneously"
	}
	return fmt.Sprintf("You rolled a %d: %s.", roll, effect)
}
