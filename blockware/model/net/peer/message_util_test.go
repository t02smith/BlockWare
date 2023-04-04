package peer

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

function stringTo32ByteArr
purpose: turns a hex string to a 32 byte array

? Test cases
success
	#1 => correct hash

failure
	#1 => invalid length
	#2 => invalid hex string

*/

func TestStringTo32ByteArr(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("hello"))
		hex := fmt.Sprintf("%x", hash)

		res, err := stringTo32ByteArr(hex)
		assert.Nil(t, err)
		assert.True(t, bytes.Equal(res[:], hash[:]))

	})

	t.Run("failure", func(t *testing.T) {
		t.Run("invalid length", func(t *testing.T) {
			res, err := stringTo32ByteArr("0121123124")
			assert.Empty(t, res)
			assert.NotNil(t, err)
			assert.Equal(t, "invalid hash length for hash 0121123124", err.Error())
		})

		t.Run("invalid hex string", func(t *testing.T) {
			res, err := stringTo32ByteArr(strings.Repeat("X", 64))
			assert.Empty(t, res)
			assert.NotNil(t, err)
		})
	})
}
