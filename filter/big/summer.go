package big

import "math/big"

// summer is a summing filter.
type summer struct {
	n int
}

// NewSummer returns a summing Filter.
// Sums n received values together and sends the result
func NewSummer(n int) Filter {
	return &summer{n: n}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (s *summer) Run(in <-chan *big.Int, out chan<- *big.Int) {
	defer close(out)

	sum := big.NewInt(0)
	for i := 0; i < s.n; i++ {
		v, ok := <-in
		if !ok {
			return
		}
		sum.Add(sum, v)
	}
	out <- sum
}
