package app

import (
	"context"
	"log/slog"

	"tg_star_miner/internal/infrastructure/tg"
	"tg_star_miner/internal/infrastructure/tg/tgbot"
	"tg_star_miner/internal/infrastructure/tg/tgui"
)

const (
	TgBotToken = "7774772315:AAEujlPzxB0c3PV2im4W0_oyXAazHd5MXRg"
)

func Run(ctx context.Context) error {

	ui := tgui.NewBuilder(
		&tgui.MenuItem{
			ID:      "start",
			Row:     0,
			Title:   "Главное меню",
			Message: "Выберите пункт меню: ",
			Inline:  true,
			OnClick: nil,
			ChildrenRows: []tgui.MenuItem{
				{
					ID:           "information",
					Row:          0,
					Title:        "Информация 🚀",
					OnClick:      nil,
					ChildrenRows: nil,
				}, {
					ID:           "action",
					Row:          0,
					Title:        "События 🚀",
					RedirectTo:   "start",
					OnClick:      nil,
					ChildrenRows: nil,
				}, {
					ID:    "qa",
					Row:   1,
					Title: "Q/A 🚀",
					OnClick: func(ctx context.Context, msg *tgbot.Message) error {
						slog.Info("handle menu command",
							slog.String("command", msg.Text()))
						return nil
					},
					ChildrenRows: nil,
				},
			},
		}, nil)

	return tg.NewBot(TgBotToken, false, ui).Run(ctx)
}
