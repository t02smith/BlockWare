/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/lib"
	"github.com/t02smith/part-iii-project/toolkit/web"
)

func main() {
	SetupConfig()
	lib.InitLogger()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lib.SetupToolkitEnvironment()
	web.StartFrontend()
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

}
