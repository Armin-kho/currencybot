package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (b *Bot) handleConfig(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	text := msg.Text

	// Admin only check
	if !b.isAdmin(msg.From.ID) {
		b.sendMessage(chatID, GetMsg(b.conf.Language, "not_admin"))
		return
	}

	if text == "/config" {
		b.sendMessage(chatID, GetMsg(b.conf.Language, "config_menu"))
		return
	}

	// parse sub‑commands
	if strings.HasPrefix(text, "/setcurrencies") {
		parts := strings.Split(text, " ")
		if len(parts) > 1 {
			currencies := parts[1:]
			settings, _ := b.loadOrNewSettings(chatID)
			settings.Currencies = currencies
			b.store.SaveChannelSettings(settings)
			b.sendMessage(chatID, fmt.Sprintf(GetMsg(b.conf.Language, "set_currencies_ok"), currencies))
		}
		return
	}

	if strings.HasPrefix(text, "/setinterval") {
		parts := strings.Split(text, " ")
		if len(parts) == 2 {
			if val, err := strconv.Atoi(parts[1]); err == nil {
				settings, _ := b.loadOrNewSettings(chatID)
				settings.IntervalMin = val
				b.store.SaveChannelSettings(settings)
				b.sendMessage(chatID, fmt.Sprintf(GetMsg(b.conf.Language, "set_interval_ok"), val))
			}
		}
		return
	}

	b.sendMessage(chatID, GetMsg(b.conf.Language, "unknown_config"))
}
