package main

import (
	"fmt"
	"os"

	"github.com/wbhemingway/wild-magic-bot/internal/surge"
)

func main() {
	tableName := "2024"
	if len(os.Args) > 1 {
		tableName = os.Args[1]
	}

	effect, err := surge.Roll(tableName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(effect)
}
