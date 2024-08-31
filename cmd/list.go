package cmd

import (
	"tasks/internal"

	"github.com/spf13/cobra"
)

var listAll bool

func init() {
	listCmd.Flags().BoolVarP(&listAll, "all", "a", false, "List all tasks including Completed tasks")
}

var listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if listAll {
			internal.LoadFromFile("tasks.csv")
			internal.ListTasks(true)
			return
		}
		internal.LoadFromFile("tasks.csv")
		internal.ListTasks(false)
	},
}
