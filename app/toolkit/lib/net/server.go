package net

import (
	"bufio"
	"fmt"
	"io"
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

func (s *TCPServer) Start(onMessage func([]string, PeerIT)) error {
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

func (s *TCPServer) listen(onMessage func([]string, PeerIT)) {
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

		p := GetPeerInstance()
		p.onConnection(con.RemoteAddr().String(), 0000, client)

		go client.listen(onMessage)
	}
}

func (c *TCPServerClient) listen(onMessage func([]string, PeerIT)) {

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

func (c *TCPServerClient) Send(command []byte) error {
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

func (c *TCPServerClient) SendString(command string) error {
	log.Printf("Sending %s\n", command)
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

func (c *TCPServerClient) Info() string {
	return fmt.Sprintf("%s", c.con.RemoteAddr())
}
