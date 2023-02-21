package testutil

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// a mock peer will form a connection with a peer for testing
type MockPeer struct {
	con net.Conn

	// network info
	reader *bufio.Reader
	writer *bufio.Writer

	// given input -> expected output
	responses map[string]string
	onReceive map[string]func()

	defaultResponse string

	// memory of all sent messages
	msgHistory []string
}

var (
	MockPeers []*MockPeer = []*MockPeer{}
)

// create a new mock peer
func StartMockPeer(peerPort uint, connect bool) (*MockPeer, error) {
	mp := &MockPeer{
		responses:       make(map[string]string),
		onReceive:       make(map[string]func()),
		defaultResponse: "",
		msgHistory:      []string{},
	}
	if connect {
		mp.Connect(peerPort)
	}

	MockPeers = append(MockPeers, mp)
	return mp, nil
}

func (m *MockPeer) Connect(peerPort uint) error {
	con, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", peerPort))
	if err != nil {
		return err
	}

	m.con = con
	m.reader = bufio.NewReader(con)
	m.writer = bufio.NewWriter(con)

	go m.listen()
	return nil
}

// listen for incoming messages
func (m *MockPeer) listen() {
	for {
		msg, err := m.reader.ReadString('\n')
		if err != nil {
			return
		}

		log.Printf("Mock peer received message %s", msg)
		m.msgHistory = append(m.msgHistory, msg)

		if res, ok := m.responses[msg]; ok {
			log.Printf("Sending %s", res)
			m.writer.WriteString(res)
		} else {
			log.Printf("Sending default response %s", m.defaultResponse)
			m.writer.WriteString(m.defaultResponse)
		}
		m.writer.Flush()

		if res, ok := m.onReceive[msg]; ok {
			res()
		}

	}
}

// add a response to a given input
func (m *MockPeer) SetResponse(input, output string) {
	m.responses[input] = output
}

func (m *MockPeer) SetOnReceive(input string, function func()) {
	m.onReceive[input] = function
}

// clear all existing set responses and message history
func (m *MockPeer) Clear() {
	m.responses = make(map[string]string)
	m.onReceive = make(map[string]func())
	m.msgHistory = []string{}
}

// add a default response for no input
func (m *MockPeer) SetDefaultResponse(res string) {
	m.defaultResponse = res
}

// send a byte arr to the peer
func (m *MockPeer) SendBytes(bytes []byte) {
	m.writer.Write(bytes)
	m.writer.Flush()
}

// send a string to the peer
func (m *MockPeer) SendString(msg string, args ...any) {
	m.writer.WriteString(fmt.Sprintf(msg, args...))
	m.writer.Flush()
}

// send a string and wait a given amount of time
func (m *MockPeer) SendStringAndWait(delay int, msg string, args ...any) {
	m.SendString(msg, args...)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

// get the last sent message to the mock peer
func (m *MockPeer) GetLastMessage() string {
	if len(m.msgHistory) == 0 {
		return ""
	}

	return m.msgHistory[len(m.msgHistory)-1]
}

func (m *MockPeer) Close() {
	m.con.Close()
}
