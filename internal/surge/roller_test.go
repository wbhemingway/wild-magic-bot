package surge

import (
	"math/rand"
	"testing"
	"time"
)

func TestRoll(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Test with a valid table name
	tableName := "2024"
	effect, err := Roll(rng, tableName)
	if err != nil {
		t.Fatalf("Expected no error for table %q, but got %v", tableName, err)
	}
	if effect == "" {
		t.Errorf("Expected an effect for table %q, but got an empty string", tableName)
	}

	// Test with an invalid table name
	invalidTableName := "invalid-table"
	_, err = Roll(rng, invalidTableName)
	if err == nil {
		t.Fatalf("Expected an error for table %q, but got nil", invalidTableName)
	}
}
