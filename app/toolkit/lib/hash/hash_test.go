package hash

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/lib"
)

func TestMain(m *testing.M) {
	viper.Set("meta.hashes.workerCount", 5)
	lib.InitLogger()

	code := m.Run()
	os.Exit(code)
}

// NewHashTree

func TestNewHashTree(t *testing.T) {

	t.Run("illegal arguments", func(t *testing.T) {
		t.Run("invalid shard size", func(t *testing.T) {
			_, err := NewHashTree(".", 0)
			if err == nil {
				t.Errorf("shard size of 0 should not be accepted")
			}
		})

		t.Run("invalid directory", func(t *testing.T) {
			_, err := NewHashTree("./test/data/fake", 1024)
			if err == nil {
				t.Errorf("invalid directories shouldn't be accepted")
			}
		})
	})

	t.Run("correct input", func(t *testing.T) {
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
	})

}

// shardFile

func TestShardFile(t *testing.T) {

	t.Run("illegal arguments", func(t *testing.T) {

		t.Run("invalid file", func(t *testing.T) {
			htf := &HashTreeFile{
				Filename:         "not-real.jpg",
				AbsoluteFilename: "./fake/path/do/no/replicate",
			}

			err := htf.shardFile(1024)
			if err == nil {
				t.Errorf("file doesn't exist and this should throw an error")
			}
		})

	})

	t.Run("valid input", func(t *testing.T) {
		htf := &HashTreeFile{
			Filename:         "test.txt",
			AbsoluteFilename: "../../test/data/testdir/test.txt",
		}

		err := htf.shardFile(11)
		if err != nil {
			t.Errorf("error sharding file %s", err)
		}

		if fmt.Sprintf("%x", htf.Hashes[0]) != "12998c017066eb0d2a70b94e6ed3192985855ce390f321bbdb832022888bd251" {
			t.Errorf("Incorrect hash")
		}
	})

}

// hashDir

func TestHash(t *testing.T) {

	t.Run("valid input", func(t *testing.T) {
		ht, err := NewHashTree("../../test/data/testdir", 1024)
		if err != nil {
			t.Fatalf("NewHashTree failed on valid input")
		}

		err = ht.Hash()
		if err != nil {
			t.Fatalf("error hashing directory %s", err)
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

		if len(rdir.Subdirs["subdir"].Files) != 1 && rdir.Subdirs["subdir"].Files["chip8.c"].Filename != "chip8.c" {
			t.Errorf("incorrect number of files or file in subdir")
		}
	})

	t.Run("vary worker count", func(t *testing.T) {
		ht, err := NewHashTree("../../test/data/testdir", 1024)
		if err != nil {
			t.Fatalf("NewHashTree failed on valid input")
		}

		for _, i := range []int{0, 1, 4, 10} {
			t.Run(fmt.Sprintf("%d workers", i), func(t *testing.T) {
				viper.Set("meta.hashes.workercount", i)
				err := ht.Hash()
				if err != nil {
					t.Fatalf("Failed with %d workers", i)
				}
			})
		}
	})

}

// verifyTree

func TestVerifyDir(t *testing.T) {

	ht1, err := NewHashTree("../../test/data/testdir", 256)
	if err != nil {
		t.Fatalf("error creating new hash tree %s", err)
	}

	err = ht1.Hash()
	if err != nil {
		t.Fatalf("error hashing hash tree %s", err)
	}

	t.Run("compare to same directory", func(t *testing.T) {
		res, err := ht1.VerifyTree(&VerifyHashTreeConfig{false, false, 0}, "../../test/data/testdir")
		if err != nil {
			t.Fatalf("error verifying tree %s", err)
		}

		if !res {
			t.Errorf("result should be correct => is false")
		}
	})

	t.Run("compare to different directory", func(t *testing.T) {
		res, err := ht1.VerifyTree(&VerifyHashTreeConfig{false, false, 0}, "../../test")
		if err != nil {
			t.Errorf("error verifying tree %s", err)
		}

		if res {
			t.Errorf("result should be incorrect => is true")
		}
	})

	t.Run("test config options", func(t *testing.T) {
		// TODO
	})

}
