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

	if tree.RootDir != nil {
		t.Errorf("by default the root directory should be nil")
	}

	if tree.ShardSize != 1024 {
		t.Errorf("shard size not set")
	}
}

// shardFile

func TestShardFileInvalidFile(t *testing.T) {
	ht, _ := NewHashTree("../test/data/testdir", 4)
	_, err := ht.shardFile("../test/data/testdir", "fake.txt")
	if err == nil {
		t.Errorf("file doesn't exist and this should throw an error")
	}
}

func TestShardFileCorrect(t *testing.T) {
	ht, _ := NewHashTree("../test/data/testdir", 11)
	f, err := ht.shardFile("../test/data/testdir", "test.txt")

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

// hashDir

func TestHashDirCorrectTree(t *testing.T) {
	ht, _ := NewHashTree("../test/data/testdir", 1024)
	err := ht.Hash()

	if err != nil {
		t.Errorf("error hashing directory %s", err)
	}

	rdir := ht.RootDir
	if rdir.Dirname != "" {
		t.Errorf("directory is root so should have an empty dirname")
	}

	if len(rdir.Files) != 2 {
		t.Errorf("incorrect number of files found in the root directory")
	}

	if len(rdir.Subdirs) != 1 {
		t.Errorf("incorrect number of subdirs found")
	}

	if len(rdir.Subdirs[0].Files) != 1 && rdir.Subdirs[0].Files[0].Filename != "chip8.c" {
		t.Errorf("incorrect number of files or file in subdir")
	}
}

// verifyTree

func TestVerifyTreeSameDir(t *testing.T) {
	ht1, _ := NewHashTree("../test/data/testdir", 256)
	ht1.Hash()

	res, err := ht1.VerifyTree(&VerifyHashTreeConfig{false, false, 0}, "../test/data/testdir")
	if err != nil {
		t.Errorf("error verifying tree %s", err)
	}

	if !res {
		t.Errorf("result should be correct => is false")
	}

}

func TestVerifyTreeDifferentDir(t *testing.T) {
	ht1, _ := NewHashTree("../test/data/testdir", 256)
	ht1.Hash()

	res, err := ht1.VerifyTree(&VerifyHashTreeConfig{false, false, 0}, "../test")
	if err != nil {
		t.Errorf("error verifying tree %s", err)
	}

	if res {
		t.Errorf("result should be incorrect => is true")
	}

}
