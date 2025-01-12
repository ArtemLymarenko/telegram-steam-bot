package handlers

import (
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/infrastructure/telegram"
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/interface/bot/messages"
	"slices"
	"strings"
)

func (handlers *BotHandlers) HelloMiddleware(ctx *telegram.RequestCtx) error {
	words := []string{"привіт", "hola", "hi", "hello"}
	payload := strings.ToLower(ctx.Update.Message.Text)

	if slices.Contains(words, payload) {
		ctx.AbortMiddleware()
		return ctx.Send(messages.MessageGreeting.Format("%s\n%s", ctx.Update.Message.From.UserName))
	}

	ctx.NextMiddleware()
	return nil
}
