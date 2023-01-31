package games

import (
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/lib"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestMain(m *testing.M) {
	testutil.SetupTestConfig()
	lib.SetupToolkitEnvironment()

	err := setupTestGame()
	if err != nil {
		log.Println(err)
		gamesTestTeardown()
		os.Exit(1)
	}

	old := verifyDomain
	mockVerifyDomain = func(domain string) (bool, error) {
		return true, nil
	}

	gamesTestSetup()
	code := m.Run()
	gamesTestTeardown()

	mockVerifyDomain = old
	os.Exit(code)
}

// create a test game and store it in long term storage
func setupTestGame() error {

	_, err := os.Stat("../../test/data/.toolkit/toolkit-1.0.4-google.com.json")
	if err == nil {
		return nil
	}

	datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()
	game, err := CreateGame("toolkit", "1.0.4", datetime, "google.com", "../../test/data/testdir", 256)

	if err != nil {
		log.Printf("Error creating game: %s\n", err)
		return err
	}

	// write game to file
	err = OutputToFile(game)
	if err != nil {
		log.Printf("Error writing game to file: %s\n", err)
		return err
	}

	return nil
}

func fetchTestGame() (*Game, error) {
	games, err := LoadGames("../../test/data/.toolkit")
	if err != nil {
		return nil, err
	}

	if len(games) == 0 {
		return nil, errors.New("No games present in the test folder")
	}

	g := games[0]
	err = g.ReadHashData()
	if err != nil {
		return nil, err
	}

	return g, nil
}

// create a test download
func setupTestDownload() (*Download, *Game, error) {
	g, err := fetchTestGame()
	if err != nil {
		return nil, nil, err
	}

	d, err := setupDownload(g)
	if err != nil && !os.IsExist(err) {
		return nil, nil, err
	}

	return d, g, nil
}

// setup/teardown functions

func gamesTestSetup() {
	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")
}

func gamesTestTeardown() {
	testutil.ClearTmp("../../")
}
