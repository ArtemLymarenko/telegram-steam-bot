package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	router *Router
}

func NewBot(token string, debug bool, router *Router) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		bot:    bot,
		router: router,
	}
}

func (bot *Bot) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		option := Option{}
		if update.Message.IsCommand() {
			option.Type = HandlerTypeCmd
			option.Route = update.Message.Command()
		} else {
			option.Type = HandlerTypeText
			option.Route = update.Message.Text
		}

		handler, err := bot.router.GetHandler(option)
		if err != nil {
			log.Println(err)
			continue
		}

		ctx := &RequestCtx{
			Update: &update,
			bot:    bot.bot,
		}
		if err := handler(ctx); err != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, err.Error())
		}

	}
}
