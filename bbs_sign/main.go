package main

import (
	"fmt"

	"os"

	"github.com/coinbase/kryptology/pkg/core/curves"
	"github.com/coinbase/kryptology/pkg/signatures/bbs"
)

func main() {

	curve := curves.BLS12381(&curves.PointBls12381G2{})
	msg1 := "Hello"
	msg2 := "Hello"
	msg3 := "Hello"

	argCount := len(os.Args[1:])

	if argCount > 0 {
		msg1 = os.Args[1]
	}
	if argCount > 1 {
		msg2 = os.Args[2]
	}
	if argCount > 2 {
		msg3 = os.Args[3]
	}

	var msgs []curves.Scalar
	msgs = append(msgs, curve.Scalar.Hash([]byte(msg1)))
	msgs = append(msgs, curve.Scalar.Hash([]byte(msg2)))
	msgs = append(msgs, curve.Scalar.Hash([]byte(msg3)))

	pk, sk, _ := bbs.NewKeys(curve)

	generators := new(bbs.MessageGenerators).Init(pk, len(msgs))

	sig, _ := sk.Sign(generators, msgs)

	sigdata, _ := sig.MarshalBinary()

	err := pk.Verify(sig, generators, msgs)

	fmt.Printf("Message 1: %s, Message 2: %s, Message 3: %s\n", msg1, msg2, msg3)
	fmt.Printf("Private key: %s\nPublic Key %s\n", sk, pk)
	fmt.Printf("Signature: %x\n", sigdata)

	if err == nil {
		fmt.Println("BBS Signature Proven")
	} else {
		fmt.Println("BBS Signature Not Proven")
	}

}
