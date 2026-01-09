package tables

import (
	"math/rand"
	"testing"
	"time"
)

func TestGetSurgeEffect2014(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Test all valid rolls
	for i := 1; i <= 100; i++ {
		effect := GetSurgeEffect2014(rng, i)
		if effect == "" {
			t.Errorf("GetSurgeEffect2014(%d) returned an empty string", i)
		}
		if effect == "Invalid roll. Please provide a roll between 1 and 100." {
			t.Errorf("GetSurgeEffect2014(%d) returned an invalid roll message", i)
		}
	}

	// Test an invalid roll
	invalidRoll := 101
	effect := GetSurgeEffect2014(rng, invalidRoll)
	if effect != "Invalid roll. Please provide a roll between 1 and 100." {
		t.Errorf("GetSurgeEffect2014(%d) returned %q, expected an invalid roll message", invalidRoll, effect)
	}
}
