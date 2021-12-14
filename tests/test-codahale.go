package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codahale/sss"
)

type CodahaleShare struct {
	data map[byte][]byte
}

func codahale(data [][]byte, n1 int, k1 int) {

	n := byte(n1)
	k := byte(k1)

	if k1 > n1 {
		fmt.Printf("Cannot do this, as k greater than n")
		os.Exit(0)
	}
	shares := make([]CodahaleShare, len(data))
	start := time.Now()

	for i := 0; i < len(data); i++ {
		key := data[i]
		share, _ := sss.Split(n, k, key)
		shares[i].data = share
	}
	elapsed := time.Since(start)
	fmt.Printf("Creating %d shares took: %s\n", len(data)*n1, elapsed)

	//key := make([]byte, 16)

	//_, _ = rand.Read(key)

	/*subset := make(map[byte][]byte, k)
	for x, y := range shares {

		fmt.Printf("Share:\t%d\t%s\n", x, hex.EncodeToString(y))
		subset[x] = y
		if len(subset) == int(k) {
			break
		}
	}*/
	start = time.Now()
	for i := 0; i < len(data); i++ {
		_ = sss.Combine(shares[i].data)

	}
	elapsed = time.Since(start)
	fmt.Printf("Combining %d shares for %d secrets took: %s\n", len(data)*n1, len(data), elapsed)

}

func test_codahale(data [][]byte, n1 int, k1 int) (timer time.Duration) {

	fmt.Printf("Codahale test	:\n")

	start := time.Now()
	codahale(data, n1, k1)
	elapsed := time.Since(start)

	fmt.Printf("Codahale took: %s\n", elapsed)
	return elapsed
}
