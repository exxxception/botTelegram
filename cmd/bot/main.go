package main

import (
	"github.com/exxxception/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	telegramBot := telegram.NewBot(bot)
	telegramBot.Start()
}
