package main

import (
	"github.com/spf13/viper"
)

func SetupConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("$HOME/.toolkit")
	viper.AddConfigPath(".")

	defaultConfig()

	viper.ReadInConfig()
	viper.SafeWriteConfig()

}

func defaultConfig() {

	// toolkit meta info
	viper.SetDefault("meta.directory", "/home/tom/.toolkit")
	viper.SetDefault("meta.hashes.directory", ".toolkit/hashes")
	viper.SetDefault("meta.hashes.shardSize", 16384)

	// user info
	viper.SetDefault("user.info.domain", "")
	viper.SetDefault("user.info.name", "")

}
