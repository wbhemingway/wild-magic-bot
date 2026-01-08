package surge

import (
	"fmt"
	"math/rand"

	"github.com/wbhemingway/wild-magic-bot/internal/surge/tables"
)

var surgeTables = map[string]func(*rand.Rand, int) string{
	"2024": tables.GetSurgeEffect,
}

// Roll rolls a d100 and returns a wild magic surge effect from the specified table.
func Roll(r *rand.Rand, tableName string) (string, error) {
	roll := r.Intn(100) + 1
	if tableFunc, ok := surgeTables[tableName]; ok {
		return tableFunc(r, roll), nil
	}
	return "", fmt.Errorf("unknown table: %s", tableName)
}
