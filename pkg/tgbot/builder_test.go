package tgbot

import (
	"context"
	"testing"
)

func TestBuilder_FindByQuery(t *testing.T) {
	menu := MenuItem{
		ID:    "general",
		Title: "Главное меню",
		OnClick: func(ctx context.Context, msg *Message) error {
			return nil
		},
		ChildrenRows: []MenuItem{
			{
				ID:    "general-p1",
				Title: "Гланое меню - П1",
				OnClick: func(ctx context.Context, msg *Message) error {
					return nil
				},
				ChildrenRows: []MenuItem{
					{
						ID:    "general-p1-p1",
						Title: "П1 - П1",
						OnClick: func(ctx context.Context, msg *Message) error {
							return nil
						},
					},
					{
						ID:    "general-p1-p2",
						Title: "П1 - П2",
						OnClick: func(ctx context.Context, msg *Message) error {
							return nil
						},
					},
				},
			}, {
				ID:    "general-p2",
				Title: "Гланое меню - П2",
				OnClick: func(ctx context.Context, msg *Message) error {
					return nil
				},
				ChildrenRows: []MenuItem{
					{
						ID:    "general-p2-p1",
						Title: "П2 - П1",
						OnClick: func(ctx context.Context, msg *Message) error {
							return nil
						},
					},
					{
						ID:    "general-p2-p1",
						Title: "П2 - П2",
						OnClick: func(ctx context.Context, msg *Message) error {
							return nil
						},
					},
				},
			},
		},
	}

	builder := NewBuilder(&menu, nil)

	item, err := builder.UserMenuFindByID("general-p1-p2")
	if err != nil {
		t.Error(err)
	}

	if item == nil {
		t.Error("userMenu item not found")
	}
}
