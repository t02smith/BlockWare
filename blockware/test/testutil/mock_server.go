package testutil

import (
	"bufio"
	"fmt"
	"net"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

/**

This will accept incoming requests and send predefined responses
back. This will be used for testing communications between peers

*/

// a mock server to form a connection with
type MockServer struct {
	listener net.Listener

	// the connection with the connected peer
	con    net.Conn
	reader *bufio.Reader
	writer *bufio.Writer

	// mocked responses input -> output
	responses      map[string]string
	messageHistory []string
}

var (
	MockServers map[uint]*MockServer = make(map[uint]*MockServer)
)

// close all open mock servers
func CloseAllMockServers() {
	util.Logger.Info("Closing all mock servers")
	for _, ms := range MockServers {
		ms.Close()
	}

	MockServers = make(map[uint]*MockServer)
}

// generate a collection of mock servers
func StartMockServers(ports []uint) error {
	for _, p := range ports {
		if _, ok := MockServers[p]; ok {
			util.Logger.Warnf("Duplicate port for mock server %d", p)
		}

		_, err := StartMockServer(p)
		if err != nil {
			CloseAllMockServers()
			return err
		}
	}

	return nil
}

// start a new mock server on a given port
func StartMockServer(port uint) (*MockServer, error) {
	util.Logger.Infof("Setting up mock server on port %d", port)
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, err
	}

	ms := &MockServer{
		listener:       ln,
		responses:      make(map[string]string),
		messageHistory: []string{},
	}

	go func() {
		con, err := ms.listener.Accept()
		defer ms.Close()
		if err != nil {
			util.Logger.Error(err)
			return
		}

		ms.con = con
		ms.reader = bufio.NewReader(con)
		ms.writer = bufio.NewWriter(con)

		for {
			util.Logger.Infof("Mock server on port %d listening", port)
			msg, err := ms.reader.ReadString('\n')
			if err != nil {
				util.Logger.Error(err)
				return
			}

			ms.messageHistory = append(ms.messageHistory, msg)
			if reply, ok := ms.responses[msg]; ok {
				ms.writer.WriteString(reply)
			} else {
				ms.writer.WriteString("ERROR;")
			}
			ms.writer.Flush()
		}
	}()

	MockServers[port] = ms
	return ms, nil
}

// close the mock server
func (ms *MockServer) Close() {
	util.Logger.Info("Closing mock server")
	ms.con.Close()
	ms.listener.Close()
}

// add a response to a given input
func (ms *MockServer) SetResponse(input, output string) {
	ms.responses[input] = output
}

// send a string to the peer
func (ms *MockServer) SendString(msg string, args ...any) {
	ms.writer.WriteString(fmt.Sprintf(msg, args...))
	ms.writer.Flush()
}

// get the last sent message to the mock peer
func (ms *MockServer) GetLastMessage() string {
	if len(ms.messageHistory) == 0 {
		return ""
	}

	return ms.messageHistory[len(ms.messageHistory)-1]
}
