package discord

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/wbhemingway/wild-magic-bot/internal/surge"
)

type InteractionType int

const (
	InteractionTypePing InteractionType = 1
	InteractionTypeApplicationCommand InteractionType = 2
)

type Interaction struct {
	Type InteractionType `json:"type"`
	Data ApplicationCommandInteractionData `json:"data"`
}

type ApplicationCommandInteractionData struct {
	Name string `json:"name"`
}

type InteractionResponseType int

const (
	InteractionResponseTypePong InteractionResponseType = 1
	InteractionResponseTypeChannelMessageWithSource InteractionResponseType = 4
)

type InteractionResponse struct {
	Type InteractionResponseType `json:"type"`
	Data InteractionResponseData `json:"data"`
}

type InteractionResponseData struct {
	Content string `json:"content"`
}

// HandleInteraction handles the interaction from Discord.
func HandleInteraction(r *rand.Rand, interaction Interaction) (InteractionResponse, error) {
	switch interaction.Type {
	case InteractionTypePing:
		return InteractionResponse{Type: InteractionResponseTypePong}, nil
	case InteractionTypeApplicationCommand:
		switch interaction.Data.Name {
		case "roll":
			effect, err := surge.Roll(r, "2024")
			if err != nil {
				return InteractionResponse{}, err
			}
			return InteractionResponse{
				Type: InteractionResponseTypeChannelMessageWithSource,
				Data: InteractionResponseData{
					Content: effect,
				},
			}, nil
		}
	}

	return InteractionResponse{}, nil
}
