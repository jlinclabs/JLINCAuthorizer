package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
)

func SignHmac(payload string, secret string) (signedJwt string, err error) {

	if !json.Valid([]byte(payload)) {
		return "", errors.New("Payload must be valid JSON")
	}

	hmacHeader := Header{
		Algorithm: "HS256",
		Type:      "JWT",
	}
	hdr, err := json.Marshal(hmacHeader)
	if err != nil {
		return "", err
	}

	toBeSigned := b64Encode(hdr) + "." + b64Encode([]byte(payload))
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(toBeSigned))

	signedJwt = toBeSigned + "." + b64Encode(mac.Sum(nil))

	return signedJwt, nil
}
