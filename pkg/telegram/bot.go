package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	b.handlerUpdates(updates)

	return nil
}

func (b *Bot) handlerUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			if err := b.handlerCallbackQuery(update.CallbackQuery); err != nil {
				log.Fatal(err)
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		if err := b.handlerMessage(update.Message); err != nil {
			log.Fatal(err)
		}
	}
}
