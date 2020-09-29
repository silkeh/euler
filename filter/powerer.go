package filter

import "math"

// Powerer filters even numbers.
type Powerer struct {
	p float64
}

// NewPowerer creates a even number Filter.
func NewPowerer(p int) Filter {
	return &Powerer{p: float64(p)}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (p *Powerer) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		out <- int(math.Pow(float64(v), p.p))
	}
}
