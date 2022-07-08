package telegram

import (
	"context"
	"net/url"

	"github.com/GrandMaster5000/go-pocket-sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"

	replyStartTemplate     = "Привет! Чтобы сохранить ссылки в своём Pocket аккаунте, для начала тебе необходимо дать мне на это доступ. Для этого переходи по ссылке:\n%s"
	replyAlreadyAuthorized = "Ты уже авторизирован. Присылай ссылку, а я её сохраню."
	replyLinkSuccessSave   = "Ссылка успешно сохранена."
	replyNotValidLink      = "Это невалидная ссылка."
	replyYouNotAuthorized  = "Ты не авторизирован. Используй команду /start."
	replyFaildSaveLink     = "Увы, не удалось сохранить ссылку. Попробуй ещё раз позже."
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, replyLinkSuccessSave)

	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		msg.Text = replyNotValidLink
		_, err = b.bot.Send(msg)
		return err
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		msg.Text = replyYouNotAuthorized
		_, err = b.bot.Send(msg)
		return err
	}

	if err = b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		msg.Text = replyFaildSaveLink
		_, err = b.bot.Send(msg)
		return err
	}

	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProccess(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, replyAlreadyAuthorized)
	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды :(")
	_, err := b.bot.Send(msg)
	return err
}
