package tg

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg_star_miner/internal/infrastructure/tg/tgui"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	ui  *tgui.Builder
}

func NewBot(token string, isDebug bool, ui *tgui.Builder) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = isDebug

	slog.Info("Authorized on account", slog.String("username", bot.Self.UserName))

	return &Bot{
		bot: bot,
		ui:  ui,
	}
}

func (b *Bot) GetBot() *tgbotapi.BotAPI {
	return b.bot
}

func (b *Bot) Run(ctx context.Context) error {
	if err := b.handleUpdates(ctx); err != nil {
		slog.Error("handle updates", slog.String("error", err.Error()))
	}

	return nil
}

func (b *Bot) handleUpdates(ctx context.Context) error {
	updateCfg := tgbotapi.NewUpdate(0)
	updateCfg.Timeout = 60

	for {
		select {
		case <-ctx.Done():
			slog.Info("handle updates", slog.String("reason", ctx.Err().Error()))
			return nil

		case update := <-b.bot.GetUpdatesChan(updateCfg):
			if update.Message == nil {
				continue
			}

			if update.Message.IsCommand() {
				cmd := update.Message.Command()

				slog.Info("handle menu command",
					slog.Int64("chat_id", update.Message.Chat.ID),
					slog.String("username", update.Message.From.UserName),
					slog.String("text", update.Message.Text),
					slog.String("command", update.Message.Command()))

				menu, err := b.ui.UserMenuFindByID(cmd)
				if err != nil {
					slog.Error("handle menu command",
						slog.String("command", update.Message.Command()),
						slog.String("error", err.Error()))
					continue
				}

				if err := b.makeAnswer(update.Message.Chat.ID, menu); err != nil {
					slog.Error("handle menu command",
						slog.String("command", update.Message.Command()),
						slog.String("error", err.Error()))
					continue
				}
			}
		}
	}
}

func (b *Bot) makeAnswer(chatID int64, menu *tgui.MenuItem) error {
	var (
		msg tgbotapi.MessageConfig
	)

	if menu.CheckRedirect() {
		menu, err := b.ui.UserMenuFindByID(string(menu.RedirectTo))
		if err != nil {
			return fmt.Errorf("redirect menu: %w", err)
		}

		msg = tgbotapi.NewMessage(chatID, menu.Message)
	} else {
		msg = tgbotapi.NewMessage(chatID, menu.Message)
	}

	if menu.Inline {
		msg.ReplyMarkup = menu.InlineKeyboard()
	} else {
		msg.ReplyMarkup = menu.ReplyKeyboard()
	}

	if _, err := b.bot.Send(msg); err != nil {
		return fmt.Errorf("send menu: %w", err)
	}
	return nil
}
