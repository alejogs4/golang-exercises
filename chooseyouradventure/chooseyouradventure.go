package chooseyouradventure

import (
	"encoding/json"
	"fmt"
	"os"
)

// Story is the representation of how story json will looks like
type Story map[string]StoryComponent

// StoryComponent is how every story chapter will be, with title a list of paraghaps and new options to go
type StoryComponent struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

// StoryOption is the reference to a new chapter
type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// StartAdventure will be the entry point to choose your adventure story
func StartAdventure() {
	storyComponents, error := parseStoryJSON("story.json")
	fmt.Println(storyComponents["intro"].Title, error)
}

func parseStoryJSON(route string) (Story, error) {
	var story Story

	file, error := os.Open(route)
	defer file.Close()
	if error != nil {
		return nil, error
	}

	jsonDecoder := json.NewDecoder(file)
	error = jsonDecoder.Decode(&story)
	if error != nil {
		return nil, error
	}

	return story, nil
}
