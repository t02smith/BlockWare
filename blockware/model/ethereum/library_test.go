package ethereum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
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
		err := ConnectToLibraryInstance(addr, testutil.Accounts[0][1])
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

	// TODO have other user upload game to blockchain

	t.Run("Purchase", func(t *testing.T) {
		err := uploadToEthereum(testGame)
		if err != nil {
			t.Fatal(err)
		}

		t.Run("game already owned", func(t *testing.T) {
			err := Purchase(net.Peer().Library(), testGame.RootHash)
			if err.Error() != fmt.Sprintf("game %x already purchased", testGame.RootHash) {
				t.Error("Already owned game not detected")
			}
		})

		t.Run("success", func(t *testing.T) {
			t.Skip()
			lib := games.NewLibrary()
			lib.SetBlockchainGame(testGame.RootHash, testGame)

			err = ConnectToLibraryInstance(contract_address, testutil.Accounts[2][1])
			if err != nil {
				t.Fatal(err)
			}

			err := Purchase(lib, testGame.RootHash)
			if err != nil {
				t.Fatal(err)
			}

			if lib.GetOwnedGame(testGame.RootHash) == nil {
				t.Fatal("Game not added to library")
			}
		})

	})

}
