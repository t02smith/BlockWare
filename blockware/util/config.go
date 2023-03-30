package util

import "github.com/spf13/viper"

func SetupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// viper.AddConfigPath("$HOME/.toolkit")
	viper.AddConfigPath(".")

	defaultConfig()
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Warnf("Error reading config file %s", err)
	}

	err = viper.SafeWriteConfig()
	if err != nil {
		Logger.Warnf("Error creating config file %s", err)
	}
}

func defaultConfig() {

	// toolkit meta info

	// where toolkit data is stored
	viper.SetDefault("meta.directory", "./test/data/.toolkit")
	viper.SetDefault("meta.log", "stdout")

	viper.SetDefault("contract.address", "0x750cf6392175f94ff5014803a0bb6b79de543337")

	// default shard size when hashing
	viper.SetDefault("meta.hashes.shardSize", 16384)

	// how many threads should be hashing at any given time
	viper.SetDefault("meta.hashes.workerCount", 5)

	// game info
	viper.SetDefault("games.installFolder", "./test/data/tmp")

	// user info
	viper.SetDefault("user.info.name", "")

	// eth
	viper.SetDefault("eth.address", "http://localhost:8545")

	// net
	viper.SetDefault("net.port", 6749)
	viper.SetDefault("net.peer.continuedownloads", true)
	viper.SetDefault("net.peer.useKnownPeers", false)
	viper.SetDefault("net.peer.serveassets", true)
	viper.SetDefault("net.peer.dovalidation", true)
}
