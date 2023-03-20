package testenv

import (
	"math/big"
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
)

// creates a test game and stores the data in the tmp folder
func CreateTestGame(t *testing.T, toRoot string) *games.Game {
	t.Helper()

	datetime := time.Date(2002, 1, 10, 0, 0, 0, 0, time.UTC).String()
	game, err := games.CreateGame(games.NewGame{
		Title:       "toolkit",
		Version:     "1.0.4",
		ReleaseDate: datetime,
		Developer:   "google.com",
		RootDir:     filepath.Join(toRoot, "./test/data/testdir"),
		Price:       big.NewInt(0),
		ShardSize:   256,
		AssetsDir:   filepath.Join(toRoot, "./test/data/assets"),
	}, nil)

	if err != nil {
		t.Fatal(err)
	}

	err = games.OutputAllGameDataToFile(game)
	if err != nil {
		t.Fatal(err)
	}

	return game
}

// setup a test download
func SetupTestDownload(t *testing.T, game *games.Game, toRoot string) {
	t.Helper()

	oldFolder := viper.GetString("games.installFolder")
	viper.Set("game.installFolder", filepath.Join(toRoot, "./test/data/tmp"))

	err := game.SetupDownload()
	if err != nil {
		t.Fatalf("Error setting up download: %s", err)
	}

	t.Cleanup(func() {
		err = game.CancelDownload()
		if err != nil {
			t.Fatalf("Error cancelling download: %s", err)
		}

		viper.Set("games.installFolder", oldFolder)
	})
}
