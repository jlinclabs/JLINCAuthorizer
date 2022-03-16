package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"

	"golang.org/x/crypto/nacl/box"
)

type Entity struct {
	SigningPublicKey     string
	SigningPrivateKey    string
	EncryptingPublicKey  string
	EncryptingPrivateKey string
	RegistrationSecret   string
}

func createEntity() ([]byte, error) {
	signPubkey, signPrivkey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}

	encPubkey, encPrivkey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	regsecret := make([]byte, 32)
	_, err = rand.Read(regsecret)

	entity, err := json.Marshal(
		Entity{
			SigningPublicKey:     b64Encode(signPubkey),
			SigningPrivateKey:    b64Encode(signPrivkey),
			EncryptingPublicKey:  b64Encode(encPubkey[:]), //[:] to convert *[32]byte to []byte
			EncryptingPrivateKey: b64Encode(encPrivkey[:]),
			RegistrationSecret:   hex.EncodeToString(regsecret)})
	if err != nil {
		return nil, err
	}

	return entity, nil
}
