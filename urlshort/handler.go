package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value, present := pathsToUrls[r.URL.Path]
		if present {
			http.Redirect(w, r, value, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

// PathsToURL ---
type PathsToURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler ---
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	//Parse Url
	var PathsToURLs []PathsToURL
	err := yaml.Unmarshal(yml, &PathsToURLs)
	if err != nil {
		return nil, err
	}
	// Build map
	MapOfPathsToURL := make(map[string]string)
	for _, value := range PathsToURLs {
		MapOfPathsToURL[value.Path] = value.URL
	}
	return MapHandler(MapOfPathsToURL, fallback), nil
}
