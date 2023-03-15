package games

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	hash "github.com/t02smith/part-iii-project/toolkit/model/manager/hashtree"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

// Game /*
type Game struct {

	// game metadata
	Title           string
	Version         string
	ReleaseDate     string
	Developer       string
	RootHash        [32]byte
	PreviousVersion [32]byte

	// IPFS
	HashTreeIPFSAddress string
	Assets              *GameAssets

	// blockchain related
	Price    *big.Int
	Uploader common.Address

	// the shard data
	data *hash.HashTree

	// a download if it exists
	Download *Download
}

// Creator

type NewGame struct {
	Title       string
	Version     string
	ReleaseDate string
	Developer   string
	RootDir     string
	Price       *big.Int
	ShardSize   uint
	AssetsDir   string
}

// CreateGame create a new instance of a game and generate a hash tree for it
func CreateGame(newGame NewGame, progress chan int) (*Game, error) {

	if newGame.ShardSize == 0 {
		util.Logger.Errorf("shard size should be > 0")
		return nil, errors.New("invalid shard size")
	}

	// check version format
	versionMatches, err := regexp.MatchString("^(\\d+\\.)*\\d+$", newGame.Version)
	if err != nil {
		util.Logger.Errorf("error matching version number to regex")
		return nil, err
	}

	if !versionMatches {
		util.Logger.Errorf("invalid version number")
		return nil, errors.New("invalid version number")
	}

	// check release date
	_, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Split(newGame.ReleaseDate, " m=")[0])
	if err != nil {
		util.Logger.Errorf("invalid release date given got %s", newGame.ReleaseDate)
		return nil, errors.New("invalid release date given")
	}

	// hash data
	tree, err := hash.NewHashTree(newGame.RootDir, newGame.ShardSize, progress)
	if err != nil {
		return nil, err
	}

	err = tree.Hash()
	if err != nil {
		return nil, err
	}

	// game root hash
	hasher := sha256.New()
	hasher.Write([]byte(newGame.Title))
	hasher.Write([]byte(newGame.Version))
	hasher.Write([]byte(newGame.ReleaseDate))
	hasher.Write([]byte(newGame.Developer))
	hasher.Write(tree.RootDir.RootHash[:])

	gameRootHash := hasher.Sum([]byte{})

	var h [32]byte
	copy(h[:], gameRootHash)

	// return value
	game := &Game{
		Title:               newGame.Title,
		Version:             newGame.Version,
		ReleaseDate:         newGame.ReleaseDate,
		Developer:           newGame.Developer,
		data:                tree,
		RootHash:            h,
		HashTreeIPFSAddress: "",
		Price:               newGame.Price,
		PreviousVersion:     [32]byte{},
		Uploader:            common.Address{},
		Assets: &GameAssets{
			AbsolutePath: newGame.AssetsDir,
		},
	}

	return game, nil
}

// IO

// GetData get a game's hash tree data and fetch it from a file if necessary
func (g *Game) GetData() (*hash.HashTree, error) {
	if g.data == nil {
		err := g.readHashData()
		if err != nil {
			return nil, err
		}
	}

	return g.data, nil
}

// read a game's hash data from a file
func (g *Game) readHashData() error {
	dir := filepath.Join(viper.GetString("meta.directory"), "hashes")
	if len(dir) == 0 {
		return errors.New(".toolkit directory not specified")
	}

	hashFileName := fmt.Sprintf("%x.hash", g.RootHash)
	data, err := hash.ReadHashTreeFromFile(filepath.Join(dir, hashFileName))
	if err != nil {
		return err
	}

	g.data = data
	return nil
}

// * IO

// OutputAllGameDataToFile output data to file
func OutputAllGameDataToFile(g *Game) error {
	gameFilename := filepath.Join(viper.GetString("meta.directory"), "games", fmt.Sprintf("%x", g.RootHash))
	err := g.OutputToFile()
	if err != nil {
		return err
	}

	// output game data
	err = g.data.OutputToFile(filepath.Join(viper.GetString("meta.directory"), "hashes", fmt.Sprintf("%x.hash", g.RootHash)))
	if err != nil {
		return err
	}

	util.Logger.Infof("Outputted game data to %s", gameFilename)
	return nil
}

