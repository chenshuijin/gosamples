package main

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func NewSecp256Key() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}
