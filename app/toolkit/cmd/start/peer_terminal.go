package start

import (
	"fmt"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/t02smith/part-iii-project/toolkit/cmd/util"
	"github.com/t02smith/part-iii-project/toolkit/lib/net"
)

// commands

type menuOption struct {
	Label    string
	Function func(*net.Peer)
}

type messageOption struct {
	Label   string
	Message string
}

func menu(p *net.Peer) {

	options := []menuOption{
		{
			Label:    "Connect",
			Function: connect,
		},
		{
			Label:    "Message",
			Function: message,
		},
		{
			Label:    "Exit",
			Function: exit,
		},
	}

	template := &promptui.SelectTemplates{
		Label:    "{{ .Label }}?",
		Active:   "\U0001F336 {{ .Label | cyan }}",
		Inactive: "{{ .Label | cyan }}",
		Selected: "\U0001F336 {{ .Label | green | cyan }}",
	}

	menuSelect := promptui.Select{
		Label:     "\nSelect an Option",
		Items:     options,
		Templates: template,
	}

	serveHost, serverPort := p.GetServerInfo()

	for {
		fmt.Printf("Welcome %s%s%s:%d%s!\n", util.ColourCyan, serveHost, util.ColourYellow, serverPort, util.ColourReset)
		i, _, err := menuSelect.Run()
		if err != nil {
			log.Printf("Error running cmd %s: %s\n", options[i].Label, err)
			continue
		}

		if options[i].Label == "Exit" {
			break
		}

		options[i].Function(p)
	}

	log.Println("Shutting down...")

}

// util

func choosePeer(p *net.Peer) (*net.PeerData, error) {

	template := &promptui.SelectTemplates{
		Label:    "{{ .Hostname }}:{{ .Port }}?",
		Active:   "\U0001F336 {{ .Hostname | yellow }}:{{ .Port | cyan }}",
		Inactive: "\U0001F336 {{ .Hostname | cyan }}:{{ .Port | cyan }}",
		Selected: "\U0001F336 {{ .Hostname | green }}:{{ .Port | green }}",
	}

	peersLs := p.GetPeers()
	peers := promptui.Select{
		Label:     "Choose a peer to send the message to",
		Items:     peersLs,
		Templates: template,
	}

	i, _, err := peers.Run()
	if err != nil {
		return nil, err
	}

	return peersLs[i], nil
}

// options

// connect to a new peer
func connect(p *net.Peer) {
	fmt.Println("Connect to another peer")

	host := promptui.Prompt{
		Label:   "Enter their hostname",
		Default: "localhost",
	}

	port := promptui.Prompt{
		Label:   "Enter their port",
		Default: "3047",
		Validate: func(s string) error {
			if _, err := strconv.ParseUint(s, 10, 32); err != nil {
				return err
			}
			return nil
		},
	}

	hostname, err := host.Run()
	if err != nil {
		log.Println("Failed to read hostname")
		return
	}

	portNo, err := port.Run()
	if err != nil {
		log.Println("Error reading port")
		return
	}

	portNoUint, err := strconv.ParseUint(portNo, 10, 32)
	if err != nil {
		log.Println("Error reading port")
		return
	}

	err = p.ConnectToPeer(hostname, uint(portNoUint))
	if err != nil {
		log.Printf("Error connecting to peer: %s\n", err)
	}
}

// send a message to an existing peer
func message(p *net.Peer) {

	// what message to send
	_ = []messageOption{
		{
			Label:   "View a peer's library",
			Message: "LIBRARY",
		},
	}

	// who to send the message to
	chosen, err := choosePeer(p)
	if err != nil {
		log.Printf("Error choosing a peer: %s\n", err)
		return
	}

	chosen.Peer.SendString("LIBRARY\n")
}

func exit(p *net.Peer) {

}
