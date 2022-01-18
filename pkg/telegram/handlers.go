package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var accessKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сохранить", "save"),
		tgbotapi.NewInlineKeyboardButtonData("Отмена", "not save"),
	),
)

func (b *Bot) handlerMessage(message *tgbotapi.Message) error {
	if message.Text != "" {
		log.Printf("[%s] %s", message.From.UserName, message.Text)
	}
	if message.Photo != nil {
		log.Printf("[%s] %s", message.From.UserName, "photo")

		photoArray := *message.Photo
		indexBigImage := len(photoArray) - 1
		photo := photoArray[indexBigImage]

		msg := tgbotapi.NewPhotoShare(message.Chat.ID, photo.FileID)
		msg.ReplyMarkup = accessKeyboard
		msg.ReplyToMessageID = message.MessageID

		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) handlerCallbackQuery(callbackQuery *tgbotapi.CallbackQuery) error {
	log.Printf("[%s] %s", callbackQuery.Message.From.UserName, callbackQuery.Data)

	if callbackQuery.Data == "save" {
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Изображение сохранено ✅")
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	} else {
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не сохранено!")
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}
