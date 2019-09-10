package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lassekv/gogames/gophercises/urlshort"
	"github.com/lassekv/gogames/gophercises/urlshort/dynamo"
)

func main() {
	mux := defaultMux()

	svc, ok := dynamo.CreateClient()
	if !ok {
		log.Fatal("Unable to create dynamo client")
	}
	dynamoGetHandler := urlshort.DynamoDBGetHandler(*svc, mux)
	dynamoPutHandler := urlshort.DynamoDBPutHandler(*svc, dynamoGetHandler)

	fmt.Println("Starting the server on :8080")
	err := http.ListenAndServe(":8080", dynamoPutHandler)
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
	fmt.Fprintln(w, "Unknown URL")
}
