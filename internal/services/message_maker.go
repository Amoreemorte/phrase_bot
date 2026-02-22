package services

import "phraser_bot/internal/models"

// commands
const (
	Start_command = "/start"
)

type MessageMaker struct {
	cfg *MessageMakerConfig
}

type MessageMakerConfig struct {
	Link string

	SendPhrase  string
	RightPhrase string
	WrongPhrase string
}

func NewMessageMaker(cfg *MessageMakerConfig) *MessageMaker {
	return &MessageMaker{cfg: cfg}
}

func (m *MessageMaker) GetSendPhraseMessage(receiverId int64) *models.MessageConfig {
	return &models.MessageConfig{
		ReceiverId: receiverId,
		Text:       m.cfg.SendPhrase,
	}
}

func (m *MessageMaker) GetRightPhraseMessage(receiverId int64) *models.MessageConfig {
	return &models.MessageConfig{
		ReceiverId: receiverId,
		Text:       m.cfg.RightPhrase + m.cfg.Link,
	}
}

func (m *MessageMaker) GetWrongPhraseMessage(receiverId int64) *models.MessageConfig {
	return &models.MessageConfig{
		ReceiverId: receiverId,
		Text:       m.cfg.WrongPhrase,
	}
}
