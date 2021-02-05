package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Get every passed task to done",
	Run: func(cmd *cobra.Command, args []string) {
		var tasksIds []int
		for _, taskID := range args {
			id, error := strconv.Atoi(taskID)
			if error != nil {
				fmt.Println("It was not possible parse: ", taskID)
			} else {
				tasksIds = append(tasksIds, id)
			}
		}

		fmt.Println(tasksIds)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
