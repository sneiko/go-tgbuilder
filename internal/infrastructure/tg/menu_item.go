package tg

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// MenuItemOnClick is a call func for a userMenu item
type MenuItemOnClick func(ctx context.Context, msg *Message) error

// MenuItemID is a unique identifier for a userMenu item
type MenuItemID string

// MenuItem represents a userMenu item
type MenuItem struct {
	ID           MenuItemID
	Row          int
	Title        string
	Message      string
	Inline       bool
	RedirectTo   string // redirect to another menu by Message field
	OnClick      MenuItemOnClick
	ChildrenRows []MenuItem
}

// CheckRedirect checks if a menu item need to be redirected
func (m *MenuItem) CheckRedirect() bool { return m.RedirectTo != "" }

// FindByQuery find by title or ID
func (m *MenuItem) FindByQuery(query string) (*MenuItem, error) {
	menu := m.FindByID(query)
	if menu != nil {
		return menu, nil
	}

	menu = m.FindByMsg(query)
	if menu != nil {
		return menu, nil
	}

	return nil, ErrNotFound
}

// FindByID finds a menu item by id
func (m *MenuItem) FindByID(id string) *MenuItem {
	if m.ID == MenuItemID(id) {
		return m
	}

	for _, child := range m.ChildrenRows {
		if child.ID == MenuItemID(id) {
			return &child
		}
		if child.FindByMsg(id) != nil {
			return child.FindByMsg(id)
		}
	}
	return nil
}

// FindByMsg finds a menu item by id
func (m *MenuItem) FindByMsg(text string) *MenuItem {
	if m.Title == text {
		return m
	}

	for _, child := range m.ChildrenRows {
		if child.Title == text {
			return &child
		}
		if child.FindByMsg(text) != nil {
			return child.FindByMsg(text)
		}
	}
	return nil
}

// InlineKeyboard builds inline a keyboard from a menu item
func (m *MenuItem) InlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	for idx := range m.ChildrenRows {
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

	for idx := range m.ChildrenRows {
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
