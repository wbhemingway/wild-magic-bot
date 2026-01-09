package surge

import (
	"fmt"
	"math/rand"

	"github.com/wbhemingway/wild-magic-bot/internal/surge/tables"
)

var surgeTables = map[string]func(*rand.Rand, int) string{
	"2024": tables.GetSurgeEffect,
	"2014": tables.GetSurgeEffect2014,
}

// Roll rolls a d100 and returns the roll, a wild magic surge effect, and an error.
func Roll(r *rand.Rand, tableName string) (int, string, error) {
	roll := r.Intn(100) + 1
	if tableFunc, ok := surgeTables[tableName]; ok {
		return roll, tableFunc(r, roll), nil
	}
	return 0, "", fmt.Errorf("unknown table: %s", tableName)
}
