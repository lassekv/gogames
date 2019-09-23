package main

import (
	"fmt"
	"os"

	"github.com/lassekv/gogames/gophercises/sitemap"
)

func main() {
	paths := os.Args[1:]
	for _, path := range paths {
		urls := sitemap.BuildSitemap(path, "www.google.com")
		for _, u := range urls {
			fmt.Printf("%v\n", u)
		}
	}
}
