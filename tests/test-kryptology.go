package main

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/coinbase/kryptology/pkg/core/curves"

	kryptology "github.com/coinbase/kryptology/pkg/sharing"
)

func Kryptology_speedTest(data [][]byte, n1 uint32, k1 uint32) {

	shares := make([][]*kryptology.ShamirShare, len(data))
	start := time.Now()

	shamir, err := kryptology.NewShamir(k1, n1, curves.ED25519())
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	for i := 0; i < len(data); i++ {
		secret := curves.ED25519().NewScalar().Hash(data[i][:])
		var err error
		shares[i], err = shamir.Split(secret, rand.Reader)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Creating %d shares took: %s\n", len(data)*int(n1), elapsed)

	start = time.Now()

	for i := 0; i < len(data); i++ {
		_, _ = shamir.Combine(shares[i]...)

	}
	elapsed = time.Since(start)
	fmt.Printf("Combining %d shares for %d secrets took: %s\n", len(data)*int(n1), len(data), elapsed)
}
func test_Kryptology(data [][]byte, n1 int, k1 int) (timer time.Duration) {

	fmt.Printf("Kryptology test:\n")

	start := time.Now()
	Kryptology_speedTest(data, uint32(n1), uint32(k1))
	elapsed := time.Since(start)

	fmt.Printf("Kryptology took: %s\n", elapsed)
	return elapsed
}
