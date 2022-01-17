package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var accessKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сохранить", "save"),
		tgbotapi.NewInlineKeyboardButtonData("Отмена", "back"),
	),
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
			log.Printf("[%s] %s", update.CallbackQuery.Message.From.UserName, "Press button")

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Ok, I remember")
			if _, err := b.bot.Send(msg); err != nil {
				log.Fatal(err)
			}
			continue
		}
		if update.Message.Text != "" {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			continue
		}
		if update.Message.Photo != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, "This is photo")

			photoArray := *update.Message.Photo
			indexBigImage := len(photoArray) - 1
			photo := photoArray[indexBigImage]

			msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, photo.FileID)
			msg.ReplyMarkup = accessKeyboard
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := b.bot.Send(msg); err != nil {
				log.Fatal(err)
			}
		}
	}
}
