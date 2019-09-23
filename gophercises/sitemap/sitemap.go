package sitemap

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/lassekv/gogames/gophercises/link"
)

func createURL(path string, host string) *url.URL {
	curURL, err := url.Parse(path)
	if err != nil {
		fmt.Printf("Unable to parse URL %v\n", path)
		return nil
	}
	if curURL.Host == "" {
		curURL.Host = host
		curURL.Scheme = "https"
	}
	return curURL
}

func getRefs(curURL url.URL) []*url.URL {
	d, err := http.Get(curURL.String())
	if err != nil {
		fmt.Printf("Unable to read URL %v. Returns error %v.\n", curURL, err)
		return nil
	}
	defer d.Body.Close()
	links, err := link.ParseHTML(d.Body)
	if err != nil {
		fmt.Printf("Unable to parse links of URL %v. %v\n", curURL, err)
		return nil
	}
	var urls = make([]*url.URL, 0, len(links))
	for _, l := range links {
		tmp := createURL(l.Href, curURL.Host)
		if tmp != nil && tmp.Host == curURL.Host {
			urls = append(urls, tmp)
		}
	}
	return urls
}

// BuildSitemap returns all URLs within a site
func BuildSitemap(entryURL string, host string) []*url.URL {
	initialURL := createURL(entryURL, host)
	urls := getRefs(*initialURL)
	// Make a queue of stuff to visit

	visitedURLs := make(map[*url.URL]int)
	visitedURLs[initialURL] = 1

	return urls
}
