package ethereum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEthereum(t *testing.T) {
	t.Skip()
	lib := lib_instance

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

		testGame, err := fetchTestGame()
		if err != nil {
			t.Fatal(err)
		}

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

}
