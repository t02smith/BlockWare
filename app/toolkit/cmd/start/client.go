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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
