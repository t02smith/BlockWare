package hash

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

// Describes a directory in terms of its files hashes
type HashTree struct {

	// the directory object for the root folder
	RootDir *HashTreeDir `json:"rootdir"`

	// what size should each shard of data be in bytes
	ShardSize uint `json:"shardsize"`

	// the local physical location of the root directory
	RootDirLocation string `json:"location"`

	// a channel for viewing the progress of a hash
	progress chan int
}

// Describes a directory with tracked files
type HashTreeDir struct {

	// the path relative to the root directory location
	Dirname string `json:"dirname"`

	// the SHA256 hash of the directories contents
	RootHash [32]byte `json:"roothash"`

	// all subdirectories recursively stored
	Subdirs map[string]*HashTreeDir `json:"subdirs"`

	// all files within this folder
	Files map[string]*HashTreeFile `json:"files"`
}

// Describes a singular tracked files
type HashTreeFile struct {
	Filename string `json:"filename"`

	AbsoluteFilename string

	// A list of SHA256 hashes that represent each shard of data in the file
	Hashes [][32]byte `json:"hashes"`

	// The hash produced by the Merkle Tree of each file
	RootHash [32]byte `json:"roothash"`
}

type VerifyHashTreeConfig struct {
	IgnoreNewFilesAndDirs     bool
	ContinueAfterError        bool
	continueAfterErrorCounter uint
}

// create a new hash tree of a directory
func NewHashTree(rootDir string, shardSize uint, progress chan int) (*HashTree, error) {
	if shardSize == 0 {
		return nil, errors.New("shard size must be greater than 0")
	}

	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		return nil, err
	}

	return &HashTree{
		RootDir:         nil,
		RootDirLocation: rootDir,
		ShardSize:       shardSize,
		progress:        progress,
	}, nil
}

func (ht1 *HashTree) Equals(ht2 *HashTree) bool {
	return ht1.ShardSize == ht2.ShardSize &&
		ht1.RootDir.Equals(ht2.RootDir)
}

// IO

