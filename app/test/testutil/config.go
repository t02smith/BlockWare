package testutil

import (
	"log"

	"github.com/spf13/viper"
)

func SetupTestConfig() {
	viper.Set("meta.hashes.workerCount", 5)
	viper.Set("meta.directory", "../../test/data/.toolkit")
	viper.Set("meta.hashes.directory", "../../test/data/.toolkit/hashes")
	viper.Set("games.installFolder", "../../test/data/tmp")
	viper.Set("games.tracker.directory", "../../test/data/tmp/tracker")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
