package ethereum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEthereum(t *testing.T) {
	lib := lib_instance
	testGame, err := fetchTestGame()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("fetch games => no games", func(t *testing.T) {
		t.Run("no games", func(t *testing.T) {
			gs, err := fetchGamesFromEthereum()
			if err != nil {
				t.Fatal(err)
			}

			if len(gs) != 0 {
				t.Fatal("No games expected")
			}
		})

		// upload game
		t.Run("upload game", func(t *testing.T) {
			err := uploadToEthereum(testGame)
			if err != nil {
				t.Fatal(err)
			}

			uploadedGame, err := lib.Games(nil, testGame.RootHash)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, testGame.Title, uploadedGame.Title, "game names not equal")
			assert.Equal(t, testGame.Version, uploadedGame.Version, "game versions not equal")
			assert.Equal(t, testGame.ReleaseDate, uploadedGame.ReleaseDate, "game release date not equal")
			assert.Equal(t, testGame.Developer, uploadedGame.Developer, "game dev not equal")
			assert.Equal(t, testGame.IPFSId, uploadedGame.IpfsAddress, "game IPFS data id not equal")

		})

		t.Run("one game", func(t *testing.T) {
			gs, err := fetchGamesFromEthereum()
			if err != nil {
				t.Fatal(err)
			}

			if len(gs) != 1 {
				t.Fatal("game not found")
			}

			// check contents
		})
	})

	t.Run("join existing network", func(t *testing.T) {
		addr := contract_address
		err := ConnectToLibraryInstance(addr)
		if err != nil {
			t.Fatal(err)
		}

		t.Run("Games exist", func(t *testing.T) {
			gs, err := fetchGamesFromEthereum()
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, 1, len(gs), "game found")

			g := gs[0]
			assert.Equal(t, testGame.Title, g.Title, "game names not equal")
			assert.Equal(t, testGame.Version, g.Version, "game versions not equal")
			assert.Equal(t, testGame.ReleaseDate, g.ReleaseDate, "game release date not equal")
			assert.Equal(t, testGame.Developer, g.Developer, "game dev not equal")
			assert.Equal(t, testGame.IPFSId, g.IPFSId, "game IPFS data id not equal")

		})
	})

}
