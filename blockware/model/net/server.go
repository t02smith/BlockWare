package net

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/**

Each peer will host its own TCP server for
communicating with other peers by exchanging
8bit commands

*/

type TCPServer struct {
	hostname string
	port     uint
	listener net.Listener
	clients  []*TCPServerClient
}

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

func (s *TCPServer) Start(onMessage func([]string, PeerIT)) error {
	util.Logger.Infof("Starting server on %s:%d", s.hostname, s.port)

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.hostname, s.port))
	if err != nil {
		util.Logger.Errorf("Error starting server: %s", err)
		return err
	}

	s.listener = ln
	util.Logger.Infof("Server listening on %s:%d", s.hostname, s.port)
	s.listen(onMessage)
	util.Logger.Infof("server started")
	return nil
}

func (s *TCPServer) listen(onMessage func([]string, PeerIT)) {
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

		p := Peer()
		p.onConnection(con.RemoteAddr().String(), 0000, client)

		go client.listen(onMessage)
	}
}

func (s *TCPServer) Close() {
	s.listener.Close()
	for _, c := range s.clients {
		c.Close()
	}
}

func (c *TCPServerClient) listen(onMessage func([]string, PeerIT)) {
	for {
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				Peer().onClose(c)
				return
			}

			util.Logger.Warnf("Malformed message from client: %s", err)
			continue
		}

		util.Logger.Infof("message received %s", msg)
		onMessage(strings.Split(msg[:len(msg)-1], ";"), c)
	}
}

func (c *TCPServerClient) Send(command []byte) error {
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

func (c *TCPServerClient) SendString(command string) error {
	util.Logger.Infof("Sending %s", command)
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

func (c *TCPServerClient) SendStringf(command string, args ...any) error {
	return c.SendString(fmt.Sprintf(command, args...))
}

func (c *TCPServerClient) Info() string {
	return fmt.Sprintf("%s", c.con.RemoteAddr())
}

func (c *TCPServerClient) Close() {
	c.con.Close()
}
