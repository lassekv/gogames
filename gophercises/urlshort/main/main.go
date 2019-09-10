package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lassekv/gogames/gophercises/urlshort"
	"github.com/lassekv/gogames/gophercises/urlshort/dynamo"
)

func main() {
	mux := http.NewServeMux()

	svc, ok := dynamo.CreateClient()
	if !ok {
		log.Fatal("Unable to create dynamo client")
	}
	dynamoListAllHandler := urlshort.DynamoDBListAllHandler(*svc, mux)
	dynamoPutHandler := urlshort.DynamoDBPutHandler(*svc, dynamoListAllHandler)
	dynamoGetHandler := urlshort.DynamoDBGetHandler(*svc, dynamoPutHandler)

	fmt.Println("Starting the server on :8080")
	err := http.ListenAndServe(":8080", dynamoGetHandler)
	if err != nil {
		log.Fatalf("error %v", err)
	}
}
