package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dsprenkels/sss-go"
)

type Share struct {
	data [][]byte
}

func sprenekls_createShares(str []byte) [][]byte {
	shares, err := sss.CreateShares(str, 5, 4)
	if err != nil {
		log.Fatalln(err)
	}
	return shares

}
func sprenekls(data [][]byte, n1 int, k1 int) []Share {

	numGen := len(data)
	shares := make([]Share, numGen)
	start := time.Now()
	for i := 0; i < numGen; i++ {
		share, err := sss.CreateShares(data[i], n1, k1)
		if err != nil {
			log.Fatalln(err)
		}
		shares[i].data = share
	}
	elapsed := time.Since(start)

	fmt.Printf("Creating %d shares took: %s\n", numGen*5, elapsed)

	start = time.Now()
	for i := 0; i < numGen; i++ {
		_, err := sss.CombineShares(shares[i].data)
		if err != nil {
			log.Fatalln(err)
		}

	}
	elapsed = time.Since(start)

	fmt.Printf("Combining %d shares for %d secrets took: %s\n", numGen*5, numGen, elapsed)

	return shares
}
func test_sprenkels(data [][]byte, n1 int, k1 int) (timer time.Duration) {

	fmt.Printf("Dsprenkels test:\n")

	start := time.Now()
	sprenekls(data, n1, k1)
	elapsed := time.Since(start)

	fmt.Printf("Dsprenkels took: %s", elapsed)
	return elapsed
}
