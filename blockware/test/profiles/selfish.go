package profiles

import (
	"strings"

	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

unreliable client:

- may or may not respond to requests randomly
- rate can be set from 0 (always) - 100 (never)

*/

var selfishPrivateKey string = testutil.Accounts[5][1]

func selfishRun() {
	p := peer.Peer()

	p.SetOnMessage(func(s []string, t tcp.TCPConnection) error {
		// ignore message but reflect back to sender
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
