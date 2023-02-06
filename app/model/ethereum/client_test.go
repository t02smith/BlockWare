package ethereum

import "testing"

func TestStartClient(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		client, err := StartClient("http://localhost:8545")
		if err != nil {
			t.Fatalf("Failed to connect to network. Is ganache running?")
		}

		client.Close()
	})

	t.Run("unknown network", func(t *testing.T) {
		t.Skip() // TODO
		_, err := StartClient("http://fake.t02smith.com")
		if err == nil {
			t.Fatalf("Ethereum network not expected to be at domain")
		}
	})

}
