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

func ensureValidCat(cat *string) {
	for _, c := range []string{"top", "new", "best"} {
		if *cat == c {
			return
		}
	}
	fmt.Fprintf(os.Stderr, "Wrong category %s.\n\n", *cat)
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	nrItems := flag.Int("nrItems", 10, "Number of items to print out")
	cat := flag.String("category", "top",
		"Category of items to show.  It should be (top|new|best)")
	showOrigURL := flag.Bool("showOrigURL", false,
		"Show URL for the story")
	showCommentURL := flag.Bool("showCommentURL", false,
		"Show URL for HN comments")
	showURLs := flag.Bool("showURLs", false,
		"Show URLs for the story and HN comments")

	flag.Parse()

	ensureValidCat(cat)

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
	if *showURLs {
		output += fmt.Sprintf("# %d %s stories\n\n", *nrItems, *cat)
		*showOrigURL = true
		*showCommentURL = true
	}
	for i := 0; i < *nrItems; i++ {
		item := <-chans[i]
		output += fmt.Sprintf("[%d] %s (%d)\n",
			i+1, item.Title, item.Score)
		if *showOrigURL {
			output += fmt.Sprintf("[%s]\n", item.Url)
		}
		if *showCommentURL {
			output += fmt.Sprintf("[%s]\n",
				fmt.Sprintf(hnItemURL+"%d", item.Id))
		}
		if *showOrigURL || *showCommentURL {
			output += "\n"
		}
	}
	fmt.Printf("%s", output)
}
