package io

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"os"
	"path/filepath"
	"testing"
)

func produceTestHashTree() (*hashTree, error) {
	ht, err := NewHashTree("../../test/data/testdir", 1024)
	if err != nil {
		return nil, err
	}

	err = ht.Hash()
	if err != nil {
		return nil, err
	}

	return ht, nil
}

func TestFindShardCorrect(t *testing.T) {
	ht, err := produceTestHashTree()
	if err != nil {
		t.Errorf("%s", err)
	}

	file := ht.RootDir.Files[0]
	found, filename, offset := ht.findShard(file.Hashes[0])

	if !found {
		t.Error("Existing shard not found")
	}

	if filename != filepath.Join("../../test/data/testdir", "architecture-diagram.png") {
		t.Errorf("Incorrect filepath returned %s", filename)
	}

	if offset != 0 {
		t.Error("Incorrect offset")
	}
}

func TestFindShardIncorrect(t *testing.T) {
	ht, err := produceTestHashTree()
	if err != nil {
		t.Errorf("%s", err)
	}

	found, _, _ := ht.findShard([32]byte{})
	if found {
		t.Errorf("Shard does not exist but states it is found")
	}
}

// Test getting a shard

func TestGetShardCorrect(t *testing.T) {
	ht, err := produceTestHashTree()
	if err != nil {
		t.Errorf("%s", err)
	}

	file := ht.RootDir.Files[0]
	data, err := ht.GetShard(file.Hashes[0])
	if err != nil {
		t.Errorf("%s", err)
	}

	hash := sha256.Sum256(data)
	if !bytes.Equal(file.Hashes[0][:], hash[:]) {
		t.Error("Incorrect shard fetched")
	}
}

func TestGetShardIncorrect(t *testing.T) {

}

// Creating empty file

func TestSetupFileSuccess(t *testing.T) {
	tmpFile := "../../test/data/tmp/skeleton.txt"
	err := setupFile(tmpFile, 64, 100)
	if err != nil {
		t.Errorf("%s", err)
	}

	file, err := os.Stat(tmpFile)
	if err != nil {
		t.Errorf("%s", err)
	}

	if file.Size() != 6400 {
		t.Errorf("Incorrect filesize.\nExpected: %d\nGot: %d", 6400, file.Size())
	}

	err = os.Remove(tmpFile)
	if err != nil {
		t.Errorf("Error removing tmp file")
	}
}

// Inserting shards

func TestInsertDataCorrect(t *testing.T) {
	// setup
	tmpFile := "../../test/data/tmp/skeleton.txt"
	err := setupFile(tmpFile, 64, 100)
	if err != nil {
		t.Errorf("%s", err)
	}

	newShard := make([]byte, 64)
	for i := range newShard {
		newShard[i] = 255
	}

	err = insertData("../../test/data/tmp/skeleton.txt", 64, 12, newShard)
	if err != nil {
		t.Errorf("%s", err)
	}

	// test change occurred
	file, _ := os.Open(tmpFile)
	defer file.Close()

	buffer := make([]byte, 64)
	reader := bufio.NewReader(file)

	file.Seek(64*12, 0)
	reader.Read(buffer)

	if !bytes.Equal(buffer, newShard) {
		t.Error("Shard not written correctly to file")
	}
}
