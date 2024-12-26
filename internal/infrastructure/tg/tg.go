package tg

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	ui  *Builder
}

func NewBot(token string, isDebug bool, ui *Builder) *Bot {
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

func (b *Bot) GetBot() *tgbotapi.BotAPI { return b.bot }

func (b *Bot) Run(ctx context.Context) error {
	if err := b.handleUpdates(ctx); err != nil {
		slog.Error("handle updates", slog.String("error", err.Error()))
	}

	return nil
}

// handleUpdates is a long-polling method for check update for bot
func (b *Bot) handleUpdates(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if ctx.Err() != nil {
			slog.Info("tg_updates close context updates", slog.String("error", ctx.Err().Error()))
			return nil
		}

		var (
			chatID   int64
			text     string
			username string
		)

		if update.Message != nil || update.CallbackQuery != nil {
			if update.Message != nil {
				chatID = update.Message.Chat.ID
				text = update.Message.Text
				username = update.Message.From.UserName
			} else {
				chatID = update.CallbackQuery.Message.Chat.ID
				text = update.CallbackQuery.Data
				username = update.CallbackQuery.From.UserName
			}
		} else {
			continue
		}

		slog.Info("handle message",
			slog.Int64("chat_id", chatID),
			slog.String("username", username),
			slog.String("text", text))

		menu, err := b.ui.UserMenuFindByQuery(text)
		if err != nil {
			slog.Error("UserMenuFindByQuery - handle menu call",
				slog.String("text", text),
				slog.String("error", err.Error()))
			continue
		}

		if menu.OnClick != nil {
			tMsg := NewMessage(b.bot, b.ui, &update)
			if err := menu.OnClick(ctx, tMsg); err != nil {
				slog.Error("OnClick - handle menu call",
					slog.String("text", text),
					slog.String("error", err.Error()))
			}
			continue
		}

		if err := b.makeAnswer(chatID, menu); err != nil {
			slog.Error("makeAnswer - handle menu call",
				slog.String("text", text),
				slog.String("error", err.Error()))
			continue
		}
	}
	return nil
}

// makeAnswer sends a message to the user
func (b *Bot) makeAnswer(chatID int64, curMenu *MenuItem) error {
	var msg tgbotapi.MessageConfig

	if curMenu.CheckRedirect() {
		newMenu, err := b.ui.UserMenuFindByQuery(curMenu.RedirectTo)
		if err != nil {
			return fmt.Errorf("UserMenuFindByQuery - %w", err)
		}

		msg = tgbotapi.NewMessage(chatID, curMenu.Message)

		curMenu = newMenu
	} else {
		msg = tgbotapi.NewMessage(chatID, curMenu.Message)
	}

	if curMenu.Inline {
		msg.ReplyMarkup = curMenu.InlineKeyboard()
	} else {
		msg.ReplyMarkup = curMenu.ReplyKeyboard()
	}

	if msg.Text == "" {
		msg.Text = "Выберите пункт меню: "
	}

	if _, err := b.bot.Send(msg); err != nil {
		return fmt.Errorf("bot send menu: %w", err)
	}
	return nil
}
