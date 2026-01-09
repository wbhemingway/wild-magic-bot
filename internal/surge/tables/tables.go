package tables

// Table defines the structure for a surge table.
type Table struct {
	Name        string
	Description string
	Roll        func(roll int) string
}

// DefaultTableName is the name of the default table to use.
const DefaultTableName = "2024"

// AvailableTables is a map of all surge tables available in the application.
var AvailableTables = map[string]Table{
	"2024": {
		Name:        "2024",
		Description: "Table from the DnD 2024 handbook",
		Roll:        GetSurgeEffect2024,
	},
	"2014": {
		Name:        "2014",
		Description: "Table from the DnD 2014 handbook.",
		Roll:        GetSurgeEffect2014,
	},
}
