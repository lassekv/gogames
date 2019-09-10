package urlshort

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/lassekv/gogames/gophercises/urlshort/dynamo"
)

// DynamoDBGetHandler Resolves the map in a DynamoDB
func DynamoDBGetHandler(client dynamodb.DynamoDB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" && req.Method != "" {
			fallback.ServeHTTP(w, req)
			return
		}
		path := req.URL.Path
		if path == "/" || path == "/favicon.ico" {
			fallback.ServeHTTP(w, req)
			return
		}
		if strings.HasPrefix(path, "/") {
			path = path[1:]
		}
		if val, ok := dynamo.GetRecord(&client, path); ok && len(val.URL) > 0 {
			http.Redirect(w, req, val.URL, 301)
		} else {
			fallback.ServeHTTP(w, req)
			return
		}
	}
}

// DynamoDBPutHandler Adds a new url mapping on the form /shorturl?url=longurl
func DynamoDBPutHandler(client dynamodb.DynamoDB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "PUT" && req.Method != "POST" {
			fallback.ServeHTTP(w, req)
			return
		}
		shortURL := req.URL.Path
		query := req.URL.Query()
		if strings.HasPrefix(shortURL, "/") {
			shortURL = shortURL[1:]
		}
		if len(query["url"][0]) == 0 {
			fallback.ServeHTTP(w, req)
			return
		}
		dynamo.PutRecord(&client, shortURL, query["url"][0])
		http.Redirect(w, req, "http://localhost:8080/", 200)
	}
}

// DynamoDBListAllHandler Returns a json formated list of all records
func DynamoDBListAllHandler(client dynamodb.DynamoDB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		allRecords := dynamo.GetAllMappings(&client)
		jsonTXT, err := json.Marshal(allRecords)
		if err != nil {
			fallback.ServeHTTP(w, req)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonTXT)
	}
}
