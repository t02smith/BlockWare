/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package start

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/lib/net"
)

var (
	peerHostname          string
	peerPort              uint
	peerGameInstallFolder string
	peerGameDataLocation  string
)

// peerCmd represents the peer command
var peerCmd = &cobra.Command{
	Use:   "peer",
	Short: "Start up a new peer",
	Long:  `Start a new peer. This is a background process that will remain running`,
	Run: func(cmd *cobra.Command, args []string) {
		// get default values
		if len(peerGameDataLocation) == 0 {
			peerGameDataLocation = viper.GetString("meta.directory")
		}

		// run peer
		_, err := net.StartPeer(peerHostname, peerPort, peerGameInstallFolder, peerGameDataLocation)
		if err != nil {
			log.Printf("Error creating peer: %s\n", err)
		}

		menu()
	},
}

func init() {
	StartCmd.AddCommand(peerCmd)

	peerCmd.Flags().StringVarP(&peerHostname, "server-hostname", "s", "localhost", "The hostname your server is running on")
	peerCmd.Flags().UintVarP(&peerPort, "server-port", "p", 3047, "What port your peer listens on")
	peerCmd.Flags().StringVarP(&peerGameInstallFolder, "install-games", "o", ".", "Where games are being installed")
	peerCmd.Flags().StringVarP(&peerGameDataLocation, "game-data", "d", "", "Where you game data is stored")

}
