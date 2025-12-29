package bot

func (b *Bot) SendAdminAlert(msg string) {
	for _, admin := range b.conf.Admins {
		b.sendMessage(admin, "⚠️ Admin Alert: "+msg)
	}
}
