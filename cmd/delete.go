package cmd

import (
	"fmt"
	"strconv"
	"tasks/internal"

	"github.com/spf13/cobra"
)

func init() {

}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "delete a task by ID",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No ID provided")
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
		if internal.DeleteTask(id) {
			fmt.Println("Deleted task:", args[0])
		} else {
			fmt.Println("Failed to delete task:", args[0])
		}
		internal.SaveToFile("tasks.csv")
	},
}
