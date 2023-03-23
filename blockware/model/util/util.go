package util

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

type Void struct{}

// creates the required directories for toolkit
func SetupToolkitEnvironment() error {

	// check for .toolkit directory
	toolkitDir := viper.GetString("meta.directory")
	if len(toolkitDir) == 0 {
		toolkitDir = ".toolkit"
	}

	// toolkit root directory
	err := CreateDirectoryIfNotExist(toolkitDir)
	if err != nil {
		return err
	}

	// look for games directory
	gamesDir := filepath.Join(toolkitDir, "games")
	err = CreateDirectoryIfNotExist(gamesDir)
	if err != nil {
		return err
	}

	// look for hash directory
	hashDir := filepath.Join(toolkitDir, "hashes")
	err = CreateDirectoryIfNotExist(hashDir)
	if err != nil {
		return err
	}

	// look for assets directory
	assetsDir := filepath.Join(toolkitDir, "assets")
	err = CreateDirectoryIfNotExist(assetsDir)
	if err != nil {
		return err
	}

	// look for assets directory
	peerDir := filepath.Join(toolkitDir, "peers")
	err = CreateDirectoryIfNotExist(peerDir)
	if err != nil {
		return err
	}

	// look for assets directory
	keyStoreDir := filepath.Join(toolkitDir, "keystores")
	err = CreateDirectoryIfNotExist(keyStoreDir)
	if err != nil {
		return err
	}

	return nil
}

// create a directory if it doesn't exist
func CreateDirectoryIfNotExist(dir string) error {
	stat, err := os.Stat(dir)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		util.Logger.Infof("Directory %s not found. Creating directory", dir)
		err = os.Mkdir(dir, 0755)

		if err != nil {
			return err
		}
	} else {
		if !stat.IsDir() {
			return fmt.Errorf("a file already exists with the name %s", dir)
		}
	}

	return nil
}

// * zip archives

// create a zip archive of a directory
func ZipDirectory(archivePath, outputLocation string) error {
	util.Logger.Debugf("Attempting to create archive %s from directory %s", outputLocation, archivePath)
	file, err := os.Create(outputLocation)
	if err != nil {
		return err
	}

	compressor := zip.NewWriter(file)
	defer compressor.Close()

	_archivePath := fmt.Sprintf("%s%s", filepath.Join(archivePath), string(os.PathSeparator))
	fileCount, dirCount := 0, 0
	err = filepath.WalkDir(_archivePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(path, _archivePath)
		if relPath == "" {
			return nil
		}

		if d.IsDir() {
			_, err = compressor.Create(relPath + "/")
			dirCount++
			return err
		}

		writer, err := compressor.Create(relPath)
		if err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		_, err = writer.Write(data)
		if err != nil {
			return err
		}

		fileCount++
		return nil
	})

	if err != nil {
		file.Close()
		util.Logger.Errorf("Error creating archive from %s: %s", archivePath, err)
		_err := os.Remove(outputLocation)
		if _err != nil {
			util.Logger.Errorf("Error clearing archive file %s", err)
		}

		return err
	}

	util.Logger.Infof("Archive created at %s from %d folders and %d files", dirCount, fileCount)
	defer file.Close()
	return compressor.Flush()
}

// unzip a .zip archive
func UnzipArchive(archive string, outputFolder string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer reader.Close()

	err = os.Mkdir(outputFolder, 0755)
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		filepath := filepath.Join(outputFolder, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(filepath, f.Mode())
		} else {
			var fileDir string
			if lastIndex := strings.LastIndex(filepath, string(os.PathSeparator)); lastIndex > -1 {
				fileDir = filepath[:lastIndex]
			}

			err = os.MkdirAll(fileDir, f.Mode())
			if err != nil {
				return err
			}

			f, err := os.Create(filepath)
			if err != nil {
				return err
			}

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}

			f.Close()
		}

		rc.Close()
	}

	return nil
}
