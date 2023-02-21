package net

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func onMessage(cmd []string, client PeerIT) {
	p := Peer()

	switch cmd[0] {

	// LIBRARY => request a list of a peers games
	case "LIBRARY":
		util.Logger.Info("Library command called")

		gameLs, err := games.LoadGames(filepath.Join(viper.GetString("meta.directory"), "games"))
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
		found, file, _, _ := gameTree.FindShard(sh)
		if !found {
			util.Logger.Errorf("Shard %x not found", sh)
		}

		download := game.GetDownload()
		if download == nil {
			util.Logger.Errorf("Download not found for %x", game.RootHash)
			return
		}

		data, err := hex.DecodeString(cmd[3])
		if err != nil {
			util.Logger.Error(err)
			return
		}

		err = download.InsertData(file.RootHash, sh, data)
		if err != nil {
			util.Logger.Error(err)
			return
		}

		// send message as confirmation
		// p.library.DownloadProgress <- &struct {
		// 	GameHash  [32]byte
		// 	BlockHash [32]byte
		// }{
		// 	GameHash:  gh,
		// 	BlockHash: sh,
		// }

		util.Logger.Infof("Successfully inserted shard %x", sh)
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

// turn a hex string (from a hash) into a 32 byte array
func stringTo32ByteArr(hexString string) ([32]byte, error) {
	var arr [32]byte
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return arr, err
	}

	copy(arr[:], data)
	return arr, nil
}
