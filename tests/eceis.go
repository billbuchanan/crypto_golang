package main

import (
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/encrypt/ecies"
	"go.dedis.ch/kyber/v3/group/edwards25519"

	"go.dedis.ch/kyber/v3/util/random"
)

var rng = random.New()

type KeyPair struct {
	publicKey  string
	privateKey string
}

func eceis(str []byte, private kyber.Scalar) (KeyPair, string, kyber.Scalar) {

	suite := edwards25519.NewBlakeSHA256Ed25519()

	public := suite.Point().Mul(private, nil)

	ciphertext, _ := ecies.Encrypt(suite, public, str, suite.Hash)

	keys := KeyPair{publicKey: public.String(), privateKey: private.String()}
	return keys, string(ciphertext[:]), private
}
