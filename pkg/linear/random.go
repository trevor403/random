package linear

import (
	"math/big"
)

var streamer, _ = big.NewInt(0).SetString("117397592171526113268558934119004209487", 10)

// Ferrier's Prime is really interested
// it is the larges prime ever computed without electronic assistance
const ferrierPrimeString = "20988936657440586486151264256610222593863921"

// 144 bit number can not use a built-in data-type
var ferrierPrime, _ = new(big.Int).SetString(ferrierPrimeString, 10)

type CongruentialGenerator struct {
	State, Stream, Multiplier *big.Int
}

func NewCongruentialGenerator(seed *big.Int) *CongruentialGenerator {
	return &CongruentialGenerator{seed, streamer, ferrierPrime}
}

func (gen *CongruentialGenerator) Seed(s int64) {
	high := new(big.Int).Lsh(big.NewInt(s), 64)
	low := big.NewInt(s)
	gen.State = new(big.Int).Add(high, low)
}

func (gen *CongruentialGenerator) Next() *big.Int {
	gen.State = new(big.Int).Mul(gen.State, gen.Multiplier)
	gen.State = new(big.Int).Add(gen.State, gen.Stream)
	return gen.State
}

func (gen *CongruentialGenerator) Int63() int64 {
	n := gen.Next()
	high := new(big.Int).Rsh(n, 64).Int64()
	low := n.Int64()
	return int64((high ^ low) >> 1)
}
