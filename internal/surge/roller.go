package surge

import (
	"fmt"
	"math/rand"

	"github.com/wbhemingway/wild-magic-bot/internal/surge/tables"
)

// Roll rolls a d100 and returns the roll, a wild magic surge effect, and an error.
func Roll(r *rand.Rand, tableName string) (int, string, error) {
	if tableName == "" {
		tableName = tables.DefaultTableName
	}

	table, ok := tables.AvailableTables[tableName]
	if !ok {
		return 0, "", fmt.Errorf("unknown table: %s", tableName)
	}

	roll := r.Intn(100) + 1
	return roll, table.Roll(r, roll), nil
}
