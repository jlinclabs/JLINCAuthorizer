package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gowebpki/jcs"
	"github.com/jlinclabs/go-jwt"
)

type Proof struct {
	Type               string `json:"type"`
	PublicKeyBase64    string `json:"publicKeyBase64"`
	Created            string `json:"created"`
	ProofPurpose       string `json:"proofPurpose"`
	VerificationMethod string `json:"verificationMethod"`
	Jws                string `json:"jws"`
}

func makeProof(zcapJson []byte, data DidData) (proof Proof, err error) {
	canonical, err := jcs.Transform(zcapJson)
	if err != nil {
		return Proof{}, err
	}
	signedJws, err := jwt.SignEdDsa(string(canonical), data.SigningPublicKey, data.SigningPrivateKey, data.DID+"#signing")
	if err != nil {
		return Proof{}, err
	}
	parts := strings.Split(signedJws, ".")
	jws := parts[0] + ".." + parts[2]

	p := Proof{
		Type:               "Ed25519Signature2018",
		PublicKeyBase64:    data.SigningPublicKey,
		Created:            fmt.Sprintf("%s", time.Now().UTC().Format(ISOStringMillisec)),
		ProofPurpose:       "authorization",
		VerificationMethod: data.DID + "#signing",
		Jws:                jws,
	}

	return p, nil
}
