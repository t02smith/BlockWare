package lib

import (
	"fmt"
	"testing"
)

// NewHashTree

func TestNewHashTreeInvalidShardSize(t *testing.T) {
	_, err := NewHashTree(".", 0)
	if err == nil {
		t.Errorf("shard size of 0 should not be accepted")
	}
}

func TestNewHashTreeInvalidDirectory(t *testing.T) {
	_, err := NewHashTree("./test/data/fake", 1024)
	if err == nil {
		t.Errorf("invalid directories shouldn't be accepted")
	}
}

func TestNewHashTreeValid(t *testing.T) {
	tree, err := NewHashTree(".", 1024)
	if err != nil {
		t.Errorf("error creating hash tree")
	}

	if tree.RootDir.Dirname != "." && tree.RootDirLocation != "." {
		t.Errorf("root directory location not set")
	}

	if tree.ShardSize != 1024 {
		t.Errorf("shard size not set")
	}
}

// shardFile

func TestShardFileInvalidFile(t *testing.T) {
	ht, _ := NewHashTree("../test/data", 4)
	_, err := ht.shardFile("../test/data", "fake.txt")
	if err == nil {
		t.Errorf("file doesn't exist and this should throw an error")
	}
}

func TestShardFileCorrect(t *testing.T) {
	ht, _ := NewHashTree("../test/data", 11)
	f, err := ht.shardFile("../test/data/", "test.txt")

	if err != nil {
		t.Errorf("error sharding file %s", err)
	}

	if f.Filename != "test.txt" {
		t.Errorf("incorrect filename stored")
	}

	if fmt.Sprintf("%x", f.Hashes[0]) != "12998c017066eb0d2a70b94e6ed3192985855ce390f321bbdb832022888bd251" {
		t.Errorf("Incorrect hash")
	}
}
