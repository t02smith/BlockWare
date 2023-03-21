package peer

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

// utility

func setupContributions(t *testing.T) (*contributions, common.Address) {
	t.Helper()
	c := newContributions()
	addr := testutil.GetAddress(testutil.Accounts[0][1])
	return c, addr
}

// tests

/*
function: newContributions
purpose: create a new contributions tracker
? Test cases
success
	#1 base case
*/

func TestNewContributions(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := newContributions()

		assert.NotNil(t, c.blocks)
		assert.Zero(t, c.total)
	})
}

/*
function: writeContributions
purpose: write the current set of contributions to file and flush memory
? Test cases
success
	#1 => no contributions to write
	#2 => create new file
	#3 => append existing file
failure
	#1 => ??
*/

func TestWriteContributions(t *testing.T) {
	c, addr := setupContributions(t)
	peersFolder := filepath.Join(viper.GetString("meta.directory"), "peers")

	t.Run("success", func(t *testing.T) {
		t.Cleanup(func() {
			os.Remove(filepath.Join(peersFolder, addr.Hex()))
		})

		t.Run("no contributions", func(t *testing.T) {
			n := c.writeContributions()
			assert.Zero(t, n)
		})

		testHashI := sha256.Sum256([]byte("hello world"))

		t.Run("create a new file", func(t *testing.T) {
			c.addContribution(addr, testHashI)

			n := c.writeContributions()
			assert.Equal(t, 1, n)
			assert.Empty(t, c.blocks)

			_, err := os.Stat(filepath.Join(peersFolder, addr.Hex()))
			assert.Nil(t, err)

			f, err := os.Open(filepath.Join(peersFolder, addr.Hex()))
			assert.Nil(t, err)
			defer f.Close()

			reader := bufio.NewReader(f)

			var buffer [32]byte
			reader.Read(buffer[:])

			fmt.Printf("%x - %x\n", testHashI, buffer)
			assert.True(t, bytes.Equal(testHashI[:], buffer[:]))
		})

		t.Run("append file", func(t *testing.T) {
			testHashII := sha256.Sum256([]byte("tom smith"))
			c.addContribution(addr, testHashII)

			n := c.writeContributions()
			assert.Equal(t, 1, n)
			assert.Empty(t, c.blocks)

			_, err := os.Stat(filepath.Join(peersFolder, addr.Hex()))
			assert.Nil(t, err)

			f, err := os.Open(filepath.Join(peersFolder, addr.Hex()))
			assert.Nil(t, err)
			defer f.Close()

			reader := bufio.NewReader(f)

			var buffer [32]byte

			reader.Read(buffer[:])
			assert.True(t, bytes.Equal(buffer[:], testHashI[:]))

			reader.Read(buffer[:])
			assert.True(t, bytes.Equal(buffer[:], testHashII[:]))

		})
	})
}

/*
function: addContribution
purpose: store a contribution from a given address
? Test cases
success
	#1 => new address
	#2 => existing address
*/

func TestAddContribution(t *testing.T) {
	c, addr := setupContributions(t)

	testHashI := sha256.Sum256([]byte("hello world"))
	testHashII := sha256.Sum256([]byte("tom smith"))

	t.Run("success", func(t *testing.T) {
		t.Run("new address", func(t *testing.T) {
			c.addContribution(addr, testHashI)

			res, ok := c.blocks[addr]
			assert.True(t, ok)
			assert.Equal(t, 1, len(res))
			assert.True(t, bytes.Equal(res[0][:], testHashI[:]))
			assert.Equal(t, 1, int(c.total))
		})

		t.Run("existing address", func(t *testing.T) {
			c.addContribution(addr, testHashII)

			res, ok := c.blocks[addr]
			assert.True(t, ok)
			assert.Equal(t, 2, len(res))
			assert.True(t, bytes.Equal(res[0][:], testHashI[:]))
			assert.True(t, bytes.Equal(res[1][:], testHashII[:]))
			assert.Equal(t, 2, int(c.total))
		})
	})
}

/*
function: GetContributionsFromPeer
purpose: retrieve a list of contributions from file
? Test cases
success
	#1 => no blocks read
	#2 => some blocks read
failure
	illegal arguments
			#1 => file for address doesn't exist
*/

func TestGetContributionsFromPeer(t *testing.T) {
	_, addr := setupContributions(t)

	t.Run("success", func(t *testing.T) {
		t.Run("no blocks read", func(t *testing.T) {
			f, err := os.Create(filepath.Join("../../../test/data/tmp/.toolkit/peers", addr.Hex()))
			if err != nil {
				t.Fatal(err)
			}
			f.Close()
			t.Cleanup(func() {
				os.Remove(filepath.Join("../../../test/data/tmp/.toolkit/peers", addr.Hex()))
			})

			res, err := GetContributionsFromPeer(addr)
			assert.Nil(t, err)
			assert.Empty(t, res)
		})

		t.Run("some blocks read", func(t *testing.T) {
			f, err := os.Create(filepath.Join("../../../test/data/tmp/.toolkit/peers", addr.Hex()))
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				os.Remove(filepath.Join("../../../test/data/tmp/.toolkit/peers", addr.Hex()))
			})

			blocks := [][32]byte{
				sha256.Sum256([]byte("hello world")),
				sha256.Sum256([]byte("tom smith")),
				sha256.Sum256([]byte("blockware")),
				sha256.Sum256([]byte("testing")),
			}

			for _, b := range blocks {
				f.Write(b[:])

			}

			f.Close()

			res, err := GetContributionsFromPeer(addr)
			assert.Nil(t, err)
			assert.Equal(t, len(blocks), len(res))
			for i, b := range blocks {
				assert.True(t, bytes.Equal(b[:], blocks[i][:]))
			}
		})
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("address file not found", func(t *testing.T) {
			addr := testutil.GetAddress(testutil.Accounts[1][1])

			res, err := GetContributionsFromPeer(addr)
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, os.ErrNotExist)
			assert.Empty(t, res)
		})
	})
}
