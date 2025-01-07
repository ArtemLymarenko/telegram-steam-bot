package v1Bot

import (
	"bot/internal/infrastructure/telegram"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetRouter() *telegram.Router {
	router := telegram.NewRouter()

	router.AddHandler(telegram.HandlerTypeCmd, "open", func(ctx *telegram.RequestCtx) error {
		numericKeyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("/close"),
			),
		)

		return ctx.ShowMarkup("Added", numericKeyboard)
	})

	router.AddHandler(telegram.HandlerTypeCmd, "close", func(ctx *telegram.RequestCtx) error {
		return ctx.CloseMarkup("Closed")
	})

	router.AddHandler(telegram.HandlerTypeText, "hello", func(ctx *telegram.RequestCtx) error {
		return ctx.Send(fmt.Sprintf("Hello %s!", ctx.Update.Message.From.UserName))
	})

	return router
}
