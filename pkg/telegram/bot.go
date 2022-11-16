package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rcomanne/telegram-bot/pkg/configuration"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func Start(config configuration.Config) *Bot {
	log.Println("Starting bot")
	botAPI, err := tgbotapi.NewBotAPI(config.Telegram.BotToken)
	if err != nil {
		log.Fatalf("Failed to start bot: %s\n", err)
	}

	log.Printf("Started bot %s\n", botAPI.Self.UserName)
	return &Bot{
		bot: botAPI,
	}
}

func (tb *Bot) SendMessage(chatId int64, message string) bool {
	log.Printf("Sending message to chat %d\n", chatId)
	msg := tgbotapi.NewMessage(chatId, message)

	_, err := tb.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return true
	}
}
