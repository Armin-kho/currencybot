package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchNavasan(apiKey string) (map[string]interface{}, error) {
	url := fmt.Sprintf("http://api.navasan.tech/latest/?api_key=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&out)
	return out, err
}
