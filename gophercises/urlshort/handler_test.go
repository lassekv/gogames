package urlshort

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseYMLValidString(t *testing.T) {
	tString := `pairs:
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	r, err := parseYAML([]byte(tString))
	if err != nil {
		t.Error("err should be nil")
	}
	if len(r) != 2 {
		t.Error("The map should contain two entries")
	}
	for _, s := range []string{"/urlshort", "/urlshort-final"} {
		if _, ok := r[s]; !ok {
			t.Errorf("The string %s was not precent in the parsed data", s)
		}
	}
}

func TestParseYMLInvalidString(t *testing.T) {
	tString := `pairs:
- path2: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	r, err := parseYAML([]byte(tString))
	if r != nil {
		t.Errorf("No data should be returned for an invalid input. The following was returned %v", r)
	}
	if err == nil {
		t.Error("there should be an error")
	}
}

func TestMapHandlerWithExistingURL(t *testing.T) {
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mux := http.NewServeMux()
	mapHandler := MapHandler(pathsToUrls, mux)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/urlshort-godoc", nil)
	if err != nil {
		t.Fatal(err)
	}
	mapHandler.ServeHTTP(w, r)
	resp := w.Result()
	if resp.StatusCode != http.StatusMovedPermanently {
		t.Error("A redirect has not been recorded")
	}
	if resp.Header.Get("Location") != "https://godoc.org/github.com/gophercises/urlshort" {
		t.Errorf("The redirect is not to %v as expected", pathsToUrls["/urlshort-godoc"])
	}
}
