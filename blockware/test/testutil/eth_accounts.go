package testutil

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/**

Run `make ganache` and you will get these addresses
[ address, private key ]

*/

var Accounts [][]string = [][]string{
	{"7AF55399f383575A4a1f22d47F7572d7A3010E63", "b3868cc6652c9279c088be8dbfe6f4ef2ab39ecc80b3c69602217fac64ed6ad4"},
	{"F5A2804a04c1704dAF418F718Eb29aD1C92418eb", "64bf0b39525b0c1be4a39b139efddeb0cbfa0cc00b8bcde269cd452eb06fe9d0"},
	{"412756437a70f0b4E57088D2180CEC1bC15bf85A", "048e2edf66cafcae2217bd111be906127a120ac8733f98755a09a320853c78bd"},
	{"5D24fF4aa073258Ff93c86aAF8a428B775207660", "ea2c8d4ed55712da948c2fbc013990df1334f763d788ae6c72adcfe2e423a6dd"},
}

func GetAddress(privKey string) common.Address {
	key, _ := crypto.HexToECDSA(privKey)
	return crypto.PubkeyToAddress(key.PublicKey)
}
