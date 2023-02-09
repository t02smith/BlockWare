package ethereum

import (
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func beforeAll() {
	util.InitLogger()
	testutil.SetupTestConfig()
	testutil.SetupTmp("../../")

	viper.Set("eth.keystore.directory", "../../test/data/tmp")
	viper.Set("eth.keystore.password", "test")
	return

	err := setupTestGame()
	if err != nil {
		util.Logger.Error(err)
		os.Exit(1)
	}

	_, _, err = StartClient("http://localhost:8545")
	if err != nil {
		util.Logger.Error(err)
	}

	_, _, err = DeployLibraryContract(testutil.Accounts[0][1])
	if err != nil {
		util.Logger.Error(err)
	}
}

func TestMain(m *testing.M) {

	beforeAll()
	code := m.Run()
	testutil.ClearTmp("../../")

	os.Exit(code)
}

//

func setupTestGame() error {

	_, err := os.Stat("../../test/data/.toolkit/toolkit-1.0.4-google.com.json")
	if err == nil {
		return nil
	}

	datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()
	game, err := games.CreateGame("toolkit", "1.0.4", datetime, "google.com", "../../test/data/testdir", 256, nil)

	if err != nil {
		log.Printf("Error creating game: %s\n", err)
		return err
	}

	// write game to file
	err = games.OutputToFile(game)
	if err != nil {
		log.Printf("Error writing game to file: %s\n", err)
		return err
	}

	return nil
}

func fetchTestGame() (*games.Game, error) {
	games, err := games.LoadGames("../../test/data/.toolkit")
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
