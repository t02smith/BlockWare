package hash

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func TestMain(m *testing.M) {
	viper.Set("meta.hashes.workerCount", 5)
	util.InitLogger()

	code := m.Run()
	os.Exit(code)
}
