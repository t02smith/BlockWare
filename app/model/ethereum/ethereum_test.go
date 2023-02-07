package ethereum

import (
	"os"
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func beforeAll() {
	util.InitLogger()
	testutil.SetupTmp("../../")
}

func TestMain(m *testing.M) {

	beforeAll()
	code := m.Run()
	testutil.ClearTmp("../../")

	os.Exit(code)
}
