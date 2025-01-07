package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type RequestCtx struct {
	Update *tgbotapi.Update
	bot    *tgbotapi.BotAPI
}

func (ctx *RequestCtx) Send(text string) error {
	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, text)
	_, err := ctx.bot.Send(msg)
	return err
}

func (ctx *RequestCtx) ShowMarkup(text string, keyboard tgbotapi.ReplyKeyboardMarkup) error {
	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboard
	_, err := ctx.bot.Send(msg)
	return err
}

func (ctx *RequestCtx) CloseMarkup(text string) error {
	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := ctx.bot.Send(msg)
	return err
}
