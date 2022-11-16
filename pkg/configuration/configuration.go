package configuration

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

// Config Application configuration options
type Config struct {
	Reddit struct {
		AppId  string `yaml:"app-id" env:"APP_ID" env-description:"App id used for authenticating to Reddit"`
		Secret string `yaml:"secret" env:"SECRET" env-description:"Secret used for authenticating to Reddit"`
	} `yaml:"reddit" env-prefix:"REDDIT_"`
	Telegram struct {
		BotToken string `yaml:"bot-token" env:"BOT_TOKEN" env-description:"Token used for authenticating the Telegram bot"`
		ChatId   string `yaml:"chat-id" env:"CHAT_ID" env-description:"Chat ID to send the messages to"`
	} `yaml:"telegram" env-prefix:"TELEGRAM_"`
}

// Args Available args for the app
type Args struct {
	ConfigPath string
}

func ProcessArgs(cfg interface{}) Args {
	var a Args

	f := flag.NewFlagSet("Telegram Bot", 1)
	f.StringVar(&a.ConfigPath, "c", "config.yaml", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	if err := f.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	return a
}
