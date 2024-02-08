package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var getTaskCmd = &cobra.Command{
	Use:   "id [id]",
	Short: "Get A Specifix Task",
	Long: ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
	},
}


func init() {
    getTasksCmd.AddCommand(getTaskCmd)
}
