package lib

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	testutil "github.com/t02smith/part-iii-project/toolkit/test/util"
)

func utilTestSetup() {
	testutil.ClearTmp("../")
	viper.Set("meta.directory", "../test/data/tmp/.toolkit")
	viper.Set("meta.hashes.directory", "../test/data/tmp/.toolkit/hashes")

}

func utilTestTeardown() {
	testutil.ClearTmp("../")
}

func TestCreateDirectoryIfNotExistCorrect(t *testing.T) {
	utilTestSetup()
	defer utilTestTeardown()

	err := createDirectoryIfNotExist("../test/data/tmp/hellothere")
	if err != nil {
		t.Error(err)
		return
	}

	f, err := os.Stat("../test/data/tmp/hellothere")
	if err != nil {
		t.Error(err)
		return
	}

	if !f.IsDir() {
		t.Error("Directory not created. File exists but is not directory")
		return
	}
}

func TestSetupToolkitEnvironment(t *testing.T) {
	utilTestSetup()
	defer utilTestTeardown()

	// run function to be tested
	err := SetupToolkitEnvironment()
	if err != nil {
		t.Error(err)
		return
	}

	// assert directories created
	_, err = os.Stat("../test/data/tmp/.toolkit")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = os.Stat("../test/data/tmp/.toolkit/hashes")
	if err != nil {
		t.Error(err)
		return
	}
}
