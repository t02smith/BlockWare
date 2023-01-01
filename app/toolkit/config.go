package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func SetupConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	defaultConfig()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config: %s\n", err)
	}

	err = viper.SafeWriteConfig()
	if err != nil {
		fmt.Printf("Error writing config: %s\n", err)
	}
}

func defaultConfig() {

	viper.SetDefault("meta.hashes.directory", ".toolkit")

}
