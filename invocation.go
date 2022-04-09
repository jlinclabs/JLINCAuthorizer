package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jlinclabs/go-jwt"
)

type Invoked struct {
	ZcapContext        string    `json:"@context"`
	ParentCapabilityID uuid.UUID `json:"parentCapabilityId"`
	InvocationId       uuid.UUID `json:"invocationId"`
	Action             string    `json:"action"`
	Created            string    `json:"created"`
	InvokerDID         string    `json:"invokerDid"`
}

type Invocation struct {
	Invoked
	InvokedProof Proof `json:"proof"`
}

func MakeInvocation(data DidData) (zcap string, err error) {
	invoked := Invoked{
		ZcapContext:        zcapContext,
		ParentCapabilityID: data.ParentCapability,
		InvocationId:       uuid.New(),
		Action:             "authorization",
		Created:            fmt.Sprintf("%s", time.Now().UTC().Format(ISOStringMillisec)),
		InvokerDID:         data.DID,
	}
	invokedJson, err := json.Marshal(invoked)
	if err != nil {
		return "", err
	}
	proof, err := makeProof(invokedJson, data)

	invocation := Invocation{
		Invoked:      invoked,
		InvokedProof: proof,
	}
	invocationJson, err := json.Marshal(invocation)
	if err != nil {
		return "", err
	}
	zcap, err = jwt.SignEdDsa(string(invocationJson), data.SigningPublicKey, data.SigningPrivateKey, data.DID+"#signing")
	return zcap, nil
}
