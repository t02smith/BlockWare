package peer

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	testenv "github.com/t02smith/part-iii-project/toolkit/test/testutil/env"
)

/*

function: fetchBlockFromLibrary
purpose: fetch a block from storage given its game and hash

? Test cases
success
	| #1 => Block found and returned

failure
	| unexpected arguments
			| #1 => game not owned by user
			| #2 => shard not in game
			| #3 => user is downloading game and doesn't have block

*/

func TestFetchBlockFromLibrary(t *testing.T) {
	game := testenv.CreateTestGame(t, "../../../")
	err := Peer().library.AddOwnedGame(game)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		Peer().library.ClearOwnedGames()
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("unexpected arguments", func(t *testing.T) {
			t.Run("game not owned by user", func(t *testing.T) {
				Peer().library.ClearOwnedGames()
				t.Cleanup(func() {
					err := Peer().library.AddOwnedGame(game)
					if err != nil {
						t.Fatal(err)
					}
				})

				_, err := fetchBlockFromLibrary(game.RootHash, sha256.Sum256([]byte("asdas")))
				assert.NotNil(t, err, err)
				assert.Equal(t, "game not in library", err.Error(), "incorrect err message")
			})

			t.Run("shard not found in game", func(t *testing.T) {
				_, err := fetchBlockFromLibrary(game.RootHash, sha256.Sum256([]byte("asdas")))
				assert.NotNil(t, err, err)
				assert.Equal(t, "block not found", err.Error(), "incorrect err message")
			})
		})
	})

	t.Run("success", func(t *testing.T) {
		t.Run("block found", func(t *testing.T) {
			ht, err := game.GetData()
			if err != nil {
				t.Fatal(err)
			}

			hash := ht.RootDir.Files["architecture-diagram.png"].Hashes[1]
			data, err := fetchBlockFromLibrary(game.RootHash, hash)
			assert.Nil(t, err, err)

			dataHash := sha256.Sum256(data)
			assert.Equal(t, dataHash[:], hash[:], "incorrect block fetched")
		})
	})
}

/*

function: onMessage
purpose: handle incoming messages

? test cases
arguments:
	| #1 a message ends with a carriage return
	| #2 error message received

failure:
	| #1 unrecognised message
*/

func TestOnMessage(t *testing.T) {
	mp, tcp := createMockPeer(t)

	t.Run("arguments", func(t *testing.T) {
		t.Run("carriage return", func(t *testing.T) {
			onMessage(strings.Split("LIBRARY\r", ";"), tcp)
			time.Sleep(25 * time.Millisecond)

			msg := mp.GetLastMessage()
			assert.Equal(t, "GAMES", msg[:5], "carriage return not resolved")
		})

		t.Run("error message", func(t *testing.T) {
			err := onMessage(strings.Split("ERROR;error message", ";"), tcp)
			assert.Nil(t, err)
		})
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("unrecognised message", func(t *testing.T) {
			err := onMessage(strings.Split("FAKE_MESSAGE;error message", ";"), tcp)
			assert.NotNil(t, err, "error expected")
			assert.Equal(t, fmt.Sprintf("unrecognised message: %s", "FAKE_MESSAGE;error message"), err.Error(), "incorrect err message")
		})
	})

}
