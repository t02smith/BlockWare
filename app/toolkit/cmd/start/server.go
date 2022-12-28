/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package start

import (
	"fmt"

	"github.com/spf13/cobra"

	toolkitNet "github.com/t02smith/part-iii-project/toolkit/lib/net"
)

var (
	startServerPort uint
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := toolkitNet.InitServer("localhost", startServerPort)
		server.Start(func(b []byte, tc *toolkitNet.TCPServerClient) {
			fmt.Printf("%x\n", b)
		})
	},
}

func init() {
	StartCmd.AddCommand(serverCmd)

	serverCmd.Flags().UintVarP(&startServerPort, "port", "p", 5656, "The port to host the server on")
}
