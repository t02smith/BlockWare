package library

import (
	"errors"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func beforeAll() {
	util.InitLogger()
	testutil.SetupTestConfig()

	_, err := peer.StartPeer(peer.Config{ContinueDownloads: false, LoadPeersFromFile: false, ServeAssets: false}, "localhost", 7887, "../../../test/data/tmp", "../../../test/data/.toolkit")
	if err != nil {
		log.Printf("Error starting test peer")
		os.Exit(1)
	}

	_, _, err = ethereum.StartClient("ws://localhost:8545")
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
	testutil.ClearTmp("../../../")

	os.Exit(code)
}

//

func setupTestGame() error {
	_, err := os.Stat("../../../test/data/.toolkit/toolkit-1.0.4-google.com.json")
	if err == nil {
		return nil
	}

	datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()
	game, err := games.CreateGame(games.NewGame{
		Title:       "toolkit",
		Version:     "1.0.4",
		ReleaseDate: datetime,
		Developer:   "google.com",
		RootDir:     "../../../test/data/testdir",
		Price:       big.NewInt(0),
		ShardSize:   256}, nil)

	if err != nil {
		log.Printf("Error creating game: %s\n", err)
		return err
	}

	// write game to file
	err = games.OutputAllGameDataToFile(game)
	if err != nil {
		log.Printf("Error writing game to file: %s\n", err)
		return err
	}

	return nil
}

func fetchTestGame() (*games.Game, error) {
	games, err := games.LoadGames("../../../test/data/.toolkit/games")
	if err != nil {
		return nil, err
	}

	if len(games) == 0 {
		return nil, errors.New("No games present in the test folder")
	}

	g := games[0]
	_, err = g.GetData()
	if err != nil {
		return nil, err
	}

	return g, nil
}
