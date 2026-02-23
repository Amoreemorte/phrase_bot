package app

import (
	"fmt"
	"os"
	"phraser_bot/internal/services"

	"github.com/stretchr/testify/assert/yaml"
)

func StartUpdateHandler(pathToConfig string, token string) {
	cfg := parseConfig(pathToConfig)
	if phrase := os.Getenv("PHRASE"); phrase != "" {
		cfg.Phrase = phrase
	}
	if link := os.Getenv("LINK"); link != "" {
		cfg.MsgsMakerConfig.Link = link
	}
	handler := services.NewUpdateHandler(
		cfg,
		token,
	)
	handler.HandleAllUpdates()
}

func parseConfig(pathToConfig string) *services.UpdateHandlerConfig {
	file, err := os.ReadFile(pathToConfig)
	if err != nil {
		fmt.Printf("Error while parsing configs: %s", err.Error())
		return services.GetDefaultUpdateHandlerConfig()
	}
	var config *services.UpdateHandlerConfig = &services.UpdateHandlerConfig{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		fmt.Printf("Error while marshalling configs: %s", err.Error())
		os.Exit(1)
	}
	return config
}
