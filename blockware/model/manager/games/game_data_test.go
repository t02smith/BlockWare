package games

import (
	"bytes"
	"testing"

	shell "github.com/ipfs/go-ipfs-api"
)

/*

function: UploadHashTree + DownloadHashTree
purpose: upload the hash tree of a game to distributed storage

? Test cases
success
	#1 => tree uploaded successfully
	#2 => tree downloaded successfully


*/

func TestUploadAndDownloadHashTree(t *testing.T) {
	t.Skip() // TODO setup IPFS daemon on github workflow

	g := fetchTestGame(t)

	smoke := t.Run("upload data", func(t *testing.T) {

		err := g.UploadHashTree()
		if err != nil {
			t.Fatal(err)
		}

		sh := shell.NewShell("localhost:5001")
		data, err := sh.Cat(g.HashTreeIPFSAddress)
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(data)

		if len(buf.Bytes()) == 0 {
			t.Fatal("no data stored on ipfs")
		}
	})

	if !smoke {
		t.Skip()
	}

	t.Run("download data", func(t *testing.T) {

		cachedData := g.data
		g.data = nil

		err := g.DownloadHashTree()
		if err != nil {
			t.Fatal(err)
		}

		if !cachedData.Equals(g.data) {
			t.Fatal("hash trees do not match")
		}

	})
}
