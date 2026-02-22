package app

import (
	"fmt"
	"os"
	"phraser_bot/internal/services"

	"github.com/stretchr/testify/assert/yaml"
)

func StartUpdateHandler(pathToConfig string, token string) {
	handler := services.NewUpdateHandler(
		parseConfig(pathToConfig),
		token,
	)
	handler.HandleAllUpdates()
}

func parseConfig(pathToConfig string) *services.UpdateHandlerConfig {
	file, err := os.ReadFile(pathToConfig)
	if err != nil {
		fmt.Printf("Error while parsing configs: %s", err.Error())
		os.Exit(1)
	}
	var config *services.UpdateHandlerConfig = &services.UpdateHandlerConfig{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		fmt.Printf("Error while marshalling configs: %s", err.Error())
		os.Exit(1)
	}
	return config
}
