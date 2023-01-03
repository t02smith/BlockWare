package lib

import (
	"bytes"
	"crypto/sha256"
	"path/filepath"
	"testing"
)

func produceTestHashTree() (*hashTree, error) {
	ht, err := NewHashTree("../test/data/testdir", 1024)
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

	if filename != filepath.Join("../test/data/testdir", "architecture-diagram.png") {
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
