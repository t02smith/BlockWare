package games

import (
	"bytes"
	"crypto/sha256"
	"testing"
	"time"
)

func TestCreateGame(t *testing.T) {
	t.Run("illegal arguments", func(t *testing.T) {

		table := []struct {
			name        string
			title       string
			version     string
			releaseDate string
			developer   string
			rootDir     string
			shardSize   uint
		}{
			{"root directory", "test-game", "1.1.1", time.Now().String(), "tcs1g20.com", "./fake/root/dir", 64},
			{"release date", "test-game", "1.1.2", "not a real time", "tcs1g20.com", ".", 64},
			{"invalid domain", "test-game", "1.1.2", time.Now().String(), "not.real.domain.t02smith.com", ".", 64},
			{"invalid shard size", "test-game", "1.1.1", time.Now().String(), "google.com", ".", 0},
		}

		for _, x := range table {
			t.Run(x.name, func(t *testing.T) {
				_, err := CreateGame(x.name, x.version, x.releaseDate, x.developer, x.rootDir, x.shardSize)
				if err == nil {
					t.Fatalf("Failed to detect illegal argument: %s", x.name)
				}
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()

		game, err := CreateGame("test-game", "1.1.1", datetime, "google.com", "../../test/data/testdir", 64)
		if err != nil {
			t.Errorf("Error creating game: %s", err)
			return
		}

		if game.Title != "test-game" {
			t.Errorf("Incorrect game name")
		}

		if game.Developer != "google.com" {
			t.Errorf("Incorrect developer")
		}

		if game.Version != "1.1.1" {
			t.Errorf("Incorrect version")
		}

		if game.ReleaseDate != datetime {
			t.Errorf("Invalid release date")
		}

		if game.data == nil {
			t.Errorf("Invalid game data")
		}
	})
}

func TestSerialise(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		gamesTestSetup()
		defer gamesTestTeardown()

		var h [32]byte
		copy(h[:], []byte("test"))

		g := &Game{
			Title:       "Test Game",
			Version:     "1.0.2",
			ReleaseDate: time.Now().String(),
			Developer:   "tcs1g20",
			RootHash:    h,
		}

		serialised, err := g.Serialise()
		if err != nil {
			t.Error(err)
			return
		}

		d, err := DeserialiseGame(serialised)
		if err != nil {
			t.Error(err)
			return
		}

		if !g.Equals(d) {
			t.Fatalf("Deserialised game not identical to original")
		}
	})
}

func TestFetchShard(t *testing.T) {

	g, err := fetchTestGame()
	if err != nil {
		t.Fatalf("error fetching test game %s", err)
	}

	type FetchShardInput struct {
		subdirs  []string
		filename string
		hashNo   int
	}

	t.Run("success", func(t *testing.T) {

		table := []FetchShardInput{
			{[]string{}, "architecture-diagram.png", 0},
			{[]string{"subdir"}, "chip8.c", 2},
		}

		for _, x := range table {
			t.Run(x.filename, func(t *testing.T) {
				dir := g.data.RootDir
				for _, sd := range x.subdirs {
					dir = dir.Subdirs[sd]
				}

				hash := dir.Files[x.filename].Hashes[x.hashNo]
				found, data, err := g.FetchShard(hash)
				if err != nil {
					t.Errorf("Error fetching shard: %s", err)
				}

				if !found {
					t.Error("error finding shard")
				}

				foundShardHash := sha256.Sum256(data)
				if !bytes.Equal(hash[:], foundShardHash[:]) {
					t.Errorf("Incorrect shard found")
					return
				}
			})
		}
	})

	t.Run("failure", func(t *testing.T) {
		// TODO
	})

	gamesTestSetup()
}