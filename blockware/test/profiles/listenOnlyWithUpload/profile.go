package profileListenOnlyUpload

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum/library"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

Profile 1:

Features:
- listen-and-respond only peer
- will upload a game to ETH and seed it
- ideal connection

Data:

- small size project (latex-template from github)
	a small scale project that will allow for the peer to download
	from many different files.

- fake large file
	TODO ^^
	A large amount of blocks that can be easily generated without
	storing them => e.g. block 0 has a content of ...0000, block 1
	is ...00001, etc.
	Will only need to store the hash with the offset to be able to
	generate it

*/

var PrivateKey string = testutil.Accounts[1][1]

func Run() {
	// if _, err := os.Stat("../../data/tmp/games/profile-2"); err != nil {
	// 	if err = testutil.GenerateLargeFolder("profile-2", "../../data/tmp/games/", 40_000_000_000, 1); err != nil {
	// 		util.Logger.Fatal(err)
	// 	}
	// }

	p := peer.Peer()

	// // * create & upload game
	// g, err := games.CreateGame(games.NewGame{
	// 	Title:       "benchmark",
	// 	Version:     "4.7.1",
	// 	ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
	// 	Developer:   "tcs1g20",
	// 	RootDir:     "../../data/tmp/games/profile-2",
	// 	Price:       big.NewInt(150),
	// 	ShardSize:   4194304,
	// 	AssetsDir:   "../../data/assets"},
	// 	nil,
	// )
	// _ = g

	// if err != nil {
	// 	util.Logger.Fatalf("Error creating game: %s", err)
	// }

	g1, err := games.CreateGame(games.NewGame{
		Title:       "website",
		Version:     "4.7.1",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "tcs1g20",
		RootDir:     "./t02smith.github.io",
		Price:       big.NewInt(150),
		ShardSize:   4194304,
		AssetsDir:   "../../data/assets"},
		nil,
	)
	_ = g1

	if err != nil {
		util.Logger.Errorf("Error creating game: %s", err)
	}

	g2, err := games.CreateGame(games.NewGame{
		Title:       "Snake",
		Version:     "1.17.4",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "Tom Smith",
		RootDir:     "../../../model",
		Price:       big.NewInt(10),
		ShardSize:   4194304,
		AssetsDir:   "../../data/assets"},
		nil,
	)
	_ = g2

	if err != nil {
		util.Logger.Errorf("Error creating game: %s", err)
	}

	for _, _g := range []*games.Game{g1, g2} {
		p.Library().AddOrUpdateOwnedGame(_g)
		if err := library.Upload(_g); err != nil {
			util.Logger.Fatalf("Error uploading game to ETH %s", err)
		}

		if err := games.OutputAllGameDataToFile(_g); err != nil {
			util.Logger.Warnf("Error saving game to file %s", err)
		}
	}

	// * listen for connections
	util.Logger.Info("PROFILE 1: listening")
	<-make(chan bool) // wait forever :/
}

func generateLargeFile() (map[[32]byte]uint32, error) {
	var out map[[32]byte]uint32
	if _, err := os.Stat("./large-file"); err == nil {
		// * read file
		f, err := os.Open("./large-file")
		if err != nil {
			return nil, err
		}
		f.Close()

		reader, err := gzip.NewReader(f)
		if err != nil {
			return nil, err
		}

		decoder := gob.NewDecoder(reader)
		err = decoder.Decode(&out)
		if err != nil {
			return nil, err
		}

		return out, nil
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	// ! will result in a 20.48 GB in size file
	// ! storing the hash data will tak 720MB RAM :/
	// ! file is stored using GZIP
	shardSize, blockCount := 8192, 2_500_00
	out = make(map[[32]byte]uint32)
	util.Logger.Warnf("This function will take a long time, Ctrl-C to cancel")
	time.Sleep(5 * time.Second)

	// * generate data
	for i := 0; i < blockCount; i++ {
		buf := make([]byte, shardSize)
		copy(buf, []byte(fmt.Sprint(i)))
		hash := sha256.Sum256(buf)
		out[hash] = uint32(i)
	}

	f, err := os.Create("./large-file")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	writer := gzip.NewWriter(f)
	encoder := gob.NewEncoder(writer)

	err = encoder.Encode(out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
