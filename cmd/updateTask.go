package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
    Use:   "update [id] [name] [description]",
    Short: "Update a tasks",
    Long: ``,
    Args: cobra.MinimumNArgs(2),

    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Update: " + strings.Join(args, " "))
    },
}

func init() {
    rootCmd.AddCommand(updateCmd)
}