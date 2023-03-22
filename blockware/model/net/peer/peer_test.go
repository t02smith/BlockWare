package peer

import (
	"bufio"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

// start peer

func TestConnectToPeer(t *testing.T) {
	// assert.Equal(t, 0, len(testPeer.peers), "Peer not tracked/connected")

	mp, it := createMockPeer(t)
	it.SendString("test message\n")
	time.Sleep(25 * time.Millisecond)

	assert.Equal(t, 1, len(Peer().peers), "Mock peer not tracked/connected")

	assert.NotEqual(t, "test message", mp.GetLastMessage(), "Test message not received got "+mp.GetLastMessage())
}

func TestLoadPeersFromFile(t *testing.T) {
	viper.Set("meta.directory", "../../../test/data/tmp")
	t.Cleanup(func() {
		viper.Set("meta.directory", "../../../test/data/.toolkit")
		testutil.ClearTmp("../../../")
	})

	t.Run("file not found", func(t *testing.T) {
		_, err := loadPeersFromFile()
		assert.NotNil(t, err, "file not found error expected")
	})

	f, err := os.Create("../../../test/data/tmp/peers.txt")
	if err != nil {
		t.Fatal(err)
	}

	writer := bufio.NewWriter(f)
	writer.WriteString("192.168.0.4:1234\n100.40.68.8:5000\n")
	writer.Flush()

	f.Close()

	t.Run("success", func(t *testing.T) {
		ps, err := loadPeersFromFile()
		assert.Nil(t, err, "no error expected loading peers")

		assert.Equal(t, 2, len(ps), "incorrect number of peers found")
		assert.Equal(t, "192.168.0.4:1234", ps[0], "incorrect peer #1")
		assert.Equal(t, "100.40.68.8:5000", ps[1], "incorrect peer #2")

	})

}

func TestSavePeersToFile(t *testing.T) {
	viper.Set("meta.directory", "../../../test/data/tmp")
	t.Cleanup(func() {
		viper.Set("meta.directory", "../../../test/data/.toolkit")
		testutil.ClearTmp("../../../")
	})

	t.Run("success", func(t *testing.T) {
		err := Peer().savePeersToFile()
		assert.Nil(t, err, "no error expected")

		_, err = loadPeersFromFile()
		assert.Nil(t, err, "no error expected")

		// assert.Equal(t, 1, len(ps), "only mock peer expected")
	})

}

func TestConnectToKnownPeers(t *testing.T) {
	t.Skip()

	t.Run("success", func(t *testing.T) {
		ports := []uint{6750, 6751, 6752}

		// * create test file
		f, err := os.Create("../../../test/data/.toolkit/peers.txt")
		if err != nil {
			t.Fatal(err)
		}

		f.WriteString("localhost:6750\nlocalhost:6751\nlocalhost:6752\n")
		f.Close()

		// * start mock servers
		err = testutil.StartMockServers(ports)
		assert.Nil(t, err, "mock servers should start succesffuly")
		time.Sleep(200 * time.Millisecond)

		p := Peer()
		p.connectToKnownPeers()

		assert.Equal(t, 3, len(p.clients), "connected to all mock servers")
		p.Broadcast("TEST_CONNECT_TO_KNOWN_PEERS\n")
		time.Sleep(10 * time.Millisecond)

		for _, port := range ports {
			ms, ok := testutil.MockServers[port]
			if !ok {
				t.Fatal("Mock server not stored => can't test")
			}

			assert.Equal(t, "TEST_CONNECT_TO_KNOWN_PEERS\n", ms.GetLastMessage(), "received test message")
		}

		testutil.CloseAllMockServers()
	})

}
