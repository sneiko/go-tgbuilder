package main

import (
	"context"
	"log/slog"

	"github.com/sneiko/go-tgbuilder/pkg/fsm"
	"github.com/sneiko/go-tgbuilder/pkg/fsm/fsmdb"
	"github.com/sneiko/go-tgbuilder/pkg/tg"
)

const (
	TgBotToken = "7774772315:AAEujlPzxB0c3PV2im4W0_oyXAazHd5MXRg"
)

var (
	AdminIds = []int64{7774772315}
)

func Run(ctx context.Context) error {
	fsm := fsm.New(fsmdb.NewInMem())

	ui := tg.NewBuilder(
		&tg.MenuItem{
			ID:      "/start",
			Title:   "Главное меню",
			Message: "Выберите пункт меню: ",
			Inline:  true,
			ChildrenRows: []tg.MenuItem{
				{
					ID:     "information",
					Row:    0,
					Title:  "Информация 🚀",
					Inline: true,
					ChildrenRows: []tg.MenuItem{
						{
							ID:         "bot",
							Title:      "Назад",
							RedirectTo: "/start",
						}, {
							ID:    "admin",
							Title: "admin test",
						}, {
							ID:    "test",
							Title: "test",
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

	return tg.NewBot(TgBotToken, false, ui, fsm).Run(ctx)
}
