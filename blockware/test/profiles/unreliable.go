package profiles

import (
	"math/rand"
	"time"

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

var unrelaiblePrivateKey string = testutil.Accounts[3][1]

const unreliableRate int = 50

func unreliableRun() {
	p := peer.Peer()
	r := rand.New(rand.NewSource(time.Now().Unix()))

	p.SetOnMessage(func(s []string, t tcp.TCPConnection) error {

		if unreliableRate < r.Intn(100) || s[0] != "BLOCK" {
			util.Logger.Debug("Responding to requests")
			return peer.OnMessage(s, t)
		} else {
			util.Logger.Debug("Not responding to requests")
		}

		return nil
	})

	if err := SetupGame(); err != nil {
		util.Logger.Fatal(err)
	}

	// * listen for connections
	util.Logger.Info("PROFILE 2: listening")
	<-make(chan bool) // wait forever :/
}
