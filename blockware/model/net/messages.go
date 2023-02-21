package net

import (
	"errors"
	"fmt"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

// process a message received from a peer
func onMessage(cmd []string, client PeerIT) {
	switch cmd[0] {

	// LIBRARY => request a list of a peers games
	case "LIBRARY":
		err := handleLIBRARY(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error handling LIBRARY message: %s", err)
		}
		return

	// GAMES => a list of users games
	case "GAMES":
		err := handleGAMES(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error handling GAMES message: %s", err)
		}

		return

	// BLOCK <game hash> <hash> => Request a block of data from a user
	case "BLOCK":
		err := handleBLOCK(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error handling BLOCK message %s", err)
		}
		return

	// SEND_BLOCK <game hash> <hash> <shard> => Download a shard off of a user
	case "SEND_BLOCK":
		err := handleSEND_BLOCK(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error handling SEND_BLOCK message %s", err)
		}
		return

	// ERROR <msg> => used to send an error message following a command
	case "ERROR":
		util.Logger.Errorf("Error received %s", cmd[1])
		return
	}
}

// * DOWNLOADS

// listen for incoming download requests
func (p *peer) listenToDownloadRequests() {
	util.Logger.Info("Listening for incoming download requests")
	go func() {
		for request := range p.library.RequestDownload {
			fmt.Println(request)
		}
	}()
}

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
func refreshLibraries() {
	util.Logger.Info("Refreshing peer library data")
	Peer().Broadcast(generateLIBRARY())
	time.Sleep(50 * time.Second)
}
