/*
Copyright © 2022 Thomas Smith tcs1g20@soton.ac.uk
*/
package main

import (
	"embed"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/controller"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	model "github.com/t02smith/part-iii-project/toolkit/model/util"
	"github.com/t02smith/part-iii-project/toolkit/test/profiles"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	if profileRan := profiles.ProfileCLI(); profileRan {
		os.Exit(0)
	}

	// * setup
	util.Logger.Info("No profile selected => Running default application")
	util.SetupConfig()
	defer viper.WriteConfig()

	model.SetupToolkitEnvironment()

	// * model setup

	// ? start peer
	err := startPeer()
	if err != nil {
		log.Fatalf("Error starting peer %s", err)
	}
	defer peer.Peer().Close()

	// ? start ETH client
	_, _, err = ethereum.StartClient(viper.GetString("eth.address"))
	if err != nil {
		util.Logger.Fatalf("Error starting eth client %s", err)
	}
	defer ethereum.CloseEthClient()

	// * Wails application
	controller := controller.NewController()

	util.Logger.Info("Starting GUI")
	err = wails.Run(&options.App{
		Title:  "blockware",
		Width:  1300,
		Height: 850,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: controller.Startup,
		Bind: []interface{}{
			controller,
		},
		Logger: nil,
	})

	if err != nil {
		util.Logger.Error("Error:", err.Error())
	}

	lib := peer.Peer().Library()
	lib.Lock()
	for _, g := range lib.GetOwnedGames() {
		g.OutputToFile()
	}
	lib.Unlock()

}

func startPeer() error {
	util.Logger.Info("Attempting to start peer")
	_, err := peer.StartPeer(
		peer.Config{
			ContinueDownloads:  viper.GetBool("net.peer.continuedownloads"),
			LoadPeersFromFile:  false,
			ServeAssets:        viper.GetBool("net.peer.serve_assets"),
			SkipValidation:     viper.GetBool("net.peer.skip_validation"),
			TrackContributions: true,
			MaxConnections:     viper.GetUint("net.peer.max_connections"),
		},
		"localhost",
		viper.GetUint("net.port"),
		viper.GetString("games.installFolder"),
		viper.GetString("meta.directory"),
	)

	return err
}
