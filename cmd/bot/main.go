package main

import (
	"log"

	"github.com/GrandMaster5000/go-pocket-sdk"
	"github.com/GrandMaster5000/tg-bot-pocket/pkg/repository/boltdb"
	"github.com/GrandMaster5000/tg-bot-pocket/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost/")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
