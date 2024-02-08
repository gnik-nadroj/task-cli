package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
    Use:   "create [name] [description]",
    Short: "Create a tasks",
    Long: ``,
    Args: cobra.MinimumNArgs(2),

    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Create: " + strings.Join(args, " "))
    },
}

func init() {
    rootCmd.AddCommand(createCmd)
}