package tcp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

/*

function: InitTCPClient
purpose: start a tcp client

? Test cases
success
	#1 => connection made

failure
	#1 => can't connect to host
*/

func TestInitTCPClient(t *testing.T) {
	s, err := testutil.StartMockServer(4045)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { s.Close() })

	t.Run("success", func(t *testing.T) {
		var received []string
		c, err := InitTCPClient("localhost", 4045, func(s []string, t TCPConnection) error {
			received = s
			return nil
		}, func(s string, u uint, t TCPConnection) {}, func(t TCPConnection) {})

		t.Run("fields created", func(t *testing.T) {
			assert.Equal(t, "localhost", c.hostname)
			assert.Equal(t, uint(4045), c.port)
			assert.NotNil(t, c.con)
			assert.NotNil(t, c.reader)
			assert.NotNil(t, c.writer)
			assert.False(t, c.closed)
		})

		t.Run("test message", func(t *testing.T) {
			assert.Nil(t, err)
			time.Sleep(25 * time.Millisecond)

			c.SendString("hello world\n")
			time.Sleep(25 * time.Millisecond)
			assert.Equal(t, "hello world\n", s.GetLastMessage())

			s.SendString("hello world II\n")
			time.Sleep(25 * time.Millisecond)
			assert.Equal(t, "ERROR", received[0])
		})

	})

	t.Run("failure", func(t *testing.T) {
		_, err := InitTCPClient("localhost", 9999, func(s []string, t TCPConnection) error { return nil }, func(s string, u uint, t TCPConnection) {}, func(t TCPConnection) {})
		assert.NotNil(t, err)
	})
}

/*

function: TCPClient.listen
purpose: listen to an connection channel

? Test cases
success
	#1 => message received
	#1 => connection closed

*/

func TestTCPClientListen(t *testing.T) {
	s, err := testutil.StartMockServer(4046)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { s.Close() })

	t.Run("success", func(t *testing.T) {
		var received []string
		closed := false

		c, err := InitTCPClient("localhost", 4046, func(s []string, t TCPConnection) error {
			received = s
			return nil
		}, func(s string, u uint, t TCPConnection) {}, func(t TCPConnection) {
			closed = true
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { c.Close() })

		t.Run("message received", func(t *testing.T) {
			s.SendString("this is a test message\n")
			time.Sleep(25 * time.Millisecond)
			assert.Equal(t, "this is a test message", received[0])
		})

		t.Run("connection closed", func(t *testing.T) {
			s.Close()
			time.Sleep(25 * time.Millisecond)
			assert.True(t, closed)
		})
	})
}
