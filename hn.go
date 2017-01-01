package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	hnAPIURL  = "https://hacker-news.firebaseio.com/v0/"
	hnItemURL = "https://news.ycombinator.com/item?id="
)

type hnItem struct {
	Id    int
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

func fetchItem(id int, c chan hnItem) {
	var item hnItem

	body := fetchURL(fmt.Sprintf(hnAPIURL+"item/%d.json", id))
	if err := json.Unmarshal(body, &item); err != nil {
		panic(fmt.Sprintf("error while unmarshal item %s: %s",
			id, err))
	}

	c <- item
}

func main() {
	nrListItems := flag.Int("nrItems", 10, "Number of items to print out")
	verbose := flag.Bool("verbose", false, "Print out verbose information")

	flag.Parse()

	var topStories []int

	body := fetchURL(hnAPIURL + "topstories.json")

	if err := json.Unmarshal(body, &topStories); err != nil {
		panic(fmt.Sprintf("error while unmarshal topstories: %s",
			err))
	}

	var chans = make([]chan hnItem, *nrListItems)
	for idx, id := range topStories[:*nrListItems] {
		chans[idx] = make(chan hnItem)
		go fetchItem(id, chans[idx])
	}

	output := ""
	for i := 0; i < *nrListItems; i++ {
		item := <-chans[i]
		output += fmt.Sprintf("[%d] %s (%d)\n", i+1, item.Title, item.Score)
		if *verbose {
			output += fmt.Sprintf("[%s]\n[%s]\n\n", item.Url,
				fmt.Sprintf(hnItemURL+"%d", item.Id))
		}
	}
	fmt.Printf("%s", output)
}
