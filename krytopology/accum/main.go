package main

import (
	"fmt"
	"os"

	"github.com/coinbase/kryptology/pkg/accumulator"
	"github.com/coinbase/kryptology/pkg/core/curves"
)

func main() {

	msg := "Hello"
	argCount := len(os.Args[1:])

	if argCount > 0 {
		msg = os.Args[1]
	}

	curve := curves.BLS12381(&curves.PointBls12381G1{})
	var seed [32]byte
	key, _ := new(accumulator.SecretKey).New(curve, seed[:])

	acc, _ := new(accumulator.Accumulator).New(curve)

	element := curve.Scalar.Hash([]byte(msg))

	rtn, _ := acc.MarshalBinary()

	fmt.Printf("Message to prove: %s", msg)
	fmt.Printf("\nNonce: %x", seed)
	datakey, _ := key.MarshalBinary()
	fmt.Printf("\nKey: %x", datakey)
	fmt.Printf("\nBefore add: %x\n", rtn)

	_, _ = acc.Add(key, element)

	rtn, _ = acc.MarshalBinary()

	fmt.Printf("\nAfter add: %x\n", rtn)

	_, _ = acc.Remove(key, element)
	rtn, _ = acc.MarshalBinary()

	fmt.Printf("\nAfter remove: %x\n", rtn)

}
