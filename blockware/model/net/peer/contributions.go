package peer

import (
	"bufio"
	"errors"
	"fmt"
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

var contributionThresholdBeforeWrite uint = 50

type contributions struct {
	// what blocks have they sent us from which games
	games map[common.Address]map[[32]byte][][32]byte

	// to protect concurrent access of the blocks map
	lock sync.Mutex

	// currently loaded number of contributions
	total uint
}

// create a new contributions tracker
func newContributions() *contributions {
	return &contributions{
		games: make(map[common.Address]map[[32]byte][][32]byte),
		total: 0,
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
	for addr, games := range c.games {
		if len(games) == 0 {
			continue
		}

		if _, err := os.Stat(filepath.Join(peersDirectory, addr.Hex())); errors.Is(err, os.ErrNotExist) {
			if err := os.Mkdir(filepath.Join(peersDirectory, addr.Hex()), 0755); err != nil {
				util.Logger.Error(err)
				return -1
			}
		}

		for hash, blocks := range games {
			f, err := os.OpenFile(filepath.Join(peersDirectory, addr.Hex(), fmt.Sprintf("%x", hash)), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				util.Logger.Warnf("error writing to peer file %s", err)
			}

			writer := bufio.NewWriter(f)
			for _, b := range blocks {
				writer.Write(b[:])
				counter++
			}

			writer.Flush()
			f.Close()
		}

	}
	c.games = make(map[common.Address]map[[32]byte][][32]byte)
	c.total = 0

	util.Logger.Debugf("Written %d contributions to file", counter)
	return counter
}

// add a new contribution from a given address
func (c *contributions) addContribution(addr common.Address, game, block [32]byte) {
	c.lock.Lock()

	if games, ok := c.games[addr]; ok {
		if _, ok := games[game]; ok {
			c.games[addr][game] = append(c.games[addr][game], block)
		} else {
			c.games[addr][game] = [][32]byte{block}
		}
	} else {
		c.games[addr] = make(map[[32]byte][][32]byte)
		c.games[addr][game] = [][32]byte{block}

	}

	c.total++

	if c.total >= contributionThresholdBeforeWrite {
		c.lock.Unlock()
		c.writeContributions()
	} else {
		c.lock.Unlock()
	}
}

func GetContributionsFromPeer(addr common.Address, game [32]byte) ([][32]byte, error) {
	var blocks [][32]byte

	f, err := os.Open(filepath.Join(viper.GetString("meta.directory"), "peers", addr.Hex(), fmt.Sprintf("%x", game)))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return [][32]byte{}, nil
		}

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

// request a peer's contribution for a given game
func (pd *peerData) RequestContributionsReceipt(game [32]byte) {
	pd.Peer.SendString(generateREQ_RECEIPT(game))
}
