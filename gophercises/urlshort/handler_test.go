package urlshort

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lassekv/gogames/gophercises/urlshort/dynamo"
)

func TestPutDynamoDB(t *testing.T) {
	mux := http.NewServeMux()
	svc, ok := dynamo.CreateClient()
	if !ok {
		t.Errorf("Unable to create dynamo client!")
	}
	putHandler := DynamoDBPutHandler(*svc, mux)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("PUT", "/test?url=http://hejsa.dk", nil)
	if err != nil {
		t.Fatal(err)
	}
	putHandler.ServeHTTP(w, r)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Error("The record was not")
	}
}
