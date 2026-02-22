package main

import (
	"os"
	"phraser_bot/internal/app"
)

func main() {
	app.StartUpdateHandler(
		"./config.yaml",
		os.Getenv("TOKEN"),
	)
}
