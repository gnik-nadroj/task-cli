package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getTasksCmd = &cobra.Command{
    Use:   "get",
    Short: "Get all tasks",
    Long: ``,
    Args: cobra.MinimumNArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Get All Taks")
    },
}

func init() {
    rootCmd.AddCommand(getTasksCmd)
}