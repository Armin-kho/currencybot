package bot

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/Armin-kho/currencybot/db"
)

type Bot struct {
	api  *tgbotapi.BotAPI
	conf *db.Config
	store *db.Store
}

func NewBot(conf *db.Config, store *db.Store) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		return nil, err
	}
	log.Printf("Authorized bot: %s", bot.Self.UserName)
	return &Bot{api: bot, conf: conf, store: store}, nil
}