// output a hash tree to a json file
func (ht *HashTree) OutputToFile(filename string) error {
	util.Logger.Infof("outputting to file %s\n", filename)
	e, err := json.Marshal(ht)
	if err != nil {
		log.Println(err)
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(string(e))
	writer.Flush()
	return nil
}

// read a hash tree from a json file
func ReadHashTreeFromFile(filename string) (*HashTree, error) {
	util.Logger.Infof("Attempting to read hash data from %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	ht := &HashTree{
		RootDirLocation: filename,
	}

	if err := json.Unmarshal(data, &ht); err != nil {
		return nil, err
	}

	util.Logger.Infof("Hash data read from %s successfully", filename)
	return ht, nil
}

// Hashing functions

func (ht *HashTree) Hash() error {
	startTime := time.Now()
	fileCount, err := ht.buildTree()
	if err != nil {
		return err
	}

	if ht.progress != nil {
		ht.progress <- fileCount

	}

	wc := viper.GetInt("meta.hashes.workercount")
	if wc <= 0 {
		wc = 1
	}

	wg, fileIn, errors := hasherPool(wc, fileCount, ht.ShardSize, ht.progress)
	err = ht.RootDir.shardData(fileIn, ht.ShardSize)
	if err != nil {
		util.Logger.Error(err)
	}

	wg.Wait()
	close(fileIn)
	close(errors)

	endTime := time.Now()
	util.Logger.Infof("Directory %s hashed in %dms", ht.RootDirLocation, endTime.Sub(startTime).Milliseconds())
	return nil
}

// Data collection

// build a directory tree and count the number of files
func (ht *HashTree) buildTree() (int, error) {
	util.Logger.Infof("Building tree of directory %s\n", ht.RootDirLocation)

	ht.RootDir = &HashTreeDir{
		Dirname: "",
		Subdirs: make(map[string]*HashTreeDir),
		Files:   make(map[string]*HashTreeFile),
	}

	return ht.RootDir.traverseDirectory(ht.RootDirLocation)
}

// perform a search to build a directory tree
func (htd *HashTreeDir) traverseDirectory(absolutePath string) (int, error) {
	util.Logger.Infof("Hashing directory %s\n", htd.Dirname)

	counter := 0
	dir, err := os.Open(filepath.Join(absolutePath, htd.Dirname))
	if err != nil {
		return counter, err
	}
	defer dir.Close()

	dirContents, err := dir.Readdirnames(0)
	if err != nil {
		return counter, err
	}

	// look in the current directory
	for _, name := range dirContents {
		f, err := os.Stat(filepath.Join(absolutePath, htd.Dirname, name))
		if err != nil {
			return counter, err
		}

		// the data object is a directory
		if f.IsDir() {
			htd.Subdirs[name] = &HashTreeDir{
				Dirname: name,
				Files:   make(map[string]*HashTreeFile),
				Subdirs: make(map[string]*HashTreeDir),
			}

			continue
		}

		// the data object is a file
		htd.Files[name] = &HashTreeFile{
			Filename:         name,
			AbsoluteFilename: filepath.Join(absolutePath, htd.Dirname, name),
		}
		counter++
	}

	// look in subdirectories
	for _, subdir := range htd.Subdirs {
		fileCount, err := subdir.traverseDirectory(filepath.Join(filepath.Join(absolutePath, htd.Dirname)))
		if err != nil {
			return counter, err
		}

		counter += fileCount
	}

	return counter, nil
}

// hashing

func (htd *HashTreeDir) shardData(fileIn chan *HashTreeFile, shardSize uint) error {
	for _, f := range htd.Files {
		fileIn <- f
	}

	for _, subdir := range htd.Subdirs {
		err := subdir.shardData(fileIn, shardSize)
		if err != nil {
			return err
		}
	}

	return nil
}

func (htf *HashTreeFile) shardFile(shardSize uint) error {
	htf.Hashes = [][32]byte{}

	stat, err := os.Stat(htf.AbsoluteFilename)
	if err != nil {
		return err
	}

	if stat.Size() == 0 {
		htf.Hashes = append(htf.Hashes, [32]byte{})
		htf.RootHash = CalculateRootHash(htf.Hashes)
		return nil
	}

	file, err := os.Open(htf.AbsoluteFilename)
	if err != nil {
		return err
	}

	defer file.Close()

	buffer := make([]byte, shardSize)
	reader := bufio.NewReader(file)

	for {
		_, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		hash := sha256.Sum256(buffer)
		htf.Hashes = append(htf.Hashes, hash)
		for i := range buffer {
			buffer[i] = 0
		}
	}

	htf.RootHash = CalculateRootHash(htf.Hashes)
	return nil
}

func (ht *HashTree) VerifyTree(config *VerifyHashTreeConfig, chosenDirectory string) (bool, error) {
	if ht.RootDir == nil {
		return false, errors.New("hash tree not found to compare given directory to")
	}

	util.Logger.Infof("Verifying directory %s\n", chosenDirectory)
	return ht.verifyDir(config, chosenDirectory, "", ht.RootDir)
}

func (ht *HashTree) verifyDir(config *VerifyHashTreeConfig, currentDir string, directoryBeingVerified string, htDir *HashTreeDir) (bool, error) {
	util.Logger.Infof("verifying directory %s/%s\n", currentDir, directoryBeingVerified)
	file, err := os.Open(filepath.Join(currentDir, directoryBeingVerified))
	if err != nil {
		return false, err
	}

	defer file.Close()
	list, err := file.Readdirnames(0)
	if err != nil {
		return false, err
	}

	for _, name := range list {
		f, err := os.Stat(filepath.Join(currentDir, directoryBeingVerified, name))
		if err != nil {
			return false, err
		}

		if f.IsDir() {

			// check the subdir should exist
			subdir := func() *HashTreeDir {
				for _, htd := range htDir.Subdirs {
					if htd.Dirname == name {
						return htd
					}
				}
				return nil
			}()

			if subdir == nil {
				if config.IgnoreNewFilesAndDirs {
					continue
				}

				util.Logger.Warnf("Unexpected directory %s/%s/%s\n", currentDir, directoryBeingVerified, name)
				if !config.ContinueAfterError {
					return false, nil
				}
			}

			subdirRes, err := ht.verifyDir(config, filepath.Join(currentDir, directoryBeingVerified), name, subdir)
			if err != nil {
				return false, err
			}

			if !subdirRes {
				if !config.ContinueAfterError {
					return false, nil
				}
			}

		} else {

			// check file should exist
			fileExists := func() *HashTreeFile {
				for _, htf := range htDir.Files {
					if htf.Filename == name {
						return htf
					}
				}
				return nil
			}()

			if fileExists == nil {
				if config.IgnoreNewFilesAndDirs {
					continue
				}

				util.Logger.Warnf("Unexpected file %s/%s/%s\n", currentDir, directoryBeingVerified, name)
				if !config.ContinueAfterError {
					return false, nil
				}
			}

			// compare hashes
			fileRes, err := ht.verifyFile(fileExists, filepath.Join(currentDir, directoryBeingVerified), name)
			if err != nil {
				return false, err
			}

			if !fileRes {
				if !config.ContinueAfterError {
					return false, nil
				}
			}
		}
	}

	return true, nil
}

func (ht *HashTree) verifyFile(htf *HashTreeFile, currentDirectory string, filename string) (bool, error) {
	file, err := os.Open(filepath.Join(currentDirectory, filename))
	if err != nil {
		return false, err
	}
	defer file.Close()

	buffer := make([]byte, ht.ShardSize)
	reader := bufio.NewReader(file)
	counter := 0

	util.Logger.Infof("\tSharding file '%s'\n", filename)
	for {
		_, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return false, err
		}

		hash := sha256.Sum256(buffer)
		if !bytes.Equal(hash[:], htf.Hashes[counter][:]) {
			util.Logger.Errorf("Incorrect hash found in file %s/%s at block %d\n", currentDirectory, filename, counter)
			return false, nil
		}

		counter++
		for i := range buffer {
			buffer[i] = 0
		}
	}

	return true, nil
}

