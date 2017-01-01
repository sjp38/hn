package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	nrItems := flag.Int("nrItems", 10, "Number of items to print out")
	verbose := flag.Bool("verbose", false, "Print out verbose information")
	cat := flag.String("category", "top",
		"Category of items to show.  It should be top, new, or best")

	flag.Parse()

	validCat := false
	for _, c := range []string{"top", "new", "best"} {
		if *cat == c {
			validCat = true
			break
		}
	}
	if !validCat {
		fmt.Fprintf(os.Stderr, "Wrong category %s\n", *cat)
		flag.PrintDefaults()
		os.Exit(2)
	}

	var topStories []int

	body := fetchURL(hnAPIURL + *cat + "stories.json")

	if err := json.Unmarshal(body, &topStories); err != nil {
		panic(fmt.Sprintf("error while unmarshal topstories: %s",
			err))
	}

	var chans = make([]chan hnItem, *nrItems)
	for idx, id := range topStories[:*nrItems] {
		chans[idx] = make(chan hnItem)
		go fetchItem(id, chans[idx])
	}

	output := ""
	if *verbose {
		output += fmt.Sprintf("# %d %s stories\n\n", *nrItems, *cat)
	}
	for i := 0; i < *nrItems; i++ {
		item := <-chans[i]
		output += fmt.Sprintf("[%d] %s (%d)\n",
			i+1, item.Title, item.Score)
		if *verbose {
			output += fmt.Sprintf("[%s]\n[%s]\n\n", item.Url,
				fmt.Sprintf(hnItemURL+"%d", item.Id))
		}
	}
	fmt.Printf("%s", output)
}
