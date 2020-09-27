package filter

// Fractor filters numbers that are a factor of another number.
type Fractor struct {
	n int
}

// NewFractor creates a factor-finding Filter.
// This filters numbers of which `n` is a factor.
func NewFractor(n int) Filter {
	return &Fractor{n: n}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (f *Fractor) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		if v%f.n == 0 {
			out <- v
		}
	}
}
