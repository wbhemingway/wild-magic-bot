package effects

import (
	"strings"
	"testing"
)

func TestRollDie(t *testing.T) {
	sides := 6
	roll := RollDie(sides)
	if roll < 1 || roll > sides {
		t.Errorf("RollDie(%d) returned %d, which is out of the expected range [1, %d]", sides, roll, sides)
	}
}

func TestHandleCreatureSummon(t *testing.T) {
	result := HandleCreatureSummon()
	if !strings.HasPrefix(result, "A creature that is Friendly toward you appears") {
		t.Errorf("Unexpected prefix for HandleCreatureSummon result: %s", result)
	}
	creatures := []string{"Modron Duodrone", "Flumph", "Modron Monodrone", "Unicorn"}
	found := false
	for _, creature := range creatures {
		if strings.Contains(result, creature) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("HandleCreatureSummon result did not contain an expected creature: %s", result)
	}
}

func TestHandleSubjectedEffect(t *testing.T) {
	result := HandleSubjectedEffect()
	if !strings.HasPrefix(result, "You are subjected to an effect that lasts for 1 minute") {
		t.Errorf("Unexpected prefix for HandleSubjectedEffect result: %s", result)
	}
}

func TestHandleRandomSpell(t *testing.T) {
	result := HandleRandomSpell()
	if !strings.HasPrefix(result, "You cast a random spell.") {
		t.Errorf("Unexpected prefix for HandleRandomSpell result: %s", result)
	}
}

func TestHandleLuckyUnlucky(t *testing.T) {
	result := HandleLuckyUnlucky()
	if !strings.HasPrefix(result, "You rolled a") {
		t.Errorf("Unexpected prefix for HandleLuckyUnlucky result: %s", result)
	}
}
