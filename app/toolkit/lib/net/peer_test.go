package net

import (
	"testing"
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

// func TestConnectToPeer(t *testing.T) {
// 	mockClient := InitServer("localhost", 9009)

// 	go mockClient.Start(func(s []string, pi PeerIT) {
// 		log.Printf("Mock Client received: %s", strings.Join(s, ";"))
// 		pi.SendString("hello")
// 	})

// 	p, err := StartPeer("localhost", 7887, "../../test/data/tmp", "../../test/data")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	time.Sleep(300 * time.Millisecond)
// 	err = p.ConnectToPeer("localhost", 9009)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	time.Sleep(300 * time.Millisecond)
// 	if len(p.clients) == 0 {
// 		t.Error("client not added to peer")
// 		return
// 	}

// 	peerClient := p.clients[0]
// 	peerClient.SendString("hello-there-test;")
// 	time.Sleep(300 * time.Millisecond)

// 	t.Error("")
// }
