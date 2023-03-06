package ethereum

import (
	"os"
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func TestMain(m *testing.M) {
	beforeAll()
	code := m.Run()
	afterAll()

	os.Exit(code)
}

func beforeAll() {
	util.InitLogger()
	testutil.SetupTestConfig()

	_, _, err := StartClient("ws://localhost:8545")
	if err != nil {
		util.Logger.Fatal(err)
	}
}

func afterAll() {
	testutil.ClearTmp("../../../")
}
