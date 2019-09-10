package dynamo

import "testing"

func TestGetRecord(t *testing.T) {
	svc, ok := CreateClient()
	if !ok {
		t.Error("something went wrong!")
	}
	shortURL := "github"
	res, ok := GetRecord(svc, shortURL)
	if !ok {
		t.Error("something went wrong!")
	}
	expectedURL := "https://www.github.com/"
	if res.URL != expectedURL {
		t.Errorf("return '%v', but expected '%v'", res.URL, expectedURL)
	}
	if res.ShortURL != shortURL {
		t.Errorf("return '%v', but expected '%v'", res.ShortURL, shortURL)
	}
}

func TestPutRecord(t *testing.T) {
	svc, ok := CreateClient()
	if !ok {
		t.Error("something went wrong!")
	}
	shortURL := "wired"
	URL := "https://www.wired.com/"
	ok = PutRecord(svc, shortURL, URL)
	if !ok {
		t.Error("something went wrong inserting the record!")
	}

	res, ok := GetRecord(svc, shortURL)
	if res.URL != URL {
		t.Errorf("return '%v' but expected '%v'", res.URL, URL)
	}
}
