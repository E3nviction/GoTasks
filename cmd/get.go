package cmd

import (
	"fmt"
	"tasks/internal"

	"github.com/spf13/cobra"
)

func init() {

}

var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "get a task by name",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No name provided")
			return
		}
		if len(args) > 1 {
			fmt.Println("Too many arguments")
			return
		}
		internal.LoadFromFile("tasks.csv")
		fmt.Println(internal.GetTask(args[0]))
	},
}
