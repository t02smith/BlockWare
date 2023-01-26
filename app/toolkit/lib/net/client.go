package net

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

	p := GetPeerInstance()
	p.onConnection(serverHostname, serverPort, client)

	go client.listen(onMessage)
	return client, nil
}

func (c *TCPClient) listen(onMessage func([]string, PeerIT)) {
	for {
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				GetPeerInstance().onClose(c)
				return
			}

			log.Printf("Malformed message from client: %s", err)
			continue
		}

		log.Printf("message received %s\n", msg)
		onMessage(strings.Split(msg[:len(msg)-1], ";"), c)
	}
}

func (c *TCPClient) Send(command []byte) error {
	_, err := c.writer.Write(command)
	if err != nil {
		log.Printf("Error sending message %s", err)
		return err
	}

	err = c.writer.Flush()
	if err != nil {
		log.Printf("Error sending message %s", err)
		return err
	}

	return nil
}

func (c *TCPClient) SendString(command string) error {
	_, err := c.writer.WriteString(command)
	if err != nil {
		log.Printf("Error sending message %s", err)
		return err
	}

	err = c.writer.Flush()
	if err != nil {
		log.Printf("Error sending message %s", err)
		return err
	}

	return nil
}

func (c *TCPClient) Info() string {
	return fmt.Sprintf("%s:%d", c.hostname, c.port)
}
