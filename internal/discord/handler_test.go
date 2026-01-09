package discord

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestHandleInteraction_Ping(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	interaction := Interaction{
		Type: InteractionTypePing,
	}

	response, err := HandleInteraction(rng, interaction)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if response.Type != InteractionResponseTypePong {
		t.Errorf("Expected response type %v, but got %v", InteractionResponseTypePong, response.Type)
	}
}

func TestHandleInteraction_RollCommand_Default(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	interaction := Interaction{
		Type: InteractionTypeApplicationCommand,
		Data: ApplicationCommandInteractionData{
			Name: "roll",
		},
	}

	response, err := HandleInteraction(rng, interaction)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if response.Type != InteractionResponseTypeChannelMessageWithSource {
		t.Errorf("Expected response type %v, but got %v", InteractionResponseTypeChannelMessageWithSource, response.Type)
	}

	if response.Data.Content == "" {
		t.Error("Expected response data to have content, but it was empty")
	}
	if strings.Contains(response.Data.Content, "\n\n") {
		t.Error("Expected a single roll, but found multiple")
	}
}

func TestHandleInteraction_RollCommand_CountTwo(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	interaction := Interaction{
		Type: InteractionTypeApplicationCommand,
		Data: ApplicationCommandInteractionData{
			Name: "roll",
			Options: []ApplicationCommandInteractionDataOption{
				{
					Name:  "count",
					Type:  4, // INTEGER
					Value: float64(2),
				},
			},
		},
	}

	response, err := HandleInteraction(rng, interaction)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if response.Type != InteractionResponseTypeChannelMessageWithSource {
		t.Errorf("Expected response type %v, but got %v", InteractionResponseTypeChannelMessageWithSource, response.Type)
	}

	if !strings.Contains(response.Data.Content, "\n\n") {
		t.Error("Expected two rolls separated by a double newline, but didn't find one")
	}
}
