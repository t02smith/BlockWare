package games

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	model "github.com/t02smith/part-iii-project/toolkit/model/util"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

var testGameRootHash [32]byte

func TestMain(m *testing.M) {
	util.InitLogger(true)
	testutil.SetupTestConfig()
	model.SetupToolkitEnvironment()

	// gamesTestSetup()
	code := m.Run()
	// gamesTestTeardown()

	os.Exit(code)
}

// create a test game and store it in long term storage
func setupTestGame(t *testing.T) {
	datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()
	game, err := CreateGame(NewGame{"toolkit", "1.0.4", datetime, "google.com", "../../../test/data/testdir", big.NewInt(0), 256, "../../../test/data/assets"}, nil)

	if err != nil {
		t.Fatal(err)
	}

	testGameRootHash = game.RootHash

	// write game to file
	err = OutputAllGameDataToFile(game)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.Remove(filepath.Join("meta.directory", "games", fmt.Sprintf("%x", game.RootHash)))
		os.Remove(filepath.Join("meta.directory", "hashes", fmt.Sprintf("%x.hash", game.RootHash)))
	})

}

func fetchTestGame(t *testing.T) *Game {
	games, err := LoadGames("../../../test/data/.toolkit/games")
	if err != nil {
		t.Fatal(err)
	}

	if len(games) == 0 {
		setupTestGame(t)
		return fetchTestGame(t)
	}

	var testGame *Game
	for _, g := range games {
		if bytes.Equal(g.RootHash[:], testGameRootHash[:]) {
			testGame = g
		}
	}

	if testGame == nil {
		setupTestGame(t)
		return fetchTestGame(t)
	}

	err = testGame.readHashData()
	if err != nil {
		t.Fatal(err)
	}

	return testGame
}

// create a test download
func setupTestDownload(t *testing.T) *Game {
	g := fetchTestGame(t)

	err := g.SetupDownload()
	if err != nil && !os.IsExist(err) {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		g.CancelDownload()
	})

	return g
}

// setup/teardown functions

func gamesTestSetup() {
	testutil.ClearTmp("../../../")
}

func gamesTestTeardown() {
	testutil.ClearTmp("../../../")
}
