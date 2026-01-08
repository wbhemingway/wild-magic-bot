package p

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"wild-magic-bot/internal/surge"
)

var DiscordPublicKey = os.Getenv("DISCORD_PUBLIC_KEY")

func HandleDiscordWebHook(w http.ResponseWriter, r *http.Request) {
	if !verifySignature(w, r){
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var interaction interaction
	if err := json.Unmarshal(bodyBytes, &interaction); err != nil {
		log.Printf("Error unmarshalling request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusInternalServerError)
		return
	}

	if interaction.Type == 1 {
		sendJSON(w, interactionResponse{Type: 1})
	}

	// Process the request body here
}

type Interaction struct {
	Type int `json:"type"`
	Data struct {
		Name string `json:"name"`
		Options []struct {
			Name string `json:"name"`
			Value string `json:"value"`
		} `json:"options"`
	} `json:"data"`
}

type InteractionResponse struct {
	Type int                      `json:"type"`
	Data *InteractionResponseData `json:"data,omitempty"`
}

type InteractionResponseData struct {
	Content string `json:"content"`
}
