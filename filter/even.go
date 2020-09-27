package filter

// Even filters even numbers.
type Even struct{}

// NewEven creates a even number Filter.
func NewEven() Filter {
	return &Even{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (e *Even) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		if v%2 == 0 {
			out <- v
		}
	}
}
