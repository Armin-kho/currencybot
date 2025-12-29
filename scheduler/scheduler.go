package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/Armin-kho/currencybot/bot"
	"github.com/Armin-kho/currencybot/db"
	"github.com/Armin-kho/currencybot/scraper"
)

func (s *Scheduler) Start() {
	for {
		channels := s.store.AllChannelKeys()
		for _, key := range channels {
			settings, err := s.store.LoadChannelSettings(key)
			if err != nil {
				log.Println("Scheduler load settings:", err)
				continue
			}

			if time.Now().Before(settings.DisabledUntil) {
				continue
			}

			// Check interval alignment
			if !isAligned(settings.IntervalMin) {
				continue
			}

			// Fetch prices
			var priceData map[string]string
			if settings.UseNavasan && settings.NavasanAPIKey != "" {
				navData, err := scraper.FetchNavasan(settings.NavasanAPIKey)
				if err == nil {
					priceData = ParseNavasan(navData)
				}
			}
			if priceData == nil {
				priceData, err = scraper.ScrapeBonbast()
				if err != nil {
					s.bot.SendAdminAlert(fmt.Sprintf("Scrape failed for chan %d error: %v", settings.ChatID, err))
					continue
				}
			}

			// Prepare message
			msg := FormatMessage(settings, priceData, s.bot.Language)

			// Only send if changed
			if PricesChanged(settings.LastPosted, priceData, settings.Currencies) {
				if settings.UseSell && !settings.UseBuy {
					msg = OnlySell(msg)
				} else if settings.UseBuy && !settings.UseSell {
					msg = OnlyBuy(msg)
				}
				s.bot.PostUpdate(settings.ChatID, msg)
				settings.LastPosted = ExtractLast(settings.Currencies, priceData)
				s.store.SaveChannelSettings(settings)
			}
		}

		time.Sleep(time.Minute)
	}
}

// Helper: check interval alignment
func isAligned(interval int) bool {
	min := time.Now().Minute()
	return (min % interval) == 0
}
