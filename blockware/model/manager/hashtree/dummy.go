package hash

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

dummies are empty files that are act as placeholders whilst data is
being downloaded. It means we can easily insert shards into storage
without having to worry about what position the shard is in the file.

One advantage is that if we cannot generate all the dummy files then
the user cannot have enough storage.

*/

func (ht *HashTree) CreateDummyFiles(rootDir, title string, onCreate func(string, *HashTreeFile)) error {
	err := os.Mkdir(filepath.Join(rootDir, title), 0655)
	if err != nil && !os.IsExist(err) {
		return err
	}

	files := ht.ListFiles()
	var wg sync.WaitGroup
	wg.Add(len(files))

	toCreate := make(chan *HashTreeFile, 5)
	go func() {
		for _, f := range files {
			toCreate <- f
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for f := range toCreate {
				fileLocation := filepath.Join(rootDir, title, f.AbsoluteFilename)
				err := setupFile(fileLocation, ht.ShardSize, len(f.Hashes), f.Size)
				if err != nil {
					util.Logger.Errorf("error creating %s: %s", fileLocation, err)
				}

				onCreate(fileLocation, f)
				wg.Done()
			}
		}()
	}

	wg.Wait()
	close(toCreate)
	return nil

}

/*
Create a dummy file
A dummy file with N shards of size B will be filled with N*B NULL bytes
to be filled in later
*/
func setupFile(filename string, shardSize uint, shardCount int, fileSize int) error {
	util.Logger.Debugf("Creating dummy file %s", filename)

	err := os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	emptyBuffer := make([]byte, shardSize)
	writer := bufio.NewWriter(file)

	for i := 0; i < shardCount; i++ {
		if i == shardCount-1 && fileSize%int(shardSize) != 0 {
			_, err = writer.Write(emptyBuffer[:fileSize%int(shardSize)])
		} else {
			_, err = writer.Write(emptyBuffer)
		}

		if err != nil {
			return err
		}

		writer.Flush()

	}

	util.Logger.Debugf("dummy file %s created", filename)
	return nil
}

// InsertData Insert a shard of data into a given dummy file
func InsertData(filename string, shardSize, offset uint, data []byte) error {
	if len(data) != int(shardSize) {
		return errors.New("data should be the same length as the byte size")
	}

	file, err := os.OpenFile(filename, os.O_RDWR, 0755)
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
