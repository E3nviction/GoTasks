package cmd

import (
	"fmt"
	"strconv"
	"tasks/internal"

	"github.com/spf13/cobra"
)

func init() {

}

var moveCmd = &cobra.Command{
	Use:   "move [id] [new-id]",
	Short: "change the id of a task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No IDs provided")
			return
		}
		if len(args) < 2 {
			fmt.Println("No new ID provided")
			return
		}
		if len(args) > 2 {
			fmt.Println("Too many arguments")
			return
		}
		internal.LoadFromFile("tasks.csv")
		internal.ClearFile("tasks.csv")
		id, err := strconv.Atoi(args[0])
		id2, err2 := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID:", args[0])
			return
		}
		if err2 != nil {
			fmt.Println("Invalid ID:", args[1])
			return
		}

		if id == id2 {
			fmt.Println("IDs must be different")
			return
		}
		if internal.MoveTask(id, id2) {
			fmt.Println("Moved task:", args[0], "to", args[1])
		} else {
			fmt.Println("Failed to move task:", args[0])
		}
		internal.SaveToFile("tasks.csv")
	},
}
