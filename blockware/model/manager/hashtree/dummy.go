package hash

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

dummies are empty files that are act as placeholders whilst data is
being downloaded. It means we can easily insert shards into storage
without having to worry about what position the shard is in the file.

One advantage is that if we cannot generate all the dummy files then
the user cannot have enough storage.

*/

// CreateDummyFiles /*
func (ht *HashTree) CreateDummyFiles(rootDir, title string, onCreate func(string, *HashTreeFile)) error {
	err := os.Mkdir(filepath.Join(rootDir, title), 0777)
	if err != nil && !os.IsExist(err) {
		return err
	}

	err = ht.createDummyFilesFromDirectory(ht.RootDir, filepath.Join(rootDir, title), onCreate)
	if err != nil {
		return err
	}

	return nil

}

// traverse a directory and create any dummy files and recursively traverse sub-dirs
func (ht *HashTree) createDummyFilesFromDirectory(dir *HashTreeDir, path string, onCreate func(string, *HashTreeFile)) error {
	util.Logger.Debugf("Generating files for folder %s", filepath.Join(path, dir.Dirname))

	if len(dir.Dirname) > 0 {
		err := os.Mkdir(filepath.Join(path, dir.Dirname), 0777)
		if err != nil {
			return err
		}
	}

	// generate files
	for _, f := range dir.Files {
		fileLocation := filepath.Join(path, dir.Dirname, f.Filename)
		err := setupFile(fileLocation, ht.ShardSize, uint(len(f.Hashes)))
		if err != nil {
			return err
		}
		onCreate(fileLocation, f)
	}

	// generate sub-dirs
	newPath := filepath.Join(path, dir.Dirname)
	for _, d := range dir.Subdirs {
		err := ht.createDummyFilesFromDirectory(d, newPath, onCreate)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
Create a dummy file
A dummy file with N shards of size B will be filled with N*B NULL bytes
to be filled in later
*/
func setupFile(filename string, shardSize, shardCount uint) error {
	util.Logger.Debugf("Creating dummy file %s", filename)
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

	util.Logger.Debugf("dummy file %s created", filename)
	return nil
}

// InsertData Insert a shard of data into a given dummy file
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

	util.Logger.Debugf("Writing shard to %s:%d", filename, offset)
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

	util.Logger.Debugf("shard %s:%d written successfully", filename, offset)
	return nil
}
