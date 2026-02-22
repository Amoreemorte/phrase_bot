package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type PhotoSize struct {
	// FileID identifier for this file, which can be used to download or reuse
	// the file
	FileID string `json:"file_id"`
}

func PhotoFromTgPhotoSize(tgPhoto *tgbotapi.PhotoSize) *PhotoSize {
	return &PhotoSize{
		FileID: tgPhoto.FileID,
	}
}

type Update struct {
	*Message
	*CallbackQuery
}

func UpdateFromTgUpdate(tgUpdate *tgbotapi.Update) *Update {
	update := &Update{}
	if tgUpdate.Message != nil {
		update.Message = MessageFromTgMessage(tgUpdate.Message)
	}
	if tgUpdate.CallbackQuery != nil {
		update.CallbackQuery = CallbackQueryFromTgCallbackQuery(tgUpdate.CallbackQuery)
	}
	return update
}

type Message struct {
	SenderId int64
	Text     string
	Photo    []*PhotoSize

	isCommand bool
}

func MessageFromTgMessage(tgMsg *tgbotapi.Message) *Message {
	msg := &Message{
		SenderId: tgMsg.From.ID,
		Text:     tgMsg.Text,
	}
	photos := make([]*PhotoSize, 0, len(tgMsg.Photo))
	for _, photo := range tgMsg.Photo {
		photos = append(photos, PhotoFromTgPhotoSize(&photo))
	}
	msg.Photo = photos
	msg.SetCommand(tgMsg.IsCommand())
	return msg
}

func (m *Message) IsCommand() bool {
	return m.isCommand
}

func (m *Message) SetCommand(isCommand bool) {
	m.isCommand = isCommand
}

type CallbackQuery struct {
	SenderId int64
	Data     string
	UserName string
}

func CallbackQueryFromTgCallbackQuery(tgCallback *tgbotapi.CallbackQuery) *CallbackQuery {
	callback := &CallbackQuery{
		SenderId: tgCallback.From.ID,
		Data:     tgCallback.Data,
		UserName: tgCallback.From.UserName,
	}
	return callback
}

type MessageConfig struct {
	ReceiverId int64
	Text       string
}

func TgMessageFromMessageConfig(cfg *MessageConfig) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(cfg.ReceiverId, cfg.Text)
	return &msg
}
