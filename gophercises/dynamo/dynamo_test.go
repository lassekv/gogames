package dynamo

import "testing"

func TestDynamoCon(t *testing.T) {
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
