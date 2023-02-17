package net

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func onMessage(cmd []string, client PeerIT) {
	p := GetPeerInstance()

	switch cmd[0] {

	// LIBRARY => request a list of a peers games
	case "LIBRARY":
		util.Logger.Info("Library command called")

		gameLs, err := games.LoadGames(viper.GetString("meta.directory"))
		if err != nil {
			util.Logger.Errorf("Error loading games: %s\n", err)
			return
		}

		gameStr, err := gameListToMessage(gameLs)
		if err != nil {
			util.Logger.Errorf("Error serialising games: %s\n", err)
			return
		}

		client.SendString(gameStr)
		return

	// GAMES => a list of users games
	case "GAMES":
		util.Logger.Infof("Games command called")

		ls, err := gameMessageToGameList(cmd)
		if err != nil {
			util.Logger.Errorf("Error reading games: %s\n", err)
			return
		}

		if peer, ok := p.peers[client]; ok {
			util.Logger.Infof("Client found. Updating library")
			peer.Library = ls
		}

		return

	// BLOCK <game hash> <hash> => Request a block of data from a user
	case "BLOCK":
		util.Logger.Infof("Block command called for block %s", cmd[2])

		gh, err := stringTo32ByteArr(cmd[1])
		if err != nil {
			util.Logger.Errorf("Error reading game hash on BLOCK cmd: %s", err)
			return
		}

		sh, err := stringTo32ByteArr(cmd[2])
		if err != nil {
			util.Logger.Errorf("Error reading shard hash on BLOCK cmd: %s", err)
			return
		}

		found, data, err := p.library.FindBlock(gh, sh)
		if err != nil {
			util.Logger.Errorf("Error finding block %s", err)
			return
		}

		if !found {
			util.Logger.Warnf("Block %x not found", sh)
			client.SendStringf("ERROR;Block %x not found\n", sh)
			return
		}

		client.SendString(fmt.Sprintf("SEND_BLOCK;%x;%x;%x\n", gh, sh, data))
		return

	// SEND_BLOCK <game hash> <hash> <shard> => Download a shard off of a user
	case "SEND_BLOCK":
		util.Logger.Infof("SEND_BLOCK => Block received")
		// * parse input
		gh, err := stringTo32ByteArr(cmd[1])
		if err != nil {
			util.Logger.Errorf("Error reading game hash on BLOCK cmd: %s", err)
			return
		}

		sh, err := stringTo32ByteArr(cmd[2])
		if err != nil {
			util.Logger.Errorf("Error reading shard hash on BLOCK cmd: %s", err)
			return
		}

		// * fetch game and find location
		game := p.library.GetOwnedGame(gh)
		gameTree, err := game.GetData()
		if err != nil {
			util.Logger.Errorf("Error loading game data %s", err)
			return
		}

		download := game.GetDownload()
		if download == nil {
			util.Logger.Errorf("Download not found for %x", game.RootHash)
			return
		}

		_, file, _, _ := gameTree.FindShard(sh)
		_, ok := download.Progress[file.RootHash].BlocksRemaining[sh]

		if !ok {
			util.Logger.Warnf("Block %x not needed for download", sh)
			return
		}

		// * insert the shard
		data, err := hex.DecodeString(cmd[3])
		if err != nil {
			util.Logger.Error(err)
			return
		}

		// * verify shard
		dataHash := sha256.Sum256(data)
		if !bytes.Equal(dataHash[:], sh[:]) {
			util.Logger.Errorf("Data given does not match expected hash\ngot %x\nexpected %x", dataHash, sh)
			return
		}

		err = game.InsertShard(sh, data)
		if err != nil {
			util.Logger.Errorf("error inserting shard %x: %s", sh, err)
		}

		delete(download.Progress[file.RootHash].BlocksRemaining, sh)
		err = download.Serialise(fmt.Sprintf("%x", game.RootHash))
		if err != nil {
			util.Logger.Errorf("Error serialising download %s", err)
		}
		return

	// ERROR <msg> => used to send an error message following a command
	case "ERROR":
		util.Logger.Errorf("Error received %s", cmd[1])
		return

	}
}

//

// turns a game library to a message to send to a peer
func gameListToMessage(games []*games.Game) (string, error) {
	var buf bytes.Buffer
	buf.WriteString("GAMES;")

	for i, g := range games {
		encoded, err := g.Serialise()
		if err != nil {
			return "", nil
		}

		buf.WriteString(encoded)
		if i != len(games)-1 {
			buf.WriteString(";")
		}
	}

	buf.WriteString("\n")
	return buf.String(), nil
}

// takes a game library message and converts it to a list of games
func gameMessageToGameList(parts []string) ([]*games.Game, error) {
	gameLs := []*games.Game{}

	for i := 1; i < len(parts); i++ {
		g, err := games.DeserialiseGame(parts[i])

		if err == io.EOF {
			return gameLs, nil
		}

		if err != nil {
			return nil, err
		}

		gameLs = append(gameLs, g)
	}

	return gameLs, nil
}

// fetch a block given a game identifier and a shard
func fetchBlock(gameHash, shardHash [32]byte) ([]byte, error) {
	p := GetPeerInstance()

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

func stringTo32ByteArr(hexString string) ([32]byte, error) {
	var arr [32]byte
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return arr, err
	}

	copy(arr[:], data)
	return arr, nil
}