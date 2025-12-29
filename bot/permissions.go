package bot

func (b *Bot) isAdmin(userID int64) bool {
	for _, admin := range b.conf.Admins {
		if admin == userID {
			return true
		}
	}
	return false
}
