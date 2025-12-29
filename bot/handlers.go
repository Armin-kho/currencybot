package bot

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Start command
func (b *Bot) handleStart(msg *tgbotapi.Message) {
	b.sendMessage(msg.Chat.ID, GetMsg(b.conf.Language, "start"))
	b.sendMenu(msg.Chat.ID)
}

// Send main menu
func (b *Bot) sendMenu(chatID int64) {
	menu := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(MainMenuButtons(b.conf.Language)...),
	)
	b.sendMessageWithKeyboard(chatID, GetMsg(b.conf.Language, "menu_text"), menu)
}

// Main update loop
func (b *Bot) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				b.handleStart(update.Message)
			default:
				b.sendMessage(update.Message.Chat.ID, GetMsg(b.conf.Language, "unknown"))
			}
		} else if update.CallbackQuery != nil {
			b.handleCallback(update.CallbackQuery)
		}
	}
}

// Callback handler
func (b *Bot) handleCallback(cb *tgbotapi.CallbackQuery) {
	data := cb.Data
	chatID := cb.Message.Chat.ID

	switch {
	case data == "conf":
		b.sendMessage(chatID, GetMsg(b.conf.Language, "conf_msg"))
	case data == "stats":
		b.sendMessage(chatID, b.getStats())
	case data == "clear":
		b.clearCache()
	case data == "lang_en":
		b.conf.Language = "en"
		b.sendMessage(chatID, "Language set to English")
	case data == "lang_fa":
		b.conf.Language = "fa"
		b.sendMessage(chatID, "زبان به فارسی تنظیم شد")
	default:
		b.sendMessage(chatID, "Unsupported action")
	}

	b.api.Request(tgbotapi.NewCallback(cb.ID, ""))
}
