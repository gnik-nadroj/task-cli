package cmd

import (
	"fmt"
	"net/rpc/jsonrpc"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
    Use:   "create [name] [description]",
    Short: "Create a tasks",
    Long: ``,
    Args: cobra.MinimumNArgs(2),

    Run: func(cmd *cobra.Command, args []string) {
      client, err := jsonrpc.Dial("tcp", "localhost:1234")
      if err != nil {
        fmt.Println("Error connecting to RPC server:", err)
        return
      }
      defer client.Close()

      var reply int
      err = client.Call("TaskService.CreateTask", , &reply)
      if err != nil {
        fmt.Println("Error calling RPC method:", err)
        return
      }

      fmt.Println("Result:", reply)
    },
}

func init() {
    rootCmd.AddCommand(createCmd)
}