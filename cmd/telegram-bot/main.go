package main

import (
	"github.com/go-co-op/gocron"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rcomanne/telegram-bot/pkg/configuration"
	"github.com/rcomanne/telegram-bot/pkg/reddit"
	"github.com/rcomanne/telegram-bot/pkg/telegram"
	"log"
	"time"
)

func main() {
	// Load in the configuration
	var config configuration.Config
	args := configuration.ProcessArgs(&config)
	if err := cleanenv.ReadConfig(args.ConfigPath, &config); err != nil {
		log.Fatal(err)
	}

	// Start the telegram bot
	bot := telegram.Start(config)

	// Start the CRON scheduler for sending the messages
	bs := gocron.NewScheduler(time.UTC)

	// Test loop
	if _, err := bs.Cron("*/1 * * * 1-5").Do(func() {
		log.Println("Test run")
		if _, err := reddit.GetPostFromSubreddit("bbq"); err != nil {
			log.Println(err)
		}
		bot.SendMessage(-1001389755191, "CRON!")
	}); err != nil {
		log.Printf("Failed sending coffee message %bs\n", err)
	}

	// Send message every morning for coffee at 8:00
	if _, err := bs.Cron("0 8 * * 1-5").Do(func() {
		log.Println("Sending morning coffee message")
		bot.SendMessage(-1001389755191, "CRON!")
	}); err != nil {
		log.Printf("Failed sending coffee message %bs\n", err)
	}

	// Send message every day for lunch at 11:30
	if _, err := bs.Cron("30 11 * * 1-5").Do(func() {
		log.Println("Sending lunch message")
		bot.SendMessage(-1001389755191, "CRON!")
	}); err != nil {
		log.Printf("Failed sending coffee message %bs\n", err)
	}

	bs.StartBlocking()
}
