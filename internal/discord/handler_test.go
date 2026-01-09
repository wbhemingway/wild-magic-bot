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

	if !strings.HasPrefix(response.Data.Content, "(Roll: ") {
		t.Errorf("Expected content to start with '(Roll: ', but got %q", response.Data.Content)
	}
	if strings.Contains(response.Data.Content, "\n\n") {
		t.Error("Expected a single roll, but found multiple")
	}
}

func TestHandleInteraction_RollCommand_MultipleCount(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	interaction := Interaction{
		Type: InteractionTypeApplicationCommand,
		Data: ApplicationCommandInteractionData{
			Name: "roll",
			Options: []ApplicationCommandInteractionDataOption{
				{
					Name:  "count",
					Type:  4, // INTEGER
					Value: float64(5),
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

	if count := strings.Count(response.Data.Content, "\n\n"); count != 4 {
		t.Errorf("Expected 4 double newline separators for 5 rolls, but found %d", count)
	}
	if !strings.HasPrefix(response.Data.Content, "(Roll: ") {
		t.Errorf("Expected content to start with '(Roll: ', but got %q", response.Data.Content)
	}
}

func TestHandleInteraction_RollCommand_Table2014(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	interaction := Interaction{
		Type: InteractionTypeApplicationCommand,
		Data: ApplicationCommandInteractionData{
			Name: "roll",
			Options: []ApplicationCommandInteractionDataOption{
				{
					Name:  "table",
					Type:  3, // STRING
					Value: "2014",
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

	if !strings.HasPrefix(response.Data.Content, "(Roll: ") {
		t.Errorf("Expected content to start with '(Roll: ', but got %q", response.Data.Content)
	}
}
