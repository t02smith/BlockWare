package start

import (
	"fmt"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/t02smith/part-iii-project/toolkit/cmd/util"
	"github.com/t02smith/part-iii-project/toolkit/cmd/view"
	"github.com/t02smith/part-iii-project/toolkit/lib/net"
)

// commands

type menuOption struct {
	Label    string
	Function func()
}

type messageOption struct {
	Label   string
	Message string
}

func menuTemplate(options []menuOption) (*menuOption, error) {
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

	i, _, err := menuSelect.Run()
	if err != nil {
		return nil, err
	}

	return &options[i], nil
}

func menu() {
	p := net.GetPeerInstance()

	options := []menuOption{
		{
			Label:    "Connect",
			Function: connect,
		},
		{
			Label:    "Peers",
			Function: peers,
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

	serveHost, serverPort := p.GetServerInfo()

	for {
		fmt.Printf("Welcome %s%s%s:%d%s!\n", util.ColourCyan, serveHost, util.ColourYellow, serverPort, util.ColourReset)
		res, err := menuTemplate(options)
		if err != nil {
			log.Printf("Error running cmd %s: %s\n", res.Label, err)
			continue
		}

		if res.Label == "Exit" {
			break
		}

		res.Function()
	}

	log.Println("Shutting down...")

}

// util

func choosePeer() (net.PeerIT, error) {
	p := net.GetPeerInstance()

	template := &promptui.SelectTemplates{
		Label:    "{{ .Hostname }}:{{ .Port }}?",
		Active:   "\U0001F336 {{ .Hostname | yellow }}:{{ .Port | cyan }}",
		Inactive: "{{ .Hostname | cyan }}:{{ .Port | cyan }}",
		Selected: "\U0001F336 {{ .Hostname | green }}:{{ .Port | green }}",
	}

	peersLs := []*net.PeerData{}
	for _, v := range p.GetPeers() {
		peersLs = append(peersLs, v)
	}

	peers := promptui.Select{
		Label:     "Select a peer",
		Items:     peersLs,
		Templates: template,
	}

	i, _, err := peers.Run()
	if err != nil {
		return nil, err
	}

	return peersLs[i].Peer, nil
}

// options

// connect to a new peer
func connect() {
	p := net.GetPeerInstance()
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
func message() {

	// what message to send
	msgOptions := []messageOption{
		{
			Label:   "Get library",
			Message: "LIBRARY\n",
		},
		{
			Label:   "Back to menu",
			Message: "RETURN",
		},
	}

	// who to send the message to
	chosenPeer, err := choosePeer()
	if err != nil {
		log.Printf("Error choosing a peer: %s\n", err)
		return
	}

	//
	template := &promptui.SelectTemplates{
		Label:    "{{ .Label }}",
		Active:   "\U0001F336 {{ .Label | yellow }}",
		Inactive: "{{ .Label | cyan }}",
		Selected: "\U0001F336 {{ .Label | green }}",
	}

	chooseMessagePrompt := promptui.Select{
		Label:     "Choose a peer to send the message to",
		Items:     msgOptions,
		Templates: template,
	}

	chosenMessage, _, err := chooseMessagePrompt.Run()
	if err != nil {
		log.Printf("Error choosing a message: %s\n", err)
		return
	}

	if msgOptions[chosenMessage].Message == "RETURN" {
		return
	}

	log.Printf("Sending %s", msgOptions[chosenMessage].Message)
	chosenPeer.SendString(msgOptions[chosenMessage].Message)
}

func peers() {
	chosenPeer, err := choosePeer()
	if err != nil {
		log.Printf("Error choosing a peer: %s\n", err)
		return
	}

	p := net.GetPeerInstance()
	options := []menuOption{
		{
			Label: "View library",
			Function: func() {
				view.OutputGamesTable(p.GetPeer(chosenPeer).Library)
			},
		},
	}

	res, err := menuTemplate(options)
	if err != nil {
		log.Printf("Error running cmd %s: %s\n", res.Label, err)
		return
	}

	res.Function()

}

func exit() {

}
