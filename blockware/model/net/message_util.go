package net

import (
	"encoding/hex"
	"io"

	"github.com/t02smith/part-iii-project/toolkit/model/games"
)

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
