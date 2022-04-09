package jwt

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}
type EdDsaHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
	JWK       Jwk    `json:"jwk"`
}
type Jwk struct {
	Kty string `json:"kty"`
	Crv string `json:"crv"`
	X   string `json:"x"`
	Kid string `json:"kid"`
}

type Jwt struct {
	Head      Header
	Payload   map[string]interface{}
	Signature []byte
}
type EdDsaJwt struct {
	Head      EdDsaHeader
	Payload   map[string]interface{}
	Signature []byte
}
