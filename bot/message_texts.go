package bot

var BotTexts = map[string]map[string]string{
	"en": {
		"not_admin":         "You are not an admin.",
		"config_menu":       "Configuration commands:\n/setcurrencies USD EUR GBP\n/setinterval 5",
		"set_currencies_ok": "Currencies updated: %v",
		"set_interval_ok":   "Interval updated to %d minutes",
		"unknown_config":    "Unknown configuration command.",
	},
	"fa": {
		"not_admin":         "شما مدیر نیستید.",
		"config_menu":       "فرمان‌های تنظیمات:\n/setcurrencies USD EUR GBP\n/setinterval 5",
		"set_currencies_ok": "ارزها به‌روزرسانی شدند: %v",
		"set_interval_ok":   "فاصله زمانی تنظیم شد به %d دقیقه",
		"unknown_config":    "فرمان پیکربندی نامشخص است.",
	},
}
