package services

import (
	"fmt"
	"phraser_bot/internal/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdateHandler struct {
	msgMaker *MessageMaker
	phraser  *Phraser

	cfg *UpdateHandlerConfig
	bot *tgbotapi.BotAPI

	updatesChan chan (tgbotapi.Update)
}

type UpdateHandlerConfig struct {
	WorkersNum      int
	WaitingTime     time.Duration
	UpdatesChanSize int

	Phrase          string
	MsgsMakerConfig *MessageMakerConfig
}

func NewUpdateHandler(cfg *UpdateHandlerConfig, token string) *UpdateHandler {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return nil
	}
	return &UpdateHandler{
		cfg:         cfg,
		updatesChan: make(chan tgbotapi.Update, cfg.UpdatesChanSize),
		bot:         bot,
		msgMaker:    NewMessageMaker(cfg.MsgsMakerConfig),
		phraser:     NewPhraser(cfg.Phrase),
	}
}

func (h *UpdateHandler) HandleAllUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := h.bot.GetUpdatesChan(u)
	// start worket pool
	for i := 0; i < h.cfg.WorkersNum; i++ {
		go func() {
			for update := range h.updatesChan {
				h.handleUpdate(models.UpdateFromTgUpdate(&update))
			}
		}()
	}

	var waitingTime time.Duration
	for update := range updates {
		waitingTime = 0
		processed := false
		for !processed {
			select {
			case h.updatesChan <- update:
				processed = true
			default:
				time.Sleep(h.cfg.WaitingTime)
				waitingTime += h.cfg.WaitingTime
				continue
			}
		}
	}
}

func (h *UpdateHandler) handleUpdate(update *models.Update) {
	if update.Message == nil {
		return
	}

	msgs := []*models.MessageConfig{}
	text := update.Message.Text
	switch text {
	case Start_command:
		msgs = append(msgs, h.msgMaker.GetSendPhraseMessage(update.Message.SenderId))
	default:
		if h.phraser.IsPhrasesMatch(text) {
			msgs = append(msgs, h.msgMaker.GetRightPhraseMessage(update.Message.SenderId))
		} else {
			msgs = append(msgs, h.msgMaker.GetWrongPhraseMessage(update.Message.SenderId))
		}
	}

	for _, msg := range msgs {
		h.bot.Send(models.TgMessageFromMessageConfig(msg))
	}
}
