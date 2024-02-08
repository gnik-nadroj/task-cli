package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "task-cli",
    Short: "A simple command line apps to manage your Tacks",
    Long: `A tasks cli that provide: ...`,

    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello From Task Cli")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
