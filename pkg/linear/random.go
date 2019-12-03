package linear

// Algorithm is a Permutation Congruential Generators (PCG)
// |- Implementation is PCG-XSH-RR built on top of a 64-bit LCG
//   |- Page 37 http://www.pcg-random.org/pdf/toms-oneill-pcg-family.pdf
// |- Multiplier from L'Ecuye tables
//   |- Page 10 https://www.ams.org/journals/mcom/1999-68-225/S0025-5718-99-00996-5/S0025-5718-99-00996-5.pdf
// |- Increment from the PCG reference implementation

const (
	multiplier uint64 = 2862933555777941757
	increment  uint64 = 1442695040888963407

	mask uint64 = (1 << 64) - 1
)

type Pcg32 struct {
	state uint64
}

func NewPcg32(seed int64) *Pcg32 {
	pcg := &Pcg32{}
	pcg.Seed(seed)

	return pcg
}

func (pcg *Pcg32) Seed(s int64) {
	seed := uint64(s)
	pcg.state = increment + (seed & mask)
	pcg.Step()
}

func (pcg *Pcg32) Step() {
	pcg.state = (pcg.state * multiplier) + (increment & mask)
}

func (pcg *Pcg32) Next() uint32 {
	pcg.Step()
	state := pcg.state
	rot := uint32(state >> 59)
	shift := uint32(((state >> 18) ^ state) >> 27)
	output := (shift >> rot) | (shift << (32 - rot))
	return output
}

func (pcg *Pcg32) Int63() int64 {
	return int64(pcg.Uint64() >> 1)
}

func (pcg *Pcg32) Uint64() uint64 {
	high := pcg.Next()
	low := pcg.Next()

	return uint64(high)<<32 | uint64(low)
}
