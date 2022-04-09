package jwt

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
)

func SignEdDsa(payload, publicKey, privateKey, didKeyUrl string) (signedJwt string, err error) {

	if !json.Valid([]byte(payload)) {
		return "", errors.New("Payload must be valid JSON")
	}

	jwk := Jwk{
		Kty: "OKP",
		Crv: "Ed25519",
		X:   publicKey,
		Kid: didKeyUrl,
	}
	edDsaHeader := EdDsaHeader{
		Algorithm: "EdDSA",
		Type:      "JWT",
		JWK:       jwk,
	}
	hdr, err := json.Marshal(edDsaHeader)
	if err != nil {
		return "", err
	}

	toBeSigned := b64Encode(hdr) + "." + b64Encode([]byte(payload))
	signer := b64Decode(privateKey)
	signature := ed25519.Sign(signer, []byte(toBeSigned))

	signedJwt = toBeSigned + "." + b64Encode(signature)
	return signedJwt, nil
}
