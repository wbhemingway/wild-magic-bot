package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wbhemingway/wild-magic-bot/internal/surge"
	"github.com/wbhemingway/wild-magic-bot/internal/surge/tables"
)

var (
	tableName = flag.String("table", tables.DefaultTableName, "The surge table to use (e.g., 2024, 2014).")
	rollCount = flag.Int("count", 1, "The number of times to roll (1-5).")
)

func main() {
	flag.Parse()

	if *rollCount > 5 {
		*rollCount = 5
	}
	if *rollCount < 1 {
		*rollCount = 1
	}

	for i := 0; i < *rollCount; i++ {
		roll, effect, err := surge.Roll(*tableName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("(Roll: %d) %s\n", roll, effect)
	}
}
