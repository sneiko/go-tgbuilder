package app

import (
	"context"
	"log/slog"

	"tg_star_miner/internal/infrastructure/tg"
)

const (
	TgBotToken = "7774772315:AAEujlPzxB0c3PV2im4W0_oyXAazHd5MXRg"
)

var (
	AdminIds = []int64{7774772315}
)

func Run(ctx context.Context) error {

	ui := tg.NewBuilder(
		&tg.MenuItem{
			ID:      "/start",
			Title:   "Главное меню",
			Message: "Выберите пункт меню: ",
			Inline:  true,
			ChildrenRows: []tg.MenuItem{
				{
					ID:    "information",
					Row:   0,
					Title: "Информация 🚀",
					ChildrenRows: []tg.MenuItem{
						{
							ID:         "bot",
							Title:      "Назад",
							RedirectTo: "/start",
						},
					},
				}, {
					ID:         "action",
					Row:        0,
					Title:      "События 🚀",
					RedirectTo: "/start",
				}, {
					ID:    "qa",
					Row:   1,
					Title: "Q/A 🚀",
					OnClick: func(ctx context.Context, msg *tg.Message) error {
						slog.Info("Q/A menu command",
							slog.String("command", msg.Text()))

						if err := msg.SendText("Q/A"); err != nil {
							return err
						}
						return nil
					},
				},
			},
		}, nil)

	return tg.NewBot(TgBotToken, false, ui).Run(ctx)
}
