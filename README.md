# Wild Magic Surge Bot

A Go-powered Discord bot and command-line tool for rolling on a Wild Magic Surge table. This project provides a lightweight, fast, and easy-to-use way to bring the chaos of wild magic to your game.

It is designed with a clean architecture, separating the core logic from the application entry points, and can be used either as a simple CLI or as a full-featured Discord bot.

## Features

- **Discord Bot (`wms-bot`):** A modern, webhook-based Discord bot that responds to a `/roll` slash command.
- **CLI (`wms`):** A simple, no-config command-line tool to get a surge effect directly in your terminal.
- **Easy Setup (`wms-register`):** An interactive, user-friendly utility to register bot commands and securely configure credentials.
- **Configurable:** Supports configuration via environment variables, a local `.env` file for development, or a shared config file for installed binaries.
- **Deployable:** Comes with a `Dockerfile` ready for easy deployment to container platforms like Google Cloud Run.

## Commands

### `/roll`
Rolls on the Wild Magic Surge table.

**Options:**
- `count` (Optional): The number of times to roll. Can be `1-5`. Defaults to `1`.
- `table` (Optional): The surge table to use. Can be `2024` or `2014`. Defaults to `2024`.

---

## 1. Discord Application Setup

Before you can run the bot, you need to create an application in Discord's developer portal.

**[Click here to go to the Discord Developer Portal](https://discord.com/developers/applications)**

1.  Click **"New Application"**.
2.  Give your application a name (e.g., "Wild Magic Bot").
3.  Navigate to the **"Bot"** tab on the left.
4.  From this page, you will need three pieces of information for the configuration steps below:
    *   **Application ID:** Found on the "General Information" page.
    *   **Public Key:** Found on the "General Information" page.
    *   **Bot Token:** Found on the "Bot" page. Click "Reset Token" to generate one. **Treat your token like a password and never share it.**

---

## 2. Using the CLI (`wms`)

This is the simplest way to use the project.

### Installation

> [!NOTE]
> This method requires [Go to be installed](https://go.dev/doc/install) on your system.

```bash
go install github.com/wbhemingway/wild-magic-bot/cmd/wms@latest
```

### Usage
The CLI uses flags to specify the table and the number of rolls.

```bash
# Roll once on the default (2024) table
wms

# Roll 3 times on the 2014 table
wms -count=3 -table=2014

# Get help and see all available flags
wms -h
```

---

## 3. Running the Discord Bot

There are two primary ways to run the bot, depending on your goal.

### For End-Users (Recommended)

This method uses `go install` and the interactive setup utility.

**1. Installation:**

> [!NOTE]
> This method requires [Go to be installed](https://go.dev/doc/install) on your system.

Install the bot server and the registration utility.
```bash
go install github.com/wbhemingway/wild-magic-bot/cmd/wms-bot@latest
go install github.com/wbhemingway/wild-magic-bot/cmd/wms-register@latest
```

**2. Configuration:**
Run the registration utility for the first time. It will guide you through an interactive setup.
```bash
wms-register
```
It will prompt you for your `Application ID`, `Public Key`, and `Bot Token`, and then ask for permission to save them to a configuration file in a safe location.

If you ever need to change your credentials, you can run the command with the `--reconfigure` flag:
```bash
wms-register --reconfigure
```

**3. Running the Bot:**
Once configured, you can start the bot server:
```bash
wms-bot
```
The server will start listening for interactions from Discord. You will need to deploy this to a public server for it to work. See the "Deployment" section below.

### For Developers / Contributors

This method is for those who want to contribute to the project's code.

**1. Clone the Repository:**
```bash
git clone https://github.com/wbhemingway/wild-magic-bot.git
cd wild-magic-bot
```

**2. Create a Configuration File:**
Create a file named `.env` in the root of the project. This file is ignored by git. Add your credentials to it:
```
# .env file
DISCORD_APP_ID="your_app_id_here"
DISCORD_BOT_TOKEN="your_bot_token_here"
DISCORD_PUBLIC_KEY="your_public_key_here"
PORT=8080
```
*You can use the `.env.example` file as a template.*

**3. Register Commands & Run the Bot:**
Use `go run` to execute the commands directly from the source code.
```bash
# Register the slash commands with Discord (only needs to be done once)
go run ./cmd/wms-register

# Run the bot server locally
go run ./cmd/wms-bot
```
When running locally, you will need a tool like `ngrok` to expose your local server to the internet and provide the public URL to Discord.

---

## 4. Deployment to Google Cloud Run

The `wms-bot` is a stateless web server, which is perfect for serverless platforms like Google Cloud Run. The included `Dockerfile` makes this process straightforward.

**Prerequisites:**
*   A Google Cloud Platform account.
*   The `gcloud` command-line tool installed and configured.
*   Docker installed and running on your machine.

**Steps:**

1.  **Enable APIs:** Make sure you have the Cloud Run and Artifact Registry APIs enabled for your GCP project.

2.  **Build the Container Image:**
    From the project's root directory, build the Docker image using Google Cloud Build. This will build your image and push it to Artifact Registry.
    ```bash
    # Replace [PROJECT_ID] with your GCP Project ID
    gcloud builds submit --tag gcr.io/[PROJECT_ID]/wms-bot
    ```

3.  **Deploy to Cloud Run:**
    Deploy the image you just pushed.
    ```bash
    gcloud run deploy wms-bot \
      --image gcr.io/[PROJECT_ID]/wms-bot \
      --platform managed \
      --region us-central1 \
      --allow-unauthenticated
    ```

4.  **Configure Secrets:**
    A Discord bot requires a public key. You must provide this to your Cloud Run service as an environment variable.
    *   Go to your `wms-bot` service in the Google Cloud Run console.
    *   Click "Edit & Deploy New Revision".
    *   Go to the "Variables & Secrets" tab.
    *   Add an environment variable named `DISCORD_PUBLIC_KEY` and set its value to your bot's public key.
    *   Click "Deploy".

5.  **Update Discord:**
    *   In the Cloud Run console, find the URL for your service. It will look like `https://wms-bot-[hash]-uc.a.run.app`.
    *   Go to your application in the Discord Developer Portal.
    *   In the "General Information" page, paste your service URL into the **"Interactions Endpoint URL"** field. Make sure to append `/api/interactions` to the end of the URL.
    *   Example: `https://wms-bot-....a.run.app/api/interactions`

Your bot is now deployed and ready to accept commands.

---
## A Note on Development

This project was built as an experiment in agent-driven development. The majority of the code, architecture, tests, and documentation were generated by an AI coding agent based on a series of high-level prompts and iterative conversation.
