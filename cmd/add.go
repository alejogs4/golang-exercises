package cmd

import (
	"log"
	"strings"

	"example.com/alejogs4/learning/taskspersistance"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task in the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		boltProvider := taskspersistance.BoltProvider{}

		error := boltProvider.CreateTask(task)
		if error != nil {
			log.Fatal(error)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
