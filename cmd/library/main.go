package main

import "C"

import (
	"math/rand"
	"time"

	"github.com/trevor403/random/pkg/linear"
)

// global rand
var rng *rand.Rand

func init() {
	seed := time.Now().Unix()

	lcg := linear.NewPcg32(seed)
	rng = rand.New(lcg)
}

//export GetRandomFromGo
func GetRandomFromGo() C.ulonglong {
	val := rng.Uint64()
	return C.ulonglong(val)
}

//export GetRandomBytesFromGo
func GetRandomBytesFromGo(buf []byte) {
	rng.Read(buf)
}

// stub
func main() {}
