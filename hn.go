package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	hnAPIURL  = "https://hacker-news.firebaseio.com/v0/"
	hnItemURL = "https://news.ycombinator.com/item?id="
)

var (
	nrItems = flag.Int("nrItems", 10, "Number of items to print out")
	cat     = flag.String("category", "top",
		"Category of items to show.  It should be (top|new|best)")
	showOrigURL = flag.Bool("showOrigURL", false,
		"Show URL for the story")
	showCommentURL = flag.Bool("showCommentURL", false,
		"Show URL for HN comments")
	showURLs = flag.Bool("showURLs", false,
		"Show URLs for the story and HN comments")
	showTitle = flag.Bool("showTitle", false,
		"Show Title")
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

	flag.Parse()

	ensureValidCat(cat)

	var storyIDs []int

	body := fetchURL(hnAPIURL + *cat + "stories.json")
	if err := json.Unmarshal(body, &storyIDs); err != nil {
		panic(fmt.Sprintf("error while unmarshal topstories: %s",
			err))
	}

	if *nrItems > len(storyIDs) {
		fmt.Fprintf(os.Stderr, "Too many items required: %d > %d.\n",
			*nrItems, len(storyIDs))
		os.Exit(1)
	}

	var chans = make([]chan hnItem, *nrItems)
	for idx, id := range storyIDs[:*nrItems] {
		chans[idx] = make(chan hnItem)
		go fetchItem(id, chans[idx])
	}

	output := ""
	if *showTitle {
		title := fmt.Sprintf("%s %d stories", *cat, *nrItems)
		output += title + "\n"
		title = strings.Repeat("=", len(title))
		output += title + "\n\n"
	}

	if *showURLs {
		*showOrigURL = true
		*showCommentURL = true
	}

	for i := 0; i < *nrItems; i++ {
		item := <-chans[i]
		output += fmt.Sprintf("[%d] %s (%d)\n",
			i+1, item.Title, item.Score)
		if *showOrigURL {
			output += fmt.Sprintf("(%s)\n", item.Url)
		}
		if *showCommentURL {
			output += fmt.Sprintf("(%s)\n",
				fmt.Sprintf(hnItemURL+"%d", item.Id))
		}
		if *showOrigURL || *showCommentURL {
			output += "\n"
		}
	}
	fmt.Printf("%s", output)
}
