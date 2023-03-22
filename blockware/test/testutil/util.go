package testutil

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

func ClearTmp(toRoot string) {
	tmp, err := os.Open(filepath.Join(toRoot, "./test/data/tmp"))
	if err != nil {
		log.Println(err)
		return
	}

	ls, err := tmp.Readdirnames(0)
	if err != nil {
		log.Println(err)
		return
	}

	for _, name := range ls {
		if name != ".gitkeep" {
			err = os.RemoveAll(filepath.Join(toRoot, "./test/data/tmp", name))
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func GenerateLargeFolder(name, root string, fileSize, fileCount uint) error {
	err := os.MkdirAll(filepath.Join(root, name), 0755)
	if err != nil {
		return err
	}

	rand.Seed(17072847)
	data := make([]byte, fileSize)
	for j := 0; j < int(fileSize); j++ {
		data[j] = byte(rand.Intn(256))
	}

	for i := 0; i < int(fileCount); i++ {
		f, err := os.Create(filepath.Join(root, name, fmt.Sprintf("%d.txt", i)))
		if err != nil {
			return err
		}

		writer := bufio.NewWriter(f)
		_, err = writer.Write(data)
		if err != nil {
			return err
		}

		err = writer.Flush()
		if err != nil {
			return err
		}

		f.Close()
	}

	return nil
}
