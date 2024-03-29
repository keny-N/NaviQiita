package qiita

import (
	"encoding/json"
	"net/http"
)

type QiitaItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func FetchQiitaItems() ([]QiitaItem, error) {
	resp, err := http.Get("https://qiita.com/api/v2/items?page=1&per_page=5")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var items []QiitaItem
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, err
	}

	return items, nil
}
