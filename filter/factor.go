package filter

// Factor filters numbers that are a factor of another number.
type Factor struct {
	n int
}

// NewFactor creates a factor-finding Filter.
// This filters numbers that are a factor of `n`
func NewFactor(n int) Filter {
	return &Factor{n: n}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (f *Factor) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		if f.n%v == 0 {
			out <- v
		}
	}
}
