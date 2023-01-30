package net

import (
	"log"
	"os"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

var (
	testPeer       *peer
	mockPeer       *testutil.MockPeer
	mockPeerClient PeerIT
)

func beforeAll() {
	tp, err := StartPeer("localhost", 7887, "../../test/data/tmp", "../../test/data")
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
}

func beforeEach() {
	mockPeer.Clear()
}

func afterAll() {
	mockPeer.Close()
	testPeer.Close()
}
