package main

import (
	"fmt"
	"time"

	"github.com/SSSaaS/sssa-golang"
)

func SSSaaS_CreateShares(data []byte, n1 int, k1 int) []string {
	shares := make([]string, len(data))
	shares, err := sssa.Create(k1, n1, string(data[:]))
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return shares
}

func SSSaaS_CombineShares(shares []string) string {
	var s string
	s, err := sssa.Combine(shares)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	return s
}

func SSSaaS_speedTest(data [][]byte, n1 int, k1 int) {

	shares := make([][]string, len(data))
	start := time.Now()
	for i := 0; i < len(data); i++ {
		shares[i], _ = sssa.Create(n1, k1, string(data[i][:]))
	}
	elapsed := time.Since(start)
	fmt.Printf("Creating %d shares took: %s\n", len(data)*n1, elapsed)

	start = time.Now()

	for i := 0; i < len(data); i++ {
		_, _ = sssa.Combine(shares[i])
	}
	elapsed = time.Since(start)
	fmt.Printf("Combining %d shares for %d secrets took: %s\n", len(data)*n1, len(data), elapsed)
}
func test_SSSaaS(data [][]byte, n1 int, k1 int) (timer time.Duration) {

	fmt.Printf("SSSaaS test:\n")

	start := time.Now()
	SSSaaS_speedTest(data, n1, k1)
	elapsed := time.Since(start)

	fmt.Printf("SSSaaS took: %s\n", elapsed)
	return elapsed
}
