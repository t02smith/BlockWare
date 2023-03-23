package peer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
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
