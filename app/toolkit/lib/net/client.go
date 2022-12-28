package net

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type TCPClient struct {

	// server details
	hostname string
	port     uint

	// server connection objects
	con    net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func InitTCPClient(serverHostname string, serverPort uint) (*TCPClient, error) {
	log.Printf("Attempting to open connection to %s:%d", serverHostname, serverPort)
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverHostname, serverPort))
	if err != nil {
		log.Printf("Error connecting to %s:%d: %s", serverHostname, serverPort, err)
		return nil, err
	}

	client := &TCPClient{
		hostname: serverHostname,
		port:     serverPort,
		con:      con,
		reader:   bufio.NewReader(con),
		writer:   bufio.NewWriter(con),
	}

	return client, nil
}

func (c *TCPClient) Send(command []byte) error {
	_, err := c.writer.Write(command)
	if err != nil {
		log.Printf("Error sending message %s", err)
	}

	err = c.writer.Flush()
	if err != nil {
		log.Printf("Error sending message %s", err)
	}

	return nil
}
