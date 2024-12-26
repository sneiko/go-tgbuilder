package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message represents a message
type Message struct {
	*tgbotapi.Message
	bot     *tgbotapi.BotAPI
	builder *Builder
}

// NewMessage creates a new message
func NewMessage(bot *tgbotapi.BotAPI, builder *Builder, msg *tgbotapi.Message) *Message {
	return &Message{
		Message: msg,
		bot:     bot,
		builder: builder,
	}
}

// Text returns the text of the message
func (m *Message) Text() string {
	return m.Message.Text
}

// ChatID returns the chat ID
func (m *Message) ChatID() int64 { return m.Message.Chat.ID }

// SendText sends a text message
func (m *Message) SendText(text string) error {
	msg := tgbotapi.NewMessage(m.ChatID(), text)
	_, err := m.bot.Send(msg)
	return err
}

// SendMenuItem sends a menu item
func (m *Message) SendMenuItem(id string) error {
	menu, err := m.builder.UserMenuFindByID(id)
	if err != nil {
		menu, err = m.builder.AdminMenuFindByID(id)
		if err != nil {
			return err
		}
	}
	msg := tgbotapi.NewMessage(m.ChatID(), menu.Message)

	if menu.Inline {
		msg.ReplyMarkup = menu.InlineKeyboard()
	} else {
		msg.ReplyMarkup = menu.ReplyKeyboard()
	}

	if msg.Text == "" {
		msg.Text = "Выберите пункт меню: "
	}

	_, err = m.bot.Send(msg)
	return err
}
