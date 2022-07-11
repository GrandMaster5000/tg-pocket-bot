package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, replyUnknownError)
	switch err {
	case errInvalidURL:
		msg.Text = replyNotValidLink
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = replyYouNotAuthorized
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = replyFaildSaveLink
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
