# Go JWT

The [JLINC protocol](https://protocol.jlinc.org/) uses JSON WebTokens of type [JWS](https://www.rfc-editor.org/rfc/rfc7515) as a compact standards-based way to transmit SISAs and SISA Events between servers.

This package supports HS256 and EdDSA/Ed25519 algorithms only and tries to be a simple as possible to enable comprehensive security audit.

## Expected Usage
## Installation and  Usage

```golang
import "github.com/jlinclabs/go-jwt"
// then install it into your app
go mod tidy
```
### Sign with HMAC/SHA256
```golang
/*
PayloadJSON must be a JSON string. For compatibility it should be created
without unnecessary spaces, as is done by encoding/json.Marshal or
javascript's JSON.stringify().

SecretString can be any string. For security it should be a long random string.
*/

jwt.SignHmac(PayloadJSON, SecretString string) (jsonWebtoken string, err error)
```


### Sign with EdDSA/Ed25519
```golang
/*
PayloadJSON must be a JSON string. For compatibility it should be created
without unnecessary spaces, as is done by encoding/json.Marshal or
javascript's JSON.stringify().

PublicKey and SecretKey are an Ed25519 keypair as created for example by
crypto/ed25519.GenerateKey or libsodium's crypto_sign_keypair.

The DIDKeyUrl argument is expected to be a DID url of the form {DID}#signing 
and is placed in the JWT header under jwk.kid, i.e. a JSON-WebKey key-ID
(https://www.rfc-editor.org/rfc/rfc7515#section-4.1.4).
*/

jwt.SignEdDsa(PayloadJSON, publicKey, privateKey, didKeyUrl string) (jsonWebtoken string, err error)
```

### Verify a HMAC/SHA256 signed JWT
```golang
/*
Presented with a valid HMAC/SHA256 signed JWT and the secret it was signed with,
returns the JSON string payload.
*/

jwt.VerifyHmac(signedJwt string, secret []byte) (payloadJson string, err error)
```


### Verify an EdDSA/Ed25519 signed JWT
```golang
/*
JWTs created with this package's SignEdDsa method will contain the public key that
validates the signature in the JWT's header under the jwk.x key.
See https://tools.ietf.org/html/rfc8037#section-2.

If the public key is not available that way, perhaps because the JWT was created
by a different application, then it must be supplied by the second argument.
Otherwise the publicKey argument must be nil. 
If the public key is present in both places, the supplied argument will be used.

On success returns the JSON string payload.
*/

jwt.VerifyEdDsa(signedJwt string, publicKey []byte) (payloadJson string, err error)
```
