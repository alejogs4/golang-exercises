package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}

// Execute will be the entry point for task CLI
// As soon as users type "task" in their terminals
func Execute() {
	if error := RootCmd.Execute(); error != nil {
		log.Fatal(error)
	}
}
