package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message represents a message
type Message struct {
	*tgbotapi.Update
	bot     *tgbotapi.BotAPI
	builder *Builder
}

// NewMessage creates a new message
func NewMessage(bot *tgbotapi.BotAPI, builder *Builder, update *tgbotapi.Update) *Message {
	return &Message{
		Update:  update,
		bot:     bot,
		builder: builder,
	}
}

// Text returns the text of the message
func (m *Message) Text() string {
	if m.Message != nil {
		return m.Message.Text
	}

	if m.CallbackQuery != nil {
		return m.CallbackQuery.Data
	}

	return ""
}

// ChatID returns the chat ID
func (m *Message) ChatID() int64 {
	if m.Message != nil {
		return m.Message.Chat.ID
	}

	if m.CallbackQuery != nil {
		return m.CallbackQuery.Message.Chat.ID
	}

	return -1
}

// SendText sends a text message
func (m *Message) SendText(text string) error {
	msg := tgbotapi.NewMessage(m.ChatID(), text)
	_, err := m.bot.Send(msg)
	return err
}
