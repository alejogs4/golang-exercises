package chooseyouradventure

import (
	"encoding/json"
	"io"
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

// NewStoryFromJSON will return a story map from a file reader that it's a json file
func NewStoryFromJSON(reader io.Reader) (Story, error) {
	var story Story

	jsonDecoder := json.NewDecoder(reader)
	error := jsonDecoder.Decode(&story)
	if error != nil {
		return nil, error
	}

	return story, nil
}
