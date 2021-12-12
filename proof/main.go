package main

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/coinbase/kryptology/pkg/core"
	"github.com/coinbase/kryptology/pkg/core/curves"
)

func main() {

	mymsg := "Hello"

	argCount := len(os.Args[1:])

	if argCount > 0 {
		mymsg = os.Args[1]
	}

	curve := curves.ED25519()
	// priv =x, pub=xG

	x := curve.Scalar.Random(crand.Reader)
	Y := curve.Point.Generator().Mul(x)

	r := curve.Scalar.Random(crand.Reader)

	C1 := curve.Point.Generator().Mul(r)

	M := curve.Point.Hash([]byte(mymsg))

	C2 := Y.Mul(r).Add(M)

	//	Bob will derive key for AEAD

	t := Y.Mul(r)
	aeadKey, _ := core.FiatShamir(new(big.Int).SetBytes(t.ToAffineCompressed()))
	block, _ := aes.NewCipher([]byte(aeadKey))

	aesGcm, _ := cipher.NewGCM(block)
	// add = C1 || C2
	aad := C1.ToAffineUncompressed()
	aad = append(aad, C2.ToAffineUncompressed()...)
	var nonce [12]byte
	_, _ = crand.Read(nonce[:])
	aead := aesGcm.Seal(nil, nonce[:], []byte(mymsg), aad)

	////////////////////////////////////////////////////////////////////////////
	// Now Alice will decrypt. Using aead, C1, C2, nonce, and private key (x) //
	////////////////////////////////////////////////////////////////////////////

	t = C1.Mul(x)
	aeadKey, _ = core.FiatShamir(new(big.Int).SetBytes(t.ToAffineCompressed()))
	block, _ = aes.NewCipher([]byte(aeadKey))
	aesGcm, _ = cipher.NewGCM(block)

	aad1 := C1.ToAffineUncompressed()
	aad1 = append(aad1, C2.ToAffineUncompressed()...)
	msg, _ := aesGcm.Open(nil, nonce[:], aead, aad1)

	fmt.Printf("Orginal message:\t%s\n", mymsg)
	fmt.Printf("\nAlice's private key:\t%x\n", x.Bytes())
	fmt.Printf("Alice's public key:\t%x\n", Y.ToAffineCompressed())
	fmt.Printf("\nC1:\t%x\n", C1.ToAffineCompressed())
	fmt.Printf("C2:\t%x\n", C2.ToAffineCompressed())
	fmt.Printf("Message to point:\t%x\n", M.ToAffineCompressed())
	fmt.Printf("AAD:\t%x\n", aad1)
	fmt.Printf("Nonce:\t\t%x\n", nonce)

	fmt.Printf("Encrypted:\t%x\n", aead)

	fmt.Printf("\n === Alice receives AEAD, Nonce, C1 and C2 and recovers message ===")
	fmt.Printf("\nRecovered message:\t%s\n", msg)

}
