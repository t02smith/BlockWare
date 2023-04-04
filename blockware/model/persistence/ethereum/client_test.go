package ethereum

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestStartClient(t *testing.T) {
	t.Skip()
	t.Run("success", func(t *testing.T) {
		client, _, err := StartClient("ws://localhost:8545")
		if err != nil {
			t.Fatalf("Failed to connect to network. Is ganache running?")
		}

		client.Close()
	})

	t.Run("unknown network", func(t *testing.T) {
		t.Skip() // TODO
		_, _, err := StartClient("ws://fake.t02smith.com")
		if err == nil {
			t.Fatalf("Ethereum network not expected to be at domain")
		}
	})

}

/*

 */

func TestGenerateAuthInstance(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		auth, err := GenerateAuthInstance(testutil.Accounts[4][1])
		assert.Nil(t, err)
		assert.Zero(t, auth.Value.Cmp(big.NewInt(0)))
		assert.True(t, auth.GasLimit == uint64(3000000))
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("invalid private key", func(t *testing.T) {
			tops, err := GenerateAuthInstance("010")
			assert.Nil(t, tops)
			assert.NotNil(t, err)
		})
	})
}

/*

 */

func TestAddress(t *testing.T) {
	t.Run("private key found", func(t *testing.T) {
		addr := Address()
		assert.NotEqual(t, common.Address{}, addr)
	})

	t.Run("private key not found", func(t *testing.T) {
		old := _privateKey
		_privateKey = nil
		t.Cleanup(func() { _privateKey = old })

		addr := Address()
		assert.Equal(t, common.Address{}, addr)
	})
}
