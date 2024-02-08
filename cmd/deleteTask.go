package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
    Use:   "delete [id]",
    Short: "Create a tasks",
    Long: ``,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Delete: " + strings.Join(args, " "))
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}