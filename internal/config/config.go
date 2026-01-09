package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the applications.
type Config struct {
	AppID     string
	BotToken  string
	PublicKey string
	Port      string
}

// Load attempts to load configuration from various sources.
func Load() (Config, error) {
	// 1. Try loading from project-local .env file (for developers)
	_ = godotenv.Load()

	// 2. Check shell environment variables
	if os.Getenv("DISCORD_APP_ID") != "" && os.Getenv("DISCORD_BOT_TOKEN") != "" {
		return loadFromEnv(), nil
	}

	// 3. Check for user-specific config file
	configDir, err := os.UserConfigDir()
	if err == nil {
		configFile := filepath.Join(configDir, "wms", "config.env")
		if _, err := os.Stat(configFile); err == nil {
			// Clear existing env vars before loading to ensure pristine load
			os.Unsetenv("DISCORD_APP_ID")
			os.Unsetenv("DISCORD_BOT_TOKEN")
			os.Unsetenv("DISCORD_PUBLIC_KEY")
			os.Unsetenv("PORT")

			if err := godotenv.Load(configFile); err == nil {
				return loadFromEnv(), nil
			}
		}
	}

	return Config{}, fmt.Errorf("could not find credentials in environment or config file")
}

// InteractiveSetup prompts the user for credentials and optionally saves them.
func InteractiveSetup() (Config, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your DISCORD_APP_ID: ")
	appID, _ := reader.ReadString('\n')
	appID = strings.TrimSpace(appID)

	fmt.Print("Please enter your DISCORD_BOT_TOKEN: ")
	botToken, _ := reader.ReadString('\n')
	botToken = strings.TrimSpace(botToken)

	fmt.Print("Please enter your DISCORD_PUBLIC_KEY: ")
	publicKey, _ := reader.ReadString('\n')
	publicKey = strings.TrimSpace(publicKey)

	if appID == "" || botToken == "" || publicKey == "" {
		return Config{}, fmt.Errorf("all fields must be provided")
	}

	cfg := Config{
		AppID:     appID,
		BotToken:  botToken,
		PublicKey: publicKey,
		Port:      "8080", // Default port
	}

	fmt.Print("\nWould you like to save this configuration for next time? (y/n): ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))

	if answer == "y" || answer == "yes" {
		saveConfiguration(cfg)
	}

	return cfg, nil
}

// saveConfiguration saves the given config to the user's config directory.
func saveConfiguration(cfg Config) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("Warning: Could not find user config directory. %v\n", err)
		return
	}

	wmsConfigDir := filepath.Join(configDir, "wms")
	if err := os.MkdirAll(wmsConfigDir, 0755); err != nil {
		fmt.Printf("Warning: Could not create config directory at %s. %v\n", wmsConfigDir, err)
		return
	}

	configFile := filepath.Join(wmsConfigDir, "config.env")
	fileContent := fmt.Sprintf(
		"DISCORD_APP_ID=%s\nDISCORD_BOT_TOKEN=%s\nDISCORD_PUBLIC_KEY=%s\nPORT=%s\n",
		cfg.AppID, cfg.BotToken, cfg.PublicKey, cfg.Port,
	)

	if err := os.WriteFile(configFile, []byte(fileContent), 0600); err != nil {
		fmt.Printf("Warning: Failed to save config file at %s. %v\n", configFile, err)
	} else {
		fmt.Printf("Successfully saved configuration to %s\n", configFile)
	}
}

// loadFromEnv loads config from the current environment variables.
func loadFromEnv() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return Config{
		AppID:     os.Getenv("DISCORD_APP_ID"),
		BotToken:  os.Getenv("DISCORD_BOT_TOKEN"),
		PublicKey: os.Getenv("DISCORD_PUBLIC_KEY"),
		Port:      port,
	}
}
