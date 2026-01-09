package discord

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/wbhemingway/wild-magic-bot/internal/surge"
)

type InteractionType int

const (
	InteractionTypePing               InteractionType = 1
	InteractionTypeApplicationCommand InteractionType = 2
)

type Interaction struct {
	Type InteractionType                   `json:"type"`
	Data ApplicationCommandInteractionData `json:"data"`
}

type ApplicationCommandInteractionData struct {
	Name    string                                    `json:"name"`
	Options []ApplicationCommandInteractionDataOption `json:"options"`
}

type ApplicationCommandInteractionDataOption struct {
	Name  string      `json:"name"`
	Type  int         `json:"type"`
	Value interface{} `json:"value"`
}

type InteractionResponseType int

const (
	InteractionResponseTypePong                     InteractionResponseType = 1
	InteractionResponseTypeChannelMessageWithSource InteractionResponseType = 4
)

type InteractionResponse struct {
	Type InteractionResponseType `json:"type"`
	Data InteractionResponseData `json:"data"`
}

type InteractionResponseData struct {
	Content string `json:"content"`
}

// CommandHandler is a function that handles a slash command.
type CommandHandler func(r *rand.Rand, data ApplicationCommandInteractionData) (InteractionResponseData, error)

var commandHandlers = map[string]CommandHandler{
	"roll": rollCommandHandler,
}

func rollCommandHandler(r *rand.Rand, data ApplicationCommandInteractionData) (InteractionResponseData, error) {
	rollCount := 1
	tableName := "2024"

	for _, opt := range data.Options {
		switch opt.Name {
		case "count":
			if val, ok := opt.Value.(float64); ok {
				rollCount = int(val)
			}
		case "table":
			if val, ok := opt.Value.(string); ok {
				tableName = val
			}
		}
	}

	if rollCount > 5 {
		rollCount = 5
	}
	if rollCount < 1 {
		rollCount = 1
	}

	var results []string
	for i := 0; i < rollCount; i++ {
		roll, effect, err := surge.Roll(r, tableName)
		if err != nil {
			return InteractionResponseData{}, err
		}
		results = append(results, fmt.Sprintf("(Roll: %d) %s", roll, effect))
	}

	return InteractionResponseData{Content: strings.Join(results, "\n\n")}, nil
}

// HandleInteraction handles the interaction from Discord.
func HandleInteraction(r *rand.Rand, interaction Interaction) (InteractionResponse, error) {
	switch interaction.Type {
	case InteractionTypePing:
		return InteractionResponse{Type: InteractionResponseTypePong}, nil
	case InteractionTypeApplicationCommand:
		if handler, ok := commandHandlers[interaction.Data.Name]; ok {
			data, err := handler(r, interaction.Data)
			if err != nil {
				return InteractionResponse{}, err
			}
			return InteractionResponse{
				Type: InteractionResponseTypeChannelMessageWithSource,
				Data: data,
			}, nil
		}
		return InteractionResponse{}, fmt.Errorf("unknown command: %s", interaction.Data.Name)
	}

	return InteractionResponse{}, nil
}
