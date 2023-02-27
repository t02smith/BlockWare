package net

import (
	"encoding/hex"
	"fmt"
)

// turn a hex string (from a hash) into a 32 byte array
func stringTo32ByteArr(hexString string) ([32]byte, error) {
	if len(hexString) != 64 {
		return [32]byte{}, fmt.Errorf("invalid hash length for hash %s", hexString)
	}

	var arr [32]byte
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return arr, err
	}

	copy(arr[:], data)
	return arr, nil
}
