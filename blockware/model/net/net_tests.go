package net

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

var (
	testPeer       *peer
	mockPeer       *testutil.MockPeer
	mockPeerClient PeerIT
)

func beforeAll() {
	util.InitLogger()
	tp, err := StartPeer("localhost", 7887, "../../test/data/tmp", "../../test/data/.toolkit")
	if err != nil {
		log.Printf("Error starting test peer")
		os.Exit(1)
	}
	testPeer = tp

	time.Sleep(25 * time.Millisecond)
	mockPeer, err = testutil.StartMockPeer(7887)
	if err != nil {
		log.Printf("Error starting mock peer")
		os.Exit(1)
	}
	time.Sleep(25 * time.Millisecond)

	mockPeerClient = testPeer.server.clients[0]

	// config
	viper.Set("meta.directory", "../../test/data/.toolkit")
	viper.Set("meta.hashes.directory", "../../test/data/.toolkit/hashes")
	viper.Set("games.installFolder", "../../test/data/tmp")
	viper.Set("games.tracker.directory", "../../test/data/tmp/tracker")

	testutil.SetupTmp("../../")
}

func beforeEach() {
	mockPeer.Clear()
}

func afterAll() {
	mockPeer.Close()
	testPeer.Close()
}

//

func fetchTestGame() (*games.Game, error) {
	games, err := games.LoadGames("../../test/data/.toolkit")
	if err != nil {
		return nil, err
	}

	if len(games) == 0 {
		return nil, errors.New("no games present in the test folder")
	}

	g := games[0]
	err = g.ReadHashData()
	if err != nil {
		return nil, err
	}

	return g, nil
}
