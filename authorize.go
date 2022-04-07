package main

import (
	"fmt"
	"log"
	"runtime"
)

///////// Setup
const didServerUrl = "http://localhost:5001"

//const didServerUrl = "https://didserver.jlinc.org"
const zcapContext = "https://jlinc.org/zcap/v1"

type TargetData struct {
	TargetUri string
}
type DidData struct {
	SigningPublicKey     string `json:"signingPublicKey"`
	SigningPrivateKey    string `json:"signingPrivateKey"`
	EncryptingPublicKey  string `json:"encryptingPublicKey"`
	EncryptingPrivateKey string `json:"encryptingPrivateKey"`
	RegistrationSecret   string `json:"registrationSecret"`
	DID                  string `json:"did"`
}

// Available services
var targets = map[string]TargetData{
	"badbirders": TargetData{TargetUri: "https://bad-birders.jlinc.io/zcap-login?zcap="},
	"catwalkers": TargetData{TargetUri: "https://cat-walkers.jlinc.io/zcap-login?zcap="},
	"dopedogs":   TargetData{TargetUri: "https://dope-dogs.jlinc.io/zcap-login?zcap="},
}

/////////

func Authorize(service string) (target string, err error) {
	opsys := runtime.GOOS
	if opsys == "darwin" {
		target, err := authorizeDarwin(service)
		return target, err
	} else {
		return opsys, fmt.Errorf("The %s operating system is not yet supported.", opsys)
	}
}

func authorizeDarwin(service string) (target string, err error) {
	zcap, err := MacStoredData(service)
	log.Printf("zcap: %v", zcap)
	target = targets[service].TargetUri
	return target, err
}
