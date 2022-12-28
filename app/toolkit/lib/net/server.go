package net

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

func (s *TCPServer) Start(onMessage func([]byte, *TCPServerClient)) error {
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

func (s *TCPServer) listen(onMessage func([]byte, *TCPServerClient)) {
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

func (c *TCPServerClient) listen(onMessage func([]byte, *TCPServerClient)) {
	buffer := make([]byte, 1)

	for {
		n, err := c.reader.Read(buffer)
		if err != nil {
			log.Printf("Malformed message from client: %s", err)
			return
		}

		if n == 0 {
			continue
		}

		log.Printf("message received %x", buffer)
		onMessage(buffer, c)
	}
}
