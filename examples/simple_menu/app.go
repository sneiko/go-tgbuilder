package main

import (
	"context"
	"log/slog"

	"github.com/sneiko/go-tgbuilder/pkg/tgbot"
	"github.com/sneiko/go-tgbuilder/pkg/tgfsm"
	"github.com/sneiko/go-tgbuilder/pkg/tgfsmdb"
)

const (
	TgBotToken = "7774772315:AAEujlPzxB0c3PV2im4W0_oyXAazHd5MXRg"
)

func Run(ctx context.Context) error {
	fsm := tgfsm.New(tgfsmdb.NewInMem())

	ui := tgbot.NewBuilder(
		&tgbot.MenuItem{
			ID:      "/start",
			Title:   "Главное меню",
			Message: "Выберите пункт меню: ",
			Inline:  true,
			ChildrenRows: []tgbot.MenuItem{
				{
					ID:     "information",
					Row:    0,
					Title:  "Информация 🚀",
					Inline: true,
					ChildrenRows: []tgbot.MenuItem{
						{
							ID:         "bot",
							Title:      "Назад",
							RedirectTo: "/start",
						}, {
							ID:    "test-1",
							Title: "test 1",
						}, {
							ID:    "test-2",
							Title: "test 2",
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
					OnClick: func(ctx context.Context, msg *tgbot.Message) error {
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

	return tgbot.NewBot(TgBotToken, false, ui, fsm).Run(ctx)
}
