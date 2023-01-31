package games

import (
	"bytes"
	"crypto/sha256"
	"testing"
	"time"
)

// Create Game

func TestCreateGameSuccess(t *testing.T) {
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

}

func TestCreateGameInvalidRootDir(t *testing.T) {
	_, err := CreateGame("test-game", "1.1.1", time.Now().String(), "tcs1g20.com", "./fake/root/dir", 64)
	if err == nil {
		t.Errorf("Fake directory not detected")
	}
}

func TestCreateGameInvalidArguments(t *testing.T) {
	_, err := CreateGame("test-game", "-@123", time.Now().String(), "tcs1g20.com", ".", 64)
	if err == nil {
		t.Errorf("Invalid version number accepted")
	}

	_, err = CreateGame("test-game", "1.1.2", "not a real time", "tcs1g20.com", ".", 64)
	if err == nil {
		t.Errorf("Invalid datetime accepted")
	}

	_, err = CreateGame("test-game", "1.1.2", time.Now().String(), "not.real.domain.t02smith.com", ".", 64)
	if err == nil {
		t.Errorf("Invalid domain accepted")
	}

}

func TestSerialise(t *testing.T) {
	gamesTestSetup()
	defer gamesTestTeardown()

	g := &Game{
		Title:       "Test Game",
		Version:     "1.0.2",
		ReleaseDate: time.Now().String(),
		Developer:   "tcs1g20",
		RootHash:    []byte("test"),
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

	// compare original and deserialised
	if g.Title != d.Title ||
		g.Version != d.Version ||
		g.ReleaseDate != d.ReleaseDate ||
		g.Developer != d.Developer ||
		!bytes.Equal(g.RootHash, d.RootHash) {

		t.Error("Deserialised game not identical to original")
	}
}

func TestFetchShard(t *testing.T) {

	// load game made from our test directory
	gs, err := LoadGames("../../test/data/.toolkit")
	if err != nil {
		t.Errorf("games failed to load: %s", err)
		return
	}

	if len(gs) == 0 {
		t.Errorf("no games found. Make sure games_test:setupTestGame has been run")
		return
	}

	g := gs[0]
	err = g.ReadHashData()
	if err != nil {
		t.Errorf("Error reading hash data from game %s", err)
		return
	}

	// fetch shard from storage and compare with expected
	hash := g.data.RootDir.Files["architecture-diagram.png"].Hashes[0]
	foundShard, err := g.FetchShard(hash)
	if err != nil {
		t.Errorf("Error fetching shard: %s", err)
	}

	foundShardHash := sha256.Sum256(foundShard)
	if !bytes.Equal(hash[:], foundShardHash[:]) {
		t.Errorf("Incorrect shard found")
		return
	}

}
