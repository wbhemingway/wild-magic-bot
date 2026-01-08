package surge

import (
	"fmt"
	"math/rand"
)

// Roll rolls a d100 and returns a wild magic surge effect from the specified table.
func Roll(tableName string) (string, error) {
	roll := rand.Intn(100) + 1
	if tableName == "2024" {
		return GetSurgeEffect(roll), nil
	}
	return "", fmt.Errorf("unknown table: %s", tableName)
}
