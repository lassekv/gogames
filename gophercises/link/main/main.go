package main

import (
	"fmt"
	"github.com/lassekv/gogames/gophercises/link"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			f.Close()
			panic(err)
		}
		links, err := link.ParseHTML(f)
		f.Close()
		for _, l := range links {
			fmt.Println(l)
		}
	}
}
