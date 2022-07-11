package main

import (
	"log"

	"github.com/GrandMaster5000/go-pocket-sdk"
	"github.com/GrandMaster5000/tg-bot-pocket/pkg/config"
	"github.com/GrandMaster5000/tg-bot-pocket/pkg/repository"
	"github.com/GrandMaster5000/tg-bot-pocket/pkg/repository/boltdb"
	"github.com/GrandMaster5000/tg-bot-pocket/pkg/server"
	"github.com/GrandMaster5000/tg-bot-pocket/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)

	bot, err := tgbotapi.NewBotAPI("5398722889:AAG_4JX-CtNRuy6zmRnZhADHZ1FYwMJRW1g")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("102691-0a439896c314807226ed2a9")
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost")

	authorizationServer := server.NewAuthorizationServer(pocketClient, tokenRepository, "https://t.me/pocket_tolik_grandmaster_bot")

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}

}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessToken))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
