# Phraser Bot
# Description

  **Phraser Bot** is a Telegram bot that validates user messages against predefined phrases. Upon receiving a correct phrase, the bot responds with a corresponding link; incorrect phrases trigger an error message.

  The bot implements a strictly utilitarian function: providing access to a link based on keyword verification.

# Features
- **Concurrent Update Processing**: Implements a worker pool pattern for handling Telegram updates, ensuring the bot remains responsive under high load by distributing requests across multiple goroutines

- **Modular Architecture**: Designed for easy adaptation to other Telegram Bot API-compatible platforms with minimal code changes

- **Configurable Phrase Matching**: YAML-based configuration for flexible phrase-response pairs management

- **Multistage build**: create light ~29.88 Mb image (alpine)

# Quick Start
## Prerequisites
Docker installed on your system

Telegram Bot Token (obtain from @BotFather)

## Installation
1. Clone the repository
```bash
git clone <repository-url>
cd phraser_bot
```
2. Configure the bot
```bash
# Copy the example configuration
cp config.yaml.example config.yaml

# Edit config.yaml with your desired phrases and responses
```

3. Build and run with Docker
```bash
# Build the Docker image
docker build --build-arg USE_EXAMPLE=true -t phraser_bot:latest .

# Run the container
docker run --rm -e TOKEN="YOUR_BOT_TOKEN" -e PHRASE="YOUR_PHRASE" -e LINK="YOUR_LINK" phraser_bot
```
## Environment Variables
|Variable	  |Description	                    |Required|
|-----------|---------------------------------|--------|
|TOKEN      |Telegram Bot API token	          |Yes     |
|PHRASE     |Phrase to check                  |No      |
|LINK       |Link to send                     |No      |
|USE_EXAMPLE|Config.yaml makes from example   |No      |

## Configuration File (config.yaml)
```yaml
# Num of workers
WorkersNum: 20
# If the bot is overloaded, the worker takes a delay of this duration
WaitingTime: 2s
# Size of updates chan buffer
UpdatesChanSize: 40
# Phrase to get link. Case-insensitive
Phrase: "test_phrase"
MessageMakerConfig:
  Link:        "https://t.me/test_link"
  SendPhrase:  "Я ожидаю твоего слова"
  RightPhrase: "Ты оказался прав, вот то, что я обещал "
  WrongPhrase: "я постараюсь забыть это"
```

# Project Structure
```text
phraser_bot/
├── cmd/
│   └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   └── models/
│       └── models.go
├── services/
│   ├── message_maker.go
│   ├── phraser.go
│   ├── phraser_test.go
│   └── update_handler.go
├── .dockerignore
├── .gitignore
├── config.yaml
├── config.yaml.example
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

# Development
Local Build (without Docker)
bash
## Build the binary
go build -o phraser_bot ./cmd

## Run Tests
``` bash
go test ./...
```

# License
MIT License

# Contributing
Contributions are welcome! Please submit a Pull Request.