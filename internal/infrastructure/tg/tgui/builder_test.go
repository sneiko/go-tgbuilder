package tgui

import (
	"context"
	"testing"

	"tg_star_miner/internal/infrastructure/tg/tgbot"
)

func TestBuilder_FindByQuery(t *testing.T) {
	menu := MenuItem{
		ID:    "general",
		Title: "Главное меню",
		OnClick: func(ctx context.Context, msg *tgbot.Message) error {
			return nil
		},
		Children: []MenuItem{
			{
				ID:    "general-p1",
				Title: "Гланое меню - П1",
				OnClick: func(ctx context.Context, msg *tgbot.Message) error {
					return nil
				},
				Children: []MenuItem{
					{
						ID:    "general-p1-p1",
						Title: "П1 - П1",
						OnClick: func(ctx context.Context, msg *tgbot.Message) error {
							return nil
						},
					},
					{
						ID:    "general-p1-p2",
						Title: "П1 - П2",
						OnClick: func(ctx context.Context, msg *tgbot.Message) error {
							return nil
						},
					},
				},
			}, {
				ID:    "general-p2",
				Title: "Гланое меню - П2",
				OnClick: func(ctx context.Context, msg *tgbot.Message) error {
					return nil
				},
				Children: []MenuItem{
					{
						ID:    "general-p2-p1",
						Title: "П2 - П1",
						OnClick: func(ctx context.Context, msg *tgbot.Message) error {
							return nil
						},
					},
					{
						ID:    "general-p2-p1",
						Title: "П2 - П2",
						OnClick: func(ctx context.Context, msg *tgbot.Message) error {
							return nil
						},
					},
				},
			},
		},
	}

	builder := Build(context.Background(), menu)

	item, err := builder.FindByQuery("general-p1-p2")
	if err != nil {
		t.Error(err)
	}

	if item == nil {
		t.Error("menu item not found")
	}
}
