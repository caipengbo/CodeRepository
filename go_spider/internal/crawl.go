package internal

import (
	"fmt"
)

type Url struct {
	url   string
	depth int
}

func fetch(maxDepth int, unseenUrls chan Url, urlsList chan []Url) {
	for url := range unseenUrls {
		foundUrls, err := ParseUrl(url, maxDepth)
		if err != nil {
			fmt.Printf("skip url %s: %s", url.url, err.Error())
			continue
		}
		if len(foundUrls) > 0 {
			go func() {
				urlsList <- foundUrls
			}()
		}
	}
}

func crawl(urls []Url, config *Config) {
	urlsList := make(chan []Url) // lists of URLs
	unseenUrls := make(chan Url) // de-duplicated URLs
	// add the seed urls to urlsList
	go func() { urlsList <- urls }()

	// create worker routines to fetch unseen link
	for i := 0; i < config.Spider.ThreadCount; i++ {
		go fetch(config.Spider.MaxDepth, unseenUrls, urlsList)
	}

	// de-duplicates urlsList items and sends the unseen to the crawlers
	seen := make(map[string]bool)
	for list := range urlsList {
		for _, u := range list {
			if !seen[u.url] {
				seen[u.url] = true
				unseenUrls <- u
			}
		}
	}
}
