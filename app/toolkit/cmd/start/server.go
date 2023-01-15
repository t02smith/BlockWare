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
	Short: "Start a new server for incoming connections",
	Long:  ``,
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
