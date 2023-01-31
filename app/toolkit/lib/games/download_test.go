package games

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestSetupDownload(t *testing.T) {
	gamesTestSetup()
	defer gamesTestTeardown()

	_, _, err := setupTestDownload()
	if err != nil {
		t.Error(err)
	}
}

func TestSerialize(t *testing.T) {
	gamesTestSetup()
	defer gamesTestTeardown()

	d, _, err := setupTestDownload()
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
	gamesTestSetup()
	defer gamesTestTeardown()

	d, _, err := setupTestDownload()
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
	if !bytes.Equal(d.GameRootHash[:], d2.GameRootHash[:]) {
		t.Error("Games not the same")
		return
	}
}

func TestFindBlock(t *testing.T) {
	t.SkipNow()
	gamesTestSetup()
	defer gamesTestTeardown()

	d, _, err := setupTestDownload()
	if err != nil {
		t.Error(err)
	}

	missingHash := sha256.Sum256([]byte("hello"))

	// request with no peers connected
	err = d.FindBlock(missingHash)
	if err == nil {
		t.Error("This function should error if no peers are connected")
	}

	// * connect a new peer

	// request an unknown hash
	err = d.FindBlock(missingHash)
	if err == nil {
		t.Error("This block doesn't exist and should error")
	}

}
