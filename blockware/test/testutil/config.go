package testutil

import (
	"log"

	"github.com/spf13/viper"
)

func SetupTestConfig() {
	viper.Set("meta.hashes.workerCount", 5)
	viper.Set("meta.directory", "../../../test/data/.toolkit")
	viper.Set("games.installFolder", "../../../test/data/tmp")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
