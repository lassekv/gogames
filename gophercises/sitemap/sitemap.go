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
	fmt.Printf("Retrieving data from %v\n", curURL)
	d, err := http.Get(curURL.String())
	if err != nil {
		// Figure out how to do correct error handling here.
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
func transformToMap(urls []*url.URL) map[url.URL]int {
	val := make(map[url.URL]int)
	for _, v := range urls {
		val[*v] = 1
	}
	return val
}

// BuildSitemap returns all URLs within a site
func BuildSitemap(entryURL string, host string) []url.URL {
	initialURL := createURL(entryURL, host)
	toVisit := transformToMap(getRefs(*initialURL))
	visitedURLs := make(map[url.URL]int)
	visitedURLs[*initialURL] = 1
	for len(toVisit) > 0 {
		for el := range toVisit {
			if _, ok := visitedURLs[el]; ok {
				delete(toVisit, el)
				continue
			}
			visitedURLs[el] = 1
			newUrls := getRefs(el)
			for _, v := range newUrls {
				_, exist1 := visitedURLs[*v]
				_, exist2 := toVisit[*v]
				if !(exist1 || exist2) {
					toVisit[el] = 1
				}
			}
			delete(toVisit, el)
			if len(toVisit) == 1 {
				for el := range toVisit {
					fmt.Println(el)
				}
			}
		}
	}
	result := make([]url.URL, 0, len(visitedURLs))
	for k := range visitedURLs {
		result = append(result, k)
	}
	return result
}
