package net

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	beforeAll()
	code := m.Run()
	afterAll()

	os.Exit(code)
}

// start peer

func TestStartPeer(t *testing.T) {
	beforeEach()

	singleton := GetPeerInstance()
	if singleton == nil || testPeer != singleton {
		t.Error("singleton not set")
		return
	}

	p, err := StartPeer("localhost", 5685, "../../test/data/tmp", "../../test/data")
	if err != nil {
		t.Error(err)
		return
	}

	if singleton == nil || p != singleton {
		t.Error("singleton should not be changed once instantiated")
		return
	}
}

func TestConnectToPeer(t *testing.T) {
	beforeEach()

	if len(testPeer.peers) == 0 {
		t.Error("Peer not tracked/connected")
		return
	}

	mpClient := testPeer.server.clients[0]
	mpClient.SendString("test message\n")
	time.Sleep(25 * time.Millisecond)

	if mockPeer.GetLastMessage() != "test message\n" {
		t.Error("Test message not received")
		return
	}
}
