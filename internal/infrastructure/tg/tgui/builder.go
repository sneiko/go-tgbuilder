package tgui

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

// UserMenuFindByID finds a userMenu item by query
func (b *Builder) UserMenuFindByID(id string) (*MenuItem, error) {
	res := b.userMenu.FindByID(id)
	if res == nil {
		return nil, ErrNotFound
	}
	return res, nil
}

// AdminMenuFindByID finds a adminMenu item by query
func (b *Builder) AdminMenuFindByID(id string) (*MenuItem, error) {
	res := b.adminMenu.FindByID(id)
	if res == nil {
		return nil, ErrNotFound
	}
	return res, nil
}
