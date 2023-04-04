package peer

import (
	"crypto/sha256"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

/*

function: validatePeer
purpose: initiate a validation check for a peer

? Test cases
success
	#1 => user not validated
	#2 => user already validated

*/

func TestValidatePeer(t *testing.T) {
	mp, con := createMockPeer(t)

	pd, ok := Peer().peers[con]
	if !ok {
		t.Fatal("mock peer not setup correctly")
	}

	t.Run("success", func(t *testing.T) {
		t.Run("user not validated", func(t *testing.T) {
			t.Cleanup(func() {
				mp.Clear()
				pd.Validator = nil
			})

			pd.ValidatePeer()
			time.Sleep(25 * time.Millisecond)

			assert.Equal(t, generateVALIDATE_REQ(pd.Validator.Message()), mp.GetLastMessage())
		})

		t.Run("user already validated", func(t *testing.T) {
			pd.Validator = &ethereum.AddressValidator{}
			pd.Validator.SetValid(true)

			pd.ValidatePeer()
			time.Sleep(25 * time.Millisecond)
			assert.Empty(t, mp.GetLastMessage())
		})
	})

}

/*

function: checkOwnership
purpose: checks whether a peer owns a given game

? Test cases

preamble
	#1 => peer not started validation => validate

success
	#1 => peer owns game
	#2 => peer doesn't own game
	#3 => unkown & check blockchain

failure
	#1 => peer doesn't own game

*/

func TestCheckOwnership(t *testing.T) {
	mp, err := testutil.StartMockPeer(7887, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { mp.Close() })
	time.Sleep(50 * time.Millisecond)

	client := Peer().server.Clients()[0]
	pd := &peerData{
		Peer:      client,
		Library:   make(map[[32]byte]ownership),
		Validator: ethereum.GenerateAddressValidation(),
	}

	game := sha256.Sum256([]byte("hello"))

	smoke := t.Run("preamble", func(t *testing.T) {
		t.Run("peer not validated", func(t *testing.T) {
			pd.Validator.SetValid(false)
			t.Cleanup(func() {
				pd.Validator.SetValid(true)
			})

			res, err := pd.checkOwnership(game)
			assert.Nil(t, err)
			assert.False(t, res)
		})
	})

	if !smoke {
		t.Fatal("failed peer validation")
	}
	time.Sleep(50 * time.Millisecond)

	t.Run("success", func(t *testing.T) {
		t.Run("game owned", func(t *testing.T) {
			pd.Library[game] = owned
			t.Cleanup(func() { pd.Library[game] = unknown })

			res, err := pd.checkOwnership(game)
			assert.Nil(t, err)
			assert.True(t, res)
		})

		t.Run("game not owned", func(t *testing.T) {
			pd.Library[game] = notOwned
			t.Cleanup(func() { pd.Library[game] = unknown })

			res, err := pd.checkOwnership(game)
			assert.Nil(t, err)
			assert.False(t, res)
		})
	})
}
