package client

import (
	"lkurnia/notifier/lib/request"
)

type Notifier struct {
	Telegram request.TelegramRequest
}

func (n *Notifier) SetupTelegram(apiKey string, chatId string) {
	n.Telegram = request.NewTelegramRequest(apiKey, chatId)
}
