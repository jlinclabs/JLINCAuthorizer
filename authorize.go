package main

import (
	"fmt"
	"runtime"

	"github.com/google/uuid"
)

///////// Setup
const didServerUrl = "http://localhost:5001"

//const didServerUrl = "https://didserver.jlinc.org"
const zcapContext = "https://jlinc.org/zcap/v1"

const ISOStringMillisec = "2006-01-02T15:04:05.999Z"

type TargetData struct {
	TargetUri string
}
type DidData struct {
	ParentCapability     uuid.UUID `json:"parentCapability"`
	SigningPublicKey     string    `json:"signingPublicKey"`
	SigningPrivateKey    string    `json:"signingPrivateKey"`
	EncryptingPublicKey  string    `json:"encryptingPublicKey"`
	EncryptingPrivateKey string    `json:"encryptingPrivateKey"`
	RegistrationSecret   string    `json:"registrationSecret"`
	DID                  string    `json:"did"`
}

// Available services
var targets = map[string]TargetData{
	"badbirders": TargetData{TargetUri: "https://bad-birders.jlinc.io/login?zcap="},
	"catwalkers": TargetData{TargetUri: "https://cat-walkers.jlinc.io/login?zcap="},
	"dopedogs":   TargetData{TargetUri: "https://dope-dogs.jlinc.io/login?zcap="},
}

/////////

func Authorize(service string) (target string, err error) {
	opsys := runtime.GOOS
	if opsys == "darwin" {
		target, err := authorizeDarwin(service)
		if err != nil {
			return "", err
		}
		return target, nil
	} else {
		return opsys, fmt.Errorf("The %s operating system is not yet supported.", opsys)
	}
}

func authorizeDarwin(service string) (target string, err error) {
	zcap, err := MacStoredData(service)
	if err != nil {
		return "", err
	}
	target = targets[service].TargetUri + zcap
	return target, nil
}
