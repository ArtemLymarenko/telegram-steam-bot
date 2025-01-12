package handlers

import (
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/infrastructure/telegram"
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
	return ctx.CloseMarkup("Keyboard has been closed successfully!")
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
	//Add check games logic

	return ctx.Send("Games you provided don't have a discount now...\nTry again later.")
}

func (handlers *BotHandlers) InlineEchoQuery(ctx *telegram.RequestCtx) error {
	articles := []telegram.Article{
		{
			Url:        "https://shared.akamai.steamstatic.com/store_item_assets/steam/apps/1172470/8249072b14153cdb6bb65e2357f24d86daf7d965/capsule_184x69.jpg?t=1734541502",
			Title:      "art 1",
			Desc:       "desc art 1",
			TextToSend: ctx.Update.InlineQuery.Query,
		},
		{
			Url:        "https://shared.akamai.steamstatic.com/store_item_assets/steam/apps/1172470/8249072b14153cdb6bb65e2357f24d86daf7d965/capsule_184x69.jpg?t=1734541502",
			Title:      "art 2",
			Desc:       "desc art 2",
			TextToSend: ctx.Update.InlineQuery.Query,
		},
		{
			Url:        "https://shared.akamai.steamstatic.com/store_item_assets/steam/apps/1172470/8249072b14153cdb6bb65e2357f24d86daf7d965/capsule_184x69.jpg?t=1734541502",
			Title:      "art 2",
			Desc:       "desc art 2",
			TextToSend: ctx.Update.InlineQuery.Query,
		},
	}

	ctx.SendInlineQueryArticle(articles)
	return nil
}
