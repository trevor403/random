package main

// Permutation congruential generators

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/trevor403/random/pkg/linear"
)

func getRandRange(min, max int, rng *rand.Rand) int {
	base := rng.Int() % max // enforce max
	adjusted := base + min  // enforce min

	return adjusted
}

func main() {
	// use time as seed
	seed := time.Now().Unix()

	pcg := linear.NewPcg32(seed)
	rng := rand.New(pcg)

	fmt.Println("==== Get Random numbers (64bits):")
	for i := 0; i < 10; i++ {
		val := rng.Uint64()
		fmt.Printf("[ %d ]\t%d\n", i, val)
	}

	fmt.Println("==== Get Random number (1-10):")
	for i := 0; i < 10; i++ {
		val := getRandRange(1, 10, rng)
		fmt.Printf("[ %d ]\t%d\n", i, val)
	}
}
