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
	viper.Set("games.installFolder", "../../test/data/tmp")

	testutil.SetupTmp("../../")
}

func beforeEach() {
	mockPeer.Clear()
}

func afterAll() {
	mockPeer.Close()
	testPeer.Close()
}