// OutputToFile output game metadata and download info to file
func (g *Game) OutputToFile() error {
	filename := filepath.Join(viper.GetString("meta.directory"), "games", fmt.Sprintf("%x", g.RootHash))
	util.Logger.Debugf("Outputting game data to %s", filename)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := gzip.NewWriter(f)
	encoder := gob.NewEncoder(writer)

	if g.Download != nil {
		g.Download.progressLock.Lock()
	}

	err = encoder.Encode(g)
	if err != nil {
		return err
	}

	if g.Download != nil {
		g.Download.progressLock.Unlock()
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	util.Logger.Debugf("Successfully outputted game data to %s", filename)
	return nil
}

// LoadGames load all the game data stored in storage
func LoadGames(gameDataLocation string) ([]*Game, error) {

	// does the directory to load from exist?
	_, err := os.Stat(gameDataLocation)
	if err != nil {
		return nil, err
	}

	// Parse the games
	var games []*Game

	dir, err := os.Open(gameDataLocation)
	if err != nil {
		return nil, err
	}

	gameList, err := dir.Readdirnames(0)
	dir.Close()

	if err != nil {
		return nil, err
	}

	for _, game := range gameList {
		f, err := os.Stat(filepath.Join(gameDataLocation, game))
		if err != nil || f.IsDir() {
			continue
		}

		gameFile, err := os.Open(filepath.Join(gameDataLocation, game))
		if err != nil {
			continue
		}

		reader, err := gzip.NewReader(gameFile)
		if err != nil {
			gameFile.Close()
			continue
		}

		decoder := gob.NewDecoder(reader)

		var gm Game
		err = decoder.Decode(&gm)
		if err != nil {
			gameFile.Close()
			continue
		}

		if gm.Download != nil {
			gm.Download.inserterPool = shardInserterPool(5, gm.Download)
		}

		gameFile.Close()
		games = append(games, &gm)
	}

	return games, nil
}

// Serialisation

// Serialise Turns a game into a base64 encoded, gzip compressed byte stream
func (g *Game) Serialise() (string, error) {

	// encode
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)

	err := e.Encode(*g)
	if err != nil {
		return "", err
	}

	// compress
	compressed := bytes.Buffer{}
	compressor := gzip.NewWriter(&compressed)

	compressor.Write(b.Bytes())
	compressor.Close()

	return base64.StdEncoding.EncodeToString(compressed.Bytes()), nil
}

// DeserializeGame Takes a serialised game and turns it into a struct
func DeserializeGame(data string) (*Game, error) {
	g := &Game{}

	// from base64
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	b := bytes.Buffer{}
	_, err = b.Write(decodedData)
	if err != nil {
		return nil, err
	}

	// decompress
	decompressor, err := gzip.NewReader(&b)
	if err != nil {
		return nil, err
	}
	decompressor.Close()

	decompressed := bytes.Buffer{}
	io.Copy(&decompressed, decompressor)

	// decode
	d := gob.NewDecoder(&decompressed)
	err = d.Decode(g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Equals compare two games
func (g *Game) Equals(g2 *Game) bool {
	return g.Title == g2.Title &&
		g.Version == g2.Version &&
		g.Developer == g2.Developer &&
		g.ReleaseDate == g2.ReleaseDate &&
		bytes.Equal(g.RootHash[:], g2.RootHash[:])

}

// FetchShard Fetch the shard for a given hash
func (g *Game) FetchShard(hash [32]byte) (bool, []byte, error) {
	hashtree, err := g.GetData()
	if err != nil {
		return false, nil, err
	}

	found, data, err := hashtree.GetShard(hash)
	if err != nil {
		return false, nil, err
	}

	if !found {
		return false, nil, nil
	}

	return true, data, nil
}

// GetDownload Get a games download
func (g *Game) GetDownload() *Download {
	return g.Download
}
