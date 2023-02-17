package games

import (
	"bytes"
	"testing"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestUploadAndDownloadDataToIPFS(t *testing.T) {
	testutil.LongTest(t)
	t.Skip() // TODO setup IPFS daemon on github workflow

	g, err := fetchTestGame()
	if err != nil {
		t.Fatal(err)
	}

	smoke := t.Run("upload data", func(t *testing.T) {

		err := g.UploadDataToIPFS()
		if err != nil {
			t.Fatal(err)
		}

		sh := shell.NewShell("localhost:5001")
		data, err := sh.Cat(g.IPFSId)
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

		err := g.ReadDataFromIPFS()
		if err != nil {
			t.Fatal(err)
		}

		if !cachedData.Equals(g.data) {
			t.Fatal("hash trees do not match")
		}

	})
}