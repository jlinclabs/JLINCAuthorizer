package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"strings"
)

func VerifyHmac(signedJwt string, secret []byte) (payloadJson string, err error) {

	sliced := strings.Split(signedJwt, ".")
	if len(sliced) != 3 {
		return "", errors.New("Not a valid JWT format")
	}

	var hmacHeader Header
	err = json.Unmarshal(b64Decode(sliced[0]), &hmacHeader)
	if hmacHeader.Algorithm != "HS256" {
		return "", errors.New("HS256 algorithm not specified in header")
	}

	wasSigned := sliced[0] + "." + sliced[1]
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(wasSigned))
	sigShouldBe := b64Encode(mac.Sum(nil))
	if sigShouldBe != sliced[2] {
		return "", errors.New("Signature not verified")
	}

	return string(b64Decode(sliced[1])), nil
}
