package db

func (s *Store) AllChannelKeys() []int64 {
	keys := []int64{}
	s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := string(item.Key())
			if strings.HasPrefix(key, "channel_") {
				idStr := strings.TrimPrefix(key, "channel_")
				id, _ := strconv.ParseInt(idStr, 10, 64)
				keys = append(keys, id)
			}
		}
		return nil
	})
	return keys
}
