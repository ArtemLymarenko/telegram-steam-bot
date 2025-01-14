package handlers

import (
	"context"
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/infrastructure/telegram"
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/interface/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type GamesClient interface {
	GetUserGames(context.Context, dto.GetUserGamesRequest) (*dto.GetUserGamesResponse, error)
	AddUserGame(context.Context, dto.AddUserGameRequest) (dto.AddUserGameResponse, error)
	SearchGamesByName(context.Context, dto.SearchGameRequest) (*dto.SearchGameResponse, error)
}

type BotHandlers struct {
	gamesClient GamesClient
}

func NewBotHandlers(gamesClient GamesClient) *BotHandlers {
	return &BotHandlers{
		gamesClient: gamesClient,
	}
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
	text, _ := ctx.GetCmdText()
	return ctx.Send(text)
}

func (handlers *BotHandlers) AddGame(ctx *telegram.RequestCtx) error {
	return ctx.Send("Game has been added successfully!")
}

func (handlers *BotHandlers) CheckMyGames(ctx *telegram.RequestCtx) error {
	return ctx.Send("Games you provided don't have a discount now...\nTry again later.")
}

func (handlers *BotHandlers) ChooseUserGameToAdd(ctx *telegram.RequestCtx) error {
	gameName := ctx.Update.InlineQuery.Query
	if gameName == "" {
		ctx.SendInlineQueryArticle([]telegram.Article{})
		return nil
	}

	games, err := handlers.gamesClient.SearchGamesByName(
		context.Background(),
		dto.SearchGameRequest{Name: gameName},
	)
	if err != nil {
		return ctx.Send("Nothing is found!")
	}

	articles := prepareArticlesForGameSearch(games.Games)
	ctx.SendInlineQueryArticle(articles)
	return nil
}
