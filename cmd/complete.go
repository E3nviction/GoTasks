package cmd

import (
	"fmt"
	"strconv"
	"tasks/internal"

	"github.com/spf13/cobra"
)

func init() {

}

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "complete a task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No id provided")
			return
		}
		if len(args) > 1 {
			fmt.Println("Too many arguments")
			return
		}
		internal.LoadFromFile("tasks.csv")
		internal.ClearFile("tasks.csv")
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID:", args[0])
			return
		}
		if internal.CompleteTask(id) {
			fmt.Println("Completed task:", args[0])
		} else {
			fmt.Println("Failed to complete task:", args[0])
		}
		internal.SaveToFile("tasks.csv")
	},
}
