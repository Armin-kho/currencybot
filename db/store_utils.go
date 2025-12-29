package db

import (
	"encoding/json"
	"fmt"
)

func (s *Store) SaveChannelSettings(settings *ChannelSettings) error {
	key := fmt.Sprintf("channel_%d", settings.ChatID)
	data, _ := json.Marshal(settings)
	return s.Set(key, data)
}

func (s *Store) LoadChannelSettings(chatID int64) (*ChannelSettings, error) {
	key := fmt.Sprintf("channel_%d", chatID)
	data, err := s.Get(key)
	if err != nil {
		return nil, err
	}
	var settings ChannelSettings
	json.Unmarshal(data, &settings)
	return &settings, nil
}
