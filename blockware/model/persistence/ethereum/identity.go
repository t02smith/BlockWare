package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/util"

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
	// AddrValidationTimeoutAfter how long an address has to respond after being sent a verification
	AddrValidationTimeoutAfter time.Duration = 5 * time.Minute
)

// AddressValidator /*
type AddressValidator struct {
	// the public key the other party claims to have
	PublicKey *ecdsa.PublicKey

	// the message that is sent to confirm
	message []byte

	// the signature received
	signature []byte

	// When the validation has to be responded to by
	// otherwise the validation will have to be restarted
	expiry time.Time

	// signature valid
	valid bool
}

// GenerateAddressValidation start a new validation
func GenerateAddressValidation() *AddressValidator {
	util.Logger.Info("Generating address validation")
	return &AddressValidator{
		message: []byte(time.Now().String()),
		expiry:  time.Now().Add(AddrValidationTimeoutAfter),
		valid:   false,
	}
}

// ProduceAddressValidation produce a validation for a given message
func ProduceAddressValidation(message []byte) ([]byte, error) {
	if len(message) == 0 {
		return nil, fmt.Errorf("validation message should not be empty")
	}

	util.Logger.Info("Generating signature")
	hash := crypto.Keccak256Hash(message)
	signature, err := crypto.Sign(hash.Bytes(), _privateKey)
	if err != nil {
		return nil, err
	}

	util.Logger.Info("Generated signature")
	return signature, nil
}

// CheckAddressValidation check whether a received signature matches an expected one
func CheckAddressValidation(validator *AddressValidator, receivedSig []byte) (bool, error) {
	util.Logger.Info("Checking signature")
	if validator.expiry.Before(time.Now()) {
		return false, fmt.Errorf("validation expired on %s", validator.expiry.String())
	}

	pubKey, err := crypto.SigToPub(crypto.Keccak256Hash(validator.Message()).Bytes(), receivedSig)
	if err != nil {
		return false, err
	}

	validator.PublicKey = pubKey
	validator.signature = receivedSig
	validator.valid = crypto.VerifySignature(
		crypto.FromECDSAPub(validator.PublicKey),
		crypto.Keccak256Hash(validator.message).Bytes(),
		receivedSig[:len(receivedSig)-1])

	util.Logger.Infof("Signature valid: %v", validator.valid)
	return validator.valid, nil
}

// getter for message field
func (a *AddressValidator) Message() []byte {
	return a.message
}

// getter for Valild field
func (a *AddressValidator) Valid() bool {
	return a.valid
}

// force set a validation
func (a *AddressValidator) SetValid(value bool) {
	a.valid = value
}
