package tgui

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg_star_miner/internal/infrastructure/tg/tgbot"
)

// MenuItemOnClick is a call func for a userMenu item
type MenuItemOnClick func(ctx context.Context, msg *tgbot.Message) error

// MenuItemID is a unique identifier for a userMenu item
type MenuItemID string

// MenuItem represents a userMenu item
type MenuItem struct {
	ID           MenuItemID
	Row          int
	Title        string
	Message      string
	Inline       bool
	RedirectTo   MenuItemID
	OnClick      MenuItemOnClick
	ChildrenRows []MenuItem
}

// CheckRedirect checks if a menu item need to be redirected
func (m *MenuItem) CheckRedirect() bool { return m.RedirectTo != "" }

// FindByID finds a menu item by id
func (m *MenuItem) FindByID(id string) *MenuItem {
	if m.ID == MenuItemID(id) {
		return m
	}

	for _, child := range m.ChildrenRows {
		if child.ID == MenuItemID(id) {
			return &child
		}
		if child.FindByID(id) != nil {
			return child.FindByID(id)
		}
	}
	return nil
}

// InlineKeyboard builds inline a keyboard from a menu item
func (m *MenuItem) InlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	for idx, _ := range m.ChildrenRows {
		var keyboardRow []tgbotapi.InlineKeyboardButton

		for _, item := range m.ChildrenRows {
			if item.Row != idx {
				continue
			}

			keyboardRow = append(keyboardRow, tgbotapi.NewInlineKeyboardButtonData(item.Title, string(item.ID)))
		}

		if len(keyboardRow) != 0 {
			keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(keyboardRow...))
		}
	}

	return tgbotapi.NewInlineKeyboardMarkup(keyboard...)
}

// ReplyKeyboard builds inline a keyboard from a menu item
func (m *MenuItem) ReplyKeyboard() tgbotapi.ReplyKeyboardMarkup {
	var keyboard [][]tgbotapi.KeyboardButton

	for idx, _ := range m.ChildrenRows {
		var keyboardRow []tgbotapi.KeyboardButton
		for _, item := range m.ChildrenRows {
			if item.Row != idx {
				continue
			}

			keyboardRow = append(keyboardRow, tgbotapi.NewKeyboardButton(item.Title))
		}

		if len(keyboardRow) != 0 {
			keyboard = append(keyboard, tgbotapi.NewKeyboardButtonRow(keyboardRow...))
		}
	}

	return tgbotapi.NewReplyKeyboard(keyboard...)
}
