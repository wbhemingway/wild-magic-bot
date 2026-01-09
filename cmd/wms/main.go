package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/wbhemingway/wild-magic-bot/internal/surge"
)

var (
	tableName = flag.String("table", "2024", "The surge table to use (e.g., 2024, 2014).")
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

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < *rollCount; i++ {
		roll, effect, err := surge.Roll(rng, *tableName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("(Roll: %d) %s\n", roll, effect)
	}
}
