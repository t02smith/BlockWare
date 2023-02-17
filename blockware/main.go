/*
Copyright © 2022 Thomas Smith tcs1g20@soton.ac.uk
*/
package main

import (
	"embed"
	"log"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model"
	"github.com/t02smith/part-iii-project/toolkit/model/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	util.InitLogger()
	SetupConfig()
	defer viper.WriteConfig()

	model.SetupToolkitEnvironment()

	// model setup
	err := startPeer()
	if err != nil {
		log.Fatalf("Error starting peer %s", err)
	}

	_, _, err = ethereum.StartClient(viper.GetString("eth.address"))
	if err != nil {
		util.Logger.Fatalf("Error starting eth client %s", err)
	}
	defer ethereum.CloseEthClient()

	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "blockware",
		Width:  1300,
		Height: 850,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// VIPER CONFIG

func SetupConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// viper.AddConfigPath("$HOME/.toolkit")
	viper.AddConfigPath(".")

	defaultConfig()

	err := viper.SafeWriteConfig()
	if err != nil {
		util.Logger.Warnf("Error creating config file %s", err)
	}

	viper.ReadInConfig()
}

func defaultConfig() {

	// toolkit meta info

	// where toolkit data is stored
	viper.SetDefault("meta.directory", "./test/data/.toolkit")

	viper.SetDefault("meta.log", "stdout")

	// max amount of threads to use
	viper.SetDefault("meta.threadPoolSize", 10)

	// where hashes are stored inside the toolkit folder
	viper.SetDefault("meta.hashes.directory", "./test/data/.toolkit/hashes")

	// default shard size when hashing
	viper.SetDefault("meta.hashes.shardSize", 16384)

	// how many threads should be hashing at any given time
	viper.SetDefault("meta.hashes.workerCount", 5)

	// game info
	viper.SetDefault("games.installFolder", "./test/data/tmp")
	viper.SetDefault("games.tracker.directory", "./test/data/.toolkit/tracker")

	// user info
	viper.SetDefault("user.info.domain", "t02smith.com")
	viper.SetDefault("user.info.name", "tom")

	// eth
	viper.SetDefault("eth.keystore.directory", "./test/data/.toolkit")
	viper.SetDefault("eth.keystore.password", "password")
	viper.SetDefault("eth.address", "ws://localhost:8545")
}

func startPeer() error {
	util.Logger.Info("Attempting to start peer")
	_, err := net.StartPeer(
		"localhost",
		6748,
		viper.GetString("games.installFolder"),
		viper.GetString("meta.directory"),
	)

	return err
}