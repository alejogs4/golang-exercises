package main

import (
	"log"
	"path/filepath"

	"example.com/alejogs4/learning/cmd"
	"example.com/alejogs4/learning/taskspersistance"
	"github.com/mitchellh/go-homedir"
)

// import "example.com/alejogs4/learning/chooseyouradventure"

// "example.com/alejogs4/learning/quizgame"
// "example.com/alejogs4/learning/urlshortener"

func main() {
	currentHomeDir, _ := homedir.Dir()
	error := taskspersistance.Init(filepath.Join(currentHomeDir, "tasks.db"))
	if error != nil {
		log.Fatal(error)
	}

	cmd.Execute()
}
