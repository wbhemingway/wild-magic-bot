package surge

import (
	"testing"
)

func TestRoll(t *testing.T) {
	// Test with a valid table name
	tableName := "2024"
	_, effect, err := Roll(tableName)
	if err != nil {
		t.Fatalf("Expected no error for table %q, but got %v", tableName, err)
	}
	if effect == "" {
		t.Errorf("Expected an effect for table %q, but got an empty string", tableName)
	}

	// Test with an invalid table name
	invalidTableName := "invalid-table"
	_, _, err = Roll(invalidTableName)
	if err == nil {
		t.Fatalf("Expected an error for table %q, but got nil", invalidTableName)
	}
}
