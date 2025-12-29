package db

import "time"

type ChannelSettings struct {
	ChatID        int64    `json:"chat_id"`
	Currencies    []string `json:"currencies"`
	UseSell       bool     `json:"use_sell"`
	UseBuy        bool     `json:"use_buy"`
	IntervalMin   int      `json:"interval_min"`
	LastPosted    map[string]string `json:"last_posted"`
	NavasanAPIKey string   `json:"navasan_api_key"`
	UseNavasan    bool     `json:"use_navasan"`
	TemplateID    string   `json:"template_id"`
	DisabledUntil time.Time `json:"disabled_until"`
}
