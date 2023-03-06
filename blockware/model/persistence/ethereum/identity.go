package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

/*

When a user buys a game, their address is stored in our contract.
Each address will have a corresponding public/private key pair and we will
need to use this to verify whether a user is authorised to access the game's
content.

To verify that someone owns the private key they will need to perform a handshake:

VERIFIER: requests a signed piece of data
PROVER: responds with the signed message
VERIFIER: checks the signed message with the expected content

*/

const (

	// how long an address has to respond after being sent a verification
	ADDR_VALIDATION_TIMEOUT_AFTER time.Duration = 5 * time.Minute
)

/*
This type is used to track the progress of another party verifying
their address using a signed piece of data
*/
type addressValidator struct {
	// the public key the other party claims to have
	PublicKey *ecdsa.PublicKey

	// the message that is sent to confirm
	Message []byte

	// the signature received
	Signature []byte

	// When the validation has to be responded to by
	// otherwise the validation will have to be restarted
	Expiry time.Time
}

// start a new validation
func GenerateAddressValidation(pubKey *ecdsa.PublicKey) *addressValidator {
	return &addressValidator{
		PublicKey: pubKey,
		Message:   []byte(time.Now().String()),
		Expiry:    time.Now().Add(ADDR_VALIDATION_TIMEOUT_AFTER),
	}
}

// produce a validation for a given message
func ProduceAddressValidation(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	if len(message) == 0 {
		return nil, fmt.Errorf("validation message should not be empty")
	}

	hash := crypto.Keccak256Hash(message)
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// check whether a received signature matches an expected one
func CheckAddressValidation(validator *addressValidator, receivedSig []byte) (bool, error) {
	if validator.Expiry.Before(time.Now()) {
		return false, fmt.Errorf("validation expired on %s", validator.Expiry.String())
	}

	validator.Signature = receivedSig
	return crypto.VerifySignature(
		crypto.FromECDSAPub(validator.PublicKey),
		crypto.Keccak256Hash(validator.Message).Bytes(),
		receivedSig[:len(receivedSig)-1]), nil
}
