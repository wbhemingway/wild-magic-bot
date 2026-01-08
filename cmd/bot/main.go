package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/wbhemingway/wild-magic-bot/internal/discord"
)

var discordPublicKey ed25519.PublicKey

func main() {
	key := os.Getenv("DISCORD_PUBLIC_KEY")
	if key == "" {
		log.Fatal("DISCORD_PUBLIC_KEY environment variable not set")
	}

	decodedKey, err := hex.DecodeString(key)
	if err != nil {
		log.Fatalf("Error decoding public key: %v", err)
	}
	discordPublicKey = ed25519.PublicKey(decodedKey)

	http.HandleFunc("/api/interactions", HandleDiscordWebHook)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleDiscordWebHook(w http.ResponseWriter, r *http.Request) {
	if !discord.VerifySignature(r, discordPublicKey) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var interaction discord.Interaction
	if err := json.Unmarshal(body, &interaction); err != nil {
		log.Printf("Error unmarshalling request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusInternalServerError)
		return
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	response, err := discord.HandleInteraction(rng, interaction)
	if err != nil {
		log.Printf("Error handling interaction: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sendJSON(w, response)
}

func sendJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
