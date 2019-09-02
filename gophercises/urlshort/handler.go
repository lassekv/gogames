package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if val, ok := pathsToUrls[path]; ok {
			http.Redirect(w, req, val, 301)
		} else {
			fallback.ServeHTTP(w, req)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathMap, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(bytes []byte) (map[string]string, error) {
	type pair struct {
		Path string
		URL  string
	}

	type pairs struct {
		Pairs []pair
	}

	var prs pairs

	err := yaml.UnmarshalStrict(bytes, &prs)
	if err != nil {
		return nil, err
	}
	res := make(map[string]string, len(prs.Pairs))
	for _, r := range prs.Pairs {
		res[r.Path] = r.URL
	}
	return res, nil
}
