package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/wbhemingway/wild-magic-bot/internal/config"
	"github.com/wbhemingway/wild-magic-bot/internal/surge/tables"
)

// Command represents a Discord application command.
type Command struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Options     []Option `json:"options,omitempty"`
}

// Option represents an option for a command.
type Option struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        int      `json:"type"` // 3 for STRING, 4 for INTEGER
	Required    bool     `json:"required"`
	Choices     []Choice `json:"choices,omitempty"`
}

// Choice represents a choice for a command option.
type Choice struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

var reconfigure = flag.Bool("reconfigure", false, "Force the interactive setup to re-enter credentials.")

func main() {
	flag.Parse()

	var cfg config.Config
	var err error

	if *reconfigure {
		cfg, err = config.InteractiveSetup()
	} else {
		cfg, err = config.Load()
	}

	if err != nil {
		log.Println("Could not load configuration. Starting interactive setup...")
		cfg, err = config.InteractiveSetup()
		if err != nil {
			log.Fatalf("Failed to setup configuration: %v", err)
		}
	}

	// Dynamically build table choices
	var tableChoices []Choice
	for _, table := range tables.AvailableTables {
		tableChoices = append(tableChoices, Choice{
			Name:  fmt.Sprintf("%s - %s", table.Name, table.Description),
			Value: table.Name,
		})
	}

	// Define the /roll command
	commands := []Command{
		{
			Name:        "roll",
			Description: "Roll on the Wild Magic Surge table.",
			Options: []Option{
				{
					Name:        "count",
					Description: "The number of times to roll (1-5). Defaults to 1.",
					Type:        4, // INTEGER
					Required:    false,
				},
				{
					Name:        "table",
					Description: fmt.Sprintf("The surge table to use. Defaults to %s.", tables.DefaultTableName),
					Type:        3, // STRING
					Required:    false,
					Choices:     tableChoices,
				},
			},
		},
	}

	// Register the commands
	url := fmt.Sprintf("https://discord.com/api/v10/applications/%s/commands", cfg.AppID)

	body, err := json.Marshal(commands)
	if err != nil {
		log.Fatalf("Error marshalling commands: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+cfg.BotToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Print the response
	log.Printf("Status Code: %d", resp.StatusCode)
	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}

	prettyResult, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(prettyResult))
	log.Println("Successfully registered commands!")
}
