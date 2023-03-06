/*
Copyright © 2022 Thomas Smith tcs1g20@soton.ac.uk
*/
package main

import (
	"embed"
	"flag"
	"log"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/controller"
	"github.com/t02smith/part-iii-project/toolkit/model"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/test/profiles"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	util.InitLogger()

	// * FLAGS
	profile := flag.Uint("profile", 0, "Run the application as a profile. (default off | see profiles.go for details)")
	contractAddr := flag.String("contract", "", "The address of the deployed contract")

	flag.Parse()

	// ? run a profile
	if *profile != 0 {
		// if *contractAddr == "" {
		// 	util.Logger.Fatal("invalid contract address")
		// }

		util.Logger.Infof("profile %d selected", *profile)
		err := profiles.RunProfile(profiles.Profile(*profile), *contractAddr)
		if err != nil {
			util.Logger.Fatalf("Error running profile %d: %s", *profile, err)
		}
	}

	// * setup
	util.Logger.Info("No profile selected => Running default application")
	SetupConfig()
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
	})

	if err != nil {
		util.Logger.Error("Error:", err.Error())
	}

}

// VIPER CONFIG

func SetupConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// viper.AddConfigPath("$HOME/.toolkit")
	viper.AddConfigPath(".")

	defaultConfig()
	err := viper.ReadInConfig()
	if err != nil {
		util.Logger.Warnf("Error reading config file %s", err)
	}

	err = viper.SafeWriteConfig()
	if err != nil {
		util.Logger.Warnf("Error creating config file %s", err)
	}

}

func defaultConfig() {

	// toolkit meta info

	// where toolkit data is stored
	viper.SetDefault("meta.directory", "./test/data/.toolkit")

	viper.SetDefault("meta.log", "stdout")

	// max amount of threads to use
	viper.SetDefault("meta.threadPoolSize", 10)

	// where hashes are stored inside the toolkit folder

	// default shard size when hashing
	viper.SetDefault("meta.hashes.shardSize", 16384)

	// how many threads should be hashing at any given time
	viper.SetDefault("meta.hashes.workerCount", 5)

	// game info
	viper.SetDefault("games.installFolder", "./test/data/tmp")

	// user info
	viper.SetDefault("user.info.domain", "t02smith.com")
	viper.SetDefault("user.info.name", "tom")

	// eth
	viper.SetDefault("eth.address", "ws://localhost:8545")

	// net
	viper.SetDefault("net.port", 6749)
	viper.SetDefault("net.useKnownPeers", false)
}

func startPeer() error {
	util.Logger.Info("Attempting to start peer")
	_, err := peer.StartPeer(
		peer.PeerConfig{
			ContinueDownloads: true,
			LoadPeersFromFile: true,
			ServeAssets:       true,
		},
		"localhost",
		viper.GetUint("net.port"),
		viper.GetString("games.installFolder"),
		viper.GetString("meta.directory"),
	)

	return err
}
