package chooseyouradventure

import (
	"net/http"
	"strings"
	"text/template"
)

func createStoryHandler(story Story) http.Handler {
	return storyHandler{story: story}
}

type storyHandler struct {
	story Story
}

func (handler storyHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	urlPath := strings.TrimSpace(request.URL.Path)
	storyTemplate := template.Must(template.New("").Parse(HTMLStoryTemplate))

	if urlPath == "/" || urlPath == "" {
		urlPath = "/intro"
	}

	urlPath = urlPath[1:]
	chapter, ok := handler.story[urlPath]
	if ok {
		error := storyTemplate.Execute(response, chapter)

		if error != nil {
			http.Error(response, "Something went wrong", http.StatusInternalServerError)
		}
	} else {
		notFoundTemplate := template.Must(template.New("").Parse(NotFoundStoryTemplate))
		error := notFoundTemplate.Execute(response, handler.story)

		if error != nil {
			http.Error(response, "Something went wrong", http.StatusInternalServerError)
		}
	}
}
