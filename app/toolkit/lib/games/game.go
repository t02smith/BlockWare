package games

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"

	hashIO "github.com/t02smith/part-iii-project/toolkit/lib/io"
)

type Game struct {

	// game metadata
	title       string
	version     string
	releaseDate string // iso-format
	developer   string // domain

	// the shard data
	data *hashIO.HashTree
}

// Creator

func CreateGame(title, version, releaseDate, developer, rootDir string, shardSize uint) (*Game, error) {

	// check version format
	versionMatches, err := regexp.MatchString("^(\\d+\\.)*\\d+$", version)
	if err != nil {
		log.Println("error matching version number to regex")
		return nil, err
	}

	if !versionMatches {
		log.Println("invalid version number")
		return nil, errors.New("invalid version number")
	}

	// check release date
	_, err = time.Parse("2006-01-02 15:04:05 -0700 MST", releaseDate)
	if err != nil {
		log.Println("invalid release date given")
		return nil, errors.New("invalid release date given")
	}

	// check domain has SSL certificate
	domainCorrect, err := verifyDomain(developer)
	if err != nil {
		return nil, err
	}

	if !domainCorrect {
		return nil, errors.New("invalid domain specified")
	}

	// hash data
	tree, err := hashIO.NewHashTree(rootDir, shardSize)
	if err != nil {
		return nil, err
	}

	// return value
	game := &Game{
		title:       title,
		version:     version,
		releaseDate: releaseDate,
		developer:   developer,
		data:        tree,
	}

	return game, nil
}

func verifyDomain(domain string) (bool, error) {

	conf := &tls.Config{}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", domain), conf)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return false, nil
	}

	return true, nil
}
