package peer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

This file defines useful communication functions for the peer
such as how it should respond to incoming messages and when it
should contact other peers with messages

*/

const (
	MaxAttemptsDownloadRequest int = 32
)

// process a message received from a peer
func onMessage(cmd []string, client tcp.TCPConnection) error {
	if cmd[0][len(cmd[0])-1] == '\r' {
		cmd[0] = cmd[0][:len(cmd[0])-1]
	}

	pd := Peer().peers[client]
	if !Peer().config.SkipValidation && !pd.Validator.Valid() && cmd[0] != "VALIDATE_REQ" && cmd[0] != "VALIDATE_RES" {
		util.Logger.Warnf("Peer not validated => discarding message")
	}

	switch cmd[0] {

	// LIBRARY => request a list of a peers games
	case "LIBRARY":
		err := handleLIBRARY(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error handling LIBRARY message: %s", err)
			client.SendString(generateERROR(err.Error()))
		}
		return nil

	// GAMES => a list of users games
	case "GAMES":
		err := handleGAMES(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error handling GAMES message: %s", err)
			client.SendString(generateERROR(err.Error()))
		}

		return nil

	// BLOCK <game hash> <hash> => Request a block of data from a user
	case "BLOCK":
		err := handleBLOCK(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error handling BLOCK message %s", err)
			client.SendString(generateERROR(err.Error()))
		}
		return nil

	// SEND_BLOCK <game hash> <hash> <shard> => Download a shard off of a user
	case "SEND_BLOCK":
		err := handleSEND_BLOCK(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error handling SEND_BLOCK message %s", err)
			client.SendString(generateERROR(err.Error()))
		}
		return nil

	// ERROR <msg> => used to send an error message following a command
	case "ERROR":
		util.Logger.Errorf("Error received %s", cmd[1])
		return nil

	// VALIDATE_REQ <message> => request a signed message to prove identity
	case "VALIDATE_REQ":
		err := handleVALIDATE_REQ(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error reading validation request: %s", err)
		}
		return nil

	// VALIDATE_RES <signed message> => validate someone's public key
	case "VALIDATE_RES":
		err := handleVALIDATE_RES(cmd, client)
		if err != nil {
			util.Logger.Warnf("Error reading validation response: %s", err)
		}
		return nil
	}

	return fmt.Errorf("unrecognised message: %s", strings.Join(cmd, ";"))
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
func (p *peer) refreshLibraries() {
	util.Logger.Info("Refreshing peer library data")
	p.Broadcast(generateLIBRARY())
}