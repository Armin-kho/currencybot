package scheduler

import (
	"time"
	"github.com/Armin-kho/currencybot/bot"
	"github.com/Armin-kho/currencybot/db"
)

type Scheduler struct {
	conf *db.Config
	store *db.Store
	bot   *bot.Bot
}

func NewScheduler(conf *db.Config, store *db.Store, b *bot.Bot) *Scheduler {
	return &Scheduler{conf: conf, store: store, bot: b}
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		// scraper & post logic will come next
	}
}
