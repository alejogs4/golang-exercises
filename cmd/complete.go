package cmd

import (
	"fmt"
	"strconv"

	"example.com/alejogs4/learning/taskspersistance"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Remove given task from to do tasks",
	Run: func(cmd *cobra.Command, args []string) {
		for _, id := range args {
			taskID, error := strconv.Atoi(id)

			if error != nil {
				fmt.Printf("Not possible to find a task with id %d\n", taskID)
			} else {
				boltProvider := taskspersistance.BoltProvider{}
				boltProvider.DeleteTask(taskID)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(completeCmd)
}
