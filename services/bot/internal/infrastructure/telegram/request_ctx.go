package telegram

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand/v2"
	"strconv"
)

type RequestCtx struct {
	Update          *tgbotapi.Update
	bot             *tgbotapi.BotAPI
	cmdText         string
	currentIndex    int
	abortProcessing bool
}

func (ctx *RequestCtx) GetCmdText() (string, error) {
	if ctx.cmdText != "" {
		return ctx.cmdText, nil
	}

	return "", errors.New("no args available")
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

func (ctx *RequestCtx) NextMiddleware() {
	if ctx.abortProcessing {
		return
	}

	ctx.currentIndex++
}

func (ctx *RequestCtx) AbortMiddleware() {
	ctx.abortProcessing = true
}

type Article struct {
	Url        string
	Title      string
	Desc       string
	TextToSend string
}

func (ctx *RequestCtx) mapToResultArticle(articles []Article) []interface{} {
	var outArticles []interface{}
	for _, article := range articles {
		outArticles = append(outArticles, tgbotapi.InlineQueryResultArticle{
			Type:        "article",
			ID:          strconv.Itoa(int(rand.Int64())),
			ThumbURL:    article.Url,
			Title:       article.Title,
			Description: article.Desc,
			InputMessageContent: tgbotapi.InputTextMessageContent{
				Text:      fmt.Sprintf("<b>%s</b>", article.TextToSend),
				ParseMode: "html",
			},
		})
	}

	return outArticles
}

func (ctx *RequestCtx) SendInlineQueryArticle(articles []Article) {
	text := ctx.Update.InlineQuery.Query
	if text == "" {
		text = "echo"
	}

	items := ctx.mapToResultArticle(articles)

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: ctx.Update.InlineQuery.ID,
		IsPersonal:    true,
		Results:       items,
	}

	if _, err := ctx.bot.Request(inlineConf); err != nil {
		log.Println(err)
	}
}
