package ethereum

import (
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

/*

function: ProduceAddressValidation
purpose: create a signed message given some data content

? Test cases
success:
	| #1 signed message produced and correct

failure:
	| illegal arguments
			| #1 invalid private key
			| #2 empty message

*/

func TestProduceTestValidation(t *testing.T) {
	privateKey := testutil.Accounts[0][1]
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		t.Fatalf("error parsing test private key %s", err)
	}
	_privateKey = privateKeyECDSA

	publicKeyBytes := crypto.FromECDSAPub(&privateKeyECDSA.PublicKey)

	t.Run("success", func(t *testing.T) {
		t.Run("msg produced and correct", func(t *testing.T) {
			message, hash := []byte("hello world"), crypto.Keccak256Hash([]byte("hello world"))
			sig, err := ProduceAddressValidation(message)
			assert.Nil(t, err, "no error expected")

			assert.True(t, crypto.VerifySignature(publicKeyBytes, hash.Bytes(), sig[:len(sig)-1]), "public keys do not match")
		})

	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("empty message", func(t *testing.T) {
				_, err := ProduceAddressValidation([]byte{})
				assert.NotNil(t, err, "error message expected")
				assert.Equal(t, "validation message should not be empty", err.Error(), "invalid error message")
			})

			t.Run("invalid private key", func(t *testing.T) {
				// TODO
			})
		})
	})

}

/*

function: CheckAddressValidation
purpose: Verify that a given signature matches the expected one

? Test cases
success
	| #1 expected signature sent
	| #2 incorrect signature sent
					| #1 invalid public key
					| #2 invalid message

failure
	| illegal arguments
			| #1 validation is expired

*/

func TestCheckAddressValidation(t *testing.T) {
	privateKey := testutil.Accounts[0][1]
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		t.Fatalf("error parsing test private key %s", err)
	}
	_privateKey = privateKeyECDSA

	validator := GenerateAddressValidation()

	t.Run("success", func(t *testing.T) {
		t.Run("valid signature", func(t *testing.T) {
			sig, err := ProduceAddressValidation(validator.message)
			assert.Nil(t, err, "no err expected")

			res, err := CheckAddressValidation(validator, sig)
			assert.Nil(t, err, "no err expected")
			assert.True(t, res, "signature should pass validation")
		})

		t.Run("invalid signature", func(t *testing.T) {
			t.Run("invalid key", func(t *testing.T) {
				privKeyII, err := crypto.HexToECDSA(testutil.Accounts[1][1])
				if err != nil {
					t.Fatalf("error parsing test private key %s", err)
				}

				oldPrivKey := _privateKey
				_privateKey = privKeyII
				t.Cleanup(func() {
					_privateKey = oldPrivKey
				})

				sig, err := ProduceAddressValidation(validator.message)
				assert.Nil(t, err, "no err expected")

				res, err := CheckAddressValidation(validator, sig)
				assert.Nil(t, err, "no err expected")
				assert.False(t, res, "should be rejected => invalid signature")
			})

			t.Run("invalid message", func(t *testing.T) {
				sig, err := ProduceAddressValidation([]byte("asuidghasgdasddash"))
				assert.Nil(t, err, "no err expected")

				res, err := CheckAddressValidation(validator, sig)
				assert.Nil(t, err, "no err expected")
				assert.False(t, res, "should be rejected => invalid message")
			})
		})
	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("expired validation", func(t *testing.T) {
				originalTime := validator.expiry
				validator.expiry = time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC)
				t.Cleanup(func() {
					validator.expiry = originalTime
				})

				res, err := CheckAddressValidation(validator, validator.message)
				assert.False(t, res, "output should be false")
				assert.NotNil(t, err, "Expected error message about expiry")
				assert.Equal(t, fmt.Sprintf("validation expired on %s", validator.expiry.String()), err.Error(), "incorrect err message")
			})
		})
	})

}
