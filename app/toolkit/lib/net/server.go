package net

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

func (s *TCPServer) Start(onMessage func([]string, *TCPServerClient)) error {
	log.Printf("Starting server on %s:%d", s.hostname, s.port)
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.hostname, s.port))
	if err != nil {
		log.Printf("Error starting server: %s", err)
		return err
	}

	s.listener = ln
	log.Printf("Server listening on %s:%d", s.hostname, s.port)
	s.listen(onMessage)
	return nil
}

func (s *TCPServer) listen(onMessage func([]string, *TCPServerClient)) {
	for {
		con, err := s.listener.Accept()
		if err != nil {
			log.Printf("Error connecting to client: %s", err)
			continue
		}

		log.Printf("Client joined: %s", con.RemoteAddr())
		client := &TCPServerClient{
			con:    con,
			reader: bufio.NewReader(con),
			writer: bufio.NewWriter(con),
		}
		s.clients = append(s.clients, client)

		go client.listen(onMessage)
	}
}

func (c *TCPServerClient) listen(onMessage func([]string, *TCPServerClient)) {

	for {
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			log.Printf("Malformed message from client: %s", err)
			return
		}

		log.Printf("message received %s\n", msg)
		onMessage(strings.Split(msg, ";"), c)
	}
}

func (c *TCPServerClient) send(msg string) {
	c.writer.WriteString(msg)
	c.writer.Flush()
}
