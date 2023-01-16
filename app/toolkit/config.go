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

	// where toolkit data is stored
	viper.SetDefault("meta.directory", "/home/tom/.toolkit")

	// max amount of threads to use
	viper.SetDefault("meta.threadPoolSize", 10)

	// where hashes are stored inside the toolkit folder
	viper.SetDefault("meta.hashes.directory", "/hashes")

	// default shard size when hashing
	viper.SetDefault("meta.hashes.shardSize", 16384)

	// user info
	viper.SetDefault("user.info.domain", "")
	viper.SetDefault("user.info.name", "")

}
