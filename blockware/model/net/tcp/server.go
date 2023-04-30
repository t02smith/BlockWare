package tcp

import (
	"bufio"
	"errors"
	"fmt"
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

	// functions
	onMessage    func([]string, TCPConnection) error
	onConnection func(string, uint, TCPConnection)
	onClose      func(TCPConnection)
}

/*
Represents a single TCP connection maintained by the server.
Messages from this will be listened for on their own process

TCPServerClient implements the PeerIT interface
*/
type TCPServerClient struct {
	con net.Conn

	// communication channels
	reader *bufio.Reader
	writer *bufio.Writer

	// reference to server
	server *TCPServer
	closed bool
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
	s.onMessage = onMessage
	s.onConnection = onConnection
	s.onClose = onClose

	util.Logger.Infof("Server listening on %s:%d", s.hostname, s.port)
	s.listen()
	util.Logger.Infof("server started")
	return nil
}

// listen for incoming connections and setup processes to listen
// for incoming messages from them
func (s *TCPServer) listen() {
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
			closed: false,
			server: s,
		}
		s.clients = append(s.clients, client)

		s.onConnection(con.RemoteAddr().String(), 0000, client)
		go client.listen()
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
func (c *TCPServerClient) listen() {
	for {
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			util.Logger.Warnf("Malformed message from client: %s", err)
			break
		}

		util.Logger.Debugf("message received from %s: %s", c.Info(), msg[:len(msg)-1])
		err = c.server.onMessage(strings.Split(msg[:len(msg)-1], ";"), c)
		if err != nil {
			util.Logger.Warn(err)
		}
	}

	c.server.onClose(c)
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
	if c.closed {
		return nil
	}

	c.closed = true
	return c.con.Close()
}

// get the hostname and port of the server
func (s *TCPServer) GetServerInfo() (string, uint) {
	return s.hostname, s.port
}

// update handlers
func (s *TCPServer) SetOnMessage(onMessage func([]string, TCPConnection) error) {
	s.onMessage = onMessage
}
