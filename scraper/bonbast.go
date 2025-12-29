package scraper

import (
	"errors"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const bonbastURL = "https://www.bonbast.com"

func ScrapeBonbast() (map[string]string, error) {
	resp, err := http.Get(bonbastURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch bonbast")
	}
	data := make(map[string]string)
	z := html.NewTokenizer(resp.Body)
	var lastKey string
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		token := z.Token()
		if token.Type == html.TextToken {
			txt := strings.TrimSpace(token.Data)
			if txt != "" {
				data[lastKey] = txt
			}
		}
	}
	return data, nil
}
