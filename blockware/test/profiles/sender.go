package profiles

import (
	"strings"

	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

var senderPrivateKey string = testutil.Accounts[4][1]

func senderRun() {
	p := peer.Peer()

	p.SetOnMessage(func(s []string, t tcp.TCPConnection) error {
		// handle incoming message
		if err := peer.OnMessage(s, t); err != nil {
			return err
		}

		// reflect request back at peer
		t.SendString(strings.Join(s, ";") + "\n")

		return nil
	})

	if err := SetupGame(); err != nil {
		util.Logger.Fatal(err)
	}

	// * listen for connections
	util.Logger.Info("PROFILE 2: listening")
	<-make(chan bool) // wait forever :/
}
