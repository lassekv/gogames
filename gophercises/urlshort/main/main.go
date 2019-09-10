package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lassekv/gogames/gophercises/dynamo"

	"github.com/lassekv/gogames/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `pairs:
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	svc, ok := dynamo.CreateClient()
	if !ok {
		log.Fatal("Unable to create dynamo client")
	}
	dynamoHandler := urlshort.DynamoDBHandler(*svc, yamlHandler)

	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", dynamoHandler)
	if err != nil {
		log.Fatalf("error %v", err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "URL not found in maps")
}
