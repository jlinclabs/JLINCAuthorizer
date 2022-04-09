package jwt

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
)

func getHash(j string) []byte {
	h := sha256.New()
	h.Write([]byte(j))
	return h.Sum(nil)
}

func getByteHash(j []byte) []byte {
	h := sha256.New()
	h.Write(j)
	return h.Sum(nil)
}

func b64Decode(s string) []byte {
	decoded, _ := base64.RawURLEncoding.DecodeString(s)
	return decoded
}

func b64Encode(h []byte) string {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.RawURLEncoding, &buf)
	encoder.Write(h)
	encoder.Close()
	return buf.String()
}
