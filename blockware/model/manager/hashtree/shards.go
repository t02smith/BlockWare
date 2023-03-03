package hash

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

A shard is stored as a SHA256 hash of a sequential list of bytes.
Hashing allows us to easily generate a unique fingerprint of each
shard in a given file.

*/

// locate a given shard within a hash tree by looking through all files
func (t *HashTree) GetShard(hash [32]byte) (bool, []byte, error) {
	util.Logger.Infof("Looking for shard %x in %s", hash, t.RootDirLocation)
	found, _, location, offset := t.FindShard(hash)
	if !found {
		return false, nil, nil
	}

	util.Logger.Infof("Shard found at %s - piece %d", location, offset)
	data, err := t.readShard(filepath.Join(t.RootDirLocation, location), offset)
	if err != nil {
		return false, nil, err
	}

	return true, data, nil
}

// LOCATING SHARDS

// looks for a given shard in the entire hash tree
// return <hashFound> <file struct> <relative filename> <shard position in file>
func (t *HashTree) FindShard(hash [32]byte) (bool, *HashTreeFile, string, int) {
	return t.RootDir.findShard(hash)
}

// auxillary function for FindShard => looks in particular directory
func (td *HashTreeDir) findShard(hash [32]byte) (bool, *HashTreeFile, string, int) {
	for _, f := range td.Files {
		for i, h := range f.Hashes {
			if bytes.Equal(hash[:], h[:]) {
				return true, f, filepath.Join(td.Dirname, f.Filename), i
			}
		}
	}

	for _, sd := range td.Subdirs {
		found, htf, filename, offset := sd.findShard(hash)
		if found {
			return true, htf, filepath.Join(td.Dirname, filename), offset
		}

	}

	return false, nil, "", -1

}

// READING SHARD

// Reads a shard from a given file
func (t *HashTree) readShard(filename string, offset int) ([]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := make([]byte, t.ShardSize)
	reader := bufio.NewReader(file)

	_, err = file.Seek(int64(offset*int(t.ShardSize)), 0)
	if err != nil {
		return nil, err
	}

	_, err = reader.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil

}
