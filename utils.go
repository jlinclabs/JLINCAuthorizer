package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"

	"github.com/shengdoushi/base58"
)

func zeroPrefixed(c []byte, n int) bool {
	prefixed := true
	for i := 0; i < n; i++ {
		if c[i] != 0 {
			prefixed = false
		}
	}
	return prefixed
}

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

func b58Decode(s string) []byte {
	decoded, _ := base58.Decode(s, base58.BitcoinAlphabet)
	return decoded
}

func b58Encode(h []byte) string {
	return base58.Encode(h, base58.BitcoinAlphabet)
}

func b58tob64(s string) string {
	return b64Encode(b58Decode(s))
}
