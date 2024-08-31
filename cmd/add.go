package cmd

import (
	"fmt"
	"tasks/internal"

	"github.com/spf13/cobra"
)

func init() {
}

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No task provided")
			return
		}
		if len(args) > 1 {
			fmt.Println("Too many arguments")
			return
		}
		internal.LoadFromFile("tasks.csv")
		if internal.AddTask(args[0]) {
			fmt.Println("Added task:", args[0])
		} else {
			fmt.Println("Failed to add task:", args[0])
		}

		internal.SaveToFile("tasks.csv")
	},
}
