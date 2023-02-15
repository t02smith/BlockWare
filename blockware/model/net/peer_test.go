package net

import (
	"os"
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestMain(m *testing.M) {
	beforeAll()
	code := m.Run()
	afterAll()

	os.Exit(code)
}

// start peer

func TestConnectToPeer(t *testing.T) {
	testutil.ShortTest(t)
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
