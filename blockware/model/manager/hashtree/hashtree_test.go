package hash

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func createTestHashTree(t *testing.T) *HashTree {
	ht1, err := NewHashTree("../../../test/data/testdir", 256, nil)
	if err != nil {
		t.Fatalf("error creating new hash tree %s", err)
	}

	err = ht1.Hash()
	if err != nil {
		t.Fatalf("error hashing hash tree %s", err)
	}

	return ht1
}

// NewHashTree

func TestNewHashTree(t *testing.T) {

	t.Run("illegal arguments", func(t *testing.T) {
		t.Run("invalid shard size", func(t *testing.T) {
			_, err := NewHashTree(".", 0, nil)
			if err == nil {
				t.Errorf("shard size of 0 should not be accepted")
			}
		})

		t.Run("invalid directory", func(t *testing.T) {
			_, err := NewHashTree("./test/data/fake", 1024, nil)
			if err == nil {
				t.Errorf("invalid directories shouldn't be accepted")
			}
		})
	})

	t.Run("correct input", func(t *testing.T) {
		tree, err := NewHashTree(".", 1024, nil)
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
		empty := [256]byte{}

		table := []struct {
			name     string
			htf      *HashTreeFile
			expected [32]byte
		}{
			{
				"normal file",
				&HashTreeFile{
					Filename:         "test.txt",
					AbsoluteFilename: "../../../test/data/testdir/test.txt",
				},
				[32]byte{2, 96, 221, 220, 74, 251, 255, 33, 146, 79, 118, 176, 189, 239, 210, 55, 87, 12, 16, 175, 90, 116, 199, 69, 81, 247, 149, 65, 223, 135, 190, 163},
			},
			{
				"empty file",
				&HashTreeFile{
					Filename:         "EMPTY",
					AbsoluteFilename: "../../../test/data/testfiles/EMPTY",
				},
				sha256.Sum256(empty[:]),
			},
		}

		for _, x := range table {
			t.Run(x.name, func(t *testing.T) {
				err := x.htf.shardFile(256)
				if err != nil {
					t.Errorf("error sharding file %s", err)
				}

				if !bytes.Equal(x.htf.Hashes[0][:], x.expected[:]) {
					t.Errorf("Incorrect hash for %s. expected %x, got %x", x.name, x.expected, x.htf.Hashes[0])
				}
			})
		}

	})
}

// hashDir

func TestHash(t *testing.T) {

	t.Run("valid input", func(t *testing.T) {
		ht, err := NewHashTree("../../../test/data/testdir", 1024, nil)
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
		ht, err := NewHashTree("../../../test/data/testdir", 1024, nil)
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
	ht1 := createTestHashTree(t)

	t.Run("compare to same directory", func(t *testing.T) {
		res, err := ht1.VerifyTree(&VerifyHashTreeConfig{false, false, 0}, "../../../test/data/testdir")
		if err != nil {
			t.Fatalf("error verifying tree %s", err)
		}

		if !res {
			t.Errorf("result should be correct => is false")
		}
	})

	t.Run("compare to different directory", func(t *testing.T) {
		res, err := ht1.VerifyTree(&VerifyHashTreeConfig{false, false, 0}, "../../../test")
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

/*

function: HashTree.Equals
purpose: compare two hash trees

? Test cases
success
	#1 => equal
	#2 => not equal

*/

func TestHashTreeEquals(t *testing.T) {
	ht := createTestHashTree(t)

	t.Run("success", func(t *testing.T) {
		t.Run("equal", func(t *testing.T) {
			assert.True(t, ht.Equals(ht))
		})

		t.Run("not equal", func(t *testing.T) {
			ht2 := *ht
			ht2.ShardSize = 104
			assert.False(t, ht.Equals(&ht2))
		})
	})
}

/*

function: HashTree.OutputToFile
purpose: serialise and output a hash tree to a file

? Test cases
success
	#1 => file successfully outputted

failure
	#1 => directory doesn't exist

*/

func TestHashTreeOutputToFile(t *testing.T) {
	ht := createTestHashTree(t)

	t.Run("success", func(t *testing.T) {
		err := ht.OutputToFile("../../../test/data/tmp/output-test")
		assert.Nil(t, err)
		t.Cleanup(func() {
			os.Remove("../../../test/data/tmp/output-test")
		})

		stat, err := os.Stat("../../../test/data/tmp/output-test")
		assert.Nil(t, err)
		assert.NotZero(t, stat.Size())

		fromFile, err := ReadHashTreeFromFile("../../../test/data/tmp/output-test")
		assert.Nil(t, err)
		assert.True(t, ht.Equals(fromFile))
	})

	t.Run("failure", func(t *testing.T) {
		err := ht.OutputToFile("./fake/directory/pls/fail")
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, os.ErrNotExist)
	})
}

/*

function: ReadHashTreeFromFile
purpose: deserialise a hash tree from a file

? Test cases
success
	#1 => correct hash tree parsed

failure
	#1 => file not found

*/

func TestReadHashTreeFromFile(t *testing.T) {
	ht := createTestHashTree(t)

	t.Run("success", func(t *testing.T) {
		if err := ht.OutputToFile("../../../test/data/tmp/test-readhtfromfile"); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			os.Remove("../../../test/data/tmp/test-readhtfromfile")
		})

		ht2, err := ReadHashTreeFromFile("../../../test/data/tmp/test-readhtfromfile")
		assert.Nil(t, err)
		assert.True(t, ht.Equals(ht2))
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("file not found", func(t *testing.T) {
			ht2, err := ReadHashTreeFromFile("./fake/test/file")
			assert.Nil(t, ht2)
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, os.ErrNotExist)
		})
	})
}

/*

function: HashTreeDir.Equals
purpose: compare two hash tree directories

? Test cases
success
	#1 => equal
	not equal
		#1 => different files
		#2 => different dirs

*/

func TestHashTreeDirEquals(t *testing.T) {
	htd := createTestHashTree(t).RootDir

	t.Run("success", func(t *testing.T) {
		t.Run("equal", func(t *testing.T) {
			assert.True(t, htd.Equals(htd))
		})

		t.Run("not equal", func(t *testing.T) {
			t.Run("quick checks", func(t *testing.T) {
				t.Run("root hashes", func(t *testing.T) {
					htd2 := *htd
					htd2.RootHash = sha256.Sum256([]byte("hello there world"))
					assert.False(t, htd.Equals(&htd2))
				})

				t.Run("file count", func(t *testing.T) {
					htd2 := *htd
					htd2.Files = make(map[string]*HashTreeFile)
					assert.False(t, htd.Equals(&htd2))
				})

				t.Run("subdir count", func(t *testing.T) {
					htd2 := *htd
					htd2.Subdirs = make(map[string]*HashTreeDir)
					assert.False(t, htd.Equals(&htd2))
				})
			})

			t.Run("different files", func(t *testing.T) {
				htd2 := *htd
				htd2.Files = make(map[string]*HashTreeFile)
				htd2.Files["test.txt"] = &HashTreeFile{
					RootHash: sha256.Sum256([]byte("test")),
				}
				assert.False(t, htd.Equals(&htd2))
			})

			t.Run("different subdirs", func(t *testing.T) {
				htd2 := *htd
				htd2.Subdirs = make(map[string]*HashTreeDir)
				htd2.Subdirs["subdir"] = &HashTreeDir{
					RootHash: sha256.Sum256([]byte("test")),
				}
				assert.False(t, htd.Equals(&htd2))
			})
		})
	})

}
