package games

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestSetupDownload(t *testing.T) {
	gamesTestSetup()
	defer gamesTestTeardown()

	_, err := setupTestDownload()
	if err != nil {
		t.Error(err)
	}
}

func TestSerialization(t *testing.T) {
	t.Cleanup(gamesTestTeardown)

	t.Run("serialize", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			gamesTestSetup()
			g, err := setupTestDownload()
			if err != nil {
				t.Fatal(err)
			}

			d := g.download

			err = d.Serialise(fmt.Sprintf("%x", g.RootHash))
			if err != nil {
				t.Fatalf("Error serialising download: %s", err)
			}

			f, err := os.Stat(filepath.Join(viper.GetString("games.tracker.directory"), fmt.Sprintf("%x", g.RootHash)))
			if err != nil {
				t.Fatal(err)
			}

			if f.Size() == 0 {
				t.Fatalf("No contents stored in tracker file")
			}
		})

	})

	t.Run("deserialize", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			gamesTestSetup()

			g, err := setupTestDownload()
			if err != nil {
				t.Fatal(err)
			}

			d := g.download

			err = d.Serialise(fmt.Sprintf("%x", g.RootHash))
			if err != nil {
				t.Fatalf("Error serialising download: %s", err)
			}

			_, err = DeserializeDownload(g.RootHash)
			if err != nil {
				t.Fatalf("Error deserializing download: %s", err)
			}

			// // compare downloads
			// if !bytes.Equal(d.GameRootHash[:], d2.GameRootHash[:]) {
			// 	t.Fatal("Games not the same")
			// }
		})
	})

}

func TestFindBlock(t *testing.T) {

	t.Cleanup(gamesTestTeardown)

	t.Run("success", func(t *testing.T) {
		gamesTestSetup()
		// TODO
	})

}

// func TestFindBlock2(t *testing.T) {
// 	t.SkipNow()
// 	gamesTestSetup()
// 	defer gamesTestTeardown()

// 	d, _, err := setupTestDownload()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	missingHash := sha256.Sum256([]byte("hello"))

// 	// request with no peers connected
// 	err = d.FindBlock(missingHash)
// 	if err == nil {
// 		t.Error("This function should error if no peers are connected")
// 	}

// 	// * connect a new peer

// 	// request an unknown hash
// 	err = d.FindBlock(missingHash)
// 	if err == nil {
// 		t.Error("This block doesn't exist and should error")
// 	}

// }
