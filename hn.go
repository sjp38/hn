package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	HNAPIURL  = "https://hacker-news.firebaseio.com/v0/"
	HNItemURL = "https://news.ycombinator.com/item?id="
)

type hnItem struct {
	Url   string
	Score int
	Title string
}

func fetchURL(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("error while get %s: %s", url, err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(fmt.Sprintf("failed to read body: %s", err))
	}

	return body
}

func fetchItem(id int) hnItem {
	var item hnItem

	body := fetchURL(fmt.Sprintf(HNAPIURL+"item/%d.json", id))
	if err := json.Unmarshal(body, &item); err != nil {
		panic(fmt.Sprintf("error while unmarshal item %s: %s",
			id, err))
	}

	return item
}

func main() {
	var bestStories []int

	body := fetchURL(HNAPIURL + "beststories.json")

	if err := json.Unmarshal(body, &bestStories); err != nil {
		panic(fmt.Sprintf("error while unmarshal beststories: %s",
			err))
	}

	for idx, id := range bestStories {
		item := fetchItem(id)
		fmt.Printf("[%d] %s (%d)\n[%s]\n[%s]\n\n",
			idx, item.Title, item.Score, item.Url,
			fmt.Sprintf(HNItemURL+"%d", id))

		if idx >= 9 {
			break
		}
	}
}
