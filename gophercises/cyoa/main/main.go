package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/lassekv/gogames/gophercises/cyoa"
)

type KnownAdventures map[string]cyoa.Adventure

func writeTemplate(templateFile string, w io.Writer, o interface{}) error {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}
	err = t.Execute(w, o)
	if err != nil {
		return err
	}
	return nil
}

type StoryArcContent struct {
	Adventure string
	Story     cyoa.StoryArc
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func (ka KnownAdventures) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		err := writeTemplate("list_all.tmpl", w, ka)
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}
	urlParts := deleteEmpty(strings.Split(strings.TrimSpace(req.URL.Path), "/"))
	if len(urlParts) != 2 {
		w.WriteHeader(404)
		return
	}

	if adventure, ok := ka[urlParts[0]]; ok {
		if storyArc, ok := adventure[urlParts[1]]; ok {
			sa := StoryArcContent{Adventure: urlParts[0], Story: storyArc}
			err := writeTemplate("show_storyarc.tmpl", w, sa)
			if err != nil {
				w.WriteHeader(500)
			}
			return
		}
	}
	w.WriteHeader(404)
}

func main() {

	adv, err := cyoa.ReadAdventure("adventures/gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	ka := KnownAdventures{"gopher": adv}
	log.Fatal(http.ListenAndServe("localhost:8080", ka))
}
