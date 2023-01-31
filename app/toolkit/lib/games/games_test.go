package games

import (
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

// setup/teardown functions

func gamesTestSetup() {
	testutil.ClearTmp("../../")
	testutil.SetupTmp("../../")
}

func gamesTestTeardown() {
	testutil.ClearTmp("../../")
}
