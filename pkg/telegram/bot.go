package telegram

import (
	"errors"
	"log"

	"github.com/GrandMaster5000/go-pocket-sdk"
	"github.com/GrandMaster5000/tg-bot-pocket/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	pocketClient    *pocket.Client
	tokenRepository repository.TokenRepository
	redirectURL     string
}

func NewBot(
	bot *tgbotapi.BotAPI,
	pocketClient *pocket.Client,
	tr repository.TokenRepository,
	redirectURL string,
) *Bot {
	return &Bot{
		bot:             bot,
		pocketClient:    pocketClient,
		tokenRepository: tr,
		redirectURL:     redirectURL,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChanel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}
		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChanel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)
	if updates == nil {
		return nil, errors.New("updates chan return nil")
	}

	return updates, nil
}
