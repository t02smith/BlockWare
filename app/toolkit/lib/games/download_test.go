package games

import (
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/util"
)

func TestSetupDownload(t *testing.T) {
	games, err := LoadGames("../../test/data/.toolkit")
	if err != nil {
		t.Error(err)
		return
	}

	if len(games) == 0 {
		t.Error("No games present in the test folder")
		return
	}

	g := games[0]
	err = g.ReadHashData()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = SetupDownload(g)
	if err != nil {
		t.Error(err)
		return
	}

	util.ClearTmp("../../")
}
