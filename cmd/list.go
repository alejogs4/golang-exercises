package cmd

import (
	"fmt"
	"log"

	"example.com/alejogs4/learning/taskspersistance"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all undone tasks",
	Run: func(cmd *cobra.Command, args []string) {
		boltProvider := taskspersistance.BoltProvider{}
		tasks, error := boltProvider.ListAllTasks()

		if error != nil {
			log.Fatal(error)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete")
			return
		}

		for _, task := range tasks {
			fmt.Printf("%d %s\n", task.Key, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
