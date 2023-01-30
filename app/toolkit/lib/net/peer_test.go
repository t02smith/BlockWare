package net

import (
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

// start peer

func TestStartPeer(t *testing.T) {
	p, err := StartPeer("localhost", 7887, "../../test/data/tmp", "../../test/data")
	if err != nil {
		t.Error(err)
		return
	}

	if singleton == nil || p != singleton {
		t.Error("singleton not set")
	}

	p, err = StartPeer("localhost", 5685, "../../test/data/tmp", "../../test/data")
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
	p, err := StartPeer("localhost", 7887, "../../test/data/tmp", "../../test/data")
	if err != nil {
		t.Error(err)
		return
	}

	mp, err := testutil.StartMockPeer(7887)
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(25 * time.Millisecond)
	if len(p.peers) == 0 {
		t.Error("Peer not tracked/connected")
		return
	}

	mpClient := p.server.clients[0]
	mpClient.SendString("test message\n")
	time.Sleep(25 * time.Millisecond)

	if mp.GetLastMessage() != "test message\n" {
		t.Error("Test message not received")
		return
	}
}
