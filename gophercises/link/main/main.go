package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/lassekv/gogames/gophercises/link"
)

func main() {
	paths := os.Args[1:]
	for _, path := range paths {
		var content io.Reader
		if strings.HasPrefix(path, "https") || strings.HasPrefix(path, "http") {
			resp, err := http.Get(path)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			content = resp.Body

		} else {
			resp, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer resp.Close()
			content = resp
		}

		links, err := link.ParseHTML(content)
		if err != nil {
			panic(err)
		}
		for _, l := range links {
			fmt.Printf("Extracted %s with text %s\n", l.Href, l.Text)
		}
	}
}
