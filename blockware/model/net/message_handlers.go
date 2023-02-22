package net

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

This file defines specific interactions/commands. For each command it:
1. has a generator function for creating a structured message given the
   required set of parameters. This reduces code duplication.
2. has a handler to deal with an incoming message from another peer,

Current handlers are:
LIBRARY => request a users game library
GAMES => receive a users game library
BLOCK => request a block of data from a user
SEND_BLOCK => receive a block of data from a user
ERROR => an error message to send to a peer
				 TODO properly deal with incoming errors

*/

// LIBRARY

func generateLIBRARY() string {
	return "LIBRARY\n"
}

func handleLIBRARY(cmd []string, client PeerIT) error {
	util.Logger.Info("Library command called")

	gameLs, err := games.LoadGames(filepath.Join(viper.GetString("meta.directory"), "games"))
	if err != nil {
		return err
	}

	return client.SendString(generateGAMES(gameLs...))
}

// GAMES <game hash>, <game hash>, ...

func generateGAMES(games ...*games.Game) string {
	var buf bytes.Buffer
	buf.WriteString("GAMES;")

	for i, g := range games {
		_, err := buf.WriteString(fmt.Sprintf("%x", g.RootHash))
		if err != nil {
			util.Logger.Warnf("Errors writing root hash for game %x => skipping from message", g.RootHash)
			continue
		}

		if i != len(games)-1 {
			buf.WriteString(";")
		}
	}

	buf.WriteString("\n")
	return buf.String()
}

func handleGAMES(cmd []string, client PeerIT) error {
	util.Logger.Infof("Games command called")

	pData, ok := Peer().peers[client]
	if !ok {
		return errors.New("unknown peer")
	}

	for i := 1; i < len(cmd); i++ {
		var hash [32]byte

		receivedHash, err := hex.DecodeString(cmd[i])
		if err != nil {
			return err
		}

		copy(hash[:], receivedHash)
		pData.Library[hash] = true
	}

	return nil
}

// BLOCK <game hash> <block hash>

func generateBLOCK(gameHash, blockHash [32]byte) string {
	return fmt.Sprintf("BLOCK;%x;%x\n", gameHash, blockHash)
}

func handleBLOCK(cmd []string, client PeerIT) error {
	util.Logger.Infof("Block command called for block %s", cmd[2])

	gh, err := stringTo32ByteArr(cmd[1])
	if err != nil {
		return fmt.Errorf("error reading game hash on BLOCK cmd: %s", err)
	}

	sh, err := stringTo32ByteArr(cmd[2])
	if err != nil {
		return fmt.Errorf("error reading shard hash on BLOCK cmd: %s", err)
	}

	found, data, err := Peer().library.FindBlock(gh, sh)
	if err != nil {
		return fmt.Errorf("error finding block %s", err)
	}

	if !found {
		client.SendStringf(generateERROR("Block %x not found", sh))
		return fmt.Errorf("block %x not found", sh)
	}

	client.SendString(generateSEND_BLOCK(gh, sh, data))
	return nil
}

// SEND_BLOCK <game hash> <shard hash> <shard data>

func generateSEND_BLOCK(gameHash, shardHash [32]byte, data []byte) string {
	return fmt.Sprintf("SEND_BLOCK;%x;%x;%x\n", gameHash, shardHash, data)
}

func handleSEND_BLOCK(cmd []string, client PeerIT) error {
	util.Logger.Infof("SEND_BLOCK => Block received")
	// * parse input
	gh, err := stringTo32ByteArr(cmd[1])
	if err != nil {
		return fmt.Errorf("error reading game hash on BLOCK cmd: %s", err)
	}

	sh, err := stringTo32ByteArr(cmd[2])
	if err != nil {
		return fmt.Errorf("error reading shard hash on BLOCK cmd: %s", err)
	}

	// * fetch game and find location
	game := Peer().library.GetOwnedGame(gh)
	gameTree, err := game.GetData()
	if err != nil {
		return fmt.Errorf("error loading game data %s", err)
	}
	found, file, _, _ := gameTree.FindShard(sh)
	if !found {
		return fmt.Errorf("shard %x not found", sh)
	}

	download := game.GetDownload()
	if download == nil {
		return fmt.Errorf("download not found for %x", game.RootHash)
	}

	data, err := hex.DecodeString(cmd[3])
	if err != nil {
		return err
	}

	err = download.InsertData(file.RootHash, sh, data)
	if err != nil {
		return err
	}

	// send message as confirmation
	// Peer().library.DownloadProgress <- games.DownloadRequest{
	// 	GameHash:  gh,
	// 	BlockHash: sh,
	// }

	util.Logger.Infof("Successfully inserted shard %x", sh)
	return nil
}

// ERROR <msg>

func generateERROR(msg string, args ...any) string {
	return fmt.Sprintf("ERROR;%s\n", fmt.Sprintf(msg, args...))
}