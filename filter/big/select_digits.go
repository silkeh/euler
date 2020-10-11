package big

import (
	"math/big"
)

// selectDigits from a string.
type selectDigits struct {
	n      int
	digits string
}

// NewSelectDigits returns a selectDigits Filter
// Reads a value of n digits from the string digits
func NewSelectDigits(digits string, n int) Filter {
	return &selectDigits{n: n, digits: digits}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (s *selectDigits) Run(in <-chan *big.Int, out chan<- *big.Int) {
	defer close(out)

	for vBig := range in {
		v := int(vBig.Int64())
		number := new(big.Int)
		number.SetString(s.digits[s.n*v:s.n*(v+1)], 10)
		out <- number
	}
}
