package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Main menu buttons
func MainMenuButtons(lang string) []tgbotapi.InlineKeyboardButton {
	return []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData(GetMsg(lang, "btn_configure"), "conf"),
		tgbotapi.NewInlineKeyboardButtonData(GetMsg(lang, "btn_stats"), "stats"),
		tgbotapi.NewInlineKeyboardButtonData(GetMsg(lang, "btn_clear"), "clear"),
	}
}

// Language selection
func LangButtons() []tgbotapi.InlineKeyboardButton {
	return []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("English", "lang_en"),
		tgbotapi.NewInlineKeyboardButtonData("فارسی", "lang_fa"),
	}
}
