package jwt

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"strings"
)

func VerifyEdDsa(signedJwt string, publicKey []byte) (payloadJson string, err error) {

	sliced := strings.Split(signedJwt, ".")
	if len(sliced) != 3 {
		return "", errors.New("Not a valid JWT format")
	}

	var edDsaHeader EdDsaHeader
	err = json.Unmarshal(b64Decode(sliced[0]), &edDsaHeader)
	if edDsaHeader.Algorithm != "EdDSA" {
		return "", errors.New("EdDSA algorithm not specified in header")
	}
	// if publicKey wasn't supplied in the arguments, we look for it in the header jwk
	if publicKey == nil {
		publicKey = b64Decode(edDsaHeader.JWK.X)
	}
	// if it's not there either we fail
	if len(publicKey) != ed25519.PublicKeySize {
		return "", errors.New("No valid public key found")
	}

	wasSigned := sliced[0] + "." + sliced[1]
	if verified := ed25519.Verify(publicKey, []byte(wasSigned), b64Decode(sliced[2])); !verified {
		return "", errors.New("Signature not verified")
	}

	return string(b64Decode(sliced[1])), nil
}
