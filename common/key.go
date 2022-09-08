package common

import (
	"crypto/ecdsa"
	"encoding/hex"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Key represents an Ethereum public-private keypair.
type Key struct {
	priv    *ecdsa.PrivateKey
	pub     ecdsa.PublicKey
	address ethcommon.Address
}

// GenerateKey returns a new randomly-generated *Key.
func GenerateKey() (*Key, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return &Key{
		priv:    priv,
		pub:     priv.PublicKey,
		address: crypto.PubkeyToAddress(priv.PublicKey),
	}, nil
}

// NewKeyFromPrivateKeyString returns a new *Key given a hex-encoded private key.
func NewKeyFromPrivateKeyString(pk string) (*Key, error) {
	pkBytes, err := hex.DecodeString(pk)
	if err != nil {
		return nil, err
	}

	priv, err := crypto.ToECDSA(pkBytes)
	if err != nil {
		return nil, err
	}

	return &Key{
		priv:    priv,
		pub:     priv.PublicKey,
		address: crypto.PubkeyToAddress(priv.PublicKey),
	}, nil
}

func NewKeyFromPrivateKey(priv *ecdsa.PrivateKey) *Key {
	return &Key{
		priv:    priv,
		pub:     priv.PublicKey,
		address: crypto.PubkeyToAddress(priv.PublicKey),
	}
}

// Sign signs the given digest and returns a 65-byte signature in (r,s,v) format.
func (k *Key) Sign(digest [32]byte) ([]byte, error) {
	sig, err := crypto.Sign(digest[:], k.priv)
	if err != nil {
		return nil, err
	}

	// Ethereum wants 27/28 for v
	sig[64] += 27
	return sig, nil
}

// Address returns the Ethereum address of the given key.
func (k *Key) Address() ethcommon.Address {
	return k.address
}
