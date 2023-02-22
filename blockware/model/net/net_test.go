package net

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

var (
	testPeer       *peer
	mockPeer       *testutil.MockPeer
	mockPeerClient PeerIT
)

/**

! The test mock servers won't be active for all tests
  so when running these tests, several warning logs
	about it will be sent. Ignore these

*/

func beforeAll() {
	util.InitLogger()

	// config
	viper.Set("meta.directory", "../../test/data/.toolkit")
	viper.Set("games.installFolder", "../../test/data/tmp")

	tp, err := StartPeer(PeerConfig{false, false}, "localhost", 7887, "../../test/data/tmp", "../../test/data/.toolkit")
	if err != nil {
		log.Printf("Error starting test peer")
		os.Exit(1)
	}
	testPeer = tp

	time.Sleep(25 * time.Millisecond)
	mockPeer, err = testutil.StartMockPeer(7887, true)
	if err != nil {
		log.Printf("Error starting mock peer")
		os.Exit(1)
	}
	time.Sleep(25 * time.Millisecond)

	mockPeerClient = testPeer.server.clients[0]
}

func beforeEach() {
	mockPeer.Clear()
}

func afterAll() {
	if mockPeer != nil {
		mockPeer.Close()
	}

	testPeer.Close()
	testutil.ClearTmp("../../")
}