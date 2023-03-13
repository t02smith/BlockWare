package ignore

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
)

/*
Ignore given folders/files when creating a hash tree
Standard .ignore implementation
*/

type Ignore []string

func (i Ignore) Allowed(rootDir, path string) bool {
	oldDir, _ := os.Getwd()
	if err := os.Chdir(rootDir); err != nil {
		return false
	}
	defer os.Chdir(oldDir)

	for _, matcher := range i {
		if path == matcher {
			return false
		}

		files, err := filepath.Glob(matcher)
		if err != nil {
			return false
		}

		for _, f := range files {
			if f == path {
				return false
			}
		}
	}

	return true
}

func ReadIgnoreFromFile(filename string) (Ignore, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out Ignore
	reader := bufio.NewReader(f)
	for {
		ln, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		if ln[len(ln)-1] == '\n' {
			ln = ln[:len(ln)-1]
		}

		if ln[len(ln)-1] == '\r' {
			ln = ln[:len(ln)-1]
		}

		out = append(out, ln)
	}

	return out, nil
}
