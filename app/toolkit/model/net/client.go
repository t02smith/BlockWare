package net

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/t02smith/part-iii-project/toolkit/model"
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
	model.Logger.Infof("Attempting to open connection to %s:%d", serverHostname, serverPort)
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverHostname, serverPort))
	if err != nil {
		model.Logger.Errorf("Error connecting to %s:%d: %s", serverHostname, serverPort, err)
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

			model.Logger.Warnf("Malformed message from client: %s", err)
			continue
		}

		model.Logger.Infof("message received %s\n", msg)
		onMessage(strings.Split(msg[:len(msg)-1], ";"), c)
	}
}

func (c *TCPClient) Send(command []byte) error {
	_, err := c.writer.Write(command)
	if err != nil {
		model.Logger.Errorf("Error sending message %s", err)
		return err
	}

	err = c.writer.Flush()
	if err != nil {
		model.Logger.Errorf("Error sending message %s", err)
		return err
	}

	return nil
}

func (c *TCPClient) SendString(command string) error {
	_, err := c.writer.WriteString(command)
	if err != nil {
		model.Logger.Errorf("Error sending message %s", err)
		return err
	}

	err = c.writer.Flush()
	if err != nil {
		model.Logger.Errorf("Error sending message %s", err)
		return err
	}

	return nil
}

func (c *TCPClient) SendStringf(command string, args ...any) error {
	return c.SendString(fmt.Sprintf(command, args...))
}

func (c *TCPClient) Info() string {
	return fmt.Sprintf("%s:%d", c.hostname, c.port)
}

func (c *TCPClient) Close() {
	c.con.Close()
}