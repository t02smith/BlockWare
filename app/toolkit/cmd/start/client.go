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
	startClientHostname string
	startClientPort     uint
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Form a connection with a peer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := toolkitNet.InitTCPClient(startClientHostname, startClientPort)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	StartCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringVarP(&startClientHostname, "hostname", "l", "localhost", "The hostname of the server you're connecting to")
	clientCmd.Flags().UintVarP(&startClientPort, "port", "p", 5656, "The port the server is running on")
}
