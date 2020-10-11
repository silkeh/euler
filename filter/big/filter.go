package big

import "math/big"

// Filter defines interface for filters that process *big.Int
type Filter interface {
	// Run against an input and output channel until the input channel is closed.
	// The output channel is closed on return.
	Run(in <-chan *big.Int, out chan<- *big.Int)
}
