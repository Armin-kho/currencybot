package bot

import "github.com/Armin-kho/currencybot/db"

func (b *Bot) loadOrNewSettings(chatID int64) (*db.ChannelSettings, error) {
	settings, err := b.store.LoadChannelSettings(chatID)
	if err != nil {
		settings = &db.ChannelSettings{
			ChatID:        chatID,
			Currencies:    []string{},
			UseSell:       true,
			UseBuy:        false,
			IntervalMin:   5,
			LastPosted:    make(map[string]string),
			NavasanAPIKey: "",
			UseNavasan:    false,
			TemplateID:    "default",
		}
	}
	return settings, nil
}
