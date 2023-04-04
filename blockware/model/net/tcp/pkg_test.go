package tcp

import (
	"os"
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

func TestMain(m *testing.M) {
	util.InitLogger(true)

	os.Exit(m.Run())
}
