package peer

import (
	"crypto/sha256"
	"testing"

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
