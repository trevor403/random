package main

// Permutation congruential generators

import (
	"math/rand"
	"os"
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

	buf := make([]byte, 1<<16)
	for {
		rng.Read(buf)
		os.Stdout.Write(buf)
	}
}
