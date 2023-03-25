package peer

import (
	"errors"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

This file defines useful communication functions for the peer
such as how it should respond to incoming messages and when it
should contact other peers with messages

*/

const (
	MaxAttemptsDownloadRequest int           = 32
	minRefreshPeriod           time.Duration = time.Minute * 5
)

// process a message received from a peer
func onMessage(cmd []string, client tcp.TCPConnection) error {
	if cmd[0][len(cmd[0])-1] == '\r' {
		cmd[0] = cmd[0][:len(cmd[0])-1]
	}

	pd := Peer().GetPeer(client)
	if !Peer().config.SkipValidation && !pd.Validator.Valid() && cmd[0] != "VALIDATE_REQ" && cmd[0] != "VALIDATE_RES" {
		util.Logger.Warnf("Peer not validated => discarding message")
	}

	handler := getHandlerByCommand(cmd[0])
	if err := handler(cmd, client); err != nil {
		util.Logger.Warnf("Error handling %s command: %s", cmd[0], err)
		return err
	}

	return nil

}

// * DOWNLOADS

// fetch a block given a game identifier and a shard
func fetchBlockFromLibrary(gameHash, shardHash [32]byte) ([]byte, error) {
	p := Peer()

	// find the game
	g := p.library.GetOwnedGame(gameHash)
	if g == nil {
		return nil, errors.New("game not in library")
	}

	d := g.GetDownload()

	// game is being downloaded and we may not have block
	if d != nil {
		if _, ok := d.Progress[shardHash]; ok {
			return nil, errors.New("shard hasn't been downloaded yet")
		}
	}

	// fetch the block
	found, b, err := g.FetchShard(shardHash)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("block not found")
	}

	return b, nil
}

// * UTIL

// refresh all known peer's game libraries
func (p *peer) RefreshLibraries() {
	util.Logger.Info("Refreshing peer library data")
	p.Broadcast(generateLIBRARY())
}
