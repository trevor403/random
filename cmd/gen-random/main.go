package main

// linear congruential generators

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"log"

	"github.com/trevor403/random/pkg/linear"
)

func getRandRange(min, max int, rng *rand.Rand) int {
	base := rng.Int() % max // enforce max
	adjusted := base + min  // enforce min

	return adjusted
}

func main() {
	// use time as seed
	seed := big.NewInt(time.Now().Unix())

	lcg := linear.NewCongruentialGenerator(seed)
	rng := rand.New(lcg)

	for i := 0; i < 20; i++ {
		val := rng.Uint64()
		log.Printf("Random number (64bits):\t%d", val)
	}

	fmt.Println("") // spacer

	for i := 0; i < 20; i++ {
		val := getRandRange(1, 10, rng)
		log.Printf("Random number (1-10):\t%d", val)
	}
}
