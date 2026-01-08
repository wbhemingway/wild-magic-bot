package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/wbhemingway/wild-magic-bot/internal/surge"
)

func main() {
	tableName := "2024"
	if len(os.Args) > 1 {
		tableName = os.Args[1]
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	effect, err := surge.Roll(rng, tableName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(effect)
}
