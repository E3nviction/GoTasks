package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(completeCmd)
}

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
