package peer

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

const (
	PEER_PORT uint = 7887
)

/**

! The test mock servers won't be active for all tests
  so when running these tests, several warning logs
	about it will be sent. Ignore these

*/

func TestMain(m *testing.M) {
	beforeAll()
	code := m.Run()
	afterAll()

	os.Exit(code)
}

func beforeAll() {
	util.InitLogger(true)

	// config
	viper.Set("meta.directory", "../../../test/data/tmp/.toolkit")
	viper.Set("games.installFolder", "../../../test/data/tmp")
	viper.Set("meta.hashes.workerCount", 5)

	_, err := StartPeer(Config{false, false, false, true, false, 5}, "localhost", 7887, "../../../test/data/tmp", "../../../test/data/tmp/.toolkit")
	if err != nil {
		log.Printf("Error starting test peer")
		os.Exit(1)
	}
}

func afterAll() {
	Peer().Close()
	testutil.ClearTmp("../../../")
}

// utility

// create a new mock peer to test the peer
func createMockPeer(t *testing.T) (*testutil.MockPeer, tcp.TCPConnection) {
	t.Helper()
	mp, err := testutil.StartMockPeer(PEER_PORT, true)
	if err != nil {
		t.Fatalf("Error starting mock peer: %s", err)
	}
	time.Sleep(25 * time.Millisecond)

	t.Cleanup(func() {
		mp.Close()
	})

	clients := Peer().server.Clients()
	util.Logger.Info("Mock peer started")
	return mp, clients[len(clients)-1]
}
