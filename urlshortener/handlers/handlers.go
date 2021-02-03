package handlers

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(urls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if destination, ok := urls[r.URL.Path]; ok {
			http.Redirect(rw, r, destination, http.StatusFound)
			return
		}

		fallback.ServeHTTP(rw, r)
	}
}

func YamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsUrlsFromYaml, error := parseYamlBytes(yamlBytes)
	if error != nil {
		return nil, error
	}

	urlsMap := buildUrlsMap(pathsUrlsFromYaml)
	return MapHandler(urlsMap, fallback), nil
}

func parseYamlBytes(yamlData []byte) ([]pathUrls, error) {
	var yamlInformation []pathUrls
	error := yaml.Unmarshal(yamlData, &yamlInformation)

	if error != nil {
		return nil, error
	}

	return yamlInformation, nil
}

func buildUrlsMap(urls []pathUrls) map[string]string {
	urlsMap := make(map[string]string)
	for _, path := range urls {
		urlsMap[path.Path] = path.URL
	}

	return urlsMap
}

type pathUrls struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
