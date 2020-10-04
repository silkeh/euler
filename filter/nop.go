package filter

// nop defines a filter that does no operation.
type nop struct{}

// NewNOP returns a filter that pipes the input channel to the output channel without any operation.
func NewNOP() Filter {
	return &nop{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (n *nop) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		out <- v
	}
}
