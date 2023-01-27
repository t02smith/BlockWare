package games

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/test/util"
)

func setupTestDownload() (*Download, error) {
	util.ClearTmp("../../")
	util.SetupTmp("../../")

	games, err := LoadGames("../../test/data/.toolkit")
	if err != nil {
		return nil, err
	}

	if len(games) == 0 {
		return nil, errors.New("No games present in the test folder")
	}

	g := games[0]
	err = g.ReadHashData()
	if err != nil {
		return nil, err
	}

	d, err := SetupDownload(g)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func TestSetupDownload(t *testing.T) {
	_, err := setupTestDownload()
	if err != nil {
		t.Error(err)
	}
}

func TestSerialize(t *testing.T) {
	d, err := setupTestDownload()
	if err != nil {
		t.Error(err)
		return
	}

	err = d.Serialise()
	if err != nil {
		t.Errorf("Error serialising download: %s", err)
		return
	}

	f, err := os.Stat(filepath.Join(viper.GetString("games.tracker.directory"), fmt.Sprintf("%x", d.GameRootHash)))
	if err != nil {
		t.Error(err)
		return
	}

	if f.Size() == 0 {
		t.Error("No contents stored in tracker file")
		return
	}
}

func TestDeserialiseDownload(t *testing.T) {
	d, err := setupTestDownload()
	if err != nil {
		t.Error(err)
	}

	err = d.Serialise()
	if err != nil {
		t.Errorf("Error serialising download: %s", err)
		return
	}

	d2, err := DeserializeDownload(d.GameRootHash)
	if err != nil {
		t.Errorf("Error deserializing download: %s", err)
		return
	}

	// compare downloads
	if !bytes.Equal(d.GameRootHash, d2.GameRootHash) {
		t.Error("Games not the same")
		return
	}
}
