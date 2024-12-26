package tgbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Message struct {
	msg *tgbotapi.Message
}

func NewMessage(msg *tgbotapi.Message) *Message {
	return &Message{msg: msg}
}

func (m *Message) Text() string {
	return m.msg.Text
}
