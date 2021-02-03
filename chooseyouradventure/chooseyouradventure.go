package chooseyouradventure

import (
	"fmt"
	"net/http"
	"os"
)

// StartAdventure will be the entry point to choose your adventure story
func StartAdventure() {
	storyComponents, error := parseStoryJSON("story.json")
	if error != nil {
		fmt.Println("Error getting stories")
		return
	}

	http.ListenAndServe(":3000", createStoryHandler(storyComponents))
}

func parseStoryJSON(route string) (Story, error) {
	file, error := os.Open(route)
	defer file.Close()
	if error != nil {
		return nil, error
	}

	story, error := NewStoryFromJSON(file)
	if error != nil {
		return nil, error
	}

	return story, nil
}
