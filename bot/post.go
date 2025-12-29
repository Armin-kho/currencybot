package bot

import "fmt"

// Post update (new message)
func (b *Bot) PostUpdate(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	b.api.Send(msg)
}

// Format message using templates
func FormatMessage(settings *db.ChannelSettings, prices map[string]string, lang string) string {
	out := ""
	for _, cur := range settings.Currencies {
		val := prices[cur]
		out += fmt.Sprintf("%s %s\n", cur, val)
	}
	return out
}
