package tcp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func startTestTCPServer(t *testing.T, port uint) *TCPServer {
	t.Helper()
	s := InitServer("localhost", port)

	go func() {
		err := s.Start(func(s []string, t TCPConnection) error {
			return nil
		}, func(s string, u uint, t TCPConnection) {}, func(t TCPConnection) {})
		if err != nil {
			panic("err starting server")
		}

		t.Cleanup(func() { s.Close() })
	}()

	return s
}

/*

function: InitServer
purpose: creates a new non-started server

? Test cases
success
	#1 => base case

*/

func TestInitServer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := InitServer("localhost", 9476)

		assert.Equal(t, "localhost", s.hostname)
		assert.Equal(t, uint(9476), s.port)
		assert.Empty(t, s.clients)
	})
}

/*

function: TCPSever.Start
purpose: start a new server listener

? Test cases
success
	#1 => server started and can be connected to

failure
	#1 => server already running on port
*/

func TestSeverStart(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := startTestTCPServer(t, 7804)

		time.Sleep(25 * time.Millisecond)
		mp, err := testutil.StartMockPeer(7804, true)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			mp.Close()
		})
		time.Sleep(25 * time.Millisecond)

		assert.Equal(t, 1, len(s.clients))
		s.clients[0].SendString("hello world\n")
		time.Sleep(25 * time.Millisecond)
		assert.Equal(t, "hello world\n", mp.GetLastMessage())
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("duplicate port", func(t *testing.T) {
			startTestTCPServer(t, 7805)
			s2 := InitServer("localhost", 7805)

			time.Sleep(25 * time.Millisecond)
			err := s2.Start(func(s []string, t TCPConnection) error { return nil }, func(s string, u uint, t TCPConnection) {}, func(t TCPConnection) {})
			assert.NotNil(t, err)
		})
	})
}

/*

function: TCPServer.listen
purpose: listen to incoming connections

? Test cases
success
	#1 => new client joined successfully
	#2 => listening to messages from client

*/

func TestTCPServerListen(t *testing.T) {
	s := InitServer("localhost", 8505)

	clientJoined := false
	msgReceived := []string{}
	go func() {
		err := s.Start(func(s []string, t TCPConnection) error {
			msgReceived = s
			return nil
		}, func(s string, u uint, t TCPConnection) { clientJoined = true }, func(t TCPConnection) {})
		if err != nil {
			return
		}

		t.Cleanup(func() { s.Close() })
	}()

	t.Run("success", func(t *testing.T) {
		mp, err := testutil.StartMockPeer(8505, true)
		assert.Nil(t, err)

		t.Run("client joined", func(t *testing.T) {
			time.Sleep(25 * time.Millisecond)
			assert.True(t, clientJoined)
			assert.Equal(t, 1, len(s.clients))
		})

		t.Run("listening to messages", func(t *testing.T) {
			mp.SendStringAndWait(25, "hello world\n")
			assert.Equal(t, "hello world", msgReceived[0])
		})
	})
}

/*

function: TCPServerClient.SendString
purpose: send a message to a server's client

? Test cases
success
	#1 => the message is sent
	#2 => the wrapper function SendStringf

*/

func TestTCPServerClientSendString(t *testing.T) {
	s := startTestTCPServer(t, 8509)

	time.Sleep(25 * time.Millisecond)
	mp, err := testutil.StartMockPeer(8509, true)
	assert.Nil(t, err)
	t.Cleanup(func() {
		mp.Close()
	})
	time.Sleep(25 * time.Millisecond)

	t.Run("success", func(t *testing.T) {
		t.Run("message sent", func(t *testing.T) {
			s.clients[0].SendString("hello world\n")
			time.Sleep(25 * time.Millisecond)
			assert.Equal(t, "hello world\n", mp.GetLastMessage())
		})

		t.Run("SendStringf", func(t *testing.T) {
			s.clients[0].SendStringf("hello world my name is %s, the year is %d\n", "tom", 2023)
			time.Sleep(25 * time.Millisecond)
			assert.Equal(t, "hello world my name is tom, the year is 2023\n", mp.GetLastMessage())
		})
	})

}

/*

function: IsClient
purpose: Is the given connection someone who connected to us

? Test cases
success
	#1 => yes
	#2 => no

*/

func TestIsClient(t *testing.T) {
	s := startTestTCPServer(t, 9045)

	t.Run("success", func(t *testing.T) {
		t.Run("yes", func(t *testing.T) {
			mp, err := testutil.StartMockPeer(9045, true)
			t.Cleanup(func() { mp.Close() })
			assert.Nil(t, err)

			time.Sleep(25 * time.Millisecond)
			assert.True(t, s.IsClient(s.clients[0]))
		})

		t.Run("no", func(t *testing.T) {
			client, err := InitTCPClient("localhost", 9045, func(s []string, t TCPConnection) error { return nil }, func(s string, u uint, t TCPConnection) {}, func(t TCPConnection) {})
			assert.Nil(t, err)
			t.Cleanup(func() { client.Close() })
			time.Sleep(25 * time.Millisecond)
			assert.False(t, s.IsClient(client))
		})
	})
}
