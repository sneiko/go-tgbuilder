package tgui

import (
	"context"

	"tg_star_miner/internal/infrastructure/tg/tgbot"
)

// MenuItemOnClick is a call func for a menu item
type MenuItemOnClick func(ctx context.Context, msg *tgbot.Message) error

// MenuItemID is a unique identifier for a menu item
type MenuItemID string

// MenuItem represents a menu item
type MenuItem struct {
	ID       MenuItemID
	Title    string
	OnClick  MenuItemOnClick
	Children []MenuItem
}

// NewMenuItem creates a new menu item
func NewMenuItem(id MenuItemID, title string, onClick MenuItemOnClick, children []MenuItem) MenuItem {
	return MenuItem{
		ID:       id,
		Title:    title,
		OnClick:  onClick,
		Children: children,
	}
}

func (m *MenuItem) FindByQuery(query MenuItemID) *MenuItem {
	for _, child := range m.Children {
		if child.ID == query {
			return &child
		}

		if child.FindByQuery(query) != nil {
			return child.FindByQuery(query)
		}
	}
	return nil
}
