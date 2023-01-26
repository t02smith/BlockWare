package games

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	viper.Set("meta.hashes.workerCount", 5)
	viper.Set("meta.hashes.directory", "../../test/data/.toolkit/hashes")
	viper.Set("games.installFolder", "../../test/data/tmp")

	old := verifyDomain
	mockVerifyDomain = func(domain string) (bool, error) {
		return true, nil
	}

	code := m.Run()

	mockVerifyDomain = old
	os.Exit(code)
}

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

	deserialised, err := DeserialiseGame(serialised)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(deserialised)
}
