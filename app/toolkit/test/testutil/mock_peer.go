package testutil

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// a mock peer will form a connection with a peer for testing
type MockPeer struct {

	// network info
	reader *bufio.Reader
	writer *bufio.Writer

	// given input -> expected output
	responses       map[string]string
	defaultResponse string

	// memory of all sent messages
	msgHistory []string
}

// create a new mock peer
func StartMockPeer(peerPort uint) (*MockPeer, error) {

	con, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", peerPort))
	if err != nil {
		return nil, err
	}

	mp := &MockPeer{
		reader:          bufio.NewReader(con),
		writer:          bufio.NewWriter(con),
		responses:       make(map[string]string),
		defaultResponse: "",
		msgHistory:      []string{},
	}

	go mp.listen()
	return mp, nil
}

// listen for incoming messages
func (m *MockPeer) listen() {
	for {
		msg, err := m.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}

			log.Printf("error reading message to mock peer %s", err)
			continue
		}

		log.Printf("Mock peer received message %s", msg)
		m.msgHistory = append(m.msgHistory, msg)

		if res, ok := m.responses[msg]; ok {
			m.writer.WriteString(res)
		} else {
			m.writer.WriteString(m.defaultResponse)
		}

		m.writer.Flush()
	}
}

// add a response to a given input
func (m *MockPeer) SetResponse(input, output string) {
	m.responses[input] = output
}

// clear all existing set responses and message history
func (m *MockPeer) Clear() {
	m.responses = make(map[string]string)
	m.msgHistory = []string{}
}

// add a default response for no input
func (m *MockPeer) SetDefaultResponse(res string) {
	m.defaultResponse = res
}

// get the last sent message to the mock peer
func (m *MockPeer) GetLastMessage() string {
	if len(m.msgHistory) == 0 {
		return ""
	}

	return m.msgHistory[len(m.msgHistory)-1]
}
