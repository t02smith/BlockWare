package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*
A TCP client will form a TCP connection with a server
socket and listen in for incoming messages from it.

TCPClient implements the PeerIT interface to allow
for an abstraction by the Peer object.
*/
type TCPClient struct {

	// server details
	hostname string
	port     uint

	// server connection objects
	con    net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

// generate a new TCP client to conenct to a server
func InitTCPClient(
	serverHostname string,
	serverPort uint,
	onMessage func([]string, TCPConnection),
	onConnection func(string, uint, TCPConnection),
	onClose func(TCPConnection)) (*TCPClient, error) {
	util.Logger.Infof("Attempting to open connection to %s:%d", serverHostname, serverPort)
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverHostname, serverPort))
	if err != nil {
		util.Logger.Errorf("Error connecting to %s:%d: %s", serverHostname, serverPort, err)
		return nil, err
	}

	client := &TCPClient{
		hostname: serverHostname,
		port:     serverPort,
		con:      con,
		reader:   bufio.NewReader(con),
		writer:   bufio.NewWriter(con),
	}

	onConnection(serverHostname, serverPort, client)

	go client.listen(onMessage, onClose)
	return client, nil
}

// listen for messages from the server
// onMessage is a handler that is called when a message is received
func (c *TCPClient) listen(onMessage func([]string, TCPConnection), onClose func(TCPConnection)) {
	for {
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				onClose(c)
				return
			}

			util.Logger.Warnf("Malformed message from client: %s", err)
			break
		}

		util.Logger.Infof("message received %s\n", msg)
		onMessage(strings.Split(msg[:len(msg)-1], ";"), c)
	}
}

// Send a message in bytes to the server
func (c *TCPClient) Send(command []byte) error {
	_, err := c.writer.Write(command)
	if err != nil {
		util.Logger.Errorf("Error sending message %s", err)
		return err
	}

	err = c.writer.Flush()
	if err != nil {
		util.Logger.Errorf("Error sending message %s", err)
		return err
	}

	return nil
}

// Send a string message to the server
func (c *TCPClient) SendString(command string) error {
	_, err := c.writer.WriteString(command)
	if err != nil {
		util.Logger.Errorf("Error sending message %s", err)
		return err
	}

	err = c.writer.Flush()
	if err != nil {
		util.Logger.Errorf("Error sending message %s", err)
		return err
	}

	return nil
}

// Send a string message with parameters
// wrapper function using fmt.Sprintf
func (c *TCPClient) SendStringf(command string, args ...any) error {
	return c.SendString(fmt.Sprintf(command, args...))
}

// get information about the server connection
func (c *TCPClient) Info() string {
	return fmt.Sprintf("%s:%d", c.hostname, c.port)
}

// close the tcp client
func (c *TCPClient) Close() {
	c.con.Close()
}
