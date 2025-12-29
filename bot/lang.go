package bot

var Messages = map[string]map[string]string{
	"en": {
		"start": "Welcome! Use the buttons to configure.",
	},
	"fa": {
		"start": "خوش آمدید! از دکمه‌ها برای تنظیم استفاده کنید.",
	},
}

func GetMsg(lang, key string) string {
	if m, ok := Messages[lang]; ok {
		return m[key]
	}
	return Messages["en"][key]
}
