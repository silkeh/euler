package filter

import (
	"math"
)

// divisorCount counts the number of divisors of received values.
type divisorCount struct{}

// NewDivisorCount returns a Filter that counts the number of divisors
func NewDivisorCount() Filter {
	return &divisorCount{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (dc *divisorCount) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		out <- dc.countDivisors(v)
	}
}

// countDivisors of the input n and return as integer.
func (dc *divisorCount) countDivisors(n int) int {
	if n == 1 {
		return 1
	}
	nSqrt := math.Sqrt(float64(n))
	nSqrtRound := math.Round(nSqrt)
	count := 2
	// if the square root is an integer, add to count.
	if nSqrtRound == nSqrt {
		count++
	}
	iMax := int(nSqrtRound)
	for i := 2; i < iMax; i++ {
		if n%i == 0 {
			count += 2
		}
	}
	return count
}
