package games

import (
	"bytes"
	"errors"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

var testGameRootHash [32]byte

func TestMain(m *testing.M) {
	util.InitLogger(true)
	testutil.SetupTestConfig()
	model.SetupToolkitEnvironment()

	err := setupTestGame()
	if err != nil {
		log.Println(err)
		gamesTestTeardown()
		os.Exit(1)
	}

	// gamesTestSetup()
	code := m.Run()
	// gamesTestTeardown()

	os.Exit(code)
}

// create a test game and store it in long term storage
func setupTestGame() error {
	datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()
	game, err := CreateGame(NewGame{"toolkit", "1.0.4", datetime, "google.com", "../../../test/data/testdir", big.NewInt(0), 256, "../../../test/data/assets"}, nil)

	if err != nil {
		log.Printf("Error creating game: %s\n", err)
		return err
	}

	testGameRootHash = game.RootHash

	// write game to file
	err = OutputAllGameDataToFile(game)
	if err != nil {
		log.Printf("Error writing game to file: %s\n", err)
		return err
	}

	return nil
}

func fetchTestGame() (*Game, error) {
	games, err := LoadGames("../../../test/data/.toolkit/games")
	if err != nil {
		return nil, err
	}

	if len(games) == 0 {
		return nil, errors.New("No games present in the test folder")
	}

	var testGame *Game
	for _, g := range games {
		if bytes.Equal(g.RootHash[:], testGameRootHash[:]) {
			testGame = g
		}
	}
	if testGame == nil {
		util.Logger.Fatalf("error finding test game")
	}

	err = testGame.readHashData()
	if err != nil {
		return nil, err
	}

	return testGame, nil
}

// create a test download
func setupTestDownload(t *testing.T) (*Game, error) {
	g, err := fetchTestGame()
	if err != nil {
		return nil, err
	}

	err = g.SetupDownload()
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	t.Cleanup(func() {
		g.CancelDownload()
	})

	return g, nil
}

// setup/teardown functions

func gamesTestSetup() {
	testutil.ClearTmp("../../../")
}

func gamesTestTeardown() {
	testutil.ClearTmp("../../../")
}
