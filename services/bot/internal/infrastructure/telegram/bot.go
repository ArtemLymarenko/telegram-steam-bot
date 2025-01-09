package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
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

func (b *Bot) getRequestArgs(text string) []string {
	msg := strings.Fields(text)
	if len(msg) == 1 {
		return nil
	}
	return msg[1:]
}

func (b *Bot) handleRequestAsync(update *tgbotapi.Update) error {
	ctx := &RequestCtx{
		Update: update,
		bot:    b.bot,
	}

	b.router.launchMiddlewares(ctx)

	msgText := update.Message.Text
	option := Option{
		Type:  HandlerTypeText,
		Route: msgText,
	}

	if update.Message.IsCommand() {
		option.Type = HandlerTypeCmd
		option.Route = update.Message.Command()

		if !b.router.HasHandler(option) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid handler route, command does not exists!")
			if _, err := b.bot.Send(msg); err != nil {
				log.Println(err)
			}
		}

		ctx.args = b.getRequestArgs(msgText)
	}

	handler, err := b.router.GetHandler(option)
	if err != nil {
		return err
	}

	return handler(ctx)
}

func (b *Bot) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		go func() {
			err := b.handleRequestAsync(&update)
			if err != nil {
				log.Printf("[%s] %s", update.Message.From.UserName, err.Error())
			}
		}()
	}
}
