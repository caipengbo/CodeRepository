package internal

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

func ParseUrl(u Url, maxDepth int) ([]Url, error) {
	var urls []Url
	fmt.Printf("ParseUrl: %s \n", u.url)
	if u.depth < maxDepth {
		list, err := parse(u.url)
		if err != nil {
			return nil, err
		}
		depth := u.depth + 1
		for _, urlString := range list {
			urls = append(urls, Url{urlString, depth})
		}
	}
	return urls, nil
}

func parse(rawUrl string) ([]string, error) {
	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return nil, fmt.Errorf("invalid url: %s", err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", rawUrl, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get error: %s", err.Error())
	}
	defer resp.Body.Close()

	r, err := charset.NewReader(resp.Body, "UTF-8")

	if err != nil {
		fmt.Printf("Err: %s", err.Error())
		return nil, err
	}
	//all, err := ioutil.ReadAll(r)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Print(string(all))

	links := extract(r, rawUrl)
	return links, nil
}

// extract collect all absolute web url. It may contain duplicate links.
func extract(htmlReader io.Reader, baseUrl string) []string {
	// remove last '/' in base url
	if baseUrl[len(baseUrl)-1:] == "/" {
		baseUrl = baseUrl[:len(baseUrl)-1]
	}
	var urls []string
	page := html.NewTokenizer(htmlReader)
	// iterate all html tags
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return urls
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					url := trimAnchor(attr.Val)
					if path.IsAbs(url) {
						url = baseUrl + url
					}
					urls = append(urls, url)
				}
			}
		}
	}
}

// trimAnchor remove anchor from the link
func trimAnchor(link string) string {
	if strings.Contains(link, "#") {
		var index int
		for n, str := range link {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return link[:index]
	}
	return link
}
