package tgui

import "context"

// Builder is a ui/router builder
type Builder struct {
	menu MenuItem
}

// Build creates a new  ui/router builder
func Build(ctx context.Context, menu MenuItem) *Builder {
	return &Builder{menu: menu}
}

// FindByQuery finds a menu item by query
func (b *Builder) FindByQuery(query MenuItemID) (*MenuItem, error) {
	res := b.menu.FindByQuery(query)
	if res == nil {
		return nil, ErrNotFound
	}
	return res, nil
}
