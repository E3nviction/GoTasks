package cmd

import (
	"fmt"
	"strconv"
	"tasks/internal"

	"github.com/spf13/cobra"
)

func init() {

}

var editCmd = &cobra.Command{
	Use:   "edit [id] [description]",
	Short: "edit a task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No ID and description provided")
			return
		}
		if len(args) < 2 {
			fmt.Println("No description provided")
			return
		}
		if len(args) > 2 {
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
		if internal.EditTask(id, args[1]) {
			fmt.Println("Edited task:", args[0])
		} else {
			fmt.Println("Failed to edit task:", args[0])
		}

		internal.SaveToFile("tasks.csv")
	},
}
