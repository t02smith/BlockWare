package ethereum

import (
	"os"
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

func beforeAll() {
	util.InitLogger()
}

func TestMain(m *testing.M) {

	beforeAll()
	code := m.Run()

	os.Exit(code)
}
