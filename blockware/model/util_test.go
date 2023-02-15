package model

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	u "github.com/t02smith/part-iii-project/toolkit/util"
)

func TestMain(m *testing.M) {
	u.InitLogger()
	os.Exit(m.Run())
}

func utilTestSetup() {
	testutil.ClearTmp("../")
	viper.Set("meta.directory", "../test/data/tmp/.toolkit")
	viper.Set("meta.hashes.directory", "../test/data/tmp/.toolkit/hashes")

}

func utilTestTeardown() {
	testutil.ClearTmp("../")
}

func TestCreateDirectoryIfNotExistCorrect(t *testing.T) {
	testutil.ShortTest(t)
	utilTestSetup()
	defer utilTestTeardown()

	err := CreateDirectoryIfNotExist("../test/data/tmp/hellothere")
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
	testutil.ShortTest(t)
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

	_, err = os.Stat("../test/data/tmp/.toolkit/tracker")
	if err != nil {
		t.Error(err)
		return
	}
}
