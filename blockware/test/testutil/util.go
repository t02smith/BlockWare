package testutil

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func SetupTmp(toRoot string) error {

	err := os.Mkdir(filepath.Join(toRoot, "./test/data/tmp/tracker"), 0775)
	if err != nil {
		return err
	}

	return nil
}

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

func ShortTest(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
}

func LongTest(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
}
