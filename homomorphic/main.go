package main

import (
	"fmt"

	"github.com/coinbase/kryptology/pkg/core/curves"
	"github.com/coinbase/kryptology/pkg/verenc/elgamal"
)

func main() {

	curve := curves.ED25519()
	pk, sk, _ := elgamal.NewKeys(curve)

	x := curve.Scalar.New(10)
	//	Y := curve.Point.Generator().Mul(x)

	res, _ := pk.HomomorphicEncrypt(x)
	fmt.Printf("%x", res)

	dec, _ := res.Decrypt(sk)
	fmt.Printf("%x", dec)

}
