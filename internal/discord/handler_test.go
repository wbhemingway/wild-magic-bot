package discord

import (
	"math/rand"
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

func TestHandleInteraction_RollCommand(t *testing.T) {
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
}
