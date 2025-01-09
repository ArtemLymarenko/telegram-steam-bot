package handlers

import (
	"bot/internal/infrastructure/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type BotHandlers struct{}

func NewBotHandlers() *BotHandlers {
	return &BotHandlers{}
}

func (handlers *BotHandlers) Open(ctx *telegram.RequestCtx) error {
	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/close"),
		),
	)

	return ctx.ShowMarkup("Keyboard has been opened successfully!", numericKeyboard)
}

func (handlers *BotHandlers) Close(ctx *telegram.RequestCtx) error {
	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/close"),
		),
	)

	return ctx.ShowMarkup("Keyboard has been closed successfully!", numericKeyboard)
}

func (handlers *BotHandlers) Help(ctx *telegram.RequestCtx) error {
	args, _ := ctx.GetArgs()
	return ctx.Send(strings.Join(args, " "))
}

func (handlers *BotHandlers) AddGame(ctx *telegram.RequestCtx) error {
	args, err := ctx.GetArgs()
	if err != nil {
		if len(args) != 1 {
			return ctx.Send("Should be exactly one argument!")
		}

		return ctx.Send("Failed to get arguments!")
	}

	url := args[0]
	if url == "" {
		return ctx.Send("Invalid url!")
	}

	//Add game logic

	return ctx.Send("Game has been added successfully!")
}

func (handlers *BotHandlers) CheckMyGames(ctx *telegram.RequestCtx) error {
	//Add game logic

	return ctx.Send("Games you provided don't have a discount now...\nTry again later.")
}
