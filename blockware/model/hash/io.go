package hash

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path/filepath"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

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

// DOWNLOADING A FILE

func (t *HashTree) CreateDummyFiles(rootDir, title string, onCreate func(string, *HashTreeFile)) error {

	err := os.Mkdir(filepath.Join(rootDir, title), 0777)
	if err != nil && !os.IsExist(err) {
		return err
	}

	err = t.createDummyFilesFromDirectory(t.RootDir, filepath.Join(rootDir, title), onCreate)
	if err != nil {
		return err
	}

	return nil

}

func (t *HashTree) createDummyFilesFromDirectory(dir *HashTreeDir, path string, onCreate func(string, *HashTreeFile)) error {
	util.Logger.Infof("Generating files for folder %s", filepath.Join(path, dir.Dirname))

	if len(dir.Dirname) > 0 {
		err := os.Mkdir(filepath.Join(path, dir.Dirname), 0777)
		if err != nil {
			return err
		}
	}

	// generate files
	for _, f := range dir.Files {
		fileLocation := filepath.Join(path, dir.Dirname, f.Filename)
		util.Logger.Infof("Creating dummy for %s", fileLocation)
		err := setupFile(fileLocation, t.ShardSize, uint(len(f.Hashes)))
		if err != nil {
			return err
		}
		onCreate(fileLocation, f)
	}

	// generate subdirs
	newPath := filepath.Join(path, dir.Dirname)
	for _, d := range dir.Subdirs {
		err := t.createDummyFilesFromDirectory(d, newPath, onCreate)
		if err != nil {
			return err
		}
	}

	return nil
}

func setupFile(filename string, shardSize, shardCount uint) error {
	util.Logger.Infof("Creating empty file %s", filename)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	emptyBuffer := make([]byte, shardSize)
	writer := bufio.NewWriter(file)

	for i := 0; i < int(shardCount); i++ {
		_, err := writer.Write(emptyBuffer)
		writer.Flush()
		if err != nil {
			return err
		}
	}

	util.Logger.Infof("%s created", filename)
	return nil
}

func InsertData(filename string, shardSize, offset uint, data []byte) error {
	if len(data) != int(shardSize) {
		return errors.New("data should be the same length as the byte size")
	}

	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(int64(offset*shardSize), 0)
	if err != nil {
		return err
	}

	util.Logger.Infof("Writing shard to %s:%d", filename, offset)
	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		util.Logger.Error(err)
		return err
	}

	err = writer.Flush()
	if err != nil {
		util.Logger.Error(err)
		return err
	}

	util.Logger.Infof("shard %s:%d written successfully", filename, offset)
	return nil
}