// Utility

func CalculateRootHash(hashes [][32]byte) [32]byte {
	oldLayer, newLayer := hashes, [][32]byte{}

	for len(oldLayer) != 1 {

		// duplicate the last element if there an odd number
		if len(oldLayer)%2 == 1 {
			oldLayer = append(oldLayer, oldLayer[len(oldLayer)-1])
		}

		// hash each pair
		for i := 0; i < len(oldLayer); i += 2 {
			newLayer = append(newLayer, sha256.Sum256(append(oldLayer[i][:], oldLayer[i+1][:]...)))
		}

		oldLayer = newLayer
		newLayer = [][32]byte{}
	}

	return oldLayer[0]

}

func (ht *HashTree) GetProgress() chan int {
	return ht.progress
}

func (htd1 *HashTreeDir) Equals(htd2 *HashTreeDir) bool {
	if len(htd1.Files) != len(htd2.Files) ||
		len(htd1.Subdirs) != len(htd2.Subdirs) ||
		!bytes.Equal(htd1.RootHash[:], htd2.RootHash[:]) ||
		htd1.Dirname != htd2.Dirname {
		return false
	}

	for filename, htf1 := range htd1.Files {
		if htf2, ok := htd2.Files[filename]; ok {
			if !htf1.Equals(htf2) {
				return false
			}

		} else {
			return false
		}
	}

	for subdir, sd1 := range htd1.Subdirs {
		if sd2, ok := htd2.Subdirs[subdir]; ok {
			if !sd1.Equals(sd2) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func (htf1 *HashTreeFile) Equals(htf2 *HashTreeFile) bool {
	if !bytes.Equal(htf1.RootHash[:], htf2.RootHash[:]) ||
		htf1.Filename != htf2.Filename ||
		len(htf1.Hashes) != len(htf2.Hashes) {
		return false
	}

	for i, h1 := range htf1.Hashes {
		if !bytes.Equal(h1[:], htf2.Hashes[i][:]) {
			return false
		}
	}

	return true
}
