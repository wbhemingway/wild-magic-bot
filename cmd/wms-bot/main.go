package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/wbhemingway/wild-magic-bot/internal/config"
	"github.com/wbhemingway/wild-magic-bot/internal/discord"
)

var discordPublicKey ed25519.PublicKey

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Could not load configuration: %v. Please run `wms-register --reconfigure` or setup your environment.", err)
	}

	if cfg.PublicKey == "" {
		log.Fatal("DISCORD_PUBLIC_KEY is not set in the configuration.")
	}

	decodedKey, err := hex.DecodeString(cfg.PublicKey)
	if err != nil {
		log.Fatalf("Error decoding public key: %v", err)
	}
	discordPublicKey = ed25519.PublicKey(decodedKey)

	http.HandleFunc("/api/interactions", HandleDiscordWebHook)
	log.Printf("Listening on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}

func HandleDiscordWebHook(w http.ResponseWriter, r *http.Request) {
	signature := r.Header.Get("X-Signature-Ed25519")
	timestamp := r.Header.Get("X-Signature-Timestamp")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !discord.VerifySignature(signature, timestamp, body, discordPublicKey) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
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
