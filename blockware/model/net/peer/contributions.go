package peer

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*
We will need to track the data that is sent to us from
each peer so that their contribution across the network
can be tracked.
*/

var contributionThresholdBeforeWrite uint = 5

type contributions struct {
	// what blocks have they sent us
	blocks map[common.Address][][32]byte

	// to protect concurrent access of the blocks map
	lock sync.Mutex

	// currently loaded number of contributions
	total uint
}

// create a new contributions tracker
func newContributions() *contributions {
	return &contributions{
		blocks: make(map[common.Address][][32]byte),
		total:  0,
	}
}

// write all existing contributions to file
func (c *contributions) writeContributions() int {
	peersDirectory := filepath.Join(viper.GetString("meta.directory"), "peers")
	c.lock.Lock()
	defer c.lock.Unlock()
	util.Logger.Debugf("Writing %d contributions to file", c.total)

	if c.total == 0 {
		return 0
	}

	counter := 0
	for addr, blocks := range c.blocks {
		if len(blocks) == 0 {
			continue
		}

		f, err := os.OpenFile(filepath.Join(peersDirectory, addr.Hex()), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			util.Logger.Warnf("error writing to file peer file %s", err)
		}

		for _, b := range blocks {
			_, err := f.Write(b[:])
			if err != nil {
				util.Logger.Warnf("err writing to peer file %s", err)
				continue
			}

			counter++
		}

		f.Close()
	}
	c.blocks = make(map[common.Address][][32]byte)
	c.total = 0

	util.Logger.Debugf("Written %d contributions to file", counter)
	return counter
}

// add a new contribution from a given address
func (c *contributions) addContribution(addr common.Address, block [32]byte) {
	c.lock.Lock()

	if blocks, ok := c.blocks[addr]; ok {
		c.blocks[addr] = append(blocks, block)
	} else {
		c.blocks[addr] = [][32]byte{block}
	}

	c.total++

	if c.total >= contributionThresholdBeforeWrite {
		c.lock.Unlock()
		c.writeContributions()
	} else {
		c.lock.Unlock()
	}
}

func GetContributionsFromPeer(addr common.Address) ([][32]byte, error) {
	var blocks [][32]byte

	f, err := os.Open(filepath.Join(viper.GetString("meta.directory"), "peers", addr.Hex()))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var buffer [1024]byte

	for {
		n, err := f.Read(buffer[:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		for i := 0; i < n/32; i++ {
			var block [32]byte
			copy(block[:], buffer[i:(i+32)])

			blocks = append(blocks, block)
		}
	}

	return blocks, nil
}
