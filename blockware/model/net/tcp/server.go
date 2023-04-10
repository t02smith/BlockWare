package tcp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*
*

A TCP server will manage many concurrent TCP connections made
to it and listen in for any messages they send.
*/
type TCPServer struct {

	// server details
	hostname string
	port     uint

	// incoming connection details
	listener net.Listener
	clients  []*TCPServerClient
}

/*
Represents a single TCP connection maintained by the server.
Messages from this will be listened for on their own process

TCPServerClient implements the PeerIT interface
*/
type TCPServerClient struct {
	con    net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

//

// generate a new server object
func InitServer(hostname string, port uint) *TCPServer {
	return &TCPServer{
		hostname: hostname,
		port:     port,
		listener: nil,
		clients:  []*TCPServerClient{},
	}
}

// start a new TCP server and listen for incoming connections
func (s *TCPServer) Start(
	onMessage func([]string, TCPConnection) error,
	onConnection func(string, uint, TCPConnection),
	onClose func(TCPConnection)) error {

	util.Logger.Infof("Starting server on %s:%d", s.hostname, s.port)

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.hostname, s.port))
	if err != nil {
		util.Logger.Errorf("Error starting server: %s", err)
		return err
	}

	s.listener = ln
	util.Logger.Infof("Server listening on %s:%d", s.hostname, s.port)
	s.listen(onMessage, onConnection, onClose)
	util.Logger.Infof("server started")
	return nil
}

// listen for incoming connections and setup processes to listen
// for incoming messages from them
func (s *TCPServer) listen(
	onMessage func([]string, TCPConnection) error,
	onConnection func(string, uint, TCPConnection),
	onClose func(TCPConnection)) {

	for {
		con, err := s.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}

			util.Logger.Errorf("Error connecting to client: %s", err)
			return
		}

		util.Logger.Infof("Client joined: %s", con.RemoteAddr())
		client := &TCPServerClient{
			con:    con,
			reader: bufio.NewReader(con),
			writer: bufio.NewWriter(con),
		}
		s.clients = append(s.clients, client)

		onConnection(con.RemoteAddr().String(), 0000, client)
		go client.listen(onMessage, onClose)
	}
}

// close a TCP server and all client connections
func (s *TCPServer) Close() {
	if s.listener != nil {
		s.listener.Close()
	}
	for _, c := range s.clients {
		c.Close()
	}
}

func (s *TCPServer) Clients() []*TCPServerClient {
	return s.clients
}

// listen for messages from a specific TCP connection
func (c *TCPServerClient) listen(onMessage func([]string, TCPConnection) error, onClose func(TCPConnection)) {
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

		// util.Logger.Debugf("message received %s", msg)
		err = onMessage(strings.Split(msg[:len(msg)-1], ";"), c)
		if err != nil {
			util.Logger.Warn(err)
		}
	}

	c.Close()
}

// send a string message to a given client
func (c *TCPServerClient) SendString(command string) error {
	util.Logger.Debugf("Sending %s", command)
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

// wrapper around SendString using fmt.Sprintf
func (c *TCPServerClient) SendStringf(command string, args ...any) error {
	return c.SendString(fmt.Sprintf(command, args...))
}

// get information about a given TCP connection
func (c *TCPServerClient) Info() string {
	return c.con.RemoteAddr().String()
}

// close a connection with a TCP client
func (c *TCPServerClient) Close() error {
	return c.con.Close()
}

func (s *TCPServer) IsClient(con TCPConnection) bool {
	for _, con2 := range s.clients {
		if con == con2 {
			return true
		}
	}

	return false
}
