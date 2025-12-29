package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			b.handleMessage(update.Message)
		}
	}
}

func (b *Bot) handleMessage(msg *tgbotapi.Message) {
	switch msg.Text {
	case "/start":
		resp := GetMsg(b.conf.Language, "start")
		b.api.Send(tgbotapi.NewMessage(msg.Chat.ID, resp))
	}
}
