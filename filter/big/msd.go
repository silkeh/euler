package big

import (
	"math/big"
)

// msd returns the n most significant digits.
type msd struct {
	n int
}

// NewMSD returns a most significant digits filter.
// n defines the number of digits that pass through this filter.
func NewMSD(n int) Filter {
	return &msd{n: n}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (m *msd) Run(in <-chan *big.Int, out chan<- *big.Int) {
	defer close(out)

	for v := range in {
		vString := v.String()
		bigOutput, ok := new(big.Int).SetString(vString[0:m.n], 10)
		if !ok {
			return
		}
		out <- bigOutput
	}
}
