package tgbot

// Builder is a ui/router builder
type Builder struct {
	userMenu  *MenuItem
	adminMenu *MenuItem
}

// NewBuilder creates a new  ui/router builder
func NewBuilder(userMenu, adminMenu *MenuItem) *Builder {
	return &Builder{
		userMenu:  userMenu,
		adminMenu: adminMenu,
	}
}

// UserMenuFindByQuery finds a menu item by title or ID
func (b *Builder) UserMenuFindByQuery(query string) (*MenuItem, error) {
	return b.userMenu.FindByQuery(query)
}

// UserMenuFindByID finds a menu item by title
func (b *Builder) UserMenuFindByID(id string) (*MenuItem, error) {
	res := b.userMenu.FindByID(id)
	if res == nil {
		return nil, ErrNotFound
	}
	return res, nil
}

// UserMenuFindByMsg finds a menu item by title
func (b *Builder) UserMenuFindByMsg(text string) (*MenuItem, error) {
	res := b.userMenu.FindByMsg(text)
	if res == nil {
		return nil, ErrNotFound
	}
	return res, nil
}

// AdminMenuFindByID finds a adminMenu item by query
func (b *Builder) AdminMenuFindByID(text string) (*MenuItem, error) {
	res := b.adminMenu.FindByMsg(text)
	if res == nil {
		return nil, ErrNotFound
	}
	return res, nil
}
