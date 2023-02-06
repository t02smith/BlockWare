/*
Copyright © 2022 Thomas Smith tcs1g20@soton.ac.uk
*/
package main

import (
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func main() {
	SetupConfig()
	util.InitLogger()

	client, err := ethereum.StartClient("http://localhost:8545")
	if err != nil {
		util.Logger.Fatal(err)
	}

	err = ethereum.DeployContracts(client, testutil.Accounts[0][1])
	if err != nil {
		util.Logger.Fatal(err)
	}

	return
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	// model.SetupToolkitEnvironment()

	// // model setup
	// err := startPeer()
	// if err != nil {
	// 	log.Fatalf("Error starting peer %s", err)
	// }

	// //
	// view.StartApp()
}

// VIPER CONFIG

func SetupConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// viper.AddConfigPath("$HOME/.toolkit")
	viper.AddConfigPath(".")

	defaultConfig()
	viper.SafeWriteConfig()

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

	// web interface
	viper.SetDefault("web.port", 8080)
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
