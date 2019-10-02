package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/lassekv/gogames/gophercises/sitemap"
)

func main() {
	paths := os.Args[1:]
	for _, path := range paths {
		curURL, _ := url.Parse(path)
		urls := sitemap.BuildSitemap(path, curURL.Host)
		fmt.Printf("Returned %d urls\n", len(urls))
		for _, u := range urls {
			fmt.Printf("%v\n", u)
		}
	}
}
