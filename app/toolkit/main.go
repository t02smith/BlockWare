/*
Copyright © 2022 Thomas Smith tcs1g20@soton.ac.uk
*/
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/controller"
	"github.com/t02smith/part-iii-project/toolkit/model"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
)

func main() {
	SetupConfig()
	model.InitLogger()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	model.SetupToolkitEnvironment()
	startGin()
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

// start gin server

func startGin() {
	err := startPeer()
	if err != nil {
		model.Logger.Panicf("Failed to start peer %s", err)
	}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/css", "static/css")
	r.Static("/img", "static/img")

	r.LoadHTMLGlob("templates/**/*.html")
	controller.Router(r)

	model.Logger.Info("server started")
	r.Run()
}

func startPeer() error {
	_, err := net.StartPeer(
		"localhost",
		6748,
		viper.GetString("games.installFolder"),
		viper.GetString("meta.directory"),
	)

	return err
}
